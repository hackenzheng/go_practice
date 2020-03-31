package main

import (
	"crypto/md5"
	"fmt"
)

const _SALT string = "123"

func SecretMd5(mediaId string) string {

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(mediaId))
	result := Md5Inst.Sum([]byte(""))
	md5str := fmt.Sprintf("%x", result) //将[]byte转成16进制
	return md5str
}


func main() {
	mediaId := "123"

	result := SecretMd5(mediaId + _SALT)
	fmt.Println(result)

}
