package dsignatures

import (
	"asymmetric-encr/internal/gears/encryption/asymmetric"
	"asymmetric-encr/internal/gears/hash"
	"fmt"
	"slices"
)

func Sign(msg []byte, encryptor *asymmetric.RSAEncryptor) ([]byte, error) {
	digest := hash.SHA256Hash(msg)
	msgAndDigest := append(msg, digest...)
	return encryptor.Encrypt(msgAndDigest)
}

func Verify(signature []byte, encryptor *asymmetric.RSAEncryptor) (bool, error) {
	decrSignature, err := encryptor.Decrypt(signature)
	if err != nil {
		return false, err
	}

	decrHash := make([]byte, 0)
	decrData := make([]byte, 0)

	decrHash = append(decrHash, decrSignature[len(decrSignature)-32:]...)
	decrData = append(decrData, decrSignature[:len(decrSignature)-32]...)

	decryptedMsgHash := hash.SHA256Hash(decrData)
	fmt.Println(decrHash)
	fmt.Println(decryptedMsgHash)
	return slices.Equal(decryptedMsgHash, decrHash), nil
}
