package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrNotFound          = errors.New("not found in database")
	ErrNotEnoughGoldEggs = errors.New("not enough gold eggs")
)

type ConnectionSettings struct {
	Host          string
	Port          string
	DatabaseName  string
	User          string
	Password      string
	IsReadReplica bool
}

func OpenSQLConnections(dbConnectionSettings []ConnectionSettings) ([]*sql.DB, error) {
	result := make([]*sql.DB, 0)

	for _, setting := range dbConnectionSettings {
		dbConnection, err := sql.Open(
			"postgres", fmt.Sprintf(
				"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
				setting.Host,
				setting.Port,
				setting.DatabaseName,
				setting.User,
				setting.Password,
			),
		)
		if err != nil {
			return nil, err
		}

		result = append(result, dbConnection)
	}

	return result, nil
}

func NormalizeNotFound(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}

	return err
}
