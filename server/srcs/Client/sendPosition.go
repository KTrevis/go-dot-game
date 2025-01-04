package client

import (
	"server/database/Character"
)

func (this *Client) sendPosition(character *character.Character) {
	this.manager.mutex.RLock()
	defer this.manager.mutex.RUnlock()
	chunks := character.GetSurroundingChunks()
	characterChunk := character.GetChunk()

	for _, curr := range this.manager.Clients {
		if curr.character == nil || curr.character == character {
			continue
		}

		currChunk := curr.character.GetChunk()
		if *currChunk == *characterChunk {
			// check if the two players are in the same chunk
			curr.sendMessage("MOVE_PLAYER", &Dict{
				"character": character,
			})
			this.sendMessage("MOVE_PLAYER", &Dict{
				"character": curr.character,
			})
			continue
		}

		// check every surrounding chunks
		for _, chunk := range *chunks {
			// if the character is not in the current surrounding chunk,
			// check the next one
			if chunk != *currChunk {
				continue
			}

			curr.sendMessage("MOVE_PLAYER", &Dict{
				"character": character,
			})
			this.sendMessage("MOVE_PLAYER", &Dict{
				"character": curr.character,
			})
			break
		}
	}
}
