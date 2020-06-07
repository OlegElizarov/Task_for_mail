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

//import _ "net/http/pprof"
//import _ "net/http"

const MAX_GOR = 5 //total count of goroutines is MAX_GOR+1(func main is goroutine too)

func Definer(str string, ch chan int, job chan struct{}, group *sync.WaitGroup) {
	defer group.Done()

	if matchedURL, _ := regexp.MatchString(`^http.`, str); matchedURL {
		//handle URL
		val, err := lib.NumberInURL(str)
		if err != nil {
			log.Print(err)
			<-job
			return
		}
		fmt.Printf("Count for %s : %v\n", str, val)
		ch <- val
		<-job
		return
	}
	if matchedFile, _ := regexp.MatchString(`^(.+)/([^/]+)$`, str); matchedFile {
		//handle File
		val, err := lib.NumberInFile(str)
		if err != nil {
			log.Print(err)
			<-job
			return
		}
		fmt.Printf("Count for %s : %v\n", str, val)
		ch <- val
		<-job
		return
	}
	//handle not file and not URL
	log.Print("Wrong input param")
	<-job
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
	ch := make(chan int, len(input))    //chan of results
	job := make(chan struct{}, MAX_GOR) //chan for goroutine limit

	sum := 0
	var wg sync.WaitGroup
	wg.Add(len(input))

	for _, val := range input {
		select {
		case job <- struct{}{}:
			go Definer(val, ch, job, &wg)
			//fmt.Println(" Gor numb ", runtime.NumGoroutine())
		}
	}
	wg.Wait()
	close(ch)

	for v := range ch {
		sum += v
	}
	fmt.Printf("Total: %v\n", sum)
}
