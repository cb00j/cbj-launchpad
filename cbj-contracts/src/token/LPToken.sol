// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";

contract LPToken is ERC20, Initializable {
    string private _customName;
    string private _customSymbol;

    constructor() ERC20("", "") {}

    function initialize(
        string memory name_,
        string memory symbol_
    ) public initializer {
        _customName = name_;
        _customSymbol = symbol_;
    }

    function name() public view override returns (string memory) {
        return _customName;
    }

    function symbol() public view override returns (string memory) {
        return _customSymbol;
    }

    function mint(address to, uint256 amount) external {
        super._mint(to, amount);
    }
}
