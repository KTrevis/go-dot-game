package client

import (
	"encoding/json"
	"errors"
	"fmt"
	gamemaps "server/maps"
)

func (this *Client) getMap() error {
	if !this.authenticated {
		const msg = "tried to get map without authenticating"
		fmt.Printf("%s %s", this.getFormattedIP(), msg)
		return errors.New(msg)
	}

	var data struct {
		Map string
	}

	err := json.Unmarshal([]byte(this.body), &data)

	if err != nil {
		this.disconnect("invalid payload")
		return fmt.Errorf("unmarshal failed: %s", this.body)
	}

	gamemap := gamemaps.GetMap(data.Map)

	if gamemap == nil {
		const msg =  "could not find requested map"
		this.sendMessage("GET_MAP", &Dictionary{
			"error": msg,
		})
		return errors.New(msg)
	}

	this.sendMessage("GET_MAP", &Dictionary{
		"map": gamemap,
	})
	return nil
}
