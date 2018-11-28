package utils

import (
	"strings"
	"crypto/md5"
	"fmt"
)

func StringKick(src string, args ... string) (result string) {
	result = src
	for _, s := range args {
		result = strings.Replace(result, s, "", -1)
	}
	return result
}

/**
 * 计算字符串md5值
 */
func Md5Encode(str string) (strMd5 string) {
	strMd5Bytes := md5.Sum([]byte(str))
	strMd5 = fmt.Sprintf("%x", strMd5Bytes)
	return strMd5
}
