package postman

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

const modelsItemName = "Models"

var errInvalidNode = errors.New("node is invalid")
var errNotFound = errors.New("not found")

// Schema is the postman schema.
type Schema struct {
	node       map[string]interface{}
	modelsItem *map[string]interface{}
}

// ParseFile parses a file by filename, unmarshal data and returns schema.
func ParseFile(filename string) Schema {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	schema := Schema{
		node: make(map[string]interface{}),
	}

	err = json.Unmarshal(data, &schema.node)
	if err != nil {
		panic(err)
	}

	return schema
}

// GetModels returns a models from the schema.
func (s *Schema) GetModels() (string, error) {
	item, err := s.findItemByName(modelsItemName)
	if err != nil {
		return "", fmt.Errorf("Cannot find models node: %s", err)
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
	data, err := json.Marshal(s.node)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}
