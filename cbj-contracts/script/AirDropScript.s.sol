// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {BaseScript} from "./BaseScript.s.sol";
import {CBJToken} from "../src/CBJToken.sol";
import {AirDrop} from "../src/AirDrop.sol";
import {console2} from "forge-std/console2.sol";

contract AirDropScript is BaseScript {
    function run() public broadcaster {
        bytes32 _salt = bytes32(
            0x0000000000000000000000000000000000000000000000000000000000000001
        );
        CBJToken token = new CBJToken{salt: _salt}(deployer, 1000_000 * 1e18);
        saveContract("CBJToken", address(token));
        console2.log(
            "CBJToken balance of deployer:",
            token.balanceOf(deployer)
        );
        AirDrop airDrop = new AirDrop{salt: _salt}(address(token));
        saveContract("AirDrop", address(airDrop));
        require(
            token.transfer(address(airDrop), 1000_000 * 1e18),
            "initial transfer failed"
        );
    }
}
