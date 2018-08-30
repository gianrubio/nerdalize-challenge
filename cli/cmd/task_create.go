package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"net/http"

	"github.com/gianrubio/nerdalize-challenge/api/models"
	"github.com/spf13/cobra"
)

const (
	TASK_LIST = "http://localhost:3000/api/v1/task"
)

var DockerImage string

func init() {
	taskCreate.Flags().StringVarP(&DockerImage, "docker-image", "", "", "Docker Image")
	taskCmd.AddCommand(taskCreate)

}

var taskCreate = &cobra.Command{
	Use:   "create",
	Short: "Create a task",
	Run: func(cmd *cobra.Command, args []string) {
		createTasks()
	},
}

func createTasks() {
	task := models.Task{
		DockerImage: DockerImage,
	}
	jsonStr, err := json.Marshal(task)
	check(err)

	resp, err := http.Post(TASK_LIST, "application/json", bytes.NewBuffer(jsonStr))
	check(err)
	if resp.StatusCode != http.StatusCreated {
		fmt.Println(fmt.Printf("ERROR: Wrong status code returned, expected: 201, got %v", resp.StatusCode))
		os.Exit(1)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Printf("Task successfully created\n%v", string(content))
}
func check(err error) {
	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %v", err))
		os.Exit(1)
	}
}
