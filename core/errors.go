package core

import (
	"errors"
	"net/http"
)

var (
	InvalidEnvironmentVariableError = errors.New("Faltando alguma variavel de ambiente") // FIXME: Nathan fix this shit, I don't have pt layout setup
	InvalidParametersError          = errors.New("Parametros utilizados são invalidos")
	EntityAlreadyExistsError        = errors.New("Item ja foi cadastrado")
	MissingAuthorizationTokenError  = errors.New("Usuario nao autenticado")
	MissingRefreshTokenError        = errors.New("Nao foi possivel atualizar token")
	AuthorizationTokenExpiredError  = errors.New("Token de autorização expirado")
	RefreshTokenExpiredError        = errors.New("Token de autorização expirado")
)

func HandleHTTPStatus(write http.ResponseWriter, err error) {
	errMessage := err.Error()

	switch err.(error) {
	case MissingAuthorizationTokenError:
		{
			http.Error(write, errMessage, 499) // ESRI: Token required (unofficial)
		}
	case AuthorizationTokenExpiredError:
		{
			http.Error(write, errMessage, http.StatusUnauthorized)
		}
	case MissingRefreshTokenError:
		{
			http.Error(write, errMessage, 498) // ESRI: Invalid token (unofficial)
		}
	case InvalidParametersError:
		{
			http.Error(write, errMessage, http.StatusBadRequest)
		}
	case EntityAlreadyExistsError:
		{
			http.Error(write, errMessage, http.StatusConflict)
		}

	default:
		{
			http.Error(write, errMessage, http.StatusInternalServerError)
		}
	}
}
