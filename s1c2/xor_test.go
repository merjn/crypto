package xor

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"
)

func TestProduceXorCombination(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	expected := "746865206b696420646f6e277420706c6179"

	result, err := ProduceXorCombination(input)
	if err != nil {
		t.Error(err)
	}

	if !strings.EqualFold(result, expected) {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestGenerateXoredBuffer(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	expected := []byte{116, 104, 101, 32, 107, 105, 100, 32, 100, 111, 110, 39, 116, 32, 112, 108, 97, 121}

	decoded, err := hex.DecodeString(input)
	if err != nil {
		t.Fatal(err)
	}

	decodedXor, err := hex.DecodeString(XorAgainst)
	if err != nil {
		t.Fatal(err)
	}

	result := GenerateXoredBuffer(decoded, decodedXor)
	if !bytes.EqualFold(expected, result) {
		t.Error("unable to generate correct xored buffer")
	}
}
