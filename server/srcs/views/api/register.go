package api

import (
	"context"
	"log"
	"server/database"

	"github.com/gin-gonic/gin"
)

type RegistrationForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
}

func Register(c *gin.Context, db *database.DBPool) {
	var form RegistrationForm

	err := c.Bind(&form)
	if err != nil {
		c.HTML(401, "index.html", gin.H{"msg": "invalid payload"})
		return
	}

	if form.Password != form.ConfirmPassword {
		c.HTML(401, "index.html", gin.H{"msg": "passwords do not match"})
		return
	}

	var user database.User
	user.Username = form.Username
	user.Password = form.Password

	conn, _ := db.Acquire(context.TODO())
	defer conn.Release()

	err = user.CreateAccount(conn)

	if err != nil {
		c.HTML(401, "index.html", gin.H{"msg": err.Error()})
		return
	}

	log.Printf("account %s created", user.Username)
	c.HTML(200, "index.html", gin.H{"msg": "account created"})
}
