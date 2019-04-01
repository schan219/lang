package cli

import (
	"fmt"
	"os"
	_"strconv"
	_"errors"
	_"path/filepath"

	"github.com/manifoldco/promptui"
)

type InitConfig struct {
	InputFile   		*os.File
	OutputPath  		string
	Contents        	[]byte
}

func Start() *InitConfig {
	var conf *InitConfig

	getInputFile(conf)
	getOutputPath(conf)

	return conf
}

func getInputFile(conf *InitConfig) {
	var fileName string
	var prompt promptui.Prompt
	var err error

	fmt.Println("Enter the file path as relative, i.e. ./file.lang")

	for {
		prompt = promptui.Prompt{Label: "Input File "}
		fileName, err = prompt.Run()

		// Skip opening if error
		if err != nil {
			panic(fmt.Sprintf("Prompt failed %v\n", err))
		}

		// Try to open, to see if there is a file.
		conf.InputFile, err = os.Open(fileName)

		if err == nil {
			break
		}
		fmt.Println("File not found. Please enter a valid file")
	}

}

func getOutputPath(conf *InitConfig) {

	prompt := promptui.Prompt{Label: "Output File Name"}

	outputFile, err := prompt.Run()
	prompt = promptui.Prompt{IsConfirm: true}

	if err != nil {
		panic(fmt.Sprintf("Prompt failed %v\n", err))
	}

	conf.OutputPath = outputFile
}