package shellcommand

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

var stdoutBuf, stderrBuf bytes.Buffer


func ListDir() (error, []string) {
	var res []string
    cmd := exec.Command("ls")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return err, nil
    }

    // Print the output
    fmt.Println(string(stdout))
	s := string(stdout)
	rowString := strings.Split(s, "\n")
	for i := range rowString {
		res = append(res, rowString[i])
	}

	return nil, res
}

func SystemDiskUsage() (error, []map[string]string) {
	var keyValArr []map[string]string
	cmd := exec.Command("df", "-h")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return err, nil
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

	return nil, keyValArr
}

func Pwd() (error, *string) {
	
    cmd := exec.Command("pwd")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return err, nil
    }

    // Print the output
    fmt.Println(string(stdout))
	res := string(stdout)

	return nil, &res
}

func Cp(src string, dest string) (error, *string) {
	
    cmd := exec.Command("cp", src, dest)
    stdout, err := cmd.Output()
	
    if err != nil {
        fmt.Println(err.Error())
        return err, nil
    }    
	res := string(stdout)
	fmt.Println(stdout);
	return nil, &res
}

func Cd(dest string) (error, *string) {
	// Execute 'cd' command followed by 'pwd' to verify the directory change in the same shell process
	cmd := exec.Command("bash", "-c", fmt.Sprintf("cd %s && pwd", dest))

	// Capture the output
	output, err := cmd.Output()
	if err != nil {
		return err, nil
	}

	// Return the output as a string
	res := string(output)
	return nil, &res
}