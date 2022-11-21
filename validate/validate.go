package validate

import (
	"encoding/json"
	"reflect"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func ValidateAgainstSchema(input []byte, schema []byte) error {
	sch, err := jsonschema.CompileString("schema.json", string(schema))

	if err != nil {
		return err
	}

	var m map[string]interface{}

	if err = json.Unmarshal(input, &m); err != nil {
		return err
	}

	removeNulls(m)

	return sch.Validate(m)

}

// https://gist.github.com/ribice/074ad38d9f2fc5c88b20663659988d19#file-trim-nulls-go
func removeNulls(m map[string]interface{}) {
	val := reflect.ValueOf(m)
	for _, e := range val.MapKeys() {
		v := val.MapIndex(e)
		if v.IsNil() {
			delete(m, e.String())
			continue
		}
		switch t := v.Interface().(type) {
		// If key is a JSON object (Go Map), use recursion to go deeper
		case map[string]interface{}:
			removeNulls(t)
		}
	}
}
