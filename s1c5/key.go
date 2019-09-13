package encrypt

type Key struct {
	Value string
	pos   int
}

func (k *Key) Current() string {
	return string(k.Value[k.pos])
}

func (k *Key) MovePos() {
	if (k.pos + 1) == len(k.Value) {
		k.pos = 0
	} else {
		k.pos++
	}
}
