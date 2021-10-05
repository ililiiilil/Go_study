package longestName

import "fmt"

func LongestName() {
	names := make([]string, 0, 3)

	var name string

	for {
		fmt.Scanf("%s", &name)
		if name == "1" {
			break
		} else {
			names = append(names, name)
		}
	}

	var result string = names[0]

	for i := 1; i < len(names); i++ {
		if len(names[i]) > len(result) {
			result = names[i]
		}
	}

	fmt.Println(result, len(result))
}
