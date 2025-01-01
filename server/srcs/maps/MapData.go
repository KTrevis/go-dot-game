package gamemaps

import (
	"encoding/json"
	"os"
	"strings"
)

type Vector2 struct {
	x int
	y int
}

type MapData struct {
	Data []int
	Map [][]Vector2
	Height int
	Width int
}

func getFile(path string) *MapData {
	var data MapData

	Data, _ := os.ReadFile(path)
	str := string(Data)
	str = str[strings.Index(str, "[") + 1:]
	str = str[:strings.LastIndex(str, "]")]
	str = str[:strings.LastIndex(str, "]")]
	json.Unmarshal([]byte(str), &data)
	return &data
}

func NewMapData(path string) *MapData {
	data := getFile(path)

	for i := 0; i < len(data.Data); {
		line := make([]Vector2, 0, data.Width)

		for range data.Width {
			line = append(line, Vector2{
				x: (data.Data[i] - 1) % 11,
				y: (data.Data[i] - 1) / 11,
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
