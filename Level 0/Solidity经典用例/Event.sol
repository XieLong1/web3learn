// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Event {
    //事件声明
    //最多可以索引3个参数。
    //通过索引参数对日志进行过滤
    event Log(address indexed sender, string message);
    event AnotherLog();

    function test() public {
        emit Log(msg.sender, "Hello World!");
        emit Log(msg.sender, "Hello EVM!");
        emit AnotherLog();
    }
}