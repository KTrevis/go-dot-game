package client

func (this *Client) inGame() error {
	this.sendNearChunks()
	return nil
}
