package pinot

import (
	"github.com/spaghettifunk/go-pinot/pkg/models"
	"gopkg.in/go-playground/validator.v9"
)

// ValidateRole implements validator.Func
func ValidateRole(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case string(models.BROKER):
		return true
	case string(models.SERVER):
		return true
	default:
		return false
	}
}
