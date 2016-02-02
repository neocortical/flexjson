package flexjson

import (
	"encoding/json"
	"strconv"
)

type Example struct {
	Desc     string          `json:"desc"`
	Sprocket SprocketWrapper `json:"sprocket,omitempty"`
}

type Sprocket struct {
	ID    int    `json:"id"`
	Size  string `json:"size,omitempty"`
	Gears int    `json:"gears,omitempty"`
}

type SprocketWrapper struct {
	Sprocket
	Partial bool `json:"partial,omitempty"`
}

func (w *SprocketWrapper) UnmarshalJSON(data []byte) error {
	if id, err := strconv.Atoi(string(data)); err == nil {
		w.ID = id
		w.Partial = true
		return nil
	}
	return json.Unmarshal(data, &w.Sprocket)
}

func (w SprocketWrapper) MarshalJSON() ([]byte, error) {
	if w.ID == 0 {
		return []byte("null"), nil
	}
	if w.Partial {
		return []byte(strconv.Itoa(w.ID)), nil
	}
	return json.Marshal(w.Sprocket)
}
