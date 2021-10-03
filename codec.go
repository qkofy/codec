package codec

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/qkofy/log"
)

var IsZero bool

func init() {
	IsZero = true
}

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func JsonEncode(i interface{}) string {
	buf := bytes.NewBuffer([]byte{})
	jsn := json.NewEncoder(buf)

	jsn.SetEscapeHTML(false)

	err := jsn.Encode(i)

	if err != nil {
		log.Error(err)
	}

	return buf.String()
}

func JsonDecode(s string, i interface{}) {
	buf := bytes.NewBufferString(s)
	jsn := json.NewDecoder(buf)

	jsn.UseNumber()

	err := jsn.Decode(i)

	if err != nil {
		log.Error(err)
	}
}

func Base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Base64Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)

	if err != nil {
		log.Error(err)

		return nil
	}

	return data
}

func zeroPadding(data []byte, size int) []byte {
	tmp := bytes.Repeat([]byte{0}, size - len(data) % size)

	return append(data, tmp...)
}

func pkcsPadding(data []byte, size int) []byte {
	num := size - len(data) % size
	tmp := bytes.Repeat([]byte{byte(num)}, num)

	return append(data, tmp...)
}

func padding(data []byte, size int) []byte {
	if IsZero {
		return zeroPadding(data, size)
	} else {
		return pkcsPadding(data, size)
	}
}

func unPadding(data []byte) []byte {
	size := len(data)

	return data[:(size - int(data[size-1]))]
}

func aesCipher(key string) (cipher.Block, []byte) {
	keys := []byte(key)

	block, err := aes.NewCipher(keys)

	if err != nil {
		log.Error(err)
	}

	return block, keys
}

func AesEncrypt(src []byte, key string) []byte {
	block, keys := aesCipher(key)

	size := block.BlockSize()
	data := padding(src, size)
	mode := cipher.NewCBCEncrypter(block, keys[:size])
	text := make([]byte, len(data))

	mode.CryptBlocks(text, data)

	return text
}

func AesDecrypt(src []byte, key string) []byte {
	block, keys := aesCipher(key)

	size := block.BlockSize()
	mode := cipher.NewCBCDecrypter(block, keys[:size])
	data := make([]byte, len(src))

	mode.CryptBlocks(data, src)

	return unPadding(data)
}