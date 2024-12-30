package api

import (
	"log"
	"server/database"

	"github.com/gin-gonic/gin"
)

type RegistrationForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
}

func Register(c *gin.Context, db *database.DB) {
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

	err = user.CreateAccount(db)

	if err != nil {
		c.HTML(401, "index.html", gin.H{"msg": err.Error()})
		return
	}

	log.Printf("account %s created", user.Username)
	c.HTML(200, "index.html", gin.H{"msg": "account created"})
}
