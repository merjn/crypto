package singlecharacter

type score struct {
	Letter byte
	Total  float64
}

var LetterFrequency = map[rune]float64{
	'a': 0.8167,
	'b': 0.1492,
	'c': 0.2782,
	'd': 0.4253,
	'e': 0.12702,
	'f': 0.228,
	'g': 0.2015,
	'h': 0.6094,
	'i': 0.6966,
	'j': 0.153,
	'k': 0.772,
	'l': 0.4025,
	'm': 0.2406,
	'n': 0.6749,
	'o': 0.7507,
	'p': 0.1929,
	'q': 0.095,
	'r': 0.5987,
	's': 0.6327,
	't': 0.9056,
	'u': 0.2758,
	'v': 0.978,
	'w': 0.2360,
	'x': 0.150,
	'y': 0.1974,
	'z': 0.074,
}

func FindBestScore(input []byte) score {
	var bestScore score
	for l := 0; l <= 255; l++ {
		letter := byte(l)

		var result []byte
		for _, i := range input {
			xored := i ^ letter
			result = append(result, xored)
		}

		s := Score(string(result))

		if bestScore.Total < s {
			bestScore = score{Letter: letter, Total: s}
		}

	}

	return bestScore
}

func Score(input string) float64 {
	var score float64
	for _, str := range input {
		result, ok := LetterFrequency[str]
		if ok {
			score += result
		}
	}
	return score
}
