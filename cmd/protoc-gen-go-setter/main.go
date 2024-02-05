package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/octavore/protoc-tools/setterpb"
)

func main() {
	plugin := &Plugin{}
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				plugin.Generate(gen, f)
			}
		}
		return nil
	})
}

type Plugin struct {
}

func (s *Plugin) Generate(gen *protogen.Plugin, file *protogen.File) {
	g := &lazyFile{gen: gen, in: file}
	isProto2 := file.Desc.Syntax() == protoreflect.Proto2
	for _, msg := range file.Messages {
		msgName := msg.GoIdent.GoName
		setAllOpt, ok := proto.GetExtension(msg.Desc.Options(), setterpb.E_AllFields).(bool)
		genAll := ok && setAllOpt
		for _, f := range msg.Fields {
			if f.Oneof != nil {
				fieldOpt, ok := proto.GetExtension(f.Oneof.Desc.Options(), setterpb.E_OneofField).(*setterpb.SetterFieldOptions)
				includeField := ok && fieldOpt != nil && fieldOpt.Include
				excludeField := ok && fieldOpt != nil && fieldOpt.Exclude
				if !genAll && !includeField {
					continue
				}
				if genAll && excludeField {
					continue
				}

				oneofStructIdent := msgName + "_" + f.GoName
				g.P("func (m *", msgName, ") Set", f.GoName, "(v ", goType(g, f, false), ") {")
				g.P("  m.", f.Oneof.GoName, " = &", oneofStructIdent, "{", f.GoName, ": v}")
				g.P("}")
				g.P()
			} else {
				fieldOpt, ok := proto.GetExtension(f.Desc.Options(), setterpb.E_Field).(*setterpb.SetterFieldOptions)
				includeField := ok && fieldOpt != nil && fieldOpt.Include
				excludeField := ok && fieldOpt != nil && fieldOpt.Exclude
				if !genAll && !includeField {
					continue
				}
				if genAll && excludeField {
					continue
				}

				g.P("func (m *", msgName, ") Set", f.GoName, "(v ", goType(g, f, isProto2), ") {")
				g.P("  m.", f.GoName, " = v")
				g.P("}")
				g.P()
			}
		}

		for _, f := range msg.Oneofs {
			fieldOpt, ok := proto.GetExtension(f.Desc.Options(), setterpb.E_OneofField).(*setterpb.SetterFieldOptions)
			includeField := ok && fieldOpt != nil && fieldOpt.Include
			excludeField := ok && fieldOpt != nil && fieldOpt.Exclude
			if !genAll && !includeField {
				continue
			}
			if genAll && excludeField {
				continue
			}
			g.P("func (m *", msgName, ") Set", f.GoName, "(v is", f.GoIdent, ") {")
			g.P("  m.", f.GoName, " = v")
			g.P("}")
			g.P()
		}
	}
}

func goType(g *lazyFile, f *protogen.Field, mustPtr bool) string {
	if f.Desc.IsMap() {
		mapKey := f.Message.Fields[0]
		mapValue := f.Message.Fields[1]
		return fmt.Sprintf(
			"map[%s]%s",
			goType(g, mapKey, false),
			goType(g, mapValue, false),
		)
	}

	if f.Message != nil {
		returnTypePrefix := "*"
		if f.Desc.Cardinality() == protoreflect.Repeated {
			returnTypePrefix = "[]*"
		}
		return returnTypePrefix + g.QualifiedGoIdent(f.Message.GoIdent)
	}

	if f.Enum != nil {
		returnTypePrefix := "*"
		if f.Desc.Cardinality() == protoreflect.Repeated {
			returnTypePrefix = "[]*"
		}
		return returnTypePrefix + g.QualifiedGoIdent(f.Enum.GoIdent)
	}

	returnTypePrefix := ""
	if f.Desc.Cardinality() == protoreflect.Repeated {
		returnTypePrefix = "[]"
	} else if mustPtr {
		returnTypePrefix = "*"
	}
	return returnTypePrefix + f.Desc.Kind().String()

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
	filename := l.in.GeneratedFilenamePrefix + "_setter.pb.go"
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
