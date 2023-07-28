package modules

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(str string) string {
	hash := sha256.New()

	// 写入要计算哈希的数据
	hash.Write([]byte(str))

	// 计算哈希值
	hashed := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	sha256Str := hex.EncodeToString(hashed)
	return sha256Str
}
