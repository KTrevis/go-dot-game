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
	Position	utils.Vector2i
	Tiles		[]utils.Vector2i
	Name		string
}

type ChunkHandler struct {
	chunks map[utils.Vector2i]*Chunk
}

func NewChunkHandler() *ChunkHandler {
	handler := &ChunkHandler{}
	return handler
}

func (this *ChunkHandler) AddNewMap(chunk *Chunk) {
	_, ok := this.chunks[chunk.Position]

	if ok {
		panic("tried to replace a chunk with another one")
	}

	this.chunks[chunk.Position] = chunk
}
