package gtools

import (
    "log"
    "os"
)

var (
    Logger    *log.Logger
    accesslog *log.Logger
)


func createLogfile(filename string) (*os.File, error) {
    var f *os.File
    var e error
    f, e = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0644))
    if e != nil {
        panic(e)
        return nil, e
    }
    return f, nil
}


