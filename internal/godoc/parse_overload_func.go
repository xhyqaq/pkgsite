package godoc

import (
	"go/ast"
	"go/doc"
	"sort"
	"strings"
)

// FindOverloadFuncThenAdd  Find overloaded functions and update function names
func FindOverloadFuncThenAdd(d *doc.Package) {
	if len(d.Funcs) == 0 {
		return
	}
	var overloadFuncName = make(map[string]string)
	for _, constO := range d.Consts {
		for _, name := range constO.Names {
			if strings.Contains(name, "Gopo_") {
				for _, spec := range constO.Decl.Specs {
					if vs, ok := spec.(*ast.ValueSpec); ok {
						for _, name := range vs.Names {
							if strings.Contains(name.Name, "Gopo_") {
								n := strings.TrimPrefix(name.Name, "Gopo_")
								for _, v := range vs.Values {
									if bas, ok := v.(*ast.BasicLit); ok {
										overloadFuncs := strings.Split(bas.Value, ",")
										for i2 := range overloadFuncs {
											overloadFuncName[strings.ReplaceAll(overloadFuncs[i2], "\"", "")] = n
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	var overloadFunc = make([]*doc.Func, 0)
	for _, funcO := range d.Funcs {
		if name, ok := overloadFuncName[funcO.Name]; ok {
			newFunc := &doc.Func{}
			newFunc.Doc = funcO.Doc
			newFunc.Recv = funcO.Recv
			newFunc.Orig = funcO.Orig
			newFunc.Decl = funcO.Decl
			newFunc.Level = funcO.Level
			newFunc.Examples = funcO.Examples
			newFunc.Name = name
			overloadFunc = append(overloadFunc, newFunc)
		}
	}
	for _, funs := range overloadFunc {
		d.Funcs = append(d.Funcs, funs)
	}

	sort.Slice(d.Funcs, func(i, j int) bool {
		return d.Funcs[i].Name < d.Funcs[j].Name
	})
}
