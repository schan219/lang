package help

import (
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"

	"fmt"
	"io/ioutil"
)

const DEFINITION_FILE = "./defs.yaml"

type Description struct {
	Name string `yaml:"name"`
	Signature string `yaml:"signature"`
	Description string `yaml:"description"`
}

func DescribeCommand(cmd string) {
	defs := readDefs()
 
	for _,def := range defs {
		if cmd == def.Name {
			showCMD(def.Name, def.Signature, def.Description)
		}
	}
}

func showCMD(cmd string, sig string, desc string) {
	red := color.New(color.FgRed).PrintfFunc()
	yellow := color.New(color.FgYellow).PrintfFunc()

	red("Command: ")
	fmt.Println(cmd)
	red("Fn Signature: ")
	fmt.Println(sig)
	yellow("Description: ")
	fmt.Println(desc)
	fmt.Printf("\n")
}

func readDefs() []Description {
	data, err := ioutil.ReadFile(DEFINITION_FILE)

	if err != nil {
		panic("Error reading in the definitions file..")
	}

	desc := &[]Description{}
	yaml.Unmarshal(data, desc)

	return *desc
}

