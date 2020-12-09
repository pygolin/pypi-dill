package __diff
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/os"
	// _ "github.com/pygolin/stdlib/__builtin__"
	// _ "github.com/pygolin/stdlib/types"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAttributeError := πg.InternStr("AttributeError")
		ßBuiltinFunctionType := πg.InternStr("BuiltinFunctionType")
		ßFalse := πg.InternStr("False")
		ßFunctionType := πg.InternStr("FunctionType")
		ßHAS_NUMPY := πg.InternStr("HAS_NUMPY")
		ßImportError := πg.InternStr("ImportError")
		ßMaskedConstant := πg.InternStr("MaskedConstant")
		ßModuleType := πg.InternStr("ModuleType")
		ßNone := πg.InternStr("None")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßTrue := πg.InternStr("True")
		ß_ := πg.InternStr("_")
		ß__class__ := πg.InternStr("__class__")
		ß__contains__ := πg.InternStr("__contains__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__import__ := πg.InternStr("__import__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__len__ := πg.InternStr("__len__")
		ß_imp := πg.InternStr("_imp")
		ßany := πg.InternStr("any")
		ßbuiltins := πg.InternStr("builtins")
		ßbuiltins_types := πg.InternStr("builtins_types")
		ßcopy := πg.InternStr("copy")
		ßcore := πg.InternStr("core")
		ßdict := πg.InternStr("dict")
		ßdifference := πg.InternStr("difference")
		ßdont_memo := πg.InternStr("dont_memo")
		ßenviron := πg.InternStr("environ")
		ßfrozenset := πg.InternStr("frozenset")
		ßget := πg.InternStr("get")
		ßget_attrs := πg.InternStr("get_attrs")
		ßget_seq := πg.InternStr("get_seq")
		ßgetattr := πg.InternStr("getattr")
		ßgetrefcount := πg.InternStr("getrefcount")
		ßhas_changed := πg.InternStr("has_changed")
		ßhasattr := πg.InternStr("hasattr")
		ßid := πg.InternStr("id")
		ßid_to_obj := πg.InternStr("id_to_obj")
		ßint := πg.InternStr("int")
		ßitems := πg.InternStr("items")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßma := πg.InternStr("ma")
		ßmemo := πg.InternStr("memo")
		ßmemorise := πg.InternStr("memorise")
		ßmod := πg.InternStr("mod")
		ßmodules := πg.InternStr("modules")
		ßndarray := πg.InternStr("ndarray")
		ßnumpy := πg.InternStr("numpy")
		ßos := πg.InternStr("os")
		ßpath_importer_cache := πg.InternStr("path_importer_cache")
		ßpop := πg.InternStr("pop")
		ßrelease_gone := πg.InternStr("release_gone")
		ßset := πg.InternStr("set")
		ßshape := πg.InternStr("shape")
		ßsimple := πg.InternStr("simple")
		ßsize := πg.InternStr("size")
		ßstr := πg.InternStr("str")
		ßsys := πg.InternStr("sys")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßtypes := πg.InternStr("types")
		ßvalues := πg.InternStr("values")
		ßwhats_changed := πg.InternStr("whats_changed")
		ßzip := πg.InternStr("zip")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.BaseException
		_ = πTemp003
		var πTemp004 *πg.Traceback
		_ = πTemp004
		var πTemp005 bool
		_ = πTemp005
		var πTemp006 []πg.Param
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Dict
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
		var πTemp018 bool
		_ = πTemp018
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 9: goto Label9
			case 2: goto Label2
			case 10: goto Label10
			case 5: goto Label5
			default: panic("unexpected function state")
			}
			// line 9: """
			πF.SetLineno(9)
			// line 9: """
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nModule to show if an object has changed since it was memorised\n").ToObject()); πE != nil {
				continue
			}
			// line 13: import os
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: import sys
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: import types
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtypes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: try:
			πF.SetLineno(16)
			πF.PushCheckpoint(2)
			// line 17: import numpy
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "numpy"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßnumpy.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: HAS_NUMPY = True
			πF.SetLineno(18)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHAS_NUMPY.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp003, πTemp004 = πF.ExcInfo()
			goto Label3
			// line 19: except:
			πF.SetLineno(19)
		Label3:
			// line 20: HAS_NUMPY = False
			πF.SetLineno(20)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHAS_NUMPY.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 21: try:
			πF.SetLineno(21)
			πF.PushCheckpoint(5)
			// line 22: import builtins
			πF.SetLineno(22)
			if πTemp002, πE = πg.ImportModule(πF, "builtins"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßbuiltins.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label4
		Label5:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp003, πTemp004 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label6
			}
			πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
			continue
			// line 23: except ImportError:
			πF.SetLineno(23)
		Label6:
			// line 24: import __builtin__ as builtins
			πF.SetLineno(24)
			if πTemp002, πE = πg.ImportModule(πF, "__builtin__"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßbuiltins.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label4
		Label4:
			// line 27: getrefcount = getattr(sys, 'getrefcount', lambda x:0)
			πF.SetLineno(27)
			πTemp002 = πF.MakeArgs(3)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			πTemp002[1] = ßgetrefcount.ToObject()
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 27: getrefcount = getattr(sys, 'getrefcount', lambda x:0)
					πF.SetLineno(27)
					πR = πg.NewInt(0).ToObject()
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			πTemp002[2] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßgetrefcount.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 32: memo = {}
			πF.SetLineno(32)
			πTemp008 = πg.NewDict()
			πTemp001 = πTemp008.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmemo.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 33: id_to_obj = {}
			πF.SetLineno(33)
			πTemp008 = πg.NewDict()
			πTemp001 = πTemp008.ToObject()
			if πE = πF.Globals().SetItem(πF, ßid_to_obj.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 35: builtins_types = set((str, list, dict, set, frozenset, int))
			πF.SetLineno(35)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple6(πTemp007, πTemp009, πTemp010, πTemp011, πTemp012, πTemp013).ToObject()
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßbuiltins_types.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 36: dont_memo = set(id(i) for i in (memo, sys.modules, sys.path_importer_cache,
			πF.SetLineno(36)
			πTemp002 = πF.MakeArgs(1)
			πTemp006 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
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
						if πTemp003, πE = πg.ResolveGlobal(πF, ßmemo); πE != nil {
							continue
						}
						if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmodules, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßpath_importer_cache, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
							continue
						}
						if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßenviron, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πg.ResolveGlobal(πF, ßid_to_obj); πE != nil {
							continue
						}
						πTemp002 = πg.NewTuple5(πTemp003, πTemp005, πTemp006, πTemp007, πTemp004).ToObject()
						if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
							continue
						}
						πF.PushCheckpoint(2)
						πTemp008 = false
					Label1:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp008 {
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
							πTemp009 = !isStop
						} else {
							πTemp009 = true
							µi = πTemp002
						}
						if πE != nil || !πTemp009 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 36: dont_memo = set(id(i) for i in (memo, sys.modules, sys.path_importer_cache,
						πF.SetLineno(36)
						πTemp010 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
							continue
						}
						πTemp010[0] = µi
						if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp010)
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
			if πTemp007, πE = πTemp001.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßdont_memo.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 40: def get_attrs(obj):
			πF.SetLineno(40)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "obj", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("get_attrs", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
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
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 41: """
					πF.SetLineno(41)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßbuiltins_types); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp007).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp006, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp008 == πTemp006).ToObject()
					πTemp003 = πTemp005
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßbuiltins_types); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, πTemp006, µobj); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp009).ToObject()
					πTemp003 = πTemp005
				Label2:
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 44: if type(obj) in builtins_types \
					πF.SetLineno(44)
				Label3:
					// line 46: return
					πF.SetLineno(46)
					πR = πg.None
					continue
					goto Label4
				Label4:
					// line 47: try:
					πF.SetLineno(47)
					πF.PushCheckpoint(6)
					// line 48: return obj.__dict__
					πF.SetLineno(48)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobj, ß__dict__, nil); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					πF.PopCheckpoint()
					goto Label5
				Label6:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp010, πTemp011 = πF.ExcInfo()
					goto Label7
					// line 49: except:
					πF.SetLineno(49)
				Label7:
					// line 50: return
					πF.SetLineno(50)
					πR = πg.None
					continue
					πF.RestoreExc(nil, nil)
					goto Label5
				Label5:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßget_attrs.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 41: """
			πF.SetLineno(41)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n    Gets all the attributes of an object though its __dict__ or return None\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßget_attrs); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 53: def get_seq(obj, cache={str: False, frozenset: False, list: True, set: True,
			πF.SetLineno(53)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "obj", Def: nil}
			πTemp008 = πg.NewDict()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πTemp010, πTemp011); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πTemp010, πTemp011); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πTemp010, πTemp011); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πTemp010, πTemp011); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πTemp010, πTemp011); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πTemp010, πTemp011); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πTemp010, πTemp011); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßModuleType, nil); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πTemp011, πTemp010); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßFunctionType, nil); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πTemp011, πTemp010); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßBuiltinFunctionType, nil); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πTemp011, πTemp010); πE != nil {
				continue
			}
			πTemp010 = πTemp008.ToObject()
			πTemp006[1] = πg.Param{Name: "cache", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("get_seq", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µcache *πg.Object = πArgs[1]; _ = µcache
				var µo_type *πg.Object = πg.UnboundLocal; _ = µo_type
				var µhsattr *πg.Object = πg.UnboundLocal; _ = µhsattr
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.BaseException
				_ = πTemp002
				var πTemp003 *πg.Traceback
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []*πg.Object
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
				var πTemp011 bool
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 57: """
					πF.SetLineno(57)
					// line 60: try:
					πF.SetLineno(60)
					πF.PushCheckpoint(2)
					// line 61: o_type = obj.__class__
					πF.SetLineno(61)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobj, ß__class__, nil); πE != nil {
						continue
					}
					µo_type = πTemp001
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp002, πTemp003 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp002.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					πE = πF.Raise(πTemp002.ToObject(), nil, πTemp003.ToObject())
					continue
					// line 62: except AttributeError:
					πF.SetLineno(62)
				Label3:
					// line 63: o_type = type(obj)
					πF.SetLineno(63)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005[0] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µo_type = πTemp006
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 64: hsattr = hasattr
					πF.SetLineno(64)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					µhsattr = πTemp001
					if πE = πg.CheckLocal(πF, µo_type, "o_type"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcache, "cache"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µcache, µo_type); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßHAS_NUMPY); πE != nil {
						continue
					}
					πTemp001 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µo_type, "o_type"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßndarray, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßma, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp010, ßcore, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßMaskedConstant, nil); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple2(πTemp009, πTemp010).ToObject()
					if πTemp011, πE = πg.Contains(πF, πTemp007, µo_type); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(πTemp011).ToObject()
					πTemp001 = πTemp006
				Label5:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005[0] = µobj
					πTemp005[1] = ß__contains__.ToObject()
					if πE = πg.CheckLocal(πF, µhsattr, "hsattr"); πE != nil {
						continue
					}
					if πTemp006, πE = µhsattr.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label7
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005[0] = µobj
					πTemp005[1] = ß__iter__.ToObject()
					if πE = πg.CheckLocal(πF, µhsattr, "hsattr"); πE != nil {
						continue
					}
					if πTemp006, πE = µhsattr.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label7
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005[0] = µobj
					πTemp005[1] = ß__len__.ToObject()
					if πE = πg.CheckLocal(πF, µhsattr, "hsattr"); πE != nil {
						continue
					}
					if πTemp006, πE = µhsattr.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label7
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µo_type, "o_type"); πE != nil {
						continue
					}
					πTemp005[0] = µo_type
					πTemp005[1] = ß__contains__.ToObject()
					if πE = πg.CheckLocal(πF, µhsattr, "hsattr"); πE != nil {
						continue
					}
					if πTemp006, πE = µhsattr.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label7
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µo_type, "o_type"); πE != nil {
						continue
					}
					πTemp005[0] = µo_type
					πTemp005[1] = ß__iter__.ToObject()
					if πE = πg.CheckLocal(πF, µhsattr, "hsattr"); πE != nil {
						continue
					}
					if πTemp006, πE = µhsattr.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label7
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µo_type, "o_type"); πE != nil {
						continue
					}
					πTemp005[0] = µo_type
					πTemp005[1] = ß__len__.ToObject()
					if πE = πg.CheckLocal(πF, µhsattr, "hsattr"); πE != nil {
						continue
					}
					if πTemp006, πE = µhsattr.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp006
				Label7:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					goto Label9
					// line 65: if o_type in cache:
					πF.SetLineno(65)
				Label4:
					if πE = πg.CheckLocal(πF, µo_type, "o_type"); πE != nil {
						continue
					}
					πTemp001 = µo_type
					if πE = πg.CheckLocal(πF, µcache, "cache"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µcache, πTemp001); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label11
					}
					goto Label12
					// line 66: if cache[o_type]:
					πF.SetLineno(66)
				Label11:
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005[0] = µobj
					πTemp005[1] = ßcopy.ToObject()
					if πE = πg.CheckLocal(πF, µhsattr, "hsattr"); πE != nil {
						continue
					}
					if πTemp001, πE = µhsattr.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label13
					}
					goto Label14
					// line 67: if hsattr(obj, "copy"):
					πF.SetLineno(67)
				Label13:
					// line 68: return obj.copy()
					πF.SetLineno(68)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobj, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πR = πTemp006
					continue
					goto Label14
				Label14:
					// line 69: return obj
					πF.SetLineno(69)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πR = µobj
					continue
					goto Label12
				Label12:
					goto Label10
					// line 70: elif HAS_NUMPY and o_type in (numpy.ndarray, numpy.ma.core.MaskedConstant):
					πF.SetLineno(70)
				Label6:
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µobj, ßshape, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µobj, ßsize, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp006
				Label15:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label16
					}
					goto Label17
					// line 71: if obj.shape and obj.size:
					πF.SetLineno(71)
				Label16:
					// line 72: return obj
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πR = µobj
					continue
					goto Label18
				Label17:
					// line 74: return []
					πF.SetLineno(74)
					πTemp005 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp005...).ToObject()
					πR = πTemp001
					continue
					goto Label18
				Label18:
					goto Label10
					// line 75: elif hsattr(obj, "__contains__") and hsattr(obj, "__iter__") \
					πF.SetLineno(75)
				Label8:
					// line 78: cache[o_type] = True
					πF.SetLineno(78)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcache, "cache"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µo_type, "o_type"); πE != nil {
						continue
					}
					πTemp007 = µo_type
					if πE = πg.SetItem(πF, µcache, πTemp007, πTemp006); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005[0] = µobj
					πTemp005[1] = ßcopy.ToObject()
					if πE = πg.CheckLocal(πF, µhsattr, "hsattr"); πE != nil {
						continue
					}
					if πTemp001, πE = µhsattr.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label19
					}
					goto Label20
					// line 79: if hsattr(obj, "copy"):
					πF.SetLineno(79)
				Label19:
					// line 80: return obj.copy()
					πF.SetLineno(80)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobj, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πR = πTemp006
					continue
					goto Label20
				Label20:
					// line 81: return obj
					πF.SetLineno(81)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πR = µobj
					continue
					goto Label10
				Label9:
					// line 83: cache[o_type] = False
					πF.SetLineno(83)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcache, "cache"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µo_type, "o_type"); πE != nil {
						continue
					}
					πTemp007 = µo_type
					if πE = πg.SetItem(πF, µcache, πTemp007, πTemp006); πE != nil {
						continue
					}
					// line 84: return None
					πF.SetLineno(84)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label10
				Label10:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßget_seq.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 57: """
			πF.SetLineno(57)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n    Gets all the items in a sequence or return None\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßget_seq); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 87: def memorise(obj, force=False):
			πF.SetLineno(87)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "force", Def: πTemp011}
			πTemp010 = πg.NewFunction(πg.NewCode("memorise", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µforce *πg.Object = πArgs[1]; _ = µforce
				var µobj_id *πg.Object = πg.UnboundLocal; _ = µobj_id
				var µid_ *πg.Object = πg.UnboundLocal; _ = µid_
				var µg *πg.Object = πg.UnboundLocal; _ = µg
				var µattrs_id *πg.Object = πg.UnboundLocal; _ = µattrs_id
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µseq_id *πg.Object = πg.UnboundLocal; _ = µseq_id
				var µmem *πg.Object = πg.UnboundLocal; _ = µmem
				var πTemp001 []*πg.Object
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 []πg.Param
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
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
					default: panic("unexpected function state")
					}
					// line 88: """
					πF.SetLineno(88)
					// line 93: obj_id = id(obj)
					πF.SetLineno(93)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µobj_id = πTemp003
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßmemo); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, πTemp007, µobj_id); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(πTemp008).ToObject()
					πTemp003 = πTemp006
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µforce); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(!πTemp008).ToObject()
					πTemp003 = πTemp006
				Label2:
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßdont_memo); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, πTemp006, µobj_id); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp005).ToObject()
					πTemp002 = πTemp003
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 94: if obj_id in memo and not force or obj_id in dont_memo:
					πF.SetLineno(94)
				Label3:
					// line 95: return
					πF.SetLineno(95)
					πR = πg.None
					continue
					goto Label4
				Label4:
					// line 96: id_ = id
					πF.SetLineno(96)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					µid_ = πTemp002
					// line 97: g = get_attrs(obj)
					πF.SetLineno(97)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßget_attrs); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µg = πTemp003
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µg == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 98: if g is None:
					πF.SetLineno(98)
				Label5:
					// line 99: attrs_id = None
					πF.SetLineno(99)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µattrs_id = πTemp002
					goto Label7
				Label6:
					// line 101: attrs_id = dict((key,id_(value)) for key, value in g.items())
					πF.SetLineno(101)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
								if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µg, ßitems, nil); πE != nil {
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
									µkey = πTemp003
									µvalue = πTemp006
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 101: attrs_id = dict((key,id_(value)) for key, value in g.items())
								πF.SetLineno(101)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp007 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp007[0] = µvalue
								if πE = πg.CheckLocal(πF, µid_, "id_"); πE != nil {
									continue
								}
								if πTemp003, πE = µid_.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp002 = πg.NewTuple2(µkey, πTemp003).ToObject()
								πF.PushCheckpoint(4)
								return πTemp002, nil
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
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µattrs_id = πTemp006
					goto Label7
				Label7:
					// line 103: s = get_seq(obj)
					πF.SetLineno(103)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßget_seq); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µs = πTemp006
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µs == πTemp006).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					πTemp001[1] = ßitems.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					πTemp001[1] = ß__len__.ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label10
					}
					goto Label11
					// line 104: if s is None:
					πF.SetLineno(104)
				Label8:
					// line 105: seq_id = None
					πF.SetLineno(105)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µseq_id = πTemp003
					goto Label12
					// line 106: elif hasattr(s, "items"):
					πF.SetLineno(106)
				Label9:
					// line 107: seq_id = dict((id_(key),id_(value)) for key, value in s.items())
					πF.SetLineno(107)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µs, ßitems, nil); πE != nil {
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
									µkey = πTemp003
									µvalue = πTemp006
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 107: seq_id = dict((id_(key),id_(value)) for key, value in s.items())
								πF.SetLineno(107)
								πTemp007 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp007[0] = µkey
								if πE = πg.CheckLocal(πF, µid_, "id_"); πE != nil {
									continue
								}
								if πTemp003, πE = µid_.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp007 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp007[0] = µvalue
								if πE = πg.CheckLocal(πF, µid_, "id_"); πE != nil {
									continue
								}
								if πTemp006, πE = µid_.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
								πF.PushCheckpoint(4)
								return πTemp002, nil
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
					πTemp001[0] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µseq_id = πTemp007
					goto Label12
					// line 108: elif not hasattr(s, "__len__"): #XXX: avoid TypeError from unexpected case
					πF.SetLineno(108)
				Label10:
					// line 109: seq_id = None
					πF.SetLineno(109)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µseq_id = πTemp006
					goto Label12
				Label11:
					// line 111: seq_id = [id_(i) for i in s]
					πF.SetLineno(111)
					πTemp009 = make([]πg.Param, 0)
					πTemp007 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µs); πE != nil {
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
								// line 111: seq_id = [id_(i) for i in s]
								πF.SetLineno(111)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp005[0] = µi
								if πE = πg.CheckLocal(πF, µid_, "id_"); πE != nil {
									continue
								}
								if πTemp004, πE = µid_.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(4)
								return πTemp004, nil
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
					if πTemp010, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ListType.Call(πF, πg.Args{πTemp010}, nil); πE != nil {
						continue
					}
					µseq_id = πTemp006
					goto Label12
				Label12:
					// line 113: memo[obj_id] = attrs_id, seq_id
					πF.SetLineno(113)
					if πE = πg.CheckLocal(πF, µattrs_id, "attrs_id"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µseq_id, "seq_id"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(µattrs_id, µseq_id).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp006); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveGlobal(πF, ßmemo); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					πTemp012 = µobj_id
					if πE = πg.SetItem(πF, πTemp011, πTemp012, πTemp010); πE != nil {
						continue
					}
					// line 114: id_to_obj[obj_id] = obj
					πF.SetLineno(114)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µobj); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveGlobal(πF, ßid_to_obj); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					πTemp011 = µobj_id
					if πE = πg.SetItem(πF, πTemp010, πTemp011, πTemp006); πE != nil {
						continue
					}
					// line 115: mem = memorise
					πF.SetLineno(115)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßmemorise); πE != nil {
						continue
					}
					µmem = πTemp006
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(µg != πTemp010).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label13
					}
					goto Label14
					// line 116: if g is not None:
					πF.SetLineno(116)
				Label13:
					// line 117: [mem(value) for key, value in g.items()]
					πF.SetLineno(117)
					πTemp009 = make([]πg.Param, 0)
					πTemp010 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
								if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µg, ßitems, nil); πE != nil {
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
									µkey = πTemp003
									µvalue = πTemp006
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 117: [mem(value) for key, value in g.items()]
								πF.SetLineno(117)
								πTemp007 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp007[0] = µvalue
								if πE = πg.CheckLocal(πF, µmem, "mem"); πE != nil {
									continue
								}
								if πTemp002, πE = µmem.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πF.PushCheckpoint(4)
								return πTemp002, nil
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
					if πTemp011, πE = πTemp010.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ListType.Call(πF, πg.Args{πTemp011}, nil); πE != nil {
						continue
					}
					goto Label14
				Label14:
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(µs != πTemp011).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label15
					}
					goto Label16
					// line 119: if s is not None:
					πF.SetLineno(119)
				Label15:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					πTemp001[1] = ßitems.ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp011); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label17
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					πTemp001[1] = ß__len__.ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp011); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label18
					}
					goto Label19
					// line 120: if hasattr(s, "items"):
					πF.SetLineno(120)
				Label17:
					// line 121: [(mem(key), mem(item))
					πF.SetLineno(121)
					πTemp009 = make([]πg.Param, 0)
					πTemp011 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µitem *πg.Object = πg.UnboundLocal; _ = µitem
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
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µs, ßitems, nil); πE != nil {
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
									µkey = πTemp003
									µitem = πTemp006
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 121: [(mem(key), mem(item))
								πF.SetLineno(121)
								πTemp007 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp007[0] = µkey
								if πE = πg.CheckLocal(πF, µmem, "mem"); πE != nil {
									continue
								}
								if πTemp003, πE = µmem.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp007 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
									continue
								}
								πTemp007[0] = µitem
								if πE = πg.CheckLocal(πF, µmem, "mem"); πE != nil {
									continue
								}
								if πTemp006, πE = µmem.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
								πF.PushCheckpoint(4)
								return πTemp002, nil
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
					if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ListType.Call(πF, πg.Args{πTemp012}, nil); πE != nil {
						continue
					}
					goto Label20
					// line 124: if hasattr(s, '__len__'):
					πF.SetLineno(124)
				Label18:
					// line 125: [mem(item) for item in s]
					πF.SetLineno(125)
					πTemp009 = make([]πg.Param, 0)
					πTemp012 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µitem *πg.Object = πg.UnboundLocal; _ = µitem
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
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µs); πE != nil {
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
									µitem = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 125: [mem(item) for item in s]
								πF.SetLineno(125)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
									continue
								}
								πTemp005[0] = µitem
								if πE = πg.CheckLocal(πF, µmem, "mem"); πE != nil {
									continue
								}
								if πTemp004, πE = µmem.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(4)
								return πTemp004, nil
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
					if πTemp013, πE = πTemp012.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ListType.Call(πF, πg.Args{πTemp013}, nil); πE != nil {
						continue
					}
					goto Label20
				Label19:
					// line 126: else: mem(s)
					πF.SetLineno(126)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πE = πg.CheckLocal(πF, µmem, "mem"); πE != nil {
						continue
					}
					if πTemp006, πE = µmem.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label20
				Label20:
					goto Label16
				Label16:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmemorise.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 88: """
			πF.SetLineno(88)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("\n    Adds an object to the memo, and recursively adds all the objects\n    attributes, and if it is a container, its items. Use force=True to update\n    an object already in the memo. Updating is not recursively done.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßmemorise); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 129: def release_gone():
			πF.SetLineno(129)
			πTemp006 = make([]πg.Param, 0)
			πTemp011 = πg.NewFunction(πg.NewCode("release_gone", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µitop *πg.Object = πg.UnboundLocal; _ = µitop
				var µmp *πg.Object = πg.UnboundLocal; _ = µmp
				var µsrc *πg.Object = πg.UnboundLocal; _ = µsrc
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 130: itop, mp, src = id_to_obj.pop, memo.pop, getrefcount
					πF.SetLineno(130)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid_to_obj); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmemo); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetrefcount); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(πTemp003, πTemp004, πTemp002).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µitop = πTemp002
					µmp = πTemp003
					µsrc = πTemp004
					// line 131: [(itop(id_), mp(id_)) for id_, obj in list(id_to_obj.items())
					πF.SetLineno(131)
					πTemp005 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µid_ *πg.Object = πg.UnboundLocal; _ = µid_
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
								if πTemp003, πE = πg.ResolveGlobal(πF, ßid_to_obj); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßitems, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
									continue
								}
								πTemp002[0] = πTemp003
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
										continue
									}
									µid_ = πTemp004
									µobj = πTemp007
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp002[0] = µobj
								if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
									continue
								}
								if πTemp004, πE = µsrc.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp003, πE = πg.LT(πF, πTemp004, πg.NewInt(4).ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 131: [(itop(id_), mp(id_)) for id_, obj in list(id_to_obj.items())
								πF.SetLineno(131)
							Label4:
								// line 131: [(itop(id_), mp(id_)) for id_, obj in list(id_to_obj.items())
								πF.SetLineno(131)
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µid_, "id_"); πE != nil {
									continue
								}
								πTemp002[0] = µid_
								if πE = πg.CheckLocal(πF, µitop, "itop"); πE != nil {
									continue
								}
								if πTemp004, πE = µitop.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µid_, "id_"); πE != nil {
									continue
								}
								πTemp002[0] = µid_
								if πE = πg.CheckLocal(πF, µmp, "mp"); πE != nil {
									continue
								}
								if πTemp007, πE = µmp.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πTemp003 = πg.NewTuple2(πTemp004, πTemp007).ToObject()
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
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
						continue
					}
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßrelease_gone.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 135: def whats_changed(obj, seen=None, simple=False, first=True):
			πF.SetLineno(135)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "seen", Def: πTemp013}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "simple", Def: πTemp013}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "first", Def: πTemp013}
			πTemp012 = πg.NewFunction(πg.NewCode("whats_changed", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µseen *πg.Object = πArgs[1]; _ = µseen
				var µsimple *πg.Object = πArgs[2]; _ = µsimple
				var µfirst *πg.Object = πArgs[3]; _ = µfirst
				var µobj_id *πg.Object = πg.UnboundLocal; _ = µobj_id
				var µchngd *πg.Object = πg.UnboundLocal; _ = µchngd
				var µid_ *πg.Object = πg.UnboundLocal; _ = µid_
				var µattrs *πg.Object = πg.UnboundLocal; _ = µattrs
				var µchanged *πg.Object = πg.UnboundLocal; _ = µchanged
				var µobj_attrs *πg.Object = πg.UnboundLocal; _ = µobj_attrs
				var µobj_get *πg.Object = πg.UnboundLocal; _ = µobj_get
				var µkey *πg.Object = πg.UnboundLocal; _ = µkey
				var µo *πg.Object = πg.UnboundLocal; _ = µo
				var µitems *πg.Object = πg.UnboundLocal; _ = µitems
				var µseq_diff *πg.Object = πg.UnboundLocal; _ = µseq_diff
				var µobj_seq *πg.Object = πg.UnboundLocal; _ = µobj_seq
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µj *πg.Object = πg.UnboundLocal; _ = µj
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Dict
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 []πg.Param
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 35: goto Label35
					case 36: goto Label36
					case 41: goto Label41
					case 42: goto Label42
					case 22: goto Label22
					case 23: goto Label23
					default: panic("unexpected function state")
					}
					// line 136: """
					πF.SetLineno(136)
					if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µfirst); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 144: if first:
					πF.SetLineno(144)
				Label1:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßbuiltins); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Contains(πF, πTemp004, ß_.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label3
					}
					goto Label4
					// line 146: if "_" in builtins.__dict__:
					πF.SetLineno(146)
				Label3:
					// line 147: del builtins._
					πF.SetLineno(147)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbuiltins); πE != nil {
						continue
					}
					if πE = πg.DelAttr(πF, πTemp002, ß_); πE != nil {
						continue
					}
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µseen == πTemp003).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label5
					}
					goto Label6
					// line 148: if seen is None:
					πF.SetLineno(148)
				Label5:
					// line 149: seen = {}
					πF.SetLineno(149)
					πTemp005 = πg.NewDict()
					πTemp002 = πTemp005.ToObject()
					µseen = πTemp002
					goto Label6
				Label6:
					goto Label2
				Label2:
					// line 151: obj_id = id(obj)
					πF.SetLineno(151)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					µobj_id = πTemp003
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Contains(πF, µseen, µobj_id); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label7
					}
					goto Label8
					// line 153: if obj_id in seen:
					πF.SetLineno(153)
				Label7:
					if πE = πg.CheckLocal(πF, µsimple, "simple"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µsimple); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label9
					}
					goto Label10
					// line 154: if simple:
					πF.SetLineno(154)
				Label9:
					// line 155: return any(seen[obj_id])
					πF.SetLineno(155)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					πTemp002 = µobj_id
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µseen, πTemp002); πE != nil {
						continue
					}
					πTemp006[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßany); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πR = πTemp003
					continue
					goto Label10
				Label10:
					// line 156: return seen[obj_id]
					πF.SetLineno(156)
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					πTemp002 = µobj_id
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µseen, πTemp002); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdont_memo); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Contains(πF, πTemp003, µobj_id); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmemo); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Contains(πF, πTemp003, µobj_id); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label12
					}
					goto Label13
					// line 159: if obj_id in dont_memo:
					πF.SetLineno(159)
				Label11:
					// line 160: seen[obj_id] = [{}, False]
					πF.SetLineno(160)
					πTemp006 = make([]*πg.Object, 2)
					πTemp005 = πg.NewDict()
					πTemp002 = πTemp005.ToObject()
					πTemp006[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp006[1] = πTemp002
					πTemp002 = πg.NewList(πTemp006...).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					πTemp004 = µobj_id
					if πE = πg.SetItem(πF, µseen, πTemp004, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsimple, "simple"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µsimple); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label14
					}
					goto Label15
					// line 161: if simple:
					πF.SetLineno(161)
				Label14:
					// line 162: return False
					πF.SetLineno(162)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label15
				Label15:
					// line 163: return seen[obj_id]
					πF.SetLineno(163)
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					πTemp002 = µobj_id
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µseen, πTemp002); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label13
					// line 164: elif obj_id not in memo:
					πF.SetLineno(164)
				Label12:
					if πE = πg.CheckLocal(πF, µsimple, "simple"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µsimple); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label16
					}
					goto Label17
					// line 165: if simple:
					πF.SetLineno(165)
				Label16:
					// line 166: return True
					πF.SetLineno(166)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label18
				Label17:
					πTemp006 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp002, πE = πg.Add(πF, πg.NewStr("Object not memorised ").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp006[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					// line 168: raise RuntimeError("Object not memorised " + str(obj))
					πF.SetLineno(168)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label18
				Label18:
					goto Label13
				Label13:
					// line 170: seen[obj_id] = ({}, False)
					πF.SetLineno(170)
					πTemp005 = πg.NewDict()
					πTemp003 = πTemp005.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					πTemp004 = µobj_id
					if πE = πg.SetItem(πF, µseen, πTemp004, πTemp003); πE != nil {
						continue
					}
					// line 172: chngd = whats_changed
					πF.SetLineno(172)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßwhats_changed); πE != nil {
						continue
					}
					µchngd = πTemp002
					// line 173: id_ = id
					πF.SetLineno(173)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					µid_ = πTemp002
					// line 176: attrs = get_attrs(obj)
					πF.SetLineno(176)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßget_attrs); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					µattrs = πTemp003
					if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µattrs == πTemp003).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label19
					}
					goto Label20
					// line 177: if attrs is None:
					πF.SetLineno(177)
				Label19:
					// line 178: changed = {}
					πF.SetLineno(178)
					πTemp005 = πg.NewDict()
					πTemp002 = πTemp005.ToObject()
					µchanged = πTemp002
					goto Label21
				Label20:
					// line 180: obj_attrs = memo[obj_id][0]
					πF.SetLineno(180)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					πTemp004 = µobj_id
					if πTemp009, πE = πg.ResolveGlobal(πF, ßmemo); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
						continue
					}
					µobj_attrs = πTemp003
					// line 181: obj_get = obj_attrs.get
					πF.SetLineno(181)
					if πE = πg.CheckLocal(πF, µobj_attrs, "obj_attrs"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj_attrs, ßget, nil); πE != nil {
						continue
					}
					µobj_get = πTemp002
					// line 182: changed = dict((key,None) for key in obj_attrs if key not in attrs)
					πF.SetLineno(182)
					πTemp006 = πF.MakeArgs(1)
					πTemp010 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
								if πE = πg.CheckLocal(πF, µobj_attrs, "obj_attrs"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µobj_attrs); πE != nil {
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
									µkey = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Contains(πF, µattrs, µkey); πE != nil {
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
								// line 182: changed = dict((key,None) for key in obj_attrs if key not in attrs)
								πF.SetLineno(182)
							Label4:
								// line 182: changed = dict((key,None) for key in obj_attrs if key not in attrs)
								πF.SetLineno(182)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp004 = πg.NewTuple2(µkey, πTemp005).ToObject()
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
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					µchanged = πTemp004
					if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µattrs, ßitems, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, πTemp008); πE != nil {
						continue
					}
					πF.PushCheckpoint(23)
					πTemp001 = false
				Label22:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp001 {
						πF.PopCheckpoint()
						goto Label24
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
							continue
						}
						µkey = πTemp008
						µo = πTemp009
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(22)            
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp006[0] = µo
					if πE = πg.CheckLocal(πF, µid_, "id_"); πE != nil {
						continue
					}
					if πTemp009, πE = µid_.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp006[0] = µkey
					if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp006[1] = πTemp012
					if πE = πg.CheckLocal(πF, µobj_get, "obj_get"); πE != nil {
						continue
					}
					if πTemp012, πE = µobj_get.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp008, πE = πg.NE(πF, πTemp009, πTemp012); πE != nil {
						continue
					}
					πTemp004 = πTemp008
					if πTemp011, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label25
					}
					πTemp006 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp006[0] = µo
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					πTemp006[1] = µseen
					if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006[2] = πTemp008
					if πTemp008, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp006[3] = πTemp008
					if πE = πg.CheckLocal(πF, µchngd, "chngd"); πE != nil {
						continue
					}
					if πTemp008, πE = µchngd.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp004 = πTemp008
				Label25:
					if πTemp011, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label26
					}
					goto Label27
					// line 184: if id_(o) != obj_get(key, None) or chngd(o, seen, True, False):
					πF.SetLineno(184)
				Label26:
					// line 185: changed[key] = o
					πF.SetLineno(185)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µo); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µchanged, "changed"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp008 = µkey
					if πE = πg.SetItem(πF, µchanged, πTemp008, πTemp004); πE != nil {
						continue
					}
					goto Label27
				Label27:
					continue
				Label23:
					if πE != nil || πR != nil {
						continue
					}
				Label24:
					goto Label21
				Label21:
					// line 188: items = get_seq(obj)
					πF.SetLineno(188)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßget_seq); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					µitems = πTemp004
					// line 189: seq_diff = False
					πF.SetLineno(189)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µseq_diff = πTemp003
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(µitems != πTemp008).ToObject()
					πTemp003 = πTemp004
					if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp001 {
						goto Label28
					}
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					πTemp006[0] = µitems
					πTemp006[1] = ß__len__.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp003 = πTemp008
				Label28:
					if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label29
					}
					goto Label30
					// line 190: if (items is not None) and (hasattr(items, '__len__')):
					πF.SetLineno(190)
				Label29:
					// line 191: obj_seq = memo[obj_id][1]
					πF.SetLineno(191)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					πTemp008 = µobj_id
					if πTemp012, πE = πg.ResolveGlobal(πF, ßmemo); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, πTemp012, πTemp008); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
						continue
					}
					µobj_seq = πTemp004
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					πTemp006[0] = µitems
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj_seq, "obj_seq"); πE != nil {
						continue
					}
					πTemp006[0] = µobj_seq
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp003, πE = πg.NE(πF, πTemp008, πTemp009); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label31
					}
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006[0] = µobj
					πTemp006[1] = ßitems.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp001, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label32
					}
					goto Label33
					// line 192: if (len(items) != len(obj_seq)):
					πF.SetLineno(192)
				Label31:
					// line 193: seq_diff = True
					πF.SetLineno(193)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µseq_diff = πTemp003
					goto Label34
					// line 194: elif hasattr(obj, "items"):  # dict type obj
					πF.SetLineno(194)
				Label32:
					// line 195: obj_get = obj_seq.get
					πF.SetLineno(195)
					if πE = πg.CheckLocal(πF, µobj_seq, "obj_seq"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj_seq, ßget, nil); πE != nil {
						continue
					}
					µobj_get = πTemp003
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µitems, ßitems, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, πTemp008); πE != nil {
						continue
					}
					πF.PushCheckpoint(36)
					πTemp001 = false
				Label35:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp001 {
						πF.PopCheckpoint()
						goto Label37
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
							continue
						}
						µkey = πTemp008
						µitem = πTemp009
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(35)            
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp006[0] = µitem
					if πE = πg.CheckLocal(πF, µid_, "id_"); πE != nil {
						continue
					}
					if πTemp009, πE = µid_.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp006 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp007[0] = µkey
					if πE = πg.CheckLocal(πF, µid_, "id_"); πE != nil {
						continue
					}
					if πTemp012, πE = µid_.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp006[0] = πTemp012
					if πE = πg.CheckLocal(πF, µobj_get, "obj_get"); πE != nil {
						continue
					}
					if πTemp012, πE = µobj_get.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp008, πE = πg.NE(πF, πTemp009, πTemp012); πE != nil {
						continue
					}
					πTemp004 = πTemp008
					if πTemp011, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label38
					}
					πTemp006 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp006[0] = µkey
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					πTemp006[1] = µseen
					if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006[2] = πTemp008
					if πTemp008, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp006[3] = πTemp008
					if πE = πg.CheckLocal(πF, µchngd, "chngd"); πE != nil {
						continue
					}
					if πTemp008, πE = µchngd.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp004 = πTemp008
					if πTemp011, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label38
					}
					πTemp006 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp006[0] = µitem
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					πTemp006[1] = µseen
					if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006[2] = πTemp008
					if πTemp008, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp006[3] = πTemp008
					if πE = πg.CheckLocal(πF, µchngd, "chngd"); πE != nil {
						continue
					}
					if πTemp008, πE = µchngd.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp004 = πTemp008
				Label38:
					if πTemp011, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label39
					}
					goto Label40
					// line 197: if id_(item) != obj_get(id_(key)) \
					πF.SetLineno(197)
				Label39:
					// line 200: seq_diff = True
					πF.SetLineno(200)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µseq_diff = πTemp004
					// line 201: break
					πF.SetLineno(201)
					πTemp001 = true
					continue
					goto Label40
				Label40:
					continue
				Label36:
					if πE != nil || πR != nil {
						continue
					}
				Label37:
					goto Label34
				Label33:
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					πTemp006[0] = µitems
					if πE = πg.CheckLocal(πF, µobj_seq, "obj_seq"); πE != nil {
						continue
					}
					πTemp006[1] = µobj_seq
					if πTemp004, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp003, πE = πg.Iter(πF, πTemp008); πE != nil {
						continue
					}
					πF.PushCheckpoint(42)
					πTemp001 = false
				Label41:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp001 {
						πF.PopCheckpoint()
						goto Label43
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
							continue
						}
						µi = πTemp008
						µj = πTemp009
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(41)            
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp006[0] = µi
					if πE = πg.CheckLocal(πF, µid_, "id_"); πE != nil {
						continue
					}
					if πTemp009, πE = µid_.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.NE(πF, πTemp009, µj); πE != nil {
						continue
					}
					πTemp004 = πTemp008
					if πTemp011, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label44
					}
					πTemp006 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp006[0] = µi
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					πTemp006[1] = µseen
					if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006[2] = πTemp008
					if πTemp008, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp006[3] = πTemp008
					if πE = πg.CheckLocal(πF, µchngd, "chngd"); πE != nil {
						continue
					}
					if πTemp008, πE = µchngd.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp004 = πTemp008
				Label44:
					if πTemp011, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label45
					}
					goto Label46
					// line 204: if id_(i) != j or chngd(i, seen, True, False):
					πF.SetLineno(204)
				Label45:
					// line 205: seq_diff = True
					πF.SetLineno(205)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µseq_diff = πTemp004
					// line 206: break
					πF.SetLineno(206)
					πTemp001 = true
					continue
					goto Label46
				Label46:
					continue
				Label42:
					if πE != nil || πR != nil {
						continue
					}
				Label43:
					goto Label34
				Label34:
					goto Label30
				Label30:
					// line 207: seen[obj_id] = changed, seq_diff
					πF.SetLineno(207)
					if πE = πg.CheckLocal(πF, µchanged, "changed"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µseq_diff, "seq_diff"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µchanged, µseq_diff).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj_id, "obj_id"); πE != nil {
						continue
					}
					πTemp008 = µobj_id
					if πE = πg.SetItem(πF, µseen, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsimple, "simple"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µsimple); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label47
					}
					goto Label48
					// line 208: if simple:
					πF.SetLineno(208)
				Label47:
					// line 209: return changed or seq_diff
					πF.SetLineno(209)
					if πE = πg.CheckLocal(πF, µchanged, "changed"); πE != nil {
						continue
					}
					πTemp003 = µchanged
					if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label49
					}
					if πE = πg.CheckLocal(πF, µseq_diff, "seq_diff"); πE != nil {
						continue
					}
					πTemp003 = µseq_diff
				Label49:
					πR = πTemp003
					continue
					goto Label48
				Label48:
					// line 210: return changed, seq_diff
					πF.SetLineno(210)
					if πE = πg.CheckLocal(πF, µchanged, "changed"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µseq_diff, "seq_diff"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µchanged, µseq_diff).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßwhats_changed.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 136: """
			πF.SetLineno(136)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("\n    Check an object against the memo. Returns a list in the form\n    (attribute changes, container changed). Attribute changes is a dict of\n    attribute name to attribute value. container changed is a boolean.\n    If simple is true, just returns a boolean. None for either item means\n    that it has not been checked yet\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßwhats_changed); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
				continue
			}
			// line 213: def has_changed(*args, **kwds):
			πF.SetLineno(213)
			πTemp006 = make([]πg.Param, 0)
			πTemp013 = πg.NewFunction(πg.NewCode("has_changed", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp006, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var πTemp001 *πg.Object
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
					// line 214: kwds['simple'] = True  # ignore simple if passed in
					πF.SetLineno(214)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					πTemp003 = ßsimple.ToObject()
					if πE = πg.SetItem(πF, µkwds, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 215: return whats_changed(*args, **kwds)
					πF.SetLineno(215)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßwhats_changed); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwds); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßhas_changed.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 217: __import__ = __import__
			πF.SetLineno(217)
			if πTemp014, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__import__.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 220: def _imp(*args, **kwds):
			πF.SetLineno(220)
			πTemp006 = make([]πg.Param, 0)
			πTemp014 = πg.NewFunction(πg.NewCode("_imp", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/__diff.py", πTemp006, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var µbefore *πg.Object = πg.UnboundLocal; _ = µbefore
				var µmod *πg.Object = πg.UnboundLocal; _ = µmod
				var µafter *πg.Object = πg.UnboundLocal; _ = µafter
				var µm *πg.Object = πg.UnboundLocal; _ = µm
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
				var πTemp009 *πg.Object
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
					// line 221: """
					πF.SetLineno(221)
					// line 225: before = set(sys.modules.keys())
					πF.SetLineno(225)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmodules, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µbefore = πTemp003
					// line 226: mod = __import__(*args, **kwds)
					πF.SetLineno(226)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, nil, µargs, nil, µkwds); πE != nil {
						continue
					}
					µmod = πTemp003
					// line 227: after = set(sys.modules.keys()).difference(before)
					πF.SetLineno(227)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbefore, "before"); πE != nil {
						continue
					}
					πTemp001[0] = µbefore
					πTemp004 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmodules, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßdifference, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µafter = πTemp003
					if πE = πg.CheckLocal(πF, µafter, "after"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µafter); πE != nil {
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
						µm = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 229: memorise(sys.modules[m])
					πF.SetLineno(229)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					πTemp003 = µm
					if πTemp008, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßmodules, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp007
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmemorise); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 230: return mod
					πF.SetLineno(230)
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					πR = µmod
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_imp.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 221: """
			πF.SetLineno(221)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("\n    Replaces the default __import__, to allow a module to be memorised\n    before the user can change it\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ß_imp); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
				continue
			}
			// line 232: builtins.__import__ = _imp
			πF.SetLineno(232)
			if πTemp015, πE = πg.ResolveGlobal(πF, ß_imp); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πTemp015); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßbuiltins); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp017, ß__import__, πTemp016); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(2)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßbuiltins); πE != nil {
				continue
			}
			πTemp002[0] = πTemp015
			πTemp002[1] = ß_.ToObject()
			if πTemp015, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πTemp005, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label7
			}
			goto Label8
			// line 233: if hasattr(builtins, "_"):
			πF.SetLineno(233)
		Label7:
			// line 234: del builtins._
			πF.SetLineno(234)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßbuiltins); πE != nil {
				continue
			}
			if πE = πg.DelAttr(πF, πTemp015, ß_); πE != nil {
				continue
			}
			goto Label8
		Label8:
			if πTemp016, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßmodules, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp017, ßvalues, nil); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp015, πE = πg.Iter(πF, πTemp017); πE != nil {
				continue
			}
			πF.PushCheckpoint(10)
			πTemp005 = false
		Label9:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp005 {
				πF.PopCheckpoint()
				goto Label11
			}
			if πTemp016, πE = πg.Next(πF, πTemp015); πE != nil {
				isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
				if exc != nil {
					πE = exc
				} else if isStop {
					πE = nil
					πF.RestoreExc(nil, nil)
				}
				πTemp018 = !isStop
			} else {
				πTemp018 = true
				if πE = πF.Globals().SetItem(πF, ßmod.ToObject(), πTemp016); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp018 {
				continue
			}
			πF.PushCheckpoint(9)            
			// line 239: memorise(mod)
			πF.SetLineno(239)
			πTemp002 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßmod); πE != nil {
				continue
			}
			πTemp002[0] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßmemorise); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			continue
		Label10:
			if πE != nil || πR != nil {
				continue
			}
		Label11:
			// line 240: release_gone()
			πF.SetLineno(240)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßrelease_gone); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, nil, nil); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("__diff", Code)
}

