// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Counter {
    uint256 private count;

    event CountIncreased(uint256 newCount);

    function increment() external {
        count += 1;
        emit CountIncreased(count);
    }

    function getCount() external view returns (uint256) {
        return count;
    }
}
