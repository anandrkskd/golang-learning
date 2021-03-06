package main

import (
    "fmt"
    "os"
)

func main() {
    if fileExists("~/.meshery/meshery.yaml") {
        fmt.Println("Example file exists")
    } else {
        fmt.Println("Example file does not exist (or is a directory)")
    }
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}
