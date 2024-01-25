package asciiArt

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Font = map[rune][]string

func GetArt(str, font string) string {
	fontMap, err := getFontMap(font)
	if err != nil {
		fmt.Println(err)
	}

	newArt := createArt(str, fontMap)
	return newArt
}

func getFontMap(fileName string) (Font, error) {
	file, err := os.Open("./" + fileName + ".txt")
	if err != nil {
		return nil, errors.New("Error opening file: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	font := Font{}

	for i, j := 0, 31; scanner.Scan(); i++ {
		if i%9 == 0 {
			j++
		}
		key := rune(j)
		font[key] = append(font[key], scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("Error reading file: " + err.Error())
	}

	return font, nil
}

func createArt(str string, font Font) string {
	res := ""
	lines := strings.Split(str, "\n")

	for _, line := range lines {
		if line == "" {
			res += "\n"
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, c := range line {
				res += font[c][i]
			}
			res += "\n"
		}
	}
	return res
}
