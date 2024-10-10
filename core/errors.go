package core

import "errors"

var (
  InvalidEnvironmentVariableError = errors.New("Faltando variavel de ambiente")
  InvalidParametersError = errors.New("Parametros utilizados são invalidos")
  EntityAlreadyExistsError = errors.New("Item ja existe")
)
