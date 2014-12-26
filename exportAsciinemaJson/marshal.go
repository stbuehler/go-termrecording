package exportAsciinemaJson

import (
	"encoding/json"
)

func Stringify(v interface{}) string {
	str, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(str)
}
