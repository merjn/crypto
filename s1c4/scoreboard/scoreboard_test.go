package scoreboard_test

import "testing"
import "encoding/hex"
import "cryptopals/scoreboard"

func TestFindBestScore(t *testing.T) {
	expected := "X"
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decoded, err := hex.DecodeString(input)
	if err != nil {
		t.Fatal(err) // sanity check
	}

	score, err := scoreboard.FindBestScore(decoded)
	if err != nil {
		t.Fatal(err)
	}

	if string(score.Character) != expected {
		t.Error("invalid key")
	}
}
