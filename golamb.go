package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"bufio"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "golamb"
	app.Usage = "CLI for create project of lambda with golang from https://github.com/golamb"
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello friend!\nThis is Project for create aws-lambda-go\nRead Doc at https://github.com/aws/aws-lambda-go/\n")
		return nil
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Thar be no %q here.\n", command)
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"-i"},
			Usage:   "initial project",
			Action: func(c *cli.Context) error {
				args := c.Args()
				if len(args) == 2 {
					projectType := c.Args().First()
					projectName := args[1]
					dir, _ := os.Getwd()
					createProject(dir, projectName, projectType)
				} else {
					fmt.Println("Init fail pls use \"golamb init [project_type] [project_name]")
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func createProject(dir string, projectName string, projectType string) {
	switch projectType {
	case "api-gateway":
		githubURL := "https://github.com/golamb/golamb-api-gateway-template.git"
		err := cloneProjectTemplate(githubURL, projectName)
		afterClone(dir, projectName, err)

	case "simple":
		githubURL := "https://github.com/golamb/golamb-simple-template.git"
		err := cloneProjectTemplate(githubURL, projectName)
		afterClone(dir, projectName, err)
	default:
		fmt.Printf("Template 404 not found!!")
	}
}

func cloneProjectTemplate(githubURL, projectName string) error {
	cmd := exec.Command("git", "clone", githubURL, projectName)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		fmt.Println(errb.String())
	}
	return err
}

func afterClone(dir, projectName string, err error) {
	if err == nil {
		removeUselessFile(dir, projectName)
		updateDeployScritp(projectName)
		fmt.Printf("cd %s\ndep ensure\n", projectName)
	}
}

func removeUselessFile(dir string, projectName string) {
	gitFolderInProject := "./" + projectName + "/.git"
	exec.Command("rm", "-rf", gitFolderInProject).Run()
}

func updateDeployScritp(projectName string) {
	reads, _ := readLines("./" + projectName + "/deploy.sh")
	var newDeploy []string
	for _, readline := range reads {
		matched, _ := regexp.MatchString("<PROJECT_NAME>", readline)
		if matched {
			newDeploy = append(newDeploy, strings.Replace(readline, "<PROJECT_NAME>", projectName, -1))
		} else {
			newDeploy = append(newDeploy, readline)
		}
	}
	writeLines(newDeploy, "./"+projectName+"/deploy.sh")
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
