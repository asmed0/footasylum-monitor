package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"math/rand"
	"time"
)

var (
	charset     = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	oauthSecret = []byte{55, 66, 53, 68, 55, 66, 67, 50, 52, 66, 53, 67, 52, 69, 51, 65, 56, 48, 70, 66, 66, 67, 50, 65, 49, 49, 53, 54, 69, 52, 51, 55}
)

func AESEncrypt(src string, key []byte, generatedIV string) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(generatedIV))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func generateIV(length int) string {
	return StringWithCharset(length, charset)
}
