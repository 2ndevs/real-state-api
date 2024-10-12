package libs

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
  InvalidTokenError = errors.New("Token isn't valid token")
  UnableToParseTokenError = errors.New("Unable to parse token")
)

type JWT struct{}

type CreateJWTParams struct {
	Sub  uint
	Role uint
}

func (JWT) Generate(params CreateJWTParams) (*string, error) {
	secret := os.Getenv("JWT_SECRET")
	if len(secret) == 0 {
		log.Panic("Missing JWT Secret environment variable")
	}

	constructor := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":   params.Sub,
		"role": params.Role,
    "exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	parsed, err := constructor.SignedString(secret)
	if err != nil {
		return nil, errors.Join(UnableToParseTokenError, err)
	}

	return &parsed, nil
}

func (JWT) Parse(value string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	if len(secret) == 0 {
		log.Panic("Missing JWT Secret environment variable")
	}

  token, err := jwt.Parse(value, func(constructor *jwt.Token) (any, error) {
    return secret, nil
  })
  if err != nil {
    return nil, err
  }

  if !token.Valid {
    return nil, InvalidTokenError
  }

	return token, nil
}
