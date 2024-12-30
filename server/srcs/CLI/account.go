package cli

import (
	"context"
	"server/database"
)

func (this *CLI) createAccount() {
	if len(this.split) != 4 {
		this.sendMessage("command usage: account create <username> <password>")
		return
	}

	var user database.User
	user.Username = this.split[2]
	user.Password = this.split[3]

	conn, _ := this.Manager.DB.Acquire(context.TODO())
	defer conn.Release()

	err := user.CreateAccount(conn)

	if err != nil {
		this.sendMessage("error: " + err.Error())
		return
	}
	this.sendMessage("account created")
}

func (this *CLI) account() {
	m := map[string]func() {
		"create": this.createAccount,
	}

	if f := this.validArg(m); f != nil {
		f()
	}
}
