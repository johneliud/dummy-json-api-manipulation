package utils

import (
	"fmt"
	"os"
	"strconv"
)

// GetPort checks for the number of arguments used to run the program's server and validates the choice of port used.
func GetPort() string {
	var port string

	switch len(os.Args) {
	case 1:
		port = ":8080"
	case 2:
		portNum, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error converting %v to int: %v\n", os.Args[1], err)
			os.Exit(1)
		}

		if portNum < 1024 || portNum > 65535 {
			fmt.Println("Use a range between 1024 and 65535")
			os.Exit(1)
		}
		port = ":" + strconv.Itoa(portNum)
	default:
		fmt.Println(`Usage: "go run main.go" OR "go run main.go [PORT]`)
		os.Exit(1)
	}
	return port
}
