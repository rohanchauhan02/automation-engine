package util

import (
	"context"
	"encoding/json"
	"io"
	"os"

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

// ReadJSON is a function to read json in path into pointer
func ReadJSON(path string, pointer interface{}) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(byteValue), pointer)
	return nil

}
