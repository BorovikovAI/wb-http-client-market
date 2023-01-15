package database

import (
	"github.com/miladibra10/vjson"
	"wb/rest-api/pkg/logging"
)

var _ Validator = Client{}
var _ Validator = Market{}

type Validator interface {
	ValidateForList([]byte, *logging.Logger) error
	ValidateForCreate([]byte, *logging.Logger) error
	ValidateForUpdate([]byte, *logging.Logger) error
	ValidateForDelete([]byte, *logging.Logger) error
}

func (c Client) ValidateForList(data []byte, logger *logging.Logger) error {
	clientSchema := vjson.NewSchema(
		vjson.String("last_name").Required().MinLength(1).MaxLength(20),
	)

	err := clientSchema.ValidateBytes(data)
	if err != nil {
		logger.Warningf("validation fail: %v", err)
		return err
	}

	return nil
}

func (c Client) ValidateForCreate(data []byte, logger *logging.Logger) error {
	clientSchema := vjson.NewSchema(
		vjson.String("last_name").Required().MinLength(1).MaxLength(20),
		vjson.String("first_name").Required().MinLength(1).MaxLength(20),
		vjson.String("patronymic").Required().MinLength(1).MaxLength(20),
		vjson.Integer("age").Range(1, 120),
		vjson.String("registration_date").Required().MinLength(10).MaxLength(10),
	)

	err := clientSchema.ValidateBytes(data)
	if err != nil {
		logger.Warningf("validation fail: %v", err)
		return err
	}

	return nil
}

func (c Client) ValidateForUpdate(data []byte, logger *logging.Logger) error {
	clientSchema := vjson.NewSchema(
		vjson.String("id").Required().MinLength(1),
		vjson.String("last_name").Required().MinLength(1).MaxLength(20),
		vjson.String("first_name").Required().MinLength(1).MaxLength(20),
		vjson.String("patronymic").Required().MinLength(1).MaxLength(20),
		vjson.Integer("age").Range(1, 120),
		vjson.String("registration_date").Required().MinLength(10).MaxLength(10),
	)

	err := clientSchema.ValidateBytes(data)
	if err != nil {
		logger.Warningf("validation fail: %v", err)
		return err
	}

	return nil
}

func (c Client) ValidateForDelete(data []byte, logger *logging.Logger) error {
	clientSchema := vjson.NewSchema(
		vjson.String("id").Required().MinLength(1),
	)

	err := clientSchema.ValidateBytes(data)
	if err != nil {
		logger.Warningf("validation fail: %v", err)
		return err
	}

	return nil
}

func (m Market) ValidateForList(data []byte, logger *logging.Logger) error {
	marketSchema := vjson.NewSchema(
		vjson.String("name").Required().MinLength(1).MaxLength(20),
	)

	err := marketSchema.ValidateBytes(data)
	if err != nil {
		logger.Warningf("validation fail: %v", err)
		return err
	}

	return nil
}

func (m Market) ValidateForCreate(data []byte, logger *logging.Logger) error {
	marketSchema := vjson.NewSchema(
		vjson.String("name").Required().MinLength(1).MaxLength(20),
		vjson.String("address").Required().MinLength(1).MaxLength(50),
		vjson.Boolean("active").Required(),
		vjson.String("owner").MinLength(1).MaxLength(20),
	)

	err := marketSchema.ValidateBytes(data)
	if err != nil {
		logger.Warningf("validation fail: %v", err)
		return err
	}

	return nil
}

func (m Market) ValidateForUpdate(data []byte, logger *logging.Logger) error {
	marketSchema := vjson.NewSchema(
		vjson.String("id").Required().MinLength(1),
		vjson.String("name").Required().MinLength(1).MaxLength(20),
		vjson.String("address").Required().MinLength(1).MaxLength(50),
		vjson.Boolean("active").Required(),
		vjson.String("owner").MinLength(1).MaxLength(20),
	)

	err := marketSchema.ValidateBytes(data)
	if err != nil {
		logger.Warningf("validation fail: %v", err)
		return err
	}

	return nil
}

func (m Market) ValidateForDelete(data []byte, logger *logging.Logger) error {
	marketSchema := vjson.NewSchema(
		vjson.String("id").Required().MinLength(1),
	)

	err := marketSchema.ValidateBytes(data)
	if err != nil {
		logger.Warningf("validation fail: %v", err)
		return err
	}

	return nil
}
