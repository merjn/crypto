package main

import "cryptopals/scoreboard"
import "encoding/hex"
import "io/ioutil"
import "log"
import "cryptopals/decoder"
import "fmt"
import "strings"

type Winner struct {
	Sentence  []byte
	Character byte
	Score     float64
}

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
			winner = Winner{Sentence: decoded, Character: score.Character, Score: score.Score}
		}
	}

	fmt.Println(winner.Character)
	fmt.Println(winner.Sentence)
	fmt.Println(string(decoder.Decode(string(winner.Sentence), winner.Character)))
}
