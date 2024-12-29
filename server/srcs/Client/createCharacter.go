package client

import "encoding/json"

type data struct {
	class	string
	name	string
}

func (this *Client) createCharacter() {
	var data data

	json.Unmarshal([]byte(this.body), &data)
}
