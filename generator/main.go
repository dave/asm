package main

import (
	"log"
	"github.com/dave/jennifer/jen"
			)

func main(){
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {
	f := jen.NewFile("x86")
	if err := f.Save("./x86/generated.go"); err != nil {
		return err
	}
	return nil
}
