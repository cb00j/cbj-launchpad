// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract FarmingCBJ is Ownable {
    using SafeERC20 for IERC20;

    struct UserInfo {
        uint256 amount; // How many LP tokens the user has provided.
        uint256 rewardDebt; // Reward debt
    }

    struct PoolInfo {
        IERC20 lpToken; // Address of LP token contract.
        uint256 allocPoint; // How many allocation points assigned to this pool.
        uint256 accERC20PerShare; // Accumulated ERC20 per share, times 1e18.
        uint256 lastRewardTime; // Last time that ERC20 reward distribution occurs.
        uint256 totalDeposits; // Total amount of LP tokens deposited in the pool
    }

    // the reward token
    IERC20 public token;

    // farming start time
    uint256 public startTime;
    // farming end time
    uint256 public endTime;

    // the total amount that's paid out as rewards
    uint256 public paidOut;

    // reward per second,it's shared by all pools, the actual reward each pool gets is based on its allocation points and total allocation points
    uint256 public rewardPerSecond;

    // Info of each pool.
    PoolInfo[] public poolInfo;
    // Info of each user that stakes LP tokens.
    // mapping of pool id => user address => UserInfo
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

    constructor(
        address _initialOwner,
        IERC20 _token,
        uint256 _rewardPerSecond,
        uint256 _startTime,
        uint256 _endTime
    ) Ownable(_initialOwner) {
        require(
            _initialOwner != address(0),
            "initial owner cannot be zero address"
        );
        require(_endTime > _startTime, "end time must be after start time");
        require(
            _rewardPerSecond > 0,
            "reward per second must be greater than 0"
        );
        require(
            address(_token) != address(0),
            "reward token address cannot be zero"
        );
        token = _token;
        rewardPerSecond = _rewardPerSecond;
        startTime = _startTime;
        endTime = _endTime;
    }

    function deposit(uint256 _pid, uint256 _amount) public {
        PoolInfo storage pool = poolInfo[_pid];
        UserInfo storage user = userInfo[_pid][msg.sender];

        // new deposit will change the reward distribution, so we need to update the pool
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
        user.rewardDebt += (user.amount * pool.accERC20PerShare) / 1e18;

        pool.lpToken.safeTransferFrom(msg.sender, address(this), _amount);

        emit Deposit(msg.sender, _pid, _amount);
    }

    function addPool(
        IERC20 _lpToken,
        uint256 _allocPoint,
        bool _withUpate
    ) public onlyOwner {
        require(
            address(_lpToken) != address(0),
            "LP token address cannot be zero"
        );

        // add a new pool may change the reward distribution(totalAllocPoint may be updated), so we need to update all pools
        if (_withUpate) {
            massUpdatePools();
        }

        totalAllocPoint += _allocPoint;

        uint256 lastRewardTime = block.timestamp > startTime
            ? block.timestamp
            : startTime;

        poolInfo.push(
            PoolInfo({
                lpToken: _lpToken,
                allocPoint: _allocPoint,
                accERC20PerShare: 0,
                lastRewardTime: lastRewardTime,
                totalDeposits: 0
            })
        );
    }

    function massUpdatePools() private {
        uint256 length = poolInfo.length;
        for (uint256 pid = 0; pid < length; ++pid) {
            updatePool(pid);
        }
    }

    function updatePool(uint256 _pid) private {
        PoolInfo storage pool = poolInfo[_pid];

        // check farming if ended
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
}
