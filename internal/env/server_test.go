package env_test

import (
	"github.com/oleksandr-pol/simple-go-service/internal/env"
)

func ExampleNewServer() {
	dbConf := &env.DbConfig{DbHostName: "localhost", DbHostPort: 5432, DbUserName: "test", DbPassword: "test", DbName: "test"}
	env.NewServer(dbConf)
}
