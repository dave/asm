package main

import (
	"fmt"

	"strings"

	"regexp"

	"os"

	"github.com/dave/asm/generator/x86spec"
	"github.com/dave/jennifer/jen"
)

const UNSAFE_PACKAGE = "github.com/dave/asm/generator/unsafe-stub"

func main() {
	if err := run(); err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		os.Exit(1)
	}
}
func run() error {

	config := &x86spec.Config{
		DebugPage: "120,121,122,123,124,125,126,127,128,129,130",
	}
	instructions := x86spec.Load(config)

	//

	f := jen.NewFile("x86")
	f.ImportName(UNSAFE_PACKAGE, "unsafe")

	for _, ins := range instructions {
		ins := utils{Instruction: ins}
		fmt.Printf("%#v\n", ins.Instruction)

		// Add the comment with function name and instruction description
		f.Commentf("%s: %s", ins.function(), ins.Desc)

		for i, p := range ins.params() {
			f.Commentf("%s: %s", p, ins.Args[i])
		}

		f.Commentf("%s#page=%d", config.URL, ins.Page)

		// Add the Go function
		f.Func().Id(ins.function()).ParamsFunc(func(g *jen.Group) {
			for i, p := range ins.params() {
				g.Id(p).Do(func(s *jen.Statement) {
					if i == len(ins.Args)-1 {
						s.Interface() // all params are `interface{}`, so only add type to last item
					}
				})
			}
		}).Block(
			jen.Qual(UNSAFE_PACKAGE, "Asm").CallFunc(func(g *jen.Group) {
				g.Lit(ins.instruction())
				if len(ins.Args) == 0 {
					g.Nil()
				} else {
					for _, p := range ins.params() {
						g.Id(p)
					}
				}
			}),
		)
	}
	if err := f.Save("./x86/generated.go"); err != nil {
		return err
	}
	return nil
}

type utils struct {
	*x86spec.Instruction
}

var functionRegex = regexp.MustCompile(`[^a-zA-z0-9]+`)

func (u utils) function() string {
	return functionRegex.ReplaceAllString(u.GoSyntax, "_")
}

func (u utils) instruction() string {
	if strings.ContainsAny(u.GoSyntax, " ") {
		return u.GoSyntax[:strings.Index(u.GoSyntax, " ")]
	}
	return u.GoSyntax
}

func (u utils) params() []string {
	if !strings.Contains(u.Syntax, " ") {
		return nil
	}
	argsString := u.Syntax[strings.Index(u.Syntax, " ")+1:] // remove the instruction from the start
	argsString = strings.Replace(argsString, ",", "", -1)   // remove commas
	var args []string
	for _, raw := range strings.Fields(argsString) {
		args = append(args, functionRegex.ReplaceAllString(raw, "_"))
	}
	return args
}
