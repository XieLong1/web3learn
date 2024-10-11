// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract ViewAndPure {
    uint256 public x = 1;

    // Promise not to modify the state. 承诺不修改状态。
    function addToX(uint256 y) public view returns (uint256) {
        return x + y;
    }

    // Promise not to modify or read from the state. 承诺不修改或读取状态。
    function add(uint256 i, uint256 j) public pure returns (uint256) {
        return i + j;
    }
}