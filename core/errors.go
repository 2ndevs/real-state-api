package core

import (
	"errors"
	"net/http"
)

var (
	InvalidEnvironmentVariableError      = errors.New("Faltando alguma variavel de ambiente") // FIXME: Nathan fix this shit, I don't have pt layout setup
	InvalidParametersError               = errors.New("Parametros utilizados são invalidos")
	EntityAlreadyExistsError             = errors.New("Item ja foi cadastrado")
	MissingAuthorizationTokenError       = errors.New("Usuario nao autenticado")
	MissingRefreshTokenError             = errors.New("Nao foi possivel atualizar token")
	AuthorizationTokenExpiredError       = errors.New("Token de autorização expirado")
	RefreshTokenExpiredError             = errors.New("Token de autorização expirado")
	PasswordEncryptionError              = errors.New("Não foi possivel encriptar a senha")
	InvalidPasswordError                 = errors.New("Senha incorreta")
	InvalidEmailError                    = errors.New("Email não foi encontrado")
	UnableToPersistToken                 = errors.New("Não foi possivel criar token, tente novamente")
	UnableToPersistTokenButEntityCreated = errors.New("Foi criado, mas não foi criado um token de login, tente logar novamente")
	FallbackError                        = errors.New("Ocorreu um erro no servidor")
)

func HandleHTTPStatus(write http.ResponseWriter, err error) {
	errMessage := err.Error()

	switch err.(error) {
	case
		InvalidEmailError,
		InvalidPasswordError,
		MissingAuthorizationTokenError,
		MissingRefreshTokenError:
		{
			http.Error(write, errMessage, http.StatusUnauthorized)
		}
	case AuthorizationTokenExpiredError:
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
			http.Error(write, FallbackError.Error(), http.StatusInternalServerError)
		}
	}
}
