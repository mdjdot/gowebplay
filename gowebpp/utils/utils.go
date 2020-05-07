package utils

import (
	"crypto/md5"
	"fmt"
	"gowebpp/confs"
	"io"
	"mime/multipart"
)

const securityStr = "1qaz@WSX"

// MD5 加密字符串
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum([]byte(securityStr)))
}

// MD5File 加密文件
func MD5File(f multipart.File) (string, error) {
	h := md5.New()
	_, err := io.Copy(h, f)
	if err != nil {
		confs.Logger.Printf("加密文件出错，错误：%v", err)
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
