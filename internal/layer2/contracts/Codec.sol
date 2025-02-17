// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

enum State {
    Created,
    Pending,
    Completed,
    Failed
}

struct Operation {
    uint64 taskId; // 8 bytes
    State state; // 1 byte
    bytes extraData;
}


contract Codec {
    constructor(Operation[] memory opts,uint256 nonce, uint256 chainId) {}
    
    event operations(Operation[] opts, uint256 nonce, uint256 chainId);
}
