package validate

import (
	"fmt"
	"net/http"
	"unicode"
)

func ValidateUsername(username string) (bool, int, error) {
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return false, http.StatusOK, fmt.Errorf("в имени пользователя разрешены только буквенно-цифровые символы")
		}
	}
	if 5 <= len(username) && len(username) <= 50 {
		return true, http.StatusOK, nil
	}
	return false, http.StatusOK, fmt.Errorf("длина имени пользователя должна быть больше 4 и меньше 51 символа")
}
