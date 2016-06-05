package symbol

import (
	"io"
	"log"
	"os"
	"path"
)

func IsFileExist(s string) bool {
	if _, err := os.Stat(s); os.IsNotExist(err) {
		return false
	}

	return true
}

func CopyFile(source string, dest string, force bool) (err error) {
	if IsFileExist(dest) && force == false {
		return nil
	}

	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}
		log.Println("copy file => ", dest)
	}

	return
}

func CopyDir(source string, dest string, force bool) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := path.Join(source, obj.Name())

		destinationfilepointer := path.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer, force)
			if err != nil {
				return err
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer, force)
			if err != nil {
				return err
			}
		}

	}

	return
}
