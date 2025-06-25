# GigaWallet Integration Setup

This document explains how to configure dogecoin-bridge-relayer to use GigaWallet for transaction sending.

## Overview

GigaWallet is a wallet service for Dogecoin transactions that provides a simple HTTP API interface. Through configuration, the original transaction sending functionality using btcd can be replaced with GigaWallet.

## Configuration

### Environment Variables

Set the following environment variables when starting the relayer:

```bash
# Enable GigaWallet
export USE_GIGAWALLET=true

# GigaWallet service address
export GIGAWALLET_URL=http://localhost:8080

# GigaWallet account ID
export GIGAWALLET_ACCOUNT_ID=user123
```

### Configuration Description

- `USE_GIGAWALLET`: Set to `true` to enable GigaWallet, set to `false` to use the original btcd/Fireblocks
- `GIGAWALLET_URL`: HTTP address of the GigaWallet service
- `GIGAWALLET_ACCOUNT_ID`: Account ID (foreignID) in GigaWallet

## How It Works

### Transaction Sending Process

1. **Transaction Building**: The system still uses the original logic to build BTC transactions
2. **Address Extraction**: Extract target addresses and amounts from transaction outputs
3. **GigaWallet Request**: Convert transactions to GigaWallet payout requests
4. **Status Checking**: Periodically check transaction status until confirmed

### Address Handling

- The system automatically identifies target addresses in transaction outputs (non-change addresses)
- Supports P2PKH, P2SH, P2WPKH and other address formats
- Skips OP_RETURN outputs

### Status Monitoring

- Periodically query GigaWallet transaction status
- Supported statuses: `pending`, `completed`, `failed`, `cancelled`, `rejected`
- Automatically handle transaction confirmation and failure scenarios

## Examples

### Starting GigaWallet

```bash
# Start GigaWallet service
docker run -p 8080:8080 gigawallet/gigawallet:latest
```

### Starting Relayer

```bash
# Set environment variables
export USE_GIGAWALLET=true
export GIGAWALLET_URL=http://localhost:8080
export GIGAWALLET_ACCOUNT_ID=myaccount

# Start relayer
./dogecoin-bridge-relayer
```

### Testing Transactions

```bash
# Use curl to test GigaWallet API
curl -X POST http://localhost:8080/account/myaccount/payout \
  -H "Content-Type: application/json" \
  -d '{
    "address": "DJ7zktMLiEbYZEMZP9ReinB3EcgP1cWwZV",
    "amount": 10.5,
    "label": "test-payment"
  }'
```

## Notes

1. **Address Format**: Ensure target addresses are valid Dogecoin addresses
2. **Amount Unit**: GigaWallet uses DOGE as the unit (floating point), the system automatically converts from satoshis
3. **Network Configuration**: Ensure GigaWallet connects to the correct Dogecoin network (mainnet/testnet)
4. **Error Handling**: The system automatically handles network errors and API errors
5. **Status Synchronization**: Periodically check transaction status to ensure synchronization with blockchain

## Troubleshooting

### Common Issues

1. **Connection Error**: Check if GigaWallet service is running normally
2. **Address Parsing Error**: Ensure transaction outputs contain valid addresses
3. **Insufficient Balance**: Check GigaWallet account balance
4. **Network Error**: Check network connection and firewall settings

### Logs

View relayer logs for detailed error information:

```bash
tail -f logs/relayer.log
```

## Migration Guide

Migrating from btcd/Fireblocks to GigaWallet:

1. **Backup Configuration**: Backup current configuration files
2. **Setup GigaWallet**: Follow the steps above to setup GigaWallet
3. **Update Configuration**: Set environment variables to enable GigaWallet
4. **Test**: Verify functionality in test environment
5. **Switch**: Enable GigaWallet in production environment
6. **Monitor**: Monitor transaction status and system performance 