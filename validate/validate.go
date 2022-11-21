package validate

import (
	"encoding/json"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func ValidateAgainstSchema(input []byte, schema []byte) error {
	sch, err := jsonschema.CompileString("schema.json", string(schema))

	if err != nil {
		return err
	}

	var v interface{}

	if err = json.Unmarshal(input, &v); err != nil {
		return err
	}

	return sch.Validate(v)

}
