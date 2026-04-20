// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;
import {BaseScript} from "./BaseScript.s.sol";
import {FarmingCBJ} from "../src/farming/FarmingCBJ.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {CBJToken} from "../src/token/CBJToken.sol";

contract FarmingCBJScript is BaseScript {
    function run() public broadcaster {
        FarmingCBJ farmingCBJ = new FarmingCBJ{salt: bytes32(0)}(
            deployer,
            IERC20(getAddress("CBJToken")),
            1e18, // 每秒 1 个 CBJ 作为奖励
            block.timestamp + 60, // 1 分钟后开始
            block.timestamp + 365 days // 1 年后结束
        );

        CBJToken rewardToken = CBJToken(getAddress("CBJToken"));
        rewardToken.approve(address(farmingCBJ), 1000_000 * 1e18); // 授权 FarmingCBJ 合约可以转移奖励代币

        farmingCBJ.fund(1000_000 * 1e18); // 预先注资 100 万个 CBJ 作为奖励
        saveContract("FarmingCBJ", address(farmingCBJ));
    }
}
