package encryption

import (
	"crypto/aes"
	"errors"
	"fmt"
)

// =================== ECB ======================
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(key)
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)

	for i := len(origData); i < len(plain); i++ {
		plain[i] = 0
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

//func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte, err error) {
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println(r)
//			err = errors.New("recover")
//		}
//	}()
//	cipher, _ := aes.NewCipher(key)
//	decrypted = make([]byte, len(encrypted))
//	//
//	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
//		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
//	}
//	decrypted = removeNullValue(decrypted)
//	return
//}

func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("AesDecryptECB recover error: %v\n", r)
			err = errors.New("AesDecryptECB recover")
		}
	}()
	cipher, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("----->AesDecryptECB error: %v", err)
	}
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	if len(decrypted) > 0 {
		return removeNullValue(decrypted), nil
	}
	return []byte{}, err
}

func removeNullValue(slice []byte) []byte {
	var output []byte
	for _, element := range slice {
		// 0 代表 nil，不能包含nil值
		if element != 0 { //if condition satisfies add the elements in new slice
			output = append(output, element)
		}
	}
	return output //slice with no nil-values
}
