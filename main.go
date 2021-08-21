package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sync"
	"time"
)

func main() {

	// declaration, initialization

	fmt.Println("varaibles .....")
	var number int
	number = 12
	fmt.Println("number", number)

	active := true // declare and initialize w/o explicit type declaration
	fmt.Println("active", active)

	const MAX int = 0

	// arrays
	fmt.Println("\nslices (arrays) .....")
	array := []int{34, 56, 34, 25, 17, 72, 75, 74, 83, 10, 90}
	fmt.Println("subset .....", array[4:8])

	// multiple return types
	fmt.Println("\nMULTIPLE RETURN FUNCTIONS .....")
	min, max, _, e := getSats(array) // blank identifier _ for the vars we do not care about the value for
	fmt.Println(min, max)

	//  error checking , no try catch
	fmt.Println("\nERRORS .....")

	array = []int{}
	min, max, _, e = getSats(array)
	if e != nil {
		fmt.Println(e)
	}

	type person struct {
		firstname string
		lastname  string
		age       int
	}
	fmt.Println("\nSTRUCTS .... ")

	sam := &person{"John", "Smith", 25}
	fmt.Println("person struct", sam.firstname, sam.lastname)

	fmt.Println("\nGO ROUTINES .... ")
	fmt.Println("loopGoRoutunes word count", loopGoRoutunes())
	fmt.Println("loop word count", loop())

}

// No public /private
func getSats(slice []int) (int, int, int, error) {

	if len(slice) == 0 || slice == nil { // return errors instead of throwing
		return 0, 0, 0, errors.New("error:empty slice")
	}
	max := 0
	min := 0
	sum := 0

	for i := 0; i < len(slice); i++ { // curly brace on same line

		if slice[i] > max || i == 0 {
			max = slice[i]
		}
		if slice[i] < min || i == 0 {
			min = slice[i]
		}
		sum = sum + slice[i]
	}
	return min, max, sum, nil

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

var loops = 20

func loop() int {
	total := 0
	defer timeTrack(time.Now(), "loop")
	for i := 0; i < loops; i++ {
		total = total + wordCount(readFile("file2.txt"))
	}
	return total
}

func loopGoRoutunes() int {
	var wg sync.WaitGroup
	wg.Add(loops)
	total := 0
	defer timeTrack(time.Now(), "loopWithRoutunes")
	for i := 0; i < loops; i++ {
		go func(i int) {
			defer wg.Done()
			total = total + wordCount(readFile("file1.txt"))
		}(i)
	}
	wg.Wait()
	return total
}

func readFile(name string) string {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("Err")
	}
	return (string(content))
}

func wordCount(value string) int {
	re := regexp.MustCompile(`[\S]+`)
	results := re.FindAllString(value, -1)
	return len(results)
}
