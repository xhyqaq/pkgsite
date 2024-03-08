package godoc

import (
	"fmt"
	"github.com/google/uuid"
	"go/ast"
	"go/doc"
	"strings"
)

// FindOverloadFuncThenAdd  Find overloaded functions and update function names
func FindOverloadFuncThenAdd(p *Package) {
	var overloadFuncName = make(map[string]string)
	var overLoadFunc = make([]*ast.FuncDecl, 0)
	for _, f := range p.encPackage.Files {
		for _, decl := range f.AST.Decls {
			// First it will traverse the constants
			if g, ok := decl.(*ast.GenDecl); ok {
				for _, specs := range g.Specs {
					if vs, ok := specs.(*ast.ValueSpec); ok {
						for _, name := range vs.Names {
							if strings.Contains(name.Name, "Gopo_") {
								for _, v := range vs.Values {
									if bas, ok := v.(*ast.BasicLit); ok {
										overloadFuncs := strings.Split(bas.Value, ",")
										for i2 := range overloadFuncs {
											overloadFuncName[strings.ReplaceAll(overloadFuncs[i2], "\"", "")] = strings.TrimPrefix(name.Name, "Gopo_")
										}
									}
								}
							}
						}
					}
				}
			}
			// Secondly find the overloaded function and rename it
			// Because the same function name will be filtered, a specific character set is required
			if d, ok := decl.(*ast.FuncDecl); ok {
				if k, ok := overloadFuncName[d.Name.Name]; ok {
					name := fmt.Sprintf("%s!%s", k, uuid.NewString())
					newFunc := &ast.FuncDecl{}
					overLoadFunc = append(overLoadFunc, newFunc)
					newFunc.Body = d.Body
					newFunc.Doc = d.Doc
					newFunc.Recv = d.Recv
					newFunc.Type = d.Type
					newFunc.Recv = d.Recv
					newFunc.Type = d.Type
					newFunc.Name = &ast.Ident{}
					newFunc.Name.NamePos = d.Name.NamePos
					newFunc.Name.Obj = d.Name.Obj
					newFunc.Name.Name = name
				}
			}
		}
	}
	for _, newFunc := range overLoadFunc {
		for _, f := range p.encPackage.Files {
			f.AST.Decls = append(f.AST.Decls, newFunc)
		}
	}
}

// RestoreFuncName Overloaded function restores function name
func RestoreFuncName(funcs []*doc.Func) {
	for _, f := range funcs {
		f.Name = isOverloadFuncThenRestoresName(f.Name)
	}
}

// RestoreFuncDeclName Overloaded function decl restores function name
func RestoreFuncDeclName(funcs []*File) {
	for _, f := range funcs {
		for _, dec := range f.AST.Decls {
			if fun, ok := dec.(*ast.FuncDecl); ok {
				name := fun.Name.Name
				fun.Name.Name = isOverloadFuncThenRestoresName(name)
			}
		}
	}
}

func isOverloadFuncThenRestoresName(name string) string {
	if strings.Contains(name, "!") {
		return strings.SplitN(name, "!", 2)[0]
	}
	return name
}
