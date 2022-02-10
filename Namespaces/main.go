package main

import "os"
import "fmt"
// docker        run image <command> <parameters>
// go run main.go run       <command> <parameters>

func main()  {
   switch os.Args[1] {
   case "run":
	  run()
	default:
	  panic("bad command")
   }	
}

func run()  {
	fmt.Printf("Running %v\n",os.Args[2:])

}
func must(err error)  {
	if err != nil {
		panic(err)
	}
}