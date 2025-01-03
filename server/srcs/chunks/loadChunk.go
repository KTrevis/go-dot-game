package chunks

import (
	"encoding/json"
	"fmt"
	"os"
	"server/utils"
)

type MapData struct {
	Name		string
	Position	utils.Vector2i
	Tiles 		[]int
}

func readFile(filename string) *MapData {
	data := &MapData{}

	filename = fmt.Sprintf("./chunks/%s.tmj", filename)
	file, err := os.ReadFile(filename)
	
	if err != nil {
		fmt.Printf("MapData getData: failed to read file")
		return nil
	}

	err = json.Unmarshal(file, &data)

	if err != nil {
		fmt.Printf("MapData getData: unmarshal failed: %s", string(file))
		return nil
	}
	return data
}

func loadChunk(filename string) *Chunk {
	const MAP_SIZE = 50
	const SHEET_SIZE = 11

	data := readFile(filename)

	if data == nil {
		return nil
	}

	chunk := &Chunk{
		Name: data.Name,
		Position: data.Position,
		Tiles: make([][]utils.Vector2i, 0, MAP_SIZE),
	}


	for i := 0; i < len(data.Tiles); {
		line := make([]utils.Vector2i, 0, MAP_SIZE)

		for range MAP_SIZE {
			line = append(line, utils.Vector2i{
				X: (data.Tiles[i] - 1) % SHEET_SIZE,
				Y: (data.Tiles[i] - 1) / SHEET_SIZE,
			})
			i++
			if i >= len(data.Tiles) {
				break
			}
		}
		chunk.Tiles = append(chunk.Tiles, line)

		if i >= len(data.Tiles) {
			break
		}
	}
	return chunk
}
