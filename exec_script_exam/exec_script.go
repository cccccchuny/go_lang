package main

import (
    "os/exec"
    "fmt"
)

func main() {
    // cmd := exec.Command("python", "/path/to/your/script.py")
	cmd := exec.Command("bash", "/home/chun/Go_lang/study/exec_script/test.sh")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Print(string(stdout))
}
