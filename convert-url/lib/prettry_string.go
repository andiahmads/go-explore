package lib

import (
	"bytes"
	"encoding/json"
)

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer

	// joining the string by separator

	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}

	return prettyJSON.String(), nil

}
