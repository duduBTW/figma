package app

import (
	"encoding/json"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Float32Map map[float32]float32

func (fm Float32Map) MarshalJSON() ([]byte, error) {
	temp := make(map[string]float32, len(fm))
	for k, v := range fm {
		temp[strconv.FormatFloat(float64(k), 'f', -1, 32)] = v
	}
	return json.Marshal(temp)
}

func (fm *Float32Map) UnmarshalJSON(b []byte) error {
	temp := make(map[string]float32)
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}
	result := make(map[float32]float32, len(temp))
	for k, v := range temp {
		f, err := strconv.ParseFloat(k, 32)
		if err != nil {
			return err
		}
		result[float32(f)] = v
	}
	*fm = result
	return nil
}

type Float32Vec2Map map[float32]rl.Vector2

func (fm Float32Vec2Map) MarshalJSON() ([]byte, error) {
	temp := make(map[string]rl.Vector2, len(fm))
	for k, v := range fm {
		temp[strconv.FormatFloat(float64(k), 'f', -1, 32)] = v
	}
	return json.Marshal(temp)
}

func (fm *Float32Vec2Map) UnmarshalJSON(b []byte) error {
	temp := make(map[string]rl.Vector2)
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}
	result := make(map[float32]rl.Vector2, len(temp))
	for k, v := range temp {
		f, err := strconv.ParseFloat(k, 32)
		if err != nil {
			return err
		}
		result[float32(f)] = v
	}
	*fm = result
	return nil
}
