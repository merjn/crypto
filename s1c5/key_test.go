package encrypt_test

import "testing"
import "encrypt"

func TestKey(t *testing.T) {
	key := encrypt.Key{Value: "ICE"}
	counter := 0

	// Iterate and increment counter by 1
	for i := 0; i < 10000; i++ {
		// reset counter if the value is too high.
		if counter == 3 {
			counter = 0
		}

		var expected string
		if counter == 0 {
			expected = "I"
		} else if counter == 1 {
			expected = "C"
		} else if counter == 2 {
			expected = "E"
		} else {
			t.Fatal("unable to set expected")
		}

		currentKey := key.Current()
		if currentKey != expected {
			t.Errorf("expected key to be %s, got %s", expected, currentKey)
		}

		// move position; we're done
		key.MovePos()
		counter++
	}
}
