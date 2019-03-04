package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// Lectura host location
const LECTURA string = "@lec.cs.arizona.edu:"
// CS352 lectura assignments paths
const LECTURA_ASSIGNMENTS_PATHS string = "/home/cs352/spring19/assignments/"

func main() {
	newProject := flag.NewFlagSet("new", flag.ExitOnError)
	newProjectName := newProject.String("name", "", "The name of the new project(must be of the form assg*)")
	newProjectDestination := newProject.String("dest", "", "Filepath to new project(defaults to current directory)")
	if len(os.Args) < 2 {
		newProject.Usage()
		os.Exit(1)
	}
	switch command := os.Args[1]; command {
	case "new":
		err := newProject.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if newProject.Parsed() {
		if *newProjectName == "" {
			newProject.PrintDefaults()
			os.Exit(1)
		}
		createNewProject(*newProjectName, *newProjectDestination)
	}
}

// GetUserNameAndPassword gets the users netid and password for lectura
func getUserNameAndPassword() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("NetID: ")
	netId, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return netId, password
}

// CreateNewProject attempts to create a new project on the users local machine.
// It attempts to copy the assignment from /home/cs352/spring19/assignments/ on lectura
func createNewProject(name string, dest string) {

}
