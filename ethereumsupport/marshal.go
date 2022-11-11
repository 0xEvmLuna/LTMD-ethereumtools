package ethereumsupport

import (
	"encoding/json"
)

func Marshal(any interface{}) ([]byte, error) {
	bytesData, err := json.Marshal(any)
	return bytesData, err
}
