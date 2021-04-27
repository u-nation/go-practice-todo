package presentation

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"

	validatorv10 "github.com/go-playground/validator/v10"
)

var validator = validatorv10.New()

type Validator interface {
	Validate() error
}

func BindAndValidate(ctx echo.Context, v interface{}) error {
	if err := ctx.Bind(v); err != nil {
		return xerrors.Errorf("failed to Bind Param: %w", err)
	}

	if err := validator.Struct(v); err != nil {
		return xerrors.Errorf("field validation failed: %w", err)
	}

	if v, ok := v.(Validator); ok {
		if err := v.Validate(); err != nil {
			return xerrors.Errorf("struct validation failed: %w", err)
		}
	}

	return nil
}
