package infrastructure

import (
	"github.com/gin-gonic/gin"

	// uc "github.com/set2002satoshi/8-4/interfaces/controllers/user"
	// bc "github.com/set2002satoshi/8-4/interfaces/controllers/blog"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: ":8080",
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	// user := r.Gin.Group("/api")
	// {
	// 	usersController := uc.NewUsersController(r.DB)
	// 	user.GET("/users", func(c *gin.Context) { usersController.FindAll(c) })
	// }
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
