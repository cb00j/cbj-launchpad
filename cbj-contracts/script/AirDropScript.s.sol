// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {BaseScript} from "./BaseScript.s.sol";
import {CBJToken} from "../src/CBJToken.sol";
import {AirDrop} from "../src/AirDrop.sol";

contract AirDropScript is BaseScript {
    function run() public broadcaster {
        CBJToken token = new CBJToken(1000_000 * 1e18);
        saveContract("CBJToken", address(token));
        AirDrop airDrop = new AirDrop(address(token));
        saveContract("AirDrop", address(airDrop));
        token.transfer(address(airDrop), 1000_000 * 1e18);
    }
}
