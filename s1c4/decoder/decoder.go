// Package decoder is responsible for decoding single-character xored data.
package decoder

import "encoding/hex"

// DecodeHex decodes the hex-encoded value and passes it to Decode.
func DecodeHex(input string, key byte) (string, error) {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	return string(Decode(string(decoded), key)), nil
}

// Decode iterates over input as a sequence of bytes and stores the decoded value to result.
func Decode(input string, key byte) []byte {
	var result []byte
	for _, in := range []byte(input) {
		result = append(result, in^key)
	}

	return result
}
