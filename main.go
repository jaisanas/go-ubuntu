package shellcommand

import (
	"bytes"
	"fmt"
	"os/exec"
)

var stdoutBuf, stderrBuf bytes.Buffer


func ListDir() {
	cmd := exec.Command("ls")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(string(stdout))
}

func SystemDiskUsage() {
	cmd := exec.Command("df", "-h")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(string(stdout))
}