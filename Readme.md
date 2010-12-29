
# GoResque

GoResque is a Go implementation of the core libraries and worker code for the Resque task management system.  See https://github.com/defunkt/resque for more details on Resque.

GoResque allows you to create workers in Go that take items from the queue and process them, just as the regular Resque workers written in Ruby would.  Go workers are compiled to native code, and have the potential to be many times faster than Ruby workers for some operations.

## Changelog

*  12/29/2010 - Initial revision.  Library to enable reading from Resque queue, and enable creating Resque workers.

## Philosophy

Ruby is dynamic, Go isn't.

In Ruby/Resque, the worker code instantiates the appropriate worker class based on the class that was enqueued.  This class is used to perform the work/task.

In Go, the worker is statically instantiated. See the main() in the example/worker.go as an example. I kept the class.perform(args) concept, so it'd look familiar to people writing both Ruby and Go workers.

* type FlavorSaver struct {
	Id int
 }
*
* func (self *FlavorSaver) perform(args []interface{}) os.Error {
	for i, val := range args {
		fmt.Println("arg #", i, val)
	}
	fmt.Println("I would have done something big here!")
	return nil
 }
*

* func main() {

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
*

## todo

There are many things left to do (in no particular order)

*	Logging at the worker level
*	Tests
*	Clean up method/function return values for consistency
*	Better(any) error handling, putting task back on queue, or in failed queue


## running tests

You could run the tests with gotest, once I get around to writing some.

## author

Brian Ketelsen  https://github.com/bketelsen  bketelsen@gmail.com

## credits

This library uses Redis.go, forked from https://github.com/hoisie/redis.go

I forked it and compiled against my fork to guard against unknown breaking changes.  As of the time of writing, I have made no modifications to my fork.

## license

MIT

