package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/murphyschaff/go-cli"
)

func main() {
	fmt.Println("manage_commands: Create and manage command lists for use in go-cli")

	args := os.Args[1:]

	if len(args) == 0 {
		//create mode
		fmt.Print("Enter a name for the program: ")
		name := validate("Program Name")
		path := "./" + name + ".json"

		list := cli.CommandList{Path: path}

		for {
			AddModule(&list)
			if !YN("Would you like to add another module?") {
				break
			}
		}
		fmt.Println("Saving to file...")
		list.Save()
		fmt.Println("Creation complete. Closing program")
	} else if len(args) == 1 {
		//edit mode
		file, err := os.Stat(args[0])
		if err != nil && file.IsDir() {
			fmt.Println("Invalid filetype entered. Please enter a filepath to a JSON file")
			os.Exit(1)
		}
		list := cli.Load(args[0])

		Edit(&list)
		list.Save()
		fmt.Println("Changes saved, exiting program.")

	} else {
		fmt.Println("Invalid input arguments entered. Please enter the path to the file you want to edit, or nothing to create a new file")
		os.Exit(1)
	}
}

// validates a given value for option
func validate(option string) string {
	validate := true
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	var choice string
	scanner.Scan()
	choice = scanner.Text()
	for validate {
		fmt.Printf("Are you sure you want '%s' for '%s'\n[Y]Yes, [N]No: ", choice, option)
		scanner.Scan()
		input = scanner.Text()
		input = strings.ToUpper(input)
		if input == "Y" {
			validate = false
		} else if input == "N" {
			fmt.Printf("Please enter a new value for %s: ", option)
			scanner.Scan()
			choice = scanner.Text()
		} else {
			fmt.Println("Invalid key entered.")
		}
	}
	return choice
}

func YN(message string) bool {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s\n[Y]Yes [N]No: ", message)
		scanner.Scan()
		input = scanner.Text()
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

func GetInt(rng int) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := scanner.Text()
		val, err := strconv.Atoi(input)
		if err == nil && val <= rng {
			return val
		} else {
			fmt.Printf("Please enter an integer value less than %d", rng)
		}
	}
}

func AddCommand(module *cli.CommandModule) {
	fmt.Print("Please enter the name of the command: ")
	name := validate("Command Name")
	fmt.Printf("Please enter a description for %s: ", name)
	description := validate("Command Description")
	fmt.Printf("Please enter the usage for %s: ", name)
	usage := validate("Command Usage")
	fmt.Printf("Please enter the function name %s will call: ", name)
	function := validate("Command Function")

	command := cli.Command{Name: name, Description: description, Usage: usage, Function: function}

	module.Commands = append(module.Commands, command)

}

func AddModule(list *cli.CommandList) {
	fmt.Print("Enter a name for the module: ")
	name := validate("Module Name")
	fmt.Printf("Enter a description for the %s module: ", name)
	description := validate("Description")

	module := cli.CommandModule{Name: name, Description: description}
	for {
		AddCommand(&module)
		if !YN("Would you like to add another command to the " + name + " module?") {
			list.Modules = append(list.Modules, module)
			break
		}
	}
}

func Edit(list *cli.CommandList) {
	edit := true
	for edit {
		fmt.Println("Please select the value next to the item you wish to edit:")
		counter := 1
		for _, module := range list.Modules {
			fmt.Printf("[%d]: %s\n", counter, module.Name)
			counter++
		}
		fmt.Printf("[%d]: Exit\n", counter)
		choice := GetInt(counter)
		if choice == counter {
			if YN("Would you like to save changes and exit?") {
				edit = false
			}
		} else {
			EditModule(&list.Modules[choice-1])
		}
	}
}

func EditModule(module *cli.CommandModule) {
	for {
		fmt.Println("Please select the value next to the item you wish to edit:")
		fmt.Printf("[1]: Module Name (%s)\n[2]: Module Description (%s)", module.Name, module.Description)
		counter := 3
		for _, commands := range module.Commands {
			fmt.Printf("[%d]: %s\n", counter, commands.Name)
			counter++
		}
		fmt.Printf("[%d]: Exit Module", counter)
		choice := GetInt(counter)
		if choice == 1 {
			module.Name = validate("Module Name")
		} else if choice == 2 {
			module.Description = validate("Module Description")
		} else if choice == counter {
			return
		} else {
			EditCommand(&module.Commands[choice-1])
		}
	}
}

func EditCommand(command *cli.Command) {
	for {
		fmt.Printf("[1]: Command Name (%s)\n[2]: Command Description (%s)\n[3]: Command Usage (%s)\n[4]: Command Function (%s)\n[5]: Exit\n", command.Name, command.Description, command.Usage, command.Function)
		choice := GetInt(5)
		switch choice {
		case 1:
			command.Name = validate("Command Name")
			break
		case 2:
			command.Description = validate("Command Description")
			break
		case 3:
			command.Usage = validate("Command Usage")
			break
		case 4:
			command.Function = validate("Command Function")
			break
		case 5:
			return
		}
	}
}
