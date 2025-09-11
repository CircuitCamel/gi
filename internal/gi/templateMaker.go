package gi

import (
	"os"
)

func TemplateMaker(arr []string) []byte {
	var resultArr []byte
	for _, v := range arr {
		fileCont, err := os.ReadFile("templates/" + v + ".gitignore")
		if err != nil {
			return []byte(v + " is not an option, use command \"list\" to show templates")
		}
		resultArr = append(resultArr, fileCont...)
	}
	return resultArr
}
