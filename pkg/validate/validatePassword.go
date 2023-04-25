package validate

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrLength8                = errors.New("пароль должен быть не менее 8 символов")
	ErrMissingNumeric         = errors.New("пароль должен состоять как минимум из 2 цифр")
	ErrMissingAlphabetic      = errors.New("пароль должен содержать не менее 4 буквенных символов")
	ErrMissingSpecial         = errors.New("пароль должен содержать как минимум 1 специальный символ")
	ErrMissingUppercase       = errors.New("пароль должен содержать как минимум 1 символ в верхнем регистре")
	ErrEqualAdjacent          = errors.New("пароль не должен содержать соседние символы с одинаковым значением")
	ErrInvalidCharCombination = errors.New("пароль содержит недопустимую комбинацию символов: 'asdf', 'qwerty', '1234' or '98765'")
	ErrConsecutive            = errors.New("пароль содержит значения, следующие друг за другом, 1234, 3456, abcd, efgh")
)

func Validate(password string) error {
	if len(password) < 8 {
		return ErrLength8
	}

	var numerics int
	var alphabetic int
	var specials int
	var uppercase int
	var buffer [4]rune
	var equal int
	var consecutive int

	for i, p := range password {
		buffer[i%4] = p

		if isNumericChar(p) {
			numerics++
		}
		if hasUppercase := isAlphabeticUppercase(p); hasUppercase || isAlphabeticLowercase(p) {
			if hasUppercase {
				uppercase++
			}
			alphabetic++
		}
		if isSpecialChar(p) {
			specials++
		}

		for _, c := range buffer {
			if isSpecialChar(p) {
				continue
			}
			if p == c {
				equal++
			}
			if p-c <= 3 && p-c >= -3 {
				consecutive++
			}
		}

		if equal == 4 {
			return ErrEqualAdjacent
		}
		if consecutive == 4 {
			return ErrConsecutive
		}
		equal = 0
		consecutive = 0
	}
	if numerics < 2 {
		return ErrMissingNumeric
	}
	if alphabetic < 4 {
		return ErrMissingAlphabetic
	}
	if uppercase < 1 {
		return ErrMissingUppercase
	}
	if specials < 1 {
		return ErrMissingSpecial
	}

	if strings.Index(password, "asdf") != -1 {
		return ErrInvalidCharCombination
	}

	if strings.Index(password, "qwerty") != -1 {
		return ErrInvalidCharCombination
	}

	return nil
}

func isAlphabeticLowercase(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func isAlphabeticUppercase(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func isNumericChar(c rune) bool {
	return c >= '0' && c <= '9'
}

func isSpecialChar(c rune) bool {
	return (c >= '!' && c <= '/') || (c >= ':' && c <= '@') || (c >= '[' && c <= '`') || (c >= '{' && c <= '~')
}

func ValidatePassword(password string) (bool, int, error) {
	err := Validate(password)

	if err != nil {
		return false, http.StatusOK, fmt.Errorf("%s", err)
	}

	return true, http.StatusOK, nil
}
