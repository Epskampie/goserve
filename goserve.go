package main

import (
	"flag"
	"golivereload/print"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var cyan func(a ...interface{}) string = color.New(color.FgCyan).SprintFunc()
var red func(a ...interface{}) string = color.New(color.FgRed).SprintFunc()
var yellow func(a ...interface{}) string = color.New(color.FgYellow).SprintFunc()

func main() {

	setupFlags(flag.CommandLine)
	flag.Parse()

	print.ShowDebug = params.debug

	// Change rootPath to working dir if not set
	cwd, err := os.Getwd()
	if err != nil {
		print.Fatal(red("Could not get current working dir.", err))
	}
	if params.rootPath == "" {
		params.rootPath = cwd
	}

	if !strings.HasSuffix(params.rootPath, string(os.PathSeparator)) {
		params.rootPath += string(os.PathSeparator)
	}

	// Check rootPath
	fileInfo, err := os.Stat(params.rootPath)
	if err != nil {
		print.Fatal(red(err))
	}
	if !fileInfo.IsDir() {
		print.Fatal(cyan(params.rootPath), red("is not a directory"))
	}

	port := strconv.Itoa(params.intPort)

	print.Line("Serving files from:", cyan(params.rootPath), "on:", cyan("http://localhost:"+port))
	http.Handle("/", http.FileServer(http.Dir(params.rootPath)))

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			print.Fatal("Port", port, "already in use")
		} else {
			print.Fatal("ListenAndServe: " + err.Error())
		}
	}
}
