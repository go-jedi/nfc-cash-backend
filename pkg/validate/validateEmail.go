package validate

import (
	"fmt"
	"net/http"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailverifier.NewVerifier()
)

func ValidateEmail(email string) (bool, int, error) {
	ret, err := verifier.Verify(email)
	if err != nil {
		return false, http.StatusOK, fmt.Errorf("проверить адрес электронной почты не удалось, ошибка, %s", err)
	}
	if !ret.Syntax.Valid {
		return false, http.StatusOK, fmt.Errorf("синтаксис адреса электронной почты недействителен")
	}
	return true, http.StatusOK, nil
}
