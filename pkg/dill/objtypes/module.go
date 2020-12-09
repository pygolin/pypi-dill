package objtypes
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/pkg/dill"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/objtypes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßNameError := πg.InternStr("NameError")
		ß__doc__ := πg.InternStr("__doc__")
		ß_type := πg.InternStr("_type")
		ßkeys := πg.InternStr("keys")
		ßobjects := πg.InternStr("objects")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 bool
		_ = πTemp005
		var πTemp006 bool
		_ = πTemp006
		var πTemp007 *πg.BaseException
		_ = πTemp007
		var πTemp008 *πg.Traceback
		_ = πTemp008
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			case 2: goto Label2
			case 5: goto Label5
			default: panic("unexpected function state")
			}
			// line 8: """
			πF.SetLineno(8)
			// line 8: """
			πF.SetLineno(8)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nall Python Standard Library object types (currently: CH 1-15 @ 2.7)\nand some other common object types (i.e. numpy.ndarray)\n\nto load more objects and types, use dill.load_types()\n").ToObject()); πE != nil {
				continue
			}
			// line 16: from dill import objects
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "dill"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßobjects); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßobjects.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßobjects); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßkeys, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
				continue
			}
			πF.PushCheckpoint(2)
			πTemp005 = false
		Label1:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp005 {
				πF.PopCheckpoint()
				goto Label3
			}
			if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
				isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
				if exc != nil {
					πE = exc
				} else if isStop {
					πE = nil
					πF.RestoreExc(nil, nil)
				}
				πTemp006 = !isStop
			} else {
				πTemp006 = true
				if πE = πF.Globals().SetItem(πF, ß_type.ToObject(), πTemp003); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp006 {
				continue
			}
			πF.PushCheckpoint(1)            
			// line 18: exec("%s = type(objects['%s'])" % (_type,_type))
			πF.SetLineno(18)
			πE = πF.RaiseType(πg.NotImplementedErrorType, "exec is not available on Grumpy. Maybe never be.")
			continue
			continue
		Label2:
			if πE != nil || πR != nil {
				continue
			}
		Label3:
			// line 20: del objects
			πF.SetLineno(20)
			if πE = πg.DelVar(πF, πF.Globals(), ßobjects); πE != nil {
				continue
			}
			// line 21: try:
			πF.SetLineno(21)
			πF.PushCheckpoint(5)
			// line 22: del _type
			πF.SetLineno(22)
			if πE = πg.DelVar(πF, πF.Globals(), ß_type); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label4
		Label5:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label6
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 23: except NameError:
			πF.SetLineno(23)
		Label6:
			// line 24: pass
			πF.SetLineno(24)
			πF.RestoreExc(nil, nil)
			goto Label4
		Label4:
		}
		return nil, πE
	})
	πg.RegisterModule("objtypes", Code)
}

