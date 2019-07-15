package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"sort"
	"time"
)

var (
	rander = rand.New(rand.NewSource(time.Now().UnixNano()))
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

// xxx--
func MapSignMD5(dataMap map[string]string, key string) string {
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

	return MD5String(writer.String())
}
