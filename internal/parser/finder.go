package parser

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
)

func findInterfaceInPackage(pkg *types.Package, name string) (*types.Interface, error) {
	scope := pkg.Scope()
	if scope == nil {
		return nil, fmt.Errorf("no scope")
	}
	obj := scope.Lookup(name)
	if obj == nil {
		return nil, fmt.Errorf("no object")
	}
	tp := obj.Type()
	if tp == nil {
		return nil, fmt.Errorf("no type")
	}
	iface, isInterface := tp.Underlying().(*types.Interface)
	if !isInterface {
		return nil, fmt.Errorf("not interface")
	}
	return iface.Complete(), nil
}

func findPackagePo(pkgs []*packages.Package) (*types.Package, error) {
	if len(pkgs) == 0 {
		return nil, fmt.Errorf("empty package list")
	}
	var po *types.Package
	packages.Visit(pkgs,
		// pre
		func(pkg *packages.Package) bool {
			pkg.Types.Complete()
			for _, imp := range pkg.Types.Imports() {
				if imp.Path() == poPackageName {
					po = imp
					return false
				}
			}
			return true
		}, // post
		func(pkg *packages.Package) {})
	if po == nil {
		return nil, fmt.Errorf("did not find po")
	}
	return po, nil
}
