package decoder

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(input string) (string, error) {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(decoded), nil
}
