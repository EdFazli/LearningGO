package main

import (
	"archive/zip"
	"bytes"
	"fmt"
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
}

//>>>>>>>>>>>>>>>>>>>>>>>>MAIN FUNCTION<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<//
