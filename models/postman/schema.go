package postman

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/madappgang/postman-doc-generator/models"
)

const modelsItemName = "Models"

var errInvalidNode = errors.New("node is invalid")
var errNotFound = errors.New("not found")
var modelPattern = regexp.MustCompile(`#+\s(\w+)[^#]*`)

// Schema is the postman schema.
type Schema struct {
	node       map[string]interface{}
	modelsItem *map[string]interface{}
}

// ParseFile parses a file by filename, unmarshal data and returns schema.
func ParseFile(filename string) (Schema, error) {
	schema := Schema{
		node: make(map[string]interface{}),
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return schema, err
	}

	err = json.Unmarshal(data, &schema.node)

	return schema, err
}

// createModelsItem creates an item in the schema to store models.
// * should be used if this element is not present
func (s *Schema) createModelsItem() (*map[string]interface{}, error) {
	item, err := s.addItem(modelsItemName, "")
	if err != nil {
		return nil, err
	}

	s.modelsItem = item

	return item, nil
}

// getModels returns all models as a string from the schema.
func (s *Schema) getModels() (string, error) {
	item, err := s.findItemByName(modelsItemName)
	if err != nil {
		if err == errNotFound {
			item, err = s.createModelsItem()
			if err != nil {
				return "", fmt.Errorf("Cannot create models node: %s", err)
			}
		} else {
			return "", fmt.Errorf("Cannot find models node: %s", err)
		}
	}

	description, ok := (*item)["description"].(string)
	if !ok {
		return "", fmt.Errorf("Cannot get models: %s", errInvalidNode)
	}

	return description, nil
}

// findItemByName searches item by name, and returns pointer to it.
func (s *Schema) findItemByName(name string) (*map[string]interface{}, error) {
	items, ok := s.node["item"].([]interface{})
	if !ok {
		return nil, errInvalidNode
	}

	for _, tmp := range items {
		item, ok := tmp.(map[string]interface{})
		if !ok {
			return nil, errInvalidNode
		}

		itemName, ok := item["name"].(string)

		if itemName == name {
			return &item, nil
		}
	}

	return nil, errNotFound
}

// addItem creates and adds an item to items array.
func (s *Schema) addItem(name, description string) (*map[string]interface{}, error) {
	items, ok := s.node["item"].([]interface{})
	if !ok {
		return nil, errInvalidNode
	}

	item := make(map[string]interface{})
	item["name"] = name
	item["description"] = description
	item["item"] = make([]interface{}, 0, 0)

	s.node["item"] = append(items, item)

	return &item, nil
}

// SetModels replaces models with new ones.
// If models item is missing, creates it.
func (s *Schema) SetModels(models string) {
	if s.modelsItem != nil {
		(*s.modelsItem)["description"] = models
	} else {
		item, err := s.findItemByName(modelsItemName)
		if err != nil && err != errNotFound {
			panic(err)
		} else if err != nil && err == errNotFound {
			item, err = s.addItem(modelsItemName, "")
			if err != nil {
				panic(err)
			}
		}

		s.modelsItem = item

		(*s.modelsItem)["description"] = models
	}
}

// Save saves the schema to file and returns nil.
func (s *Schema) Save(filename string) error {
	data, err := json.MarshalIndent(s.node, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

// AddModels adds new models to schema
// Non nil verbose error returns if something goes wrong
func (s *Schema) AddModels(newModels []models.Model) error {
	models, err := s.getModelsMap()
	if err != nil {
		return err
	}

	for _, newModel := range newModels {
		models[newModel.Name] = newModel.String()
	}

	var str string
	for _, model := range models {
		str += model
	}
	s.SetModels(str)

	return nil
}

// getModelsMap returns all models as map from schema
// Non nil verbose error returns if something goes wrong
func (s *Schema) getModelsMap() (map[string]string, error) {
	m := make(map[string]string)

	models, err := s.getModels()
	if err != nil {
		return nil, err
	}

	for _, submatches := range modelPattern.FindAllSubmatch([]byte(models), -1) {
		model := string(submatches[0])
		modelName := string(submatches[1])
		m[modelName] = model
	}

	return m, nil
}
