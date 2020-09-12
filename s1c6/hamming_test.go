package decoder_test

import "testing"
import "decoder"


func TestHammingDistance(t *testing.T) {
	input := [2]string{"this is a test", "wokka wokka!!!"}
	expected := 37

	distance, err := decoder.HammingDistance(input[0], input[1])
	if err != nil {
		t.Fatal(err)
	}

	if int(distance) != expected {
		t.Errorf("expected distance to be %d, got %d", expected, distance)
	}
}