package commonutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

// Encrypt return encrypted byte
// 128 16byte/196 24byte /256 32byte
// support only aes256 only
func Encrypt(src, key, lv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error", err)
	}
	if src == nil {
		fmt.Println("data is nil")
	}
	ecb := cipher.NewCBCEncrypter(block, lv)
	src = PKCS5Padding(src, block.BlockSize())
	crypted := make([]byte, len(src))
	ecb.CryptBlocks(crypted, src)

	return crypted
}

// Decrypt return decrypted byte
// 128 16byte/196 24byte /256 32byte
// support only aes256 only
func Decrypt(crypt, key, lv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.New("key is nil")
	}
	if len(crypt) == 0 {
		return "", errors.New("plain content empty")
	}
	if (len(crypt) % block.BlockSize()) != 0 {
		return "", errors.New("input not nil block")
	}
	ecb := cipher.NewCBCDecrypter(block, lv)
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return string(PKCS5Trimming(decrypted)), nil
}

// PKCS5Padding return byte
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5Trimming return byte
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
