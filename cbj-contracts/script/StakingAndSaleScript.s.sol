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
        address sale = salesFactory.deploySale();

        CBJSale cbjSale = CBJSale(payable(sale));

        // ===== 时间排期(UTC 基准,对齐 6 天周期)=====
        uint256 day0 = block.timestamp - (block.timestamp % 1 days); // 当天UTC 0点
        uint256 day1 = day0 + 1 days; // 第2天:销售
        uint256 day2 = day0 + 2 days; // 第3天:解锁/vesting

        // 1. setSaleParams(先设,含 saleEnd)
        cbjSale.setSaleParams(
            getAddress("CBJToken"), // _token
            deployer, // _saleOwner
            5e17, // _tokenPriceInETH = 0.5 ETH
            10000 * 1e18, // _amountOfTokensToSell
            day1 + 1 days - 1, // _saleEnd: 第2天 23:59:59
            day2, // _tokensUnlockTime: 第3天 0点
            1000 * 1e18, // _maxParticipation
            1000 // _portionVestingPrecision
        );

        // 2. setSaleStart(后设,要求 > now 且 < saleEnd)
        //    销售开始:第2天 0点(day1)。day1 > now ✓,day1 < saleEnd(day1结束) ✓
        cbjSale.setSaleStart(day1);

        // 3. setRegistrationTime:报名第1天(部署后1分钟 ~ 第1天结束)
        cbjSale.setRegistrationTime(
            block.timestamp + 60, // 报名开始(满足 > now)
            day1 - 1 // 报名结束:第1天 23:59:59
        );

        // 4. setVestingParams:4 批从第3天(day2)起,每天一批
        uint256[] memory unlockTimes = new uint256[](4);
        unlockTimes[0] = day2; // 第3天
        unlockTimes[1] = day2 + 1 days; // 第4天
        unlockTimes[2] = day2 + 2 days; // 第5天
        unlockTimes[3] = day2 + 3 days; // 第6天

        uint256[] memory percentPerPortion = new uint256[](4);
        percentPerPortion[0] = 100;
        percentPerPortion[1] = 200;
        percentPerPortion[2] = 300;
        percentPerPortion[3] = 400;

        cbjSale.setVestingParams(unlockTimes, percentPerPortion, 30 days);

        // 5. 注资 + 存币
        CBJToken cbjToken = CBJToken(getAddress("CBJToken"));
        cbjToken.approve(address(allocationStaking), 1000_000 * 1e18);
        cbjToken.approve(address(cbjSale), 1000_000 * 1e18);

        cbjSale.depositTokens(1000_000 * 1e18);

        allocationStaking.fund(1000_000 * 1e18);
        allocationStaking.addPool(100, IERC20(getAddress("LP-CBJ")), true);

        saveContract("AllocationStaking", address(allocationStaking));
        saveContract("SalesFactory", address(salesFactory));
        saveContract("CBJSale", address(sale));
    }
}
