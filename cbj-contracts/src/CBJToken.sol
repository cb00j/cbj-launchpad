// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract CBJToken is ERC20, Ownable {
    constructor(uint256 initialSupply) ERC20("CBJ", "CBJ") Ownable(msg.sender) {
        _mint(msg.sender, initialSupply);
    }
}
