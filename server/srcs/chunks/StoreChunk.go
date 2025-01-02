package chunks

import (
	"encoding/json"
	"fmt"
	"os"
	"server/utils"
)

func readFile(filename string) []int {
	var data struct {
		Tiles []int
	}

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
	return data.Tiles
}

func StoreChunk(filename string) [][]utils.Vector2i {
	data := readFile(filename)
	var vec [][]utils.Vector2i
	const MAP_SIZE = 50
	const SHEET_COLUMNS = 11

	if data == nil {
		return nil
	}

	for i := 0; i < len(data); {
		line := make([]utils.Vector2i, 0, MAP_SIZE)

		for range MAP_SIZE {
			line = append(line, utils.Vector2i{
				X: (data[i] - 1) % SHEET_COLUMNS,
				Y: (data[i] - 1) / SHEET_COLUMNS,
			})
			i++
			if i >= len(data) {
				break
			}
		}

		vec = append(vec, line)
		if i >= len(data) {
			break
		}
	}
	return vec
}
