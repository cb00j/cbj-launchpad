// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {BaseScript} from "./BaseScript.s.sol";
import {LPToken} from "../src/token/LPToken.sol";
import {Clones} from "@openzeppelin/contracts/proxy/Clones.sol";
import {console2} from "forge-std/console2.sol";

contract LPTokenScript is BaseScript {
    function run() public broadcaster {
        // 1. 部署逻辑合约（只需部署一次）
        LPToken impl = new LPToken{salt: bytes32(0)}();
        saveContract("LPToken", address(impl));
        // 2. 定义我们想要的 LP 列表
        string[1] memory lpTokens = ["LP-USDT/CBJ"];

        for (uint i = 0; i < lpTokens.length; i++) {
            // 使用名字作为 salt 的一部分，确保每个 LP 地址不同但固定
            bytes32 salt = keccak256(abi.encodePacked(lpTokens[i]));
            // 3. 使用 CREATE2 克隆代理合约
            // Clones.cloneDeterministic 是 OpenZeppelin 提供的 CREATE2 克隆方法
            address lpToken = Clones.cloneDeterministic(address(impl), salt);
            // 4. 初始化代理合约
            LPToken(lpToken).initialize(lpTokens[i], lpTokens[i]);
            vm.label(lpToken, lpTokens[i]);
            saveContract(lpTokens[i], lpToken);
            // 测试：顺便给测试账号铸造点币
            LPToken(lpToken).mint(deployer, 100 * 1e18);
        }
    }
}
