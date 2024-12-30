package cli

import (
	"server/database"
	"strings"
)

func (this *CLI) createAccount(split []string) {
	if len(split) != 4 {
		this.sendMessage("command usage: account create <username> <password>")
		return
	}

	var user database.User
	user.Username = split[2]
	user.Password = split[3]
	err := user.CreateAccount(this.Manager.DB)

	if err != nil {
		this.sendMessage("error: " + err.Error())
		return
	}
	this.sendMessage("account created")
}

func getArgumentList(m map[string]func([]string)) string {
	str := "\n"

	for key := range m {
		str += "- " + key + "\n"
	}

	str = strings.TrimSuffix(str, "\n")
	return str
}

func (this *CLI) account(split []string) {
	m := make(map[string]func([]string))
	m["create"] = this.createAccount

	if len(split) < 2 {
		msg := "invalid argument. arguments possible: "
		msg += getArgumentList(m)
		this.sendMessage(msg)
		return
	}

	fn, ok := m[split[1]]
	if ok {
		fn(split)
	} else {
		msg := "invalid argument. arguments possible: "
		msg += getArgumentList(m)
		this.sendMessage(msg)
	}
}
