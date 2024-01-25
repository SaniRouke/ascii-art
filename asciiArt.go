package asciiArt

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Style = map[rune][]string

func GetArt(str, style string) string {
	styleMap, err := getStyleMap(style)
	if err != nil {
		fmt.Println(err)
	}

	newArt := createArt(str, styleMap)
	return newArt
}

func getStyleMap(fileName string) (Style, error) {
	file, err := os.Open("./" + fileName + ".txt")
	if err != nil {
		return nil, errors.New("Error opening file:" + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	style := Style{}

	for i, j := 0, 31; scanner.Scan(); i++ {
		if i%9 == 0 {
			j++
		}
		key := rune(j)
		style[key] = append(style[key], scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("Error reading file:" + err.Error())
	}

	return style, nil
}

func createArt(str string, style Style) string {
	fmt.Println(str)
	res := ""
	lines := strings.Split(str, "\n")

	for _, line := range lines {
		if line == "" {
			res += "\n"
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, c := range line {
				res += style[c][i]
			}
			res += "\n"
		}
	}
	return res
}
