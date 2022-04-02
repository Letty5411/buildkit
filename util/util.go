package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func LettyPrettyDump(obj interface{})  {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(obj)
	fmt.Printf("LETTYHINT %s\n", buf.String())
}