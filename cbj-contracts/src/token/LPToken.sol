// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";

/**
 * this is a LP Token template, we will deploy a logic contract and then clone it for each LP token we want to create. 
 * Each clone will have its own name and symbol, but they will all share the same code and logic, which saves gas and deployment costs.
 * 
 * for more details, please check the deployment script: script/LPTokenScript.s.sol
 * 
 * @title 
 * @author 
 * @notice 
 */
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
