package go_cli

import (
	"fmt"
	"strings"
)

type Interface interface {
	//implemented with BaseInterface
	Initialize(path string) error        //initializes interface
	Run() error                          //runs interface, waits for user input
	ListAll() error                      //lists all commands in all modules
	ListModule(module_name string) error //lists all commands in given module

	Query(query string) error //NEEDED: where the commands are matched to functions in program
}

type BaseInterface struct {
	Commands    CommandList
	ProgramName string
}

// initializes new command line interface
func (b BaseInterface) Initialize(path string) error {
	fmt.Println("Initializing interface...")
	Commands := CommandList{Path: path}
	err := Commands.AddModule(path)
	if err != nil {
		return fmt.Errorf("unable to start interface: %s", err)
	}
	fmt.Println("Commands found, loaded. Starting interface...")
	return b.Run()
}

func (b BaseInterface) Run() error {
	var input string
	for {
		fmt.Print(b.ProgramName + ">: ")
		fmt.Scanln(&input)

		query := strings.Split(input, " ")
		len_query := len(query)
		//check built in commands first
		if len_query < 3 {
			if query[0] == "list" {
				if len_query == 2 {
					b.ListModule(query[1])
				} else {
					b.List()
				}
			}
		}
		//run query
	}
}

func (b *BaseInterface) Query(query string) error {

}
