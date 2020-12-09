package settings
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/pkg/pickle"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/settings.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßDEFAULT_PROTOCOL := πg.InternStr("DEFAULT_PROTOCOL")
		ßFalse := πg.InternStr("False")
		ßHIGHEST_PROTOCOL := πg.InternStr("HIGHEST_PROTOCOL")
		ßImportError := πg.InternStr("ImportError")
		ß__doc__ := πg.InternStr("__doc__")
		ßbyref := πg.InternStr("byref")
		ßfmode := πg.InternStr("fmode")
		ßignore := πg.InternStr("ignore")
		ßprotocol := πg.InternStr("protocol")
		ßrecurse := πg.InternStr("recurse")
		ßsettings := πg.InternStr("settings")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.BaseException
		_ = πTemp004
		var πTemp005 *πg.Traceback
		_ = πTemp005
		var πTemp006 bool
		_ = πTemp006
		var πTemp007 *πg.Dict
		_ = πTemp007
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 8: """
			πF.SetLineno(8)
			// line 8: """
			πF.SetLineno(8)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nglobal settings for Pickler\n").ToObject()); πE != nil {
				continue
			}
			// line 12: try:
			πF.SetLineno(12)
			πF.PushCheckpoint(2)
			// line 13: from pickle import DEFAULT_PROTOCOL
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "pickle"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßDEFAULT_PROTOCOL); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDEFAULT_PROTOCOL.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp004, πTemp005 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label3
			}
			πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 14: except ImportError:
			πF.SetLineno(14)
		Label3:
			// line 15: from pickle import HIGHEST_PROTOCOL as DEFAULT_PROTOCOL
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "pickle"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßHIGHEST_PROTOCOL); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDEFAULT_PROTOCOL.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 17: settings = {
			πF.SetLineno(17)
			πTemp007 = πg.NewDict()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßDEFAULT_PROTOCOL); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ßprotocol.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ßbyref.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ßfmode.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ßrecurse.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ßignore.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πTemp007.ToObject()
			if πE = πF.Globals().SetItem(πF, ßsettings.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 27: del DEFAULT_PROTOCOL
			πF.SetLineno(27)
			if πE = πg.DelVar(πF, πF.Globals(), ßDEFAULT_PROTOCOL); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("settings", Code)
}

