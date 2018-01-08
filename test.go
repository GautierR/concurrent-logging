package main

import (
	"concurrent-logging/packageA"
	"os"
	"io"
	"log"
	"concurrent-logging/packageB"
	"strconv"
	"fmt"
	"time"
)

func parrallel_log(ID int) {

	fileName := "log_" + strconv.Itoa(ID) + ".txt"
	if _, err := os.Stat(fileName); err == nil {
		os.Remove(fileName)
	}
	logFile, err := os.OpenFile(fileName, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Millisecond * 50 * time.Duration(ID))

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	for i := 0; i < 10000; i++{
		packageA.LogA(ID)
		packageB.LogB(ID)
	}

}


func main() {
	
	// Create goroutines
	fmt.Printf("Hello from main! \n")

	for i := 0; i < 3; i++  {
		go parrallel_log(i)
	}

	time.Sleep(time.Millisecond * 1000) // if i take this line out all i will see is hello and bye gopher!
	fmt.Println("Bye gopher!")

}


