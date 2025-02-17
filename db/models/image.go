package db

import (
	"encoding/json"

	jsonschema "github.com/alecthomas/jsonschema"
	log "github.com/sirupsen/logrus"
	gjs "github.com/xeipuuv/gojsonschema"
)

// Image is a resource type that provides metadata for endpoint images in use within Lessons
type Image struct {
	Slug string `json:"Slug" yaml:"slug" jsonschema:"Unique identifier for this image,pattern=^[A-Za-z0-9\\-]*$"`

	Description string `json:"Description" yaml:"description" jsonschema:"Description of this image"`

	// Temporary measure to grant privileges to endpoints selectively
	Privileged bool `json:"Privileged" yaml:"privileged" jsonschema:"Should this image be granted admin privileges?"`

	// Used to allow authors to know which interfaces are available, and in which order they'll be connected
	NetworkInterfaces []string `json:"NetworkInterfaces" yaml:"networkInterfaces" jsonschema:"minItems=1"`

	SSHUser     string `json:"SSHUser" yaml:"sshUser" jsonschema:"minLength=1,description=Username for SSH connections"`
	SSHPassword string `json:"SSHPassword" yaml:"sshPassword" jsonschema:"minLength=1,Password for SSH Connections"`

	ConfigUser     string `json:"ConfigUser" yaml:"configUser" jsonschema:"minLength=1,description=Username for configuration scripts"`
	ConfigPassword string `json:"ConfigPassword" yaml:"configPassword" jsonschema:"minLength=1,Password for configuration scripts"`
}

// GetSchema returns a Schema to be used in creation wizards
func (i Image) GetSchema() *jsonschema.Schema {
	return jsonschema.Reflect(i)
}

// JSValidate uses an Antidote resource's struct properties and tags to construct a jsonschema
// document, and then validates that instance's values against that schema.
func (i Image) JSValidate() bool {

	// Load JSON Schema document for type
	schemaReflect := jsonschema.Reflect(i)
	b, _ := json.Marshal(schemaReflect)
	schemaLoader := gjs.NewStringLoader(string(b))
	schema, _ := gjs.NewSchema(schemaLoader)

	// Load instance JSON document
	b, err := json.Marshal(i)
	if err != nil {
		log.Error(err)
		return false
	}
	documentLoader := gjs.NewStringLoader(string(b))

	// Perform validation
	result, err := schema.Validate(documentLoader)
	if err != nil {
		log.Error(err)
		return false
	}

	validationErrors := result.Errors()
	for j := range validationErrors {
		log.Errorf("Validation error - %s", validationErrors[j].String())
	}

	return result.Valid()
}
