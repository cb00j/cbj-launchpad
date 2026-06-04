// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

interface IAllocationStaking {
    // how many LP tokens the user has provided
    function deposited(
        uint256 _pid,
        address _user
    ) external view returns (uint256);

    // set the unlock time for a user's staked tokens in a specific pool
    function setTokensUnlockTime(
        uint256 _pid,
        address _user,
        uint256 _unlockTime
    ) external;
}
