package decoder_test

import "testing"
import "cryptopals/decoder"
import "strings"

var (
	expected = "Cooking MC's like a pound of bacon"
	key      = byte('X')
)

func TestHexDecode(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	output, err := decoder.DecodeHex(input, key)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.EqualFold(expected, output) {
		t.Errorf("output %s does not match expected %s", output, expected)
	}
}

func TestDecode(t *testing.T) {
	input := []byte{27, 55, 55, 51, 49, 54, 63, 120, 21, 27, 127, 43, 120, 52, 49, 51, 61, 120, 57, 120, 40, 55, 45, 54, 60, 120, 55, 62, 120, 58, 57, 59, 55, 54}
	output := decoder.Decode(string(input), key)

	if !strings.EqualFold(expected, output) {
		t.Errorf("output %s does not match expected %s", output, expected)
	}
}
