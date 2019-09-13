package encrypt_test

import "testing"
import "strings"
import "encrypt"

type Expected struct {
	Poem   []byte
	Result string
}

func TestEncrypt(t *testing.T) {
	expected := Expected{
		Poem: []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`),
		Result: "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f",
	}

	key := encrypt.Key{Value: "ICE"}

	encryption := encrypt.New(key)

	result := encryption.Encrypt(expected.Poem)
	if !strings.EqualFold(expected.Result, result) {
		t.Errorf("expected value to be %s, got %s", expected.Result, result)
	}
}
