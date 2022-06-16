package codegen

import (
	"encoding/json"
	"fmt"
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
servers:
  - url: http://api.example.com/v1
    description: Optional server description, e.g. Main (production) server
  - url: http://staging-api.example.com
    description: Optional server description, e.g. Internal staging server for testing
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

func TestName(t *testing.T) {
	bytes := []byte("123")
	var numberInterface interface{}
	err := json.Unmarshal(bytes, &numberInterface)
	assert.NoError(t, err)

	result := reflect.ValueOf(numberInterface)
	fmt.Println(result)
	fmt.Println(result.Kind())
}

func TestStuff(t *testing.T) {
	var yamlFile interface{}
	err := yaml.Unmarshal([]byte(fixture), &yamlFile)
	assert.NoError(t, err)

	recurse(reflect.ValueOf(yamlFile))

	// Get a spec from the test definition in this file:
	//swagger, err := openapi3.NewLoader().LoadFromData([]byte(fixture))
	//assert.NoError(t, err)
	//
	//// Run our code generation:
	//code, err := Generate(swagger, Configuration{
	//	PackageName: "dontcare",
	//	Generate: GenerateOptions{
	//		ChiServer:     false,
	//		EchoServer:    false,
	//		GinServer:     false,
	//		GorillaServer: false,
	//		Client:        true,
	//		Models:        false,
	//		EmbeddedSpec:  false,
	//	},
	//	Compatibility:     CompatibilityOptions{},
	//	OutputOptions:     OutputOptions{},
	//	ImportMapping:     nil,
	//	AdditionalImports: nil,
	//})
	//assert.NoError(t, err)
	//assert.NotEmpty(t, code)
}

func recurse(value reflect.Value) {
	fmt.Println("value", value, "kind", value.Kind())

	switch value.Kind() {
	case reflect.Map:
		fmt.Println("do something with a map")
		for _, key := range value.MapKeys() {
			fmt.Println("key", key)
			recurse(value.MapIndex(key))
		}

	case reflect.String:
		fmt.Println("do something with a string")
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fmt.Println("do something with an array")
	case reflect.Interface:
		fmt.Println("do something with an interface?")
		theMap := (map[string]interface{})(value)

		wtf := value.(map[string]interface{})


		if theMap, ok := ; ok {

		}


	}
}
