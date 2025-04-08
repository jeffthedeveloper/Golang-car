package validations

import (
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	h "github.com/jhonyaltoe/go-car-shop/api/helpers"
)

func Validate(c interface{}) *h.CustomError {
	en := en.New()
	var uni *ut.UniversalTranslator = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(c)

	if err != nil {
		return &h.CustomError{
			Message:    err.(validator.ValidationErrors)[0].Translate(trans),
			StatusCode: http.StatusBadRequest,
			Method:     "validations.Validate",
			Err:        err,
		}
	}

	return nil
}
