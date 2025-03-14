package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
)

// 패키지 내부용 전역 변수 (소문자)
var (
	trans  ut.Translator
	engine *validator.Validate
)

// 패키지 초기화
func init() {
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)

	var found bool
	trans, found = uni.GetTranslator("en")
	if !found {
		panic("English validator of translator not found")
	}

	engine = validator.New(validator.WithRequiredStructEnabled())

	if err := en_translations.RegisterDefaultTranslations(engine, trans); err != nil {
		panic(fmt.Sprintf("Failed to register translations: %v", err))
	}
}

func Validate(s any) error {
	err := engine.Struct(s)

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		var errorMessages []string
		for _, e := range validationErrors {
			errorMessages = append(errorMessages, e.Translate(trans))
		}
		return apperr.New("ERR_VALIDATION_FAILED").WithStatus(400).WithData(errorMessages)
	}

	return nil
}
