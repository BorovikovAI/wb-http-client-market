package database

import (
	"encoding/json"
	"github.com/google/uuid"
	"wb/rest-api/pkg/logging"
)

var _ Model = &Client{}
var _ Model = &Market{}

type Model interface {
	Marshal(*logging.Logger) ([]byte, error)
	GetList(*Database) ([]Model, error)
	Insert(*Database) (string, error)
	Update(*Database) error
	Delete(*Database) error
}

type Client struct {
	Id               *string `json:"id,omitempty"`
	LastName         string  `json:"last_name"`
	FirstName        string  `json:"first_name"`
	Patronymic       string  `json:"patronymic"`
	Age              *int    `json:"age,omitempty"`
	RegistrationDate string  `json:"registration_date"`
}

func (c Client) Marshal(logger *logging.Logger) ([]byte, error) {
	data, err := json.Marshal(c)
	if err != nil {
		logger.Warningf("failed to marshal client: %v", err)
		return nil, err
	}

	return data, nil
}

const (
	getByLastNameClient = "SELECT id, last_name, first_name, patronymic, age, registration_date FROM clients WHERE last_name = $1"
	insertClient        = "INSERT INTO clients (id, last_name, first_name, patronymic, age, registration_date) VALUES ($1, $2, $3, $4, $5,$6)"
	updClient           = "UPDATE clients SET last_name=$1, first_name=$2, patronymic=$3, age=$4, registration_date=$5 WHERE id=$6"
	deleteClient        = "DELETE FROM clients WHERE id = $1"
)

func (c Client) GetList(db *Database) ([]Model, error) {
	result := make([]Model, 0)
	rows, err := db.Conn.Query(getByLastNameClient, c.LastName)
	if err != nil {
		db.logger.Warningf("failed to get client by LastName: %v", err)
		return nil, err
	}

	for rows.Next() {
		client := Client{}
		err = rows.Scan(
			&client.Id,
			&client.LastName,
			&client.FirstName,
			&client.Patronymic,
			&client.Age,
			&client.RegistrationDate)
		if err != nil {
			db.logger.Warningf("failed to scan row: %v", err)
			return nil, err
		}
		result = append(result, client)
	}

	return result, nil
}

func (c Client) Insert(db *Database) (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		db.logger.Warningf("failed to get new uuid: %v", err)
		return "", err
	}

	id := uid.String()
	_, err = db.Conn.Exec(insertClient,
		id,
		c.LastName,
		c.FirstName,
		c.Patronymic,
		c.Age,
		c.RegistrationDate)
	if err != nil {
		db.logger.Warningf("failed to insert client: %v", err)
		return "", err
	}

	return id, err
}

func (c Client) Update(db *Database) error {
	_, err := db.Conn.Exec(updClient,
		c.LastName,
		c.FirstName,
		c.Patronymic,
		c.Age,
		c.RegistrationDate,
		c.Id)
	if err != nil {
		db.logger.Warningf("failed to update client: %v", err)
	}

	return err
}

func (c Client) Delete(db *Database) error {
	_, err := db.Conn.Exec(deleteClient, c.Id)
	if err != nil {
		db.logger.Warningf("failed to delete client: %v", err)
	}

	return err
}

type Market struct {
	Id      *string `json:"id,omitempty"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Active  bool    `json:"active"`
	Owner   *string `json:"owner,omitempty"`
}

func (m Market) Marshal(logger *logging.Logger) ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		logger.Warningf("failed to marshal market: %v", err)
		return nil, err
	}
	return data, nil
}

const (
	getByNameMarket = "SELECT id, name, address, active, owner FROM markets WHERE name = $1"
	insertMarket    = "INSERT INTO markets (id, name, address, active, owner) VALUES ($1, $2, $3, $4, $5)"
	updMarket       = "UPDATE markets SET name=$1, address=$2, active=$3, owner=$4 WHERE id=$5"
	deleteMarket    = "DELETE FROM markets WHERE id = $1"
)

func (m Market) GetList(db *Database) ([]Model, error) {
	result := make([]Model, 0)
	rows, err := db.Conn.Query(getByNameMarket, m.Name)
	if err != nil {
		db.logger.Warningf("failed to get market by Name: %v", err)
		return nil, err
	}

	for rows.Next() {
		market := Market{}
		err = rows.Scan(
			&market.Id,
			&market.Name,
			&market.Address,
			&market.Active,
			&market.Owner)
		if err != nil {
			db.logger.Warningf("failed to scan row: %v", err)
			return nil, err
		}
		result = append(result, market)
	}

	return result, nil
}

func (m Market) Insert(db *Database) (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		db.logger.Warningf("failed to get new uuid: %v", err)
		return "", err
	}

	id := uid.String()
	_, err = db.Conn.Exec(insertMarket,
		id,
		m.Name,
		m.Address,
		m.Active,
		m.Owner)
	if err != nil {
		db.logger.Warningf("failed to insert market: %v", err)
		return "", err
	}

	return id, err
}

func (m Market) Update(db *Database) error {
	_, err := db.Conn.Exec(updMarket,
		m.Name,
		m.Address,
		m.Active,
		m.Owner,
		m.Id)
	if err != nil {
		db.logger.Warningf("failed to update market: %v", err)
	}

	return err
}

func (m Market) Delete(db *Database) error {
	_, err := db.Conn.Exec(deleteMarket, m.Id)
	if err != nil {
		db.logger.Warningf("failed to delete market: %v", err)
	}

	return err
}
