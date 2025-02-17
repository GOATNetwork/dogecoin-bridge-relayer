// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract TaskPayload {
    enum AddressCategory {
        BTC,
        EVM,
        SOL,
        SUI,
        EVM_TSS
    }

    event TransferRequest (
        string fromAddress,
        string toAddress,
        bytes32 ticker,
        uint64 chainId,
        uint256 amount,
        uint256 offset
    );

    event ConsolidationRequest (
        string fromAddress,
        bytes32 ticker,
        uint64 chainId,
        uint256 amount,
        uint256 offset
    );

    event WalletCreationRequest(
        address _userAddr,
        uint32 _account,
        AddressCategory _chain,
        uint32 _index
    );

    event WithdrawalRequest(
        address userAddress,
        uint64 chainId,
        bytes32 ticker,
        string toAddress,
        uint256 amount,
        uint256 offset
    );

    event DepositRequest(
        address userAddress,    
        uint64 chainId,
        bytes32 ticker,
        string depositAddress,
        uint256 amount,
        string txHash,
        uint256 blockHeight,
        uint256 logIndex
    );

    event TaskResult(uint8 version, bool success, uint8 errorCode);

    event WalletCreationResult(
        uint8 version,
        bool success,
        uint8 errorCode,
        string walletAddress
    );

    event DepositResult(uint8 version, bool success, uint8 errorCode);

    event WithdrawalResult(uint8 version, bool success, uint8 errorCode);
}
