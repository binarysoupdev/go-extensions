package file

import (
	"io"
	"os"

	"github.com/binarysoupdev/go-extensions/errors"
)

func Copy(dest string, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return errors.Chain(err, "error opening source file")
	}
	defer s.Close()

	d, err := os.Create(dest)
	if err != nil {
		return errors.Chain(err, "error creating destination file")
	}
	defer d.Close()

	_, err = io.Copy(d, s)
	if err != nil {
		return errors.Chain(err, "error copying data")
	}

	return nil
}
