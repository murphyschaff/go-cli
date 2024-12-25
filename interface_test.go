package go_cli

import (
	"io"
	"os"
	"testing"
)

func TestInterface(t *testing.T) {

	test_interface := BaseInterface{ProgramName: "Test Interface", CommandPath: "./test_commands.json"}

	reader, writer := io.Pipe()
	os.Stdin = reader
	defer func() { os.Stdin = os.Stdin }()

}
