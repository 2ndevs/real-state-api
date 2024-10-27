package core

import (
	"errors"
	"net/http"
)

var (
	InvalidEnvironmentVariableError      = errors.New("Faltando alguma variável de ambiente")
	InvalidParametersError               = errors.New("Parâmetros utilizados são inválidos")
	EntityAlreadyExistsError             = errors.New("Item já foi cadastrado")
	MissingAuthorizationTokenError       = errors.New("Usuário não autenticado")
	MissingRefreshTokenError             = errors.New("Não foi possível atualizar o token")
	AuthorizationTokenExpiredError       = errors.New("Token de autorização expirado")
	RefreshTokenExpiredError             = errors.New("Token de autorização expirado")
	PasswordEncryptionError              = errors.New("Não foi possível encriptar a senha")
	InvalidPasswordError                 = errors.New("Dados de login incorretos")
	InvalidEmailError                    = errors.New("Dados de login incorretos")
	UnableToPersistToken                 = errors.New("Não foi possível criar o token, tente novamente")
	UnableToPersistTokenButEntityCreated = errors.New("Usuario criado, mas não foi criado um token de login. Tente logar novamente")
	FallbackError                        = errors.New("Ocorreu um erro no servidor")
	NotFoundError                        = errors.New("Não foi possível encontrar o item solicitado")
	ImageUploadError                     = errors.New("Não foi possivel dar upload na imagem")
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
	case UnableToPersistToken, UnableToPersistTokenButEntityCreated:
		{
			http.Error(write, errMessage, http.StatusInternalServerError)
		}
	case NotFoundError:
		{
			http.Error(write, errMessage, http.StatusNotFound)
		}

	default:
		{
			http.Error(write, FallbackError.Error(), http.StatusInternalServerError)
		}
	}
}
