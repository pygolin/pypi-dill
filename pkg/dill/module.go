package dill
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/dill/info"
	// _ "github.com/pygolin/stdlib/importlib"
	// _ "github.com/pygolin/stdlib/dill/detect"
	// _ "github.com/pygolin/stdlib/dill/_objects"
	// _ "github.com/pygolin/stdlib/dill/_dill"
	// _ "github.com/pygolin/stdlib/collections"
	// _ "github.com/pygolin/stdlib/dill/objtypes"
	// _ "github.com/pygolin/stdlib/dill/settings"
	// _ "github.com/pygolin/stdlib/dill/source"
	// _ "github.com/pygolin/stdlib/dill/temp"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßCONTENTS_FMODE := πg.InternStr("CONTENTS_FMODE")
		ßDEFAULT_PROTOCOL := πg.InternStr("DEFAULT_PROTOCOL")
		ßFILE_FMODE := πg.InternStr("FILE_FMODE")
		ßFalse := πg.InternStr("False")
		ßHANDLE_FMODE := πg.InternStr("HANDLE_FMODE")
		ßHIGHEST_PROTOCOL := πg.InternStr("HIGHEST_PROTOCOL")
		ßImportError := πg.InternStr("ImportError")
		ßNone := πg.InternStr("None")
		ßOrderedDict := πg.InternStr("OrderedDict")
		ßPickler := πg.InternStr("Pickler")
		ßPicklingError := πg.InternStr("PicklingError")
		ßTrue := πg.InternStr("True")
		ßType := πg.InternStr("Type")
		ßUnpickler := πg.InternStr("Unpickler")
		ßUnpicklingError := πg.InternStr("UnpicklingError")
		ß__author__ := πg.InternStr("__author__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__license__ := πg.InternStr("__license__")
		ß__version__ := πg.InternStr("__version__")
		ß_extend := πg.InternStr("_extend")
		ß_revert_extension := πg.InternStr("_revert_extension")
		ßcheck := πg.InternStr("check")
		ßcitation := πg.InternStr("citation")
		ßcopy := πg.InternStr("copy")
		ßdetect := πg.InternStr("detect")
		ßdict := πg.InternStr("dict")
		ßdump := πg.InternStr("dump")
		ßdump_session := πg.InternStr("dump_session")
		ßdumps := πg.InternStr("dumps")
		ßextend := πg.InternStr("extend")
		ßfailures := πg.InternStr("failures")
		ßfind := πg.InternStr("find")
		ßkeys := πg.InternStr("keys")
		ßlicense := πg.InternStr("license")
		ßlist := πg.InternStr("list")
		ßload := πg.InternStr("load")
		ßload_session := πg.InternStr("load_session")
		ßload_types := πg.InternStr("load_types")
		ßloads := πg.InternStr("loads")
		ßmsg := πg.InternStr("msg")
		ßobjects := πg.InternStr("objects")
		ßodict := πg.InternStr("odict")
		ßpickle := πg.InternStr("pickle")
		ßpickles := πg.InternStr("pickles")
		ßpop := πg.InternStr("pop")
		ßreadme := πg.InternStr("readme")
		ßregister := πg.InternStr("register")
		ßregistered := πg.InternStr("registered")
		ßreload := πg.InternStr("reload")
		ßsettings := πg.InternStr("settings")
		ßsource := πg.InternStr("source")
		ßsucceeds := πg.InternStr("succeeds")
		ßtemp := πg.InternStr("temp")
		ßthis_version := πg.InternStr("this_version")
		ßtrace := πg.InternStr("trace")
		ßtypes := πg.InternStr("types")
		ßupdate := πg.InternStr("update")
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
		var πTemp007 *πg.BaseException
		_ = πTemp007
		var πTemp008 []πg.Param
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 8: goto Label8
			case 2: goto Label2
			case 11: goto Label11
			case 5: goto Label5
			case 14: goto Label14
			default: panic("unexpected function state")
			}
			// line 10: try:
			πF.SetLineno(10)
			πF.PushCheckpoint(2)
			// line 11: from .info import this_version as __version__
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "dill.info"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßthis_version); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__version__.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 12: from .info import readme as __doc__, license as __license__
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "dill.info"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßreadme); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßlicense); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__license__.ToObject(), πTemp003); πE != nil {
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
			// line 13: except ImportError:
			πF.SetLineno(13)
		Label3:
			// line 14: msg = """First run 'python setup.py build' to build dill."""
			πF.SetLineno(14)
			// line 14: msg = """First run 'python setup.py build' to build dill."""
			πF.SetLineno(14)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("First run 'python setup.py build' to build dill.").ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmsg.ToObject(), πg.NewStr("First run 'python setup.py build' to build dill.").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßmsg); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 15: raise ImportError(msg)
			πF.SetLineno(15)
			πE = πF.Raise(πTemp003, nil, nil)
			continue
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 17: __author__ = 'Mike McKerns'
			πF.SetLineno(17)
			if πE = πF.Globals().SetItem(πF, ß__author__.ToObject(), πg.NewStr("Mike McKerns").ToObject()); πE != nil {
				continue
			}
			// line 19: __doc__ = """
			πF.SetLineno(19)
			if πTemp003, πE = πg.ResolveGlobal(πF, ß__doc__); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πg.NewStr("\n").ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: __license__ = """
			πF.SetLineno(22)
			if πTemp003, πE = πg.ResolveGlobal(πF, ß__license__); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πg.NewStr("\n").ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__license__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: from ._dill import dump, dumps, load, loads, dump_session, load_session, \
			πF.SetLineno(25)
			if πTemp002, πE = πg.ImportModule(πF, "dill._dill"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßdump); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdump.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßdumps); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdumps.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßload); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßload.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßloads); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßloads.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßdump_session); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdump_session.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßload_session); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßload_session.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßPickler); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPickler.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßUnpickler); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnpickler.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßregister); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßregister.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßcopy); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcopy.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßpickle); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpickle.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßpickles); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpickles.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßcheck); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcheck.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßHIGHEST_PROTOCOL); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHIGHEST_PROTOCOL.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßDEFAULT_PROTOCOL); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDEFAULT_PROTOCOL.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßPicklingError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPicklingError.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßUnpicklingError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnpicklingError.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßHANDLE_FMODE); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHANDLE_FMODE.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßCONTENTS_FMODE); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCONTENTS_FMODE.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßFILE_FMODE); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFILE_FMODE.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 29: from . import source, temp, detect
			πF.SetLineno(29)
			if πTemp002, πE = πg.ImportModule(πF, "dill.source"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßsource.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "dill.temp"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtemp.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "dill.detect"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßdetect.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 32: from .settings import settings
			πF.SetLineno(32)
			if πTemp002, πE = πg.ImportModule(πF, "dill.settings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßsettings); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsettings.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 35: detect.trace(False)
			πF.SetLineno(35)
			πTemp002 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßdetect); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßtrace, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 37: try:
			πF.SetLineno(37)
			πF.PushCheckpoint(5)
			// line 38: from importlib import reload
			πF.SetLineno(38)
			if πTemp002, πE = πg.ImportModule(πF, "importlib"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßreload); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreload.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label4
		Label5:
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
				goto Label6
			}
			πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 39: except ImportError:
			πF.SetLineno(39)
		Label6:
			// line 40: try:
			πF.SetLineno(40)
			πF.PushCheckpoint(8)
			// line 41: from imp import reload
			πF.SetLineno(41)
			if πTemp002, πE = πg.ImportModule(πF, "imp"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßreload); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreload.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label7
		Label8:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp005 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label9
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 42: except ImportError:
			πF.SetLineno(42)
		Label9:
			// line 43: pass
			πF.SetLineno(43)
			πF.RestoreExc(nil, nil)
			goto Label7
		Label7:
			πF.RestoreExc(nil, nil)
			goto Label4
		Label4:
			// line 46: try:
			πF.SetLineno(46)
			πF.PushCheckpoint(11)
			// line 47: from collections import OrderedDict as odict
			πF.SetLineno(47)
			if πTemp002, πE = πg.ImportModule(πF, "collections"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßOrderedDict); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßodict.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label10
		Label11:
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
				goto Label12
			}
			πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 48: except ImportError:
			πF.SetLineno(48)
		Label12:
			// line 49: try:
			πF.SetLineno(49)
			πF.PushCheckpoint(14)
			// line 50: from ordereddict import OrderedDict as odict
			πF.SetLineno(50)
			if πTemp002, πE = πg.ImportModule(πF, "ordereddict"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßOrderedDict); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßodict.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label13
		Label14:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp005 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label15
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 51: except ImportError:
			πF.SetLineno(51)
		Label15:
			// line 52: odict = dict
			πF.SetLineno(52)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßodict.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label13
		Label13:
			πF.RestoreExc(nil, nil)
			goto Label10
		Label10:
			// line 53: objects = odict()
			πF.SetLineno(53)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßodict); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßobjects.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 60: from . import objtypes as types
			πF.SetLineno(60)
			if πTemp002, πE = πg.ImportModule(πF, "dill.objtypes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtypes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 62: def load_types(pickleable=True, unpickleable=True):
			πF.SetLineno(62)
			πTemp008 = make([]πg.Param, 2)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp008[0] = πg.Param{Name: "pickleable", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp008[1] = πg.Param{Name: "unpickleable", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("load_types", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__init__.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickleable *πg.Object = πArgs[0]; _ = µpickleable
				var µunpickleable *πg.Object = πArgs[1]; _ = µunpickleable
				var µ_objects *πg.Object = πg.UnboundLocal; _ = µ_objects
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 63: """load pickleable and/or unpickleable types to ``dill.types``
					πF.SetLineno(63)
					// line 78: from . import _objects
					πF.SetLineno(78)
					if πTemp002, πE = πg.ImportModule(πF, "dill._objects"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					µ_objects = πTemp001
					if πE = πg.CheckLocal(πF, µpickleable, "pickleable"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µpickleable); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 79: if pickleable:
					πF.SetLineno(79)
				Label1:
					// line 80: objects.update(_objects.succeeds)
					πF.SetLineno(80)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_objects, "_objects"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µ_objects, ßsucceeds, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßobjects); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label3
				Label2:
					// line 82: [objects.pop(obj,None) for obj in _objects.succeeds]
					πF.SetLineno(82)
					πTemp005 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__init__.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µ_objects, "_objects"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µ_objects, ßsucceeds, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp003 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp003 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp004 = !isStop
								} else {
									πTemp004 = true
									µobj = πTemp002
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 82: [objects.pop(obj,None) for obj in _objects.succeeds]
								πF.SetLineno(82)
								πTemp005 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp005[0] = µobj
								if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp005[1] = πTemp002
								if πTemp002, πE = πg.ResolveGlobal(πF, ßobjects); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(4)
								return πTemp002, nil
							Label4:
								πTemp006 = πSent
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
						continue
					}
					goto Label3
				Label3:
					if πE = πg.CheckLocal(πF, µunpickleable, "unpickleable"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µunpickleable); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					goto Label5
					// line 83: if unpickleable:
					πF.SetLineno(83)
				Label4:
					// line 84: objects.update(_objects.failures)
					πF.SetLineno(84)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_objects, "_objects"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µ_objects, ßfailures, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßobjects); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label6
				Label5:
					// line 86: [objects.pop(obj,None) for obj in _objects.failures]
					πF.SetLineno(86)
					πTemp005 = make([]πg.Param, 0)
					πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__init__.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µ_objects, "_objects"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µ_objects, ßfailures, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp003 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp003 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp004 = !isStop
								} else {
									πTemp004 = true
									µobj = πTemp002
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 86: [objects.pop(obj,None) for obj in _objects.failures]
								πF.SetLineno(86)
								πTemp005 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp005[0] = µobj
								if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp005[1] = πTemp002
								if πTemp002, πE = πg.ResolveGlobal(πF, ßobjects); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(4)
								return πTemp002, nil
							Label4:
								πTemp006 = πSent
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
						continue
					}
					goto Label6
				Label6:
					// line 87: objects.update(_objects.registered)
					πF.SetLineno(87)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_objects, "_objects"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µ_objects, ßregistered, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßobjects); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 88: del _objects
					πF.SetLineno(88)
					if πE = πg.CheckLocal(πF, µ_objects, "_objects"); πE != nil {
						continue
					}
					µ_objects = πg.UnboundLocal
					// line 90: [types.__dict__.pop(obj) for obj in list(types.__dict__.keys()) \
					πF.SetLineno(90)
					πTemp005 = make([]πg.Param, 0)
					πTemp007 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__init__.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								πTemp002 = πF.MakeArgs(1)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__dict__, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßkeys, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
								πTemp002[0] = πTemp004
								if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
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
									µobj = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp002 = πF.MakeArgs(1)
								πTemp002[0] = ßType.ToObject()
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µobj, ßfind, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.NE(πF, πTemp007, πTemp004); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 90: [types.__dict__.pop(obj) for obj in list(types.__dict__.keys()) \
								πF.SetLineno(90)
							Label4:
								// line 90: [types.__dict__.pop(obj) for obj in list(types.__dict__.keys()) \
								πF.SetLineno(90)
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp002[0] = µobj
								if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__dict__, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßpop, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πF.PushCheckpoint(6)
								return πTemp004, nil
							Label6:
								πTemp003 = πSent
								goto Label5
							Label5:
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
						continue
					}
					// line 93: reload(types)
					πF.SetLineno(93)
					πTemp002 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßreload); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßload_types.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 63: """load pickleable and/or unpickleable types to ``dill.types``
			πF.SetLineno(63)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("load pickleable and/or unpickleable types to ``dill.types``\n\n    ``dill.types`` is meant to mimic the ``types`` module, providing a\n    registry of object types.  By default, the module is empty (for import\n    speed purposes). Use the ``load_types`` function to load selected object\n    types to the ``dill.types`` module.\n\n    Args:\n        pickleable (bool, default=True): if True, load pickleable types.\n        unpickleable (bool, default=True): if True, load unpickleable types.\n\n    Returns:\n        None\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßload_types); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 95: def extend(use_dill=True):
			πF.SetLineno(95)
			πTemp008 = make([]πg.Param, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp008[0] = πg.Param{Name: "use_dill", Def: πTemp009}
			πTemp003 = πg.NewFunction(πg.NewCode("extend", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__init__.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µuse_dill *πg.Object = πArgs[0]; _ = µuse_dill
				var µ_revert_extension *πg.Object = πg.UnboundLocal; _ = µ_revert_extension
				var µ_extend *πg.Object = πg.UnboundLocal; _ = µ_extend
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 96: '''add (or remove) dill types to/from the pickle registry
					πF.SetLineno(96)
					// line 108: from ._dill import _revert_extension, _extend
					πF.SetLineno(108)
					if πTemp002, πE = πg.ImportModule(πF, "dill._dill"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ß_revert_extension); πE != nil {
						continue
					}
					µ_revert_extension = πTemp003
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ß_extend); πE != nil {
						continue
					}
					µ_extend = πTemp003
					if πE = πg.CheckLocal(πF, µuse_dill, "use_dill"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µuse_dill); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 109: if use_dill: _extend()
					πF.SetLineno(109)
				Label1:
					// line 109: if use_dill: _extend()
					πF.SetLineno(109)
					if πE = πg.CheckLocal(πF, µ_extend, "_extend"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_extend.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label3
				Label2:
					// line 110: else: _revert_extension()
					πF.SetLineno(110)
					if πE = πg.CheckLocal(πF, µ_revert_extension, "_revert_extension"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_revert_extension.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label3
				Label3:
					// line 111: return
					πF.SetLineno(111)
					πR = πg.None
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßextend.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 96: '''add (or remove) dill types to/from the pickle registry
			πF.SetLineno(96)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("add (or remove) dill types to/from the pickle registry\n\n    by default, ``dill`` populates its types to ``pickle.Pickler.dispatch``.\n    Thus, all ``dill`` types are available upon calling ``'import pickle'``.\n    To drop all ``dill`` types from the ``pickle`` dispatch, *use_dill=False*.\n\n    Args:\n        use_dill (bool, default=True): if True, extend the dispatch table.\n\n    Returns:\n        None\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßextend); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 113: extend()
			πF.SetLineno(113)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßextend); πE != nil {
				continue
			}
			if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
				continue
			}
			// line 115: def license():
			πF.SetLineno(115)
			πTemp008 = make([]πg.Param, 0)
			πTemp009 = πg.NewFunction(πg.NewCode("license", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__init__.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 116: """print license"""
					πF.SetLineno(116)
					// line 117: print (__license__)
					πF.SetLineno(117)
					πTemp001 = make([]*πg.Object, 1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß__license__); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 118: return
					πF.SetLineno(118)
					πR = πg.None
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßlicense.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 116: """print license"""
			πF.SetLineno(116)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("print license").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßlicense); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 120: def citation():
			πF.SetLineno(120)
			πTemp008 = make([]πg.Param, 0)
			πTemp010 = πg.NewFunction(πg.NewCode("citation", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__init__.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 121: """print citation"""
					πF.SetLineno(121)
					// line 122: print (__doc__[-485:-115])
					πF.SetLineno(122)
					πTemp001 = make([]*πg.Object, 1)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(485).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(115).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ß__doc__); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 123: return
					πF.SetLineno(123)
					πR = πg.None
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcitation.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 121: """print citation"""
			πF.SetLineno(121)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("print citation").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßcitation); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 125: del odict
			πF.SetLineno(125)
			if πE = πg.DelVar(πF, πF.Globals(), ßodict); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("dill", Code)
}

