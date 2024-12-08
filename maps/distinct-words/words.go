package words

import "strings"

func countDistinctWords(messages []string) int {
	res := make(map[string]int)

	for _, message := range messages {
		// split & convert to lowercase
		splitM := strings.Fields(strings.ToLower(message))
		for _, msg := range splitM {
			res[msg]++
		}
	}
	return len(res) 
}

// map doesn't consider duplicate keys
/* ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
fmt.Println(len(ages)) // 2 */
