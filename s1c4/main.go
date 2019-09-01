package main

import "fmt"
import "io/ioutil"
import "log"
import "bytes"
import "encoding/hex"

// main iterates over file.txt and calls the decoder package to
// check which line has been encoded by single byte xor.
func main() {
	out, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	separator := []byte("\n")
	lines := bytes.Split(out, separator)
	for _, line := range lines {
		decoded := hex.EncodeToString(line)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(decoded)
	}
}