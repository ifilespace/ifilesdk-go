package util

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
)

func GetSignToken(keyid, keysecret string, now int64) string {
	token := Sha1(keyid + keysecret + strconv.FormatInt(now, 10))
	return token
}

// 对字符串进行sha1 计算
func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
