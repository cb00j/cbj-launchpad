// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract AirDrop {
    using SafeERC20 for IERC20;

    IERC20 public token;

    mapping(address => bool) wasClaimed;

    uint256 public totalClaimedAmount;

    uint256 private constant DEFAULT_AMOUNT_PER_CLAIM = 100 * 1e18;

    uint256 public amountPerClaim;

    event TokenAirDropped(address indexed user, uint256 amount);

    constructor(address _token, uint256 _amountPerClaim) {
        require(_token != address(0), "AirDrop: invalid address");
        token = IERC20(_token);
        amountPerClaim = _amountPerClaim < DEFAULT_AMOUNT_PER_CLAIM
            ? DEFAULT_AMOUNT_PER_CLAIM
            : _amountPerClaim;
    }

    function claim() external {
        require(msg.sender == tx.origin, "AirDrop: not EOA");
        require(!wasClaimed[msg.sender], "AirDrop: already claimed");

        wasClaimed[msg.sender] = true;
        token.safeTransfer(msg.sender, amountPerClaim);
        totalClaimedAmount += amountPerClaim;

        emit TokenAirDropped(msg.sender, amountPerClaim);
    }
}
