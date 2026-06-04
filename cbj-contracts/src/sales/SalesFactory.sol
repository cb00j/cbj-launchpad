// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {CBJSale} from "./CBJSale.sol";

// a factory for creating sales contracts
contract SalesFactory is Ownable {
    // address of the allocation staking contract
    address public allocationStaking;
    // mapping to track if a sale was created through this factory
    mapping(address => bool) public isSaleCreatedThroughFactory;

    // list of all sales created through this factory;
    // Expose so query can be possible only by position as well
    address[] public allSales;

    event SaleDeployed(address saleAddress);

    event SaleOwnerAndTokenSetInFactory(
        address sale,
        address indexed saleToken,
        address indexed saleOwner
    );

    constructor(address admin, address _allocationStaking) Ownable(admin) {
        allocationStaking = _allocationStaking;
    }

    function setAllocationStaking(
        address _allocationStaking
    ) external onlyOwner {
        require(
            _allocationStaking != address(0),
            "Invalid allocation staking address"
        );
        allocationStaking = _allocationStaking;
    }

    function deploySale() external onlyOwner returns (address) {
        require(allocationStaking != address(0), "Allocation staking not set");

        CBJSale sale = new CBJSale{salt: bytes32(0)}(
            msg.sender,
            allocationStaking
        );

        isSaleCreatedThroughFactory[address(sale)] = true;
        allSales.push(address(sale));

        emit SaleDeployed(address(sale));
        return address(sale);
    }

    function getNumberOfSalesDeployed() external view returns (uint256) {
        return allSales.length;
    }

    function GetLastDeployedSale() external view returns (address) {
        require(allSales.length > 0, "No sales deployed");
        return allSales[allSales.length - 1];
    }

    function getAllSales(
        uint startIndex,
        uint endIndex
    ) external view returns (address[] memory) {
        address[] memory sales = new address[](endIndex - startIndex);
        uint index = 0;

        for (uint i = startIndex; i < endIndex; i++) {
            sales[index++] = allSales[i];
        }
        return sales;
    }
}
