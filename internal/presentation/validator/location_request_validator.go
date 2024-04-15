package validators

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func LatitudeValidation(fl validator.FieldLevel) bool {
	if latitude, ok := fl.Field().Interface().(string); ok {
		match := regexp.MustCompile(
			`^(-?(90(\.0{1,7})?|[1-8]?\d(\.\d{1,7})?))$`,
		).MatchString(latitude)

		return match && len(latitude) <= 11
	}
	return false
}

func LongitudeValidation(fl validator.FieldLevel) bool {
	if longitude, ok := fl.Field().Interface().(string); ok {
		match := regexp.MustCompile(
			`^(-?(180(\.0{1,7})?|1[0-7]\d(\.\d{1,7})?|\d{1,2}(\.\d{1,7})?))$`,
		).MatchString(longitude)

		return match && len(longitude) <= 12
	}
	return false
}

func RegisterCustomValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("latitude", LatitudeValidation)
		v.RegisterValidation("longitude", LongitudeValidation)
	}
}
