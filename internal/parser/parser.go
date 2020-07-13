package parser

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
)

const (
	poPackageName = "github.com/go-po/po"
)

func ParsePkg(pkgName string) error {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo,
	}, pkgName)
	if err != nil {
		return err
	}

	// find po
	poPkg, err := findPackagePo(pkgs)
	if err != nil {
		return err
	}

	commandHandlerInterface, err := findInterfaceInPackage(poPkg, "CommandHandler")
	if err != nil {
		return err
	}

	handlerInterface, err := findInterfaceInPackage(poPkg, "Handler")
	if err != nil {
		return err
	}

	var commands []types.Object
	var views []types.Object
	packages.Visit(pkgs,
		// pre
		func(pkg *packages.Package) bool {
			for _, name := range pkg.Types.Scope().Names() {
				obj := pkg.Types.Scope().Lookup(name)
				V := types.NewPointer(obj.Type())
				if types.Implements(V, commandHandlerInterface) {
					commands = append(commands, obj)
					continue
				}
				if types.Implements(V, handlerInterface) {
					views = append(views, obj)
					continue
				}
			}
			return true
		},
		// post
		func(pkg *packages.Package) {})

	for _, handler := range commands {
		p("COMMAND: %s", handler.Name())
	}

	for _, view := range views {
		p("VIEW: %s %s", view.Pkg(), view.Name())
	}

	return nil
}

func p(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println()
}
