package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	taskCmd.AddCommand(taskList)
}

var taskList = &cobra.Command{
	Use:   "list",
	Short: "List nerdalize tasks",
	Run: func(cmd *cobra.Command, args []string) {
		listTasks()
	},
}

func listTasks() {
	resp, err := http.Get(TASK_LIST)
	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %v", err))
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(fmt.Printf("ERROR: Wrong status code returned, expected: 200, got %v", resp.StatusCode))
		os.Exit(1)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %v", err))
		os.Exit(1)
	}
	fmt.Printf("Task successfully created\n%v", string(content))
}
