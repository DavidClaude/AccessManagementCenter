package utils

import (
	"strings"
	"crypto/md5"
	"fmt"
	"encoding/base64"
)

func StringKick(src string, args ... string) (result string) {
	result = src
	for _, s := range args {
		result = strings.Replace(result, s, "", -1)
	}
	return result
}

func Base64EncodeBytesToString(data []byte) (str64 string) {
	str64 = base64.StdEncoding.EncodeToString(data)
	return str64
}

func Base64DecodeStringToBytes(str64 string) (data []byte, err error) {
	data, err = base64.StdEncoding.DecodeString(str64)
	return data, err
}

func Base64DecodeBytesToString(data []byte) (str string, err error)  {
	var strBytes []byte
	_, err = base64.StdEncoding.Decode(strBytes, data)
	if err != nil {
		return "", err
	}
	return string(strBytes), nil
}

/**
 * 计算字符串md5值
 */
func Md5Encode(str string) (strMd5 string) {
	strMd5Bytes := md5.Sum([]byte(str))
	strMd5 = fmt.Sprintf("%x", strMd5Bytes)
	return strMd5
}
