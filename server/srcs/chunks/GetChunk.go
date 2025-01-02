package chunks

import (
	"server/database"
	"server/utils"
)

func GetChunk(db *database.DB, chunkPos utils.Vector2i) *Chunk {
	chunk := &Chunk{}
	const query = "SELECT ()"
	// row := db.QueryRow(context.TODO())
	return chunk
}
