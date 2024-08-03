package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"syscall"
)

func statx(file string) {
	var stat syscall.Stat_t
	err := syscall.Stat(file, &stat)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	logStatx(file, stat)
}

func logStatx(file string, stat syscall.Stat_t) {
	reflectValue := reflect.ValueOf(stat)
    typeOfExample := reflectValue.Type()

	fmt.Printf("------------------------------\nFile: %v\n", file)
    for i := 0; i < reflectValue.NumField(); i++ {
        field := reflectValue.Field(i)
        fmt.Printf("%-13s -> %v \n", typeOfExample.Field(i).Name, field.Interface())
    }
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("You should provide at least one file")
	}

	for _, arg := range args[1:] {
		statx(arg)
	}
}