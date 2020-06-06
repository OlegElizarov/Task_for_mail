package main

import (
	"Task_for_mail/lib"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	_ "runtime"
	"sync"
)

func Definer(str string, ch chan int, group *sync.WaitGroup) {
	defer group.Done()
	if matchedURL, _ := regexp.MatchString(`^http.`, str); matchedURL {
		//handle URL
		val, err := lib.NumberInURL(str)
		if err != nil {
			fmt.Print(err)
			close(ch)
			return
		}
		fmt.Println(str, " URL ", val)
		ch <- val
		return
	}
	if matchedFile, _ := regexp.MatchString(`^(.+)/([^/]+)$`, str); matchedFile {
		//handle File
		val, err := lib.NumberInFile(str)
		if err != nil {
			fmt.Print(err)
			close(ch)
			return
		}
		fmt.Println(str, " File ", val)
		ch <- val
		return
	}
	return
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
	//input = append(input, "./test.txt", "./try.txt")
	ch := make(chan int, len(input))
	sum := 0
	var wg sync.WaitGroup
	wg.Add(len(input))

	for _, val := range input {
		go Definer(val, ch, &wg)
	}
	//fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	close(ch)
	for v := range ch {
		sum += v
	}
	fmt.Println(sum)
}
