package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	zr, _ := zip.OpenReader("evil.zip")
	defer zr.Close()

	fmt.Printf("%+#v", zr.File[0])
	for _, f := range zr.File {
		r, _ := f.Open()
		data, _ := ioutil.ReadAll(r)
		os.WriteFile(f.Name, data, 0644)
	}
}

// github.com/patrickhener/go-evilarc
// go-evilarc -out evil.gz -depth 1 -platform unix somefile.txt
