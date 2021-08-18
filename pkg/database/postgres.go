package database

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
)

func PostgresHealth() bool {
	db, err := sql.Open("postgres", postgresStringConnection())
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	if err != nil {
		return false
	}

	return true
}

func postgresStringConnection() string {
	user, _ := viper.Get("POSTGRES_USER").(string)
	pwd, _ := viper.Get("POSTGRES_PWD").(string)
	name, _ := viper.Get("POSTGRES_NAME").(string)

	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pwd, name)
}
