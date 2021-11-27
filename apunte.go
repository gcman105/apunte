package main

import (
    "os"
    "fmt"
    "gcman105/apunte/util"

    //"log"
)

func main() {
    progName := os.Args[0]
    fmt.Println("Arg 0 is", progName)
    for i, a := range os.Args[1:] {
        fmt.Printf("Arg %d is %s\n", i+1, a) 
    }

    config, err := util.LoadConfig(".")

    if err != nil { // Handle errors reading the config file
        panic(fmt.Errorf("fatal error config file: %w ", err))
    }

    fmt.Println(config)

    inpath := config.INpath
    outpath := config.OUTpath

    fmt.Printf("IN=%s  OUT=%s\n", inpath, outpath)
    fmt.Printf("A=%s  B=%s\n", config.Paths.Apath, config.Paths.Bpath)
}                                     
