package main

import (
	"fmt"

	"strings"

	"regexp"

	"os"

	"sort"

	"reflect"

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

var argnames = map[string]string{
	"":              "arg",
	"1":             "v1",
	"3":             "v3",
	"<XMM0>":        "xmm",
	"AL":            "al",
	"AL/AX/EAX/RAX": "al",
	"AX":            "ax",
	"AX/EAX/RAX":    "ax",
	"CL":            "cl",
	"CS":            "cs",
	"DS":            "ds",
	"DX":            "dx",
	"EAX":           "eax",
	"ES":            "es",
	"EVEX.vvvv":     "evex",
	"FS":            "fs",
	"GS":            "gs",
	"ModRM:r/m":     "rm",
	"ModRM:reg":     "reg",
	"ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)": "reg",
	"Moffs":  "moffs",
	"Offset": "offset",
	"RAX":    "rax",
	"SS":     "ss",
	"ST(0)":  "st0",
	"ST(i)":  "sti",
	"VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)": "vex",
	"VEX.vvvv":                 "vex",
	"VectorReg(R): VSIB:index": "v",
	"imm16":                    "imm",
	"imm8":                     "imm",
	"imm8/16/32":               "imm",
	"imm8/16/32/64":            "imm",
	"imm8[7:4]":                "imm",
	"implicit XMM0":            "implicit",
	"iw":                       "iw",
	"opcode + rd":              "opcode",
	"vsib":                     "vsib",
	"vvvv":                     "vvvv",
}

func run() error {

	config := &x86spec.Config{
	//DebugPage: "214",
	}
	instructions := x86spec.Load(config)

	allargs := map[string]map[string]bool{}

	grouped := map[string]map[string][]*utils{} // name -> op/en -> instructions
	for _, ins := range instructions {
		if ins.Page == 0 {
			continue
		}
		if ins.Name == "zmm3/m512," {
			// TODO: That's this?
			continue
		}
		if grouped[ins.Name] == nil {
			grouped[ins.Name] = map[string][]*utils{}
		}
		op := strings.Replace(ins.OpEn, "-", "_", -1)
		op = strings.Replace(op, " ", "_", -1)
		grouped[ins.Name][op] = append(grouped[ins.Name][op], &utils{ins})
		//fmt.Printf("*** %s %s %#v\n", ins.Name, ins.OpEn, ins)
	}

	fmt.Println("*** ALLARGS")
	for _, v := range keys(allargs) {
		fmt.Printf("\"%s\": \"\"\n", v)
	}
	fmt.Println("ALLARGS ***")

	f := jen.NewFile("x86")
	f.ImportName(UNSAFE_PACKAGE, "unsafe")

	for _, name := range keys(grouped) {
		byop := grouped[name]
		for _, op := range keys(grouped[name]) {
			instructions := grouped[name][op]
			ins := instructions[0]

			fname := name
			if len(byop) > 1 {
				fname = name + "_" + op
			}

			// Add the comment with function name and instruction description
			f.Comment(fname)
			descriptions := map[string]bool{}
			for _, ins := range instructions {
				if descriptions[ins.Desc] {
					continue
				}
				f.Comment(ins.Desc)
				descriptions[ins.Desc] = true
			}
			/*
				buf := &bytes.Buffer{}
				w := tabwriter.NewWriter(buf, 1, 4, 1, ' ', 0)
				for _, ins := range instructions {
					fmt.Fprintf(w, "//\t%s\t%s\t%s\n", ins.Opcode, ins.Desc)
				}
				w.Flush()
				f.Comment(buf.String())
			*/
			if len(ins.params()) > 0 {
				f.Comment("")
				for i, p := range ins.params() {
					f.Commentf("%s: %s", p, ins.Args[i])
				}
			}
			f.Comment("")
			f.Commentf("Documentation: %s#page=%d", config.URL, ins.Page)

			// Add the Go function
			f.Func().Id(fname).ParamsFunc(func(g *jen.Group) {
				for i, p := range ins.params() {
					g.Id(p).Do(func(s *jen.Statement) {
						if i == len(ins.Args)-1 {
							s.Interface() // all params are `interface{}`, so only add type to last item
						}
					})
				}
			}).Block(
				jen.Qual(UNSAFE_PACKAGE, "Asm").CallFunc(func(g *jen.Group) {
					g.Lit(ins.Name)
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

	var args []string
	for _, arg := range u.Args {
		arg = strings.TrimSuffix(arg, " (r)")
		arg = strings.TrimSuffix(arg, " (w)")
		arg = strings.TrimSuffix(arg, " (r, w)")
		if argnames[arg] == "" {
			panic("unknown arg " + arg + " in " + u.Name)
		}
		args = append(args, argnames[arg])
	}

	return args
}

func keys(i interface{}) []string {
	var keys []string
	for _, v := range reflect.ValueOf(i).MapKeys() {
		keys = append(keys, v.String())
	}
	sort.Strings(keys)
	return keys
}
