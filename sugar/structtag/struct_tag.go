package structtag

import (
	"reflect"
	"strings"
)

// GetNameFromTag extracts name from tag by specified key and returns it.
// If key is missing in tag returns empty string.
func GetNameFromTag(tag, key string) string {
	st := reflect.StructTag(tag)
	name := st.Get(key)

	return strings.Split(name, ",")[0]
}
