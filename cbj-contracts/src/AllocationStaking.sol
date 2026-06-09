// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

import {ISalesFactory} from "./interfaces/ISalesFactory.sol";
import {IAllocationStaking} from "./interfaces/IAllocationStaking.sol";

contract AllocationStaking is
    OwnableUpgradeable,
    ReentrancyGuard,
    IAllocationStaking
{
    using SafeERC20 for IERC20;

    struct UserInfo {
        uint256 amount; // How many LP tokens the user has provided
        uint256 rewardDebt; // Reward debt. Current reward debt when user joined staking.
        uint256 tokensUnlockTime; // If user registered for sale, returns when tokens are getting unlocked
        address[] salesRegistered; // List of sales user registered for
    }

    struct PoolInfo {
        IERC20 lpToken; // Address of LP token contract
        uint256 allocPoint; // How many allocation points assigned to this pool. ERC20s to distribute per block
        uint256 lastRewardTime; // Last time that ERC20s distribution occurs.
        uint256 accERC20PerShare; // // Accumulated ERC20 per share, times 1e18.
        uint256 totalDeposits; // Total amount of tokens deposited at the moment (staked)
    }

    // Address of ERC20 token contract
    IERC20 public token;
    // staking start time
    uint256 public startTime;
    // staking end time
    uint256 public endTime;
    // The total amount of ERC20 that's paid out as reward
    uint256 public paidOut;
    // ERC20 tokens rewarded per second
    uint256 public rewardPerSecond;
    // Total rewards added to staking pool
    uint256 public totalRewards;
    // Address of sales factory contract
    ISalesFactory public salesFactory;

    // Info of each pool.
    PoolInfo[] public poolInfo;
    // Info of each user that stakes LP tokens.
    mapping(uint256 => mapping(address => UserInfo)) public userInfo;
    // Total allocation points. Must be the sum of all allocation points in all pools.
    uint256 public totalAllocPoint;

    event Deposit(address indexed user, uint256 indexed pid, uint256 amount);
    event Withdraw(address indexed user, uint256 indexed pid, uint256 amount);
    event EmergencyWithdraw(
        address indexed user,
        uint256 indexed pid,
        uint256 amount
    );
    event CompoundedEarnings(
        address indexed user,
        uint256 indexed pid,
        uint256 amountAdded,
        uint256 totalDeposited
    );

    // Restricting calls to only verified sales
    modifier onlyVerifiedSales() {
        require(
            salesFactory.isSaleCreatedThroughFactory(msg.sender),
            "Sale not created through factory."
        );
        _;
    }

    function initialize(
        address _initialOwner,
        address _token,
        uint256 _rewardPerSecond,
        address _salesFactory
    ) public initializer {
        require(
            _initialOwner != address(0),
            "initial owner cannot be zero address"
        );
        require(
            _rewardPerSecond > 0,
            "reward per second must be greater than 0"
        );
        require(
            address(_token) != address(0),
            "reward token address cannot be zero"
        );
        __Ownable_init(_initialOwner);
        token = IERC20(_token);
        rewardPerSecond = _rewardPerSecond;
        salesFactory = ISalesFactory(_salesFactory);
        startTime = block.timestamp + 60; // 1 分钟后开始
        endTime = startTime + 365 days; // 1 年后结束
    }

    function pending(
        uint256 _pid,
        address _user
    ) public view returns (uint256) {
        PoolInfo storage pool = poolInfo[_pid];
        UserInfo storage user = userInfo[_pid][_user];

        uint256 accERC20PerShare = pool.accERC20PerShare;

        // if current block time is greater than the last reward time, we need to calculate the pending reward with the current accERC20PerShare and the time elapsed since the last reward time
        if (block.timestamp > pool.lastRewardTime && pool.totalDeposits != 0) {
            uint256 lastTimestamp = block.timestamp > endTime
                ? endTime
                : block.timestamp;

            // the pool's lastRewardTime may be updated after endTime
            uint256 lastRewardTime = pool.lastRewardTime > endTime
                ? endTime
                : pool.lastRewardTime;

            uint256 pendingTime = lastTimestamp - lastRewardTime;

            uint256 poolReward = (pendingTime *
                rewardPerSecond *
                pool.allocPoint) / totalAllocPoint;

            accERC20PerShare += (poolReward * 1e18) / pool.totalDeposits;
        }

        // the reward calculation fomala: user.amount * accERC20PerShare - user.rewardDebt
        return (user.amount * accERC20PerShare) / 1e18 - user.rewardDebt;
    }

    function totalPending() public view returns (uint256) {
        if (block.timestamp <= startTime) {
            return 0;
        }
        uint256 lastTimestamp = block.timestamp > endTime
            ? endTime
            : block.timestamp;

        return (lastTimestamp - startTime) * rewardPerSecond - paidOut;
    }

    // Function where owner can set sales factory in case of upgrading some of smart-contracts
    function setSalesFactory(address _salesFactory) external onlyOwner {
        require(
            _salesFactory != address(0),
            "Sales factory address cannot be zero"
        );
        salesFactory = ISalesFactory(_salesFactory);
    }

    function poolLength() external view returns (uint256) {
        return poolInfo.length;
    }

    function addPool(
        uint256 _allocPoint,
        IERC20 _lpToken,
        bool _withUpdate
    ) external onlyOwner {
        require(
            address(_lpToken) != address(0),
            "LP token address cannot be zero"
        );
        if (_withUpdate) {
            massUpdatePools();
        }
        uint256 lastRewardTime = block.timestamp > startTime
            ? block.timestamp
            : startTime;
        totalAllocPoint += _allocPoint;
        poolInfo.push(
            PoolInfo({
                lpToken: _lpToken,
                allocPoint: _allocPoint,
                lastRewardTime: lastRewardTime,
                accERC20PerShare: 0,
                totalDeposits: 0
            })
        );
    }

    function fund(uint256 _amount) public {
        require(_amount > 0, "fund amount must be greater than 0");
        require(block.timestamp < endTime, "funding has already ended");
        token.safeTransferFrom(msg.sender, address(this), _amount);
        endTime += _amount / rewardPerSecond;
        totalRewards += _amount;
    }

    // update the given pool's reward allocation point. Can only be called by the owner.
    function setPool(
        uint256 _pid,
        uint256 _allocPoint,
        bool _withUpdate
    ) public onlyOwner {
        if (_withUpdate) {
            massUpdatePools();
        }

        totalAllocPoint =
            totalAllocPoint -
            poolInfo[_pid].allocPoint +
            _allocPoint;
        poolInfo[_pid].allocPoint = _allocPoint;
    }

    function massUpdatePools() private {
        uint256 length = poolInfo.length;
        for (uint256 pid = 0; pid < length; ++pid) {
            updatePool(pid);
        }
    }

    function updatePool(uint256 _pid) private {
        PoolInfo storage pool = poolInfo[_pid];

        // check staking if ended
        uint256 lastTime = block.timestamp > endTime
            ? endTime
            : block.timestamp;

        // the pool info has already been updated after the last reward time, so we can skip the update
        if (lastTime <= pool.lastRewardTime) {
            return;
        }

        // if there is no deposit in the pool, we just update the lastRewardTime and return
        uint256 lpSupply = pool.totalDeposits;
        if (lpSupply == 0) {
            pool.lastRewardTime = lastTime;
            return;
        }

        // calculate the reward for the pool
        uint256 timeElapsed = lastTime - pool.lastRewardTime;
        uint256 poolReward = (timeElapsed * rewardPerSecond * pool.allocPoint) /
            totalAllocPoint;

        // update the accumulated reward per share for the pool
        pool.accERC20PerShare =
            pool.accERC20PerShare +
            (poolReward * 1e18) /
            lpSupply;

        pool.lastRewardTime = lastTime;
    }

    // return the given user's deposited amount for the given pool
    // override the deposited function in IAllocationStaking interface
    function deposited(
        uint256 _pid,
        address _user
    ) public view override returns (uint256) {
        return userInfo[_pid][_user].amount;
    }

    // only verified sales can set tokens unlock time for users who registered for the sale
    // override the setTokensUnlockTime function in IAllocationStaking interface
    function setTokensUnlockTime(
        uint256 _pid,
        address _user,
        uint256 _unlockTime
    ) external override onlyVerifiedSales {
        UserInfo storage user = userInfo[_pid][_user];
        // Require that tokens are currently unlocked
        require(
            user.tokensUnlockTime <= block.timestamp,
            "Tokens unlock time is already set"
        );
        user.tokensUnlockTime = _unlockTime;
        // Add sale to the array of sales user registered for.
        user.salesRegistered.push(msg.sender);
    }

    function deposit(uint256 _pid, uint256 _amount) public {
        PoolInfo storage pool = poolInfo[_pid];
        UserInfo storage user = userInfo[_pid][msg.sender];

        // new deposit will change the reward distribution, so we need to update the pool
        // in this method,the pool will be updated before the user info is updated, so the user can get the reward for the previous deposit with the current accERC20PerShare
        updatePool(_pid);

        // if user has staked before,should settle current reward first with current accERC20PerShare
        if (user.amount > 0) {
            uint256 peddingReward = (user.amount * pool.accERC20PerShare) /
                1e18 -
                user.rewardDebt;
            token.safeTransfer(msg.sender, peddingReward);
            paidOut += peddingReward;
        }

        pool.totalDeposits += _amount;
        user.amount += _amount;
        user.rewardDebt = (user.amount * pool.accERC20PerShare) / 1e18;

        pool.lpToken.safeTransferFrom(msg.sender, address(this), _amount);

        emit Deposit(msg.sender, _pid, _amount);
    }

    // withdraw LP tokens from the pool;
    // calculate and transfer the pending reward to the user
    function withdraw(uint256 _pid, uint256 _amount) public nonReentrant {
        UserInfo storage user = userInfo[_pid][msg.sender];
        PoolInfo storage pool = poolInfo[_pid];

        require(user.amount >= _amount, "can't withdraw more than deposit");

        // calulate accumulated reward per share and update the pool
        updatePool(_pid);

        // calulate the pending reward for the user with the current accERC20PerShare
        uint256 reward = (user.amount * pool.accERC20PerShare) /
            1e18 -
            user.rewardDebt;

        // update user info and pool info before transfer the reward and LP tokens to prevent reentrancy attack
        user.amount -= _amount;
        user.rewardDebt = (user.amount * pool.accERC20PerShare) / 1e18;
        pool.totalDeposits -= _amount;
        paidOut += reward;

        // withdraw the LP tokens and reward to the user
        token.safeTransfer(msg.sender, reward);
        pool.lpToken.safeTransfer(msg.sender, _amount);

        emit Withdraw(msg.sender, _pid, _amount);
    }

    // Withdraw without caring about rewards. EMERGENCY ONLY
    function emergencyWithdraw(uint256 _pid) public nonReentrant {
        UserInfo storage user = userInfo[_pid][msg.sender];
        PoolInfo storage pool = poolInfo[_pid];

        uint256 amount = user.amount;

        // update user info and pool info before transfer the LP tokens to prevent reentrancy attack
        user.amount = 0;
        user.rewardDebt = 0;
        pool.totalDeposits -= amount;

        // withdraw the LP tokens to the user, but no reward
        pool.lpToken.safeTransfer(msg.sender, amount);

        emit EmergencyWithdraw(msg.sender, _pid, amount);
    }

    // compund the pending reward to the user's deposit.
    function compoundEarnings(uint256 _pid) public nonReentrant {
        UserInfo storage user = userInfo[_pid][msg.sender];
        PoolInfo storage pool = poolInfo[_pid];

        // calulate accumulated reward per share and update the pool
        updatePool(_pid);

        // calulate the pending reward for the user with the current accERC20PerShare
        uint256 reward = (user.amount * pool.accERC20PerShare) /
            1e18 -
            user.rewardDebt;

        require(reward > 0, "No pending reward to compound");

        // update user info and pool info before transfer the reward to prevent reentrancy attack
        user.amount += reward;
        user.rewardDebt = (user.amount * pool.accERC20PerShare) / 1e18;
        pool.totalDeposits += reward;
        paidOut += reward;

        emit CompoundedEarnings(msg.sender, _pid, reward, user.amount);
    }

    // Function to fetch deposits and earnings at one call for multiple users for passed pool id.
    function getPendingAndDepositedForUsers(
        uint256 _pid,
        address[] calldata _users
    )
        external
        view
        returns (
            uint256[] memory pendingRewards,
            uint256[] memory depositedAmounts
        )
    {
        pendingRewards = new uint256[](_users.length);
        depositedAmounts = new uint256[](_users.length);

        for (uint256 i = 0; i < _users.length; i++) {
            address userAddress = _users[i];
            pendingRewards[i] = deposited(_pid, userAddress);
            depositedAmounts[i] = pending(_pid, userAddress);
        }
    }
}
