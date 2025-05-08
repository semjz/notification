package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"notification/domain/notify"
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

func ValidatePayloadStructure(body []byte) *json.Decoder {
	var raw map[string]json.RawMessage
	json.Unmarshal(body, &raw)
	delete(raw, "type")

	body, _ = json.Marshal(raw)

	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	return decoder
}

func ValidatePayloadFields(payload notify.NotifyPayload) []string {
	if err := Validate.Struct(payload); err != nil {
		errs := make([]string, 0)
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, err := range err.(validator.ValidationErrors) {
				errs = append(errs, err.Field()+" is "+err.Tag())
			}
		} else {
			errs = append(errs, err.Error())
		}
		return errs
	}
	return nil
}

func StructToMap(v interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &result)
	return result, err
}
