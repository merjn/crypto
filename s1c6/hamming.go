package decoder

import "errors"
import "fmt"

func HammingDistance(str1, str2 string) (int, error) {
	if len(str1) != len(str2) {
		return 0, errors.New("string lengths do no match")
	}

	// 

	var modifications int
	for index, _ := range str1 {
		fmt.Println(str1[index]^str2[index])
	} 

	return modifications, nil
} 