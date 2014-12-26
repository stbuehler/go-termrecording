package exportAsciinemaJson

import (
	"encoding/json"
	// tsm "github.com/stbuehler/go-termrecording/libtsm"
)

type Cell struct {
	Text  string
	Brush Brush
}

func (cell Cell) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{cell.Text, cell.Brush})
}
