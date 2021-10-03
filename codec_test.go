package codec

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(`Md5("md5") =`, Md5("md5"))
}

func TestJson(t *testing.T) {
	data := map[string]interface{}{"abc": "123", "xyz": 123}
	fmt.Println("data:", data)

	dst := JsonEncode(data)
	fmt.Println("JsonEncode:", dst)

	var src interface{}
	JsonDecode(dst, &src)
	fmt.Println("JsonDecode:", src)
}

func TestBase64(t *testing.T) {
	str := "base64"
	fmt.Println("str:", str)

	dst := Base64Encode([]byte(str))
	fmt.Println("Base64Encode:", dst)

	src := Base64Decode(dst)
	fmt.Println("Base64Decode:", string(src))
}

func TestAes(t *testing.T) {
	key := "1bc29b36f623ba82aaf6724fd3b16718"
	str := `{"abc":"123","xyz":123}`
	fmt.Println("str:", str, "key:", key)

	dst := AesEncrypt([]byte(str), key)
	fmt.Println("AesEncrypt:", string(dst))

	src := AesDecrypt(dst, key)
	fmt.Println("AesDecrypt:", string(src))
}