package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"strconv"
)

func errorAndExit(err error) {
	if err == nil {
		return
	}
	log.Error(err)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		errorAndExit(errors.New("Syntax: dungeon width height"))
	}

	width, err := strconv.Atoi(os.Args[1])
	errorAndExit(err)

	height, err := strconv.Atoi(os.Args[2])
	errorAndExit(err)

	dungeon := NewDungeon(width, height)
	dungeon.Generate()
	fmt.Println(dungeon.String())
}
