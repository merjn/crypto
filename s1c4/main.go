package main

import "cryptopals/scoreboard"
import "encoding/hex"
import "io/ioutil"
import "log"
import "cryptopals/decoder"
import "fmt"
import "strings"

type Winner struct {
	Sentence string
	Character byte
	Score float64
}

// main iterates over file.txt and calls the decoder package to
// check which line has been encoded by single byte xor.
func main() {
	in, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}


	lines := strings.Split(string(in), "\r\n")
	var winner Winner

	for _, line := range lines {
		decoded, err := hex.DecodeString(strings.TrimSpace(line))
		if err != nil {
			log.Fatalln(err)
		}

		score, err := scoreboard.FindBestScore(decoded)
		if err != nil {
			log.Fatalln(err)
		}


		if score.Score > winner.Score {
			winner = Winner { Sentence: string(decoded), Character: score.Character, Score: score.Score }
		}
	}

	fmt.Println(winner.Character)
	fmt.Println(winner.Sentence)
	fmt.Println(decoder.Decode(winner.Sentence, winner.Character))
}