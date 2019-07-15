package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var (
	rander = rand.New(rand.NewSource(time.Now().UnixNano()))
)

const (
	SignTypeMD5        = "MD5"
	SignTypeHMACSHA256 = "HMAC-SHA256"
)

func RandomString(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lettersLength := len(letters)

	result := make([]rune, ln)

	for i := range result {
		result[i] = letters[rander.Intn(lettersLength)]
	}

	return string(result)
}

// MD5String md5 digest in string
func MD5String(plain string) string {
	cipher := MD5([]byte(plain))
	return hex.EncodeToString(cipher)
}

// MD5 md5 digest
func MD5(plain []byte) []byte {
	md5Ctx := md5.New()
	md5Ctx.Write(plain)
	cipher := md5Ctx.Sum(nil)
	return cipher
}

func Sha256(plain []byte) []byte {
	sha256Ctx := sha256.New()
	sha256Ctx.Write(plain)
	cipher := sha256Ctx.Sum(nil)
	return cipher
}

func Sha256String(plain string) string {
	cipher := Sha256([]byte(plain))
	return hex.EncodeToString(cipher)
}

// xxx--
func GenerateMapSign(dataMap map[string]string, signType string, key string) (string, error) {
	keys := make([]string, 0, len(dataMap))

	for key := range dataMap {
		keys = append(keys, key)
	}

	// 对keys排序
	sort.Strings(keys)

	var writer bytes.Buffer
	for _, key := range keys {
		if dataMap[key] == "" {
			continue
		}

		writer.WriteString(key)
		writer.WriteString("=")
		writer.WriteString(dataMap[key])
		writer.WriteString("&")
	}
	writer.WriteString("key=" + key)

	if signType == SignTypeMD5 {
		return MD5String(writer.String()), nil
	} else if signType == SignTypeHMACSHA256 {
		return Sha256String(writer.String()), nil
	}

	return "", fmt.Errorf("invalid sign_type: %s", signType)
}
