package signer

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signer struct {
	// it's private key, so we don't export it
	privateKey *ecdsa.PrivateKey
}

func NewSigner(privateKey string) (*Signer, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}
	return &Signer{privateKey: key}, nil
}

func (s *Signer) GetSign(hexStr string) (string, error) {
	// convert hex string to bytes
	data, err := hexutil.Decode(hexStr)
	if err != nil {
		return "", fmt.Errorf("decode hex failed: %w", err)
	}

	// hexStr has been decoded to bytes, now we need to hash it using Keccak256
	dataHash := crypto.Keccak256(data)

	// EIP-191: adding the prefix and length to the hash before signing
	cryptoKey := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(dataHash))), dataHash)

	// sign the hash with the private key using ECDSA
	signature, err := crypto.Sign(cryptoKey, s.privateKey)
	if err != nil {
		return "", fmt.Errorf("sign failed: %w", err)
	}

	signature[64] += 27 // make it compatible with Ethereum

	return hexutil.Encode(signature), nil
}
