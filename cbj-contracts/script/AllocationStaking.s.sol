// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;
import {BaseScript} from "./BaseScript.s.sol";
import {AllocationStaking} from "../src/AllocationStaking.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {CBJToken} from "../src/token/CBJToken.sol";

contract AllocationStakingScript is BaseScript {
    function run() public broadcaster {
        AllocationStaking allocationStaking = new AllocationStaking{
            salt: bytes32(0)
        }();

        allocationStaking.initialize(
            deployer,
            getAddress("CBJToken"),
            1e18,
            getAddress("CBJToken")
        );
    }
}
