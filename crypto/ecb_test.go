package crypto

import (
	"bytes"
	"crypto/aes"
	"testing"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Test_ecb(t *testing.T) {
	txt := "abcdefgh12345678Key"
	key := "abcdefgh12345678"

	dest := make([]byte, (len(txt)/len(key)+1)*len(key))

	aesCipher, _ := aes.NewCipher([]byte(key))
	encrypter := NewECBEncrypter(aesCipher)
	ciphertext := PKCS7Padding([]byte(txt), aes.BlockSize)
	encrypter.CryptBlocks(dest, ciphertext)

	decrypter := NewECBDecrypter(aesCipher)
	p := make([]byte, len(dest))
	decrypter.CryptBlocks(p, dest)

	// de
	b, err := DecryptDoPKCS5UnPadding(dest, key)
	if err != nil || string(b) != txt {
		t.Error(string(b), txt, err)
	}
}
