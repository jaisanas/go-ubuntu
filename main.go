package shellcommand

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
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

func SystemDiskUsage() map[string]string {
	var keyValArr []map[string]string
	cmd := exec.Command("df", "-h")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return nil
    }

    fmt.Println(string(stdout))
	s := string(stdout);
	
	res := strings.Split(s, "\n")
	for i := range res {
		row := strings.Split(res[i], " ")
		counter := 0;
		rowMap := make(map[string]string)
		for j := range row {
			//ignore the empty string after splitted by the blank space
			if row[j] != "" {
				switch counter {
				case 0:
					rowMap["filesystem"] = row[j];
				case 1:
					rowMap["size"] = row[j];
				case 2:
					rowMap["used"] = row[j];
				case 3:
					rowMap["avail"] = row[j];
				case 4:
					rowMap["capacity"] = row[j];
				case 5:
					rowMap["iused"] = row[j];
				case 6:
					rowMap["ifree"] = row[j];
				case 7:
					rowMap["mountedOn"] = row[j];
				}
				counter = counter + 1
			}
 		}

		// only append the data and ignore the header table
		if rowMap["avail"] != "Avail" {
			keyValArr = append(keyValArr, rowMap)
		}
	}

	return keyValArr[]
}