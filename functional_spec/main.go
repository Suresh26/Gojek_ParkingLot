package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"parking_lot/functional_spec/parking"
	slot "parking_lot/functional_spec/slot"
	"strings"
)

func main() {
	slots := &slot.Slots{}
	if len(os.Args) > 1 && os.Args[1] != "" {
		// fileName := flag.String("fpath", "fixtures/file_input.txt", "Read the file") //Debug mode
		fileName := flag.String("fpath", os.Args[1], "Read the file")
		flag.Parse()
		cmdFile, err := os.Open(*fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer cmdFile.Close()
		scanner := bufio.NewScanner(cmdFile)
		for scanner.Scan() {
			cmdInput := scanner.Text()
			cmdInput = strings.TrimRight(cmdInput, "\n")
			slots = getString(slots, cmdInput)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		cmdInput, _ := reader.ReadString('\n')
		cmdInput = strings.TrimSuffix(cmdInput, "\n")
		if cmdInput != "" {
			slots = getString(slots, cmdInput)
		}
	}
}

func getString(slots *slot.Slots, data string) *slot.Slots {

	split := strings.SplitN(data, " ", 2)
	str := ""
	if len(split) > 1 {
		str = split[1]
		str = strings.TrimSuffix(split[1], "\r")
	}
	message := ""
	slots, message = parking.ReadFileCmd(split[0], str, slots)
	fmt.Println(message)
	return slots
}
