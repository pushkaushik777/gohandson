package main

import (
	"fmt"
	os "os/exec"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		out, err := os.Command("ls", "ltr").Output()
		if err != nil {
			fmt.Printf("%s", err)
		}
		o := string(out[:])
		fmt.Println(o)
	}
}
