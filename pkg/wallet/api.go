/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package wallet

// Wallet interface
type Wallet interface {
	Crypto
	Pack
}

// Crypto interface
type Crypto interface {

	// CreateSigningKey create a new public/private signing keypair.
	//
	// Returns:
	//
	// KeyInfo: A `KeyInfo` representing the keypair
	//
	// error: error
	CreateSigningKey(metadata map[string]string) (KeyInfo, error)

	// GetSigningKey Fetch info for a signing keypair.
	//
	// Args:
	//
	// verkey: The verification key of the keypair
	//
	// Returns:
	//
	// KeyInfo: A `KeyInfo` representing the keypair
	//
	// error: error
	GetSigningKey(verKey string) (KeyInfo, error)

	// SignMessage sign a message using the private key associated with a given verification key.
	//
	// Args:
	//
	// message: The message to sign
	//
	// fromVerKey: Sign using the private key related to this verification key
	//
	// Returns:
	//
	// []byte: The signature
	//
	// error: error
	SignMessage(message []byte, fromVerKey string) ([]byte, error)

	// DecryptMessage decrypt message
	//
	// Args:
	//
	// encMessage: The encrypted message content
	//
	// toVerKey:The verification key of the recipient.
	//
	// []byte: Decrypted message content
	//
	// string: The sender verification key
	//
	// error: error
	DecryptMessage(encMessage []byte, toVerKey string) ([]byte, string, error)
}

// Pack provide methods to pack and unpack msg
type Pack interface {
	// PackMessage Pack a message for one or more recipients.
	//
	// Args:
	//
	// envelope: The message to pack
	//
	// Returns:
	//
	// []byte: The packed message
	//
	// error: error
	PackMessage(envelope *Envelope) ([]byte, error)

	// UnpackMessage Unpack a message.
	//
	// Args:
	//
	// encMessage: The encrypted message
	//
	// Returns:
	//
	// envelope: unpack message
	//
	// error: error
	UnpackMessage(encMessage []byte) (*Envelope, error)
}

// KeyInfo contains public and private key
type KeyInfo interface {
	GetVerificationKey() string
	GetKeyMetadata() map[string]string
}

// Envelope contain msg,FromVerKey and ToVerKeys
type Envelope struct {
	Message    []byte
	FromVerKey string
	ToVerKeys  string
}
