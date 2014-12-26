package exportAsciinemaJson

import (
	"encoding/json"
)

type Line struct {
	Cells []Cell
}

func (line Line) MarshalJSON() ([]byte, error) {
	return json.Marshal(line.Cells)
}

func (line Line) Equal(oldLine Line) bool {
	if len(line.Cells) != len(oldLine.Cells) {
		return false
	}
	for i, cell := range line.Cells {
		if cell != oldLine.Cells[i] {
			return false
		}
	}
	return true
}
