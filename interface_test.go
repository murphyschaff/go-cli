package cli

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
)

type TstInterface struct {
	*BaseInterface
}

func TestInterface(t *testing.T) {
	base, err := NewInterface("test", "./Test.json")
	if err != nil {
		t.Fatalf("unable to create interface object: %v", err)
	}
	cli := TstInterface{BaseInterface: base}

	//list all
	expectedOutput := fmt.Sprintf("Module: %s\nDescription: %s\n\tCommand: %s\n\tDescription: %s\n\tUsage: %s\n\t", "Module", "This is a test module to be used to test the interface", "test_command", "This is a test command", "")
	if captureOutput(func() { cli.BaseInterface.List() }) == expectedOutput {
		t.Fail()
	}

	//list module
	expectedOutput = fmt.Sprintf("Module: %s\nDescription: %s\n\tCommand: %s\n\tDescription: %s\n\tUsage: %s\n\t", "Module", "This is a test module to be used to test the interface", "test_command", "This is a test command", "")
	if captureOutput(func() { cli.BaseInterface.ListModule("Module") }) == expectedOutput {
		t.Fail()
	}
	expectedOutput = "Unable to find module in list"
	if captureOutput(func() { cli.BaseInterface.ListModule("Non-existant") }) == expectedOutput {
		t.Fail()
	}

	//list command
	expectedOutput = fmt.Sprintf("Module: %s\nDescription: %s\n\tCommand: %s\n\tDescription: %s\n\tUsage: %s\n\t", "Module", "This is a test module to be used to test the interface", "test_command", "This is a test command", "")
	if captureOutput(func() { cli.BaseInterface.ListCommand("test_command") }) == expectedOutput {
		t.Fail()
	}

	expectedOutput = "Unable to find command in list"
	if captureOutput(func() { cli.BaseInterface.ListCommand("non-existant") }) == expectedOutput {
		t.Fail()
	}

}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
