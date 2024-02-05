package main

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	plugin := &Plugin{}
	protogen.Options{
		ParamFunc: func(name, value string) error {
			fmt.Fprintln(os.Stderr, name, value)
			return nil
		},
	}.Run(func(gen *protogen.Plugin) error {
		fmt.Fprintln(os.Stderr, gen.Request.GetParameter())
		for _, f := range gen.Files {
			if f.Generate {
				plugin.GenerateServices(gen, f)
			}
		}
		return nil
	})
}

type Plugin struct{}

func (s *Plugin) GenerateServices(gen *protogen.Plugin, file *protogen.File) {
	g := &lazyFile{gen: gen, in: file}
	for _, svc := range file.Services {
		g.P()
		g.P("type ", svc.GoName, " interface {")
		for _, met := range svc.Methods {
			g.P(fmt.Sprintf(
				"  %s(*%s) (*%s, error)",
				met.GoName,
				g.QualifiedGoIdent(met.Input.GoIdent),
				g.QualifiedGoIdent(met.Output.GoIdent),
			))
		}
		g.P("}")
	}
}

type lazyFile struct {
	gen *protogen.Plugin
	in  *protogen.File
	out *protogen.GeneratedFile
}

func (l *lazyFile) ensure() {
	if l.out != nil {
		return
	}
	filename := l.in.GeneratedFilenamePrefix + "_svc.pb.go"
	l.out = l.gen.NewGeneratedFile(filename, l.in.GoImportPath)
	l.out.P("package ", l.in.GoPackageName)
	l.out.P()
}

func (l *lazyFile) QualifiedGoIdent(ident protogen.GoIdent) string {
	l.ensure()
	return l.out.QualifiedGoIdent(ident)
}

func (l *lazyFile) P(args ...interface{}) {
	l.ensure()
	l.out.P(args...)
}