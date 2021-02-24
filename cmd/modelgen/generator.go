package main

import (
	"bytes"
	"fmt"
	"log"
	//"os"
	"go/format"
	"strings"
	"unicode"
)

/*Base Generator struct*/
type Generator struct {
	buf     bytes.Buffer
	pkgName string
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) PrintHeader() {
	g.Printf("// Code generated by \"ovsmodel\"; DO NOT EDIT.\n")
	// TODO: Add version
	g.Printf("\n")
	g.Printf("package %s\n", g.pkgName)
}

/*
Input is a variable definition line split in strings
*/
func (g *Generator) PrintVars(vars [][]string) {
	// print vars
	g.Printf("var (\n")
	for _, varPart := range vars {
		g.Printf("\t%s\n", strings.Join(varPart, " "))
	}
	g.Printf(")\n")

}
func (g *Generator) PrintImports(imports [][]string) {
	g.Printf("import (\n")
	for _, dep := range imports {
		if len(dep) == 1 {
			g.Printf("\t\"%s\"\n", dep[0])
		} else if len(dep) == 2 {
			g.Printf("\t%s \"%s\"\n", dep[0], dep[1])
		}
	}
	g.Printf(")\n")
}

func (g *Generator) Format() []byte {
	gen, err := format.Source(g.buf.Bytes())
	if err != nil {
		log.Printf("WARNING:", err)
	}
	return gen
}

/* tunnel_key -> TunnelKey */
func camelCase(field string) string {
	capNext := true
	orig := []rune(field)
	for i, c := range orig {
		if capNext {
			orig[i] = unicode.ToUpper(c)
		}
		capNext = c == '_' || c == '-'
	}
	return strings.ReplaceAll(string(orig), "_", "")
}