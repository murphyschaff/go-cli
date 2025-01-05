package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Command struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Usage       string `json:"usage"`
	Function    string `json:"function"`
	APIPath     string `json:"apipath"`
}

type CommandModule struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Path        string     `json:"path"`
	Commands    []*Command `json:"commands"`
}

type CommandList struct {
	Path    string
	Modules []*CommandModule
}

// save CLI commands to base file path
func (l *CommandList) Save() error {
	file, err := os.Create(l.Path)
	if err != nil {
		return fmt.Errorf("unable to open path file: %s", err)
	}
	defer file.Close()

	for _, module := range l.Modules {
		json, err := json.Marshal(module)
		if err != nil {
			return fmt.Errorf("unable to marshal json for path: %s", err)
		}
		file.Write(append(json, '\n'))
	}
	return nil
}

// loads the specified command file
func NewCommandList(path string) (*CommandList, error) {
	l := &CommandList{Path: path}
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open command file %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		var module CommandModule
		err := json.Unmarshal(line, &module)
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshal json from file: %s", err)
		}
		l.Modules = append(l.Modules, &module)
	}
	return l, nil
}

// add new commands to list, saves list
func (l *CommandList) AddModule(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("unable to open command file %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		var module *CommandModule
		err := json.Unmarshal(line, module)
		if err != nil {
			return fmt.Errorf("unable to unmarshal json from file: %s", err)
		}
		l.Modules = append(l.Modules, module)
	}
	return l.Save()
}

// removes commands associated with a given module from the list
func (l *CommandList) RemoveModule(module_name string) error {
	for i, module := range l.Modules {
		if module.Name == module_name {
			l.Modules = append(l.Modules[:i], l.Modules[i+1:]...)
			break
		}
	}
	return l.Save()
}

// searches list for the name of a command, returns an error if nothing can be found
func (l *CommandList) FindCommand(module_name string, command_name string) (*Command, error) {

	ret_command := &Command{Name: "nil"}

	for _, module := range l.Modules {
		if module.Name == module_name {
			for _, command := range module.Commands {
				if command.Name == command_name {
					ret_command = command
					break
				}
			}
			break
		}
	}

	if ret_command.Name != "nil" {
		return ret_command, nil
	} else {
		return ret_command, fmt.Errorf("unable to find command with name '%s' from module '%s'", command_name, module_name)
	}
}
