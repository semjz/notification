package setup

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var Validate = validator.New()

func init() {
	// Register custom validation function
	err := Validate.RegisterValidation("phonenumber", func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()
		re := regexp.MustCompile(`^\+?\d{12}$`)
		return re.MatchString(phone)
	})
	if err != nil {
		return
	}
}
