package main

import (
	"fmt"
	_"embed"
	"os/exec"
	"os"
	"io/ioutil"
)

//go:embed nircmd.exe
var nircmdBytes []byte

var device1 string = "Headphones"
var device2 string = "ASUS PB277"

func unpackDeps() {
	//unpack nircmd.exe dependancy if it doesnt already exist
	if !fileExists("nircmd.exe") {
		err := ioutil.WriteFile("nircmd.exe", nircmdBytes, 0770)
		if err != nil {
			fmt.Println("[ERROR] could not unpack nircmd.exe depencancy: ", err.Error())
			os.Exit(1)
		}
	}
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true 
	} else {
		return false  
	}
}

func toggleDevices() {
	var cmd *exec.Cmd
	if fileExists("toggle.flag") {
		cmd = exec.Command("nircmd.exe", "setdefaultsounddevice", "Headphones")
		os.Remove("toggle.flag")
	} else {
		cmd = exec.Command("nircmd.exe", "setdefaultsounddevice", "ASUS PB277")
		os.Create("toggle.flag")
	}
	cmd.Run()
}

func main() {
	toggleDevices()
}

