package main

import (
    "log"
    "io/ioutil"
    "fmt"
    "bufio"
    "os"
    "os/exec"

    "github.com/go-by-example/examples"
)

func listExamples() {
    dir := "./examples/"

    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    var name []byte
    for k, file := range files {
        name = []byte(file.Name())
        name = name[0:len(name)-3] // remove .go

        if k != 0 {
            fmt.Print(" | ")
        }
        fmt.Print(string(name))
    }

    fmt.Println()
}

func clearTerm() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

type fn func()
func mapExamples() (map[string]fn) {
    function := map[string]fn {
        "hello": examples.Hello,
        "values": examples.Values,
        "variables": examples.Variables,
        "constants": examples.Constants,
        "for": examples.For,
        "ifelse": examples.IfElse,
        "switch": examples.Switch,
        "arrays": examples.Arrays,
        "slices": examples.Slices,
        "maps": examples.Maps,
        "range": examples.Range,
        "functions": examples.Functions,
        "pointers": examples.Pointers,
        "structs": examples.Structs,
        "methods": examples.Methods,
        "interfaces": examples.Interfaces,
        "interfaces2": examples.Interfaces2,
        "errors": examples.Errors,
        "goroutines": examples.GoRoutines,
        "templates": examples.Templates,
    }

    return function
}

func printHeader(input string) {
    fmt.Println("#######################")
    fmt.Printf("Executing '%s'\n", input)
    fmt.Println("#######################")
    fmt.Println()
}

func printFooter() {
    fmt.Println("#######################")
    fmt.Println()

    fmt.Print("press enter to continue")
}

func main() {

    buf := bufio.NewReader(os.Stdin)

    for true {

        fmt.Println("Please type one of the following:")
        listExamples()

        fmt.Print("> ")
        input, err := buf.ReadBytes('\n')
        inputStr := string(input[0:len(input)-1]) // remove last char (enter)

        if err != nil {
            clearTerm()

            fmt.Printf("Error receiving input '%s': %s\n\n", inputStr, err.Error())

            continue
        }

        function := mapExamples()

        cb, has := function[inputStr]
        if has {
            printHeader(inputStr)

            cb()

            printFooter()
            buf.ReadBytes('\n')

            clearTerm()
        } else {
            clearTerm()

            fmt.Printf("No example available for '%s'\n\n", inputStr)
        }

    }
}
