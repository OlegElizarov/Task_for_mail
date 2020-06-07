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
		val, err := lib.NumberInURL(str)
		if err != nil {
			log.Print(err)
			return 0, err
		}
		fmt.Printf("Count for %s : %v\n", str, val)
		return val, nil
	}
	if matchedFile, _ := regexp.MatchString(`^(.+)/([^/]+)$`, str); matchedFile {
		//handle File
		val, err := lib.NumberInFile(str)
		if err != nil {
			log.Print(err)
			return 0, err
		}
		fmt.Printf("Count for %s : %v\n", str, val)
		return val, nil
	}
	log.Print("Wrong input param")
	return 0, errors.New("Wrong input param")
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
		if err != nil {
			log.Println(err)
			val = 0
		}
		sum += val
	}
	fmt.Printf("Total: %v\n", sum)
}
