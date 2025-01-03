package chunks

import "server/utils"

const (
	TOP = iota
	BOTTOM
	LEFT
	RIGHT
	CHUNK_SIZE = 50
)

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
