package client

import (
	"errors"
	"server/classes"
)

func (this *Client) getClassList() error {
	if this.authenticated == false {
		const msg = "tried to get class list without authenticating"
		this.sendMessage("GET_CLASS_LIST", &Dictionary{"error": msg})
		return errors.New(msg)
	}

	this.sendMessage("GET_CLASS_LIST", &Dictionary{"classes": classes.GetClassesName()})
	return nil
}
