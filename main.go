package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	ScanFile string `json:"ScanFile"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println("\n\x1b[31mError: Please check config.json", "\x1b[0m")
		//	fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func main() {
	fmt.Println("\n\x1b[93mBracket Missing Finder {!!!}", "\x1b[0m")
	fmt.Println("\n\x1b[35mCreated with ♥ by Zile42O", "\x1b[0m")
	start := time.Now()

	config := LoadConfiguration("config.json")

	open := 0
	closed := 0
	var openbraces [100]int
	var lastln [100]int
	meh := -1
	//reset
	for i := 0; i < 100; i++ {
		openbraces[i] = -1
		lastln[i] = -1
	}
	//scanning file
	file, err := os.Open(config.ScanFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//loop lines
	lines := 1
	ezpz := 0
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		if err := strings.HasPrefix(scanner.Text(), "//"); err == false {
			if err := strings.HasPrefix(scanner.Text(), "/*"); err == false {
				if err := strings.Contains(scanner.Text(), "{"); err != false {
					for i := 0; i < 100; i++ {
						if openbraces[i] == -1 {
							meh = i
							break
							fmt.Println(i)
						}
					}
					//fmt.Println(meh)
					open++
					openbraces[meh] = lines
					//fmt.Println(lines)
					//fmt.Println(openbraces[meh])
					lastln[ezpz] = meh
					ezpz++

				}
				if err := strings.Contains(scanner.Text(), "}"); err != false {
					closed++
					ezpz--
					openbraces[lastln[ezpz]] = -1
				}
			}
		}
		lines++
	}
	//count
	count := 0
	for i := 0; i < 100; i++ {
		if openbraces[i] != -1 {
			count++
			fmt.Println("\n\x1b[31mThere is an unclosed brace at line number: ", openbraces[i], "\x1b[0m")
		}
	}
	//end loop
	//	fmt.Println("\n\x1b[35mOpen:  ", open, "\x1b[0m")
	//	fmt.Println("\n\x1b[35mClosed:  ", closed, "\x1b[0m")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("\n\x1b[32m× Program took:\x1b[0m \x1b[33m", elapsed, "\x1b[0m")
	time.Sleep(5000 * time.Millisecond)
}
