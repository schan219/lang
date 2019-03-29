package cli
import (
	"fmt"
	"os"
	_"strconv"
	_"errors"
	_"path/filepath"

	"github.com/manifoldco/promptui"
)

func Start() {

	var outputFile string
	var prompt promptui.Prompt
	var err error

	getInputFile()

	prompt = promptui.Prompt{Label: "Output File Name"}

	outputFile, err = prompt.Run()
	prompt = promptui.Prompt{IsConfirm: true}

	if err != nil {
		fmt.Printf("Prompt failed %s, %v\n", outputFile, err)
		return
	}


}

func getInputFile () {

	var fileName string
	var prompt promptui.Prompt
	var err error
	fmt.Println("Enter the file path as relative, i.e. ./file.lang")

	for {
		prompt = promptui.Prompt{Label: "Input File "}
		fileName, err = prompt.Run()

		// Skip opening if error
		if err != nil {
			continue
		}

		// Try to open, to see if there is a file.
	    _, err := os.Open(fileName)

		if err == nil {
			break
		}
		fmt.Println("File not found. Please enter a valid file")
	}
}