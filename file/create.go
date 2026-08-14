package file

import "os"

// Create an empty file at the given path.
func CreateEmpty(path string) error {
	return os.WriteFile(path, []byte{}, 0666)
}
