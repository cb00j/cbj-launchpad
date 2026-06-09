// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {ISalesFactory} from "../interfaces/ISalesFactory.sol";
import {IAllocationStaking} from "../interfaces/IAllocationStaking.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

/**
 *
 * initialize the sale -> registration -> sale start -> sale end -> vesting -> sale owner withdraw earnings
 *
 * @title
 * @author
 * @notice
 */
contract CBJSale is ReentrancyGuard, Ownable {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;
    using SafeERC20 for IERC20;

    ISalesFactory public factory;
    IAllocationStaking public allocationStaking;

    // sale structure
    struct Sale {
        // token being sold
        IERC20 token;
        // ---status of the sale---
        // is sale created,if the sale is initailized or not
        bool isCreated;
        // are earings withdrawn, if the sale owner has withdrawn the earnings(ETH) or not
        bool earningWithdrawn;
        // is leftover withdrawn,if the leftover tokens has been withdrawn or not
        bool leftoverWithdrawn;
        // have tokens been deposited,the tokens should be deposited in the contract before sale starts
        bool tokensDeposited;
        // ---sale parameters---
        // address of sale owner
        address saleOwner;
        // price of the token quoted in ETH
        uint256 tokenPriceInETH;
        // amount of tokens to sell,unit is wei
        uint256 amountOfTokensToSell;
        // total tokens being sold,,unit is wei
        uint256 totalTokensSold;
        // total ETH raised
        uint256 totalETHRaised;
        // max participation for sigle user
        uint256 maxParticipation;
        // time limit
        // sale start time
        uint256 saleStart;
        // sale end time
        uint256 saleEnd;
        // when tokens can be withdrawn
        uint256 tokensUnlockTime;
    }

    // user's portions and withdraw records
    struct Participation {
        // amount of tokens bought
        uint256 amountBought;
        // amount of ETH paid
        uint256 amountETHPaid;
        // time when user participated
        uint256 timeParticipated;
        // array of bools to indicate if portions are withdrawn
        bool[] isPortionWithdrawn;
    }

    // registration structure
    struct Registration {
        // start time of registration
        uint256 registrationTimeStarts;
        // end time of registration
        uint256 registrationTimeEnds;
        // number of registrants
        uint256 numberOfRegistrants;
    }

    Sale public sale;
    Registration public registration;
    uint256 public numberOfParticipants;

    // Mapping user to his participation
    mapping(address => Participation) public userToParticipation;
    // Mapping if user is registered or not
    mapping(address => bool) public isRegistered;
    // mapping if user is participated or not
    mapping(address => bool) public isParticipated;
    // Times when portions are getting unlocked,e.g. [1717380000, 1719972000] means first portion unlocked time, second portion unlocked time, etc.
    uint256[] public vestingPortionsUnlockTime;
    // Percent of the participation user can withdraw,using with portionVestingPrecision; e.g. [1000, 2000, ...] means first portion can withdraw 10%, second portion can withdraw 20%, The sum of the percent should be 100%.
    uint256[] public vestingPercentPerPortion;
    // Precision for percent for portion vesting,e.g. 10000 means 1000 means above 10%,2000 means 20%, etc.
    uint256 public portionVestingPrecision;
    // Max vesting time shift,e.g. 30 days means the max time shift is 30 days.
    uint256 public maxVestingTimeShift;

    // Restricting calls only to sale owner
    modifier onlySaleOwner() {
        require(msg.sender == sale.saleOwner, "Not sale owner");
        _;
    }

    event TokensSold(address user, uint256 amount);
    event UserRegistered(address user);
    event TokenPriceSet(uint256 newPrice);
    event MaxParticipationSet(uint256 newMaxParticipation);
    event TokensWithdrawn(address user, uint256 amount);
    event SaleCreated(
        address saleOwner,
        uint256 tokenPriceInETH,
        uint256 amountOfTokensToSell,
        uint256 saleEnd
    );
    event StartTimeSet(uint256 startTime);
    event RegistrationTimeSet(
        uint256 registrationTimeStarts,
        uint256 registrationTimeEnds
    );

    // will be constructed by the Salefactory
    constructor(address _admin, address _allocationStaking) Ownable(_admin) {
        require(_admin != address(0), "Invalid admin address");
        require(
            _allocationStaking != address(0),
            "Invalid allocation staking address"
        );
        factory = ISalesFactory(msg.sender);
        allocationStaking = IAllocationStaking(_allocationStaking);
    }

    function setSaleParams(
        address _token,
        address _saleOwner,
        uint256 _tokenPriceInETH,
        uint256 _amountOfTokensToSell,
        uint256 _saleEnd,
        uint256 _tokensUnlockTime,
        uint256 _maxParticipation,
        uint256 _portionVestingPrecision
    ) external onlyOwner {
        require(!sale.isCreated, "Sale already created");
        require(
            _token != address(0) &&
                _saleOwner != address(0) &&
                _tokenPriceInETH != 0 &&
                _amountOfTokensToSell != 0 &&
                _saleEnd > block.timestamp &&
                _tokensUnlockTime > block.timestamp &&
                _maxParticipation > 0,
            "setSaleParams: Bad input"
        );

        require(_portionVestingPrecision > 0, "Should be at least 100");

        sale.token = IERC20(_token);
        sale.saleOwner = _saleOwner;
        sale.tokenPriceInETH = _tokenPriceInETH;
        sale.amountOfTokensToSell = _amountOfTokensToSell;
        sale.saleEnd = _saleEnd;
        sale.tokensUnlockTime = _tokensUnlockTime;
        sale.maxParticipation = _maxParticipation;
        sale.earningWithdrawn = false;
        sale.leftoverWithdrawn = false;
        sale.isCreated = true;

        portionVestingPrecision = _portionVestingPrecision;

        emit SaleCreated(
            sale.saleOwner,
            sale.tokenPriceInETH,
            sale.amountOfTokensToSell,
            sale.saleEnd
        );
    }

    /////////////////////////////////⬇️sale admin management functions⬇️/////////////////////////////////
    function setSaleStart(uint256 _saleStart) external onlyOwner {
        require(sale.isCreated, "Sale not created");
        require(
            _saleStart > block.timestamp && _saleStart < sale.saleEnd,
            "Sale start must be greater than current timestamp and less than sale end"
        );
        sale.saleStart = _saleStart;
        emit StartTimeSet(_saleStart);
    }

    function setVestingParams(
        uint256[] memory _unlockTimes,
        uint256[] memory _percentPerPortion,
        uint256 _maxVestingTimeShift
    ) external onlyOwner {
        require(
            vestingPercentPerPortion.length == 0 &&
                vestingPortionsUnlockTime.length == 0,
            "Vesting params already set"
        );
        require(
            _unlockTimes.length == _percentPerPortion.length,
            "Invalid params"
        );

        require(
            portionVestingPrecision > 0,
            "Safeguard for making sure setSaleParams get first called."
        );

        require(_maxVestingTimeShift <= 30 days, "Maximal shift is 30 days.");

        maxVestingTimeShift = _maxVestingTimeShift;

        uint256 sum;

        for (uint i = 0; i < _unlockTimes.length; i++) {
            vestingPortionsUnlockTime.push(_unlockTimes[i]);
            vestingPercentPerPortion.push(_percentPerPortion[i]);
            sum += _percentPerPortion[i];
        }

        require(sum == portionVestingPrecision, "Percent distribution issue.");
    }

    function shiftVestingPortionsUnlockTime(
        uint256 timeToShift
    ) external onlyOwner {
        require(
            timeToShift > 0 && timeToShift <= maxVestingTimeShift,
            "Shift must be nonzero and smaller than maxVestingTimeShift."
        );

        // Time can be shifted only once.
        maxVestingTimeShift = 0;

        for (uint i = 0; i < vestingPortionsUnlockTime.length; i++) {
            vestingPortionsUnlockTime[i] += timeToShift;
        }
    }

    // @notice     DELAY BINDED Function to retroactively set sale token address, can be called only once,
    //             after initial contract creation has passed. Added as an options for teams which
    //             are not having token at the moment of sale launch.
    function setToken(address _token) external onlyOwner {
        require(address(sale.token) == address(0), "Token already set");
        require(_token != address(0), "Invalid token address");
        sale.token = IERC20(_token);
    }

    // Function to set registration period parameters
    function setRegistrationTime(
        uint256 _registrationTimeStarts,
        uint256 _registrationTimeEnds
    ) external onlyOwner {
        require(sale.isCreated, "Sale not created");
        require(
            registration.registrationTimeStarts == 0,
            "Registration time already set"
        );
        require(
            _registrationTimeStarts > block.timestamp &&
                _registrationTimeEnds > block.timestamp,
            "Registration time start must be greater than current timestamp"
        );
        require(
            _registrationTimeStarts < _registrationTimeEnds,
            "Registration time start must be less than registration time end"
        );
        registration.registrationTimeStarts = _registrationTimeStarts;
        registration.registrationTimeEnds = _registrationTimeEnds;
        emit RegistrationTimeSet(
            _registrationTimeStarts,
            _registrationTimeEnds
        );
    }

    /// @notice     Admin function, to update token price before sale to match the closest $ desired rate.
    /// @dev        This will be updated with an oracle during the sale every N minutes, so the users will always
    ///             pay initialy set $ value of the token. This is to reduce reliance on the ETH volatility.
    function updateTokenPriceInETH(uint256 price) external onlyOwner {
        require(sale.isCreated, "Sale not created");
        require(price > 0, "Token price in ETH must be greater than zero");
        // Allowing oracle to run and change the sale value
        sale.tokenPriceInETH = price;
        emit TokenPriceSet(price);
    }

    /// @notice     Admin function to postpone the sale
    function postponeSale(uint256 timeToShift) external onlyOwner {
        require(block.timestamp < sale.saleStart, "Sale has started");
        sale.saleStart += timeToShift;
        require(
            sale.saleStart + timeToShift < sale.saleEnd,
            "Start time can not be greater than end time."
        );
    }

    /// @notice     Function to extend registration period
    function extendRegistrationTime(uint256 timeToExtend) external onlyOwner {
        require(
            registration.registrationTimeEnds + timeToExtend < sale.saleStart,
            "Registration period overflows sale start."
        );
        registration.registrationTimeEnds += timeToExtend;
    }

    /// @notice     Admin function to set max participation before sale start
    function setCap(uint256 cap) external onlyOwner {
        require(block.timestamp < sale.saleStart, "sale has started");
        require(cap > 0, "Cap must be greater than zero");
        sale.maxParticipation = cap;
        emit MaxParticipationSet(cap);
    }

    /////////////////////////////////⬇️user functions⬇️/////////////////////////////////

    /// @notice     Registration for sale.
    /// @param      signature is the message signed by the backend
    /// @param      pid is the id of the staking pool
    function registerForSale(bytes memory signature, uint256 pid) external {
        require(sale.isCreated, "Sale not created");
        require(
            block.timestamp >= registration.registrationTimeStarts &&
                block.timestamp <= registration.registrationTimeEnds,
            "Registration gate is closed."
        );
        require(!isRegistered[msg.sender], "Already registered");
        require(
            checkRegistrationSignature(signature, msg.sender),
            "Invalid signature"
        );
        isRegistered[msg.sender] = true;

        // Lock users stake
        allocationStaking.setTokensUnlockTime(
            pid,
            msg.sender,
            sale.tokensUnlockTime
        );

        // Increment number of registered users
        registration.numberOfRegistrants++;

        // Emit Registration event
        emit UserRegistered(msg.sender);
    }

    function participate(
        bytes memory signature,
        uint256 amount
    ) external payable {
        require(sale.isCreated, "Sale not created");
        // verify the timestamp
        require(
            block.timestamp >= sale.saleStart &&
                block.timestamp <= sale.saleEnd,
            "Sale is not active."
        );
        require(
            amount <= sale.maxParticipation,
            "Overflowing maximal participation for sale."
        );
        // user must have registered for the round in advance
        require(isRegistered[msg.sender], "Not registered");

        // check user haven't participated before
        require(!isParticipated[msg.sender], "User can participate only once.");

        // verify the signature
        require(
            checkParticipationSignature(signature, msg.sender, amount),
            "Invalid signature"
        );

        // Disallow contract calls.
        require(msg.sender == tx.origin, "Only direct contract calls.");

        // compute the amount of tokens user is buying
        uint256 amountOfTokensBuying = (msg.value *
            10 ** IERC20Metadata(address(sale.token)).decimals()) /
            sale.tokenPriceInETH;

        // must buy more than 0 tokens
        require(amountOfTokensBuying > 0, "Can't buy 0 tokens");

        // check in terms of user allo
        require(
            amountOfTokensBuying <= amount,
            "Trying to buy more than allowed."
        );

        // increase amount of sold tokens
        sale.totalTokensSold += amountOfTokensBuying;
        // increase amount of ETH raised
        sale.totalETHRaised += msg.value;

        bool[] memory _isPortionWithdrawn = new bool[](
            vestingPortionsUnlockTime.length
        );

        // create participation object
        Participation memory p = Participation({
            amountBought: amountOfTokensBuying,
            amountETHPaid: msg.value,
            timeParticipated: block.timestamp,
            isPortionWithdrawn: _isPortionWithdrawn
        });

        // add participation for user.
        userToParticipation[msg.sender] = p;
        // mark user is participated
        isParticipated[msg.sender] = true;
        // increment number of participants in the Sale
        numberOfParticipants++;

        emit TokensSold(msg.sender, amountOfTokensBuying);
    }

    function getSaleInfo()
        external
        view
        returns (
            uint256 totalTokensSold,
            uint256 totalETHRaised,
            uint256 amountOfTokensToSell,
            uint256 tokenPriceInETH
        )
    {
        totalTokensSold = sale.totalTokensSold;
        totalETHRaised = sale.totalETHRaised;
        amountOfTokensToSell = sale.amountOfTokensToSell;
        tokenPriceInETH = sale.tokenPriceInETH;
    }

    function withdrawTokens(uint256 portionId) external {
        require(
            block.timestamp >= sale.tokensUnlockTime,
            "Sale has not ended yet."
        );

        require(
            portionId < vestingPercentPerPortion.length,
            "Invalid portion id"
        );

        Participation memory p = userToParticipation[msg.sender];
        require(
            p.isPortionWithdrawn[portionId] == false,
            "Portion already withdrawn"
        );

        require(
            vestingPortionsUnlockTime[portionId] <= block.timestamp,
            "Portion not unlocked yet"
        );

        p.isPortionWithdrawn[portionId] = true;

        // calculate amount to withdraw
        uint256 amountToWithdraw = (p.amountBought *
            vestingPercentPerPortion[portionId]) / portionVestingPrecision;

        if (amountToWithdraw > 0) {
            // transfer tokens to user
            sale.token.safeTransfer(msg.sender, amountToWithdraw);
            emit TokensWithdrawn(msg.sender, amountToWithdraw);
        }
    }

    function withdrawMultiplePortions(uint256[] memory portionIds) external {
        require(
            block.timestamp >= sale.tokensUnlockTime,
            "Sale has not ended yet."
        );

        uint256 totalAmountToWithdraw = 0;
        Participation memory p = userToParticipation[msg.sender];

        for (uint i = 0; i < portionIds.length; i++) {
            uint256 portionId = portionIds[i];
            require(
                portionId < vestingPercentPerPortion.length,
                "Invalid portion id"
            );
            if ((p.isPortionWithdrawn[portionId]) == true) {
                continue;
            }
            p.isPortionWithdrawn[portionId] = true;
            totalAmountToWithdraw += ((p.amountBought *
                vestingPercentPerPortion[portionId]) / portionVestingPrecision);
        }

        if (totalAmountToWithdraw > 0) {
            // transfer tokens to user
            sale.token.safeTransfer(msg.sender, totalAmountToWithdraw);
            emit TokensWithdrawn(msg.sender, totalAmountToWithdraw);
        }
    }

    function getParticipation(
        address _user
    ) external view returns (uint256, uint256, uint256, bool[] memory) {
        Participation memory p = userToParticipation[_user];
        return (
            p.amountBought,
            p.amountETHPaid,
            p.timeParticipated,
            p.isPortionWithdrawn
        );
    }

    function getNumberOfRegisteredUsers() external view returns (uint256) {
        return registration.numberOfRegistrants;
    }

    function getVestingInfo()
        external
        view
        returns (uint256[] memory, uint256[] memory)
    {
        return (vestingPortionsUnlockTime, vestingPercentPerPortion);
    }

    /////////////////////////////////⬇️sale owner functions⬇️////////////////////////////////////
    /// @notice Function for owner to deposit tokens, can be called only once.
    function depositTokens(uint256 amount) external onlySaleOwner {
        require(!sale.tokensDeposited, "Tokens already deposited");
        sale.tokensDeposited = true;
        sale.token.safeTransferFrom(msg.sender, address(this), amount);
    }

    /// Function to withdraw all the earnings and the leftover of the sale contract.
    function withdrawEarningsAndLeftover() external onlySaleOwner {
        withdrawEarningsInternal();
        withdrawLeftoverInternal();
    }

    // Function to withdraw only earnings
    function withdrawEarnings() external onlySaleOwner {
        withdrawEarningsInternal();
    }

    // Function to withdraw only leftover
    function withdrawLeftover() external onlySaleOwner {
        withdrawLeftoverInternal();
    }

    /////////////////////////////////⬇️internal functions⬇️/////////////////////////////////

    // function to withdraw earnings
    function withdrawEarningsInternal() internal {
        require(block.timestamp >= sale.saleEnd, "Sale has not ended yet.");
        require(!sale.earningWithdrawn, "Earnings already withdrawn");
        sale.earningWithdrawn = true;
        uint256 amountToWithdraw = sale.totalETHRaised;
        sale.token.safeTransfer(msg.sender, amountToWithdraw);
        emit TokensWithdrawn(msg.sender, amountToWithdraw);
    }

    // Function to withdraw leftover
    function withdrawLeftoverInternal() internal {
        require(block.timestamp >= sale.saleEnd, "Sale has not ended yet.");
        require(!sale.leftoverWithdrawn, "Leftover already withdrawn");
        sale.leftoverWithdrawn = true;
        uint256 leftover = sale.amountOfTokensToSell - sale.totalTokensSold;
        if (leftover > 0) {
            sale.token.safeTransfer(msg.sender, leftover);
        }
    }

    function checkRegistrationSignature(
        bytes memory signature,
        address user
    ) internal view returns (bool) {
        bytes32 hash = keccak256(abi.encodePacked(user, address(this)));
        bytes32 messageHash = hash.toEthSignedMessageHash();
        return owner() == messageHash.recover(signature);
    }

    function checkParticipationSignature(
        bytes memory signature,
        address user,
        uint256 amount
    ) internal view returns (bool) {
        bytes32 hash = keccak256(abi.encodePacked(user, amount, address(this)));
        bytes32 messageHash = hash.toEthSignedMessageHash();
        return owner() == messageHash.recover(signature);
    }

    // Function to act as a fallback and handle receiving ETH.
    receive() external payable {}
}
