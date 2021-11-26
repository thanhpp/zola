package accountcipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

func New(key string) (entity.AccountCipher, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	return &accountCipherImpl{
		cipher: c,
		gcm:    gcm,
		key:    []byte(key),
	}, nil
}

type accountCipherImpl struct {
	cipher cipher.Block
	gcm    cipher.AEAD
	key    []byte
}

func (c accountCipherImpl) genNonce() ([]byte, error) {
	nonce := make([]byte, c.gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return nonce, nil
}

func (c accountCipherImpl) Encrypt(raw string) (string, error) {
	nonce, err := c.genNonce()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(c.gcm.Seal(nonce, nonce, []byte(raw), nil)), nil
}

func (c accountCipherImpl) Decrypt(encryptedHex string) (string, error) {
	encrypted, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return "", err
	}

	if len(encrypted) < c.gcm.NonceSize() {
		return "", errors.New("invalid ecrypted text")
	}

	nonce, ciphertext := encrypted[:c.gcm.NonceSize()], encrypted[c.gcm.NonceSize():]
	plaintext, err := c.gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
