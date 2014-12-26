package exportAsciinemaJson

import (
	"encoding/json"
	tsm "github.com/stbuehler/go-termrecording/libtsm"
)

type Brush struct {
	Foreground int
	Background int
	Bold       bool
	Underline  bool
	Inverse    bool
	Blink      bool
}

func (brush Brush) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{}
	if brush.Foreground >= 0 {
		m["fg"] = brush.Foreground
	}
	if brush.Background >= 0 {
		m["bg"] = brush.Background
	}
	if brush.Bold {
		m["bold"] = brush.Bold
	}
	if brush.Underline {
		m["underline"] = brush.Underline
	}
	if brush.Inverse {
		m["inverse"] = brush.Inverse
	}
	if brush.Blink {
		m["blink"] = brush.Blink
	}
	return json.Marshal(m)
}

var (
	/* "const" */
	rgbLevels = [...]uint8{0x00, 0x5f, 0x87, 0xaf, 0xd7, 0xff}
)

func getRGBLevel(value uint8) int {
	for i, v := range rgbLevels {
		if v == value {
			return i
		}
	}
	return -1
}

func getRGBIndex(color tsm.RGB) int {
	if color.Red == color.Green && color.Green == color.Blue && color.Red >= 8 && 0 == (int(color.Red)-8)%10 {
		// grey shades: map rgb 8, 18, ..., 248 to index 232..256
		return 232 + (int(color.Red)-8)/10
	} else {
		// try to map certain levels to index 0..5, and combine them to a
		// range of 216 colors from 16..231
		r := getRGBLevel(color.Red)
		g := getRGBLevel(color.Green)
		b := getRGBLevel(color.Blue)
		if r >= 0 && g >= 0 && b >= 0 {
			return 16 + 36*r + 6*g + b
		}
	}
	println("invalid color: ", Stringify(color))
	return -1
}

func getColorIndex(color tsm.Color) int {
	if color.Code < 0 {
		return getRGBIndex(color.RGB)
	} else if color.Code < 16 {
		return int(color.Code)
	} else {
		return -1
	}
}

func MakeBrush(attr *tsm.ScreenAttr) Brush {
	if attr == nil {
		return Brush{
			Foreground: -1,
			Background: -1,
		}
	}
	return Brush{
		Foreground: getColorIndex(attr.Foreground),
		Background: getColorIndex(attr.Background),
		Bold:       attr.Bold,
		Underline:  attr.Underline,
		Inverse:    attr.Inverse,
		Blink:      attr.Blink,
	}
}
