package file

import "os"

func CreateEmpty(path string) error {
	return os.WriteFile(path, []byte{}, 0666)
}
