package gamemaps

import (
	"encoding/json"
	"fmt"
	"os"
	"server/utils"
	"strings"
)

type MapData struct {
	Data []int
	Map [][]utils.Vector2
	Height int
	Width int
}

func getData(mapName string) *MapData {
	var data MapData

	mapName = fmt.Sprintf("./maps/%s.tmj", mapName)
	Data, err := os.ReadFile(mapName)
	
	if err != nil {
		fmt.Printf("MapData getData: failed to read file")
		return nil
	}

	str := string(Data)
	str = str[strings.Index(str, "[") + 1:]
	str = str[:strings.LastIndex(str, "]")]
	str = str[:strings.LastIndex(str, "]")]

	err = json.Unmarshal([]byte(str), &data)

	if err != nil {
		fmt.Printf("MapData getData: unmarshal failed: %s", str)
		return nil
	}
	return &data
}

func NewMapData(mapName string) *MapData {
	data := getData(mapName)

	if data == nil {
		return nil
	}

	for i := 0; i < len(data.Data); {
		line := make([]utils.Vector2, 0, data.Width)

		for range data.Width {
			line = append(line, utils.Vector2{
				X: (data.Data[i] - 1) % 11,
				Y: (data.Data[i] - 1) / 11,
			})
			i++
			if i >= len(data.Data) {
				break
			}
		}
		data.Map = append(data.Map, line)

		if i >= len(data.Data) {
			break
		}
	}
	return data
}
