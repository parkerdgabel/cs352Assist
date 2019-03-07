package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Lectura host location
const Lectura string = "@lec.cs.arizona.edu:"

// LecturaAssignmentPath is the path to the assignments directory for cs352 Spring 2019
const LecturaAssignmentPath string = "/home/cs352/spring19/assignments/"

func main() {
	newProject := flag.NewFlagSet("new", flag.ExitOnError)
	newProjectName := newProject.String("name", "", "The name of the new project(must be of the form assg*)")
	newProjectDestination := newProject.String("dest", "./", "Filepath to new project(defaults to current directory)")
	copyProject := flag.NewFlagSet("copy", flag.ExitOnError)
	copyProjectDestination := copyProject.String("dest", "", "Destination on lectura to copy the project.")
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
	case "copy":
		err := copyProject.Parse(os.Args[2:])
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
		CreateNewProject(*newProjectName, *newProjectDestination)
	} else if copyProject.Parsed() {
		CopyProject(*copyProjectDestination)
	}
}

// GetUserName gets the users netid for lectura
func GetUserName() string {
	var netID string
	if netID = os.Getenv("LECTURA_USERNAME"); netID != "" {
		return netID
	}
	fmt.Print("NetID: ")
	_, err := fmt.Scanf("%s", &netID)
	if err != nil {
		log.Fatal(err)
	}
	return netID
}

// CreateNewProject attempts to create a new project on the users local machine.
// It attempts to copy the assignment from /home/cs352/spring19/assignments/ on lectura
func CreateNewProject(name string, dest string) {
	address := fmt.Sprintf("%s%s%s%s", GetUserName(), Lectura, LecturaAssignmentPath, name)
	address = strings.TrimSuffix(address, "\n")
	scp := exec.Command("scp", "-r", address, dest)
	scp.Stdin = os.Stdin
	scp.Stderr = os.Stderr
	scp.Stdout = os.Stdout
	err := scp.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// CopyProject copies the project from local machine to lectura
// Must be called from the assignment directory or subdirectories
func CopyProject(destination string) {
	copyPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if !strings.Contains(filepath.Base(copyPath), "assg") {
		copyPath = filepath.Dir(copyPath)
		if !strings.Contains(filepath.Base(copyPath), "assg") {
			log.Fatal("Copy must be done from the project directory or subdirectory. Please change your directory.")
		}
	}
	lecturaAddress := fmt.Sprintf("%s%s%s", GetUserName(), Lectura, destination)
	lecturaAddress = strings.TrimSuffix(lecturaAddress, "\n")
	scp := exec.Command("scp", "-r", copyPath, lecturaAddress)
	scp.Stdin = os.Stdin
	scp.Stderr = os.Stderr
	scp.Stdout = os.Stdout
	err = scp.Run()
	if err != nil {
		log.Fatal(err)
	}
}
