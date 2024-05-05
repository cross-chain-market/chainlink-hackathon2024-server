package validator

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
	"os"
	"reflect"

	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/google/uuid"
)

var Validator = validator.New()

func InitValidator() {
	if err := Validator.RegisterValidation("notblank", validators.NotBlank); err != nil {
		slog.Error("failed to register validation", slog.String("error", err.Error()))
		os.Exit(1)
	}

	Validator.RegisterCustomTypeFunc(validateUUID, uuid.UUID{})
}

func validateUUID(field reflect.Value) any {
	if valuer, ok := field.Interface().(uuid.UUID); ok {
		return valuer.String()
	}

	return nil
}
