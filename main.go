package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	timeBtw := time.Now().UnixMicro()
	var timeArray []int
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		if string(b) == " " {
			timeBtw = time.Now().UnixMicro() - timeBtw;
			timeArray = append(timeArray, int(timeBtw))
			timeBtw = time.Now().UnixMicro()
		}
		fmt.Print(string(b))
	}
}

