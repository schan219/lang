package cli

import (
	"fmt"
	"os"
	"io/ioutil"
	_"strconv"
	_"errors"
	_"path/filepath"

	"github.com/manifoldco/promptui"
)

type InitConfig struct {
	InputFile   		string
	Contents			[]byte
	OutputPath  		string
}

func Start() *InitConfig {
	var conf *InitConfig = &InitConfig{}

	getInputFile(conf)
	getOutputPath(conf)

	return conf
}

func getInputFile(conf *InitConfig) {
	var prompt promptui.Prompt
	var err error

	fmt.Println("Enter the file path as relative, i.e. ./file.lang")

	for {
		prompt = promptui.Prompt{Label: "Input File "}
		conf.InputFile, err = prompt.Run()

		// Skip opening if error
		if err != nil {
			panic(fmt.Sprintf("Prompt failed %v\n", err))
		}

		// Try to open, to see if there is a file.
		_, err = os.Open(conf.InputFile)

		if err == nil {
			break
		}
		fmt.Println("File not found. Please enter a valid file")
	}

	// We must've found a valid file, so we read it in
	file, _ := os.Open(conf.InputFile)
	conf.Contents, err = ioutil.ReadAll(file)

	// If there is an error, throw it
	if err != nil {
		panic(fmt.Sprintf("Failed reading file: %s", conf.InputFile))
	}
}

func getOutputPath(conf *InitConfig) {

	prompt := promptui.Prompt{Label: "Output File "}

	outputFile, err := prompt.Run()
	prompt = promptui.Prompt{IsConfirm: true}

	if err != nil {
		panic(fmt.Sprintf("Prompt failed %v\n", err))
	}

	conf.OutputPath = outputFile
}