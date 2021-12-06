package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
)

//>>>>>>>>>>>>>>>>>>>>>>>>MAIN FUNCTION<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<//

func main() {
	fmt.Println("This is errors examples")

	//Sentinel errors - to indicate that you cannot start or continue processing
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}

	// wrapping and unwrapping errors
	errs := fileChecker("not_here.txt")
	if errs != nil {
		fmt.Println(errs)
		if wrappedErrs := errors.Unwrap(errs); wrappedErrs != nil {
			fmt.Println(wrappedErrs)
		}
	}
}

//>>>>>>>>>>>>>>>>>>>>>>>>MAIN FUNCTION<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<//

// wrapping and unwrapping errors
func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in filechecker: %w", err)
	}
	f.Close()
	return nil
}
