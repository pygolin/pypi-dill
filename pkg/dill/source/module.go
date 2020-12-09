package source
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/pkg/dill"
	// _ "github.com/pygolin/stdlib/pkg/detect"
	// _ "github.com/pygolin/stdlib/pkg/tokenize"
	// _ "github.com/pygolin/stdlib/pkg/re"
	// _ "github.com/pygolin/stdlib/pkg/inspect"
	// _ "github.com/pygolin/stdlib/pkg/sys"
	// _ "github.com/pygolin/stdlib/pkg/linecache"
	// _ "github.com/pygolin/stdlib/pkg/_dill"
	// _ "github.com/pygolin/stdlib/pkg/types"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßEllipsis := πg.InternStr("Ellipsis")
		ßEllipsisType := πg.InternStr("EllipsisType")
		ßFalse := πg.InternStr("False")
		ßIOError := πg.InternStr("IOError")
		ßImportError := πg.InternStr("ImportError")
		ßInf := πg.InternStr("Inf")
		ßNaN := πg.InternStr("NaN")
		ßNone := πg.InternStr("None")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßPY3 := πg.InternStr("PY3")
		ßSyntaxError := πg.InternStr("SyntaxError")
		ßTokenError := πg.InternStr("TokenError")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß__all__ := πg.InternStr("__all__")
		ß__builtin__ := πg.InternStr("__builtin__")
		ß__class__ := πg.InternStr("__class__")
		ß__code__ := πg.InternStr("__code__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__func__ := πg.InternStr("__func__")
		ß__globals__ := πg.InternStr("__globals__")
		ß__locals__ := πg.InternStr("__locals__")
		ß__main__ := πg.InternStr("__main__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__this_is_a_big_dummy_enclosing_function__ := πg.InternStr("__this_is_a_big_dummy_enclosing_function__")
		ß__this_is_a_big_dummy_function__ := πg.InternStr("__this_is_a_big_dummy_function__")
		ß__this_is_a_big_dummy_object__ := πg.InternStr("__this_is_a_big_dummy_object__")
		ß__this_is_a_stub_variable__ := πg.InternStr("__this_is_a_stub_variable__")
		ß_closuredimport := πg.InternStr("_closuredimport")
		ß_closuredsource := πg.InternStr("_closuredsource")
		ß_enclose := πg.InternStr("_enclose")
		ß_get_name := πg.InternStr("_get_name")
		ß_getimport := πg.InternStr("_getimport")
		ß_hascode := πg.InternStr("_hascode")
		ß_importable := πg.InternStr("_importable")
		ß_intypes := πg.InternStr("_intypes")
		ß_isinstance := πg.InternStr("_isinstance")
		ß_isstring := πg.InternStr("_isstring")
		ß_likely_import := πg.InternStr("_likely_import")
		ß_matchlambda := πg.InternStr("_matchlambda")
		ß_namespace := πg.InternStr("_namespace")
		ß_outdent := πg.InternStr("_outdent")
		ß_wrap := πg.InternStr("_wrap")
		ßall := πg.InternStr("all")
		ßappend := πg.InternStr("append")
		ßargs := πg.InternStr("args")
		ßarray := πg.InternStr("array")
		ßbasestring := πg.InternStr("basestring")
		ßbuiltins := πg.InternStr("builtins")
		ßbytes := πg.InternStr("bytes")
		ßc := πg.InternStr("c")
		ßco_code := πg.InternStr("co_code")
		ßco_filename := πg.InternStr("co_filename")
		ßco_firstlineno := πg.InternStr("co_firstlineno")
		ßcode := πg.InternStr("code")
		ßcompile := πg.InternStr("compile")
		ßcount := πg.InternStr("count")
		ßdict := πg.InternStr("dict")
		ßdumps := πg.InternStr("dumps")
		ßdumpsource := πg.InternStr("dumpsource")
		ßellipsis := πg.InternStr("ellipsis")
		ßeval := πg.InternStr("eval")
		ßexc_info := πg.InternStr("exc_info")
		ßf_code := πg.InternStr("f_code")
		ßfindsource := πg.InternStr("findsource")
		ßfreevars := πg.InternStr("freevars")
		ßfunc_code := πg.InternStr("func_code")
		ßget_current_history_length := πg.InternStr("get_current_history_length")
		ßget_history_item := πg.InternStr("get_history_item")
		ßgetattr := πg.InternStr("getattr")
		ßgetblock := πg.InternStr("getblock")
		ßgetblocks := πg.InternStr("getblocks")
		ßgetblocks_from_history := πg.InternStr("getblocks_from_history")
		ßgetfile := πg.InternStr("getfile")
		ßgetimport := πg.InternStr("getimport")
		ßgetimportable := πg.InternStr("getimportable")
		ßgetlines := πg.InternStr("getlines")
		ßgetmodule := πg.InternStr("getmodule")
		ßgetname := πg.InternStr("getname")
		ßgetsource := πg.InternStr("getsource")
		ßgetsourcefile := πg.InternStr("getsourcefile")
		ßgetsourcelines := πg.InternStr("getsourcelines")
		ßglobals := πg.InternStr("globals")
		ßglobalvars := πg.InternStr("globalvars")
		ßgroup := πg.InternStr("group")
		ßgroups := πg.InternStr("groups")
		ßhasattr := πg.InternStr("hasattr")
		ßim_func := πg.InternStr("im_func")
		ßimportable := πg.InternStr("importable")
		ßindent := πg.InternStr("indent")
		ßindentsize := πg.InternStr("indentsize")
		ßindex := πg.InternStr("index")
		ßinf := πg.InternStr("inf")
		ßint := πg.InternStr("int")
		ßisbuiltin := πg.InternStr("isbuiltin")
		ßisclass := πg.InternStr("isclass")
		ßiscode := πg.InternStr("iscode")
		ßisdynamic := πg.InternStr("isdynamic")
		ßisframe := πg.InternStr("isframe")
		ßisfrommain := πg.InternStr("isfrommain")
		ßisfunction := πg.InternStr("isfunction")
		ßisinstance := πg.InternStr("isinstance")
		ßismethod := πg.InternStr("ismethod")
		ßismodule := πg.InternStr("ismodule")
		ßistraceback := πg.InternStr("istraceback")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßlikely_import := πg.InternStr("likely_import")
		ßlinecache := πg.InternStr("linecache")
		ßlist := πg.InternStr("list")
		ßlocals := πg.InternStr("locals")
		ßlstrip := πg.InternStr("lstrip")
		ßmatch := πg.InternStr("match")
		ßnan := πg.InternStr("nan")
		ßnumpy := πg.InternStr("numpy")
		ßoutdent := πg.InternStr("outdent")
		ßoutermost := πg.InternStr("outermost")
		ßplatform := πg.InternStr("platform")
		ßpop := πg.InternStr("pop")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßrepr := πg.InternStr("repr")
		ßrsplit := πg.InternStr("rsplit")
		ßset := πg.InternStr("set")
		ßsort := πg.InternStr("sort")
		ßsorted := πg.InternStr("sorted")
		ßsplit := πg.InternStr("split")
		ßstartswith := πg.InternStr("startswith")
		ßstr := πg.InternStr("str")
		ßstrip := πg.InternStr("strip")
		ßtb_frame := πg.InternStr("tb_frame")
		ßtype := πg.InternStr("type")
		ßtypes := πg.InternStr("types")
		ßvalues := πg.InternStr("values")
		ßvarnames := πg.InternStr("varnames")
		ßwin := πg.InternStr("win")
		ßwrap2 := πg.InternStr("wrap2")
		ßwrap3 := πg.InternStr("wrap3")
		ßzip := πg.InternStr("zip")
		var πTemp001 []*πg.Object
		_ = πTemp001
		var πTemp002 *πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 []πg.Param
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 *πg.Object
		_ = πTemp017
		var πTemp018 *πg.Object
		_ = πTemp018
		var πTemp019 bool
		_ = πTemp019
		var πTemp020 *πg.Object
		_ = πTemp020
		var πTemp021 *πg.Object
		_ = πTemp021
		var πTemp022 *πg.Object
		_ = πTemp022
		var πTemp023 *πg.Object
		_ = πTemp023
		var πTemp024 *πg.Object
		_ = πTemp024
		var πTemp025 *πg.Object
		_ = πTemp025
		var πTemp026 *πg.Object
		_ = πTemp026
		var πTemp027 *πg.Object
		_ = πTemp027
		var πTemp028 *πg.Object
		_ = πTemp028
		var πTemp029 *πg.Object
		_ = πTemp029
		var πTemp030 *πg.Object
		_ = πTemp030
		var πTemp031 *πg.Object
		_ = πTemp031
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 12: """
			πF.SetLineno(12)
			// line 12: """
			πF.SetLineno(12)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nExtensions to python's 'inspect' module, which can be used\nto retrieve information from live python objects. The methods\ndefined in this module are augmented to facilitate access to \nsource code of interactively defined functions and classes,\nas well as provide access to source code for objects defined\nin a file.\n").ToObject()); πE != nil {
				continue
			}
			// line 21: __all__ = ['findsource', 'getsourcelines', 'getsource', 'indent', 'outdent', \
			πF.SetLineno(21)
			πTemp001 = make([]*πg.Object, 14)
			πTemp001[0] = ßfindsource.ToObject()
			πTemp001[1] = ßgetsourcelines.ToObject()
			πTemp001[2] = ßgetsource.ToObject()
			πTemp001[3] = ßindent.ToObject()
			πTemp001[4] = ßoutdent.ToObject()
			πTemp001[5] = ß_wrap.ToObject()
			πTemp001[6] = ßdumpsource.ToObject()
			πTemp001[7] = ßgetname.ToObject()
			πTemp001[8] = ß_namespace.ToObject()
			πTemp001[9] = ßgetimport.ToObject()
			πTemp001[10] = ß_importable.ToObject()
			πTemp001[11] = ßimportable.ToObject()
			πTemp001[12] = ßisdynamic.ToObject()
			πTemp001[13] = ßisfrommain.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 25: import linecache
			πF.SetLineno(25)
			if πTemp001, πE = πg.ImportModule(πF, "linecache"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßlinecache.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 26: import re
			πF.SetLineno(26)
			if πTemp001, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 27: from inspect import (getblock, getfile, getmodule, getsourcefile, indentsize,
			πF.SetLineno(27)
			if πTemp001, πE = πg.ImportModule(πF, "inspect"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßgetblock); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetblock.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßgetfile); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetfile.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßgetmodule); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetmodule.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßgetsourcefile); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetsourcefile.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßindentsize); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßindentsize.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßisbuiltin); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßisbuiltin.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßisclass); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßisclass.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßiscode); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßiscode.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßisframe); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßisframe.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßisfunction); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßisfunction.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßismethod); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßismethod.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßismodule); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßismodule.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßistraceback); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßistraceback.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 30: from tokenize import TokenError
			πF.SetLineno(30)
			if πTemp001, πE = πg.ImportModule(πF, "tokenize"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßTokenError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTokenError.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 32: from ._dill import PY3
			πF.SetLineno(32)
			if πTemp001, πE = πg.ImportModule(πF, "._dill"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßPY3); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPY3.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 35: def isfrommain(obj):
			πF.SetLineno(35)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("isfrommain", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 36: "check if object was built in __main__"
					πF.SetLineno(36)
					// line 37: module = getmodule(obj)
					πF.SetLineno(37)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmodule = πTemp003
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp002 = µmodule
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µmodule, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp005, ß__main__.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 38: if module and module.__name__ == '__main__':
					πF.SetLineno(38)
				Label2:
					// line 39: return True
					πF.SetLineno(39)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label3
				Label3:
					// line 40: return False
					πF.SetLineno(40)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp002
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßisfrommain.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 36: "check if object was built in __main__"
			πF.SetLineno(36)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("check if object was built in __main__").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßisfrommain); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 43: def isdynamic(obj):
			πF.SetLineno(43)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("isdynamic", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µfile *πg.Object = πg.UnboundLocal; _ = µfile
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.BaseException
				_ = πTemp004
				var πTemp005 *πg.Traceback
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 44: "check if object was built in the interpreter"
					πF.SetLineno(44)
					// line 45: try: file = getfile(obj)
					πF.SetLineno(45)
					πF.PushCheckpoint(2)
					// line 45: try: file = getfile(obj)
					πF.SetLineno(45)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetfile); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfile = πTemp003
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 46: except TypeError: file = None
					πF.SetLineno(46)
				Label3:
					// line 46: except TypeError: file = None
					πF.SetLineno(46)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µfile = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µfile, πg.NewStr("<stdin>").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label4
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisfrommain); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp007
				Label4:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label5
					}
					goto Label6
					// line 47: if file == '<stdin>' and isfrommain(obj):
					πF.SetLineno(47)
				Label5:
					// line 48: return True
					πF.SetLineno(48)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label6
				Label6:
					// line 49: return False
					πF.SetLineno(49)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp002
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßisdynamic.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 44: "check if object was built in the interpreter"
			πF.SetLineno(44)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("check if object was built in the interpreter").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßisdynamic); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
				continue
			}
			// line 52: def _matchlambda(func, line):
			πF.SetLineno(52)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			πTemp004[1] = πg.Param{Name: "line", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("_matchlambda", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µline *πg.Object = πArgs[1]; _ = µline
				var µgetcode *πg.Object = πg.UnboundLocal; _ = µgetcode
				var µfreevars *πg.Object = πg.UnboundLocal; _ = µfreevars
				var µglobalvars *πg.Object = πg.UnboundLocal; _ = µglobalvars
				var µvarnames *πg.Object = πg.UnboundLocal; _ = µvarnames
				var µdummy *πg.Object = πg.UnboundLocal; _ = µdummy
				var µlhs *πg.Object = πg.UnboundLocal; _ = µlhs
				var µrhs *πg.Object = πg.UnboundLocal; _ = µrhs
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
				var µ_f *πg.Object = πg.UnboundLocal; _ = µ_f
				var µ_lhs *πg.Object = πg.UnboundLocal; _ = µ_lhs
				var µ_rhs *πg.Object = πg.UnboundLocal; _ = µ_rhs
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 bool
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πTemp017 *πg.Object
				_ = πTemp017
				var πTemp018 *πg.Object
				_ = πTemp018
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 20: goto Label20
					default: panic("unexpected function state")
					}
					// line 53: """check if lambda object 'func' matches raw line of code 'line'"""
					πF.SetLineno(53)
					// line 54: from .detect import code as getcode
					πF.SetLineno(54)
					if πTemp002, πE = πg.ImportModule(πF, ".detect"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßcode); πE != nil {
						continue
					}
					µgetcode = πTemp003
					// line 55: from .detect import freevars, globalvars, varnames
					πF.SetLineno(55)
					if πTemp002, πE = πg.ImportModule(πF, ".detect"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßfreevars); πE != nil {
						continue
					}
					µfreevars = πTemp003
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßglobalvars); πE != nil {
						continue
					}
					µglobalvars = πTemp003
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßvarnames); πE != nil {
						continue
					}
					µvarnames = πTemp003
					// line 56: dummy = lambda : '__this_is_a_big_dummy_function__'
					πF.SetLineno(56)
					πTemp004 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 56: dummy = lambda : '__this_is_a_big_dummy_function__'
							πF.SetLineno(56)
							πR = ß__this_is_a_big_dummy_function__.ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					µdummy = πTemp001
					// line 58: lhs,rhs = line.split('lambda ',1)[-1].split(":", 1) #FIXME: if !1 inputs
					πF.SetLineno(58)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr(":").ToObject()
					πTemp002[1] = πg.NewInt(1).ToObject()
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					πTemp005 = πF.MakeArgs(2)
					πTemp005[0] = πg.NewStr("lambda ").ToObject()
					πTemp005[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µline, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
						continue
					}
					µlhs = πTemp001
					µrhs = πTemp006
					// line 59: try: #FIXME: unsafe
					πF.SetLineno(59)
					πF.PushCheckpoint(2)
					// line 60: _ = eval("lambda %s : %s" % (lhs,rhs), globals(),locals())
					πF.SetLineno(60)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µlhs, "lhs"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrhs, "rhs"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µlhs, µrhs).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("lambda %s : %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlocals); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[2] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßeval); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µ_ = πTemp003
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					goto Label3
					// line 61: except: _ = dummy
					πF.SetLineno(61)
				Label3:
					// line 61: except: _ = dummy
					πF.SetLineno(61)
					if πE = πg.CheckLocal(πF, µdummy, "dummy"); πE != nil {
						continue
					}
					µ_ = µdummy
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 63: _, code = getcode(_).co_code, getcode(func).co_code
					πF.SetLineno(63)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					πTemp002[0] = µ_
					if πE = πg.CheckLocal(πF, µgetcode, "getcode"); πE != nil {
						continue
					}
					if πTemp003, πE = µgetcode.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßco_code, nil); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πE = πg.CheckLocal(πF, µgetcode, "getcode"); πE != nil {
						continue
					}
					if πTemp003, πE = µgetcode.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßco_code, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
						continue
					}
					µ_ = πTemp003
					µcode = πTemp006
					// line 65: _f = [line.count(i) for i in freevars(func).keys()]
					πF.SetLineno(65)
					πTemp004 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
									continue
								}
								πTemp002[0] = µfunc
								if πE = πg.CheckLocal(πF, µfreevars, "freevars"); πE != nil {
									continue
								}
								if πTemp003, πE = µfreevars.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
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
									µi = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 65: _f = [line.count(i) for i in freevars(func).keys()]
								πF.SetLineno(65)
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp002[0] = µi
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µline, ßcount, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πF.PushCheckpoint(4)
								return πTemp004, nil
							Label4:
								πTemp003 = πSent
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
					if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
						continue
					}
					µ_f = πTemp001
					if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, µ_f); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp010).ToObject()
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label4
					}
					goto Label5
					// line 66: if not _f: # not in closure
					πF.SetLineno(66)
				Label4:
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µ_, µcode); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label6
					}
					goto Label7
					// line 68: if _ == code: return True
					πF.SetLineno(68)
				Label6:
					// line 68: if _ == code: return True
					πF.SetLineno(68)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label7
				Label7:
					// line 69: return False
					πF.SetLineno(69)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label5
				Label5:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
						continue
					}
					πTemp002[0] = µ_f
					if πTemp006, πE = πg.ResolveGlobal(πF, ßall); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp010, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp010).ToObject()
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label8
					}
					goto Label9
					// line 71: if not all(_f): return False  #XXX: VERY WEAK
					πF.SetLineno(71)
				Label8:
					// line 71: if not all(_f): return False  #XXX: VERY WEAK
					πF.SetLineno(71)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label9
				Label9:
					// line 73: _f = varnames(func)
					πF.SetLineno(73)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πE = πg.CheckLocal(πF, µvarnames, "varnames"); πE != nil {
						continue
					}
					if πTemp001, πE = µvarnames.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µ_f = πTemp001
					// line 74: _f = [line.count(i) for i in _f[0]+_f[1]]
					πF.SetLineno(74)
					πTemp004 = make([]πg.Param, 0)
					πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
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
								πTemp003 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetItem(πF, µ_f, πTemp003); πE != nil {
									continue
								}
								πTemp003 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, µ_f, πTemp003); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp006 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp006 {
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
									πTemp007 = !isStop
								} else {
									πTemp007 = true
									µi = πTemp002
								}
								if πE != nil || !πTemp007 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 74: _f = [line.count(i) for i in _f[0]+_f[1]]
								πF.SetLineno(74)
								πTemp008 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp008[0] = µi
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µline, ßcount, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp008)
								πF.PushCheckpoint(4)
								return πTemp003, nil
							Label4:
								πTemp002 = πSent
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
					µ_f = πTemp001
					if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
						continue
					}
					πTemp001 = µ_f
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp010 {
						goto Label10
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
						continue
					}
					πTemp002[0] = µ_f
					if πTemp011, πE = πg.ResolveGlobal(πF, ßall); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp013, πE = πg.IsTrue(πF, πTemp012); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(!πTemp013).ToObject()
					πTemp001 = πTemp007
				Label10:
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label11
					}
					goto Label12
					// line 75: if _f and not all(_f): return False  #XXX: VERY WEAK
					πF.SetLineno(75)
				Label11:
					// line 75: if _f and not all(_f): return False  #XXX: VERY WEAK
					πF.SetLineno(75)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label12
				Label12:
					// line 76: _f = [line.count(i) for i in globalvars(func).keys()]
					πF.SetLineno(76)
					πTemp004 = make([]πg.Param, 0)
					πTemp007 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
									continue
								}
								πTemp002[0] = µfunc
								if πE = πg.CheckLocal(πF, µglobalvars, "globalvars"); πE != nil {
									continue
								}
								if πTemp003, πE = µglobalvars.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
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
									µi = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 76: _f = [line.count(i) for i in globalvars(func).keys()]
								πF.SetLineno(76)
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp002[0] = µi
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µline, ßcount, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πF.PushCheckpoint(4)
								return πTemp004, nil
							Label4:
								πTemp003 = πSent
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
					if πTemp011, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp011}, nil); πE != nil {
						continue
					}
					µ_f = πTemp001
					if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
						continue
					}
					πTemp001 = µ_f
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp010 {
						goto Label13
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
						continue
					}
					πTemp002[0] = µ_f
					if πTemp012, πE = πg.ResolveGlobal(πF, ßall); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp013, πE = πg.IsTrue(πF, πTemp014); πE != nil {
						continue
					}
					πTemp011 = πg.GetBool(!πTemp013).ToObject()
					πTemp001 = πTemp011
				Label13:
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label14
					}
					goto Label15
					// line 77: if _f and not all(_f): return False  #XXX: VERY WEAK
					πF.SetLineno(77)
				Label14:
					// line 77: if _f and not all(_f): return False  #XXX: VERY WEAK
					πF.SetLineno(77)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label15
				Label15:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("lambda ").ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, µline, ßcount, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp011, πE = πg.GT(πF, πTemp014, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp011
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp010 {
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µlhs, "lhs"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πE = πg.CheckLocal(πF, µfreevars, "freevars"); πE != nil {
						continue
					}
					if πTemp012, πE = µfreevars.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp014, πE = πg.GetAttr(πF, πTemp012, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp014.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πg.Contains(πF, πTemp012, µlhs); πE != nil {
						continue
					}
					πTemp011 = πg.GetBool(πTemp013).ToObject()
					πTemp001 = πTemp011
				Label16:
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label17
					}
					goto Label18
					// line 79: if (line.count('lambda ') > 1) and (lhs in freevars(func).keys()):
					πF.SetLineno(79)
				Label17:
					// line 80: _lhs,_rhs = rhs.split('lambda ',1)[-1].split(":",1) #FIXME: if !1 inputs
					πF.SetLineno(80)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr(":").ToObject()
					πTemp002[1] = πg.NewInt(1).ToObject()
					if πTemp011, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp011
					πTemp005 = πF.MakeArgs(2)
					πTemp005[0] = πg.NewStr("lambda ").ToObject()
					πTemp005[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µrhs, "rhs"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, µrhs, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp012.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp011, πE = πg.GetItem(πF, πTemp014, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp011, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp012}}}, πTemp011); πE != nil {
						continue
					}
					µ_lhs = πTemp001
					µ_rhs = πTemp012
					// line 81: try: #FIXME: unsafe
					πF.SetLineno(81)
					πF.PushCheckpoint(20)
					// line 82: _f = eval("lambda %s : %s" % (_lhs,_rhs), globals(),locals())
					πF.SetLineno(82)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µ_lhs, "_lhs"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_rhs, "_rhs"); πE != nil {
						continue
					}
					πTemp011 = πg.NewTuple2(µ_lhs, µ_rhs).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("lambda %s : %s").ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp011
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlocals); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[2] = πTemp011
					if πTemp001, πE = πg.ResolveGlobal(πF, ßeval); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µ_f = πTemp011
					πF.PopCheckpoint()
					goto Label19
				Label20:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					goto Label21
					// line 83: except: _f = dummy
					πF.SetLineno(83)
				Label21:
					// line 83: except: _f = dummy
					πF.SetLineno(83)
					if πE = πg.CheckLocal(πF, µdummy, "dummy"); πE != nil {
						continue
					}
					µ_f = µdummy
					πF.RestoreExc(nil, nil)
					goto Label19
				Label19:
					// line 85: _, code = getcode(_f).co_code, getcode(func).co_code
					πF.SetLineno(85)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
						continue
					}
					πTemp002[0] = µ_f
					if πE = πg.CheckLocal(πF, µgetcode, "getcode"); πE != nil {
						continue
					}
					if πTemp011, πE = µgetcode.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßco_code, nil); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πE = πg.CheckLocal(πF, µgetcode, "getcode"); πE != nil {
						continue
					}
					if πTemp011, πE = µgetcode.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp014, πE = πg.GetAttr(πF, πTemp011, ßco_code, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp012, πTemp014).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp011}, πg.TieTarget{Target: &πTemp012}}}, πTemp001); πE != nil {
						continue
					}
					µ_ = πTemp011
					µcode = πTemp012
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					πTemp002[0] = µ_
					if πTemp011, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp002[0] = µcode
					if πTemp011, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.NE(πF, πTemp012, πTemp014); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label22
					}
					goto Label23
					// line 86: if len(_) != len(code): return False
					πF.SetLineno(86)
				Label22:
					// line 86: if len(_) != len(code): return False
					πF.SetLineno(86)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label23
				Label23:
					// line 88: _ = set((i,j) for (i,j) in zip(_,code) if i != j)
					πF.SetLineno(88)
					πTemp002 = πF.MakeArgs(1)
					πTemp004 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µj *πg.Object = πg.UnboundLocal; _ = µj
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
								πTemp002 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
									continue
								}
								πTemp002[0] = µ_
								if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
									continue
								}
								πTemp002[1] = µcode
								if πTemp003, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
										continue
									}
									µi = πTemp004
									µj = πTemp007
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.NE(πF, µi, µj); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 88: _ = set((i,j) for (i,j) in zip(_,code) if i != j)
								πF.SetLineno(88)
							Label4:
								// line 88: _ = set((i,j) for (i,j) in zip(_,code) if i != j)
								πF.SetLineno(88)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
									continue
								}
								πTemp003 = πg.NewTuple2(µi, µj).ToObject()
								πF.PushCheckpoint(6)
								return πTemp003, nil
							Label6:
								πTemp004 = πSent
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
					if πTemp011, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp011
					if πTemp011, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µ_ = πTemp012
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					πTemp002[0] = µ_
					if πTemp012, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp011, πE = πg.NE(πF, πTemp014, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, πTemp011); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label24
					}
					goto Label25
					// line 89: if len(_) != 1: return False #('t','\x88')
					πF.SetLineno(89)
				Label24:
					// line 89: if len(_) != 1: return False #('t','\x88')
					πF.SetLineno(89)
					if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp011
					continue
					goto Label25
				Label25:
					// line 90: return True
					πF.SetLineno(90)
					if πTemp011, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πR = πTemp011
					continue
					goto Label18
				Label18:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp002[0] = µline
					if πTemp012, πE = πg.ResolveGlobal(πF, ßindentsize); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp010, πE = πg.IsTrue(πF, πTemp014); πE != nil {
						continue
					}
					πTemp011 = πg.GetBool(!πTemp010).ToObject()
					if πTemp010, πE = πg.IsTrue(πF, πTemp011); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label26
					}
					goto Label27
					// line 92: if not indentsize(line): return False #FIXME: is this a good check???
					πF.SetLineno(92)
				Label26:
					// line 92: if not indentsize(line): return False #FIXME: is this a good check???
					πF.SetLineno(92)
					if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp011
					continue
					goto Label27
				Label27:
					// line 95: _ = _.split(_[0])  # 't' #XXX: remove matching values if starts the same?
					πF.SetLineno(95)
					πTemp002 = πF.MakeArgs(1)
					πTemp011 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µ_, πTemp011); πE != nil {
						continue
					}
					πTemp002[0] = πTemp012
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, µ_, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µ_ = πTemp012
					// line 96: _f = code.split(code[0])  # '\x88'
					πF.SetLineno(96)
					πTemp002 = πF.MakeArgs(1)
					πTemp011 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µcode, πTemp011); πE != nil {
						continue
					}
					πTemp002[0] = πTemp012
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, µcode, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µ_f = πTemp012
					// line 98: _ = dict(re.match(r'([\W\D\S])(.*)', _[i]).groups() for i in range(1,len(_)))
					πF.SetLineno(98)
					πTemp002 = πF.MakeArgs(1)
					πTemp004 = make([]πg.Param, 0)
					πTemp011 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
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
								πTemp002 = πF.MakeArgs(2)
								πTemp002[0] = πg.NewInt(1).ToObject()
								πTemp003 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
									continue
								}
								πTemp003[0] = µ_
								if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								πTemp002[1] = πTemp005
								if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp006 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp006 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp007 = !isStop
								} else {
									πTemp007 = true
									µi = πTemp004
								}
								if πE != nil || !πTemp007 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 98: _ = dict(re.match(r'([\W\D\S])(.*)', _[i]).groups() for i in range(1,len(_)))
								πF.SetLineno(98)
								πTemp002 = πF.MakeArgs(2)
								πTemp002[0] = πg.NewStr("([\\W\\D\\S])(.*)").ToObject()
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp004 = µi
								if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, µ_, πTemp004); πE != nil {
									continue
								}
								πTemp002[1] = πTemp005
								if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmatch, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßgroups, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp004, nil
							Label4:
								πTemp005 = πSent
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
					if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp012
					if πTemp012, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µ_ = πTemp014
					// line 99: _f = dict(re.match(r'([\W\D\S])(.*)', _f[i]).groups() for i in range(1,len(_f)))
					πF.SetLineno(99)
					πTemp002 = πF.MakeArgs(1)
					πTemp004 = make([]πg.Param, 0)
					πTemp012 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
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
								πTemp002 = πF.MakeArgs(2)
								πTemp002[0] = πg.NewInt(1).ToObject()
								πTemp003 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
									continue
								}
								πTemp003[0] = µ_f
								if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								πTemp002[1] = πTemp005
								if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp006 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp006 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp007 = !isStop
								} else {
									πTemp007 = true
									µi = πTemp004
								}
								if πE != nil || !πTemp007 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 99: _f = dict(re.match(r'([\W\D\S])(.*)', _f[i]).groups() for i in range(1,len(_f)))
								πF.SetLineno(99)
								πTemp002 = πF.MakeArgs(2)
								πTemp002[0] = πg.NewStr("([\\W\\D\\S])(.*)").ToObject()
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp004 = µi
								if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, µ_f, πTemp004); πE != nil {
									continue
								}
								πTemp002[1] = πTemp005
								if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmatch, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßgroups, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp004, nil
							Label4:
								πTemp005 = πSent
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
					if πTemp014, πE = πTemp012.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp014
					if πTemp014, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µ_f = πTemp015
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, µ_, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp017, πE = πTemp016.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, µ_f, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp018, πE = πTemp016.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πg.Eq(πF, πTemp017, πTemp018); πE != nil {
						continue
					}
					πTemp014 = πTemp015
					if πTemp010, πE = πg.IsTrue(πF, πTemp014); πE != nil {
						continue
					}
					if !πTemp010 {
						goto Label28
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, µ_, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp017, πE = πTemp016.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp017
					if πTemp016, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
						continue
					}
					if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_f, "_f"); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, µ_f, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp018, πE = πTemp016.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp018
					if πTemp016, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
						continue
					}
					if πTemp018, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp015, πE = πg.Eq(πF, πTemp017, πTemp018); πE != nil {
						continue
					}
					πTemp014 = πTemp015
				Label28:
					if πTemp010, πE = πg.IsTrue(πF, πTemp014); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label29
					}
					goto Label30
					// line 100: if (_.keys() == _f.keys()) and (sorted(_.values()) == sorted(_f.values())):
					πF.SetLineno(100)
				Label29:
					// line 101: return True
					πF.SetLineno(101)
					if πTemp014, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πR = πTemp014
					continue
					goto Label30
				Label30:
					// line 102: return False
					πF.SetLineno(102)
					if πTemp014, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp014
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_matchlambda.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 53: """check if lambda object 'func' matches raw line of code 'line'"""
			πF.SetLineno(53)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("check if lambda object 'func' matches raw line of code 'line'").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_matchlambda); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
			// line 105: def findsource(object):
			πF.SetLineno(105)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("findsource", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var µfile *πg.Object = πg.UnboundLocal; _ = µfile
				var µreadline *πg.Object = πg.UnboundLocal; _ = µreadline
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
				var µsys *πg.Object = πg.UnboundLocal; _ = µsys
				var µlbuf *πg.Object = πg.UnboundLocal; _ = µlbuf
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µsourcefile *πg.Object = πg.UnboundLocal; _ = µsourcefile
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µpat1 *πg.Object = πg.UnboundLocal; _ = µpat1
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
				var µpat2 *πg.Object = πg.UnboundLocal; _ = µpat2
				var µstdin *πg.Object = πg.UnboundLocal; _ = µstdin
				var µlnum *πg.Object = πg.UnboundLocal; _ = µlnum
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var µhats *πg.Object = πg.UnboundLocal; _ = µhats
				var µ_lnum *πg.Object = πg.UnboundLocal; _ = µ_lnum
				var µpat *πg.Object = πg.UnboundLocal; _ = µpat
				var µcandidates *πg.Object = πg.UnboundLocal; _ = µcandidates
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µmatch *πg.Object = πg.UnboundLocal; _ = µmatch
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.BaseException
				_ = πTemp004
				var πTemp005 *πg.Traceback
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 []πg.Param
				_ = πTemp012
				var πTemp013 []*πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 bool
				_ = πTemp015
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 64: goto Label64
					case 2: goto Label2
					case 9: goto Label9
					case 75: goto Label75
					case 76: goto Label76
					case 16: goto Label16
					case 82: goto Label82
					case 89: goto Label89
					case 90: goto Label90
					case 63: goto Label63
					default: panic("unexpected function state")
					}
					// line 106: """Return the entire source file and starting line number for an object.
					πF.SetLineno(106)
					// line 115: module = getmodule(object)
					πF.SetLineno(115)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmodule = πTemp003
					// line 116: try: file = getfile(module)
					πF.SetLineno(116)
					πF.PushCheckpoint(2)
					// line 116: try: file = getfile(module)
					πF.SetLineno(116)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001[0] = µmodule
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetfile); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfile = πTemp003
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 117: except TypeError: file = None
					πF.SetLineno(117)
				Label3:
					// line 117: except TypeError: file = None
					πF.SetLineno(117)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µfile = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp002 = µmodule
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µmodule, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp007, ß__main__.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µfile); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp008).ToObject()
					πTemp002 = πTemp003
				Label4:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label5
					}
					goto Label6
					// line 119: if module and module.__name__ == '__main__' and not file:
					πF.SetLineno(119)
				Label5:
					// line 120: try:
					πF.SetLineno(120)
					πF.PushCheckpoint(9)
					// line 121: import readline
					πF.SetLineno(121)
					if πTemp001, πE = πg.ImportModule(πF, "readline"); πE != nil {
						continue
					}
					πTemp002 = πTemp001[0]
					µreadline = πTemp002
					// line 122: err = ''
					πF.SetLineno(122)
					µerr = ß.ToObject()
					πF.PopCheckpoint()
					goto Label8
				Label9:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					goto Label10
					// line 123: except:
					πF.SetLineno(123)
				Label10:
					// line 124: import sys
					πF.SetLineno(124)
					if πTemp001, πE = πg.ImportModule(πF, "sys"); πE != nil {
						continue
					}
					πTemp002 = πTemp001[0]
					µsys = πTemp002
					// line 125: err = sys.exc_info()[1].args[0]
					πF.SetLineno(125)
					πTemp002 = πg.NewInt(0).ToObject()
					πTemp007 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, µsys, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, πTemp011, πTemp007); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp009, ßargs, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
						continue
					}
					µerr = πTemp003
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(3).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µsys, ßplatform, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp007, ßwin.ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label11
					}
					goto Label12
					// line 126: if sys.platform[:3] == 'win':
					πF.SetLineno(126)
				Label11:
					// line 127: err += ", please install 'pyreadline'"
					πF.SetLineno(127)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µerr, πg.NewStr(", please install 'pyreadline'").ToObject()); πE != nil {
						continue
					}
					µerr = πTemp002
					goto Label12
				Label12:
					πF.RestoreExc(nil, nil)
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label13
					}
					goto Label14
					// line 128: if err:
					πF.SetLineno(128)
				Label13:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					πTemp001[0] = µerr
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 129: raise IOError(err)
					πF.SetLineno(129)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label14
				Label14:
					// line 130: lbuf = readline.get_current_history_length()
					πF.SetLineno(130)
					if πE = πg.CheckLocal(πF, µreadline, "readline"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µreadline, ßget_current_history_length, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µlbuf = πTemp003
					// line 131: lines = [readline.get_history_item(i)+'\n' for i in range(1,lbuf)]
					πF.SetLineno(131)
					πTemp012 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp012, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								πTemp002 = πF.MakeArgs(2)
								πTemp002[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µlbuf, "lbuf"); πE != nil {
									continue
								}
								πTemp002[1] = µlbuf
								if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
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
									µi = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 131: lines = [readline.get_history_item(i)+'\n' for i in range(1,lbuf)]
								πF.SetLineno(131)
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp002[0] = µi
								if πE = πg.CheckLocal(πF, µreadline, "readline"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µreadline, ßget_history_item, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp003, πE = πg.Add(πF, πTemp007, πg.NewStr("\n").ToObject()); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp003, nil
							Label4:
								πTemp004 = πSent
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
					if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
						continue
					}
					µlines = πTemp002
					goto Label7
				Label6:
					// line 133: try: # special handling for class instances
					πF.SetLineno(133)
					πF.PushCheckpoint(16)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp009, πE = πg.ResolveGlobal(πF, ßisclass); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp010); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(!πTemp008).ToObject()
					πTemp002 = πTemp007
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label17
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp013 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp013[0] = µobject
					if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					πTemp001[0] = πTemp009
					if πTemp007, πE = πg.ResolveGlobal(πF, ßisclass); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp009
				Label17:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label18
					}
					goto Label19
					// line 134: if not isclass(object) and isclass(type(object)): # __class__
					πF.SetLineno(134)
				Label18:
					// line 135: file = getfile(module)
					πF.SetLineno(135)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001[0] = µmodule
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetfile); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfile = πTemp007
					// line 136: sourcefile = getsourcefile(module)
					πF.SetLineno(136)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001[0] = µmodule
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetsourcefile); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsourcefile = πTemp007
					goto Label20
				Label19:
					// line 138: file = getfile(object)
					πF.SetLineno(138)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetfile); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfile = πTemp007
					// line 139: sourcefile = getsourcefile(object)
					πF.SetLineno(139)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetsourcefile); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsourcefile = πTemp007
					goto Label20
				Label20:
					πF.PopCheckpoint()
					goto Label15
				Label16:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp007, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp007, πTemp009).ToObject()
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label21
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 140: except (TypeError, AttributeError): # fail with better error
					πF.SetLineno(140)
				Label21:
					// line 141: file = getfile(object)
					πF.SetLineno(141)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetfile); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfile = πTemp007
					// line 142: sourcefile = getsourcefile(object)
					πF.SetLineno(142)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetsourcefile); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsourcefile = πTemp007
					πF.RestoreExc(nil, nil)
					goto Label15
				Label15:
					if πE = πg.CheckLocal(πF, µsourcefile, "sourcefile"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µsourcefile); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(!πTemp008).ToObject()
					πTemp002 = πTemp007
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label22
					}
					if πTemp010, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µfile, πTemp010); πE != nil {
						continue
					}
					if πTemp014, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.SliceType.Call(πF, πg.Args{πTemp014, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetItem(πF, µfile, πTemp010); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Add(πF, πTemp011, πTemp014); πE != nil {
						continue
					}
					if πTemp007, πE = πg.NE(πF, πTemp009, πg.NewStr("<>").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp007
				Label22:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label23
					}
					goto Label24
					// line 143: if not sourcefile and file[:1] + file[-1:] != '<>':
					πF.SetLineno(143)
				Label23:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("source code not available").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 144: raise IOError('source code not available')
					πF.SetLineno(144)
					πE = πF.Raise(πTemp007, nil, nil)
					continue
					goto Label24
				Label24:
					// line 145: file = sourcefile if sourcefile else file
					πF.SetLineno(145)
					if πE = πg.CheckLocal(πF, µsourcefile, "sourcefile"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µsourcefile); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label25
					}
					if πE = πg.CheckLocal(πF, µsourcefile, "sourcefile"); πE != nil {
						continue
					}
					πTemp002 = µsourcefile
					goto Label26
				Label25:
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp002 = µfile
				Label26:
					µfile = πTemp002
					// line 147: module = getmodule(object, file)
					πF.SetLineno(147)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp001[1] = µfile
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmodule = πTemp007
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µmodule); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label27
					}
					goto Label28
					// line 148: if module:
					πF.SetLineno(148)
				Label27:
					// line 149: lines = linecache.getlines(file, module.__dict__)
					πF.SetLineno(149)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp001[0] = µfile
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmodule, ß__dict__, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlinecache); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp002, ßgetlines, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlines = πTemp002
					goto Label29
				Label28:
					// line 151: lines = linecache.getlines(file)
					πF.SetLineno(151)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp001[0] = µfile
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlinecache); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp002, ßgetlines, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlines = πTemp002
					goto Label29
				Label29:
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µlines); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label30
					}
					goto Label31
					// line 153: if not lines:
					πF.SetLineno(153)
				Label30:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("could not extract source code").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 154: raise IOError('could not extract source code')
					πF.SetLineno(154)
					πE = πF.Raise(πTemp007, nil, nil)
					continue
					goto Label31
				Label31:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßismodule); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label32
					}
					goto Label33
					// line 157: if ismodule(object):
					πF.SetLineno(157)
				Label32:
					// line 158: return lines, 0
					πF.SetLineno(158)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µlines, πg.NewInt(0).ToObject()).ToObject()
					πR = πTemp002
					continue
					goto Label33
				Label33:
					// line 161: name = pat1 = obj = ''
					πF.SetLineno(161)
					µname = ß.ToObject()
					µpat1 = ß.ToObject()
					µobj = ß.ToObject()
					// line 162: pat2 = r'^(\s*@)'
					πF.SetLineno(162)
					µpat2 = πg.NewStr("^(\\s*@)").ToObject()
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßismethod); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label34
					}
					goto Label35
					// line 164: if ismethod(object):
					πF.SetLineno(164)
				Label34:
					// line 165: name = object.__name__
					πF.SetLineno(165)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ß__name__, nil); πE != nil {
						continue
					}
					µname = πTemp002
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µname, πg.NewStr("<lambda>").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label36
					}
					goto Label37
					// line 166: if name == '<lambda>': pat1 = r'(.*(?<!\w)lambda(:|\s))'
					πF.SetLineno(166)
				Label36:
					// line 166: if name == '<lambda>': pat1 = r'(.*(?<!\w)lambda(:|\s))'
					πF.SetLineno(166)
					µpat1 = πg.NewStr("(.*(?<!\\w)lambda(:|\\s))").ToObject()
					goto Label38
				Label37:
					// line 167: else: pat1 = r'^(\s*def\s)'
					πF.SetLineno(167)
					µpat1 = πg.NewStr("^(\\s*def\\s)").ToObject()
					goto Label38
				Label38:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label39
					}
					goto Label40
					// line 168: if PY3: object = object.__func__
					πF.SetLineno(168)
				Label39:
					// line 168: if PY3: object = object.__func__
					πF.SetLineno(168)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ß__func__, nil); πE != nil {
						continue
					}
					µobject = πTemp002
					goto Label41
				Label40:
					// line 169: else: object = object.im_func
					πF.SetLineno(169)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ßim_func, nil); πE != nil {
						continue
					}
					µobject = πTemp002
					goto Label41
				Label41:
					goto Label35
				Label35:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisfunction); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label42
					}
					goto Label43
					// line 170: if isfunction(object):
					πF.SetLineno(170)
				Label42:
					// line 171: name = object.__name__
					πF.SetLineno(171)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ß__name__, nil); πE != nil {
						continue
					}
					µname = πTemp002
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µname, πg.NewStr("<lambda>").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label44
					}
					goto Label45
					// line 172: if name == '<lambda>':
					πF.SetLineno(172)
				Label44:
					// line 173: pat1 = r'(.*(?<!\w)lambda(:|\s))'
					πF.SetLineno(173)
					µpat1 = πg.NewStr("(.*(?<!\\w)lambda(:|\\s))").ToObject()
					// line 174: obj = object #XXX: better a copy?
					πF.SetLineno(174)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					µobj = µobject
					goto Label46
				Label45:
					// line 175: else: pat1 = r'^(\s*def\s)'
					πF.SetLineno(175)
					µpat1 = πg.NewStr("^(\\s*def\\s)").ToObject()
					goto Label46
				Label46:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label47
					}
					goto Label48
					// line 176: if PY3: object = object.__code__
					πF.SetLineno(176)
				Label47:
					// line 176: if PY3: object = object.__code__
					πF.SetLineno(176)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ß__code__, nil); πE != nil {
						continue
					}
					µobject = πTemp002
					goto Label49
				Label48:
					// line 177: else: object = object.func_code
					πF.SetLineno(177)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ßfunc_code, nil); πE != nil {
						continue
					}
					µobject = πTemp002
					goto Label49
				Label49:
					goto Label43
				Label43:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßistraceback); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label50
					}
					goto Label51
					// line 178: if istraceback(object):
					πF.SetLineno(178)
				Label50:
					// line 179: object = object.tb_frame
					πF.SetLineno(179)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ßtb_frame, nil); πE != nil {
						continue
					}
					µobject = πTemp002
					goto Label51
				Label51:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisframe); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label52
					}
					goto Label53
					// line 180: if isframe(object):
					πF.SetLineno(180)
				Label52:
					// line 181: object = object.f_code
					πF.SetLineno(181)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ßf_code, nil); πE != nil {
						continue
					}
					µobject = πTemp002
					goto Label53
				Label53:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßiscode); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label54
					}
					goto Label55
					// line 182: if iscode(object):
					πF.SetLineno(182)
				Label54:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					πTemp001[1] = ßco_firstlineno.ToObject()
					if πTemp007, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.IsTrue(πF, πTemp009); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label56
					}
					goto Label57
					// line 183: if not hasattr(object, 'co_firstlineno'):
					πF.SetLineno(183)
				Label56:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("could not find function definition").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 184: raise IOError('could not find function definition')
					πF.SetLineno(184)
					πE = πF.Raise(πTemp007, nil, nil)
					continue
					goto Label57
				Label57:
					// line 185: stdin = object.co_filename == '<stdin>'
					πF.SetLineno(185)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobject, ßco_filename, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp007, πg.NewStr("<stdin>").ToObject()); πE != nil {
						continue
					}
					µstdin = πTemp002
					if πE = πg.CheckLocal(πF, µstdin, "stdin"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µstdin); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label58
					}
					goto Label59
					// line 186: if stdin:
					πF.SetLineno(186)
				Label58:
					// line 187: lnum = len(lines) - 1 # can't get lnum easily, so leverage pat
					πF.SetLineno(187)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp001[0] = µlines
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Sub(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µlnum = πTemp002
					if πE = πg.CheckLocal(πF, µpat1, "pat1"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µpat1); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label61
					}
					goto Label62
					// line 188: if not pat1: pat1 = r'^(\s*def\s)|(.*(?<!\w)lambda(:|\s))|^(\s*@)'
					πF.SetLineno(188)
				Label61:
					// line 188: if not pat1: pat1 = r'^(\s*def\s)|(.*(?<!\w)lambda(:|\s))|^(\s*@)'
					πF.SetLineno(188)
					µpat1 = πg.NewStr("^(\\s*def\\s)|(.*(?<!\\w)lambda(:|\\s))|^(\\s*@)").ToObject()
					goto Label62
				Label62:
					goto Label60
				Label59:
					// line 190: lnum = object.co_firstlineno - 1
					πF.SetLineno(190)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobject, ßco_firstlineno, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µlnum = πTemp002
					// line 191: pat1 = r'^(\s*def\s)|(.*(?<!\w)lambda(:|\s))|^(\s*@)'
					πF.SetLineno(191)
					µpat1 = πg.NewStr("^(\\s*def\\s)|(.*(?<!\\w)lambda(:|\\s))|^(\\s*@)").ToObject()
					goto Label60
				Label60:
					// line 192: pat1 = re.compile(pat1); pat2 = re.compile(pat2)
					πF.SetLineno(192)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpat1, "pat1"); πE != nil {
						continue
					}
					πTemp001[0] = µpat1
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpat1 = πTemp002
					// line 192: pat1 = re.compile(pat1); pat2 = re.compile(pat2)
					πF.SetLineno(192)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpat2, "pat2"); πE != nil {
						continue
					}
					πTemp001[0] = µpat2
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpat2 = πTemp002
					// line 194: while lnum > 0: #XXX: won't find decorators in <stdin> ?
					πF.SetLineno(194)
					πF.PushCheckpoint(64)
					πTemp006 = false
				Label63:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label65
					}
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µlnum, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(63)            
					// line 195: line = lines[lnum]
					πF.SetLineno(195)
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					πTemp002 = µlnum
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					µline = πTemp007
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp001[0] = µline
					if πE = πg.CheckLocal(πF, µpat1, "pat1"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpat1, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label66
					}
					goto Label67
					// line 196: if pat1.match(line):
					πF.SetLineno(196)
				Label66:
					if πE = πg.CheckLocal(πF, µstdin, "stdin"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µstdin); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label68
					}
					goto Label69
					// line 197: if not stdin: break # co_firstlineno does the job
					πF.SetLineno(197)
				Label68:
					// line 197: if not stdin: break # co_firstlineno does the job
					πF.SetLineno(197)
					πTemp006 = true
					continue
					goto Label69
				Label69:
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µname, πg.NewStr("<lambda>").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label70
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µline, µname); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label71
					}
					goto Label72
					// line 198: if name == '<lambda>': # hackery needed to confirm a match
					πF.SetLineno(198)
				Label70:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp001[1] = µline
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_matchlambda); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label73
					}
					goto Label74
					// line 199: if _matchlambda(obj, line): break
					πF.SetLineno(199)
				Label73:
					// line 199: if _matchlambda(obj, line): break
					πF.SetLineno(199)
					πTemp006 = true
					continue
					goto Label74
				Label74:
					goto Label72
					// line 201: if name in line: # need to check for decorator...
					πF.SetLineno(201)
				Label71:
					// line 202: hats = 0
					πF.SetLineno(202)
					µhats = πg.NewInt(0).ToObject()
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Sub(πF, µlnum, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp007
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[1] = πTemp007
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[2] = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp009); πE != nil {
						continue
					}
					πF.PushCheckpoint(76)
					πTemp008 = false
				Label75:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label77
					}
					if πTemp007, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp015 = !isStop
					} else {
						πTemp015 = true
						µ_lnum = πTemp007
					}
					if πE != nil || !πTemp015 {
						continue
					}
					πF.PushCheckpoint(75)            
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_lnum, "_lnum"); πE != nil {
						continue
					}
					πTemp007 = µ_lnum
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µlines, πTemp007); πE != nil {
						continue
					}
					πTemp001[0] = πTemp009
					if πE = πg.CheckLocal(πF, µpat2, "pat2"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µpat2, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp015, πE = πg.IsTrue(πF, πTemp009); πE != nil {
						continue
					}
					if πTemp015 {
						goto Label78
					}
					goto Label79
					// line 204: if pat2.match(lines[_lnum]): hats += 1
					πF.SetLineno(204)
				Label78:
					// line 204: if pat2.match(lines[_lnum]): hats += 1
					πF.SetLineno(204)
					if πE = πg.CheckLocal(πF, µhats, "hats"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IAdd(πF, µhats, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µhats = πTemp007
					goto Label80
				Label79:
					// line 205: else: break
					πF.SetLineno(205)
					πTemp008 = true
					continue
					goto Label80
				Label80:
					continue
				Label76:
					if πE != nil || πR != nil {
						continue
					}
				Label77:
					// line 206: lnum = lnum - hats
					πF.SetLineno(206)
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhats, "hats"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µlnum, µhats); πE != nil {
						continue
					}
					µlnum = πTemp002
					// line 207: break
					πF.SetLineno(207)
					πTemp006 = true
					continue
					goto Label72
				Label72:
					goto Label67
				Label67:
					// line 208: lnum = lnum - 1
					πF.SetLineno(208)
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µlnum, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µlnum = πTemp002
					continue
				Label64:
					if πE != nil || πR != nil {
						continue
					}
				Label65:
					// line 209: return lines, lnum
					πF.SetLineno(209)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µlines, µlnum).ToObject()
					πR = πTemp002
					continue
					goto Label55
				Label55:
					// line 211: try: # turn instances into classes
					πF.SetLineno(211)
					πF.PushCheckpoint(82)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp009, πE = πg.ResolveGlobal(πF, ßisclass); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp010); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(!πTemp008).ToObject()
					πTemp002 = πTemp007
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label83
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp013 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp013[0] = µobject
					if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					πTemp001[0] = πTemp009
					if πTemp007, πE = πg.ResolveGlobal(πF, ßisclass); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp009
				Label83:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label84
					}
					goto Label85
					// line 212: if not isclass(object) and isclass(type(object)): # __class__
					πF.SetLineno(212)
				Label84:
					// line 213: object = object.__class__ #XXX: sometimes type(class) is better?
					πF.SetLineno(213)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ß__class__, nil); πE != nil {
						continue
					}
					µobject = πTemp002
					goto Label85
				Label85:
					πF.PopCheckpoint()
					goto Label81
				Label82:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label86
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 215: except AttributeError: pass
					πF.SetLineno(215)
				Label86:
					// line 215: except AttributeError: pass
					πF.SetLineno(215)
					πF.RestoreExc(nil, nil)
					goto Label81
				Label81:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisclass); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label87
					}
					goto Label88
					// line 216: if isclass(object):
					πF.SetLineno(216)
				Label87:
					// line 217: name = object.__name__
					πF.SetLineno(217)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ß__name__, nil); πE != nil {
						continue
					}
					µname = πTemp002
					// line 218: pat = re.compile(r'^(\s*)class\s*' + name + r'\b')
					πF.SetLineno(218)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Add(πF, πg.NewStr("^(\\s*)class\\s*").ToObject(), µname); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp007, πg.NewStr("\\b").ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpat = πTemp002
					// line 222: candidates = []
					πF.SetLineno(222)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcandidates = πTemp002
					πTemp001 = πF.MakeArgs(3)
					πTemp013 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp013[0] = µlines
					if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					if πTemp007, πE = πg.Sub(πF, πTemp010, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp007
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[1] = πTemp007
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[2] = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp009); πE != nil {
						continue
					}
					πF.PushCheckpoint(90)
					πTemp006 = false
				Label89:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label91
					}
					if πTemp007, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp008 = !isStop
					} else {
						πTemp008 = true
						µi = πTemp007
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(89)            
					// line 224: match = pat.match(lines[i])
					πF.SetLineno(224)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µlines, πTemp007); πE != nil {
						continue
					}
					πTemp001[0] = πTemp009
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µpat, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmatch = πTemp009
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µmatch); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label92
					}
					goto Label93
					// line 225: if match:
					πF.SetLineno(225)
				Label92:
					πTemp009 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp011 = µi
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetItem(πF, µlines, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, πTemp014, πTemp009); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Eq(πF, πTemp010, ßc.ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label94
					}
					goto Label95
					// line 227: if lines[i][0] == 'c':
					πF.SetLineno(227)
				Label94:
					// line 228: return lines, i
					πF.SetLineno(228)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple2(µlines, µi).ToObject()
					πR = πTemp007
					continue
					goto Label95
				Label95:
					// line 230: candidates.append((match.group(1), i))
					πF.SetLineno(230)
					πTemp001 = πF.MakeArgs(1)
					πTemp013 = πF.MakeArgs(1)
					πTemp013[0] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple2(πTemp010, µi).ToObject()
					πTemp001[0] = πTemp007
					if πE = πg.CheckLocal(πF, µcandidates, "candidates"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µcandidates, ßappend, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label93
				Label93:
					continue
				Label90:
					if πE != nil || πR != nil {
						continue
					}
				Label91:
					if πE = πg.CheckLocal(πF, µcandidates, "candidates"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µcandidates); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label96
					}
					goto Label97
					// line 231: if candidates:
					πF.SetLineno(231)
				Label96:
					// line 234: candidates.sort()
					πF.SetLineno(234)
					if πE = πg.CheckLocal(πF, µcandidates, "candidates"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcandidates, ßsort, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 235: return lines, candidates[0][1]
					πF.SetLineno(235)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp007 = πg.NewInt(1).ToObject()
					πTemp010 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µcandidates, "candidates"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µcandidates, πTemp010); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, πTemp011, πTemp007); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µlines, πTemp009).ToObject()
					πR = πTemp002
					continue
					goto Label98
				Label97:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("could not find class definition").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 237: raise IOError('could not find class definition')
					πF.SetLineno(237)
					πE = πF.Raise(πTemp007, nil, nil)
					continue
					goto Label98
				Label98:
					goto Label88
				Label88:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("could not find code object").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 238: raise IOError('could not find code object')
					πF.SetLineno(238)
					πE = πF.Raise(πTemp007, nil, nil)
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßfindsource.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 106: """Return the entire source file and starting line number for an object.
			πF.SetLineno(106)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Return the entire source file and starting line number for an object.\n    For interactively-defined objects, the 'file' is the interpreter's history.\n\n    The argument may be a module, class, method, function, traceback, frame,\n    or code object.  The source code is returned as a list of all the lines\n    in the file and the line number indexes a line in that list.  An IOError\n    is raised if the source code cannot be retrieved, while a TypeError is\n    raised for objects where the source code is unavailable (e.g. builtins).").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßfindsource); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
				continue
			}
			// line 241: def getblocks(object, lstrip=False, enclosing=False, locate=False):
			πF.SetLineno(241)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "lstrip", Def: πTemp008}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "enclosing", Def: πTemp008}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "locate", Def: πTemp008}
			πTemp007 = πg.NewFunction(πg.NewCode("getblocks", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µlstrip *πg.Object = πArgs[1]; _ = µlstrip
				var µenclosing *πg.Object = πArgs[2]; _ = µenclosing
				var µlocate *πg.Object = πArgs[3]; _ = µlocate
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µlnum *πg.Object = πg.UnboundLocal; _ = µlnum
				var µindent *πg.Object = πg.UnboundLocal; _ = µindent
				var µblock *πg.Object = πg.UnboundLocal; _ = µblock
				var µpat1 *πg.Object = πg.UnboundLocal; _ = µpat1
				var µpat2 *πg.Object = πg.UnboundLocal; _ = µpat2
				var µskip *πg.Object = πg.UnboundLocal; _ = µskip
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var µblocks *πg.Object = πg.UnboundLocal; _ = µblocks
				var µ_lnum *πg.Object = πg.UnboundLocal; _ = µ_lnum
				var µtarget *πg.Object = πg.UnboundLocal; _ = µtarget
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
				var µ_line *πg.Object = πg.UnboundLocal; _ = µ_line
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 []πg.Param
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 33: goto Label33
					case 34: goto Label34
					case 14: goto Label14
					case 15: goto Label15
					case 24: goto Label24
					case 31: goto Label31
					default: panic("unexpected function state")
					}
					// line 242: """Return a list of source lines and starting line number for an object.
					πF.SetLineno(242)
					// line 251: lines, lnum = findsource(object)
					πF.SetLineno(251)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßfindsource); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µlines = πTemp002
					µlnum = πTemp004
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßismodule); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 253: if ismodule(object):
					πF.SetLineno(253)
				Label1:
					if πE = πg.CheckLocal(πF, µlstrip, "lstrip"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µlstrip); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					goto Label4
					// line 254: if lstrip: lines = _outdent(lines)
					πF.SetLineno(254)
				Label3:
					// line 254: if lstrip: lines = _outdent(lines)
					πF.SetLineno(254)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp001[0] = µlines
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_outdent); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlines = πTemp003
					goto Label4
				Label4:
					// line 255: return ([lines], [0]) if locate is True else [lines]
					πF.SetLineno(255)
					if πE = πg.CheckLocal(πF, µlocate, "locate"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µlocate == πTemp004).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label5
					}
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp001[0] = µlines
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = πg.NewInt(0).ToObject()
					πTemp006 = πg.NewList(πTemp001...).ToObject()
					πTemp003 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
					πTemp002 = πTemp003
					goto Label6
				Label5:
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp001[0] = µlines
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					πTemp002 = πTemp003
				Label6:
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 258: indent = indentsize(lines[lnum])
					πF.SetLineno(258)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					πTemp002 = µlnum
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßindentsize); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µindent = πTemp003
					// line 259: block = getblock(lines[lnum:]) #XXX: catch any TokenError here?
					πF.SetLineno(259)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µlnum, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetblock); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µblock = πTemp003
					if πE = πg.CheckLocal(πF, µenclosing, "enclosing"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µenclosing); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µindent); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp002 = πTemp003
				Label7:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 261: if not enclosing or not indent:
					πF.SetLineno(261)
				Label8:
					if πE = πg.CheckLocal(πF, µlstrip, "lstrip"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µlstrip); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label10
					}
					goto Label11
					// line 262: if lstrip: block = _outdent(block)
					πF.SetLineno(262)
				Label10:
					// line 262: if lstrip: block = _outdent(block)
					πF.SetLineno(262)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
						continue
					}
					πTemp001[0] = µblock
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_outdent); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µblock = πTemp003
					goto Label11
				Label11:
					// line 263: return ([block], [lnum]) if locate is True else [block]
					πF.SetLineno(263)
					if πE = πg.CheckLocal(πF, µlocate, "locate"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µlocate == πTemp004).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label12
					}
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
						continue
					}
					πTemp001[0] = µblock
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					πTemp001[0] = µlnum
					πTemp006 = πg.NewList(πTemp001...).ToObject()
					πTemp003 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
					πTemp002 = πTemp003
					goto Label13
				Label12:
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
						continue
					}
					πTemp001[0] = µblock
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					πTemp002 = πTemp003
				Label13:
					πR = πTemp002
					continue
					goto Label9
				Label9:
					// line 265: pat1 = r'^(\s*def\s)|(.*(?<!\w)lambda(:|\s))'; pat1 = re.compile(pat1)
					πF.SetLineno(265)
					µpat1 = πg.NewStr("^(\\s*def\\s)|(.*(?<!\\w)lambda(:|\\s))").ToObject()
					// line 265: pat1 = r'^(\s*def\s)|(.*(?<!\w)lambda(:|\s))'; pat1 = re.compile(pat1)
					πF.SetLineno(265)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpat1, "pat1"); πE != nil {
						continue
					}
					πTemp001[0] = µpat1
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpat1 = πTemp002
					// line 266: pat2 = r'^(\s*@)'; pat2 = re.compile(pat2)
					πF.SetLineno(266)
					µpat2 = πg.NewStr("^(\\s*@)").ToObject()
					// line 266: pat2 = r'^(\s*@)'; pat2 = re.compile(pat2)
					πF.SetLineno(266)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpat2, "pat2"); πE != nil {
						continue
					}
					πTemp001[0] = µpat2
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpat2 = πTemp002
					// line 271: skip = 0
					πF.SetLineno(271)
					µskip = πg.NewInt(0).ToObject()
					// line 272: line = 0
					πF.SetLineno(272)
					µline = πg.NewInt(0).ToObject()
					// line 273: blocks = []; _lnum = []
					πF.SetLineno(273)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µblocks = πTemp002
					// line 273: blocks = []; _lnum = []
					πF.SetLineno(273)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µ_lnum = πTemp002
					// line 274: target = ''.join(block)
					πF.SetLineno(274)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
						continue
					}
					πTemp001[0] = µblock
					if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtarget = πTemp003
					// line 275: while line <= lnum: #XXX: repeat lnum? or until line < lnum?
					πF.SetLineno(275)
					πF.PushCheckpoint(15)
					πTemp005 = false
				Label14:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LE(πF, µline, µlnum); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(14)            
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp002 = µline
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µpat1, "pat1"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpat1, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label17
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp002 = µline
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µpat2, "pat2"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpat2, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label18
					}
					goto Label19
					// line 277: if pat1.match(lines[line]):
					πF.SetLineno(277)
				Label17:
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µskip); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label21
					}
					goto Label22
					// line 278: if not skip:
					πF.SetLineno(278)
				Label21:
					// line 279: try: code = getblock(lines[line:])
					πF.SetLineno(279)
					πF.PushCheckpoint(24)
					// line 279: try: code = getblock(lines[line:])
					πF.SetLineno(279)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µline, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetblock); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp003
					πF.PopCheckpoint()
					goto Label23
				Label24:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTokenError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label25
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 280: except TokenError: code = [lines[line]]
					πF.SetLineno(280)
				Label25:
					// line 280: except TokenError: code = [lines[line]]
					πF.SetLineno(280)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp002 = µline
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcode = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label23
				Label23:
					goto Label22
				Label22:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp003 = µline
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µlines, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßindentsize); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, πTemp004, µindent); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label26
					}
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[0] = µcode
					if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.Contains(πF, πTemp004, µtarget); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label27
					}
					goto Label28
					// line 281: if indentsize(lines[line]) > indent: #XXX: should be >= ?
					πF.SetLineno(281)
				Label26:
					// line 282: line += len(code) - skip
					πF.SetLineno(282)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[0] = µcode
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp004, µskip); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µline, πTemp002); πE != nil {
						continue
					}
					µline = πTemp003
					goto Label29
					// line 283: elif target in ''.join(code):
					πF.SetLineno(283)
				Label27:
					// line 284: blocks.append(code) # save code block as the potential winner
					πF.SetLineno(284)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[0] = µcode
					if πE = πg.CheckLocal(πF, µblocks, "blocks"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µblocks, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 285: _lnum.append(line - skip) # save the line number for the match
					πF.SetLineno(285)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µline, µskip); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µ_lnum, "_lnum"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µ_lnum, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 286: line += len(code) - skip
					πF.SetLineno(286)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[0] = µcode
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp004, µskip); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µline, πTemp002); πE != nil {
						continue
					}
					µline = πTemp003
					goto Label29
				Label28:
					// line 288: line += 1
					πF.SetLineno(288)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µline, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µline = πTemp002
					goto Label29
				Label29:
					// line 289: skip = 0
					πF.SetLineno(289)
					µskip = πg.NewInt(0).ToObject()
					goto Label20
					// line 291: elif pat2.match(lines[line]):
					πF.SetLineno(291)
				Label18:
					// line 292: try: code = getblock(lines[line:])
					πF.SetLineno(292)
					πF.PushCheckpoint(31)
					// line 292: try: code = getblock(lines[line:])
					πF.SetLineno(292)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µline, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetblock); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp003
					πF.PopCheckpoint()
					goto Label30
				Label31:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTokenError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label32
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 293: except TokenError: code = [lines[line]]
					πF.SetLineno(293)
				Label32:
					// line 293: except TokenError: code = [lines[line]]
					πF.SetLineno(293)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp002 = µline
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcode = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label30
				Label30:
					// line 294: skip = 1
					πF.SetLineno(294)
					µskip = πg.NewInt(1).ToObject()
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µcode, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(34)
					πTemp007 = false
				Label33:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label35
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp010 = !isStop
					} else {
						πTemp010 = true
						µ_line = πTemp003
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(33)            
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_line, "_line"); πE != nil {
						continue
					}
					πTemp001[0] = µ_line
					if πE = πg.CheckLocal(πF, µpat2, "pat2"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µpat2, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp010, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp010).ToObject()
					if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label36
					}
					goto Label37
					// line 296: if not pat2.match(_line): break
					πF.SetLineno(296)
				Label36:
					// line 296: if not pat2.match(_line): break
					πF.SetLineno(296)
					πTemp007 = true
					continue
					goto Label37
				Label37:
					// line 297: skip += 1
					πF.SetLineno(297)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µskip, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µskip = πTemp003
					continue
				Label34:
					if πE != nil || πR != nil {
						continue
					}
				Label35:
					// line 298: line += skip
					πF.SetLineno(298)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µline, µskip); πE != nil {
						continue
					}
					µline = πTemp002
					goto Label20
				Label19:
					// line 301: line +=1
					πF.SetLineno(301)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µline, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µline = πTemp002
					// line 302: skip = 0
					πF.SetLineno(302)
					µskip = πg.NewInt(0).ToObject()
					goto Label20
				Label20:
					continue
				Label15:
					if πE != nil || πR != nil {
						continue
					}
				Label16:
					if πE = πg.CheckLocal(πF, µblocks, "blocks"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µblocks); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label38
					}
					goto Label39
					// line 304: if not blocks:
					πF.SetLineno(304)
				Label38:
					// line 305: blocks = [block]
					πF.SetLineno(305)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
						continue
					}
					πTemp001[0] = µblock
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µblocks = πTemp002
					// line 306: _lnum = [lnum]
					πF.SetLineno(306)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					πTemp001[0] = µlnum
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µ_lnum = πTemp002
					goto Label39
				Label39:
					if πE = πg.CheckLocal(πF, µlstrip, "lstrip"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µlstrip); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label40
					}
					goto Label41
					// line 307: if lstrip: blocks = [_outdent(block) for block in blocks]
					πF.SetLineno(307)
				Label40:
					// line 307: if lstrip: blocks = [_outdent(block) for block in blocks]
					πF.SetLineno(307)
					πTemp011 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µblock *πg.Object = πg.UnboundLocal; _ = µblock
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
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
								if πE = πg.CheckLocal(πF, µblocks, "blocks"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µblocks); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp002 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp003 = !isStop
								} else {
									πTemp003 = true
									µblock = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 307: if lstrip: blocks = [_outdent(block) for block in blocks]
								πF.SetLineno(307)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
									continue
								}
								πTemp005[0] = µblock
								if πTemp004, πE = πg.ResolveGlobal(πF, ß_outdent); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(4)
								return πTemp006, nil
							Label4:
								πTemp004 = πSent
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
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
						continue
					}
					µblocks = πTemp002
					goto Label41
				Label41:
					// line 309: return (blocks, _lnum) if locate is True else blocks
					πF.SetLineno(309)
					if πE = πg.CheckLocal(πF, µlocate, "locate"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(µlocate == πTemp006).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label42
					}
					if πE = πg.CheckLocal(πF, µblocks, "blocks"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_lnum, "_lnum"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(µblocks, µ_lnum).ToObject()
					πTemp002 = πTemp004
					goto Label43
				Label42:
					if πE = πg.CheckLocal(πF, µblocks, "blocks"); πE != nil {
						continue
					}
					πTemp002 = µblocks
				Label43:
					πR = πTemp002
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßgetblocks.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 242: """Return a list of source lines and starting line number for an object.
			πF.SetLineno(242)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Return a list of source lines and starting line number for an object.\n    Interactively-defined objects refer to lines in the interpreter's history.\n\n    If enclosing=True, then also return any enclosing code.\n    If lstrip=True, ensure there is no indentation in the first line of code.\n    If locate=True, then also return the line number for the block of code.\n\n    DEPRECATED: use 'getsourcelines' instead\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßgetblocks); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
			// line 312: def getsourcelines(object, lstrip=False, enclosing=False):
			πF.SetLineno(312)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "lstrip", Def: πTemp009}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "enclosing", Def: πTemp009}
			πTemp008 = πg.NewFunction(πg.NewCode("getsourcelines", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µlstrip *πg.Object = πArgs[1]; _ = µlstrip
				var µenclosing *πg.Object = πArgs[2]; _ = µenclosing
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 πg.KWArgs
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 313: """Return a list of source lines and starting line number for an object.
					πF.SetLineno(313)
					// line 325: code, n = getblocks(object, lstrip=lstrip, enclosing=enclosing, locate=True)
					πF.SetLineno(325)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πE = πg.CheckLocal(πF, µlstrip, "lstrip"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µenclosing, "enclosing"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp003 = πg.KWArgs{
						{"lstrip", µlstrip},
						{"enclosing", µenclosing},
						{"locate", πTemp002},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetblocks); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, πTemp003); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
						continue
					}
					µcode = πTemp002
					µn = πTemp005
					// line 326: return code[-1], n[-1]
					πF.SetLineno(326)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µcode, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp006
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µn, πTemp004); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					πR = πTemp002
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßgetsourcelines.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 313: """Return a list of source lines and starting line number for an object.
			πF.SetLineno(313)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Return a list of source lines and starting line number for an object.\n    Interactively-defined objects refer to lines in the interpreter's history.\n\n    The argument may be a module, class, method, function, traceback, frame,\n    or code object.  The source code is returned as a list of the lines\n    corresponding to the object and the line number indicates where in the\n    original source file the first line of code was found.  An IOError is\n    raised if the source code cannot be retrieved, while a TypeError is\n    raised for objects where the source code is unavailable (e.g. builtins).\n\n    If lstrip=True, ensure there is no indentation in the first line of code.\n    If enclosing=True, then also return any enclosing code.").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßgetsourcelines); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 330: def getsource(object, alias='', lstrip=False, enclosing=False, \
			πF.SetLineno(330)
			πTemp004 = make([]πg.Param, 6)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp004[1] = πg.Param{Name: "alias", Def: ß.ToObject()}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "lstrip", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "enclosing", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[4] = πg.Param{Name: "force", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[5] = πg.Param{Name: "builtin", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("getsource", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µalias *πg.Object = πArgs[1]; _ = µalias
				var µlstrip *πg.Object = πArgs[2]; _ = µlstrip
				var µenclosing *πg.Object = πArgs[3]; _ = µenclosing
				var µforce *πg.Object = πArgs[4]; _ = µforce
				var µbuiltin *πg.Object = πArgs[5]; _ = µbuiltin
				var µhascode *πg.Object = πg.UnboundLocal; _ = µhascode
				var µinstance *πg.Object = πg.UnboundLocal; _ = µinstance
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µlnum *πg.Object = πg.UnboundLocal; _ = µlnum
				var µ_import *πg.Object = πg.UnboundLocal; _ = µ_import
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µ_alias *πg.Object = πg.UnboundLocal; _ = µ_alias
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
				var µskip *πg.Object = πg.UnboundLocal; _ = µskip
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 πg.KWArgs
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 []πg.Param
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 bool
				_ = πTemp014
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 35: goto Label35
					case 34: goto Label34
					default: panic("unexpected function state")
					}
					// line 332: """Return the text of the source code for an object. The source code for
					πF.SetLineno(332)
					// line 348: hascode = _hascode(object)
					πF.SetLineno(348)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_hascode); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µhascode = πTemp003
					// line 350: instance = _isinstance(object)
					πF.SetLineno(350)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_isinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µinstance = πTemp003
					// line 353: try: # fails for builtins, and other assorted object types
					πF.SetLineno(353)
					πF.PushCheckpoint(2)
					// line 354: lines, lnum = getsourcelines(object, enclosing=enclosing)
					πF.SetLineno(354)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πE = πg.CheckLocal(πF, µenclosing, "enclosing"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"enclosing", µenclosing},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetsourcelines); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µlines = πTemp002
					µlnum = πTemp005
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label3
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 355: except (TypeError, IOError): # failed to get source, resort to import hooks
					πF.SetLineno(355)
				Label3:
					if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µforce); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label4
					}
					goto Label5
					// line 356: if not force: # don't try to get types that findsource can't get
					πF.SetLineno(356)
				Label4:
					// line 357: raise
					πF.SetLineno(357)
					πE = πF.Raise(nil, nil, nil)
					continue
					goto Label5
				Label5:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label6
					}
					goto Label7
					// line 358: if not getmodule(object): # get things like 'None' and '1'
					πF.SetLineno(358)
				Label6:
					if πE = πg.CheckLocal(πF, µinstance, "instance"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µinstance); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label9
					}
					goto Label10
					// line 359: if not instance: return getimport(object, alias, builtin=builtin)
					πF.SetLineno(359)
				Label9:
					// line 359: if not instance: return getimport(object, alias, builtin=builtin)
					πF.SetLineno(359)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp001[1] = µalias
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"builtin", µbuiltin},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetimport); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					goto Label10
				Label10:
					// line 361: _import = getimport(object, builtin=builtin)
					πF.SetLineno(361)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"builtin", µbuiltin},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetimport); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µ_import = πTemp003
					// line 362: name = getname(object, force=True)
					πF.SetLineno(362)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"force", πTemp002},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µname = πTemp003
					// line 363: _alias = "%s = " % alias if alias else ""
					πF.SetLineno(363)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s = ").ToObject(), µalias); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					goto Label12
				Label11:
					πTemp002 = ß.ToObject()
				Label12:
					µ_alias = πTemp002
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µalias, µname); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label13
					}
					goto Label14
					// line 364: if alias == name: _alias = ""
					πF.SetLineno(364)
				Label13:
					// line 364: if alias == name: _alias = ""
					πF.SetLineno(364)
					µ_alias = ß.ToObject()
					goto Label14
				Label14:
					// line 365: return _import+_alias+"%s\n" % name
					πF.SetLineno(365)
					if πE = πg.CheckLocal(πF, µ_import, "_import"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_alias, "_alias"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µ_import, µ_alias); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mod(πF, πg.NewStr("%s\n").ToObject(), µname); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label8
				Label7:
					if πE = πg.CheckLocal(πF, µinstance, "instance"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µinstance); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label15
					}
					goto Label16
					// line 367: if not instance: return getimport(object, alias, builtin=builtin)
					πF.SetLineno(367)
				Label15:
					// line 367: if not instance: return getimport(object, alias, builtin=builtin)
					πF.SetLineno(367)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp001[1] = µalias
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"builtin", µbuiltin},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetimport); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					goto Label16
				Label16:
					// line 369: name = object.__class__.__name__
					πF.SetLineno(369)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ß__class__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__name__, nil); πE != nil {
						continue
					}
					µname = πTemp003
					// line 370: module = object.__module__
					πF.SetLineno(370)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobject, ß__module__, nil); πE != nil {
						continue
					}
					µmodule = πTemp002
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001 = make([]*πg.Object, 2)
					πTemp001[0] = ßbuiltins.ToObject()
					πTemp001[1] = ß__builtin__.ToObject()
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					if πTemp008, πE = πg.Contains(πF, πTemp003, µmodule); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label17
					}
					goto Label18
					// line 371: if module in ['builtins','__builtin__']:
					πF.SetLineno(371)
				Label17:
					// line 372: return getimport(object, alias, builtin=builtin)
					πF.SetLineno(372)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp001[1] = µalias
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"builtin", µbuiltin},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetimport); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					goto Label19
				Label18:
					// line 374: lines, lnum = ["%s = __import__('%s', fromlist=['%s']).%s\n" % (name,module,name,name)], 0
					πF.SetLineno(374)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple4(µname, µmodule, µname, µname).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s = __import__('%s', fromlist=['%s']).%s\n").ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πg.NewInt(0).ToObject()).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
						continue
					}
					µlines = πTemp003
					µlnum = πTemp005
					// line 375: obj = eval(lines[0].lstrip(name + ' = '))
					πF.SetLineno(375)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µname, πg.NewStr(" = ").ToObject()); πE != nil {
						continue
					}
					πTemp009[0] = πTemp002
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßlstrip, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßeval); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µobj = πTemp003
					// line 376: lines, lnum = getsourcelines(obj, enclosing=enclosing)
					πF.SetLineno(376)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πE = πg.CheckLocal(πF, µenclosing, "enclosing"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"enclosing", µenclosing},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetsourcelines); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µlines = πTemp002
					µlnum = πTemp005
					goto Label19
				Label19:
					goto Label8
				Label8:
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					if πE = πg.CheckLocal(πF, µlstrip, "lstrip"); πE != nil {
						continue
					}
					πTemp002 = µlstrip
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label20
					}
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp002 = µalias
				Label20:
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label21
					}
					goto Label22
					// line 379: if lstrip or alias:
					πF.SetLineno(379)
				Label21:
					// line 380: lines = _outdent(lines)
					πF.SetLineno(380)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp001[0] = µlines
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_outdent); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlines = πTemp003
					goto Label22
				Label22:
					if πE = πg.CheckLocal(πF, µinstance, "instance"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µinstance); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label23
					}
					goto Label24
					// line 383: if instance: #and force: #XXX: move into findsource or getsourcelines ?
					πF.SetLineno(383)
				Label23:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.Contains(πF, πTemp005, πg.NewStr("(").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label25
					}
					goto Label26
					// line 384: if '(' in repr(object): lines.append('%r\n' % object)
					πF.SetLineno(384)
				Label25:
					// line 384: if '(' in repr(object): lines.append('%r\n' % object)
					πF.SetLineno(384)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%r\n").ToObject(), µobject); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label27
				Label26:
					// line 390: lines = dumpsource(object, alias='', new=force, enclose=False)
					πF.SetLineno(390)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"alias", ß.ToObject()},
						{"new", µforce},
						{"enclose", πTemp002},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdumpsource); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlines = πTemp003
					// line 391: lines, lnum = [line+'\n' for line in lines.split('\n')][:-1], 0
					πF.SetLineno(391)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp005, πg.None}, nil); πE != nil {
						continue
					}
					πTemp012 = make([]πg.Param, 0)
					πTemp011 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp012, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µline *πg.Object = πg.UnboundLocal; _ = µline
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
								πTemp002 = πF.MakeArgs(1)
								πTemp002[0] = πg.NewStr("\n").ToObject()
								if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µlines, ßsplit, nil); πE != nil {
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
									µline = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 391: lines, lnum = [line+'\n' for line in lines.split('\n')][:-1], 0
								πF.SetLineno(391)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, µline, πg.NewStr("\n").ToObject()); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp003, nil
							Label4:
								πTemp004 = πSent
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
					if πTemp013, πE = πTemp011.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ListType.Call(πF, πg.Args{πTemp013}, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp010, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp005, πg.NewInt(0).ToObject()).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
						continue
					}
					µlines = πTemp003
					µlnum = πTemp005
					goto Label27
				Label27:
					goto Label24
				Label24:
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label28
					}
					goto Label29
					// line 395: if alias:
					πF.SetLineno(395)
				Label28:
					if πE = πg.CheckLocal(πF, µhascode, "hascode"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µhascode); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label30
					}
					if πE = πg.CheckLocal(πF, µinstance, "instance"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µinstance); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label31
					}
					goto Label32
					// line 396: if hascode:
					πF.SetLineno(396)
				Label30:
					// line 397: skip = 0
					πF.SetLineno(397)
					µskip = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µlines); πE != nil {
						continue
					}
					πF.PushCheckpoint(35)
					πTemp008 = false
				Label34:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label36
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp014 = !isStop
					} else {
						πTemp014 = true
						µline = πTemp003
					}
					if πE != nil || !πTemp014 {
						continue
					}
					πF.PushCheckpoint(34)            
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("@").ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µline, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp014, πE = πg.IsTrue(πF, πTemp010); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp014).ToObject()
					if πTemp014, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp014 {
						goto Label37
					}
					goto Label38
					// line 399: if not line.startswith('@'): break
					πF.SetLineno(399)
				Label37:
					// line 399: if not line.startswith('@'): break
					πF.SetLineno(399)
					πTemp008 = true
					continue
					goto Label38
				Label38:
					// line 400: skip += 1
					πF.SetLineno(400)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µskip, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µskip = πTemp003
					continue
				Label35:
					if πE != nil || πR != nil {
						continue
					}
				Label36:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("def ").ToObject()
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp002 = µskip
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßlstrip, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label39
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp003 = µskip
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µlines, πTemp003); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, πTemp005, πg.NewStr("lambda ").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label40
					}
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobject, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, µalias, πTemp003); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label41
					}
					goto Label42
					// line 402: if lines[skip].lstrip().startswith('def '): # we have a function
					πF.SetLineno(402)
				Label39:
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobject, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, µalias, πTemp003); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label43
					}
					goto Label44
					// line 403: if alias != object.__name__:
					πF.SetLineno(403)
				Label43:
					// line 404: lines.append('\n%s = %s\n' % (alias, object.__name__))
					πF.SetLineno(404)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobject, ß__name__, nil); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µalias, πTemp005).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("\n%s = %s\n").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label44
				Label44:
					goto Label42
					// line 405: elif 'lambda ' in lines[skip]: # we have a lambda
					πF.SetLineno(405)
				Label40:
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp003 = πg.NewInt(0).ToObject()
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("=").ToObject()
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp010 = µskip
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetItem(πF, µlines, πTemp010); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp013, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.GetItem(πF, πTemp013, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, µalias, πTemp005); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label45
					}
					goto Label46
					// line 406: if alias != lines[skip].split('=')[0].strip():
					πF.SetLineno(406)
				Label45:
					// line 407: lines[skip] = '%s = %s' % (alias, lines[skip])
					πF.SetLineno(407)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp005 = µskip
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µlines, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µalias, πTemp010).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s = %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp005 = µskip
					if πE = πg.SetItem(πF, µlines, πTemp005, πTemp003); πE != nil {
						continue
					}
					goto Label46
				Label46:
					goto Label42
					// line 409: if alias != object.__name__:
					πF.SetLineno(409)
				Label41:
					// line 410: lines.append('\n%s = %s\n' % (alias, object.__name__))
					πF.SetLineno(410)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobject, ß__name__, nil); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µalias, πTemp005).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("\n%s = %s\n").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label42
				Label42:
					goto Label33
					// line 412: if instance:
					πF.SetLineno(412)
				Label31:
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp003 = πg.NewInt(0).ToObject()
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("=").ToObject()
					if πTemp013, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp013
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetItem(πF, µlines, πTemp010); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp013, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.GetItem(πF, πTemp013, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, µalias, πTemp005); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label47
					}
					goto Label48
					// line 413: if alias != lines[-1].split('=')[0].strip():
					πF.SetLineno(413)
				Label47:
					// line 414: lines[-1] = ('%s = ' % alias) + lines[-1]
					πF.SetLineno(414)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s = ").ToObject(), µalias); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp010
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µlines, πTemp005); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πTemp010); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp010
					if πE = πg.SetItem(πF, µlines, πTemp005, πTemp003); πE != nil {
						continue
					}
					goto Label48
				Label48:
					goto Label33
				Label32:
					// line 416: name = getname(object, force=True) or object.__name__
					πF.SetLineno(416)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp001[0] = µobject
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"force", πTemp003},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label49
					}
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobject, ß__name__, nil); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label49:
					µname = πTemp002
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, µalias, µname); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label50
					}
					goto Label51
					// line 417: if alias != name:
					πF.SetLineno(417)
				Label50:
					// line 418: lines.append('\n%s = %s\n' % (alias, name))
					πF.SetLineno(418)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µalias, µname).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("\n%s = %s\n").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label51
				Label51:
					goto Label33
				Label33:
					goto Label29
				Label29:
					// line 419: return ''.join(lines)
					πF.SetLineno(419)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp001[0] = µlines
					if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßgetsource.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 332: """Return the text of the source code for an object. The source code for
			πF.SetLineno(332)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Return the text of the source code for an object. The source code for\n    interactively-defined objects are extracted from the interpreter's history.\n\n    The argument may be a module, class, method, function, traceback, frame,\n    or code object.  The source code is returned as a single string.  An\n    IOError is raised if the source code cannot be retrieved, while a\n    TypeError is raised for objects where the source code is unavailable\n    (e.g. builtins).\n\n    If alias is provided, then add a line of code that renames the object.\n    If lstrip=True, ensure there is no indentation in the first line of code.\n    If enclosing=True, then also return any enclosing code.\n    If force=True, catch (TypeError,IOError) and try to use import hooks.\n    If builtin=True, force an import for any builtins\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 422: def _hascode(object):
			πF.SetLineno(422)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("_hascode", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 423: '''True if object has an attribute that stores it's __code__'''
					πF.SetLineno(423)
					// line 424: return getattr(object,'__code__',None) or getattr(object,'func_code',None)
					πF.SetLineno(424)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					πTemp003[1] = ß__code__.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					πTemp003[1] = ßfunc_code.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
				Label1:
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_hascode.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 423: '''True if object has an attribute that stores it's __code__'''
			πF.SetLineno(423)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("True if object has an attribute that stores it's __code__").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_hascode); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 426: def _isinstance(object):
			πF.SetLineno(426)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("_isinstance", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µ_types *πg.Object = πg.UnboundLocal; _ = µ_types
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 πg.KWArgs
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 427: '''True if object is a class instance type (and is not a builtin)'''
					πF.SetLineno(427)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_hascode); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πTemp004, πE = πg.ResolveGlobal(πF, ßisclass); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πTemp004, πE = πg.ResolveGlobal(πF, ßismodule); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 428: if _hascode(object) or isclass(object) or ismodule(object):
					πF.SetLineno(428)
				Label2:
					// line 429: return False
					πF.SetLineno(429)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label3
				Label3:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πTemp004, πE = πg.ResolveGlobal(πF, ßistraceback); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πTemp004, πE = πg.ResolveGlobal(πF, ßisframe); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πTemp004, πE = πg.ResolveGlobal(πF, ßiscode); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
				Label4:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					goto Label6
					// line 430: if istraceback(object) or isframe(object) or iscode(object):
					πF.SetLineno(430)
				Label5:
					// line 431: return False
					πF.SetLineno(431)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label6
				Label6:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πTemp005, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label7
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp008[0] = µobject
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp003[0] = πTemp006
					if πTemp005, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ß__name__, nil); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = ßnumpy.ToObject()
					πTemp006 = πg.NewList(πTemp003...).ToObject()
					if πTemp007, πE = πg.Contains(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp007).ToObject()
					πTemp001 = πTemp004
				Label7:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					goto Label9
					// line 433: if not getmodule(object) and getmodule(type(object)).__name__ in ['numpy']:
					πF.SetLineno(433)
				Label8:
					// line 434: return True
					πF.SetLineno(434)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label9
				Label9:
					// line 438: _types = ('<class ',"<type 'instance'>")
					πF.SetLineno(438)
					πTemp001 = πg.NewTuple2(πg.NewStr("<class ").ToObject(), πg.NewStr("<type 'instance'>").ToObject()).ToObject()
					µ_types = πTemp001
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_types, "_types"); πE != nil {
						continue
					}
					πTemp003[0] = µ_types
					πTemp008 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp009[0] = µobject
					if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp008[0] = πTemp005
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label10
					}
					goto Label11
					// line 439: if not repr(type(object)).startswith(_types): #FIXME: weak hack
					πF.SetLineno(439)
				Label10:
					// line 440: return False
					πF.SetLineno(440)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label11
				Label11:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πTemp005, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobject, ß__module__, nil); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 2)
					πTemp003[0] = ßbuiltins.ToObject()
					πTemp003[1] = ß__builtin__.ToObject()
					πTemp006 = πg.NewList(πTemp003...).ToObject()
					if πTemp007, πE = πg.Contains(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp007).ToObject()
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp010 = πg.KWArgs{
						{"force", πTemp005},
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = ßarray.ToObject()
					πTemp005 = πg.NewList(πTemp003...).ToObject()
					if πTemp007, πE = πg.Contains(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp007).ToObject()
					πTemp001 = πTemp004
				Label12:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label13
					}
					goto Label14
					// line 441: if not getmodule(object) or object.__module__ in ['builtins','__builtin__'] or getname(object, force=True) in ['array']:
					πF.SetLineno(441)
				Label13:
					// line 442: return False
					πF.SetLineno(442)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label14
				Label14:
					// line 443: return True # by process of elimination... it's what we want
					πF.SetLineno(443)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_isinstance.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 427: '''True if object is a class instance type (and is not a builtin)'''
			πF.SetLineno(427)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("True if object is a class instance type (and is not a builtin)").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_isinstance); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 446: def _intypes(object):
			πF.SetLineno(446)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("_intypes", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µtypes *πg.Object = πg.UnboundLocal; _ = µtypes
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 447: '''check if object is in the 'types' module'''
					πF.SetLineno(447)
					// line 448: import types
					πF.SetLineno(448)
					if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µtypes = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp002[0] = µobject
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = ß.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πg.GetBool(πTemp004 != πTemp005).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 450: if type(object) is not type(''):
					πF.SetLineno(450)
				Label1:
					// line 451: object = getname(object, force=True)
					πF.SetLineno(451)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp002[0] = µobject
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"force", πTemp001},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µobject = πTemp003
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µobject, ßellipsis.ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 452: if object == 'ellipsis': object = 'EllipsisType'
					πF.SetLineno(452)
				Label3:
					// line 452: if object == 'ellipsis': object = 'EllipsisType'
					πF.SetLineno(452)
					µobject = ßEllipsisType.ToObject()
					goto Label4
				Label4:
					// line 453: return True if hasattr(types, object) else False
					πF.SetLineno(453)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtypes, "types"); πE != nil {
						continue
					}
					πTemp002[0] = µtypes
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp002[1] = µobject
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label5
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					goto Label6
				Label5:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label6:
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_intypes.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 447: '''check if object is in the 'types' module'''
			πF.SetLineno(447)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("check if object is in the 'types' module").ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ß_intypes); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
				continue
			}
			// line 456: def _isstring(object): #XXX: isstringlike better?
			πF.SetLineno(456)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("_isstring", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 457: '''check if object is a string-like type'''
					πF.SetLineno(457)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 458: if PY3: return isinstance(object, (str, bytes))
					πF.SetLineno(458)
				Label1:
					// line 458: if PY3: return isinstance(object, (str, bytes))
					πF.SetLineno(458)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßbytes); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					πTemp003[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label2
				Label2:
					// line 459: return isinstance(object, basestring)
					πF.SetLineno(459)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_isstring.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 457: '''check if object is a string-like type'''
			πF.SetLineno(457)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("check if object is a string-like type").ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ß_isstring); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
				continue
			}
			// line 462: def indent(code, spaces=4):
			πF.SetLineno(462)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "code", Def: nil}
			πTemp004[1] = πg.Param{Name: "spaces", Def: πg.NewInt(4).ToObject()}
			πTemp014 = πg.NewFunction(πg.NewCode("indent", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcode *πg.Object = πArgs[0]; _ = µcode
				var µspaces *πg.Object = πArgs[1]; _ = µspaces
				var µindent *πg.Object = πg.UnboundLocal; _ = µindent
				var µnspaces *πg.Object = πg.UnboundLocal; _ = µnspaces
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µ_indent *πg.Object = πg.UnboundLocal; _ = µ_indent
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 463: '''indent a block of code with whitespace (default is 4 spaces)'''
					πF.SetLineno(463)
					// line 464: indent = indentsize(code)
					πF.SetLineno(464)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[0] = µcode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßindentsize); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µindent = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					πTemp001[0] = µspaces
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004 == πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 465: if type(spaces) is int: spaces = ' '*spaces
					πF.SetLineno(465)
				Label1:
					// line 465: if type(spaces) is int: spaces = ' '*spaces
					πF.SetLineno(465)
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πg.NewStr(" ").ToObject(), µspaces); πE != nil {
						continue
					}
					µspaces = πTemp002
					goto Label2
				Label2:
					// line 467: nspaces = indentsize(spaces)
					πF.SetLineno(467)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					πTemp001[0] = µspaces
					if πTemp002, πE = πg.ResolveGlobal(πF, ßindentsize); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnspaces = πTemp003
					// line 469: lines = code.split('\n')
					πF.SetLineno(469)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\n").ToObject()
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcode, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlines = πTemp003
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp006[0] = µlines
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp005 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label5
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µi = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 474: _indent = indentsize(lines[i])
					πF.SetLineno(474)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp003 = µi
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µlines, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßindentsize); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µ_indent = πTemp004
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_indent, "_indent"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GT(πF, µindent, µ_indent); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					goto Label7
					// line 475: if indent > _indent: continue
					πF.SetLineno(475)
				Label6:
					// line 475: if indent > _indent: continue
					πF.SetLineno(475)
					continue
					goto Label7
				Label7:
					// line 476: lines[i] = spaces+lines[i]
					πF.SetLineno(476)
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp004 = µi
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µlines, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µspaces, πTemp008); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp008 = µi
					if πE = πg.SetItem(πF, µlines, πTemp008, πTemp004); πE != nil {
						continue
					}
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µlines, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp004, ß.ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 488: if lines[-1].strip() == '': lines[-1] = ''
					πF.SetLineno(488)
				Label8:
					// line 488: if lines[-1].strip() == '': lines[-1] = ''
					πF.SetLineno(488)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ß.ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πE = πg.SetItem(πF, µlines, πTemp003, πTemp002); πE != nil {
						continue
					}
					goto Label9
				Label9:
					// line 489: return '\n'.join(lines)
					πF.SetLineno(489)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp001[0] = µlines
					if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßindent.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 463: '''indent a block of code with whitespace (default is 4 spaces)'''
			πF.SetLineno(463)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("indent a block of code with whitespace (default is 4 spaces)").ToObject()); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßindent); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
				continue
			}
			// line 492: def _outdent(lines, spaces=None, all=True):
			πF.SetLineno(492)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "lines", Def: nil}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "spaces", Def: πTemp016}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "all", Def: πTemp016}
			πTemp015 = πg.NewFunction(πg.NewCode("_outdent", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µlines *πg.Object = πArgs[0]; _ = µlines
				var µspaces *πg.Object = πArgs[1]; _ = µspaces
				var µall *πg.Object = πArgs[2]; _ = µall
				var µindent *πg.Object = πg.UnboundLocal; _ = µindent
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µ_indent *πg.Object = πg.UnboundLocal; _ = µ_indent
				var µ_spaces *πg.Object = πg.UnboundLocal; _ = µ_spaces
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 493: '''outdent lines of code, accounting for docs and line continuations'''
					πF.SetLineno(493)
					// line 494: indent = indentsize(lines[0])
					πF.SetLineno(494)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßindentsize); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µindent = πTemp003
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µspaces == πTemp005).ToObject()
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GT(πF, µspaces, µindent); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µspaces, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 495: if spaces is None or spaces > indent or spaces < 0: spaces = indent
					πF.SetLineno(495)
				Label2:
					// line 495: if spaces is None or spaces > indent or spaces < 0: spaces = indent
					πF.SetLineno(495)
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					µspaces = µindent
					goto Label3
				Label3:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µall, "all"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µall); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label4
					}
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp006[0] = µlines
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp003 = πTemp007
					goto Label5
				Label4:
					πTemp003 = πg.NewInt(1).ToObject()
				Label5:
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(7)
					πTemp004 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label8
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp008 = !isStop
					} else {
						πTemp008 = true
						µi = πTemp003
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(6)            
					// line 498: _indent = indentsize(lines[i])
					πF.SetLineno(498)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp003 = µi
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µlines, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßindentsize); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µ_indent = πTemp005
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_indent, "_indent"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GT(πF, µspaces, µ_indent); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label9
					}
					goto Label10
					// line 499: if spaces > _indent: _spaces = _indent
					πF.SetLineno(499)
				Label9:
					// line 499: if spaces > _indent: _spaces = _indent
					πF.SetLineno(499)
					if πE = πg.CheckLocal(πF, µ_indent, "_indent"); πE != nil {
						continue
					}
					µ_spaces = µ_indent
					goto Label11
				Label10:
					// line 500: else: _spaces = spaces
					πF.SetLineno(500)
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					µ_spaces = µspaces
					goto Label11
				Label11:
					// line 501: lines[i] = lines[i][_spaces:]
					πF.SetLineno(501)
					if πE = πg.CheckLocal(πF, µ_spaces, "_spaces"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µ_spaces, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µlines, πTemp007); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πE = πg.SetItem(πF, µlines, πTemp007, πTemp003); πE != nil {
						continue
					}
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 502: return lines
					πF.SetLineno(502)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πR = µlines
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_outdent.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 493: '''outdent lines of code, accounting for docs and line continuations'''
			πF.SetLineno(493)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("outdent lines of code, accounting for docs and line continuations").ToObject()); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ß_outdent); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
				continue
			}
			// line 504: def outdent(code, spaces=None, all=True):
			πF.SetLineno(504)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "code", Def: nil}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "spaces", Def: πTemp017}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "all", Def: πTemp017}
			πTemp016 = πg.NewFunction(πg.NewCode("outdent", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcode *πg.Object = πArgs[0]; _ = µcode
				var µspaces *πg.Object = πArgs[1]; _ = µspaces
				var µall *πg.Object = πArgs[2]; _ = µall
				var µindent *πg.Object = πg.UnboundLocal; _ = µindent
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 πg.KWArgs
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 505: '''outdent a block of code (default is to strip all leading whitespace)'''
					πF.SetLineno(505)
					// line 506: indent = indentsize(code)
					πF.SetLineno(506)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[0] = µcode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßindentsize); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µindent = πTemp003
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µspaces == πTemp005).ToObject()
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GT(πF, µspaces, µindent); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µspaces, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 507: if spaces is None or spaces > indent or spaces < 0: spaces = indent
					πF.SetLineno(507)
				Label2:
					// line 507: if spaces is None or spaces > indent or spaces < 0: spaces = indent
					πF.SetLineno(507)
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					µspaces = µindent
					goto Label3
				Label3:
					if πE = πg.CheckLocal(πF, µall, "all"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µall); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 509: if not all: return code[spaces:]
					πF.SetLineno(509)
				Label4:
					// line 509: if not all: return code[spaces:]
					πF.SetLineno(509)
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µspaces, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µcode, πTemp002); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label5
				Label5:
					// line 510: return '\n'.join(_outdent(code.split('\n'), spaces=spaces, all=all))
					πF.SetLineno(510)
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewStr("\n").ToObject()
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcode, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp006[0] = πTemp003
					if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µall, "all"); πE != nil {
						continue
					}
					πTemp008 = πg.KWArgs{
						{"spaces", µspaces},
						{"all", µall},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_outdent); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp006, πTemp008); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßoutdent.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 505: '''outdent a block of code (default is to strip all leading whitespace)'''
			πF.SetLineno(505)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("outdent a block of code (default is to strip all leading whitespace)").ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßoutdent); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
				continue
			}
			// line 515: __globals__ = globals()
			πF.SetLineno(515)
			if πTemp017, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp017.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__globals__.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 516: __locals__ = locals()
			πF.SetLineno(516)
			if πTemp017, πE = πg.ResolveGlobal(πF, ßlocals); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp017.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__locals__.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 517: wrap2 = '''
			πF.SetLineno(517)
			πTemp018 = πg.NewTuple2(ß__globals__.ToObject(), ß__locals__.ToObject()).ToObject()
			if πTemp017, πE = πg.Mod(πF, πg.NewStr("\ndef _wrap(f):\n    \"\"\" encapsulate a function and it's __import__ \"\"\"\n    def func(*args, **kwds):\n        try:\n            # _ = eval(getsource(f, force=True)) #XXX: safer but less robust\n            exec getimportable(f, alias='_') in %s, %s\n        except:\n            raise ImportError('cannot import name ' + f.__name__)\n        return _(*args, **kwds)\n    func.__name__ = f.__name__\n    func.__doc__ = f.__doc__\n    return func\n").ToObject(), πTemp018); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwrap2.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 531: wrap3 = '''
			πF.SetLineno(531)
			πTemp018 = πg.NewTuple2(ß__globals__.ToObject(), ß__locals__.ToObject()).ToObject()
			if πTemp017, πE = πg.Mod(πF, πg.NewStr("\ndef _wrap(f):\n    \"\"\" encapsulate a function and it's __import__ \"\"\"\n    def func(*args, **kwds):\n        try:\n            # _ = eval(getsource(f, force=True)) #XXX: safer but less robust\n            exec(getimportable(f, alias='_'), %s, %s)\n        except:\n            raise ImportError('cannot import name ' + f.__name__)\n        return _(*args, **kwds)\n    func.__name__ = f.__name__\n    func.__doc__ = f.__doc__\n    return func\n").ToObject(), πTemp018); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwrap3.ToObject(), πTemp017); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp019, πE = πg.IsTrue(πF, πTemp017); πE != nil {
				continue
			}
			if πTemp019 {
				goto Label1
			}
			goto Label2
			// line 545: if PY3:
			πF.SetLineno(545)
		Label1:
			// line 546: exec(wrap3)
			πF.SetLineno(546)
			πE = πF.RaiseType(πg.NotImplementedErrorType, "exec is not available on Grumpy. Maybe never be.")
			continue
			goto Label3
		Label2:
			// line 548: exec(wrap2)
			πF.SetLineno(548)
			πE = πF.RaiseType(πg.NotImplementedErrorType, "exec is not available on Grumpy. Maybe never be.")
			continue
			goto Label3
		Label3:
			// line 549: del wrap2, wrap3
			πF.SetLineno(549)
			if πE = πg.DelVar(πF, πF.Globals(), ßwrap2); πE != nil {
				continue
			}
			if πE = πg.DelVar(πF, πF.Globals(), ßwrap3); πE != nil {
				continue
			}
			// line 552: def _enclose(object, alias=''): #FIXME: needs alias to hold returned object
			πF.SetLineno(552)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp004[1] = πg.Param{Name: "alias", Def: ß.ToObject()}
			πTemp017 = πg.NewFunction(πg.NewCode("_enclose", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µalias *πg.Object = πArgs[1]; _ = µalias
				var µdummy *πg.Object = πg.UnboundLocal; _ = µdummy
				var µstub *πg.Object = πg.UnboundLocal; _ = µstub
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 πg.KWArgs
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 553: """create a function enclosure around the source of some object"""
					πF.SetLineno(553)
					// line 555: dummy = '__this_is_a_big_dummy_enclosing_function__'
					πF.SetLineno(555)
					µdummy = ß__this_is_a_big_dummy_enclosing_function__.ToObject()
					// line 556: stub = '__this_is_a_stub_variable__'
					πF.SetLineno(556)
					µstub = ß__this_is_a_stub_variable__.ToObject()
					// line 557: code = 'def %s():\n' % dummy
					πF.SetLineno(557)
					if πE = πg.CheckLocal(πF, µdummy, "dummy"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("def %s():\n").ToObject(), µdummy); πE != nil {
						continue
					}
					µcode = πTemp001
					// line 558: code += indent(getsource(object, alias=stub, lstrip=True, force=True))
					πF.SetLineno(558)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp003[0] = µobject
					if πE = πg.CheckLocal(πF, µstub, "stub"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"alias", µstub},
						{"lstrip", πTemp001},
						{"force", πTemp004},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[0] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßindent); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.IAdd(πF, µcode, πTemp004); πE != nil {
						continue
					}
					µcode = πTemp001
					// line 559: code += indent('return %s\n' % stub)
					πF.SetLineno(559)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstub, "stub"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("return %s\n").ToObject(), µstub); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßindent); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.IAdd(πF, µcode, πTemp004); πE != nil {
						continue
					}
					µcode = πTemp001
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 560: if alias: code += '%s = ' % alias
					πF.SetLineno(560)
				Label1:
					// line 560: if alias: code += '%s = ' % alias
					πF.SetLineno(560)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s = ").ToObject(), µalias); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µcode, πTemp001); πE != nil {
						continue
					}
					µcode = πTemp004
					goto Label2
				Label2:
					// line 561: code += '%s(); del %s\n' % (dummy, dummy)
					πF.SetLineno(561)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdummy, "dummy"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdummy, "dummy"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(µdummy, µdummy).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s(); del %s\n").ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µcode, πTemp001); πE != nil {
						continue
					}
					µcode = πTemp004
					// line 563: return code
					πF.SetLineno(563)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πR = µcode
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_enclose.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 553: """create a function enclosure around the source of some object"""
			πF.SetLineno(553)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("create a function enclosure around the source of some object").ToObject()); πE != nil {
				continue
			}
			if πTemp020, πE = πg.ResolveGlobal(πF, ß_enclose); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp018); πE != nil {
				continue
			}
			// line 566: def dumpsource(object, alias='', new=False, enclose=True):
			πF.SetLineno(566)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp004[1] = πg.Param{Name: "alias", Def: ß.ToObject()}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "new", Def: πTemp020}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "enclose", Def: πTemp020}
			πTemp018 = πg.NewFunction(πg.NewCode("dumpsource", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µalias *πg.Object = πArgs[1]; _ = µalias
				var µnew *πg.Object = πArgs[2]; _ = µnew
				var µenclose *πg.Object = πArgs[3]; _ = µenclose
				var µdumps *πg.Object = πg.UnboundLocal; _ = µdumps
				var µpik *πg.Object = πg.UnboundLocal; _ = µpik
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
				var µstub *πg.Object = πg.UnboundLocal; _ = µstub
				var µpre *πg.Object = πg.UnboundLocal; _ = µpre
				var µmod *πg.Object = πg.UnboundLocal; _ = µmod
				var µdummy *πg.Object = πg.UnboundLocal; _ = µdummy
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 πg.KWArgs
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 567: """'dump to source', where the code includes a pickled object.
					πF.SetLineno(567)
					// line 574: from dill import dumps
					πF.SetLineno(574)
					if πTemp002, πE = πg.ImportModule(πF, "dill"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßdumps); πE != nil {
						continue
					}
					µdumps = πTemp003
					// line 575: pik = repr(dumps(object))
					πF.SetLineno(575)
					πTemp002 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp004[0] = µobject
					if πE = πg.CheckLocal(πF, µdumps, "dumps"); πE != nil {
						continue
					}
					if πTemp001, πE = µdumps.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µpik = πTemp003
					// line 576: code = 'import dill\n'
					πF.SetLineno(576)
					µcode = πg.NewStr("import dill\n").ToObject()
					if πE = πg.CheckLocal(πF, µenclose, "enclose"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µenclose); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 577: if enclose:
					πF.SetLineno(577)
				Label1:
					// line 578: stub = '__this_is_a_stub_variable__' #XXX: *must* be same _enclose.stub
					πF.SetLineno(578)
					µstub = ß__this_is_a_stub_variable__.ToObject()
					// line 579: pre = '%s = ' % stub
					πF.SetLineno(579)
					if πE = πg.CheckLocal(πF, µstub, "stub"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s = ").ToObject(), µstub); πE != nil {
						continue
					}
					µpre = πTemp001
					// line 580: new = False #FIXME: new=True doesn't work with enclose=True
					πF.SetLineno(580)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µnew = πTemp001
					goto Label3
				Label2:
					// line 582: stub = alias
					πF.SetLineno(582)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					µstub = µalias
					// line 583: pre = '%s = ' % stub if alias else alias
					πF.SetLineno(583)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µstub, "stub"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s = ").ToObject(), µstub); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					goto Label5
				Label4:
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp001 = µalias
				Label5:
					µpre = πTemp001
					goto Label3
				Label3:
					if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µnew); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp006).ToObject()
					πTemp001 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp002[0] = µobject
					if πTemp007, πE = πg.ResolveGlobal(πF, ß_isinstance); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp006, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp006).ToObject()
					πTemp001 = πTemp003
				Label6:
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label7
					}
					goto Label8
					// line 586: if not new or not _isinstance(object):
					πF.SetLineno(586)
				Label7:
					// line 587: code += pre + 'dill.loads(%s)\n' % pik
					πF.SetLineno(587)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpre, "pre"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpik, "pik"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("dill.loads(%s)\n").ToObject(), µpik); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µpre, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µcode, πTemp001); πE != nil {
						continue
					}
					µcode = πTemp003
					goto Label9
				Label8:
					// line 589: code += getsource(object.__class__, alias='', lstrip=True, force=True)
					πF.SetLineno(589)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobject, ß__class__, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"alias", ß.ToObject()},
						{"lstrip", πTemp001},
						{"force", πTemp003},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.IAdd(πF, µcode, πTemp003); πE != nil {
						continue
					}
					µcode = πTemp001
					// line 590: mod = repr(object.__module__) # should have a module (no builtins here)
					πF.SetLineno(590)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobject, ß__module__, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmod = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label10
					}
					goto Label11
					// line 591: if PY3:
					πF.SetLineno(591)
				Label10:
					// line 592: code += pre + 'dill.loads(%s.replace(b%s,bytes(__name__,"UTF-8")))\n' % (pik,mod)
					πF.SetLineno(592)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpre, "pre"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpik, "pik"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple2(µpik, µmod).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("dill.loads(%s.replace(b%s,bytes(__name__,\"UTF-8\")))\n").ToObject(), πTemp007); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µpre, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µcode, πTemp001); πE != nil {
						continue
					}
					µcode = πTemp003
					goto Label12
				Label11:
					// line 594: code += pre + 'dill.loads(%s.replace(%s,__name__))\n' % (pik,mod)
					πF.SetLineno(594)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpre, "pre"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpik, "pik"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple2(µpik, µmod).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("dill.loads(%s.replace(%s,__name__))\n").ToObject(), πTemp007); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µpre, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µcode, πTemp001); πE != nil {
						continue
					}
					µcode = πTemp003
					goto Label12
				Label12:
					goto Label9
				Label9:
					if πE = πg.CheckLocal(πF, µenclose, "enclose"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µenclose); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label13
					}
					goto Label14
					// line 597: if enclose:
					πF.SetLineno(597)
				Label13:
					// line 599: dummy = '__this_is_a_big_dummy_object__'
					πF.SetLineno(599)
					µdummy = ß__this_is_a_big_dummy_object__.ToObject()
					// line 600: dummy = _enclose(dummy, alias=alias)
					πF.SetLineno(600)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdummy, "dummy"); πE != nil {
						continue
					}
					πTemp002[0] = µdummy
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"alias", µalias},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_enclose); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µdummy = πTemp003
					// line 602: dummy = dummy.split('\n')
					πF.SetLineno(602)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("\n").ToObject()
					if πE = πg.CheckLocal(πF, µdummy, "dummy"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdummy, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µdummy = πTemp003
					// line 603: code = dummy[0]+'\n' + indent(code) + '\n'.join(dummy[-3:])
					πF.SetLineno(603)
					πTemp008 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µdummy, "dummy"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µdummy, πTemp008); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Add(πF, πTemp010, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp002[0] = µcode
					if πTemp008, πE = πg.ResolveGlobal(πF, ßindent); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Add(πF, πTemp007, πTemp010); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πTemp008, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{πTemp008, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdummy, "dummy"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µdummy, πTemp007); πE != nil {
						continue
					}
					πTemp002[0] = πTemp008
					if πTemp007, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Add(πF, πTemp003, πTemp008); πE != nil {
						continue
					}
					µcode = πTemp001
					goto Label14
				Label14:
					// line 605: return code #XXX: better 'dumpsourcelines', returning list of lines?
					πF.SetLineno(605)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πR = µcode
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdumpsource.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 567: """'dump to source', where the code includes a pickled object.
			πF.SetLineno(567)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("'dump to source', where the code includes a pickled object.\n\n    If new=True and object is a class instance, then create a new\n    instance using the unpacked class source code. If enclose, then\n    create the object inside a function enclosure (thus minimizing\n    any global namespace pollution).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp021, πE = πg.ResolveGlobal(πF, ßdumpsource); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
				continue
			}
			// line 608: def getname(obj, force=False, fqn=False): #XXX: throw(?) to raise error on fail?
			πF.SetLineno(608)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp021, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "force", Def: πTemp021}
			if πTemp021, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "fqn", Def: πTemp021}
			πTemp020 = πg.NewFunction(πg.NewCode("getname", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µforce *πg.Object = πArgs[1]; _ = µforce
				var µfqn *πg.Object = πArgs[2]; _ = µfqn
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8: goto Label8
					default: panic("unexpected function state")
					}
					// line 609: """get the name of the object. for lambdas, get the name of the pointer """
					πF.SetLineno(609)
					if πE = πg.CheckLocal(πF, µfqn, "fqn"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µfqn); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 610: if fqn: return '.'.join(_namespace(obj))
					πF.SetLineno(610)
				Label1:
					// line 610: if fqn: return '.'.join(_namespace(obj))
					πF.SetLineno(610)
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_namespace); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[0] = πTemp005
					if πTemp004, πE = πg.GetAttr(πF, πg.NewStr(".").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp005
					continue
					goto Label2
				Label2:
					// line 611: module = getmodule(obj)
					πF.SetLineno(611)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmodule = πTemp005
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µmodule); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label3
					}
					goto Label4
					// line 612: if not module: # things like "None" and "1"
					πF.SetLineno(612)
				Label3:
					if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µforce); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label5
					}
					goto Label6
					// line 613: if not force: return None
					πF.SetLineno(613)
				Label5:
					// line 613: if not force: return None
					πF.SetLineno(613)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp004
					continue
					goto Label6
				Label6:
					// line 614: return repr(obj)
					πF.SetLineno(614)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp005
					continue
					goto Label4
				Label4:
					// line 615: try:
					πF.SetLineno(615)
					πF.PushCheckpoint(8)
					// line 618: name = obj.__name__
					πF.SetLineno(618)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					µname = πTemp004
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µname, πg.NewStr("<lambda>").ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label9
					}
					goto Label10
					// line 619: if name == '<lambda>':
					πF.SetLineno(619)
				Label9:
					// line 620: return getsource(obj).split('=',1)[0].strip()
					πF.SetLineno(620)
					πTemp004 = πg.NewInt(0).ToObject()
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("=").ToObject()
					πTemp002[1] = πg.NewInt(1).ToObject()
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πTemp006, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πR = πTemp005
					continue
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µmodule, ß__name__, nil); πE != nil {
						continue
					}
					πTemp002 = make([]*πg.Object, 2)
					πTemp002[0] = ßbuiltins.ToObject()
					πTemp002[1] = ß__builtin__.ToObject()
					πTemp006 = πg.NewList(πTemp002...).ToObject()
					if πTemp001, πE = πg.Contains(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label11
					}
					goto Label12
					// line 622: if module.__name__ in ['builtins','__builtin__']:
					πF.SetLineno(622)
				Label11:
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µname, ßellipsis.ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label13
					}
					goto Label14
					// line 623: if name == 'ellipsis': name = 'EllipsisType'
					πF.SetLineno(623)
				Label13:
					// line 623: if name == 'ellipsis': name = 'EllipsisType'
					πF.SetLineno(623)
					µname = ßEllipsisType.ToObject()
					goto Label14
				Label14:
					goto Label12
				Label12:
					// line 624: return name
					πF.SetLineno(624)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πR = µname
					continue
					πF.PopCheckpoint()
					goto Label7
				Label8:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label15
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 625: except AttributeError: #XXX: better to just throw AttributeError ?
					πF.SetLineno(625)
				Label15:
					if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µforce); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label16
					}
					goto Label17
					// line 626: if not force: return None
					πF.SetLineno(626)
				Label16:
					// line 626: if not force: return None
					πF.SetLineno(626)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp004
					continue
					goto Label17
				Label17:
					// line 627: name = repr(obj)
					πF.SetLineno(627)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µname = πTemp005
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("<").ToObject()
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µname, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label18
					}
					goto Label19
					// line 628: if name.startswith('<'): # or name.split('('):
					πF.SetLineno(628)
				Label18:
					// line 629: return None
					πF.SetLineno(629)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp004
					continue
					goto Label19
				Label19:
					// line 630: return name
					πF.SetLineno(630)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πR = µname
					continue
					πF.RestoreExc(nil, nil)
					goto Label7
				Label7:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßgetname.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 609: """get the name of the object. for lambdas, get the name of the pointer """
			πF.SetLineno(609)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πg.NewStr("get the name of the object. for lambdas, get the name of the pointer ").ToObject()); πE != nil {
				continue
			}
			if πTemp022, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp022, ß__doc__, πTemp021); πE != nil {
				continue
			}
			// line 633: def _namespace(obj):
			πF.SetLineno(633)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp021 = πg.NewFunction(πg.NewCode("_namespace", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var µqual *πg.Object = πg.UnboundLocal; _ = µqual
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.BaseException
				_ = πTemp010
				var πTemp011 *πg.Traceback
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 14: goto Label14
					default: panic("unexpected function state")
					}
					// line 634: """_namespace(obj); return namespace hierarchy (as a list of names)
					πF.SetLineno(634)
					// line 646: try: #XXX: needs some work and testing on different types
					πF.SetLineno(646)
					πF.PushCheckpoint(2)
					// line 647: module = qual = str(getmodule(obj)).split()[1].strip('"').strip("'")
					πF.SetLineno(647)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("'").ToObject()
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("\"").ToObject()
					πTemp003 = πg.NewInt(1).ToObject()
					πTemp005 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006[0] = µobj
					if πTemp007, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[0] = πTemp008
					if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp007, πE = πg.GetAttr(πF, πTemp008, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmodule = πTemp004
					µqual = πTemp004
					// line 648: qual = qual.split('.')
					πF.SetLineno(648)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(".").ToObject()
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µqual, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µqual = πTemp004
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßismodule); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label3
					}
					goto Label4
					// line 649: if ismodule(obj):
					πF.SetLineno(649)
				Label3:
					// line 650: return qual
					πF.SetLineno(650)
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					πR = µqual
					continue
					goto Label4
				Label4:
					// line 652: name = getname(obj) or obj.__name__ # failing, raise AttributeError
					πF.SetLineno(652)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp003 = πTemp007
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					πTemp003 = πTemp004
				Label5:
					µname = πTemp003
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001 = make([]*πg.Object, 2)
					πTemp001[0] = ßbuiltins.ToObject()
					πTemp001[1] = ß__builtin__.ToObject()
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					if πTemp009, πE = πg.Contains(πF, πTemp004, µmodule); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label6
					}
					goto Label7
					// line 654: if module in ['builtins','__builtin__']: # BuiltinFunctionType
					πF.SetLineno(654)
				Label6:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_intypes); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label8
					}
					goto Label9
					// line 655: if _intypes(name): return ['types'] + [name]
					πF.SetLineno(655)
				Label8:
					// line 655: if _intypes(name): return ['types'] + [name]
					πF.SetLineno(655)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = ßtypes.ToObject()
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					πTemp007 = πg.NewList(πTemp001...).ToObject()
					if πTemp003, πE = πg.Add(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label9
				Label9:
					goto Label7
				Label7:
					// line 656: return qual + [name] #XXX: can be wrong for some aliased objects
					πF.SetLineno(656)
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					if πTemp003, πE = πg.Add(πF, µqual, πTemp004); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp010, πTemp011 = πF.ExcInfo()
					goto Label10
					// line 657: except: pass
					πF.SetLineno(657)
				Label10:
					// line 657: except: pass
					πF.SetLineno(657)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp001 = make([]*πg.Object, 4)
					πTemp001[0] = ßinf.ToObject()
					πTemp001[1] = ßnan.ToObject()
					πTemp001[2] = ßInf.ToObject()
					πTemp001[3] = ßNaN.ToObject()
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					if πTemp009, πE = πg.Contains(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label11
					}
					goto Label12
					// line 659: if str(obj) in ['inf','nan','Inf','NaN']: # is more, but are they needed?
					πF.SetLineno(659)
				Label11:
					// line 660: return ['numpy'] + [str(obj)]
					πF.SetLineno(660)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = ßnumpy.ToObject()
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp008
					πTemp007 = πg.NewList(πTemp001...).ToObject()
					if πTemp003, πE = πg.Add(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label12
				Label12:
					// line 662: module = getattr(obj.__class__, '__module__', None)
					πF.SetLineno(662)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__class__, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp001[1] = ß__module__.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[2] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmodule = πTemp004
					// line 663: qual = str(obj.__class__)
					πF.SetLineno(663)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__class__, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µqual = πTemp004
					// line 664: try: qual = qual[qual.index("'")+1:-2]
					πF.SetLineno(664)
					πF.PushCheckpoint(14)
					// line 664: try: qual = qual[qual.index("'")+1:-2]
					πF.SetLineno(664)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("'").ToObject()
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µqual, ßindex, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.Add(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πTemp007, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µqual, πTemp003); πE != nil {
						continue
					}
					µqual = πTemp004
					πF.PopCheckpoint()
					goto Label13
				Label14:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp010, πTemp011 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label15
					}
					πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
					continue
					// line 665: except ValueError: pass # str(obj.__class__) made the 'try' unnecessary
					πF.SetLineno(665)
				Label15:
					// line 665: except ValueError: pass # str(obj.__class__) made the 'try' unnecessary
					πF.SetLineno(665)
					πF.RestoreExc(nil, nil)
					goto Label13
				Label13:
					// line 666: qual = qual.split(".")
					πF.SetLineno(666)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(".").ToObject()
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µqual, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µqual = πTemp004
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001 = make([]*πg.Object, 2)
					πTemp001[0] = ßbuiltins.ToObject()
					πTemp001[1] = ß__builtin__.ToObject()
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					if πTemp009, πE = πg.Contains(πF, πTemp004, µmodule); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label16
					}
					goto Label17
					// line 667: if module in ['builtins','__builtin__']:
					πF.SetLineno(667)
				Label16:
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp007
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µqual, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp007, ßellipsis.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label18
					}
					goto Label19
					// line 669: if qual[-1] == 'ellipsis': qual[-1] = 'EllipsisType'
					πF.SetLineno(669)
				Label18:
					// line 669: if qual[-1] == 'ellipsis': qual[-1] = 'EllipsisType'
					πF.SetLineno(669)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, ßEllipsisType.ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp007
					if πE = πg.SetItem(πF, µqual, πTemp004, πTemp003); πE != nil {
						continue
					}
					goto Label19
				Label19:
					πTemp001 = πF.MakeArgs(1)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µqual, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_intypes); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label20
					}
					goto Label21
					// line 670: if _intypes(qual[-1]): module = 'types' #XXX: BuiltinFunctionType
					πF.SetLineno(670)
				Label20:
					// line 670: if _intypes(qual[-1]): module = 'types' #XXX: BuiltinFunctionType
					πF.SetLineno(670)
					µmodule = ßtypes.ToObject()
					goto Label21
				Label21:
					// line 671: qual = [module] + qual
					πF.SetLineno(671)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001[0] = µmodule
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp004, µqual); πE != nil {
						continue
					}
					µqual = πTemp003
					goto Label17
				Label17:
					// line 672: return qual
					πF.SetLineno(672)
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					πR = µqual
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_namespace.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 634: """_namespace(obj); return namespace hierarchy (as a list of names)
			πF.SetLineno(634)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πg.NewStr("_namespace(obj); return namespace hierarchy (as a list of names)\n    for the given object.  For an instance, find the class hierarchy.\n\n    For example:\n\n    >>> from functools import partial\n    >>> p = partial(int, base=2)\n    >>> _namespace(p)\n    ['functools', 'partial']\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp023, πE = πg.ResolveGlobal(πF, ß_namespace); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp023, ß__doc__, πTemp022); πE != nil {
				continue
			}
			// line 676: def _getimport(head, tail, alias='', verify=True, builtin=False):
			πF.SetLineno(676)
			πTemp004 = make([]πg.Param, 5)
			πTemp004[0] = πg.Param{Name: "head", Def: nil}
			πTemp004[1] = πg.Param{Name: "tail", Def: nil}
			πTemp004[2] = πg.Param{Name: "alias", Def: ß.ToObject()}
			if πTemp023, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "verify", Def: πTemp023}
			if πTemp023, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[4] = πg.Param{Name: "builtin", Def: πTemp023}
			πTemp022 = πg.NewFunction(πg.NewCode("_getimport", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µhead *πg.Object = πArgs[0]; _ = µhead
				var µtail *πg.Object = πArgs[1]; _ = µtail
				var µalias *πg.Object = πArgs[2]; _ = µalias
				var µverify *πg.Object = πArgs[3]; _ = µverify
				var µbuiltin *πg.Object = πArgs[4]; _ = µbuiltin
				var µ_alias *πg.Object = πg.UnboundLocal; _ = µ_alias
				var µ_str *πg.Object = πg.UnboundLocal; _ = µ_str
				var µ_head *πg.Object = πg.UnboundLocal; _ = µ_head
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 33: goto Label33
					default: panic("unexpected function state")
					}
					// line 677: """helper to build a likely import string from head and tail of namespace.
					πF.SetLineno(677)
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					πTemp004 = make([]*πg.Object, 2)
					πTemp004[0] = ßEllipsis.ToObject()
					πTemp004[1] = ßNotImplemented.ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					if πTemp006, πE = πg.Contains(πF, πTemp005, µtail); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp004 = make([]*πg.Object, 1)
					πTemp004[0] = ßtypes.ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					if πTemp006, πE = πg.Contains(πF, πTemp005, µhead); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					πTemp004 = make([]*πg.Object, 1)
					πTemp004[0] = ßNone.ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					if πTemp006, πE = πg.Contains(πF, πTemp005, µtail); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp004 = make([]*πg.Object, 1)
					πTemp004[0] = ßtypes.ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					if πTemp006, πE = πg.Contains(πF, πTemp005, µhead); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					πTemp001 = πTemp003
				Label3:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 685: if tail in ['Ellipsis', 'NotImplemented'] and head in ['types']:
					πF.SetLineno(685)
				Label2:
					// line 686: head = len.__module__
					πF.SetLineno(686)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß__module__, nil); πE != nil {
						continue
					}
					µhead = πTemp003
					goto Label5
					// line 687: elif tail in ['None'] and head in ['types']:
					πF.SetLineno(687)
				Label4:
					// line 688: _alias = '%s = ' % alias if alias else ''
					πF.SetLineno(688)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s = ").ToObject(), µalias); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					goto Label7
				Label6:
					πTemp001 = ß.ToObject()
				Label7:
					µ_alias = πTemp001
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µalias, µtail); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					goto Label9
					// line 689: if alias == tail: _alias = ''
					πF.SetLineno(689)
				Label8:
					// line 689: if alias == tail: _alias = ''
					πF.SetLineno(689)
					µ_alias = ß.ToObject()
					goto Label9
				Label9:
					// line 690: return _alias+'%s\n' % tail
					πF.SetLineno(690)
					if πE = πg.CheckLocal(πF, µ_alias, "_alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s\n").ToObject(), µtail); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µ_alias, πTemp003); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label5
				Label5:
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp004 = make([]*πg.Object, 2)
					πTemp004[0] = ßbuiltins.ToObject()
					πTemp004[1] = ß__builtin__.ToObject()
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					if πTemp002, πE = πg.Contains(πF, πTemp003, µhead); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label10
					}
					goto Label11
					// line 693: if head in ['builtins','__builtin__']:
					πF.SetLineno(693)
				Label10:
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µtail, ßellipsis.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					goto Label13
					// line 695: if tail == 'ellipsis': tail = 'EllipsisType'
					πF.SetLineno(695)
				Label12:
					// line 695: if tail == 'ellipsis': tail = 'EllipsisType'
					πF.SetLineno(695)
					µtail = ßEllipsisType.ToObject()
					goto Label13
				Label13:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					πTemp004[0] = µtail
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_intypes); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µbuiltin); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label15
					}
					goto Label16
					// line 696: if _intypes(tail): head = 'types'
					πF.SetLineno(696)
				Label14:
					// line 696: if _intypes(tail): head = 'types'
					πF.SetLineno(696)
					µhead = ßtypes.ToObject()
					goto Label17
					// line 697: elif not builtin:
					πF.SetLineno(697)
				Label15:
					// line 698: _alias = '%s = ' % alias if alias else ''
					πF.SetLineno(698)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label18
					}
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s = ").ToObject(), µalias); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					goto Label19
				Label18:
					πTemp001 = ß.ToObject()
				Label19:
					µ_alias = πTemp001
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µalias, µtail); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label20
					}
					goto Label21
					// line 699: if alias == tail: _alias = ''
					πF.SetLineno(699)
				Label20:
					// line 699: if alias == tail: _alias = ''
					πF.SetLineno(699)
					µ_alias = ß.ToObject()
					goto Label21
				Label21:
					// line 700: return _alias+'%s\n' % tail
					πF.SetLineno(700)
					if πE = πg.CheckLocal(πF, µ_alias, "_alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s\n").ToObject(), µtail); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µ_alias, πTemp003); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label17
				Label16:
					// line 701: else: pass # handle builtins below
					πF.SetLineno(701)
					goto Label17
				Label17:
					goto Label11
				Label11:
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µhead); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label22
					}
					goto Label23
					// line 703: if not head: _str = "import %s" % tail
					πF.SetLineno(703)
				Label22:
					// line 703: if not head: _str = "import %s" % tail
					πF.SetLineno(703)
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("import %s").ToObject(), µtail); πE != nil {
						continue
					}
					µ_str = πTemp001
					goto Label24
				Label23:
					// line 704: else: _str = "from %s import %s" % (head, tail)
					πF.SetLineno(704)
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µhead, µtail).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("from %s import %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					µ_str = πTemp001
					goto Label24
				Label24:
					// line 705: _alias = " as %s\n" % alias if alias else "\n"
					πF.SetLineno(705)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label25
					}
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr(" as %s\n").ToObject(), µalias); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					goto Label26
				Label25:
					πTemp001 = πg.NewStr("\n").ToObject()
				Label26:
					µ_alias = πTemp001
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µalias, µtail); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label27
					}
					goto Label28
					// line 706: if alias == tail: _alias = "\n"
					πF.SetLineno(706)
				Label27:
					// line 706: if alias == tail: _alias = "\n"
					πF.SetLineno(706)
					µ_alias = πg.NewStr("\n").ToObject()
					goto Label28
				Label28:
					// line 707: _str += _alias
					πF.SetLineno(707)
					if πE = πg.CheckLocal(πF, µ_str, "_str"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_alias, "_alias"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µ_str, µ_alias); πE != nil {
						continue
					}
					µ_str = πTemp001
					if πE = πg.CheckLocal(πF, µverify, "verify"); πE != nil {
						continue
					}
					πTemp001 = µverify
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label29
					}
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("dill.").ToObject()
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µhead, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp006).ToObject()
					πTemp001 = πTemp003
				Label29:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label30
					}
					goto Label31
					// line 711: if verify and not head.startswith('dill.'):# weird behavior for dill
					πF.SetLineno(711)
				Label30:
					// line 713: try: exec(_str) #XXX: check if == obj? (name collision)
					πF.SetLineno(713)
					πF.PushCheckpoint(33)
					// line 713: try: exec(_str) #XXX: check if == obj? (name collision)
					πF.SetLineno(713)
					πE = πF.RaiseType(πg.NotImplementedErrorType, "exec is not available on Grumpy. Maybe never be.")
					continue
					πF.PopCheckpoint()
					goto Label32
				Label33:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label34
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 714: except ImportError: #XXX: better top-down or bottom-up recursion?
					πF.SetLineno(714)
				Label34:
					// line 715: _head = head.rsplit(".",1)[0] #(or get all, then compare == obj?)
					πF.SetLineno(715)
					πTemp001 = πg.NewInt(0).ToObject()
					πTemp004 = πF.MakeArgs(2)
					πTemp004[0] = πg.NewStr(".").ToObject()
					πTemp004[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µhead, ßrsplit, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
						continue
					}
					µ_head = πTemp003
					if πE = πg.CheckLocal(πF, µ_head, "_head"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µ_head); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label35
					}
					goto Label36
					// line 716: if not _head: raise
					πF.SetLineno(716)
				Label35:
					// line 716: if not _head: raise
					πF.SetLineno(716)
					πE = πF.Raise(nil, nil, nil)
					continue
					goto Label36
				Label36:
					if πE = πg.CheckLocal(πF, µ_head, "_head"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µ_head, µhead); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label37
					}
					goto Label38
					// line 717: if _head != head:
					πF.SetLineno(717)
				Label37:
					// line 718: _str = _getimport(_head, tail, alias, verify)
					πF.SetLineno(718)
					πTemp004 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µ_head, "_head"); πE != nil {
						continue
					}
					πTemp004[0] = µ_head
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					πTemp004[1] = µtail
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp004[2] = µalias
					if πE = πg.CheckLocal(πF, µverify, "verify"); πE != nil {
						continue
					}
					πTemp004[3] = µverify
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_getimport); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µ_str = πTemp003
					goto Label38
				Label38:
					πF.RestoreExc(nil, nil)
					goto Label32
				Label32:
					goto Label31
				Label31:
					// line 719: return _str
					πF.SetLineno(719)
					if πE = πg.CheckLocal(πF, µ_str, "_str"); πE != nil {
						continue
					}
					πR = µ_str
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_getimport.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 677: """helper to build a likely import string from head and tail of namespace.
			πF.SetLineno(677)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πg.NewStr("helper to build a likely import string from head and tail of namespace.\n    ('head','tail') are used in the following context: \"from head import tail\"\n\n    If verify=True, then test the import string before returning it.\n    If builtin=True, then force an import for builtins where possible.\n    If alias is provided, then rename the object on import.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp024, πE = πg.ResolveGlobal(πF, ß_getimport); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp024, ß__doc__, πTemp023); πE != nil {
				continue
			}
			// line 724: def getimport(obj, alias='', verify=True, builtin=False, enclosing=False):
			πF.SetLineno(724)
			πTemp004 = make([]πg.Param, 5)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "alias", Def: ß.ToObject()}
			if πTemp024, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "verify", Def: πTemp024}
			if πTemp024, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "builtin", Def: πTemp024}
			if πTemp024, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[4] = πg.Param{Name: "enclosing", Def: πTemp024}
			πTemp023 = πg.NewFunction(πg.NewCode("getimport", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µalias *πg.Object = πArgs[1]; _ = µalias
				var µverify *πg.Object = πArgs[2]; _ = µverify
				var µbuiltin *πg.Object = πArgs[3]; _ = µbuiltin
				var µenclosing *πg.Object = πArgs[4]; _ = µenclosing
				var µoutermost *πg.Object = πg.UnboundLocal; _ = µoutermost
				var µ_obj *πg.Object = πg.UnboundLocal; _ = µ_obj
				var µqual *πg.Object = πg.UnboundLocal; _ = µqual
				var µhead *πg.Object = πg.UnboundLocal; _ = µhead
				var µtail *πg.Object = πg.UnboundLocal; _ = µtail
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µ_alias *πg.Object = πg.UnboundLocal; _ = µ_alias
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.BaseException
				_ = πTemp011
				var πTemp012 *πg.Traceback
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 25: goto Label25
					case 14: goto Label14
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 725: """get the likely import string for the given object
					πF.SetLineno(725)
					if πE = πg.CheckLocal(πF, µenclosing, "enclosing"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µenclosing); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 733: if enclosing:
					πF.SetLineno(733)
				Label1:
					// line 734: from .detect import outermost
					πF.SetLineno(734)
					if πTemp003, πE = πg.ImportModule(πF, ".detect"); πE != nil {
						continue
					}
					πTemp002 = πTemp003[1]
					if πTemp004, πE = πg.GetAttrImport(πF, πTemp002, ßoutermost); πE != nil {
						continue
					}
					µoutermost = πTemp004
					// line 735: _obj = outermost(obj)
					πF.SetLineno(735)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πE = πg.CheckLocal(πF, µoutermost, "outermost"); πE != nil {
						continue
					}
					if πTemp002, πE = µoutermost.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µ_obj = πTemp002
					// line 736: obj = _obj if _obj else obj
					πF.SetLineno(736)
					if πE = πg.CheckLocal(πF, µ_obj, "_obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µ_obj); πE != nil {
						continue
					}
					if !πTemp001 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µ_obj, "_obj"); πE != nil {
						continue
					}
					πTemp002 = µ_obj
					goto Label4
				Label3:
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002 = µobj
				Label4:
					µobj = πTemp002
					goto Label2
				Label2:
					// line 738: qual = _namespace(obj)
					πF.SetLineno(738)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_namespace); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µqual = πTemp004
					// line 739: head = '.'.join(qual[:-1])
					πF.SetLineno(739)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µqual, πTemp002); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(".").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µhead = πTemp004
					// line 740: tail = qual[-1]
					πF.SetLineno(740)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp004
					if πE = πg.CheckLocal(πF, µqual, "qual"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µqual, πTemp002); πE != nil {
						continue
					}
					µtail = πTemp004
					// line 742: try: # look for '<...>' and be mindful it might be in lists, dicts, etc...
					πF.SetLineno(742)
					πF.PushCheckpoint(6)
					// line 743: name = repr(obj).split('<',1)[1].split('>',1)[1]
					πF.SetLineno(743)
					πTemp002 = πg.NewInt(1).ToObject()
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = πg.NewStr(">").ToObject()
					πTemp003[1] = πg.NewInt(1).ToObject()
					πTemp005 = πg.NewInt(1).ToObject()
					πTemp007 = πF.MakeArgs(2)
					πTemp007[0] = πg.NewStr("<").ToObject()
					πTemp007[1] = πg.NewInt(1).ToObject()
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp008[0] = µobj
					if πTemp009, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp009, πE = πg.GetAttr(πF, πTemp010, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp006, πE = πg.GetItem(πF, πTemp010, πTemp005); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
						continue
					}
					µname = πTemp004
					// line 744: name = None # we have a 'object'-style repr
					πF.SetLineno(744)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µname = πTemp002
					πF.PopCheckpoint()
					goto Label5
				Label6:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp011, πTemp012 = πF.ExcInfo()
					goto Label7
					// line 745: except: # it's probably something 'importable'
					πF.SetLineno(745)
				Label7:
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 2)
					πTemp003[0] = ßbuiltins.ToObject()
					πTemp003[1] = ß__builtin__.ToObject()
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πTemp001, πE = πg.Contains(πF, πTemp004, µhead); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label8
					}
					goto Label9
					// line 746: if head in ['builtins','__builtin__']:
					πF.SetLineno(746)
				Label8:
					// line 747: name = repr(obj) #XXX: catch [1,2], (1,2), set([1,2])... others?
					πF.SetLineno(747)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µname = πTemp004
					goto Label10
				Label9:
					// line 749: name = repr(obj).split('(')[0]
					πF.SetLineno(749)
					πTemp002 = πg.NewInt(0).ToObject()
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("(").ToObject()
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007[0] = µobj
					if πTemp005, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
						continue
					}
					µname = πTemp004
					goto Label10
				Label10:
					πF.RestoreExc(nil, nil)
					goto Label5
				Label5:
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µname); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label11
					}
					goto Label12
					// line 752: if name: # try using name instead of tail
					πF.SetLineno(752)
				Label11:
					// line 753: try: return _getimport(head, name, alias, verify, builtin)
					πF.SetLineno(753)
					πF.PushCheckpoint(14)
					// line 753: try: return _getimport(head, name, alias, verify, builtin)
					πF.SetLineno(753)
					πTemp003 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp003[0] = µhead
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp003[1] = µname
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp003[2] = µalias
					if πE = πg.CheckLocal(πF, µverify, "verify"); πE != nil {
						continue
					}
					πTemp003[3] = µverify
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp003[4] = µbuiltin
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_getimport); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					πF.PopCheckpoint()
					goto Label13
				Label14:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp011, πTemp012 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label15
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label16
					}
					πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 754: except ImportError: pass
					πF.SetLineno(754)
				Label15:
					// line 754: except ImportError: pass
					πF.SetLineno(754)
					πF.RestoreExc(nil, nil)
					goto Label13
					// line 755: except SyntaxError:
					πF.SetLineno(755)
				Label16:
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 2)
					πTemp003[0] = ßbuiltins.ToObject()
					πTemp003[1] = ß__builtin__.ToObject()
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πTemp001, πE = πg.Contains(πF, πTemp004, µhead); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label17
					}
					goto Label18
					// line 756: if head in ['builtins','__builtin__']:
					πF.SetLineno(756)
				Label17:
					// line 757: _alias = '%s = ' % alias if alias else ''
					πF.SetLineno(757)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					if !πTemp001 {
						goto Label20
					}
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("%s = ").ToObject(), µalias); πE != nil {
						continue
					}
					πTemp002 = πTemp004
					goto Label21
				Label20:
					πTemp002 = ß.ToObject()
				Label21:
					µ_alias = πTemp002
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µalias, µname); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label22
					}
					goto Label23
					// line 758: if alias == name: _alias = ''
					πF.SetLineno(758)
				Label22:
					// line 758: if alias == name: _alias = ''
					πF.SetLineno(758)
					µ_alias = ß.ToObject()
					goto Label23
				Label23:
					// line 759: return _alias+'%s\n' % name
					πF.SetLineno(759)
					if πE = πg.CheckLocal(πF, µ_alias, "_alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("%s\n").ToObject(), µname); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µ_alias, πTemp004); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label19
				Label18:
					// line 760: else: pass
					πF.SetLineno(760)
					goto Label19
				Label19:
					πF.RestoreExc(nil, nil)
					goto Label13
				Label13:
					goto Label12
				Label12:
					// line 761: try:
					πF.SetLineno(761)
					πF.PushCheckpoint(25)
					// line 764: return _getimport(head, tail, alias, verify, builtin)
					πF.SetLineno(764)
					πTemp003 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp003[0] = µhead
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					πTemp003[1] = µtail
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp003[2] = µalias
					if πE = πg.CheckLocal(πF, µverify, "verify"); πE != nil {
						continue
					}
					πTemp003[3] = µverify
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp003[4] = µbuiltin
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_getimport); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					πF.PopCheckpoint()
					goto Label24
				Label25:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp011, πTemp012 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label26
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label27
					}
					πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 765: except ImportError:
					πF.SetLineno(765)
				Label26:
					// line 766: raise # could do some checking against obj
					πF.SetLineno(766)
					πE = πF.Raise(nil, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label24
					// line 767: except SyntaxError:
					πF.SetLineno(767)
				Label27:
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 2)
					πTemp003[0] = ßbuiltins.ToObject()
					πTemp003[1] = ß__builtin__.ToObject()
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πTemp001, πE = πg.Contains(πF, πTemp004, µhead); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label28
					}
					goto Label29
					// line 768: if head in ['builtins','__builtin__']:
					πF.SetLineno(768)
				Label28:
					// line 769: _alias = '%s = ' % alias if alias else ''
					πF.SetLineno(769)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					if !πTemp001 {
						goto Label30
					}
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("%s = ").ToObject(), µalias); πE != nil {
						continue
					}
					πTemp002 = πTemp004
					goto Label31
				Label30:
					πTemp002 = ß.ToObject()
				Label31:
					µ_alias = πTemp002
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µalias, µtail); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label32
					}
					goto Label33
					// line 770: if alias == tail: _alias = ''
					πF.SetLineno(770)
				Label32:
					// line 770: if alias == tail: _alias = ''
					πF.SetLineno(770)
					µ_alias = ß.ToObject()
					goto Label33
				Label33:
					// line 771: return _alias+'%s\n' % tail
					πF.SetLineno(771)
					if πE = πg.CheckLocal(πF, µ_alias, "_alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("%s\n").ToObject(), µtail); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µ_alias, πTemp004); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label29
				Label29:
					// line 772: raise # could do some checking against obj
					πF.SetLineno(772)
					πE = πF.Raise(nil, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label24
				Label24:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßgetimport.ToObject(), πTemp023); πE != nil {
				continue
			}
			// line 725: """get the likely import string for the given object
			πF.SetLineno(725)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πg.NewStr("get the likely import string for the given object\n\n    obj is the object to inspect\n    If verify=True, then test the import string before returning it.\n    If builtin=True, then force an import for builtins where possible.\n    If enclosing=True, get the import for the outermost enclosing callable.\n    If alias is provided, then rename the object on import.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp025, πE = πg.ResolveGlobal(πF, ßgetimport); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp024); πE != nil {
				continue
			}
			// line 775: def _importable(obj, alias='', source=None, enclosing=False, force=True, \
			πF.SetLineno(775)
			πTemp004 = make([]πg.Param, 7)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "alias", Def: ß.ToObject()}
			if πTemp025, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "source", Def: πTemp025}
			if πTemp025, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "enclosing", Def: πTemp025}
			if πTemp025, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[4] = πg.Param{Name: "force", Def: πTemp025}
			if πTemp025, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[5] = πg.Param{Name: "builtin", Def: πTemp025}
			if πTemp025, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[6] = πg.Param{Name: "lstrip", Def: πTemp025}
			πTemp024 = πg.NewFunction(πg.NewCode("_importable", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µalias *πg.Object = πArgs[1]; _ = µalias
				var µsource *πg.Object = πArgs[2]; _ = µsource
				var µenclosing *πg.Object = πArgs[3]; _ = µenclosing
				var µforce *πg.Object = πArgs[4]; _ = µforce
				var µbuiltin *πg.Object = πArgs[5]; _ = µbuiltin
				var µlstrip *πg.Object = πArgs[6]; _ = µlstrip
				var µ_import *πg.Object = πg.UnboundLocal; _ = µ_import
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µ_alias *πg.Object = πg.UnboundLocal; _ = µ_alias
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8: goto Label8
					case 24: goto Label24
					case 11: goto Label11
					default: panic("unexpected function state")
					}
					// line 777: """get an import string (or the source code) for the given object
					πF.SetLineno(777)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µsource == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 797: if source is None:
					πF.SetLineno(797)
				Label1:
					// line 798: source = True if isfrommain(obj) else False
					πF.SetLineno(798)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisfrommain); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label3
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					goto Label4
				Label3:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label4:
					µsource = πTemp001
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µsource); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					goto Label6
					// line 799: if source: # first try to get the source
					πF.SetLineno(799)
				Label5:
					// line 800: try:
					πF.SetLineno(800)
					πF.PushCheckpoint(8)
					// line 801: return getsource(obj, alias, enclosing=enclosing, \
					πF.SetLineno(801)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp004[1] = µalias
					if πE = πg.CheckLocal(πF, µenclosing, "enclosing"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlstrip, "lstrip"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"enclosing", µenclosing},
						{"force", µforce},
						{"lstrip", µlstrip},
						{"builtin", µbuiltin},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label7
				Label8:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					goto Label9
					// line 803: except: pass
					πF.SetLineno(803)
				Label9:
					// line 803: except: pass
					πF.SetLineno(803)
					πF.RestoreExc(nil, nil)
					goto Label7
				Label7:
					goto Label6
				Label6:
					// line 804: try:
					πF.SetLineno(804)
					πF.PushCheckpoint(11)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_isinstance); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label12
					}
					goto Label13
					// line 805: if not _isinstance(obj):
					πF.SetLineno(805)
				Label12:
					// line 806: return getimport(obj, alias, enclosing=enclosing, \
					πF.SetLineno(806)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp004[1] = µalias
					if πE = πg.CheckLocal(πF, µenclosing, "enclosing"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µforce); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"enclosing", µenclosing},
						{"verify", πTemp001},
						{"builtin", µbuiltin},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetimport); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp002
					continue
					goto Label13
				Label13:
					// line 809: _import = getimport(obj, enclosing=enclosing, \
					πF.SetLineno(809)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µenclosing, "enclosing"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µforce); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"enclosing", µenclosing},
						{"verify", πTemp001},
						{"builtin", µbuiltin},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetimport); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µ_import = πTemp002
					// line 811: name = getname(obj, force=True)
					πF.SetLineno(811)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"force", πTemp001},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µname = πTemp002
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label14
					}
					goto Label15
					// line 812: if not name:
					πF.SetLineno(812)
				Label14:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("object has no atribute '__name__'").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 813: raise AttributeError("object has no atribute '__name__'")
					πF.SetLineno(813)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label15
				Label15:
					// line 814: _alias = "%s = " % alias if alias else ""
					πF.SetLineno(814)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s = ").ToObject(), µalias); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					goto Label17
				Label16:
					πTemp001 = ß.ToObject()
				Label17:
					µ_alias = πTemp001
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µalias, µname); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label18
					}
					goto Label19
					// line 815: if alias == name: _alias = ""
					πF.SetLineno(815)
				Label18:
					// line 815: if alias == name: _alias = ""
					πF.SetLineno(815)
					µ_alias = ß.ToObject()
					goto Label19
				Label19:
					// line 816: return _import+_alias+"%s\n" % name
					πF.SetLineno(816)
					if πE = πg.CheckLocal(πF, µ_import, "_import"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_alias, "_alias"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µ_import, µ_alias); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mod(πF, πg.NewStr("%s\n").ToObject(), µname); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp005); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					πF.PopCheckpoint()
					goto Label10
				Label11:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					goto Label20
					// line 818: except: pass
					πF.SetLineno(818)
				Label20:
					// line 818: except: pass
					πF.SetLineno(818)
					πF.RestoreExc(nil, nil)
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µsource); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label21
					}
					goto Label22
					// line 819: if not source: # try getsource, only if it hasn't been tried yet
					πF.SetLineno(819)
				Label21:
					// line 820: try:
					πF.SetLineno(820)
					πF.PushCheckpoint(24)
					// line 821: return getsource(obj, alias, enclosing=enclosing, \
					πF.SetLineno(821)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp004[1] = µalias
					if πE = πg.CheckLocal(πF, µenclosing, "enclosing"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlstrip, "lstrip"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"enclosing", µenclosing},
						{"force", µforce},
						{"lstrip", µlstrip},
						{"builtin", µbuiltin},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label23
				Label24:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					goto Label25
					// line 823: except: pass
					πF.SetLineno(823)
				Label25:
					// line 823: except: pass
					πF.SetLineno(823)
					πF.RestoreExc(nil, nil)
					goto Label23
				Label23:
					goto Label22
				Label22:
					// line 827: obj = getname(obj, force=force)
					πF.SetLineno(827)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"force", µforce},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µobj = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, µobj); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp009).ToObject()
					πTemp001 = πTemp002
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label26
					}
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("<").ToObject()
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001 = πTemp005
				Label26:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label27
					}
					goto Label28
					// line 829: if not obj or obj.startswith('<'):
					πF.SetLineno(829)
				Label27:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("object has no atribute '__name__'").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 830: raise AttributeError("object has no atribute '__name__'")
					πF.SetLineno(830)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label28
				Label28:
					// line 831: _alias = '%s = ' % alias if alias else ''
					πF.SetLineno(831)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label29
					}
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s = ").ToObject(), µalias); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					goto Label30
				Label29:
					πTemp001 = ß.ToObject()
				Label30:
					µ_alias = πTemp001
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µalias, µobj); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label31
					}
					goto Label32
					// line 832: if alias == obj: _alias = ''
					πF.SetLineno(832)
				Label31:
					// line 832: if alias == obj: _alias = ''
					πF.SetLineno(832)
					µ_alias = ß.ToObject()
					goto Label32
				Label32:
					// line 833: return _alias+'%s\n' % obj
					πF.SetLineno(833)
					if πE = πg.CheckLocal(πF, µ_alias, "_alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s\n").ToObject(), µobj); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µ_alias, πTemp002); πE != nil {
						continue
					}
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_importable.ToObject(), πTemp024); πE != nil {
				continue
			}
			// line 777: """get an import string (or the source code) for the given object
			πF.SetLineno(777)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πg.NewStr("get an import string (or the source code) for the given object\n\n    This function will attempt to discover the name of the object, or the repr\n    of the object, or the source code for the object. To attempt to force\n    discovery of the source code, use source=True, to attempt to force the\n    use of an import, use source=False; otherwise an import will be sought\n    for objects not defined in __main__. The intent is to build a string\n    that can be imported from a python file. obj is the object to inspect.\n    If alias is provided, then rename the object with the given alias.\n\n    If source=True, use these options:\n      If enclosing=True, then also return any enclosing code.\n      If force=True, catch (TypeError,IOError) and try to use import hooks.\n      If lstrip=True, ensure there is no indentation in the first line of code.\n\n    If source=False, use these options:\n      If enclosing=True, get the import for the outermost enclosing callable.\n      If force=True, then don't test the import string before returning it.\n      If builtin=True, then force an import for builtins where possible.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp026, πE = πg.ResolveGlobal(πF, ß_importable); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp026, ß__doc__, πTemp025); πE != nil {
				continue
			}
			// line 837: def _closuredimport(func, alias='', builtin=False):
			πF.SetLineno(837)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			πTemp004[1] = πg.Param{Name: "alias", Def: ß.ToObject()}
			if πTemp026, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "builtin", Def: πTemp026}
			πTemp025 = πg.NewFunction(πg.NewCode("_closuredimport", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µalias *πg.Object = πArgs[1]; _ = µalias
				var µbuiltin *πg.Object = πArgs[2]; _ = µbuiltin
				var µre *πg.Object = πg.UnboundLocal; _ = µre
				var µfreevars *πg.Object = πg.UnboundLocal; _ = µfreevars
				var µoutermost *πg.Object = πg.UnboundLocal; _ = µoutermost
				var µfree_vars *πg.Object = πg.UnboundLocal; _ = µfree_vars
				var µfunc_vars *πg.Object = πg.UnboundLocal; _ = µfunc_vars
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
				var µfobj *πg.Object = πg.UnboundLocal; _ = µfobj
				var µsrc *πg.Object = πg.UnboundLocal; _ = µsrc
				var µencl *πg.Object = πg.UnboundLocal; _ = µencl
				var µpat *πg.Object = πg.UnboundLocal; _ = µpat
				var µmod *πg.Object = πg.UnboundLocal; _ = µmod
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µcandidate *πg.Object = πg.UnboundLocal; _ = µcandidate
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 πg.KWArgs
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 []*πg.Object
				_ = πTemp012
				var πTemp013 []πg.Param
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 838: """get import for closured objects; return a dict of 'name' and 'import'"""
					πF.SetLineno(838)
					// line 839: import re
					πF.SetLineno(839)
					if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µre = πTemp001
					// line 840: from .detect import freevars, outermost
					πF.SetLineno(840)
					if πTemp002, πE = πg.ImportModule(πF, ".detect"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßfreevars); πE != nil {
						continue
					}
					µfreevars = πTemp003
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßoutermost); πE != nil {
						continue
					}
					µoutermost = πTemp003
					// line 841: free_vars = freevars(func)
					πF.SetLineno(841)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πE = πg.CheckLocal(πF, µfreevars, "freevars"); πE != nil {
						continue
					}
					if πTemp001, πE = µfreevars.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfree_vars = πTemp001
					// line 842: func_vars = {}
					πF.SetLineno(842)
					πTemp004 = πg.NewDict()
					πTemp001 = πTemp004.ToObject()
					µfunc_vars = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfree_vars, "free_vars"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfree_vars, ßitems, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
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
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp003); πE != nil {
							continue
						}
						µname = πTemp005
						µobj = πTemp008
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp005, πE = πg.ResolveGlobal(πF, ßisfunction); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label4
					}
					goto Label5
					// line 845: if not isfunction(obj): continue
					πF.SetLineno(845)
				Label4:
					// line 845: if not isfunction(obj): continue
					πF.SetLineno(845)
					continue
					goto Label5
				Label5:
					// line 847: fobj = free_vars.pop(name)
					πF.SetLineno(847)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002[0] = µname
					if πE = πg.CheckLocal(πF, µfree_vars, "free_vars"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfree_vars, ßpop, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfobj = πTemp005
					// line 848: src = getsource(fobj)
					πF.SetLineno(848)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfobj, "fobj"); πE != nil {
						continue
					}
					πTemp002[0] = µfobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µsrc = πTemp005
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("@").ToObject()
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsrc, ßlstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					goto Label7
					// line 849: if src.lstrip().startswith('@'): # we have a decorator
					πF.SetLineno(849)
				Label6:
					// line 850: src = getimport(fobj, alias=alias, builtin=builtin)
					πF.SetLineno(850)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfobj, "fobj"); πE != nil {
						continue
					}
					πTemp002[0] = µfobj
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"alias", µalias},
						{"builtin", µbuiltin},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetimport); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µsrc = πTemp005
					goto Label8
				Label7:
					// line 852: encl = outermost(func)
					πF.SetLineno(852)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πE = πg.CheckLocal(πF, µoutermost, "outermost"); πE != nil {
						continue
					}
					if πTemp003, πE = µoutermost.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µencl = πTemp003
					// line 854: pat = r'.*[\w\s]=\s*'+getname(encl)+r'\('+getname(fobj)
					πF.SetLineno(854)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µencl, "encl"); πE != nil {
						continue
					}
					πTemp002[0] = µencl
					if πTemp010, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp008, πE = πg.Add(πF, πg.NewStr(".*[\\w\\s]=\\s*").ToObject(), πTemp011); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, πTemp008, πg.NewStr("\\(").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfobj, "fobj"); πE != nil {
						continue
					}
					πTemp002[0] = µfobj
					if πTemp008, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Add(πF, πTemp005, πTemp010); πE != nil {
						continue
					}
					µpat = πTemp003
					// line 855: mod = getname(getmodule(encl))
					πF.SetLineno(855)
					πTemp002 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µencl, "encl"); πE != nil {
						continue
					}
					πTemp012[0] = µencl
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp002[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmod = πTemp005
					// line 857: lines,_ = findsource(encl)
					πF.SetLineno(857)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µencl, "encl"); πE != nil {
						continue
					}
					πTemp002[0] = µencl
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfindsource); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp008}}}, πTemp005); πE != nil {
						continue
					}
					µlines = πTemp003
					µ_ = πTemp008
					// line 858: candidate = [line for line in lines if getname(encl) in line and \
					πF.SetLineno(858)
					πTemp013 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µline *πg.Object = πg.UnboundLocal; _ = µline
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µlines); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp002 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp003 = !isStop
								} else {
									πTemp003 = true
									µline = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp006 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µencl, "encl"); πE != nil {
									continue
								}
								πTemp006[0] = µencl
								if πTemp007, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.Contains(πF, µline, πTemp008); πE != nil {
									continue
								}
								πTemp005 = πg.GetBool(πTemp009).ToObject()
								πTemp004 = πTemp005
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if !πTemp003 {
									goto Label4
								}
								πTemp006 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
									continue
								}
								πTemp006[0] = µpat
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πTemp006[1] = µline
								if πE = πg.CheckLocal(πF, µre, "re"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µre, ßmatch, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								πTemp004 = πTemp007
							Label4:
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label5
								}
								goto Label6
								// line 858: candidate = [line for line in lines if getname(encl) in line and \
								πF.SetLineno(858)
							Label5:
								// line 858: candidate = [line for line in lines if getname(encl) in line and \
								πF.SetLineno(858)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πF.PushCheckpoint(7)
								return µline, nil
							Label7:
								πTemp004 = πSent
								goto Label6
							Label6:
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
					if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
						continue
					}
					µcandidate = πTemp003
					if πE = πg.CheckLocal(πF, µcandidate, "candidate"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µcandidate); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label9
					}
					goto Label10
					// line 860: if not candidate:
					πF.SetLineno(860)
				Label9:
					// line 861: mod = getname(getmodule(fobj))
					πF.SetLineno(861)
					πTemp002 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfobj, "fobj"); πE != nil {
						continue
					}
					πTemp012[0] = µfobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp003.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp002[0] = πTemp008
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmod = πTemp008
					// line 863: lines,_ = findsource(fobj)
					πF.SetLineno(863)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfobj, "fobj"); πE != nil {
						continue
					}
					πTemp002[0] = µfobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfindsource); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp010}}}, πTemp008); πE != nil {
						continue
					}
					µlines = πTemp003
					µ_ = πTemp010
					// line 864: candidate = [line for line in lines \
					πF.SetLineno(864)
					πTemp013 = make([]πg.Param, 0)
					πTemp008 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µline *πg.Object = πg.UnboundLocal; _ = µline
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µlines); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp002 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp003 = !isStop
								} else {
									πTemp003 = true
									µline = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp006 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µfobj, "fobj"); πE != nil {
									continue
								}
								πTemp006[0] = µfobj
								if πTemp007, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.Contains(πF, µline, πTemp008); πE != nil {
									continue
								}
								πTemp005 = πg.GetBool(πTemp009).ToObject()
								πTemp004 = πTemp005
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if !πTemp003 {
									goto Label4
								}
								πTemp006 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
									continue
								}
								πTemp006[0] = µpat
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πTemp006[1] = µline
								if πE = πg.CheckLocal(πF, µre, "re"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µre, ßmatch, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								πTemp004 = πTemp007
							Label4:
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label5
								}
								goto Label6
								// line 864: candidate = [line for line in lines \
								πF.SetLineno(864)
							Label5:
								// line 864: candidate = [line for line in lines \
								πF.SetLineno(864)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πF.PushCheckpoint(7)
								return µline, nil
							Label7:
								πTemp004 = πSent
								goto Label6
							Label6:
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
					if πTemp010, πE = πTemp008.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp010}, nil); πE != nil {
						continue
					}
					µcandidate = πTemp003
					goto Label10
				Label10:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcandidate, "candidate"); πE != nil {
						continue
					}
					πTemp002[0] = µcandidate
					if πTemp010, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp007, πE = πg.IsTrue(πF, πTemp011); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label11
					}
					goto Label12
					// line 866: if not len(candidate): raise TypeError('import could not be found')
					πF.SetLineno(866)
				Label11:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("import could not be found").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 866: if not len(candidate): raise TypeError('import could not be found')
					πF.SetLineno(866)
					πE = πF.Raise(πTemp010, nil, nil)
					continue
					goto Label12
				Label12:
					// line 867: candidate = candidate[-1]
					πF.SetLineno(867)
					if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp010
					if πE = πg.CheckLocal(πF, µcandidate, "candidate"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µcandidate, πTemp003); πE != nil {
						continue
					}
					µcandidate = πTemp010
					// line 868: name = candidate.split('=',1)[0].split()[-1].strip()
					πF.SetLineno(868)
					if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp010
					πTemp011 = πg.NewInt(0).ToObject()
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("=").ToObject()
					πTemp002[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µcandidate, "candidate"); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, µcandidate, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp016, πE = πTemp015.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp014, πE = πg.GetItem(πF, πTemp016, πTemp011); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp014, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp011.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, πTemp014, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp010, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					µname = πTemp010
					// line 869: src = _getimport(mod, name, alias=alias, builtin=builtin)
					πF.SetLineno(869)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					πTemp002[0] = µmod
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002[1] = µname
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"alias", µalias},
						{"builtin", µbuiltin},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_getimport); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp003.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µsrc = πTemp010
					goto Label8
				Label8:
					// line 870: func_vars[name] = src
					πF.SetLineno(870)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µsrc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfunc_vars, "func_vars"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp010 = µname
					if πE = πg.SetItem(πF, µfunc_vars, πTemp010, πTemp003); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					if πE = πg.CheckLocal(πF, µfunc_vars, "func_vars"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µfunc_vars); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label13
					}
					goto Label14
					// line 871: if not func_vars:
					πF.SetLineno(871)
				Label13:
					// line 872: name = outermost(func)
					πF.SetLineno(872)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πE = πg.CheckLocal(πF, µoutermost, "outermost"); πE != nil {
						continue
					}
					if πTemp001, πE = µoutermost.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µname = πTemp001
					// line 873: mod = getname(getmodule(name))
					πF.SetLineno(873)
					πTemp002 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp012[0] = µname
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp002[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmod = πTemp003
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µmod); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µname == µfunc).ToObject()
					πTemp001 = πTemp003
				Label15:
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label16
					}
					goto Label17
					// line 874: if not mod or name is func: # then it can be handled by getimport
					πF.SetLineno(874)
				Label16:
					// line 875: name = getname(func, force=True) #XXX: better key?
					πF.SetLineno(875)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"force", πTemp001},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µname = πTemp003
					// line 876: src = getimport(func, alias=alias, builtin=builtin)
					πF.SetLineno(876)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"alias", µalias},
						{"builtin", µbuiltin},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetimport); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µsrc = πTemp003
					goto Label18
				Label17:
					// line 878: lines,_ = findsource(name)
					πF.SetLineno(878)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002[0] = µname
					if πTemp001, πE = πg.ResolveGlobal(πF, ßfindsource); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp010}}}, πTemp003); πE != nil {
						continue
					}
					µlines = πTemp001
					µ_ = πTemp010
					// line 880: candidate = [line for line in lines if getname(name) in line and \
					πF.SetLineno(880)
					πTemp013 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µline *πg.Object = πg.UnboundLocal; _ = µline
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µlines); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp002 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp003 = !isStop
								} else {
									πTemp003 = true
									µline = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp006 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								πTemp006[0] = µname
								if πTemp007, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.Contains(πF, µline, πTemp008); πE != nil {
									continue
								}
								πTemp005 = πg.GetBool(πTemp009).ToObject()
								πTemp004 = πTemp005
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if !πTemp003 {
									goto Label4
								}
								πTemp006 = πF.MakeArgs(2)
								πTemp010 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								πTemp010[0] = µname
								if πTemp008, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
									continue
								}
								if πTemp011, πE = πTemp008.Call(πF, πTemp010, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp010)
								if πTemp007, πE = πg.Add(πF, πg.NewStr(".*[\\w\\s]=\\s*").ToObject(), πTemp011); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Add(πF, πTemp007, πg.NewStr("\\(").ToObject()); πE != nil {
									continue
								}
								πTemp006[0] = πTemp005
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πTemp006[1] = µline
								if πE = πg.CheckLocal(πF, µre, "re"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µre, ßmatch, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								πTemp004 = πTemp007
							Label4:
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label5
								}
								goto Label6
								// line 880: candidate = [line for line in lines if getname(name) in line and \
								πF.SetLineno(880)
							Label5:
								// line 880: candidate = [line for line in lines if getname(name) in line and \
								πF.SetLineno(880)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πF.PushCheckpoint(7)
								return µline, nil
							Label7:
								πTemp004 = πSent
								goto Label6
							Label6:
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
					if πTemp010, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp010}, nil); πE != nil {
						continue
					}
					µcandidate = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcandidate, "candidate"); πE != nil {
						continue
					}
					πTemp002[0] = µcandidate
					if πTemp010, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp006, πE = πg.IsTrue(πF, πTemp011); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label19
					}
					goto Label20
					// line 882: if not len(candidate): raise TypeError('import could not be found')
					πF.SetLineno(882)
				Label19:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("import could not be found").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 882: if not len(candidate): raise TypeError('import could not be found')
					πF.SetLineno(882)
					πE = πF.Raise(πTemp010, nil, nil)
					continue
					goto Label20
				Label20:
					// line 883: candidate = candidate[-1]
					πF.SetLineno(883)
					if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp010
					if πE = πg.CheckLocal(πF, µcandidate, "candidate"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µcandidate, πTemp001); πE != nil {
						continue
					}
					µcandidate = πTemp010
					// line 884: name = candidate.split('=',1)[0].split()[-1].strip()
					πF.SetLineno(884)
					if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp010
					πTemp011 = πg.NewInt(0).ToObject()
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("=").ToObject()
					πTemp002[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µcandidate, "candidate"); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, µcandidate, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp016, πE = πTemp015.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp014, πE = πg.GetItem(πF, πTemp016, πTemp011); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp014, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp011.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, πTemp014, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp010, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µname = πTemp010
					// line 885: src = _getimport(mod, name, alias=alias, builtin=builtin)
					πF.SetLineno(885)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					πTemp002[0] = µmod
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002[1] = µname
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"alias", µalias},
						{"builtin", µbuiltin},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_getimport); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp001.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µsrc = πTemp010
					goto Label18
				Label18:
					// line 886: func_vars[name] = src
					πF.SetLineno(886)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsrc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfunc_vars, "func_vars"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp010 = µname
					if πE = πg.SetItem(πF, µfunc_vars, πTemp010, πTemp001); πE != nil {
						continue
					}
					goto Label14
				Label14:
					// line 887: return func_vars
					πF.SetLineno(887)
					if πE = πg.CheckLocal(πF, µfunc_vars, "func_vars"); πE != nil {
						continue
					}
					πR = µfunc_vars
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_closuredimport.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 838: """get import for closured objects; return a dict of 'name' and 'import'"""
			πF.SetLineno(838)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp026}, πg.NewStr("get import for closured objects; return a dict of 'name' and 'import'").ToObject()); πE != nil {
				continue
			}
			if πTemp027, πE = πg.ResolveGlobal(πF, ß_closuredimport); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp027, ß__doc__, πTemp026); πE != nil {
				continue
			}
			// line 890: def _closuredsource(func, alias=''):
			πF.SetLineno(890)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			πTemp004[1] = πg.Param{Name: "alias", Def: ß.ToObject()}
			πTemp026 = πg.NewFunction(πg.NewCode("_closuredsource", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µalias *πg.Object = πArgs[1]; _ = µalias
				var µfreevars *πg.Object = πg.UnboundLocal; _ = µfreevars
				var µfree_vars *πg.Object = πg.UnboundLocal; _ = µfree_vars
				var µfunc_vars *πg.Object = πg.UnboundLocal; _ = µfunc_vars
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
				var µfobj *πg.Object = πg.UnboundLocal; _ = µfobj
				var µsrc *πg.Object = πg.UnboundLocal; _ = µsrc
				var µorg *πg.Object = πg.UnboundLocal; _ = µorg
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 πg.KWArgs
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 891: """get source code for closured objects; return a dict of 'name'
					πF.SetLineno(891)
					// line 897: from .detect import freevars
					πF.SetLineno(897)
					if πTemp002, πE = πg.ImportModule(πF, ".detect"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßfreevars); πE != nil {
						continue
					}
					µfreevars = πTemp003
					// line 898: free_vars = freevars(func)
					πF.SetLineno(898)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πE = πg.CheckLocal(πF, µfreevars, "freevars"); πE != nil {
						continue
					}
					if πTemp001, πE = µfreevars.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfree_vars = πTemp001
					// line 899: func_vars = {}
					πF.SetLineno(899)
					πTemp004 = πg.NewDict()
					πTemp001 = πTemp004.ToObject()
					µfunc_vars = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfree_vars, "free_vars"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfree_vars, ßitems, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
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
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp003); πE != nil {
							continue
						}
						µname = πTemp005
						µobj = πTemp008
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp005, πE = πg.ResolveGlobal(πF, ßisfunction); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label4
					}
					goto Label5
					// line 902: if not isfunction(obj):
					πF.SetLineno(902)
				Label4:
					// line 904: free_vars[name] = getsource(obj, force=True, alias=name)
					πF.SetLineno(904)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"force", πTemp003},
						{"alias", µname},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfree_vars, "free_vars"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp008 = µname
					if πE = πg.SetItem(πF, µfree_vars, πTemp008, πTemp003); πE != nil {
						continue
					}
					// line 905: continue
					πF.SetLineno(905)
					continue
					goto Label5
				Label5:
					// line 907: fobj = free_vars.pop(name)
					πF.SetLineno(907)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002[0] = µname
					if πE = πg.CheckLocal(πF, µfree_vars, "free_vars"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfree_vars, ßpop, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfobj = πTemp005
					// line 908: src = getsource(fobj, alias) # DO NOT include dependencies
					πF.SetLineno(908)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfobj, "fobj"); πE != nil {
						continue
					}
					πTemp002[0] = µfobj
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp002[1] = µalias
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µsrc = πTemp005
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("@").ToObject()
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µsrc, ßlstrip, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp008, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					goto Label7
					// line 910: if not src.lstrip().startswith('@'): #FIXME: 'enclose' in dummy;
					πF.SetLineno(910)
				Label6:
					// line 911: src = importable(fobj,alias=name)#        wrong ref 'name'
					πF.SetLineno(911)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfobj, "fobj"); πE != nil {
						continue
					}
					πTemp002[0] = µfobj
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"alias", µname},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßimportable); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µsrc = πTemp005
					// line 912: org = getsource(func, alias, enclosing=False, lstrip=True)
					πF.SetLineno(912)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp002[1] = µalias
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"enclosing", πTemp003},
						{"lstrip", πTemp005},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µorg = πTemp005
					// line 913: src = (src, org) # undecorated first, then target
					πF.SetLineno(913)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µorg, "org"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µsrc, µorg).ToObject()
					µsrc = πTemp003
					goto Label8
				Label7:
					// line 915: org = getsource(func, enclosing=True, lstrip=False)
					πF.SetLineno(915)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"enclosing", πTemp003},
						{"lstrip", πTemp005},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µorg = πTemp005
					// line 916: src = importable(fobj, alias, source=True) # include dependencies
					πF.SetLineno(916)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfobj, "fobj"); πE != nil {
						continue
					}
					πTemp002[0] = µfobj
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp002[1] = µalias
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"source", πTemp003},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßimportable); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µsrc = πTemp005
					// line 917: src = (org, src) # target first, then decorated
					πF.SetLineno(917)
					if πE = πg.CheckLocal(πF, µorg, "org"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µorg, µsrc).ToObject()
					µsrc = πTemp003
					goto Label8
				Label8:
					// line 918: func_vars[name] = src
					πF.SetLineno(918)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µsrc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfunc_vars, "func_vars"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp005 = µname
					if πE = πg.SetItem(πF, µfunc_vars, πTemp005, πTemp003); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 919: src = ''.join(free_vars.values())
					πF.SetLineno(919)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfree_vars, "free_vars"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfree_vars, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µsrc = πTemp003
					if πE = πg.CheckLocal(πF, µfunc_vars, "func_vars"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µfunc_vars); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					goto Label10
					// line 920: if not func_vars: #FIXME: 'enclose' in dummy; wrong ref 'name'
					πF.SetLineno(920)
				Label9:
					// line 921: org = getsource(func, alias, force=True, enclosing=False, lstrip=True)
					πF.SetLineno(921)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002[0] = µfunc
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp002[1] = µalias
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"force", πTemp001},
						{"enclosing", πTemp003},
						{"lstrip", πTemp005},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µorg = πTemp003
					// line 922: src = (src, org) # variables first, then target
					πF.SetLineno(922)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µorg, "org"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µsrc, µorg).ToObject()
					µsrc = πTemp001
					goto Label11
				Label10:
					// line 924: src = (src, None) # just variables        (better '' instead of None?)
					πF.SetLineno(924)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µsrc, πTemp003).ToObject()
					µsrc = πTemp001
					goto Label11
				Label11:
					// line 925: func_vars[None] = src
					πF.SetLineno(925)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsrc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfunc_vars, "func_vars"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πE = πg.SetItem(πF, µfunc_vars, πTemp003, πTemp001); πE != nil {
						continue
					}
					// line 927: return func_vars
					πF.SetLineno(927)
					if πE = πg.CheckLocal(πF, µfunc_vars, "func_vars"); πE != nil {
						continue
					}
					πR = µfunc_vars
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_closuredsource.ToObject(), πTemp026); πE != nil {
				continue
			}
			// line 891: """get source code for closured objects; return a dict of 'name'
			πF.SetLineno(891)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πg.NewStr("get source code for closured objects; return a dict of 'name'\n    and 'code blocks'").ToObject()); πE != nil {
				continue
			}
			if πTemp028, πE = πg.ResolveGlobal(πF, ß_closuredsource); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp028, ß__doc__, πTemp027); πE != nil {
				continue
			}
			// line 929: def importable(obj, alias='', source=None, builtin=True):
			πF.SetLineno(929)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "alias", Def: ß.ToObject()}
			if πTemp028, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "source", Def: πTemp028}
			if πTemp028, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "builtin", Def: πTemp028}
			πTemp027 = πg.NewFunction(πg.NewCode("importable", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µalias *πg.Object = πArgs[1]; _ = µalias
				var µsource *πg.Object = πArgs[2]; _ = µsource
				var µbuiltin *πg.Object = πArgs[3]; _ = µbuiltin
				var µtried_source *πg.Object = πg.UnboundLocal; _ = µtried_source
				var µtried_import *πg.Object = πg.UnboundLocal; _ = µtried_import
				var µsrc *πg.Object = πg.UnboundLocal; _ = µsrc
				var µ_code_stitcher *πg.Object = πg.UnboundLocal; _ = µ_code_stitcher
				var µ_src *πg.Object = πg.UnboundLocal; _ = µ_src
				var µxxx *πg.Object = πg.UnboundLocal; _ = µxxx
				var µglobalvars *πg.Object = πg.UnboundLocal; _ = µglobalvars
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.BaseException
				_ = πTemp009
				var πTemp010 *πg.Traceback
				_ = πTemp010
				var πTemp011 []πg.Param
				_ = πTemp011
				var πTemp012 []*πg.Object
				_ = πTemp012
				var πTemp013 bool
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 7: goto Label7
					case 8: goto Label8
					case 13: goto Label13
					case 24: goto Label24
					case 29: goto Label29
					case 30: goto Label30
					default: panic("unexpected function state")
					}
					// line 930: """get an importable string (i.e. source code or the import string)
					πF.SetLineno(930)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µsource == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp001 = µbuiltin
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label2
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisbuiltin); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001 = πTemp005
				Label2:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 947: if source is None:
					πF.SetLineno(947)
				Label1:
					// line 948: source = True if isfrommain(obj) else False
					πF.SetLineno(948)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisfrommain); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label5
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					goto Label6
				Label5:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label6:
					µsource = πTemp001
					goto Label4
					// line 949: elif builtin and isbuiltin(obj):
					πF.SetLineno(949)
				Label3:
					// line 950: source = False
					πF.SetLineno(950)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µsource = πTemp001
					goto Label4
				Label4:
					// line 951: tried_source = tried_import = False
					πF.SetLineno(951)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µtried_source = πTemp001
					µtried_import = πTemp001
					// line 952: while True:
					πF.SetLineno(952)
					πF.PushCheckpoint(8)
					πTemp003 = false
				Label7:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label9
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(7)            
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µsource); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					goto Label11
					// line 953: if not source: # we want an import
					πF.SetLineno(953)
				Label10:
					// line 954: try:
					πF.SetLineno(954)
					πF.PushCheckpoint(13)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_isinstance); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label14
					}
					goto Label15
					// line 955: if _isinstance(obj): # for instances, punt to _importable
					πF.SetLineno(955)
				Label14:
					// line 956: return _importable(obj, alias, source=False, builtin=builtin)
					πF.SetLineno(956)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp004[1] = µalias
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"source", πTemp001},
						{"builtin", µbuiltin},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_importable); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp002
					continue
					goto Label15
				Label15:
					// line 957: src = _closuredimport(obj, alias=alias, builtin=builtin)
					πF.SetLineno(957)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"alias", µalias},
						{"builtin", µbuiltin},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_closuredimport); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µsrc = πTemp002
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					πTemp004[0] = µsrc
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label16
					}
					goto Label17
					// line 958: if len(src) == 0:
					πF.SetLineno(958)
				Label16:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("not implemented").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 959: raise NotImplementedError('not implemented')
					πF.SetLineno(959)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label17
				Label17:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					πTemp004[0] = µsrc
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label18
					}
					goto Label19
					// line 960: if len(src) > 1:
					πF.SetLineno(960)
				Label18:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("not implemented").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 961: raise NotImplementedError('not implemented')
					πF.SetLineno(961)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label19
				Label19:
					// line 962: return list(src.values())[0]
					πF.SetLineno(962)
					πTemp001 = πg.NewInt(0).ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µsrc, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp008
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.GetItem(πF, πTemp008, πTemp001); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label12
				Label13:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					goto Label20
					// line 963: except:
					πF.SetLineno(963)
				Label20:
					if πE = πg.CheckLocal(πF, µtried_source, "tried_source"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µtried_source); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label21
					}
					goto Label22
					// line 964: if tried_source: raise
					πF.SetLineno(964)
				Label21:
					// line 964: if tried_source: raise
					πF.SetLineno(964)
					πE = πF.Raise(nil, nil, nil)
					continue
					goto Label22
				Label22:
					// line 965: tried_import = True
					πF.SetLineno(965)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µtried_import = πTemp001
					πF.RestoreExc(nil, nil)
					goto Label12
				Label12:
					goto Label11
				Label11:
					// line 967: try:
					πF.SetLineno(967)
					πF.PushCheckpoint(24)
					// line 968: src = _closuredsource(obj, alias=alias)
					πF.SetLineno(968)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"alias", µalias},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_closuredsource); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µsrc = πTemp002
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					πTemp004[0] = µsrc
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label25
					}
					goto Label26
					// line 969: if len(src) == 0:
					πF.SetLineno(969)
				Label25:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("not implemented").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 970: raise NotImplementedError('not implemented')
					πF.SetLineno(970)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label26
				Label26:
					// line 972: def _code_stitcher(block):
					πF.SetLineno(972)
					πTemp011 = make([]πg.Param, 1)
					πTemp011[0] = πg.Param{Name: "block", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("_code_stitcher", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µblock *πg.Object = πArgs[0]; _ = µblock
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 973: "stitch together the strings in tuple 'block'"
							πF.SetLineno(973)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µblock, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µblock, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µblock, πTemp001); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µblock, πTemp001); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 974: if block[0] and block[-1]: block = '\n'.join(block)
							πF.SetLineno(974)
						Label2:
							// line 974: if block[0] and block[-1]: block = '\n'.join(block)
							πF.SetLineno(974)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp005[0] = µblock
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µblock = πTemp003
							goto Label6
							// line 975: elif block[0]: block = block[0]
							πF.SetLineno(975)
						Label3:
							// line 975: elif block[0]: block = block[0]
							πF.SetLineno(975)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µblock, πTemp001); πE != nil {
								continue
							}
							µblock = πTemp003
							goto Label6
							// line 976: elif block[-1]: block = block[-1]
							πF.SetLineno(976)
						Label4:
							// line 976: elif block[-1]: block = block[-1]
							πF.SetLineno(976)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µblock, πTemp001); πE != nil {
								continue
							}
							µblock = πTemp003
							goto Label6
						Label5:
							// line 977: else: block = ''
							πF.SetLineno(977)
							µblock = ß.ToObject()
							goto Label6
						Label6:
							// line 978: return block
							πF.SetLineno(978)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πR = µblock
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					µ_code_stitcher = πTemp001
					// line 973: "stitch together the strings in tuple 'block'"
					πF.SetLineno(973)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("stitch together the strings in tuple 'block'").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_code_stitcher, "_code_stitcher"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µ_code_stitcher, ß__doc__, πTemp002); πE != nil {
						continue
					}
					// line 980: _src = _code_stitcher(src.pop(None))
					πF.SetLineno(980)
					πTemp004 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp012[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsrc, ßpop, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp004[0] = πTemp005
					if πE = πg.CheckLocal(πF, µ_code_stitcher, "_code_stitcher"); πE != nil {
						continue
					}
					if πTemp002, πE = µ_code_stitcher.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µ_src = πTemp002
					// line 981: _src = [_src] if _src else []
					πF.SetLineno(981)
					if πE = πg.CheckLocal(πF, µ_src, "_src"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µ_src); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label27
					}
					πTemp004 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µ_src, "_src"); πE != nil {
						continue
					}
					πTemp004[0] = µ_src
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp002 = πTemp005
					goto Label28
				Label27:
					πTemp004 = make([]*πg.Object, 0)
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp002 = πTemp005
				Label28:
					µ_src = πTemp002
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µsrc, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp008); πE != nil {
						continue
					}
					πF.PushCheckpoint(30)
					πTemp006 = false
				Label29:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label31
					}
					if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp013 = !isStop
					} else {
						πTemp013 = true
						µxxx = πTemp005
					}
					if πE != nil || !πTemp013 {
						continue
					}
					πF.PushCheckpoint(29)            
					// line 984: xxx = _code_stitcher(xxx)
					πF.SetLineno(984)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µxxx, "xxx"); πE != nil {
						continue
					}
					πTemp004[0] = µxxx
					if πE = πg.CheckLocal(πF, µ_code_stitcher, "_code_stitcher"); πE != nil {
						continue
					}
					if πTemp005, πE = µ_code_stitcher.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µxxx = πTemp005
					if πE = πg.CheckLocal(πF, µxxx, "xxx"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.IsTrue(πF, µxxx); πE != nil {
						continue
					}
					if πTemp013 {
						goto Label32
					}
					goto Label33
					// line 985: if xxx: _src.append(xxx)
					πF.SetLineno(985)
				Label32:
					// line 985: if xxx: _src.append(xxx)
					πF.SetLineno(985)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µxxx, "xxx"); πE != nil {
						continue
					}
					πTemp004[0] = µxxx
					if πE = πg.CheckLocal(πF, µ_src, "_src"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µ_src, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label33
				Label33:
					continue
				Label30:
					if πE != nil || πR != nil {
						continue
					}
				Label31:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_src, "_src"); πE != nil {
						continue
					}
					πTemp004[0] = µ_src
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp006, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label34
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_src, "_src"); πE != nil {
						continue
					}
					πTemp004[0] = µ_src
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.Eq(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label35
					}
					goto Label36
					// line 987: if not len(_src):
					πF.SetLineno(987)
				Label34:
					// line 988: src = ''
					πF.SetLineno(988)
					µsrc = ß.ToObject()
					goto Label37
					// line 989: elif len(_src) == 1:
					πF.SetLineno(989)
				Label35:
					// line 990: src = _src[0]
					πF.SetLineno(990)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µ_src, "_src"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µ_src, πTemp002); πE != nil {
						continue
					}
					µsrc = πTemp005
					goto Label37
				Label36:
					// line 992: src = '\n'.join(_src)
					πF.SetLineno(992)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_src, "_src"); πE != nil {
						continue
					}
					πTemp004[0] = µ_src
					if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µsrc = πTemp005
					goto Label37
				Label37:
					// line 994: from .detect import globalvars
					πF.SetLineno(994)
					if πTemp004, πE = πg.ImportModule(πF, ".detect"); πE != nil {
						continue
					}
					πTemp002 = πTemp004[1]
					if πTemp005, πE = πg.GetAttrImport(πF, πTemp002, ßglobalvars); πE != nil {
						continue
					}
					µglobalvars = πTemp005
					// line 995: obj = globalvars(obj) #XXX: don't worry about alias? recurse? etc?
					πF.SetLineno(995)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µglobalvars, "globalvars"); πE != nil {
						continue
					}
					if πTemp002, πE = µglobalvars.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µobj = πTemp002
					// line 996: obj = list(getsource(_obj,name,force=True) for (name,_obj) in obj.items() if not isbuiltin(_obj))
					πF.SetLineno(996)
					πTemp004 = πF.MakeArgs(1)
					πTemp011 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µname *πg.Object = πg.UnboundLocal; _ = µname
						var µ_obj *πg.Object = πg.UnboundLocal; _ = µ_obj
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 πg.KWArgs
						_ = πTemp008
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
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µobj, ßitems, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp004 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp004 {
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
									πTemp005 = !isStop
								} else {
									πTemp005 = true
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
										continue
									}
									µname = πTemp003
									µ_obj = πTemp006
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp007 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µ_obj, "_obj"); πE != nil {
									continue
								}
								πTemp007[0] = µ_obj
								if πTemp003, πE = πg.ResolveGlobal(πF, ßisbuiltin); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								πTemp002 = πg.GetBool(!πTemp005).ToObject()
								if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label4
								}
								goto Label5
								// line 996: obj = list(getsource(_obj,name,force=True) for (name,_obj) in obj.items() if not isbuiltin(_obj))
								πF.SetLineno(996)
							Label4:
								// line 996: obj = list(getsource(_obj,name,force=True) for (name,_obj) in obj.items() if not isbuiltin(_obj))
								πF.SetLineno(996)
								πTemp007 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µ_obj, "_obj"); πE != nil {
									continue
								}
								πTemp007[0] = µ_obj
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								πTemp007[1] = µname
								if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πTemp008 = πg.KWArgs{
									{"force", πTemp002},
								}
								if πTemp002, πE = πg.ResolveGlobal(πF, ßgetsource); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp007, πTemp008); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πF.PushCheckpoint(6)
								return πTemp003, nil
							Label6:
								πTemp002 = πSent
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
					if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µobj = πTemp008
					// line 997: obj = '\n'.join(obj) if obj else ''
					πF.SetLineno(997)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µobj); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label38
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp008, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp005 = πTemp014
					goto Label39
				Label38:
					πTemp005 = ß.ToObject()
				Label39:
					µobj = πTemp005
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µobj); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label40
					}
					goto Label41
					// line 999: if not obj: return src
					πF.SetLineno(999)
				Label40:
					// line 999: if not obj: return src
					πF.SetLineno(999)
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					πR = µsrc
					continue
					goto Label41
				Label41:
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µsrc); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label42
					}
					goto Label43
					// line 1000: if not src: return obj
					πF.SetLineno(1000)
				Label42:
					// line 1000: if not src: return obj
					πF.SetLineno(1000)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πR = µobj
					continue
					goto Label43
				Label43:
					// line 1001: return obj + src
					πF.SetLineno(1001)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µobj, µsrc); πE != nil {
						continue
					}
					πR = πTemp005
					continue
					πF.PopCheckpoint()
					goto Label23
				Label24:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					goto Label44
					// line 1002: except:
					πF.SetLineno(1002)
				Label44:
					if πE = πg.CheckLocal(πF, µtried_import, "tried_import"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µtried_import); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label45
					}
					goto Label46
					// line 1003: if tried_import: raise
					πF.SetLineno(1003)
				Label45:
					// line 1003: if tried_import: raise
					πF.SetLineno(1003)
					πE = πF.Raise(nil, nil, nil)
					continue
					goto Label46
				Label46:
					// line 1004: tried_source = True
					πF.SetLineno(1004)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µtried_source = πTemp005
					// line 1005: source = not source
					πF.SetLineno(1005)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µsource); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp006).ToObject()
					µsource = πTemp005
					πF.RestoreExc(nil, nil)
					goto Label23
				Label23:
					continue
				Label8:
					if πE != nil || πR != nil {
						continue
					}
				Label9:
					// line 1007: return
					πF.SetLineno(1007)
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
			if πE = πF.Globals().SetItem(πF, ßimportable.ToObject(), πTemp027); πE != nil {
				continue
			}
			// line 930: """get an importable string (i.e. source code or the import string)
			πF.SetLineno(930)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πg.NewStr("get an importable string (i.e. source code or the import string)\n    for the given object, including any required objects from the enclosing\n    and global scope\n\n    This function will attempt to discover the name of the object, or the repr\n    of the object, or the source code for the object. To attempt to force\n    discovery of the source code, use source=True, to attempt to force the\n    use of an import, use source=False; otherwise an import will be sought\n    for objects not defined in __main__. The intent is to build a string\n    that can be imported from a python file.\n\n    obj is the object to inspect. If alias is provided, then rename the\n    object with the given alias. If builtin=True, then force an import for\n    builtins where possible.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp029, πE = πg.ResolveGlobal(πF, ßimportable); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp029, ß__doc__, πTemp028); πE != nil {
				continue
			}
			// line 1011: def getimportable(obj, alias='', byname=True, explicit=False):
			πF.SetLineno(1011)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "alias", Def: ß.ToObject()}
			if πTemp029, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "byname", Def: πTemp029}
			if πTemp029, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "explicit", Def: πTemp029}
			πTemp028 = πg.NewFunction(πg.NewCode("getimportable", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µalias *πg.Object = πArgs[1]; _ = µalias
				var µbyname *πg.Object = πArgs[2]; _ = µbyname
				var µexplicit *πg.Object = πArgs[3]; _ = µexplicit
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 πg.KWArgs
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1012: return importable(obj,alias,source=(not byname),builtin=explicit)
					πF.SetLineno(1012)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp001[1] = µalias
					if πE = πg.CheckLocal(πF, µbyname, "byname"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µbyname); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp003).ToObject()
					if πE = πg.CheckLocal(πF, µexplicit, "explicit"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"source", πTemp002},
						{"builtin", µexplicit},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßimportable); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp005
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßgetimportable.ToObject(), πTemp028); πE != nil {
				continue
			}
			// line 1014: def likely_import(obj, passive=False, explicit=False):
			πF.SetLineno(1014)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "passive", Def: πTemp030}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "explicit", Def: πTemp030}
			πTemp029 = πg.NewFunction(πg.NewCode("likely_import", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µpassive *πg.Object = πArgs[1]; _ = µpassive
				var µexplicit *πg.Object = πArgs[2]; _ = µexplicit
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 πg.KWArgs
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1015: return getimport(obj, verify=(not passive), builtin=explicit)
					πF.SetLineno(1015)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πE = πg.CheckLocal(πF, µpassive, "passive"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µpassive); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp003).ToObject()
					if πE = πg.CheckLocal(πF, µexplicit, "explicit"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"verify", πTemp002},
						{"builtin", µexplicit},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetimport); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp005
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßlikely_import.ToObject(), πTemp029); πE != nil {
				continue
			}
			// line 1016: def _likely_import(first, last, passive=False, explicit=True):
			πF.SetLineno(1016)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "first", Def: nil}
			πTemp004[1] = πg.Param{Name: "last", Def: nil}
			if πTemp031, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "passive", Def: πTemp031}
			if πTemp031, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "explicit", Def: πTemp031}
			πTemp030 = πg.NewFunction(πg.NewCode("_likely_import", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/source.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfirst *πg.Object = πArgs[0]; _ = µfirst
				var µlast *πg.Object = πArgs[1]; _ = µlast
				var µpassive *πg.Object = πArgs[2]; _ = µpassive
				var µexplicit *πg.Object = πArgs[3]; _ = µexplicit
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 πg.KWArgs
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1017: return _getimport(first, last, verify=(not passive), builtin=explicit)
					πF.SetLineno(1017)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
						continue
					}
					πTemp001[0] = µfirst
					if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
						continue
					}
					πTemp001[1] = µlast
					if πE = πg.CheckLocal(πF, µpassive, "passive"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µpassive); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp003).ToObject()
					if πE = πg.CheckLocal(πF, µexplicit, "explicit"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"verify", πTemp002},
						{"builtin", µexplicit},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_getimport); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp005
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_likely_import.ToObject(), πTemp030); πE != nil {
				continue
			}
			// line 1018: _get_name = getname
			πF.SetLineno(1018)
			if πTemp031, πE = πg.ResolveGlobal(πF, ßgetname); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_get_name.ToObject(), πTemp031); πE != nil {
				continue
			}
			// line 1019: getblocks_from_history = getblocks
			πF.SetLineno(1019)
			if πTemp031, πE = πg.ResolveGlobal(πF, ßgetblocks); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetblocks_from_history.ToObject(), πTemp031); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("source", Code)
}

