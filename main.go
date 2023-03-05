package shellcommand

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

var stdoutBuf, stderrBuf bytes.Buffer


func Ls() {
	cmd := exec.Command("ls")

	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Start()  // Starts command asynchronously

	if err != nil {
		fmt.Printf(err.Error())
	}
}