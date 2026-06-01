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
            1e18 // 每秒 1 个 CBJ 作为奖励
        );

        CBJToken rewardToken = CBJToken(getAddress("CBJToken"));
        rewardToken.approve(address(farmingCBJ), 1000_000 * 1e18); // 授权 FarmingCBJ 合约可以转移奖励代币

        farmingCBJ.fund(1000_000 * 1e18); // 预先注资 100 万个 CBJ 作为奖励
        farmingCBJ.addPool(IERC20(getAddress("LP-USDT/CBJ")), 100, true); // 初始化一个流动性池，分配权重为 100
        saveContract("FarmingCBJ", address(farmingCBJ));
    }
}
