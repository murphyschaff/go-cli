package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("create_commands, edit and create commands for use in the go_cli interface")
	fmt.Println("Specify name of program")
	name := validate("Program Name")
	path := "./" + name + ".json"

	list := CommandList{Path: path}

}

// validates a given value for option
func validate(option string) string {
	validate := true
	var input string
	fmt.Scanln(&input)
	for validate {
		fmt.Printf("Are you sure you want '%s' for '%s'\n[Y]Yes, [N]No", option, input)
		fmt.Scanln(&input)
		input = strings.ToUpper(input)
		if input == "Y" {
			validate = false
		} else if input == "N" {
			fmt.Printf("Please enter a new value for %s", option)
			fmt.Scanln(&value)
		} else {
			fmt.Println("Invalid key entered.")
		}
	}
	return value
}

func YN(message string) bool {
	var input string
	for {
		fmt.Printf("%s\n[Y]Yes [N]No\n", message)
		fmt.Scanln(&input)
		input = strings.ToUpper(input)
		if input == "Y" {
			return true
		} else if input == "N" {
			return false
		} else {
			fmt.Println("Invalid key entered")
		}
	}
}

func AddCommand(module *CommandModule) {
	fmt.Println("Please enter the name of the command")
	name := validate("Command Name")
	fmt.Printf("Please enter a description for %s\n", name)
	description := validate("Command Description")
	fmt.Printf("Please enter the usage for %s\n", name)
	usage := validate("Command Usage")
	fmt.Printf("Please enter the function name %s will call\n", name)
	function := validate("Command Function")

	command := Command{Name: name, Description: description, Usage: usage, Function: function}

	module.Commands = append(module.Commands, command)

}

func AddModule(list *CommandList) {
	fmt.Println("Enter a name for the module")
	name := validate("Module Name")
	fmt.Printf("Enter a description for the %s module\n", name)
	description := validate("Description")

	module := CommandModule{Name: name, Description: description}
	for {
		AddCommand(&module)
		if YN("Would you like to add another command to the " + name + " module?") {
			break
		}
	}
}
