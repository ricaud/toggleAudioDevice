package main

import (
	"fmt"
	_"embed"
	"os/exec"
	"os"
	"io/ioutil"
	"time"
	"path/filepath"
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
			time.Sleep(3*time.Second)
			os.Exit(1)
		}
	}
}

func fileExists(path string) bool {
	 _, err := os.Stat(path); 
	if err != nil {
		return false 
	} 
	return true
}

func toggleDevices() {
	var cmd *exec.Cmd

	if fileExists("toggle.flag") {
		cmd = exec.Command("nircmd.exe", "setdefaultsounddevice", device1)
		os.Remove("toggle.flag")
	} else {
		cmd = exec.Command("nircmd.exe", "setdefaultsounddevice", device2)
		_, err := os.Create("toggle.flag")
		if err != nil {
			printAndSleep(err.Error())
		}
		
	}
	cmd.Run()
}

func changeWD(){
	wd, err := os.Executable()
	if err != nil {
		printAndSleep(err.Error())
	}

	err = os.Chdir(filepath.Dir(wd))
	if err != nil {
		printAndSleep(err.Error())
	}

}

func main() {
	changeWD()
	unpackDeps()
	toggleDevices()
}

func printAndSleep(msg string) {
	fmt.Println(msg)
	time.Sleep(3*time.Second)
}