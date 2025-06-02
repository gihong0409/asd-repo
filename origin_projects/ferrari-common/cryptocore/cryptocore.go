package cryptocore

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
)

// strProductKey := "eSre!inUataD 2@18 mAcSitna $Sca*"
var PROD_KEY = []byte{
	0x65, 0x53, 0x72, 0x65, 0x21, 0x69, 0x6e, 0x55, 0x61, 0x74,
	0x61, 0x44, 0x20, 0x32, 0x40, 0x31, 0x38, 0x20, 0x6d, 0x41,
	0x63, 0x53, 0x69, 0x74, 0x6e, 0x61, 0x20, 0x24, 0x53, 0x63,
	0x61, 0x2a,
}

// strProductKey := "B!y P^es$si1mo3n"
var PROD_IV = []byte{
	0x42, 0x21, 0x79, 0x20, 0x50, 0x5e, 0x65, 0x73, 0x24, 0x73,
	0x69, 0x31, 0x6d, 0x6f, 0x33, 0x6e,
}

func GenerateKeyPair() *rsa.PrivateKey {
	reader := rand.Reader
	bitSize := 1024
	key, _ := rsa.GenerateKey(reader, bitSize)
	return key
}

func PublicToPEMKey(rsaKey rsa.PublicKey) []byte {
	asn256Bytes, err := x509.MarshalPKIXPublicKey(&rsaKey)
	if err != nil {
		return nil
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn256Bytes,
	}
	return pem.EncodeToMemory(pemkey)
}

func PrivateToPEMKey(key *rsa.PrivateKey) []byte {
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	return pem.EncodeToMemory(privateKey)
}

func ProductKeyEncription(plainText []byte) string {
	var err error
	var block cipher.Block
	// Pad key
	if block, err = aes.NewCipher(PROD_KEY); err != nil {
		return ""
	}

	ciphertext := make([]byte, len(plainText)+(aes.BlockSize-len(plainText)%aes.BlockSize))

	cbc := cipher.NewCBCEncrypter(block, PROD_IV)
	plainData := PKCS5Padding(plainText, aes.BlockSize)
	cbc.CryptBlocks(ciphertext, plainData)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func ProductKeyDecrypt(strBase64EncKeyData string) string {

	encData, err := base64.StdEncoding.DecodeString(strBase64EncKeyData)
	if err != nil {
		return ""
	}

	var block cipher.Block
	// Pad key
	if block, err = aes.NewCipher(PROD_KEY); err != nil {
		return ""
	}

	plainText := make([]byte, len(encData))

	cbc := cipher.NewCBCDecrypter(block, PROD_IV)
	cbc.CryptBlocks(plainText, encData)

	// fmt.Printf("%s", ciphertext)

	return string(PKCS5Trimming(plainText))
}

func EncryptSessionKey(publicKeyPEM []byte, plainText string) string {
	var block *pem.Block

	if block, _ = pem.Decode(publicKeyPEM); block == nil || block.Type != "PUBLIC KEY" { //privatKeyData is in string format
		log.Fatal("No valid PEM data found")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		//fmt.Println(err)
	}

	label := []byte("orders")
	decryptedPKCS1v15, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey.(*rsa.PublicKey),
		[]byte(plainText),
		label,
	)
	if err != nil {
		//fmt.Println(err)
	}

	return base64.StdEncoding.EncodeToString(decryptedPKCS1v15)
}

func DecryptSessionKey(privateKeyPEM []byte, encData []byte) string {

	encData, _ = base64.StdEncoding.DecodeString(string(encData))

	var block *pem.Block

	if block, _ = pem.Decode(privateKeyPEM); block == nil || block.Type != "PRIVATE KEY" { //privatKeyData is in string format
		log.Fatal("No valid PEM data found")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		//fmt.Println(err)
	}
	label := []byte("orders")
	decryptedPKCS1v15, _ := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		encData,
		label,
	)
	return string(decryptedPKCS1v15)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	if len(encrypt) > 0 {
		padding := encrypt[len(encrypt)-1]
		return encrypt[:len(encrypt)-int(padding)]
	}
	return nil
}

func EncryptData(sessionKey string, plainText string) string {
	var err error
	var block cipher.Block
	// Pad key
	if block, err = aes.NewCipher([]byte(sessionKey)); err != nil {
		return ""
	}

	ciphertext := make([]byte, len(plainText)+(aes.BlockSize-len(plainText)%aes.BlockSize))

	cbc := cipher.NewCBCEncrypter(block, PROD_IV)
	plainData := PKCS5Padding([]byte(plainText), aes.BlockSize)
	cbc.CryptBlocks(ciphertext, plainData)

	//fmt.Printf("%s", ciphertext)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func EncryptPin(plainText string) string {
	// hash := sha256.Sum256([]byte(plainText))

	key := []byte("secret")
	h := hmac.New(sha256.New, key)
	h.Write([]byte(plainText))
	fmt.Println(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))

	// return hex.EncodeToString(hash[:])
}

func DecryptData(sessionKey string, encData []byte) string {
	var err error
	var block cipher.Block

	encData, _ = base64.StdEncoding.DecodeString(string(encData))

	// Pad key
	if block, err = aes.NewCipher([]byte(sessionKey)); err != nil {
		return ""
	}

	plainText := make([]byte, len(encData))

	cbc := cipher.NewCBCDecrypter(block, PROD_IV)
	cbc.CryptBlocks(plainText, encData)

	// fmt.Printf("%s", ciphertext)

	return string(PKCS5Trimming(plainText))
}
