package main

func BossBabyRevenge(report string) string {
	// The first action is retaliation “R” and the last shot has no revenge
	if string(rune(report[0])) == "R" || string(rune(report[len(report)-1])) == "S" {
		return "Bad boy"
	}

	shots := 0
	// revenge := 0
	for _, event := range report {
		switch true {
		case string(event) == "S":
			shots++
		// Revenge for the previous shooting
		case string(event) == "R" && shots > 0:
			shots--
			// case string(event) == "R":
			// 	revenge++
		}
	}

	// Revenge for every shot
	if shots == 0 {
		return "Good boy"
	}

	// Revenge isn't complete in every shot
	return "Bad boy"
}
