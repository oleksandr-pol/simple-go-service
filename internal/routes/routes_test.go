package routes_test

import (
	"github.com/oleksandr-pol/simple-go-service/internal/env"
	"github.com/oleksandr-pol/simple-go-service/internal/routes"
)

func ExampleRegisterRoutes() {
	dbConf := &env.DbConfig{DbHostName: "localhost", DbHostPort: 5432, DbUserName: "test", DbPassword: "test", DbName: "test"}
	s := env.NewServer(dbConf)
	routes.RegisterRoutes(s, "5s")
}
