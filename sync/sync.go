package main

import (
	"Task_for_mail/lib"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
)

func Definer(str string) (int, error) {
	if matchedURL, _ := regexp.MatchString(`^http.`, str); matchedURL {
		//handle URL
		fmt.Print(" URL ")
		val, err := lib.NumberInURL(str)
		if err != nil {
			return 0, err
		}
		return val, nil
	}
	if matchedFile, _ := regexp.MatchString(`^(.+)/([^/]+)$`, str); matchedFile {
		//handle File
		fmt.Print(" File ")
		val, err := lib.NumberInFile(str)
		if err != nil {
			return 0, err
		}
		return val, nil
	}
	return 0, nil
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if scanner.Err() != nil {
		log.Print(scanner.Err())
	}
	//input=append(input,"./test.txt")
	sum := 0
	for ind, val := range input {
		fmt.Print(ind, " : ", val)
		val, err := Definer(val)
		if err != nil && err != errors.New("Bad input param") {
			log.Println(err)
			return
		}
		fmt.Println(val)
		sum += val
	}

	fmt.Println(sum)
}
