// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Primitives {
    bool public boo = true;

    /*
    Uint代表无符号整数，即非负整数
    不同的尺寸可供选择
    Uint8的取值范围是0到2 ** 8 - 1
    Uint16的取值范围是0到2 ** 16 - 1
    …
    Uint256的取值范围是0到2 ** 256 - 1
    */
    uint8 public u8 = 1;
    uint256 public u256 = 456;
    uint256 public u = 123; // uint is an alias for uint256

    /*
    int类型允许使用负数。
    像int一样，从int8到int256也有不同的范围
    Int256的取值范围是-2 ** 255 ~ 2 ** 255 - 1
    Int128的取值范围是-2 ** 127 ~ 2 ** 127 - 1
    */
    int8 public i8 = -1;
    int256 public i256 = 456;
    int256 public i = -123; // int is same as int256

    // minimum and maximum of int
    int256 public minInt = type(int256).min;
    int256 public maxInt = type(int256).max;

    address public addr = 0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c;

    /*
    在Solidity中，数据类型byte表示一个字节序列。
    Solidity提供了两种类型的字节类型:
    -固定大小的字节数组
    -动态大小的字节数组。
    Solidity中的术语bytes表示一个动态的字节数组。
    它是byte[]的简写。
    */
    bytes1 a = 0xb5; //  [10110101]
    bytes1 b = 0x56; //  [01010110]

    // Default values
    // Unassigned variables have a default value
    bool public defaultBoo; // false
    uint256 public defaultUint; // 0
    int256 public defaultInt; // 0
    address public defaultAddr; // 0x0000000000000000000000000000000000000000
}