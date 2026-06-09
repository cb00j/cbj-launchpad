// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;
import {BaseScript} from "./BaseScript.s.sol";
import {AllocationStaking} from "../src/AllocationStaking.sol";
import {SalesFactory} from "../src/sales/SalesFactory.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {CBJToken} from "../src/token/CBJToken.sol";
import {CBJSale} from "../src/sales/CBJSale.sol";

contract StakingAndSaleScript is BaseScript {
    function run() public broadcaster {
        AllocationStaking allocationStaking = new AllocationStaking{
            salt: bytes32(0)
        }();

        SalesFactory salesFactory = new SalesFactory{salt: bytes32(0)}(
            deployer
        );

        allocationStaking.initialize(
            deployer,
            getAddress("CBJToken"),
            1e18,
            address(salesFactory)
        );

        salesFactory.setAllocationStaking(address(allocationStaking));
        address sale = salesFactory.deploySale(); // Deploy a sale through the factory

        CBJSale cbjSale = CBJSale(payable(sale));
        // initialize the sale
        cbjSale.setSaleParams(
            getAddress("CBJToken"),
            deployer,
            0.5 * 1e18,
            10000 * 1e18,
            block.timestamp + 365 days,
            block.timestamp + 365 days + 1000000,
            1000 * 1e18,
            100
        );
        cbjSale.setRegistrationTime(
            block.timestamp + 10,
            block.timestamp + 365 days
        );

        CBJToken rewardToken = CBJToken(getAddress("CBJToken"));
        rewardToken.approve(address(allocationStaking), 1000_000 * 1e18);

        allocationStaking.fund(1000_000 * 1e18);
        allocationStaking.addPool(100, IERC20(getAddress("LP-CBJ/USDT")), true); // 初始化一个流动性池，分配权重为 100

        saveContract("AllocationStaking", address(allocationStaking));
        saveContract("SalesFactory", address(salesFactory));
        saveContract("CBJSale", address(sale));
    }
}
