package pointers
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/pkg/sys"
	// _ "github.com/pygolin/stdlib/pkg/_dill"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/pointers.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßNone := πg.InternStr("None")
		ß__all__ := πg.InternStr("__all__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__len__ := πg.InternStr("__len__")
		ß_getframe := πg.InternStr("_getframe")
		ß_locate_object := πg.InternStr("_locate_object")
		ß_proxy_helper := πg.InternStr("_proxy_helper")
		ßadd := πg.InternStr("add")
		ßappend := πg.InternStr("append")
		ßat := πg.InternStr("at")
		ßchildren := πg.InternStr("children")
		ßcollect := πg.InternStr("collect")
		ßfind_chain := πg.InternStr("find_chain")
		ßgc := πg.InternStr("gc")
		ßget_referents := πg.InternStr("get_referents")
		ßget_referrers := πg.InternStr("get_referrers")
		ßhasattr := πg.InternStr("hasattr")
		ßid := πg.InternStr("id")
		ßisinstance := πg.InternStr("isinstance")
		ßparent := πg.InternStr("parent")
		ßparents := πg.InternStr("parents")
		ßpop := πg.InternStr("pop")
		ßreference := πg.InternStr("reference")
		ßrefobject := πg.InternStr("refobject")
		ßset := πg.InternStr("set")
		ßsys := πg.InternStr("sys")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 9: __all__ = ['parent', 'reference', 'at', 'parents', 'children']
			πF.SetLineno(9)
			πTemp001 = make([]*πg.Object, 5)
			πTemp001[0] = ßparent.ToObject()
			πTemp001[1] = ßreference.ToObject()
			πTemp001[2] = ßat.ToObject()
			πTemp001[3] = ßparents.ToObject()
			πTemp001[4] = ßchildren.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 11: import gc
			πF.SetLineno(11)
			if πTemp001, πE = πg.ImportModule(πF, "gc"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßgc.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 12: import sys
			πF.SetLineno(12)
			if πTemp001, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 14: from ._dill import _proxy_helper as reference
			πF.SetLineno(14)
			if πTemp001, πE = πg.ImportModule(πF, "._dill"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ß_proxy_helper); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreference.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 15: from ._dill import _locate_object as at
			πF.SetLineno(15)
			if πTemp001, πE = πg.ImportModule(πF, "._dill"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ß_locate_object); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßat.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 17: def parent(obj, objtype, ignore=()):
			πF.SetLineno(17)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "objtype", Def: nil}
			πTemp003 = πg.NewTuple0().ToObject()
			πTemp004[2] = πg.Param{Name: "ignore", Def: πTemp003}
			πTemp002 = πg.NewFunction(πg.NewCode("parent", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/pointers.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µobjtype *πg.Object = πArgs[1]; _ = µobjtype
				var µignore *πg.Object = πArgs[2]; _ = µignore
				var µdepth *πg.Object = πg.UnboundLocal; _ = µdepth
				var µchain *πg.Object = πg.UnboundLocal; _ = µchain
				var µparent *πg.Object = πg.UnboundLocal; _ = µparent
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
					// line 18: """
					πF.SetLineno(18)
					// line 29: depth = 1 #XXX: always looking for the parent (only, right?)
					πF.SetLineno(29)
					µdepth = πg.NewInt(1).ToObject()
					// line 30: chain = parents(obj, objtype, depth, ignore)
					πF.SetLineno(30)
					πTemp001 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πE = πg.CheckLocal(πF, µobjtype, "objtype"); πE != nil {
						continue
					}
					πTemp001[1] = µobjtype
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					πTemp001[2] = µdepth
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp001[3] = µignore
					if πTemp002, πE = πg.ResolveGlobal(πF, ßparents); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µchain = πTemp003
					// line 31: parent = chain.pop()
					πF.SetLineno(31)
					if πE = πg.CheckLocal(πF, µchain, "chain"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µchain, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µparent = πTemp003
					if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µparent == µobj).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 32: if parent is obj:
					πF.SetLineno(32)
				Label1:
					// line 33: return None
					πF.SetLineno(33)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 34: return parent
					πF.SetLineno(34)
					if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
						continue
					}
					πR = µparent
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßparent.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 18: """
			πF.SetLineno(18)
			// line 18: """
			πF.SetLineno(18)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n>>> listiter = iter([4,5,6,7])\n>>> obj = parent(listiter, list)\n>>> obj == [4,5,6,7]  # actually 'is', but don't have handle any longer\nTrue\n\nNOTE: objtype can be a single type (e.g. int or list) or a tuple of types.\n\nWARNING: if obj is a sequence (e.g. list), may produce unexpected results.\nParent finds *one* parent (e.g. the last member of the sequence).\n    ").ToObject()); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n>>> listiter = iter([4,5,6,7])\n>>> obj = parent(listiter, list)\n>>> obj == [4,5,6,7]  # actually 'is', but don't have handle any longer\nTrue\n\nNOTE: objtype can be a single type (e.g. int or list) or a tuple of types.\n\nWARNING: if obj is a sequence (e.g. list), may produce unexpected results.\nParent finds *one* parent (e.g. the last member of the sequence).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßparent); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 37: def parents(obj, objtype, depth=1, ignore=()): #XXX: objtype=object ?
			πF.SetLineno(37)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "objtype", Def: nil}
			πTemp004[2] = πg.Param{Name: "depth", Def: πg.NewInt(1).ToObject()}
			πTemp005 = πg.NewTuple0().ToObject()
			πTemp004[3] = πg.Param{Name: "ignore", Def: πTemp005}
			πTemp003 = πg.NewFunction(πg.NewCode("parents", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/pointers.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µobjtype *πg.Object = πArgs[1]; _ = µobjtype
				var µdepth *πg.Object = πArgs[2]; _ = µdepth
				var µignore *πg.Object = πArgs[3]; _ = µignore
				var µedge_func *πg.Object = πg.UnboundLocal; _ = µedge_func
				var µpredicate *πg.Object = πg.UnboundLocal; _ = µpredicate
				var µchain *πg.Object = πg.UnboundLocal; _ = µchain
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 38: """Find the chain of referents for obj. Chain will end with obj.
					πF.SetLineno(38)
					// line 44: edge_func = gc.get_referents # looking for refs, not back_refs
					πF.SetLineno(44)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_referents, nil); πE != nil {
						continue
					}
					µedge_func = πTemp002
					// line 45: predicate = lambda x: isinstance(x, objtype) # looking for parent type
					πF.SetLineno(45)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "x", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/pointers.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
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
							// line 45: predicate = lambda x: isinstance(x, objtype) # looking for parent type
							πF.SetLineno(45)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µobjtype, "objtype"); πE != nil {
								continue
							}
							πTemp001[1] = µobjtype
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
					µpredicate = πTemp001
					// line 47: ignore = (ignore,) if not hasattr(ignore, '__len__') else ignore
					πF.SetLineno(47)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp004[0] = µignore
					πTemp004[1] = ß__len__.ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(µignore).ToObject()
					πTemp001 = πTemp002
					goto Label2
				Label1:
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp001 = µignore
				Label2:
					µignore = πTemp001
					// line 48: ignore = (id(obj) for obj in ignore)
					πF.SetLineno(48)
					πTemp003 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/pointers.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
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
								if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µignore); πE != nil {
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
									µobj = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 48: ignore = (id(obj) for obj in ignore)
								πF.SetLineno(48)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp005[0] = µobj
								if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
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
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µignore = πTemp002
					// line 49: chain = find_chain(obj, predicate, edge_func, depth)[::-1]
					πF.SetLineno(49)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πTemp005}, nil); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µpredicate, "predicate"); πE != nil {
						continue
					}
					πTemp004[1] = µpredicate
					if πE = πg.CheckLocal(πF, µedge_func, "edge_func"); πE != nil {
						continue
					}
					πTemp004[2] = µedge_func
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					πTemp004[3] = µdepth
					if πTemp006, πE = πg.ResolveGlobal(πF, ßfind_chain); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
						continue
					}
					µchain = πTemp005
					// line 51: return chain
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µchain, "chain"); πE != nil {
						continue
					}
					πR = µchain
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßparents.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 38: """Find the chain of referents for obj. Chain will end with obj.
			πF.SetLineno(38)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Find the chain of referents for obj. Chain will end with obj.\n\n    objtype: an object type or tuple of types to search for\n    depth: search depth (e.g. depth=2 is 'grandparents')\n    ignore: an object or tuple of objects to ignore in the search\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßparents); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
				continue
			}
			// line 54: def children(obj, objtype, depth=1, ignore=()): #XXX: objtype=object ?
			πF.SetLineno(54)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "objtype", Def: nil}
			πTemp004[2] = πg.Param{Name: "depth", Def: πg.NewInt(1).ToObject()}
			πTemp006 = πg.NewTuple0().ToObject()
			πTemp004[3] = πg.Param{Name: "ignore", Def: πTemp006}
			πTemp005 = πg.NewFunction(πg.NewCode("children", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/pointers.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µobjtype *πg.Object = πArgs[1]; _ = µobjtype
				var µdepth *πg.Object = πArgs[2]; _ = µdepth
				var µignore *πg.Object = πArgs[3]; _ = µignore
				var µedge_func *πg.Object = πg.UnboundLocal; _ = µedge_func
				var µpredicate *πg.Object = πg.UnboundLocal; _ = µpredicate
				var µchain *πg.Object = πg.UnboundLocal; _ = µchain
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 55: """Find the chain of referrers for obj. Chain will start with obj.
					πF.SetLineno(55)
					// line 67: edge_func = gc.get_referrers # looking for back_refs, not refs
					πF.SetLineno(67)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_referrers, nil); πE != nil {
						continue
					}
					µedge_func = πTemp002
					// line 68: predicate = lambda x: isinstance(x, objtype) # looking for child type
					πF.SetLineno(68)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "x", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/pointers.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
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
							// line 68: predicate = lambda x: isinstance(x, objtype) # looking for child type
							πF.SetLineno(68)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µobjtype, "objtype"); πE != nil {
								continue
							}
							πTemp001[1] = µobjtype
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
					µpredicate = πTemp001
					// line 70: ignore = (ignore,) if not hasattr(ignore, '__len__') else ignore
					πF.SetLineno(70)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp004[0] = µignore
					πTemp004[1] = ß__len__.ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(µignore).ToObject()
					πTemp001 = πTemp002
					goto Label2
				Label1:
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp001 = µignore
				Label2:
					µignore = πTemp001
					// line 71: ignore = (id(obj) for obj in ignore)
					πF.SetLineno(71)
					πTemp003 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/pointers.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
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
								if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µignore); πE != nil {
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
									µobj = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 71: ignore = (id(obj) for obj in ignore)
								πF.SetLineno(71)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp005[0] = µobj
								if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
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
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µignore = πTemp002
					// line 72: chain = find_chain(obj, predicate, edge_func, depth, ignore)
					πF.SetLineno(72)
					πTemp004 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µpredicate, "predicate"); πE != nil {
						continue
					}
					πTemp004[1] = µpredicate
					if πE = πg.CheckLocal(πF, µedge_func, "edge_func"); πE != nil {
						continue
					}
					πTemp004[2] = µedge_func
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					πTemp004[3] = µdepth
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp004[4] = µignore
					if πTemp002, πE = πg.ResolveGlobal(πF, ßfind_chain); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µchain = πTemp005
					// line 74: return chain
					πF.SetLineno(74)
					if πE = πg.CheckLocal(πF, µchain, "chain"); πE != nil {
						continue
					}
					πR = µchain
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßchildren.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 55: """Find the chain of referrers for obj. Chain will start with obj.
			πF.SetLineno(55)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Find the chain of referrers for obj. Chain will start with obj.\n\n    objtype: an object type or tuple of types to search for\n    depth: search depth (e.g. depth=2 is 'grandchildren')\n    ignore: an object or tuple of objects to ignore in the search\n\n    NOTE: a common thing to ignore is all globals, 'ignore=(globals(),)'\n\n    NOTE: repeated calls may yield different results, as python stores\n    the last value in the special variable '_'; thus, it is often good\n    to execute something to replace '_' (e.g. >>> 1+1).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßchildren); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
			// line 83: def find_chain(obj, predicate, edge_func, max_depth=20, extra_ignore=()):
			πF.SetLineno(83)
			πTemp004 = make([]πg.Param, 5)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "predicate", Def: nil}
			πTemp004[2] = πg.Param{Name: "edge_func", Def: nil}
			πTemp004[3] = πg.Param{Name: "max_depth", Def: πg.NewInt(20).ToObject()}
			πTemp007 = πg.NewTuple0().ToObject()
			πTemp004[4] = πg.Param{Name: "extra_ignore", Def: πTemp007}
			πTemp006 = πg.NewFunction(πg.NewCode("find_chain", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/pointers.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µpredicate *πg.Object = πArgs[1]; _ = µpredicate
				var µedge_func *πg.Object = πArgs[2]; _ = µedge_func
				var µmax_depth *πg.Object = πArgs[3]; _ = µmax_depth
				var µextra_ignore *πg.Object = πArgs[4]; _ = µextra_ignore
				var µqueue *πg.Object = πg.UnboundLocal; _ = µqueue
				var µdepth *πg.Object = πg.UnboundLocal; _ = µdepth
				var µparent *πg.Object = πg.UnboundLocal; _ = µparent
				var µignore *πg.Object = πg.UnboundLocal; _ = µignore
				var µtarget *πg.Object = πg.UnboundLocal; _ = µtarget
				var µchain *πg.Object = πg.UnboundLocal; _ = µchain
				var µtdepth *πg.Object = πg.UnboundLocal; _ = µtdepth
				var µreferrers *πg.Object = πg.UnboundLocal; _ = µreferrers
				var µsource *πg.Object = πg.UnboundLocal; _ = µsource
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Dict
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
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
					case 1: goto Label1
					case 2: goto Label2
					case 6: goto Label6
					case 7: goto Label7
					case 11: goto Label11
					case 12: goto Label12
					default: panic("unexpected function state")
					}
					// line 84: queue = [obj]
					πF.SetLineno(84)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µqueue = πTemp002
					// line 85: depth = {id(obj): 0}
					πF.SetLineno(85)
					πTemp003 = πg.NewDict()
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πTemp003.SetItem(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003.ToObject()
					µdepth = πTemp002
					// line 86: parent = {id(obj): None}
					πF.SetLineno(86)
					πTemp003 = πg.NewDict()
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, πTemp004, πTemp002); πE != nil {
						continue
					}
					πTemp002 = πTemp003.ToObject()
					µparent = πTemp002
					// line 87: ignore = set(extra_ignore)
					πF.SetLineno(87)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µextra_ignore, "extra_ignore"); πE != nil {
						continue
					}
					πTemp001[0] = µextra_ignore
					if πTemp002, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µignore = πTemp004
					// line 88: ignore.add(id(extra_ignore))
					πF.SetLineno(88)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µextra_ignore, "extra_ignore"); πE != nil {
						continue
					}
					πTemp005[0] = µextra_ignore
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µignore, ßadd, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 89: ignore.add(id(queue))
					πF.SetLineno(89)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
						continue
					}
					πTemp005[0] = µqueue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µignore, ßadd, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 90: ignore.add(id(depth))
					πF.SetLineno(90)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					πTemp005[0] = µdepth
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µignore, ßadd, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 91: ignore.add(id(parent))
					πF.SetLineno(91)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
						continue
					}
					πTemp005[0] = µparent
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µignore, ßadd, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 92: ignore.add(id(ignore))
					πF.SetLineno(92)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp005[0] = µignore
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µignore, ßadd, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 93: ignore.add(id(sys._getframe()))  # this function
					πF.SetLineno(93)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ß_getframe, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µignore, ßadd, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 94: ignore.add(id(sys._getframe(1))) # find_chain/find_backref_chain, likely
					πF.SetLineno(94)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = πg.NewInt(1).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ß_getframe, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µignore, ßadd, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 95: gc.collect()
					πF.SetLineno(95)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgc); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßcollect, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 96: while queue:
					πF.SetLineno(96)
					πF.PushCheckpoint(2)
					πTemp007 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µqueue); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 97: target = queue.pop(0)
					πF.SetLineno(97)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µqueue, ßpop, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtarget = πTemp004
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					πTemp001[0] = µtarget
					if πE = πg.CheckLocal(πF, µpredicate, "predicate"); πE != nil {
						continue
					}
					if πTemp002, πE = µpredicate.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label4
					}
					goto Label5
					// line 98: if predicate(target):
					πF.SetLineno(98)
				Label4:
					// line 99: chain = [target]
					πF.SetLineno(99)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					πTemp001[0] = µtarget
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µchain = πTemp002
					// line 100: while parent[id(target)] is not None:
					πF.SetLineno(100)
					πF.PushCheckpoint(7)
					πTemp008 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label8
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					πTemp001[0] = µtarget
					if πTemp010, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp004 = πTemp011
					if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µparent, πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp010 != πTemp004).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(6)            
					// line 101: target = parent[id(target)]
					πF.SetLineno(101)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					πTemp001[0] = µtarget
					if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp010
					if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µparent, πTemp002); πE != nil {
						continue
					}
					µtarget = πTemp004
					// line 102: chain.append(target)
					πF.SetLineno(102)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					πTemp001[0] = µtarget
					if πE = πg.CheckLocal(πF, µchain, "chain"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µchain, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 103: return chain
					πF.SetLineno(103)
					if πE = πg.CheckLocal(πF, µchain, "chain"); πE != nil {
						continue
					}
					πR = µchain
					continue
					goto Label5
				Label5:
					// line 104: tdepth = depth[id(target)]
					πF.SetLineno(104)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					πTemp001[0] = µtarget
					if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp010
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µdepth, πTemp002); πE != nil {
						continue
					}
					µtdepth = πTemp004
					if πE = πg.CheckLocal(πF, µtdepth, "tdepth"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmax_depth, "max_depth"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µtdepth, µmax_depth); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label9
					}
					goto Label10
					// line 105: if tdepth < max_depth:
					πF.SetLineno(105)
				Label9:
					// line 106: referrers = edge_func(target)
					πF.SetLineno(106)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					πTemp001[0] = µtarget
					if πE = πg.CheckLocal(πF, µedge_func, "edge_func"); πE != nil {
						continue
					}
					if πTemp002, πE = µedge_func.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µreferrers = πTemp002
					// line 107: ignore.add(id(referrers))
					πF.SetLineno(107)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µreferrers, "referrers"); πE != nil {
						continue
					}
					πTemp005[0] = µreferrers
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µignore, ßadd, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µreferrers, "referrers"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µreferrers); πE != nil {
						continue
					}
					πF.PushCheckpoint(12)
					πTemp008 = false
				Label11:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label13
					}
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µsource = πTemp004
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(11)            
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πTemp010, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, µignore, πTemp011); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label14
					}
					goto Label15
					// line 109: if id(source) in ignore:
					πF.SetLineno(109)
				Label14:
					// line 110: continue
					πF.SetLineno(110)
					continue
					goto Label15
				Label15:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πTemp010, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, µdepth, πTemp011); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label16
					}
					goto Label17
					// line 111: if id(source) not in depth:
					πF.SetLineno(111)
				Label16:
					// line 112: depth[id(source)] = tdepth + 1
					πF.SetLineno(112)
					if πE = πg.CheckLocal(πF, µtdepth, "tdepth"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µtdepth, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πTemp012, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp012.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp011 = πTemp013
					if πE = πg.SetItem(πF, µdepth, πTemp011, πTemp010); πE != nil {
						continue
					}
					// line 113: parent[id(source)] = target
					πF.SetLineno(113)
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µtarget); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πTemp011, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp010 = πTemp012
					if πE = πg.SetItem(πF, µparent, πTemp010, πTemp004); πE != nil {
						continue
					}
					// line 114: queue.append(source)
					πF.SetLineno(114)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µqueue, ßappend, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label17
				Label17:
					continue
				Label12:
					if πE != nil || πR != nil {
						continue
					}
				Label13:
					goto Label10
				Label10:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 115: return [obj] # not found
					πF.SetLineno(115)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					πTemp002 = πg.NewList(πTemp001...).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßfind_chain.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 119: refobject = at
			πF.SetLineno(119)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßat); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßrefobject.ToObject(), πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("pointers", Code)
}

