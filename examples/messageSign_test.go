package examples

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/kalyan3104/k-chain-core-go/display"
	"github.com/kalyan3104/k-chain-core-go/hashing/keccak"
	"github.com/kalyan3104/k-chain-crypto-go/signing"
	"github.com/kalyan3104/k-chain-crypto-go/signing/ed25519"
	"github.com/stretchr/testify/require"
)

/*
	How message signing works:

	Signing:
	- a user signs the hash of the calculated payload based on the message with the private key of the address
	- data to be signed = keccakHash(prefix + len(message) + message)

	Verifying:
	- the address, the original message and the signature have to be provided
	- the hash is calculated again and the signature validity is checked based on the public key (address)
*/

// This prefix should be added when computing the hash to be signed
const signedMessagePrefix = "\x17Numbat Signed Message:\n"

var messageSigningHasher = keccak.NewKeccak()

// func TestVerifyMessageSignatureFromLedger(t *testing.T) {
// 	// these field values were obtained by using Kalyan App for Ledger
// 	address := "moa19pht2w242wcj0x9gq3us86dtjrrfe3wk8ffh5nhdemf0mce6hsms3eqdrs"
// 	message := "test message"
// 	signature := "ec7a27cb4b23641ae62e3ea96d5858c8142e20d79a6e1710037d1c27b0d138d7452a98da93c036b2b47ee587d4cb4af6ae24c358f3f5f74f85580f45e072280b"

// 	err := checkSignature(address, message, signature)
// 	require.NoError(t, err)
// }

func TestVerifyMessageSignature(t *testing.T) {
	address := "moa1qyu5wthldzr8wx5c9ucg8kjagg0jfs53s8nr3zpz3hypefsdd8ssfq94h8" // alice
	message := "custom message of Alice"
	signature := "07e3e890e53c43e549800bd13540da28459f72d88b0b34f19c91e39a9fc5134ed75bbdeea63414db874e16837020e6d08448234613e130773d776ef645e27e00"

	err := checkSignature(address, message, signature)
	require.NoError(t, err)
}

func TestMessageSigning(t *testing.T) {
	messageToSign := "custom message of Alice"
	address, hash, signature := signMessage(t, alicePrivateKeyHex, messageToSign)

	header := []string{"Parameter", "Value"}
	lines := []*display.LineData{
		display.NewLineData(false, []string{"Message to sign", messageToSign}),
		display.NewLineData(false, []string{"Bech32 Address of signer", address}),
		display.NewLineData(false, []string{"Hash that was signed", hash}),
		display.NewLineData(false, []string{"Signature", signature}),
	}

	table, _ := display.CreateTableString(header, lines)
	fmt.Println(table)
}

func checkSignature(address string, message string, signature string) error {
	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		return err
	}

	return checkMessageSignature(address, message, sigBytes)
}

func signMessage(t *testing.T, senderSeedHex string, message string) (string, string, string) {
	keyGenerator := signing.NewKeyGenerator(signingCryptoSuite)

	senderSeed, err := hex.DecodeString(senderSeedHex)
	require.Nil(t, err)

	privateKey, err := keyGenerator.PrivateKeyFromByteArray(senderSeed)
	require.Nil(t, err)

	hash := computeHashForMessage(message)

	signature, err := signer.Sign(privateKey, hash)
	require.Nil(t, err)
	require.Len(t, signature, 64)

	publicKey := privateKey.GeneratePublic()
	publicKeyBytes, err := publicKey.ToByteArray()
	require.NoError(t, err)

	address, err := addressEncoder.Encode(publicKeyBytes)
	require.NoError(t, err)

	return address, hex.EncodeToString(hash), hex.EncodeToString(signature)
}

func computeHashForMessage(message string) []byte {
	payloadForHash := fmt.Sprintf("%s%v%s", signedMessagePrefix, len(message), message)
	hash := messageSigningHasher.Compute(payloadForHash)

	return hash
}

func checkMessageSignature(address string, message string, signature []byte) error {
	hash := computeHashForMessage(message)
	suite := ed25519.NewEd25519()
	keyGen := signing.NewKeyGenerator(suite)

	addressBytes, err := addressEncoder.Decode(address)
	if err != nil {
		return err
	}

	publicKey, err := keyGen.PublicKeyFromByteArray(addressBytes)
	if err != nil {
		return err
	}

	return signer.Verify(publicKey, hash, signature)
}
