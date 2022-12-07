package main

import (
    "os"
    "path"
    "testing"
)


func TestWget(t *testing.T) {
    fileName := "google.com"
    Wget("https://google.com")    
    wd, _ := os.Getwd()
    _, err := os.Open(path.Join(wd, fileName))
    if err != nil {
        t.Errorf("A file with name  was unable to be opened")    
    }

}