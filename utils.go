package paysdk

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

// md5 digest in string
func _md5String(plain string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(plain))
	cipher := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipher)
}

// sha256 digest in string
func _sha256String(plain string) string {
	sha256Ctx := sha256.New()
	sha256Ctx.Write([]byte(plain))
	cipher := sha256Ctx.Sum(nil)
	return hex.EncodeToString(cipher)
}

// xxx-----------------------
func GenerateMapSign(m map[string]string, signType string, key string) (string, error) {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	// 对keys排序
	sort.Strings(keys)

	var writer bytes.Buffer
	for _, key := range keys {
		if m[key] == "" {
			continue
		}

		writer.WriteString(key)
		writer.WriteString("=")
		writer.WriteString(m[key])
		writer.WriteString("&")
	}
	writer.WriteString("key=" + key)

	if signType == SignTypeMD5 {
		return _md5String(writer.String()), nil
	} else if signType == SignTypeHMACSHA256 {
		return _sha256String(writer.String()), nil
	}

	return "", fmt.Errorf("invalid sign_type: %s", signType)
}

func RandomString(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lettersLength := len(letters)

	result := make([]rune, ln)

	for i := range result {
		result[i] = letters[rander.Intn(lettersLength)]
	}

	return string(result)
}
