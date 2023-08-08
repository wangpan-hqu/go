package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
)

const DefaultBase64KeyString = "Yk5DanhkbmdkaUphZGpjUQ=="

//PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)

	if length == 0 {
		return nil, errors.New("加密字符串错误！")

	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

func EncryptToBase64(originData string) string {
	return base64.StdEncoding.EncodeToString([]byte(originData))
}

func DecryptFromBase64(encryptData string) string {
	return base64.StdEncoding.EncodeToString([]byte(encryptData))
}

/*
Aes解密，key经过base64加密的
*/
func AesDecryptForBase64Key(encryptData string, base64KeyString string) (string, error) {
	return AesDecrypt(encryptData, DecryptFromBase64(base64KeyString))
}

/*
Aes加密，key经过base64加密的
*/
func AesEncryptForBase64Key(originString string, base64KeyString string) (string, error) {
	return AesEncrypt(originString, DecryptFromBase64(base64KeyString))
}

/*
AES加密<br/>
Date：2021年1月29日<br/>
@return 16进制加密字符<br/>
*/
func AesEncrypt(originString string, keyString string) (string, error) {
	originBytes := []byte(originString)
	key := []byte(keyString)

	//创建加密算法实例
	block, err := aes.NewCipher(key)

	if err != nil {
		log.Println("create chipper error")
		return "", err
	}

	blockSize := block.BlockSize()
	originBytes = PKCS7Padding(originBytes, blockSize)
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encryptData := make([]byte, len(originBytes))
	blocMode.CryptBlocks(encryptData, originBytes)
	hexString := hex.EncodeToString(encryptData)
	return hexString, nil
}

/*
AES解密<br/>
Date：2021年1月29日<br/>
@param hexString 加密后的16进制字符串<br/>
@return 解密字符<br/>
*/
func AesDecrypt(hexString string, keyString string) (string, error) {
	key := []byte(keyString)
	block, err := aes.NewCipher(key)

	if err != nil {
		log.Println("create chipper error")
		return "", err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	encryptData, err := hex.DecodeString(hexString)

	if nil != err {
		log.Printf("the hexString %s to []byte error", hexString)
		return "", err
	}

	originData := make([]byte, len(encryptData))
	blockMode.CryptBlocks(originData, encryptData)
	originData, err = PKCS7UnPadding(originData)

	if err != nil {
		return "", err
	}

	return string(originData[:]), err
}
