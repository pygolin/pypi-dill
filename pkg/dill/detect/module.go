package detect
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/dill"
	// _ "github.com/pygolin/stdlib/temp"
	// _ "github.com/pygolin/stdlib/pointers"
	// _ "github.com/pygolin/stdlib/inspect"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/source"
	// _ "github.com/pygolin/stdlib/_dill"
	// _ "github.com/pygolin/stdlib/dis"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßNone := πg.InternStr("None")
		ßPY3 := πg.InternStr("PY3")
		ßTrue := πg.InternStr("True")
		ß_GLOBAL := πg.InternStr("_GLOBAL")
		ß__all__ := πg.InternStr("__all__")
		ß__builtin__ := πg.InternStr("__builtin__")
		ß__closure__ := πg.InternStr("__closure__")
		ß__code__ := πg.InternStr("__code__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__func__ := πg.InternStr("__func__")
		ß__globals__ := πg.InternStr("__globals__")
		ß__import__ := πg.InternStr("__import__")
		ß__iter__ := πg.InternStr("__iter__")
		ß_trace := πg.InternStr("_trace")
		ßadd := πg.InternStr("add")
		ßappend := πg.InternStr("append")
		ßat := πg.InternStr("at")
		ßbaditems := πg.InternStr("baditems")
		ßbadobjects := πg.InternStr("badobjects")
		ßbadtypes := πg.InternStr("badtypes")
		ßbuiltins := πg.InternStr("builtins")
		ßcapture := πg.InternStr("capture")
		ßcell_contents := πg.InternStr("cell_contents")
		ßchildren := πg.InternStr("children")
		ßco_cellvars := πg.InternStr("co_cellvars")
		ßco_code := πg.InternStr("co_code")
		ßco_consts := πg.InternStr("co_consts")
		ßco_freevars := πg.InternStr("co_freevars")
		ßco_name := πg.InternStr("co_name")
		ßco_names := πg.InternStr("co_names")
		ßco_varnames := πg.InternStr("co_varnames")
		ßcode := πg.InternStr("code")
		ßcopy := πg.InternStr("copy")
		ßdict := πg.InternStr("dict")
		ßdir := πg.InternStr("dir")
		ßdis := πg.InternStr("dis")
		ßerrors := πg.InternStr("errors")
		ßexc_info := πg.InternStr("exc_info")
		ßf_code := πg.InternStr("f_code")
		ßfreevars := πg.InternStr("freevars")
		ßfunc_closure := πg.InternStr("func_closure")
		ßfunc_code := πg.InternStr("func_code")
		ßfunc_globals := πg.InternStr("func_globals")
		ßget := πg.InternStr("get")
		ßget_referrers := πg.InternStr("get_referrers")
		ßgetattr := πg.InternStr("getattr")
		ßgetmodule := πg.InternStr("getmodule")
		ßgetname := πg.InternStr("getname")
		ßgetsourcelines := πg.InternStr("getsourcelines")
		ßgetvalue := πg.InternStr("getvalue")
		ßglobalvars := πg.InternStr("globalvars")
		ßhasattr := πg.InternStr("hasattr")
		ßim_func := πg.InternStr("im_func")
		ßiscode := πg.InternStr("iscode")
		ßisframe := πg.InternStr("isframe")
		ßisfunction := πg.InternStr("isfunction")
		ßismethod := πg.InternStr("ismethod")
		ßistraceback := πg.InternStr("istraceback")
		ßitems := πg.InternStr("items")
		ßiteritems := πg.InternStr("iteritems")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlist := πg.InternStr("list")
		ßnestedcode := πg.InternStr("nestedcode")
		ßnestedglobals := πg.InternStr("nestedglobals")
		ßoutermost := πg.InternStr("outermost")
		ßparent := πg.InternStr("parent")
		ßparents := πg.InternStr("parents")
		ßpickles := πg.InternStr("pickles")
		ßreference := πg.InternStr("reference")
		ßreferredglobals := πg.InternStr("referredglobals")
		ßreferrednested := πg.InternStr("referrednested")
		ßset := πg.InternStr("set")
		ßsplit := πg.InternStr("split")
		ßsplitlines := πg.InternStr("splitlines")
		ßstdout := πg.InternStr("stdout")
		ßsum := πg.InternStr("sum")
		ßtb_frame := πg.InternStr("tb_frame")
		ßtrace := πg.InternStr("trace")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßupdate := πg.InternStr("update")
		ßvalues := πg.InternStr("values")
		ßvarnames := πg.InternStr("varnames")
		ßvars := πg.InternStr("vars")
		ßzip := πg.InternStr("zip")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 8: """
			πF.SetLineno(8)
			// line 8: """
			πF.SetLineno(8)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nMethods for detecting objects leading to pickling failures.\n").ToObject()); πE != nil {
				continue
			}
			// line 12: import dis
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "dis"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdis.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: from inspect import ismethod, isfunction, istraceback, isframe, iscode
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "inspect"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßismethod); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßismethod.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßisfunction); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßisfunction.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßistraceback); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßistraceback.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßisframe); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßisframe.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßiscode); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßiscode.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 14: from .pointers import parent, reference, at, parents, children
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, ".pointers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßparent); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßparent.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßreference); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreference.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßat); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßat.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßparents); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßparents.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßchildren); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßchildren.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 16: from ._dill import _trace as trace
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "._dill"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ß_trace); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtrace.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 17: from ._dill import PY3
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "._dill"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßPY3); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPY3.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 19: __all__ = ['baditems','badobjects','badtypes','code','errors','freevars',
			πF.SetLineno(19)
			πTemp002 = make([]*πg.Object, 15)
			πTemp002[0] = ßbaditems.ToObject()
			πTemp002[1] = ßbadobjects.ToObject()
			πTemp002[2] = ßbadtypes.ToObject()
			πTemp002[3] = ßcode.ToObject()
			πTemp002[4] = ßerrors.ToObject()
			πTemp002[5] = ßfreevars.ToObject()
			πTemp002[6] = ßgetmodule.ToObject()
			πTemp002[7] = ßglobalvars.ToObject()
			πTemp002[8] = ßnestedcode.ToObject()
			πTemp002[9] = ßnestedglobals.ToObject()
			πTemp002[10] = ßoutermost.ToObject()
			πTemp002[11] = ßreferredglobals.ToObject()
			πTemp002[12] = ßreferrednested.ToObject()
			πTemp002[13] = ßtrace.ToObject()
			πTemp002[14] = ßvarnames.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: def getmodule(object, _filename=None, force=False):
			πF.SetLineno(23)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "_filename", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "force", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("getmodule", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µ_filename *πg.Object = πArgs[1]; _ = µ_filename
				var µforce *πg.Object = πArgs[2]; _ = µforce
				var µgetmod *πg.Object = πg.UnboundLocal; _ = µgetmod
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var µbuiltins *πg.Object = πg.UnboundLocal; _ = µbuiltins
				var µgetname *πg.Object = πg.UnboundLocal; _ = µgetname
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 πg.KWArgs
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
					// line 24: """get the module of the object"""
					πF.SetLineno(24)
					// line 25: from inspect import getmodule as getmod
					πF.SetLineno(25)
					if πTemp002, πE = πg.ImportModule(πF, "inspect"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßgetmodule); πE != nil {
						continue
					}
					µgetmod = πTemp003
					// line 26: module = getmod(object, _filename)
					πF.SetLineno(26)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp002[0] = µobject
					if πE = πg.CheckLocal(πF, µ_filename, "_filename"); πE != nil {
						continue
					}
					πTemp002[1] = µ_filename
					if πE = πg.CheckLocal(πF, µgetmod, "getmod"); πE != nil {
						continue
					}
					if πTemp001, πE = µgetmod.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmodule = πTemp001
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001 = µmodule
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µforce); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp005).ToObject()
					πTemp001 = πTemp003
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 27: if module or not force: return module
					πF.SetLineno(27)
				Label2:
					// line 27: if module or not force: return module
					πF.SetLineno(27)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πR = µmodule
					continue
					goto Label3
				Label3:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 28: if PY3: builtins = 'builtins'
					πF.SetLineno(28)
				Label4:
					// line 28: if PY3: builtins = 'builtins'
					πF.SetLineno(28)
					µbuiltins = ßbuiltins.ToObject()
					goto Label6
				Label5:
					// line 29: else: builtins = '__builtin__'
					πF.SetLineno(29)
					µbuiltins = ß__builtin__.ToObject()
					goto Label6
				Label6:
					// line 30: builtins = __import__(builtins)
					πF.SetLineno(30)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbuiltins, "builtins"); πE != nil {
						continue
					}
					πTemp002[0] = µbuiltins
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µbuiltins = πTemp003
					// line 31: from .source import getname
					πF.SetLineno(31)
					if πTemp002, πE = πg.ImportModule(πF, ".source"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßgetname); πE != nil {
						continue
					}
					µgetname = πTemp003
					// line 32: name = getname(object, force=True)
					πF.SetLineno(32)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp002[0] = µobject
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"force", πTemp001},
					}
					if πE = πg.CheckLocal(πF, µgetname, "getname"); πE != nil {
						continue
					}
					if πTemp001, πE = µgetname.Call(πF, πTemp002, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µname = πTemp001
					// line 33: return builtins if name in vars(builtins).keys() else None
					πF.SetLineno(33)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbuiltins, "builtins"); πE != nil {
						continue
					}
					πTemp002[0] = µbuiltins
					if πTemp007, πE = πg.ResolveGlobal(πF, ßvars); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp007, πE = πg.GetAttr(πF, πTemp008, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp008, µname); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µbuiltins, "builtins"); πE != nil {
						continue
					}
					πTemp001 = µbuiltins
					goto Label8
				Label7:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label8:
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
			if πE = πF.Globals().SetItem(πF, ßgetmodule.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 24: """get the module of the object"""
			πF.SetLineno(24)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("get the module of the object").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 35: def outermost(func): # is analogous to getsource(func,enclosing=True)
			πF.SetLineno(35)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("outermost", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µ_globals *πg.Object = πg.UnboundLocal; _ = µ_globals
				var µgetsourcelines *πg.Object = πg.UnboundLocal; _ = µgetsourcelines
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µlnum *πg.Object = πg.UnboundLocal; _ = µlnum
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
				var µ_locals *πg.Object = πg.UnboundLocal; _ = µ_locals
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
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
				var πTemp006 *πg.Dict
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 []πg.Param
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 17: goto Label17
					case 19: goto Label19
					case 20: goto Label20
					case 23: goto Label23
					default: panic("unexpected function state")
					}
					// line 36: """get outermost enclosing object (i.e. the outer function in a closure)
					πF.SetLineno(36)
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
					// line 40: if PY3:
					πF.SetLineno(40)
				Label1:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßismethod); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisfunction); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					goto Label6
					// line 41: if ismethod(func):
					πF.SetLineno(41)
				Label4:
					// line 42: _globals = func.__func__.__globals__ or {}
					πF.SetLineno(42)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfunc, ß__func__, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__globals__, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					πTemp006 = πg.NewDict()
					πTemp004 = πTemp006.ToObject()
					πTemp001 = πTemp004
				Label8:
					µ_globals = πTemp001
					goto Label7
					// line 43: elif isfunction(func):
					πF.SetLineno(43)
				Label5:
					// line 44: _globals = func.__globals__ or {}
					πF.SetLineno(44)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfunc, ß__globals__, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					πTemp006 = πg.NewDict()
					πTemp004 = πTemp006.ToObject()
					πTemp001 = πTemp004
				Label9:
					µ_globals = πTemp001
					goto Label7
				Label6:
					// line 46: return #XXX: or raise? no matches
					πF.SetLineno(46)
					πR = πg.None
					continue
					goto Label7
				Label7:
					// line 47: _globals = _globals.items()
					πF.SetLineno(47)
					if πE = πg.CheckLocal(πF, µ_globals, "_globals"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µ_globals, ßitems, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µ_globals = πTemp004
					goto Label3
				Label2:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßismethod); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label10
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisfunction); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label11
					}
					goto Label12
					// line 49: if ismethod(func):
					πF.SetLineno(49)
				Label10:
					// line 50: _globals = func.im_func.func_globals or {}
					πF.SetLineno(50)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfunc, ßim_func, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßfunc_globals, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label14
					}
					πTemp006 = πg.NewDict()
					πTemp004 = πTemp006.ToObject()
					πTemp001 = πTemp004
				Label14:
					µ_globals = πTemp001
					goto Label13
					// line 51: elif isfunction(func):
					πF.SetLineno(51)
				Label11:
					// line 52: _globals = func.func_globals or {}
					πF.SetLineno(52)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfunc, ßfunc_globals, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label15
					}
					πTemp006 = πg.NewDict()
					πTemp004 = πTemp006.ToObject()
					πTemp001 = πTemp004
				Label15:
					µ_globals = πTemp001
					goto Label13
				Label12:
					// line 54: return #XXX: or raise? no matches
					πF.SetLineno(54)
					πR = πg.None
					continue
					goto Label13
				Label13:
					// line 55: _globals = _globals.iteritems()
					πF.SetLineno(55)
					if πE = πg.CheckLocal(πF, µ_globals, "_globals"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µ_globals, ßiteritems, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µ_globals = πTemp004
					goto Label3
				Label3:
					// line 57: from .source import getsourcelines
					πF.SetLineno(57)
					if πTemp003, πE = πg.ImportModule(πF, ".source"); πE != nil {
						continue
					}
					πTemp001 = πTemp003[1]
					if πTemp004, πE = πg.GetAttrImport(πF, πTemp001, ßgetsourcelines); πE != nil {
						continue
					}
					µgetsourcelines = πTemp004
					// line 58: try: lines,lnum = getsourcelines(func, enclosing=True)
					πF.SetLineno(58)
					πF.PushCheckpoint(17)
					// line 58: try: lines,lnum = getsourcelines(func, enclosing=True)
					πF.SetLineno(58)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"enclosing", πTemp001},
					}
					if πE = πg.CheckLocal(πF, µgetsourcelines, "getsourcelines"); πE != nil {
						continue
					}
					if πTemp001, πE = µgetsourcelines.Call(πF, πTemp003, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
						continue
					}
					µlines = πTemp004
					µlnum = πTemp005
					πF.PopCheckpoint()
					goto Label16
				Label17:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					goto Label18
					// line 59: except: #TypeError, IOError
					πF.SetLineno(59)
				Label18:
					// line 60: lines,lnum = [],None
					πF.SetLineno(60)
					πTemp003 = make([]*πg.Object, 0)
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
						continue
					}
					µlines = πTemp004
					µlnum = πTemp005
					πF.RestoreExc(nil, nil)
					goto Label16
				Label16:
					// line 61: code = ''.join(lines)
					πF.SetLineno(61)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp003[0] = µlines
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µcode = πTemp004
					// line 63: _locals = ((name,obj) for (name,obj) in _globals if name in code)
					πF.SetLineno(63)
					πTemp010 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µname *πg.Object = πg.UnboundLocal; _ = µname
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
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
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µ_globals, "_globals"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µ_globals); πE != nil {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp004); πE != nil {
										continue
									}
									µname = πTemp005
									µobj = πTemp006
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Contains(πF, µcode, µname); πE != nil {
									continue
								}
								πTemp004 = πg.GetBool(πTemp003).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								goto Label5
								// line 63: _locals = ((name,obj) for (name,obj) in _globals if name in code)
								πF.SetLineno(63)
							Label4:
								// line 63: _locals = ((name,obj) for (name,obj) in _globals if name in code)
								πF.SetLineno(63)
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp004 = πg.NewTuple2(µname, µobj).ToObject()
								πF.PushCheckpoint(6)
								return πTemp004, nil
							Label6:
								πTemp005 = πSent
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
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µ_locals = πTemp004
					if πE = πg.CheckLocal(πF, µ_locals, "_locals"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Iter(πF, µ_locals); πE != nil {
						continue
					}
					πF.PushCheckpoint(20)
					πTemp002 = false
				Label19:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label21
					}
					if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp011 = !isStop
					} else {
						πTemp011 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp012}, πg.TieTarget{Target: &πTemp013}}}, πTemp005); πE != nil {
							continue
						}
						µname = πTemp012
						µobj = πTemp013
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(19)            
					// line 66: try:
					πF.SetLineno(66)
					πF.PushCheckpoint(23)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πE = πg.CheckLocal(πF, µgetsourcelines, "getsourcelines"); πE != nil {
						continue
					}
					if πTemp012, πE = µgetsourcelines.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlnum, "lnum"); πE != nil {
						continue
					}
					πTemp013 = πg.NewTuple2(µlines, µlnum).ToObject()
					if πTemp005, πE = πg.Eq(πF, πTemp012, πTemp013); πE != nil {
						continue
					}
					if πTemp011, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label24
					}
					goto Label25
					// line 67: if getsourcelines(obj) == (lines,lnum): return obj
					πF.SetLineno(67)
				Label24:
					// line 67: if getsourcelines(obj) == (lines,lnum): return obj
					πF.SetLineno(67)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πR = µobj
					continue
					goto Label25
				Label25:
					πF.PopCheckpoint()
					goto Label22
				Label23:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					goto Label26
					// line 68: except: #TypeError, IOError
					πF.SetLineno(68)
				Label26:
					// line 69: pass
					πF.SetLineno(69)
					πF.RestoreExc(nil, nil)
					goto Label22
				Label22:
					continue
				Label20:
					if πE != nil || πR != nil {
						continue
					}
				Label21:
					// line 70: return #XXX: or raise? no matches
					πF.SetLineno(70)
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
			if πE = πF.Globals().SetItem(πF, ßoutermost.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 36: """get outermost enclosing object (i.e. the outer function in a closure)
			πF.SetLineno(36)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("get outermost enclosing object (i.e. the outer function in a closure)\n\n    NOTE: this is the object-equivalent of getsource(func, enclosing=True)\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßoutermost); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
				continue
			}
			// line 72: def nestedcode(func, recurse=True): #XXX: or return dict of {co_name: co} ?
			πF.SetLineno(72)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "recurse", Def: πTemp006}
			πTemp005 = πg.NewFunction(πg.NewCode("nestedcode", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µrecurse *πg.Object = πArgs[1]; _ = µrecurse
				var µnested *πg.Object = πg.UnboundLocal; _ = µnested
				var µco *πg.Object = πg.UnboundLocal; _ = µco
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
				var πTemp006 bool
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
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 73: """get the code objects for any nested functions (e.g. in a closure)"""
					πF.SetLineno(73)
					// line 74: func = code(func)
					πF.SetLineno(74)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp001[0] = µfunc
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcode); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfunc = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp001[0] = µfunc
					if πTemp003, πE = πg.ResolveGlobal(πF, ßiscode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 75: if not iscode(func): return [] #XXX: or raise? no matches
					πF.SetLineno(75)
				Label1:
					// line 75: if not iscode(func): return [] #XXX: or raise? no matches
					πF.SetLineno(75)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 76: nested = set()
					πF.SetLineno(76)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µnested = πTemp003
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfunc, ßco_consts, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						µco = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(3)            
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µco == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					goto Label7
					// line 78: if co is None: continue
					πF.SetLineno(78)
				Label6:
					// line 78: if co is None: continue
					πF.SetLineno(78)
					continue
					goto Label7
				Label7:
					// line 79: co = code(co)
					πF.SetLineno(79)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					πTemp001[0] = µco
					if πTemp003, πE = πg.ResolveGlobal(πF, ßcode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µco = πTemp004
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µco); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label8
					}
					goto Label9
					// line 80: if co:
					πF.SetLineno(80)
				Label8:
					// line 81: nested.add(co)
					πF.SetLineno(81)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					πTemp001[0] = µco
					if πE = πg.CheckLocal(πF, µnested, "nested"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µnested, ßadd, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µrecurse); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					goto Label11
					// line 82: if recurse: nested |= set(nestedcode(co, recurse=True))
					πF.SetLineno(82)
				Label10:
					// line 82: if recurse: nested |= set(nestedcode(co, recurse=True))
					πF.SetLineno(82)
					if πE = πg.CheckLocal(πF, µnested, "nested"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					πTemp007[0] = µco
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp008 = πg.KWArgs{
						{"recurse", πTemp003},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßnestedcode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp007, πTemp008); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.IOr(πF, µnested, πTemp004); πE != nil {
						continue
					}
					µnested = πTemp003
					goto Label11
				Label11:
					goto Label9
				Label9:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 83: return list(nested)
					πF.SetLineno(83)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnested, "nested"); πE != nil {
						continue
					}
					πTemp001[0] = µnested
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßnestedcode.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 73: """get the code objects for any nested functions (e.g. in a closure)"""
			πF.SetLineno(73)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("get the code objects for any nested functions (e.g. in a closure)").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßnestedcode); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
			// line 85: def code(func):
			πF.SetLineno(85)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("code", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µim_func *πg.Object = πg.UnboundLocal; _ = µim_func
				var µfunc_code *πg.Object = πg.UnboundLocal; _ = µfunc_code
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 86: '''get the code object for the given function or method
					πF.SetLineno(86)
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
					// line 90: if PY3:
					πF.SetLineno(90)
				Label1:
					// line 91: im_func = '__func__'
					πF.SetLineno(91)
					µim_func = ß__func__.ToObject()
					// line 92: func_code = '__code__'
					πF.SetLineno(92)
					µfunc_code = ß__code__.ToObject()
					goto Label3
				Label2:
					// line 94: im_func = 'im_func'
					πF.SetLineno(94)
					µim_func = ßim_func.ToObject()
					// line 95: func_code = 'func_code'
					πF.SetLineno(95)
					µfunc_code = ßfunc_code.ToObject()
					goto Label3
				Label3:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßismethod); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 96: if ismethod(func): func = getattr(func, im_func)
					πF.SetLineno(96)
				Label4:
					// line 96: if ismethod(func): func = getattr(func, im_func)
					πF.SetLineno(96)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πE = πg.CheckLocal(πF, µim_func, "im_func"); πE != nil {
						continue
					}
					πTemp003[1] = µim_func
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µfunc = πTemp004
					goto Label5
				Label5:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisfunction); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					goto Label7
					// line 97: if isfunction(func): func = getattr(func, func_code)
					πF.SetLineno(97)
				Label6:
					// line 97: if isfunction(func): func = getattr(func, func_code)
					πF.SetLineno(97)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πE = πg.CheckLocal(πF, µfunc_code, "func_code"); πE != nil {
						continue
					}
					πTemp003[1] = µfunc_code
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µfunc = πTemp004
					goto Label7
				Label7:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßistraceback); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					goto Label9
					// line 98: if istraceback(func): func = func.tb_frame
					πF.SetLineno(98)
				Label8:
					// line 98: if istraceback(func): func = func.tb_frame
					πF.SetLineno(98)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfunc, ßtb_frame, nil); πE != nil {
						continue
					}
					µfunc = πTemp001
					goto Label9
				Label9:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisframe); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label10
					}
					goto Label11
					// line 99: if isframe(func): func = func.f_code
					πF.SetLineno(99)
				Label10:
					// line 99: if isframe(func): func = func.f_code
					πF.SetLineno(99)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfunc, ßf_code, nil); πE != nil {
						continue
					}
					µfunc = πTemp001
					goto Label11
				Label11:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßiscode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					goto Label13
					// line 100: if iscode(func): return func
					πF.SetLineno(100)
				Label12:
					// line 100: if iscode(func): return func
					πF.SetLineno(100)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πR = µfunc
					continue
					goto Label13
				Label13:
					// line 101: return
					πF.SetLineno(101)
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
			if πE = πF.Globals().SetItem(πF, ßcode.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 86: '''get the code object for the given function or method
			πF.SetLineno(86)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("get the code object for the given function or method\n\n    NOTE: use dill.source.getsource(CODEOBJ) to get the source code\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßcode); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
				continue
			}
			// line 104: def referrednested(func, recurse=True): #XXX: return dict of {__name__: obj} ?
			πF.SetLineno(104)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "recurse", Def: πTemp008}
			πTemp007 = πg.NewFunction(πg.NewCode("referrednested", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µrecurse *πg.Object = πArgs[1]; _ = µrecurse
				var µatt1 *πg.Object = πg.UnboundLocal; _ = µatt1
				var µatt0 *πg.Object = πg.UnboundLocal; _ = µatt0
				var µgc *πg.Object = πg.UnboundLocal; _ = µgc
				var µfuncs *πg.Object = πg.UnboundLocal; _ = µfuncs
				var µco *πg.Object = πg.UnboundLocal; _ = µco
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
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
				var πTemp006 bool
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
					case 8: goto Label8
					case 4: goto Label4
					case 5: goto Label5
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 105: """get functions defined inside of func (e.g. inner functions in a closure)
					πF.SetLineno(105)
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
					// line 112: if PY3:
					πF.SetLineno(112)
				Label1:
					// line 113: att1 = '__code__'
					πF.SetLineno(113)
					µatt1 = ß__code__.ToObject()
					// line 114: att0 = '__func__'
					πF.SetLineno(114)
					µatt0 = ß__func__.ToObject()
					goto Label3
				Label2:
					// line 116: att1 = 'func_code' # functions
					πF.SetLineno(116)
					µatt1 = ßfunc_code.ToObject()
					// line 117: att0 = 'im_func'   # methods
					πF.SetLineno(117)
					µatt0 = ßim_func.ToObject()
					goto Label3
				Label3:
					// line 119: import gc
					πF.SetLineno(119)
					if πTemp003, πE = πg.ImportModule(πF, "gc"); πE != nil {
						continue
					}
					πTemp001 = πTemp003[0]
					µgc = πTemp001
					// line 120: funcs = set()
					πF.SetLineno(120)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfuncs = πTemp004
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
						continue
					}
					πTemp003[1] = µrecurse
					if πTemp004, πE = πg.ResolveGlobal(πF, ßnestedcode); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp002 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label6
					}
					if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
						µco = πTemp004
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(4)            
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					πTemp003[0] = µco
					if πE = πg.CheckLocal(πF, µgc, "gc"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µgc, ßget_referrers, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.Iter(πF, πTemp007); πE != nil {
						continue
					}
					πF.PushCheckpoint(8)
					πTemp006 = false
				Label7:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label9
					}
					if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
						µobj = πTemp005
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(7)            
					// line 126: _ = getattr(obj, att0, None) # ismethod
					πF.SetLineno(126)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πE = πg.CheckLocal(πF, µatt0, "att0"); πE != nil {
						continue
					}
					πTemp003[1] = µatt0
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µ_ = πTemp007
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					πTemp003[0] = µ_
					if πE = πg.CheckLocal(πF, µatt1, "att1"); πE != nil {
						continue
					}
					πTemp003[1] = µatt1
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp009 == µco).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label10
					}
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πE = πg.CheckLocal(πF, µatt1, "att1"); πE != nil {
						continue
					}
					πTemp003[1] = µatt1
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp009 == µco).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label11
					}
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					πTemp003[1] = ßf_code.ToObject()
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp009 == µco).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label12
					}
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					πTemp003[1] = ßco_code.ToObject()
					if πTemp007, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp005 = πTemp009
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(µobj == µco).ToObject()
					πTemp005 = πTemp007
				Label13:
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label14
					}
					goto Label15
					// line 127: if getattr(_, att1, None) is co: funcs.add(obj)
					πF.SetLineno(127)
				Label10:
					// line 127: if getattr(_, att1, None) is co: funcs.add(obj)
					πF.SetLineno(127)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πE = πg.CheckLocal(πF, µfuncs, "funcs"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µfuncs, ßadd, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label15
					// line 129: elif getattr(obj, att1, None) is co: funcs.add(obj)
					πF.SetLineno(129)
				Label11:
					// line 129: elif getattr(obj, att1, None) is co: funcs.add(obj)
					πF.SetLineno(129)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πE = πg.CheckLocal(πF, µfuncs, "funcs"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µfuncs, ßadd, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label15
					// line 131: elif getattr(obj, 'f_code', None) is co: funcs.add(obj)
					πF.SetLineno(131)
				Label12:
					// line 131: elif getattr(obj, 'f_code', None) is co: funcs.add(obj)
					πF.SetLineno(131)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πE = πg.CheckLocal(πF, µfuncs, "funcs"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µfuncs, ßadd, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label15
					// line 133: elif hasattr(obj, 'co_code') and obj is co: funcs.add(obj)
					πF.SetLineno(133)
				Label14:
					// line 133: elif hasattr(obj, 'co_code') and obj is co: funcs.add(obj)
					πF.SetLineno(133)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πE = πg.CheckLocal(πF, µfuncs, "funcs"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µfuncs, ßadd, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label15
				Label15:
					continue
				Label8:
					if πE != nil || πR != nil {
						continue
					}
				Label9:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 139: return list(funcs)
					πF.SetLineno(139)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfuncs, "funcs"); πE != nil {
						continue
					}
					πTemp003[0] = µfuncs
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßreferrednested.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 105: """get functions defined inside of func (e.g. inner functions in a closure)
			πF.SetLineno(105)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("get functions defined inside of func (e.g. inner functions in a closure)\n\n    NOTE: results may differ if the function has been executed or not.\n    If len(nestedcode(func)) > len(referrednested(func)), try calling func().\n    If possible, python builds code objects, but delays building functions\n    until func() is called.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßreferrednested); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
			// line 142: def freevars(func):
			πF.SetLineno(142)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("freevars", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µim_func *πg.Object = πg.UnboundLocal; _ = µim_func
				var µfunc_code *πg.Object = πg.UnboundLocal; _ = µfunc_code
				var µfunc_closure *πg.Object = πg.UnboundLocal; _ = µfunc_closure
				var µclosures *πg.Object = πg.UnboundLocal; _ = µclosures
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
				var πTemp006 *πg.Dict
				_ = πTemp006
				var πTemp007 []πg.Param
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 143: """get objects defined in enclosing code that are referred to by func
					πF.SetLineno(143)
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
					// line 146: if PY3:
					πF.SetLineno(146)
				Label1:
					// line 147: im_func = '__func__'
					πF.SetLineno(147)
					µim_func = ß__func__.ToObject()
					// line 148: func_code = '__code__'
					πF.SetLineno(148)
					µfunc_code = ß__code__.ToObject()
					// line 149: func_closure = '__closure__'
					πF.SetLineno(149)
					µfunc_closure = ß__closure__.ToObject()
					goto Label3
				Label2:
					// line 151: im_func = 'im_func'
					πF.SetLineno(151)
					µim_func = ßim_func.ToObject()
					// line 152: func_code = 'func_code'
					πF.SetLineno(152)
					µfunc_code = ßfunc_code.ToObject()
					// line 153: func_closure = 'func_closure'
					πF.SetLineno(153)
					µfunc_closure = ßfunc_closure.ToObject()
					goto Label3
				Label3:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßismethod); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 154: if ismethod(func): func = getattr(func, im_func)
					πF.SetLineno(154)
				Label4:
					// line 154: if ismethod(func): func = getattr(func, im_func)
					πF.SetLineno(154)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πE = πg.CheckLocal(πF, µim_func, "im_func"); πE != nil {
						continue
					}
					πTemp003[1] = µim_func
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µfunc = πTemp004
					goto Label5
				Label5:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisfunction); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					goto Label7
					// line 155: if isfunction(func):
					πF.SetLineno(155)
				Label6:
					// line 156: closures = getattr(func, func_closure) or ()
					πF.SetLineno(156)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πE = πg.CheckLocal(πF, µfunc_closure, "func_closure"); πE != nil {
						continue
					}
					πTemp003[1] = µfunc_closure
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
						goto Label9
					}
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp001 = πTemp004
				Label9:
					µclosures = πTemp001
					// line 157: func = getattr(func, func_code).co_freevars # get freevars
					πF.SetLineno(157)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πE = πg.CheckLocal(πF, µfunc_code, "func_code"); πE != nil {
						continue
					}
					πTemp003[1] = µfunc_code
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßco_freevars, nil); πE != nil {
						continue
					}
					µfunc = πTemp001
					goto Label8
				Label7:
					// line 159: return {}
					πF.SetLineno(159)
					πTemp006 = πg.NewDict()
					πTemp001 = πTemp006.ToObject()
					πR = πTemp001
					continue
					goto Label8
				Label8:
					// line 160: return dict((name,c.cell_contents) for (name,c) in zip(func,closures))
					πF.SetLineno(160)
					πTemp003 = πF.MakeArgs(1)
					πTemp007 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µname *πg.Object = πg.UnboundLocal; _ = µname
						var µc *πg.Object = πg.UnboundLocal; _ = µc
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
								if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
									continue
								}
								πTemp002[0] = µfunc
								if πE = πg.CheckLocal(πF, µclosures, "closures"); πE != nil {
									continue
								}
								πTemp002[1] = µclosures
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
									µname = πTemp004
									µc = πTemp007
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 160: return dict((name,c.cell_contents) for (name,c) in zip(func,closures))
								πF.SetLineno(160)
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µc, ßcell_contents, nil); πE != nil {
									continue
								}
								πTemp003 = πg.NewTuple2(µname, πTemp004).ToObject()
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
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
			if πE = πF.Globals().SetItem(πF, ßfreevars.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 143: """get objects defined in enclosing code that are referred to by func
			πF.SetLineno(143)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("get objects defined in enclosing code that are referred to by func\n\n    returns a dict of {name:object}").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßfreevars); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 163: def nestedglobals(func, recurse=True):
			πF.SetLineno(163)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "recurse", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("nestedglobals", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µrecurse *πg.Object = πArgs[1]; _ = µrecurse
				var µcapture *πg.Object = πg.UnboundLocal; _ = µcapture
				var µnames *πg.Object = πg.UnboundLocal; _ = µnames
				var µout *πg.Object = πg.UnboundLocal; _ = µout
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µco *πg.Object = πg.UnboundLocal; _ = µco
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 *πg.Type
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 bool
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 []*πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πTemp017 *πg.Object
				_ = πTemp017
				var πTemp018 πg.KWArgs
				_ = πTemp018
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 9: goto Label9
					case 10: goto Label10
					case 3: goto Label3
					case 4: goto Label4
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 164: """get the names of any globals found within func"""
					πF.SetLineno(164)
					// line 165: func = code(func)
					πF.SetLineno(165)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp001[0] = µfunc
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcode); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfunc = πTemp003
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µfunc == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 166: if func is None: return list()
					πF.SetLineno(166)
				Label1:
					// line 166: if func is None: return list()
					πF.SetLineno(166)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label2
				Label2:
					// line 167: from .temp import capture
					πF.SetLineno(167)
					if πTemp001, πE = πg.ImportModule(πF, ".temp"); πE != nil {
						continue
					}
					πTemp002 = πTemp001[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßcapture); πE != nil {
						continue
					}
					µcapture = πTemp003
					// line 168: names = set()
					πF.SetLineno(168)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µnames = πTemp003
					// line 169: with capture('stdout') as out:
					πF.SetLineno(169)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßstdout.ToObject()
					if πE = πg.CheckLocal(πF, µcapture, "capture"); πE != nil {
						continue
					}
					if πTemp002, πE = µcapture.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__exit__, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__enter__, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp005.Call(πF, πg.Args{πTemp002}, nil); πE != nil {
						continue
					}
					πF.PushCheckpoint(3)
					µout = πTemp005
					// line 170: dis.dis(func) #XXX: dis.dis(None) disassembles last traceback
					πF.SetLineno(170)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp001[0] = µfunc
					if πTemp006, πE = πg.ResolveGlobal(πF, ßdis); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßdis, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
				Label3:
					πTemp008, πTemp009 = nil, nil
					if πE != nil {
						πTemp008, πTemp009 = πF.ExcInfo()
					}
					if πTemp008 != nil {
						πTemp010 = πTemp008.Type()
						if πTemp006, πE = πTemp003.Call(πF, πg.Args{πTemp002, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
							continue
						}
					} else {
						if πTemp006, πE = πTemp003.Call(πF, πg.Args{πTemp002, πg.None, πg.None, πg.None}, nil); πE != nil {
							continue
						}
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp008 != nil && πTemp004 != true {
						πE = πF.Raise(nil, nil, nil)
						continue
					}
					if πR != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µout, ßgetvalue, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp011, ßsplitlines, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Iter(πF, πTemp011); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp004 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label6
					}
					if πTemp007, πE = πg.Next(πF, πTemp006); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp012 = !isStop
					} else {
						πTemp012 = true
						µline = πTemp007
					}
					if πE != nil || !πTemp012 {
						continue
					}
					πF.PushCheckpoint(4)            
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.Contains(πF, µline, ß_GLOBAL.ToObject()); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(πTemp012).ToObject()
					if πTemp012, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp012 {
						goto Label7
					}
					goto Label8
					// line 172: if '_GLOBAL' in line:
					πF.SetLineno(172)
				Label7:
					// line 173: name = line.split('(')[-1].split(')')[0]
					πF.SetLineno(173)
					πTemp007 = πg.NewInt(0).ToObject()
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(")").ToObject()
					if πTemp014, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp013 = πTemp014
					πTemp015 = πF.MakeArgs(1)
					πTemp015[0] = πg.NewStr("(").ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, µline, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp017, πE = πTemp016.Call(πF, πTemp015, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp015)
					if πTemp014, πE = πg.GetItem(πF, πTemp017, πTemp013); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp014, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp011, πE = πg.GetItem(πF, πTemp014, πTemp007); πE != nil {
						continue
					}
					µname = πTemp011
					// line 174: names.add(name)
					πF.SetLineno(174)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µnames, ßadd, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label8
				Label8:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp001[0] = µfunc
					πTemp001[1] = ßco_consts.ToObject()
					if πTemp007, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[2] = πTemp011
					if πTemp007, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.Iter(πF, πTemp011); πE != nil {
						continue
					}
					πF.PushCheckpoint(10)
					πTemp004 = false
				Label9:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label11
					}
					if πTemp007, πE = πg.Next(πF, πTemp006); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp012 = !isStop
					} else {
						πTemp012 = true
						µco = πTemp007
					}
					if πE != nil || !πTemp012 {
						continue
					}
					πF.PushCheckpoint(9)            
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					πTemp007 = µco
					if πTemp012, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if !πTemp012 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
						continue
					}
					πTemp007 = µrecurse
					if πTemp012, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if !πTemp012 {
						goto Label12
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					πTemp001[0] = µco
					if πTemp011, πE = πg.ResolveGlobal(πF, ßiscode); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp007 = πTemp013
				Label12:
					if πTemp012, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp012 {
						goto Label13
					}
					goto Label14
					// line 176: if co and recurse and iscode(co):
					πF.SetLineno(176)
				Label13:
					// line 177: names.update(nestedglobals(co, recurse=True))
					πF.SetLineno(177)
					πTemp001 = πF.MakeArgs(1)
					πTemp015 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					πTemp015[0] = µco
					if πTemp007, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp018 = πg.KWArgs{
						{"recurse", πTemp007},
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßnestedglobals); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp015, πTemp018); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp015)
					πTemp001[0] = πTemp011
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µnames, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label14
				Label14:
					continue
				Label10:
					if πE != nil || πR != nil {
						continue
					}
				Label11:
					// line 178: return list(names)
					πF.SetLineno(178)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					πTemp001[0] = µnames
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp007
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnestedglobals.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 164: """get the names of any globals found within func"""
			πF.SetLineno(164)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("get the names of any globals found within func").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßnestedglobals); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 180: def referredglobals(func, recurse=True, builtin=False):
			πF.SetLineno(180)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "recurse", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "builtin", Def: πTemp011}
			πTemp010 = πg.NewFunction(πg.NewCode("referredglobals", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µrecurse *πg.Object = πArgs[1]; _ = µrecurse
				var µbuiltin *πg.Object = πArgs[2]; _ = µbuiltin
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 181: """get the names of objects in the global scope referred to by func"""
					πF.SetLineno(181)
					// line 182: return globalvars(func, recurse, builtin).keys()
					πF.SetLineno(182)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp001[0] = µfunc
					if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
						continue
					}
					πTemp001[1] = µrecurse
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp001[2] = µbuiltin
					if πTemp002, πE = πg.ResolveGlobal(πF, ßglobalvars); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßreferredglobals.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 181: """get the names of objects in the global scope referred to by func"""
			πF.SetLineno(181)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("get the names of objects in the global scope referred to by func").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßreferredglobals); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 184: def globalvars(func, recurse=True, builtin=False):
			πF.SetLineno(184)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "recurse", Def: πTemp012}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "builtin", Def: πTemp012}
			πTemp011 = πg.NewFunction(πg.NewCode("globalvars", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µrecurse *πg.Object = πArgs[1]; _ = µrecurse
				var µbuiltin *πg.Object = πArgs[2]; _ = µbuiltin
				var µim_func *πg.Object = πg.UnboundLocal; _ = µim_func
				var µfunc_code *πg.Object = πg.UnboundLocal; _ = µfunc_code
				var µfunc_globals *πg.Object = πg.UnboundLocal; _ = µfunc_globals
				var µfunc_closure *πg.Object = πg.UnboundLocal; _ = µfunc_closure
				var µglobs *πg.Object = πg.UnboundLocal; _ = µglobs
				var µorig_func *πg.Object = πg.UnboundLocal; _ = µorig_func
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
				var µ_vars *πg.Object = πg.UnboundLocal; _ = µ_vars
				var µkey *πg.Object = πg.UnboundLocal; _ = µkey
				var µnested_func *πg.Object = πg.UnboundLocal; _ = µnested_func
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Dict
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πTemp011 []πg.Param
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 32: goto Label32
					case 13: goto Label13
					case 14: goto Label14
					case 21: goto Label21
					case 22: goto Label22
					case 31: goto Label31
					default: panic("unexpected function state")
					}
					// line 185: """get objects defined in global scope that are referred to by func
					πF.SetLineno(185)
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
					// line 188: if PY3:
					πF.SetLineno(188)
				Label1:
					// line 189: im_func = '__func__'
					πF.SetLineno(189)
					µim_func = ß__func__.ToObject()
					// line 190: func_code = '__code__'
					πF.SetLineno(190)
					µfunc_code = ß__code__.ToObject()
					// line 191: func_globals = '__globals__'
					πF.SetLineno(191)
					µfunc_globals = ß__globals__.ToObject()
					// line 192: func_closure = '__closure__'
					πF.SetLineno(192)
					µfunc_closure = ß__closure__.ToObject()
					goto Label3
				Label2:
					// line 194: im_func = 'im_func'
					πF.SetLineno(194)
					µim_func = ßim_func.ToObject()
					// line 195: func_code = 'func_code'
					πF.SetLineno(195)
					µfunc_code = ßfunc_code.ToObject()
					// line 196: func_globals = 'func_globals'
					πF.SetLineno(196)
					µfunc_globals = ßfunc_globals.ToObject()
					// line 197: func_closure = 'func_closure'
					πF.SetLineno(197)
					µfunc_closure = ßfunc_closure.ToObject()
					goto Label3
				Label3:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßismethod); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 198: if ismethod(func): func = getattr(func, im_func)
					πF.SetLineno(198)
				Label4:
					// line 198: if ismethod(func): func = getattr(func, im_func)
					πF.SetLineno(198)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πE = πg.CheckLocal(πF, µim_func, "im_func"); πE != nil {
						continue
					}
					πTemp003[1] = µim_func
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µfunc = πTemp004
					goto Label5
				Label5:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisfunction); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp003[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßiscode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					goto Label8
					// line 199: if isfunction(func):
					πF.SetLineno(199)
				Label6:
					// line 200: globs = vars(getmodule(sum)).copy() if builtin else {}
					πF.SetLineno(200)
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µbuiltin); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label10
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsum); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp006
					if πTemp004, πE = πg.ResolveGlobal(πF, ßvars); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp006
					goto Label11
				Label10:
					πTemp007 = πg.NewDict()
					πTemp004 = πTemp007.ToObject()
					πTemp001 = πTemp004
				Label11:
					µglobs = πTemp001
					// line 202: orig_func, func = func, set()
					πF.SetLineno(202)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µfunc, πTemp006).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
						continue
					}
					µorig_func = πTemp004
					µfunc = πTemp006
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µorig_func, "orig_func"); πE != nil {
						continue
					}
					πTemp003[0] = µorig_func
					if πE = πg.CheckLocal(πF, µfunc_closure, "func_closure"); πE != nil {
						continue
					}
					πTemp003[1] = µfunc_closure
					if πTemp006, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp004 = πTemp008
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					πTemp007 = πg.NewDict()
					πTemp006 = πTemp007.ToObject()
					πTemp004 = πTemp006
				Label12:
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(14)
					πTemp002 = false
				Label13:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label15
					}
					if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						µobj = πTemp004
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(13)            
					// line 204: _vars = globalvars(obj.cell_contents, recurse, builtin) or {}
					πF.SetLineno(204)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µobj, ßcell_contents, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp006
					if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
						continue
					}
					πTemp003[1] = µrecurse
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp003[2] = µbuiltin
					if πTemp006, πE = πg.ResolveGlobal(πF, ßglobalvars); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp004 = πTemp008
					if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label16
					}
					πTemp007 = πg.NewDict()
					πTemp006 = πTemp007.ToObject()
					πTemp004 = πTemp006
				Label16:
					µ_vars = πTemp004
					// line 205: func.update(_vars) #XXX: (above) be wary of infinte recursion?
					πF.SetLineno(205)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_vars, "_vars"); πE != nil {
						continue
					}
					πTemp003[0] = µ_vars
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfunc, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 206: globs.update(_vars)
					πF.SetLineno(206)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_vars, "_vars"); πE != nil {
						continue
					}
					πTemp003[0] = µ_vars
					if πE = πg.CheckLocal(πF, µglobs, "globs"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µglobs, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					continue
				Label14:
					if πE != nil || πR != nil {
						continue
					}
				Label15:
					// line 208: globs.update(getattr(orig_func, func_globals) or {})
					πF.SetLineno(208)
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µorig_func, "orig_func"); πE != nil {
						continue
					}
					πTemp005[0] = µorig_func
					if πE = πg.CheckLocal(πF, µfunc_globals, "func_globals"); πE != nil {
						continue
					}
					πTemp005[1] = µfunc_globals
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp006
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label17
					}
					πTemp007 = πg.NewDict()
					πTemp004 = πTemp007.ToObject()
					πTemp001 = πTemp004
				Label17:
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µglobs, "globs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µglobs, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µrecurse); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label18
					}
					goto Label19
					// line 210: if not recurse:
					πF.SetLineno(210)
				Label18:
					// line 211: func.update(getattr(orig_func, func_code).co_names)
					πF.SetLineno(211)
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µorig_func, "orig_func"); πE != nil {
						continue
					}
					πTemp005[0] = µorig_func
					if πE = πg.CheckLocal(πF, µfunc_code, "func_code"); πE != nil {
						continue
					}
					πTemp005[1] = µfunc_code
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßco_names, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfunc, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label20
				Label19:
					// line 213: func.update(nestedglobals(getattr(orig_func, func_code)))
					πF.SetLineno(213)
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µorig_func, "orig_func"); πE != nil {
						continue
					}
					πTemp010[0] = µorig_func
					if πE = πg.CheckLocal(πF, µfunc_code, "func_code"); πE != nil {
						continue
					}
					πTemp010[1] = µfunc_code
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp005[0] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßnestedglobals); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfunc, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfunc, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
						continue
					}
					πF.PushCheckpoint(22)
					πTemp002 = false
				Label21:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label23
					}
					if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						µkey = πTemp004
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(21)            
					// line 216: nested_func = globs.get(key)
					πF.SetLineno(216)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp003[0] = µkey
					if πE = πg.CheckLocal(πF, µglobs, "globs"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µglobs, ßget, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µnested_func = πTemp006
					if πE = πg.CheckLocal(πF, µnested_func, "nested_func"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µorig_func, "orig_func"); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(µnested_func == µorig_func).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label24
					}
					goto Label25
					// line 217: if nested_func is orig_func:
					πF.SetLineno(217)
				Label24:
					// line 219: continue  #XXX: globalvars(func, False)?
					πF.SetLineno(219)
					continue
					goto Label25
				Label25:
					// line 220: func.update(globalvars(nested_func, True, builtin))
					πF.SetLineno(220)
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µnested_func, "nested_func"); πE != nil {
						continue
					}
					πTemp005[0] = µnested_func
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp005[1] = πTemp004
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp005[2] = µbuiltin
					if πTemp004, πE = πg.ResolveGlobal(πF, ßglobalvars); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp006
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfunc, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					continue
				Label22:
					if πE != nil || πR != nil {
						continue
					}
				Label23:
					goto Label20
				Label20:
					goto Label9
					// line 221: elif iscode(func):
					πF.SetLineno(221)
				Label7:
					// line 222: globs = vars(getmodule(sum)).copy() if builtin else {}
					πF.SetLineno(222)
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µbuiltin); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label26
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsum); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetmodule); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp006
					if πTemp004, πE = πg.ResolveGlobal(πF, ßvars); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp006
					goto Label27
				Label26:
					πTemp007 = πg.NewDict()
					πTemp004 = πTemp007.ToObject()
					πTemp001 = πTemp004
				Label27:
					µglobs = πTemp001
					if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µrecurse); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label28
					}
					goto Label29
					// line 224: if not recurse:
					πF.SetLineno(224)
				Label28:
					// line 225: func = func.co_names # get names
					πF.SetLineno(225)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfunc, ßco_names, nil); πE != nil {
						continue
					}
					µfunc = πTemp001
					goto Label30
				Label29:
					// line 227: orig_func = func.co_name # to stop infinite recursion
					πF.SetLineno(227)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfunc, ßco_name, nil); πE != nil {
						continue
					}
					µorig_func = πTemp001
					// line 228: func = set(nestedglobals(func))
					πF.SetLineno(228)
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp005[0] = µfunc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßnestedglobals); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µfunc = πTemp004
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfunc, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
						continue
					}
					πF.PushCheckpoint(32)
					πTemp002 = false
				Label31:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label33
					}
					if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						µkey = πTemp004
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(31)            
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µorig_func, "orig_func"); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(µkey == µorig_func).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label34
					}
					goto Label35
					// line 231: if key is orig_func:
					πF.SetLineno(231)
				Label34:
					// line 233: continue  #XXX: globalvars(func, False)?
					πF.SetLineno(233)
					continue
					goto Label35
				Label35:
					// line 234: nested_func = globs.get(key)
					πF.SetLineno(234)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp003[0] = µkey
					if πE = πg.CheckLocal(πF, µglobs, "globs"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µglobs, ßget, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µnested_func = πTemp006
					// line 235: func.update(globalvars(nested_func, True, builtin))
					πF.SetLineno(235)
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µnested_func, "nested_func"); πE != nil {
						continue
					}
					πTemp005[0] = µnested_func
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp005[1] = πTemp004
					if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
						continue
					}
					πTemp005[2] = µbuiltin
					if πTemp004, πE = πg.ResolveGlobal(πF, ßglobalvars); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp006
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfunc, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					continue
				Label32:
					if πE != nil || πR != nil {
						continue
					}
				Label33:
					goto Label30
				Label30:
					goto Label9
				Label8:
					// line 237: return {}
					πF.SetLineno(237)
					πTemp007 = πg.NewDict()
					πTemp001 = πTemp007.ToObject()
					πR = πTemp001
					continue
					goto Label9
				Label9:
					// line 239: return dict((name,globs[name]) for name in func if name in globs)
					πF.SetLineno(239)
					πTemp003 = πF.MakeArgs(1)
					πTemp011 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µname *πg.Object = πg.UnboundLocal; _ = µname
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
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µfunc); πE != nil {
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
									µname = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µglobs, "globs"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Contains(πF, µglobs, µname); πE != nil {
									continue
								}
								πTemp004 = πg.GetBool(πTemp003).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								goto Label5
								// line 239: return dict((name,globs[name]) for name in func if name in globs)
								πF.SetLineno(239)
							Label4:
								// line 239: return dict((name,globs[name]) for name in func if name in globs)
								πF.SetLineno(239)
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								πTemp005 = µname
								if πE = πg.CheckLocal(πF, µglobs, "globs"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, µglobs, πTemp005); πE != nil {
									continue
								}
								πTemp004 = πg.NewTuple2(µname, πTemp006).ToObject()
								πF.PushCheckpoint(6)
								return πTemp004, nil
							Label6:
								πTemp005 = πSent
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
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp006
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßglobalvars.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 185: """get objects defined in global scope that are referred to by func
			πF.SetLineno(185)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("get objects defined in global scope that are referred to by func\n\n    return a dict of {name:object}").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßglobalvars); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 242: def varnames(func):
			πF.SetLineno(242)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("varnames", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 243: """get names of variables defined by func
					πF.SetLineno(243)
					// line 246: func = code(func)
					πF.SetLineno(246)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp001[0] = µfunc
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcode); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfunc = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp001[0] = µfunc
					if πTemp003, πE = πg.ResolveGlobal(πF, ßiscode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 247: if not iscode(func):
					πF.SetLineno(247)
				Label1:
					// line 248: return () #XXX: better ((),())? or None?
					πF.SetLineno(248)
					πTemp002 = πg.NewTuple0().ToObject()
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 249: return func.co_varnames, func.co_cellvars
					πF.SetLineno(249)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfunc, ßco_varnames, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfunc, ßco_cellvars, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßvarnames.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 243: """get names of variables defined by func
			πF.SetLineno(243)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("get names of variables defined by func\n\n    returns a tuple (local vars, local vars referrenced by nested functions)").ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßvarnames); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
				continue
			}
			// line 252: def baditems(obj, exact=False, safe=False): #XXX: obj=globals() ?
			πF.SetLineno(252)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "exact", Def: πTemp014}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "safe", Def: πTemp014}
			πTemp013 = πg.NewFunction(πg.NewCode("baditems", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µexact *πg.Object = πArgs[1]; _ = µexact
				var µsafe *πg.Object = πArgs[2]; _ = µsafe
				var µ_obj *πg.Object = πg.UnboundLocal; _ = µ_obj
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
				var πTemp006 []πg.Param
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
					// line 253: """get items in object that fail to pickle"""
					πF.SetLineno(253)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					πTemp002[1] = ß__iter__.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 254: if not hasattr(obj,'__iter__'): # is not iterable
					πF.SetLineno(254)
				Label1:
					// line 255: return [j for j in (badobjects(obj,0,exact,safe),) if j is not None]
					πF.SetLineno(255)
					πTemp006 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µj *πg.Object = πg.UnboundLocal; _ = µj
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								πTemp003 = πF.MakeArgs(4)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp003[0] = µobj
								πTemp003[1] = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
									continue
								}
								πTemp003[2] = µexact
								if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
									continue
								}
								πTemp003[3] = µsafe
								if πTemp004, πE = πg.ResolveGlobal(πF, ßbadobjects); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								πTemp002 = πg.NewTuple1(πTemp005).ToObject()
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
									µj = πTemp002
								}
								if πE != nil || !πTemp007 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp002 = πg.GetBool(µj != πTemp004).ToObject()
								if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp007 {
									goto Label4
								}
								goto Label5
								// line 255: return [j for j in (badobjects(obj,0,exact,safe),) if j is not None]
								πF.SetLineno(255)
							Label4:
								// line 255: return [j for j in (badobjects(obj,0,exact,safe),) if j is not None]
								πF.SetLineno(255)
								if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return µj, nil
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
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 256: obj = obj.values() if getattr(obj,'values',None) else obj
					πF.SetLineno(256)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					πTemp002[1] = ßvalues.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp007
					goto Label4
				Label3:
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001 = µobj
				Label4:
					µobj = πTemp001
					// line 257: _obj = [] # can't use a set, as items may be unhashable
					πF.SetLineno(257)
					πTemp002 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					µ_obj = πTemp001
					// line 258: [_obj.append(badobjects(i,0,exact,safe)) for i in obj if i not in _obj]
					πF.SetLineno(258)
					πTemp006 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
						var πTemp006 []*πg.Object
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
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µobj); πE != nil {
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
									µi = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µ_obj, "_obj"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Contains(πF, µ_obj, µi); πE != nil {
									continue
								}
								πTemp004 = πg.GetBool(!πTemp003).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								goto Label5
								// line 258: [_obj.append(badobjects(i,0,exact,safe)) for i in obj if i not in _obj]
								πF.SetLineno(258)
							Label4:
								// line 258: [_obj.append(badobjects(i,0,exact,safe)) for i in obj if i not in _obj]
								πF.SetLineno(258)
								πTemp005 = πF.MakeArgs(1)
								πTemp006 = πF.MakeArgs(4)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp006[0] = µi
								πTemp006[1] = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
									continue
								}
								πTemp006[2] = µexact
								if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
									continue
								}
								πTemp006[3] = µsafe
								if πTemp004, πE = πg.ResolveGlobal(πF, ßbadobjects); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								πTemp005[0] = πTemp007
								if πE = πg.CheckLocal(πF, µ_obj, "_obj"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µ_obj, ßappend, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(6)
								return πTemp007, nil
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
					if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
						continue
					}
					// line 259: return [j for j in _obj if j is not None]
					πF.SetLineno(259)
					πTemp006 = make([]πg.Param, 0)
					πTemp007 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µj *πg.Object = πg.UnboundLocal; _ = µj
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
								if πE = πg.CheckLocal(πF, µ_obj, "_obj"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µ_obj); πE != nil {
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
									µj = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp004 = πg.GetBool(µj != πTemp005).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								goto Label5
								// line 259: return [j for j in _obj if j is not None]
								πF.SetLineno(259)
							Label4:
								// line 259: return [j for j in _obj if j is not None]
								πF.SetLineno(259)
								if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return µj, nil
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
					if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßbaditems.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 253: """get items in object that fail to pickle"""
			πF.SetLineno(253)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("get items in object that fail to pickle").ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßbaditems); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
				continue
			}
			// line 262: def badobjects(obj, depth=0, exact=False, safe=False):
			πF.SetLineno(262)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "depth", Def: πg.NewInt(0).ToObject()}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "exact", Def: πTemp015}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "safe", Def: πTemp015}
			πTemp014 = πg.NewFunction(πg.NewCode("badobjects", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µdepth *πg.Object = πArgs[1]; _ = µdepth
				var µexact *πg.Object = πArgs[2]; _ = µexact
				var µsafe *πg.Object = πArgs[3]; _ = µsafe
				var µpickles *πg.Object = πg.UnboundLocal; _ = µpickles
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []πg.Param
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
					// line 263: """get objects that fail to pickle"""
					πF.SetLineno(263)
					// line 264: from dill import pickles
					πF.SetLineno(264)
					if πTemp002, πE = πg.ImportModule(πF, "dill"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßpickles); πE != nil {
						continue
					}
					µpickles = πTemp003
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µdepth); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 265: if not depth:
					πF.SetLineno(265)
				Label1:
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
						continue
					}
					πTemp002[1] = µexact
					if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
						continue
					}
					πTemp002[2] = µsafe
					if πE = πg.CheckLocal(πF, µpickles, "pickles"); πE != nil {
						continue
					}
					if πTemp001, πE = µpickles.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 266: if pickles(obj,exact,safe): return None
					πF.SetLineno(266)
				Label3:
					// line 266: if pickles(obj,exact,safe): return None
					πF.SetLineno(266)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label4
				Label4:
					// line 267: return obj
					πF.SetLineno(267)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πR = µobj
					continue
					goto Label2
				Label2:
					// line 268: return dict(((attr, badobjects(getattr(obj,attr),depth-1,exact,safe)) \
					πF.SetLineno(268)
					πTemp002 = πF.MakeArgs(1)
					πTemp005 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µattr *πg.Object = πg.UnboundLocal; _ = µattr
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp002[0] = µobj
								if πTemp003, πE = πg.ResolveGlobal(πF, ßdir); πE != nil {
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
									µattr = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp002 = πF.MakeArgs(3)
								πTemp007 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp007[0] = µobj
								if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
									continue
								}
								πTemp007[1] = µattr
								if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp002[0] = πTemp008
								if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
									continue
								}
								πTemp002[1] = µexact
								if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
									continue
								}
								πTemp002[2] = µsafe
								if πE = πg.CheckLocal(πF, µpickles, "pickles"); πE != nil {
									continue
								}
								if πTemp004, πE = µpickles.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(!πTemp006).ToObject()
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 268: return dict(((attr, badobjects(getattr(obj,attr),depth-1,exact,safe)) \
								πF.SetLineno(268)
							Label4:
								// line 268: return dict(((attr, badobjects(getattr(obj,attr),depth-1,exact,safe)) \
								πF.SetLineno(268)
								if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
									continue
								}
								πTemp002 = πF.MakeArgs(4)
								πTemp007 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp007[0] = µobj
								if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
									continue
								}
								πTemp007[1] = µattr
								if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp002[0] = πTemp008
								if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.Sub(πF, µdepth, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp002[1] = πTemp004
								if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
									continue
								}
								πTemp002[2] = µexact
								if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
									continue
								}
								πTemp002[3] = µsafe
								if πTemp004, πE = πg.ResolveGlobal(πF, ßbadobjects); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πTemp003 = πg.NewTuple2(µattr, πTemp008).ToObject()
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
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp006
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßbadobjects.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 263: """get objects that fail to pickle"""
			πF.SetLineno(263)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("get objects that fail to pickle").ToObject()); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßbadobjects); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
				continue
			}
			// line 271: def badtypes(obj, depth=0, exact=False, safe=False):
			πF.SetLineno(271)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "depth", Def: πg.NewInt(0).ToObject()}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "exact", Def: πTemp016}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "safe", Def: πTemp016}
			πTemp015 = πg.NewFunction(πg.NewCode("badtypes", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µdepth *πg.Object = πArgs[1]; _ = µdepth
				var µexact *πg.Object = πArgs[2]; _ = µexact
				var µsafe *πg.Object = πArgs[3]; _ = µsafe
				var µpickles *πg.Object = πg.UnboundLocal; _ = µpickles
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []πg.Param
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
					// line 272: """get types for objects that fail to pickle"""
					πF.SetLineno(272)
					// line 273: from dill import pickles
					πF.SetLineno(273)
					if πTemp002, πE = πg.ImportModule(πF, "dill"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßpickles); πE != nil {
						continue
					}
					µpickles = πTemp003
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µdepth); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 274: if not depth:
					πF.SetLineno(274)
				Label1:
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
						continue
					}
					πTemp002[1] = µexact
					if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
						continue
					}
					πTemp002[2] = µsafe
					if πE = πg.CheckLocal(πF, µpickles, "pickles"); πE != nil {
						continue
					}
					if πTemp001, πE = µpickles.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 275: if pickles(obj,exact,safe): return None
					πF.SetLineno(275)
				Label3:
					// line 275: if pickles(obj,exact,safe): return None
					πF.SetLineno(275)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label4
				Label4:
					// line 276: return type(obj)
					πF.SetLineno(276)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp003
					continue
					goto Label2
				Label2:
					// line 277: return dict(((attr, badtypes(getattr(obj,attr),depth-1,exact,safe)) \
					πF.SetLineno(277)
					πTemp002 = πF.MakeArgs(1)
					πTemp005 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µattr *πg.Object = πg.UnboundLocal; _ = µattr
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp002[0] = µobj
								if πTemp003, πE = πg.ResolveGlobal(πF, ßdir); πE != nil {
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
									µattr = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp002 = πF.MakeArgs(3)
								πTemp007 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp007[0] = µobj
								if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
									continue
								}
								πTemp007[1] = µattr
								if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp002[0] = πTemp008
								if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
									continue
								}
								πTemp002[1] = µexact
								if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
									continue
								}
								πTemp002[2] = µsafe
								if πE = πg.CheckLocal(πF, µpickles, "pickles"); πE != nil {
									continue
								}
								if πTemp004, πE = µpickles.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(!πTemp006).ToObject()
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 277: return dict(((attr, badtypes(getattr(obj,attr),depth-1,exact,safe)) \
								πF.SetLineno(277)
							Label4:
								// line 277: return dict(((attr, badtypes(getattr(obj,attr),depth-1,exact,safe)) \
								πF.SetLineno(277)
								if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
									continue
								}
								πTemp002 = πF.MakeArgs(4)
								πTemp007 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp007[0] = µobj
								if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
									continue
								}
								πTemp007[1] = µattr
								if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp002[0] = πTemp008
								if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.Sub(πF, µdepth, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp002[1] = πTemp004
								if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
									continue
								}
								πTemp002[2] = µexact
								if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
									continue
								}
								πTemp002[3] = µsafe
								if πTemp004, πE = πg.ResolveGlobal(πF, ßbadtypes); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πTemp003 = πg.NewTuple2(µattr, πTemp008).ToObject()
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
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp006
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßbadtypes.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 272: """get types for objects that fail to pickle"""
			πF.SetLineno(272)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("get types for objects that fail to pickle").ToObject()); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßbadtypes); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
				continue
			}
			// line 280: def errors(obj, depth=0, exact=False, safe=False):
			πF.SetLineno(280)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "depth", Def: πg.NewInt(0).ToObject()}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "exact", Def: πTemp017}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "safe", Def: πTemp017}
			πTemp016 = πg.NewFunction(πg.NewCode("errors", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/detect.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µdepth *πg.Object = πArgs[1]; _ = µdepth
				var µexact *πg.Object = πArgs[2]; _ = µexact
				var µsafe *πg.Object = πArgs[3]; _ = µsafe
				var µpickles *πg.Object = πg.UnboundLocal; _ = µpickles
				var µcopy *πg.Object = πg.UnboundLocal; _ = µcopy
				var µpik *πg.Object = πg.UnboundLocal; _ = µpik
				var µsys *πg.Object = πg.UnboundLocal; _ = µsys
				var µ_dict *πg.Object = πg.UnboundLocal; _ = µ_dict
				var µattr *πg.Object = πg.UnboundLocal; _ = µattr
				var µ_attr *πg.Object = πg.UnboundLocal; _ = µ_attr
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
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
				var πTemp010 *πg.Dict
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8: goto Label8
					case 9: goto Label9
					case 4: goto Label4
					case 12: goto Label12
					default: panic("unexpected function state")
					}
					// line 281: """get errors for objects that fail to pickle"""
					πF.SetLineno(281)
					// line 282: from dill import pickles, copy
					πF.SetLineno(282)
					if πTemp002, πE = πg.ImportModule(πF, "dill"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßpickles); πE != nil {
						continue
					}
					µpickles = πTemp003
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßcopy); πE != nil {
						continue
					}
					µcopy = πTemp003
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µdepth); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 283: if not depth:
					πF.SetLineno(283)
				Label1:
					// line 284: try:
					πF.SetLineno(284)
					πF.PushCheckpoint(4)
					// line 285: pik = copy(obj)
					πF.SetLineno(285)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πE = πg.CheckLocal(πF, µcopy, "copy"); πE != nil {
						continue
					}
					if πTemp001, πE = µcopy.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µpik = πTemp001
					if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µexact); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 286: if exact:
					πF.SetLineno(286)
				Label5:
					// line 287: assert pik == obj, \
					πF.SetLineno(287)
					if πE = πg.CheckLocal(πF, µpik, "pik"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µpik, µobj).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Unpickling produces %s instead of %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpik, "pik"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µpik, µobj); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					goto Label6
				Label6:
					// line 289: assert type(pik) == type(obj), \
					πF.SetLineno(289)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpik, "pik"); πE != nil {
						continue
					}
					πTemp002[0] = µpik
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp003 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Unpickling produces %s instead of %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpik, "pik"); πE != nil {
						continue
					}
					πTemp002[0] = µpik
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Eq(πF, πTemp006, πTemp007); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					// line 291: return None
					πF.SetLineno(291)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 292: except Exception:
					πF.SetLineno(292)
				Label7:
					// line 293: import sys
					πF.SetLineno(293)
					if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µsys = πTemp001
					// line 294: return sys.exc_info()[1]
					πF.SetLineno(294)
					πTemp001 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µsys, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					goto Label2
				Label2:
					// line 295: _dict = {}
					πF.SetLineno(295)
					πTemp010 = πg.NewDict()
					πTemp001 = πTemp010.ToObject()
					µ_dict = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdir); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(9)
					πTemp004 = false
				Label8:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label10
					}
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp011 = !isStop
					} else {
						πTemp011 = true
						µattr = πTemp003
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(8)            
					// line 297: try:
					πF.SetLineno(297)
					πF.PushCheckpoint(12)
					// line 298: _attr = getattr(obj,attr)
					πF.SetLineno(298)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp002[1] = µattr
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µ_attr = πTemp005
					πF.PopCheckpoint()
					goto Label11
				Label12:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
						continue
					}
					if πTemp011, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label13
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 299: except Exception:
					πF.SetLineno(299)
				Label13:
					// line 300: import sys
					πF.SetLineno(300)
					if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
						continue
					}
					πTemp003 = πTemp002[0]
					µsys = πTemp003
					// line 301: _dict[attr] = sys.exc_info()[1]
					πF.SetLineno(301)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µsys, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_dict, "_dict"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp006 = µattr
					if πE = πg.SetItem(πF, µ_dict, πTemp006, πTemp003); πE != nil {
						continue
					}
					// line 302: continue
					πF.SetLineno(302)
					continue
					πF.RestoreExc(nil, nil)
					goto Label11
				Label11:
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µ_attr, "_attr"); πE != nil {
						continue
					}
					πTemp002[0] = µ_attr
					if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
						continue
					}
					πTemp002[1] = µexact
					if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
						continue
					}
					πTemp002[2] = µsafe
					if πE = πg.CheckLocal(πF, µpickles, "pickles"); πE != nil {
						continue
					}
					if πTemp005, πE = µpickles.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp011, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp011).ToObject()
					if πTemp011, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label14
					}
					goto Label15
					// line 303: if not pickles(_attr,exact,safe):
					πF.SetLineno(303)
				Label14:
					// line 304: _dict[attr] = errors(_attr,depth-1,exact,safe)
					πF.SetLineno(304)
					πTemp002 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µ_attr, "_attr"); πE != nil {
						continue
					}
					πTemp002[0] = µ_attr
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µdepth, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
						continue
					}
					πTemp002[2] = µexact
					if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
						continue
					}
					πTemp002[3] = µsafe
					if πTemp003, πE = πg.ResolveGlobal(πF, ßerrors); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_dict, "_dict"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp006 = µattr
					if πE = πg.SetItem(πF, µ_dict, πTemp006, πTemp003); πE != nil {
						continue
					}
					goto Label15
				Label15:
					continue
				Label9:
					if πE != nil || πR != nil {
						continue
					}
				Label10:
					// line 305: return _dict
					πF.SetLineno(305)
					if πE = πg.CheckLocal(πF, µ_dict, "_dict"); πE != nil {
						continue
					}
					πR = µ_dict
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßerrors.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 281: """get errors for objects that fail to pickle"""
			πF.SetLineno(281)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("get errors for objects that fail to pickle").ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßerrors); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("detect", Code)
}

