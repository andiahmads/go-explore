package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/joncalhoun/qson"
)

func ConvertToJosn(url string) {
	newErr := errors.New("string url not valid || ex:clientid=rdfb0afc6c0640&kp")
	b, err := qson.ToJSON(url)
	if err != nil {
		fmt.Println(newErr)
	}
	p, _ := prettyprint(b)
	fmt.Println(string(p))
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
