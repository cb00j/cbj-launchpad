// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

abstract contract BaseScript is Script {
    uint256 internal deployerPrivateKey;
    address internal deployer;
    address internal factory = 0x4e59b44847b379578588920cA78FbF26c0B4956C;

    // 定义统一的部署记录文件路径
    string constant DEPLOYMENT_PATH = "./deployments/contract_addresses.json";

    function setUp() public virtual {
        deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        deployer = vm.addr(deployerPrivateKey);
    }

    // 保存合约地址
    function saveContract(string memory name, address addr) public {
        string memory network = _getNetworkName();

        // 1. 确保目录和文件存在（如果文件不存在，先创建一个空的 JSON 对象）
        if (!vm.exists(DEPLOYMENT_PATH)) {
            vm.writeFile(DEPLOYMENT_PATH, "{}");
        }

        // 2. 序列化当前合约地址
        // 这里使用一个临时的 key "addressObj" 来构建内部对象
        string memory jsonKey = "temp_key";
        string memory contractJson = vm.serializeAddress(jsonKey, name, addr);

        // 3. 将合约地址写入对应的网络 key 下
        // .networkName 表示在 JSON 根部的 networkName 字段下操作
        vm.writeJson(
            contractJson,
            DEPLOYMENT_PATH,
            string.concat(".", network)
        );

        console2.log(string.concat("Saved ", name, " on ", network, ":"), addr);
    }

    // 修改后的获取地址方法
    function getAddress(string memory name) public view returns (address) {
        string memory network = _getNetworkName();
        string memory jsonContent = vm.readFile(DEPLOYMENT_PATH);

        // 根据路径读取：.networkName.contractName
        string memory path = string.concat(".", network, ".", name);
        address result = vm.parseJsonAddress(jsonContent, path);

        return result;
    }

    // 内部辅助方法：将 ChainID 映射为网络名称
    function _getNetworkName() internal view returns (string memory) {
        uint256 chainId = block.chainid;
        if (chainId == 31337) return "local";
        if (chainId == 11155111) return "sepolia";
        if (chainId == 421614) return "arb_sepolia";
        if (chainId == 1) return "mainnet";

        return Strings.toString(chainId); // 默认返回 chainId 字符串
    }

    modifier broadcaster() {
        vm.startBroadcast(deployerPrivateKey);
        _;
        vm.stopBroadcast();
    }
}
