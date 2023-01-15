package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"wb/rest-api/internal/config"
	"wb/rest-api/pkg/logging"
)

var _ Storage = &Database{}

type Storage interface {
	GetList(Model) ([]Model, error)
	Insert(Model) (string, error)
	Delete(Model) error
	Update(Model) error
}

type Database struct {
	Conn   *sql.DB
	logger *logging.Logger
}

func NewDatabaseConnection(dbConfig config.Database, logger *logging.Logger) (Storage, error) {
	logger.Info("new db connection:", dbConfig)

	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		logger.Warningf("failed to open sql: %v", err)
		return nil, fmt.Errorf("open sql error: %v", err)
	}

	if err = db.Ping(); err != nil {
		logger.Warningf("failed to ping db: %v", err)
		return nil, fmt.Errorf("ping error: %v", err)
	}

	return &Database{
		Conn:   db,
		logger: logger,
	}, nil
}

func (db *Database) GetList(mdl Model) ([]Model, error) {
	return mdl.GetList(db)
}

func (db *Database) Insert(mdl Model) (string, error) {
	return mdl.Insert(db)
}

func (db *Database) Update(mdl Model) error {
	return mdl.Update(db)
}

func (db *Database) Delete(mdl Model) error {
	return mdl.Delete(db)
}
