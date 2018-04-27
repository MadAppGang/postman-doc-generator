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

// PrefixDirectory places the directory name at the beginning of each name in the list.
func PrefixDirectory(directory string, names []string) []string {
	if directory == "." {
		return names
	}
	ret := make([]string, len(names))
	for i, name := range names {
		ret[i] = filepath.Join(directory, name)
	}
	return ret
}
