package validations

import (
	"errors"
	"net/mail"
	"strings"

	verrors "github.com/akrck02/whisper/sdk/errors"
)

func ValidatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New(verrors.PasswordEmptyMessage)
	}

	if len(password) < 16 {
		return errors.New(verrors.PasswordShortMessage)
	}

	if !strings.ContainsAny(password, "123456789") {
		return errors.New(verrors.PasswordNoNumericMessage)
	}

	if !strings.ContainsAny(password, "*¡!¿?$%&/()@#~¬") {
		return errors.New(verrors.PasswordNoSpecialCharacterMessage)
	}

	if password == strings.ToLower(password) {
		return errors.New(verrors.PasswordNoUppercaseCharacterMessage)
	}

	if password == strings.ToUpper(password) {
		return errors.New(verrors.PasswordNoLowercaseCharacterMessage)
	}

	return nil
}

func ValidateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New(verrors.EmailEmptyMessage)
	}

	_, err := mail.ParseAddress(email)
	if nil != err {
		return err
	}

	return nil
}
