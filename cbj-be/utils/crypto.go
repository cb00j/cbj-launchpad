package utils

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var Credentials *CredentialsUtils

type CredentialsUtils struct {
	privateKey *ecdsa.PrivateKey
}

func (cu *CredentialsUtils) NewCredentialsUtils(privateKeyHex string) (string, error) {
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return "", fmt.Errorf("failed to convert private key: %w", err)
	}
	cu.privateKey = privateKey
	return "", nil
}

func (cu *CredentialsUtils) GetSign(hexStr string) (string, error) {
	// convert hex string to bytes
	data, err := hexutil.Decode(hexStr)
	if err != nil {
		return "", fmt.Errorf("decode hex failed: %w", err)
	}

	// hexStr has been decoded to bytes, now we need to hash it using Keccak256
	dataHash := crypto.Keccak256(data)

	// EIP-191: adding the prefix and length to the hash before signing
	cryptoKey := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(dataHash))), dataHash)

	// load private key
	LoadPrivateKey, err := LoadPrivateKey()
	if err != nil {
		return "", fmt.Errorf("load private key failed: %w", err)
	}

	// sign the hash with the private key using ECDSA
	signature, err := crypto.Sign(cryptoKey, LoadPrivateKey)
	if err != nil {
		return "", fmt.Errorf("sign failed: %w", err)
	}

	signature[64] += 27 // make it compatible with Ethereum

	return hexutil.Encode(signature), nil
}

func LoadPrivateKey() (*ecdsa.PrivateKey, error) {
	// load private key from environment variable
	pk := os.Getenv("SIGNER_PRIVATE_KEY")
	return crypto.HexToECDSA(strings.TrimPrefix(pk, "0x"))
}
