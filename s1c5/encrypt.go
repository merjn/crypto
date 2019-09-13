package encrypt

import "encoding/hex"

type encryption struct {
	key Key
}

func (e *encryption) Encrypt(input []byte) string {
	var result []byte
	for _, i := range input {
		result = append(result, i^e.key.Current()[0])
		e.key.MovePos()
	}

	return hex.EncodeToString(result)
}

func New(key Key) *encryption {
	return &encryption{key: key}
}
