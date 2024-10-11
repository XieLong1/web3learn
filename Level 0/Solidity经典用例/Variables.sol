// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Variables {
    // 状态变量存储在区块链上。
    string public text = "Hello";
    uint256 public num = 123;

    function doSomething() public view {
        // 局部变量不会保存到区块链中。
        uint256 i = 456;
        // 这里有一些全局变量
        uint256 timestamp = block.timestamp; // 当前块时间戳
        address sender = msg.sender; // address of the caller
    }
}