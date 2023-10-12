package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/Masterminds/sprig"
)

const templateFile = "entity.gotmpl"

type EntField struct {
	Name string
	Type string
}

type EntParam struct {
	SchemaName string
	Fields     []*EntField
}

func main() {
	targets := []any{
		ent.RelyingParty{},
		ent.User{},
	}

	for _, v := range targets {
		tgt := reflect.TypeOf(v)

		param := &EntParam{
			SchemaName: tgt.Name(),
		}

		for i := 0; i < tgt.NumField(); i++ {
			param.Fields = append(param.Fields, &EntField{
				Name: tgt.Field(i).Name,
				Type: tgt.Field(i).Type.String(),
			})
		}

		var buf bytes.Buffer
		t := template.Must(template.New(templateFile).Funcs(sprig.FuncMap()).ParseFiles(fmt.Sprintf("gen/%s", templateFile)))
		if err := t.Execute(&buf, param); err != nil {
			log.Fatal(err)
		}

		var out bytes.Buffer
		out.WriteString(fmt.Sprintf("// Code generated by app/entity/gen/gen.go; DO NOT EDIT.\npackage entity\n"))
		out.Write(buf.Bytes())

		body, err := format.Source(out.Bytes())
		if err != nil {
			log.Fatal(err)
		}

		var output = fmt.Sprintf("%s_gen.go", strings.ToLower(tgt.Name()))

		if err = os.WriteFile(output, body, 0644); err != nil {
			log.Fatal(err)
		}
	}
}
