package xor

import (
	"encoding/hex"
)

const XorAgainst = "686974207468652062756c6c277320657965"

func ProduceXorCombination(input string) (string, error) {
	decodedInput, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	decodedXor, err := hex.DecodeString(XorAgainst)
	if err != nil {
		return "", err
	}

	buffer := GenerateXoredBuffer(decodedInput, decodedXor)
	return hex.EncodeToString(buffer), err
}

func GenerateXoredBuffer(decoded, decodedXor []byte) []byte {
	var buffer []byte
	for index := range decoded {
		buffer = append(buffer, decoded[index]^decodedXor[index])
	}

	return buffer
}
