package application

import (
	"main/core"
	"main/domain/entities"
	"main/utils/libs"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type RefreshTokenService struct {
	Database *gorm.DB
	Parser   libs.JWT
}

type RefreshTokenResponse struct {
	Token        *string
	RefreshToken *string
}

func (self RefreshTokenService) Execute(refreshToken string) (*RefreshTokenResponse, error) {
	oldToken, err := self.Parser.Parse(refreshToken)
	if err != nil {
		switch err.(error) {
		case jwt.ErrTokenExpired:
			{
				return nil, core.AuthorizationTokenExpiredError
			}

		default:
			{
				return nil, core.MissingAuthorizationTokenError
			}
		}
	}

	sub, err := oldToken.Claims.GetSubject()
	if err != nil {
		return nil, err
	}

	id, err := strconv.ParseInt(sub, 32, 10)
	if err != nil {
		return nil, err
	}

	var user entities.User
	query := self.Database.Model(&entities.User{}).Where("id = ?", id)

	databaseResponse := query.First(&user)
	if databaseResponse.Error != nil {
		return nil, core.InvalidParametersError
	}

	newToken, err := self.Parser.Generate(libs.CreateJWTParams{
		Sub:  user.ID,
		Role: user.RoleID,
		Time: time.Now().Add(time.Hour * 2).Unix(),
	})
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := self.Parser.Generate(libs.CreateJWTParams{
		Sub:  user.ID,
		Role: user.RoleID,
		Time: time.Now().Add(time.Hour * 24).Unix(),
	})
	if err != nil {
		return nil, err
	}

  databaseResponse = self.Database.Model(&entities.User {}).Where("id = ?", id).Update("refresh_token", newRefreshToken)
  if databaseResponse.Error != nil {
    return nil, core.UnableToPersistToken
  }

	response := RefreshTokenResponse{
		Token:        newToken,
		RefreshToken: newRefreshToken,
	}

	return &response, nil
}
