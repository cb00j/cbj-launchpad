// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract LPToken is ERC20, Ownable {
    constructor(
        address _owner,
        uint256 initialSupply
    ) ERC20("LP-USDT2CBJ", "LP-USDT2CBJ") Ownable(_owner) {
        _mint(_owner, initialSupply);
    }
}
