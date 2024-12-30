package client

import (
	"errors"
	"server/classes"
)

func (this *Client) getClassList() error {
	if this.authenticated == false {
		const msg = "tried to get class list without authenticating"
		this.sendMessage(&Dictionary{"error": msg})
		return errors.New(msg)
	}

	this.sendMessage(&Dictionary{"classes": classes.GetClassesName()})
	return nil
}
