package codegen

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"reflect"
	"testing"
)

const fixture = `
openapi: 3.1.0
info:
  title: Sample API
  description: Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.
  version: 0.1.9
paths:
  /users:
    get:
      summary: Returns a list of users.
      description: Optional extended description in CommonMark or HTML.
      responses:
        '200':    # status code
          description: A JSON array of user names
          content:
            application/json:
              schema:
                type: object
components:
  schemas:
    Address:
      description: Address
      title: Address
      type: object
      $schema: 'http://json-schema.org/draft-07/schema#'
      properties:
        city:
          description: City
          type: string
        state:
          description: State
          type:
            - string
            - integer
          examples:
            - Some State
        street_address_line_1:
          description: Street Address Line 1
          type: string
          examples:
            - Street A
`

func TestStuff(t *testing.T) {
	var yamlFile interface{}
	err := yaml.Unmarshal([]byte(fixture), &yamlFile)
	assert.NoError(t, err)

	result := recurse(reflect.ValueOf(yamlFile))
	fmt.Println(result)

	outYamlBytes, err := yaml.Marshal(result)
	assert.NoError(t, err)

	//Get a spec from the test definition in this file:
	swagger, err := openapi3.NewLoader().LoadFromData(outYamlBytes)
	assert.NoError(t, err)

	// Run our code generation:
	code, err := Generate(swagger, Configuration{
		PackageName: "dontcare",
		Generate: GenerateOptions{
			ChiServer:     false,
			EchoServer:    false,
			GinServer:     false,
			GorillaServer: false,
			Client:        true,
			Models:        false,
			EmbeddedSpec:  false,
		},
		Compatibility:     CompatibilityOptions{},
		OutputOptions:     OutputOptions{},
		ImportMapping:     nil,
		AdditionalImports: nil,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, code)
}

func recurse(value reflect.Value) interface{} {
	switch value.Kind() {
	case reflect.Map:
		newMap := make(map[interface{}]interface{})
		for _, key := range value.MapKeys() {
			if key.Elem().String() == "type" {
				array := value.MapIndex(key).Elem()
				if array.Kind() == reflect.Slice {
					var oneOf []map[string]string

					for i := 0; i < array.Len(); i++ {
						m := make(map[string]string)
						m["type"] = array.Index(i).Elem().String()
						oneOf = append(oneOf, m)
					}
					newMap["oneOf"] = oneOf
				}
			} else {
				newMap[key.Elem().String()] = recurse(value.MapIndex(key))
			}
		}
		return newMap

	case reflect.String:
		return value.String()
	case reflect.Interface:
		return recurse(value.Elem())
	}
	return value.Kind()
}
