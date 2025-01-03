package chunks

import "server/utils"

var SPAWN = utils.Vector2i{
	X: 0,
	Y: 0,
}

const CHUNK_SIZE = 50

type Chunk struct {
	Tiles		[][]utils.Vector2i
	Position	utils.Vector2i
	Name		string
}

type ChunkHandler struct {
	Chunks map[utils.Vector2i]*Chunk
}

func NewChunkHandler() *ChunkHandler {
	handler := &ChunkHandler{
		Chunks: make(map[utils.Vector2i]*Chunk),
	}
	chunks := [...]string {
		"test",
		"test1",
		"test2",
	}

	for _, v := range chunks {
		chunk := loadChunk(v)

		if _, ok := handler.Chunks[chunk.Position]; ok {
			panic("trying to create two chunks at the same position")
		}

		handler.Chunks[chunk.Position] = chunk
	}
	return handler
}
