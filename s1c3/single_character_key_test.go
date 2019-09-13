package singlecharacter

import "testing"
import "encoding/hex"

func TestFindBestScore(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decoded, err := hex.DecodeString(input)
	if err != nil {
		t.Fatal(err) // sanity check
	} 

	score := FindBestScore(decoded)
	if string(score.Letter) != "X" {
		t.Error("invalid key")
	}
}

func TestScore(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyz"
	for _, character := range input {
		result := Score(string(character))

		value, ok := LetterFrequency[character]
		if !ok {
			t.Fatal("unable to get value from LetterFrequency") // sanity check
		}

		if value != result {
			t.Errorf("expected value to be %f, got %f", result, value)
		}
	}
}

func getFrequencyTotal() float64 {
	var total float64
	for _, f := range LetterFrequency {
		total += f
	}

	return total
}