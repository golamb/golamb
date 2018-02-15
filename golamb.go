package main

import (
	"fmt"
	"os"
	"os/exec"

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
	case "net":
		githubURL := "https://github.com/golamb/golamb-simple-template.git"
		cmd := exec.Command("git", "clone", githubURL, projectName)
		err := cmd.Run()
		if err == nil {
			removeUselessFile(dir, projectName)
			fmt.Printf("cd %s\ndep ensure\n", projectName)
		}
	case "simple":
		githubURL := "https://github.com/golamb/golamb-simple-template.git"
		cmd := exec.Command("git", "clone", githubURL, projectName)
		err := cmd.Run()
		if err == nil {
			removeUselessFile(dir, projectName)
			fmt.Printf("cd %s\ndep ensure\n", projectName)
		}
	default:
		fmt.Printf("Project not found!!")
	}
}

func removeUselessFile(dir string, projectName string) {
	gitFolderInProject := "./" + projectName + "/.git"
	rm := exec.Command("rm", "-rf", gitFolderInProject)
	rm.Run()
}
