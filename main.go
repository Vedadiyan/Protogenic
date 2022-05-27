package main

import (
	"bytes"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			GenerateServices(gen, f)
		}
		return nil
	})
}

func GenerateServices(plugin *protogen.Plugin, file *protogen.File) {
	for _, service := range file.Services {
		filename := file.GeneratedFilenamePrefix + "_" + service.GoName + ".pb.go"
		g := plugin.NewGeneratedFile(filename, file.GoImportPath)
		g.P("package ", file.GoPackageName)
		for _, method := range service.Methods {
			methodName := ToCamelCase(method.GoName)
			inputMessageName := ToCamelCase(method.Input.GoIdent.GoName)
			g.P("type ", method.GoName, " func(", inputMessageName, " *", method.Input.GoIdent.GoName, ") (*", method.Output.GoIdent.GoName, ", error)")
			g.P("func ", method.GoName, "Client ", "(conn *nats.EncodedConn, timeout time.Duration) ", method.GoName, "{")
			g.P(`panic("")`)
			g.P("}")
			g.P("func ", method.GoName, "Server ", "(conn *nats.EncodedConn, ", methodName, " ", method.GoName, ", logger Logger) *nats.Subscription", "{")
			g.P(`panic("")`)
			g.P("}")
		}
	}
}

func ToCamelCase(input string) string {
	buffer := bytes.NewBufferString("")
	buffer.WriteString(strings.ToLower(string(input[0])))
	for i := 1; i < len(input); i++ {
		buffer.WriteRune(rune(input[i]))
	}
	return buffer.String()
}
