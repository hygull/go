package main

import (
	"log"
	"os/exec"
)

func main() {
	// cmd := exec.Command("sleep", "5")
	// err := cmd.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Waiting for command to finish...")
	// err = cmd.Wait()
	// log.Printf("Command finished with error: %v", err)

	log.Println("echo command is going to execute")
	cmd1 := exec.Command("echo", "Rishikesh")
	err1 := cmd1.Start()
	log.Println(err1)
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Printf("Waiting for command to finish...")
	err1 = cmd1.Wait()
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Println(err1)
	outputByte, err1 := cmd1.Output() //I implemented this by seeing the signature at https://golang.org/src/os/exec/exec.go?s=12280:12318#L451
	log.Println(err1)
	log.Printf("COMMAND: echo Rishikesh, OUTPUT: %v", string(outputByte))
	log.Printf("Command finished with error : %v", err1)
}
