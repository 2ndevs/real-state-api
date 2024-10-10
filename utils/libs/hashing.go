package libs

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	InvalidPasswordError = errors.New("Senha incorreta")
	PasswordEncryptionError = errors.New("Não foi possivel encriptar a senha")
)

type Hashing struct{}

func (Hashing) EncryptPassword(value string) (*string, error) {
	valueAsBytes := []byte(value)

	bytes, err := bcrypt.GenerateFromPassword(valueAsBytes, 14)
	if err != nil {
		return nil, errors.Join(PasswordEncryptionError, err)
	}

	hash := string(bytes)
	return &hash, nil
}

func (Hashing) IsValidPassword(password string, hash string) error {
	passwordAsBytes := []byte(password)
	hashAsBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(
		passwordAsBytes,
		hashAsBytes,
	)
  if err != nil {
    return errors.Join(InvalidPasswordError, err)
  }

	return err
}
