package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
)

func Crypt(plaintext []byte, pkey rsa.PublicKey) (ciphertext []byte, combinedKeys []byte, hmacIntegrity []byte, err error) {
	// add padding to plaintext
	count := aes.BlockSize - (len(plaintext) % aes.BlockSize)
	padded := make([]byte, len(plaintext)+count)
	copy(padded, plaintext)
	copy(padded[len(plaintext):], bytes.Repeat([]byte{byte(count)}, count))
	plaintext = padded

	ciphertext = make([]byte, aes.BlockSize+len(plaintext))

	// aes initialization
	aesKey := make([]byte, 32)
	if _, err = io.ReadFull(rand.Reader, aesKey); err != nil {
		return nil, nil, nil, fmt.Errorf("unable to get randon aes key")
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("unable to create aes cipher: %v", err)
	}
	// iv generation
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, nil, fmt.Errorf("unable to get randon IV")
	}
	// aes crypt
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("unable to crypt block: %v", err)
	}

	// hmac initialization
	hmacKey := make([]byte, 32)
	if _, err = io.ReadFull(rand.Reader, hmacKey); err != nil {
		return nil, nil, nil, fmt.Errorf("unable to get randon aes key")
	}
	sig := hmac.New(sha256.New, hmacKey)
	sig.Write(ciphertext)
	// hmac computation
	hmacIntegrity = sig.Sum(nil)

	// protect keys with rsa
	combinedKeysRaw := append(aesKey, hmacKey...)
	combinedKeys, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, &pkey, combinedKeysRaw, nil)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("unable to crypt block: %v", err)
	}

	return ciphertext, combinedKeys, hmacIntegrity, nil
}

func Decrypt(ciphertext []byte, combinedKeys []byte, hmacIntegrity []byte, pkey rsa.PrivateKey) (plaintext []byte, err error) {
	combinedKeysRaw, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, &pkey, combinedKeys, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to decrypt block: %v", err)
	}

	// aes initialization
	aesKey := combinedKeysRaw[:32]
	aesBlock, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, fmt.Errorf("unable to create aes cipher: %v", err)
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("cipher length too short")
	}
	// find iv
	iv := ciphertext[:aes.BlockSize]
	if len(ciphertext[aes.BlockSize:])%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	// hmac initialization
	hmacKey := combinedKeysRaw[32:]
	sig := hmac.New(sha256.New, hmacKey)
	sig.Write(ciphertext)
	// hmac computation
	hmacIntegrityOnReceivedCipher := sig.Sum(nil)
	// hmac check
	if !bytes.Equal(hmacIntegrity, hmacIntegrityOnReceivedCipher) {
		return nil, errors.New("hmac verification failed")
	}

	// aes decrypt
	plaintext = make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(aesBlock, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	// remove iv
	plaintext = plaintext[aes.BlockSize:]

	// remove padding
	var errInvalidPadding = errors.New("invalid padding")
	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errInvalidPadding
	}
	c := plaintext[len(plaintext)-1]
	n := int(c)
	if n == 0 || n > len(plaintext) {
		return nil, errInvalidPadding
	}
	for i := 0; i < n; i++ {
		if plaintext[len(plaintext)-n+i] != c {
			return nil, errInvalidPadding
		}
	}
	plaintext = plaintext[:len(plaintext)-n]

	return plaintext, nil
}
