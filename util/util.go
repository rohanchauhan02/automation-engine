package util

import (
	"context"

	"github.com/fgrosse/goldi"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/rohanchauhan02/automation-engine/shared/config"
)

type CustomApplicationContext struct {
	context.Context
	Container       *goldi.Container
	SharedConfig    config.ImmutableConfigInterface
	PostgresSession *gorm.DB
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (c *CustomValidator) Validate(i interface{}) error {
	return c.Validator.Struct(i)
}

func DefaultValidator() *CustomValidator {
	return &CustomValidator{Validator: validator.New()}
}
