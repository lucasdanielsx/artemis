package database

import (
	"artemis/pkg/util"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func PostgresHealth() bool {
	db, err := sql.Open("postgres", postgresStringConnection())
	if err != nil {
		util.HandleError("Can not open a postgres connection: ", err)
		return false
	}

	errDB := db.Close()
	if errDB != nil {
		util.HandleError("Can not close a postgres connection: ", err)
		return false
	}

	return true
}

func postgresStringConnection() string {
	host, _ := viper.Get("POSTGRES_HOST").(string)
	port, _ := viper.Get("POSTGRES_PORT").(string)
	user, _ := viper.Get("POSTGRES_USER").(string)
	pwd, _ := viper.Get("POSTGRES_PWD").(string)
	name, _ := viper.Get("POSTGRES_NAME").(string)

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, name)
}
