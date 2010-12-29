package main

import (
	"goresque"
	"os"
	"fmt"
	"syscall"
)

type FlavorSaver struct {
	Id int
}


func (self *FlavorSaver) perform(args []interface{}) os.Error {
	for i, val := range args {
		fmt.Println("arg #", i, val)
	}
	fmt.Println("I would have done something big here!")
	return nil
}


func main() {

	//instantiate a new Resque
	r := goresque.NewResque("127.0.0.1", 6379, 0)
	for { // loop forever
		//pop a job off the queue the easy way
		job2, err2 := r.Reserve("flavors")
		if err2 != nil {
			fmt.Print(".")         // makes it look very Wargams-ish
			syscall.Sleep(1000000) // this is a .1 second sleep 
		} else {
			fmt.Println("return:", job2)

			switch job2.Class {
			case "FlavorSaver":
				is := new(FlavorSaver)
				is.perform(job2.Args)
			default:
				fmt.Println("Incorrect Class!")
			}
		}
	}
}
