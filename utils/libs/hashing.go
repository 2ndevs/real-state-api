package libs

import (
	"errors"
	"main/core"

	"golang.org/x/crypto/bcrypt"
)

type Hashing struct{}

func (Hashing) EncryptPassword(value string) (*string, error) {
	valueAsBytes := []byte(value)

	bytes, err := bcrypt.GenerateFromPassword(valueAsBytes, 10)
	if err != nil {
		return nil, errors.Join(core.PasswordEncryptionError, err)
	}

	hash := string(bytes)
	return &hash, nil
}

func (Hashing) IsValidPassword(password string, hash string) error {
	passwordAsBytes := []byte(password)
	hashAsBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(
		hashAsBytes,
		passwordAsBytes,
	)
  if err != nil {
    return core.InvalidPasswordError
  }

	return err
}
