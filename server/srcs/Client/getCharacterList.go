package client

import (
	"context"
	"errors"
	"fmt"
	"server/database"
)

func (this *Client) getCharacterList() error {
	if !this.authenticated {
		const msg = "tried to get characters without authenticating"
		this.sendMessage(&Dictionary{"error": msg})
		return errors.New(msg)
	}

	conn, _ := this.manager.DB.Acquire(context.TODO())
	defer conn.Release()

	var characters []database.Character

	const query = "SELECT (name, level, class) FROM users WHERE user_id=$1;"
	rows, _ := conn.Query(context.TODO(), query, this.user.ID)
	rows.Scan(&characters)
	fmt.Printf("characters: %v\n", characters)
	return nil
}
