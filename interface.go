package cli

import (
	"fmt"
	"strings"
)

type Interface interface {
	//implemented with BaseInterface
	Initialize(path string) error    //initializes interface
	Run() error                      //runs interface, waits for user input
	ListAll()                        //lists all commands in all modules
	ListModule(module_name string)   //lists all commands in given module
	ListCommand(command_name string) //lists the command and usage

	Query(query string) error //NEEDED: where the commands are matched to functions in program
}

type BaseInterface struct {
	Commands    *CommandList
	ProgramName string
	CommandPath string
}

// creates new BaseInterfae objet
func NewInterface(name string, path string) (*BaseInterface, error) {
	commands, err := Load(path)
	if err != nil {
		return nil, fmt.Errorf("unable to start interface: %s", err)
	}

	init := BaseInterface{Commands: commands, ProgramName: name, CommandPath: path}
	return &init, nil
}

// loop that runs the CLI interface
func (b *BaseInterface) Run() error {
	var input string
	for {
		fmt.Print(b.ProgramName + ">: ")
		fmt.Scanln(&input)

		query := strings.Split(input, " ")
		len_query := len(query)
		//check built in commands first
		if len_query <= 3 && query[0] == "list" {
			if len_query == 3 && query[1] == "command" {
				b.ListCommand(query[3])
			} else if len_query == 2 {
				b.ListModule(query[2])
			} else {
				b.List()
			}
		}
		if len_query == 1 && query[0] == "exit" {
			//exit CLI
			return nil
		}
		//run query
		err := b.Query(query)
		if err != nil {
			return fmt.Errorf("failed to run query: %s", err)
		}
	}
}

// list the command and usage for each command in the list
func (b *BaseInterface) List() {
	for _, module := range b.Commands.Modules {
		fmt.Printf("Module: %s\nDescription: %s\n", module.Name, module.Description)
		for _, command := range module.Commands {
			fmt.Printf("\tCommand: %s\n\tDescription: %s\n\tUsage: %s\n\t", command.Name, command.Description, command.Usage)
		}
	}
}

// list all the commands in each module
func (b *BaseInterface) ListModule(module_name string) {
	for _, module := range b.Commands.Modules {
		if module.Name == module_name {
			fmt.Printf("Module: %s\nDescription: %s\n", module.Name, module.Description)
			for _, command := range module.Commands {
				fmt.Printf("\tCommand: %s\n\tDescription: %s\n\tUsage: %s\n\t", command.Name, command.Description, command.Usage)
			}
			return
		}
	}
	fmt.Println("Unable to find module in list.")
}

// lists all commands with given name in the list
func (b *BaseInterface) ListCommand(command_name string) {
	for _, module := range b.Commands.Modules {
		for _, command := range module.Commands {
			if command.Name == command_name {
				fmt.Printf("Module: %s\nDescription: %s", module.Name, module.Description)
				fmt.Printf("\tCommand: %s\n\tDescription: %s\n\tUsage: %s\n\t", command.Name, command.Description, command.Usage)
			}
		}
	}
}

// matches given commands to functions
func (b *BaseInterface) Query(query []string) error {
	fmt.Println("Implement me.")
	return nil
}
