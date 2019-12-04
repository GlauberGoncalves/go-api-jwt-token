package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "nickname") {
		return errors.New("Nickname Já Utilizado")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email já utilizado")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Senha incorreta")
	}
	return errors.New("Detalhes incorretos")
}
