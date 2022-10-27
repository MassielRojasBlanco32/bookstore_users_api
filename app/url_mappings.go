package app

import (
	"github.com/MassielRojasBlanco32/bookstore_users_api/controllers/ping"
	"github.com/MassielRojasBlanco32/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// curl -X POST localhost:8080/users -d '{"id": 123, "first_name": "Massiel"}'
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	// Get all users with a matching status
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)
}
