package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	changeType := flag.String("t", "", "Change type, win:wsl to win, wsl:win to wsl, def:empty(auto)")
	path := flag.String("p", "", "Path to change")
	flag.Parse()

	if *path == "" {
		fmt.Println("Use wslpath -h for help.")
		return
	}

	var result string
	if *changeType == "wsl" {
		result = win2unix(*path)
	}else if *changeType == "win" {
		result = unix2win(*path)
	}else{
		var mntIndex = strings.Index(*path, "/mnt/")
		if mntIndex == 0 {
			result = unix2win(*path)
		}else{
			result = win2unix(*path)
		}
	}
	fmt.Print(result)
}

func win2unix(path string) string {
	path = strings.Replace(path, "\\\\", "/", 1)
	path = strings.ReplaceAll(path, "\\", "/")
	return "/mnt/" + path
}

func unix2win(path string) string {
	path = strings.Replace(path, "/mnt/", "", 1)
	path = strings.ReplaceAll(path, "/", "\\")
	path = strings.Replace(path, ":\\", ":\\\\", 1)
	return path
}