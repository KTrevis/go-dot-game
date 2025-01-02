package gamemaps

import (
	"encoding/json"
	"fmt"
	"os"
	"server/utils"
)

func getMapData(mapName string) []int {
	var data []int

	mapName = fmt.Sprintf("./maps/%s.tmj", mapName)
	file, err := os.ReadFile(mapName)
	
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

func GetMap(mapName string) [][]utils.Vector2i {
	data := getMapData(mapName)
	var vec [][]utils.Vector2i
	const MAP_SIZE = 50

	if data == nil {
		return nil
	}

	for i := 0; i < len(data); {
		line := make([]utils.Vector2i, 0, MAP_SIZE)

		for range MAP_SIZE {
			line = append(line, utils.Vector2i{
				X: (data[i] - 1) % 11,
				Y: (data[i] - 1) / 11,
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
