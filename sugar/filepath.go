package sugar

import (
	"log"
	"os"
	"path/filepath"
)

// IsDirectory returns true if the named file is a directory
func IsDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

// AddPathPrefix places the path at the beginning of each name in the list.
func AddPathPrefix(path string, names []string) []string {
	if path == "." {
		return names
	}
	ret := make([]string, len(names))
	for i, name := range names {
		ret[i] = filepath.Join(path, name)
	}
	return ret
}
