package client

import (
	"context"
	"errors"
)

func (this *Client) getCharacterList() error {
	if !this.authenticated {
		const msg = "tried to get characters without authenticating"
		this.sendMessage("GET_CHARACTER_LIST", &Dict{"error": msg})
		return errors.New(msg)
	}

	conn, _ := this.manager.DB.Acquire(context.TODO())
	defer conn.Release()

	const query = "SELECT (name, class, level, xp) FROM characters WHERE user_id=$1;"
	rows, _ := conn.Query(context.TODO(), query, this.user.ID)

	type data struct {
		Name	string
		Class	string
		Level 	int
		XP		int
	}

	var msg []data

	for rows.Next() {
		var data data
		rows.Scan(&data)
		msg = append(msg, data)
	}

	if msg == nil {
		msg = []data{}
	}

	this.sendMessage("GET_CHARACTER_LIST", &Dict{
		"characterList": msg,
	})

	return nil
}
