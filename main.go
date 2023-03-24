package main

import (
	"fmt" 
	"strings"
	"log"
	"os"
    "time"
	"strconv"
)

func WordCount(s string) map[string]int { 
    words := strings.Fields(s) 
    m := make(map[string]int) 
    for _, word := range words { 
        m[word] += 1 
    } 
    return m 
}

func ConcurrentWordCount(words []string) map[string]int{
	myMap := make(map[string]int) 
	ch := make(chan map[string]int)
	for _, word := range words{
		go func(s string){
			ch <- WordCount(s)
		}(word)
	}
	
	for range words{
		m := <-ch
		for lett,freq := range m {
			myMap[lett] += freq
		}
	}
	return myMap
}



func main() {




	dirName := "./DividedInput/"

	files ,err := os.ReadDir(dirName)

	f , _ := os.Create("output.txt")

	if err != nil {
        log.Fatal(err)
    }

	start := time.Now()

	for _, file := range files {

        fmt.Println("\nReading",file.Name(),"\n")

		content, err := os.ReadFile(dirName+file.Name())

		if err != nil {
			log.Fatal(err)
		}

		txt := string(content)

		// res := WordCount(txt)
		res := ConcurrentWordCount([]string{txt})

		for word,freq := range res {
			f.WriteString(word + "\t=\t" + strconv.Itoa(freq) + "\n" )
			// fmt.Printf("[ %v ]\t=\t%v",word,freq)
		}

    }

	end := time.Since(start)
    log.Printf("\n\n word count took: %s", end)

	

}