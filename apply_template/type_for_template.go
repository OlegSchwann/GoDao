package apply_template

import (
	"fmt"
	"os"
	"text/template"
)

type ReturnValueType uint8

const (
	// no return value
	Exec ReturnValueType = 1
	// 1 row
	QueryRow ReturnValueType = 2
	// multi rows
	Query ReturnValueType = 3
)

type Variable struct {
	Name string
	Type string
}

type Function struct {
	Name            string
	SQL             string
	ReturnValueType ReturnValueType
	InputArguments  []Variable
	OutputArguments []Variable // without last (err error)
}

type DotType struct {
	PackageName string
	Packages    []string
	Functions   []Function
}

func Render(tt DotType) error {
	t, err := template.New("GoDao").Parse(templateEmbedded)
	if err != nil {
		return fmt.Errorf("text/template.Template.Parse(): %w", err)
	}

	err = t.Execute(os.Stdout, tt)
	if err != nil {
		return fmt.Errorf("text/template.Template.Execute(): %w", err)
	}

	return nil
}
