package temp
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/pkg/dill"
	// _ "github.com/pygolin/stdlib/pkg/StringIO"
	// _ "github.com/pygolin/stdlib/pkg/tempfile"
	// _ "github.com/pygolin/stdlib/pkg/_dill"
	// _ "github.com/pygolin/stdlib/pkg/sys"
	// _ "github.com/pygolin/stdlib/pkg/source"
	// _ "github.com/pygolin/stdlib/pkg/contextlib"
	// _ "github.com/pygolin/stdlib/pkg/io"
	// _ "github.com/pygolin/stdlib/pkg/codecs"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/temp.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßBytesIO := πg.InternStr("BytesIO")
		ßIOError := πg.InternStr("IOError")
		ßNamedTemporaryFile := πg.InternStr("NamedTemporaryFile")
		ßNone := πg.InternStr("None")
		ßPY3 := πg.InternStr("PY3")
		ßStringIO := πg.InternStr("StringIO")
		ß__all__ := πg.InternStr("__all__")
		ß__doc__ := πg.InternStr("__doc__")
		ßalias := πg.InternStr("alias")
		ßb := πg.InternStr("b")
		ßcapture := πg.InternStr("capture")
		ßcontextlib := πg.InternStr("contextlib")
		ßcontextmanager := πg.InternStr("contextmanager")
		ßdecode := πg.InternStr("decode")
		ßdump := πg.InternStr("dump")
		ßdumpIO := πg.InternStr("dumpIO")
		ßdumpIO_source := πg.InternStr("dumpIO_source")
		ßdump_source := πg.InternStr("dump_source")
		ßeval := πg.InternStr("eval")
		ßflush := πg.InternStr("flush")
		ßgetattr := πg.InternStr("getattr")
		ßgetname := πg.InternStr("getname")
		ßgetvalue := πg.InternStr("getvalue")
		ßimportable := πg.InternStr("importable")
		ßjoin := πg.InternStr("join")
		ßlatin_1_encode := πg.InternStr("latin_1_encode")
		ßload := πg.InternStr("load")
		ßloadIO := πg.InternStr("loadIO")
		ßloadIO_source := πg.InternStr("loadIO_source")
		ßload_source := πg.InternStr("load_source")
		ßmode := πg.InternStr("mode")
		ßname := πg.InternStr("name")
		ßopen := πg.InternStr("open")
		ßpop := πg.InternStr("pop")
		ßr := πg.InternStr("r")
		ßrb := πg.InternStr("rb")
		ßread := πg.InternStr("read")
		ßsetattr := πg.InternStr("setattr")
		ßsplit := πg.InternStr("split")
		ßsplitlines := πg.InternStr("splitlines")
		ßstdout := πg.InternStr("stdout")
		ßstr := πg.InternStr("str")
		ßstrip := πg.InternStr("strip")
		ßsuffix := πg.InternStr("suffix")
		ßwrite := πg.InternStr("write")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 8: """
			πF.SetLineno(8)
			// line 8: """
			πF.SetLineno(8)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nMethods for serialized objects (or source code) stored in temporary files\nand file-like objects.\n").ToObject()); πE != nil {
				continue
			}
			// line 15: __all__ = ['dump_source', 'dump', 'dumpIO_source', 'dumpIO',\
			πF.SetLineno(15)
			πTemp001 = make([]*πg.Object, 9)
			πTemp001[0] = ßdump_source.ToObject()
			πTemp001[1] = ßdump.ToObject()
			πTemp001[2] = ßdumpIO_source.ToObject()
			πTemp001[3] = ßdumpIO.ToObject()
			πTemp001[4] = ßload_source.ToObject()
			πTemp001[5] = ßload.ToObject()
			πTemp001[6] = ßloadIO_source.ToObject()
			πTemp001[7] = ßloadIO.ToObject()
			πTemp001[8] = ßcapture.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 19: import contextlib
			πF.SetLineno(19)
			if πTemp001, πE = πg.ImportModule(πF, "contextlib"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßcontextlib.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 20: from ._dill import PY3
			πF.SetLineno(20)
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
			// line 24: def capture(stream='stdout'):
			πF.SetLineno(24)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "stream", Def: ßstdout.ToObject()}
			πTemp002 = πg.NewFunction(πg.NewCode("capture", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/temp.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstream *πg.Object = πArgs[0]; _ = µstream
				var µsys *πg.Object = πg.UnboundLocal; _ = µsys
				var µStringIO *πg.Object = πg.UnboundLocal; _ = µStringIO
				var µorig *πg.Object = πg.UnboundLocal; _ = µorig
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 4: goto Label4
						case 5: goto Label5
						default: panic("unexpected function state")
						}
						// line 25: """builds a context that temporarily replaces the given stream name
						πF.SetLineno(25)
						// line 34: import sys
						πF.SetLineno(34)
						if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
							continue
						}
						πTemp001 = πTemp002[0]
						µsys = πTemp001
						if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label1
						}
						goto Label2
						// line 35: if PY3:
						πF.SetLineno(35)
					Label1:
						// line 36: from io import StringIO
						πF.SetLineno(36)
						if πTemp002, πE = πg.ImportModule(πF, "io"); πE != nil {
							continue
						}
						πTemp001 = πTemp002[0]
						if πTemp004, πE = πg.GetAttrImport(πF, πTemp001, ßStringIO); πE != nil {
							continue
						}
						µStringIO = πTemp004
						goto Label3
					Label2:
						// line 38: from StringIO import StringIO
						πF.SetLineno(38)
						if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
							continue
						}
						πTemp001 = πTemp002[0]
						if πTemp004, πE = πg.GetAttrImport(πF, πTemp001, ßStringIO); πE != nil {
							continue
						}
						µStringIO = πTemp004
						goto Label3
					Label3:
						// line 39: orig = getattr(sys, stream)
						πF.SetLineno(39)
						πTemp002 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
							continue
						}
						πTemp002[0] = µsys
						if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
							continue
						}
						πTemp002[1] = µstream
						if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						µorig = πTemp004
						// line 40: setattr(sys, stream, StringIO())
						πF.SetLineno(40)
						πTemp002 = πF.MakeArgs(3)
						if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
							continue
						}
						πTemp002[0] = µsys
						if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
							continue
						}
						πTemp002[1] = µstream
						if πE = πg.CheckLocal(πF, µStringIO, "StringIO"); πE != nil {
							continue
						}
						if πTemp001, πE = µStringIO.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp002[2] = πTemp001
						if πTemp001, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						// line 41: try:
						πF.SetLineno(41)
						πF.PushCheckpoint(4)
						// line 42: yield getattr(sys, stream)
						πF.SetLineno(42)
						πTemp002 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
							continue
						}
						πTemp002[0] = µsys
						if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
							continue
						}
						πTemp002[1] = µstream
						if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πF.PushCheckpoint(5)
						return πTemp004, nil
					Label5:
						πTemp001 = πSent
						πF.PopCheckpoint()
						goto Label4
					Label4:
						πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
						// line 44: setattr(sys, stream, orig)
						πF.SetLineno(44)
						πTemp002 = πF.MakeArgs(3)
						if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
							continue
						}
						πTemp002[0] = µsys
						if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
							continue
						}
						πTemp002[1] = µstream
						if πE = πg.CheckLocal(πF, µorig, "orig"); πE != nil {
							continue
						}
						πTemp002[2] = µorig
						if πTemp001, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						if πTemp005 != nil {
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
						}
						if πR != nil {
							continue
						}
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcapture.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 25: """builds a context that temporarily replaces the given stream name
			πF.SetLineno(25)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("builds a context that temporarily replaces the given stream name\n\n    >>> with capture('stdout') as out:\n    ...   print \"foo!\"\n    ... \n    >>> print out.getvalue()\n    foo!\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcapture); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 23: @contextlib.contextmanager
			πF.SetLineno(23)
			πTemp001 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßcapture); πE != nil {
				continue
			}
			πTemp001[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßcontextlib); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßcontextmanager, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßcapture.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 47: def b(x): # deal with b'foo' versus 'foo'
			πF.SetLineno(47)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("b", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/temp.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µcodecs *πg.Object = πg.UnboundLocal; _ = µcodecs
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 48: import codecs
					πF.SetLineno(48)
					if πTemp002, πE = πg.ImportModule(πF, "codecs"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µcodecs = πTemp001
					// line 49: return codecs.latin_1_encode(x)[0]
					πF.SetLineno(49)
					πTemp001 = πg.NewInt(0).ToObject()
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πE = πg.CheckLocal(πF, µcodecs, "codecs"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µcodecs, ßlatin_1_encode, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßb.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 51: def load_source(file, **kwds):
			πF.SetLineno(51)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "file", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("load_source", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/temp.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfile *πg.Object = πArgs[0]; _ = µfile
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var µalias *πg.Object = πg.UnboundLocal; _ = µalias
				var µmode *πg.Object = πg.UnboundLocal; _ = µmode
				var µfname *πg.Object = πg.UnboundLocal; _ = µfname
				var µsource *πg.Object = πg.UnboundLocal; _ = µsource
				var µtag *πg.Object = πg.UnboundLocal; _ = µtag
				var µstub *πg.Object = πg.UnboundLocal; _ = µstub
				var µlocal *πg.Object = πg.UnboundLocal; _ = µlocal
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 πg.KWArgs
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Dict
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 52: """load an object that was stored with dill.temp.dump_source
					πF.SetLineno(52)
					// line 64: alias = kwds.pop('alias', None)
					πF.SetLineno(64)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßalias.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µalias = πTemp003
					// line 65: mode = kwds.pop('mode', 'r')
					πF.SetLineno(65)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßmode.ToObject()
					πTemp001[1] = ßr.ToObject()
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmode = πTemp003
					// line 66: fname = getattr(file, 'name', file) # fname=file.name or fname=file (if str)
					πF.SetLineno(66)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp001[0] = µfile
					πTemp001[1] = ßname.ToObject()
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp001[2] = µfile
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfname = πTemp003
					// line 67: source = open(fname, mode=mode, **kwds).read()
					πF.SetLineno(67)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfname, "fname"); πE != nil {
						continue
					}
					πTemp001[0] = µfname
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"mode", µmode},
					}
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, nil, πTemp004, µkwds); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsource = πTemp003
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µalias); πE != nil {
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
					// line 68: if not alias:
					πF.SetLineno(68)
				Label1:
					// line 69: tag = source.strip().splitlines()[-1].split()
					πF.SetLineno(69)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µsource, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßsplitlines, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtag = πTemp003
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtag, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, πTemp006, πg.NewStr("#NAME:").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					goto Label4
					// line 70: if tag[0] != '#NAME:':
					πF.SetLineno(70)
				Label3:
					// line 71: stub = source.splitlines()[0]
					πF.SetLineno(71)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µsource, ßsplitlines, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
						continue
					}
					µstub = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstub, "stub"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("unknown name for code: %s").ToObject(), µstub); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 72: raise IOError("unknown name for code: %s" % stub)
					πF.SetLineno(72)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label4
				Label4:
					// line 73: alias = tag[-1]
					πF.SetLineno(73)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µtag, πTemp002); πE != nil {
						continue
					}
					µalias = πTemp003
					goto Label2
				Label2:
					// line 74: local = {}
					πF.SetLineno(74)
					πTemp008 = πg.NewDict()
					πTemp002 = πTemp008.ToObject()
					µlocal = πTemp002
					// line 75: exec(source, local)
					πF.SetLineno(75)
					πE = πF.RaiseType(πg.NotImplementedErrorType, "exec is not available on Grumpy. Maybe never be.")
					continue
					// line 76: _ = eval("%s" % alias, local)
					πF.SetLineno(76)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s").ToObject(), µalias); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µlocal, "local"); πE != nil {
						continue
					}
					πTemp001[1] = µlocal
					if πTemp002, πE = πg.ResolveGlobal(πF, ßeval); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µ_ = πTemp003
					// line 77: return _
					πF.SetLineno(77)
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					πR = µ_
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßload_source.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 52: """load an object that was stored with dill.temp.dump_source
			πF.SetLineno(52)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("load an object that was stored with dill.temp.dump_source\n\n    file: filehandle\n    alias: string name of stored object\n    mode: mode to open the file, one of: {'r', 'rb'}\n\n    >>> f = lambda x: x**2\n    >>> pyfile = dill.temp.dump_source(f, alias='_f')\n    >>> _f = dill.temp.load_source(pyfile)\n    >>> _f(4)\n    16\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßload_source); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
			// line 79: def dump_source(object, **kwds):
			πF.SetLineno(79)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("dump_source", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/temp.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var µimportable *πg.Object = πg.UnboundLocal; _ = µimportable
				var µgetname *πg.Object = πg.UnboundLocal; _ = µgetname
				var µtempfile *πg.Object = πg.UnboundLocal; _ = µtempfile
				var µalias *πg.Object = πg.UnboundLocal; _ = µalias
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µfile *πg.Object = πg.UnboundLocal; _ = µfile
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
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 80: """write object source to a NamedTemporaryFile (instead of dill.dump)
					πF.SetLineno(80)
					// line 111: from .source import importable, getname
					πF.SetLineno(111)
					if πTemp002, πE = πg.ImportModule(πF, ".source"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßimportable); πE != nil {
						continue
					}
					µimportable = πTemp003
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßgetname); πE != nil {
						continue
					}
					µgetname = πTemp003
					// line 112: import tempfile
					πF.SetLineno(112)
					if πTemp002, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µtempfile = πTemp001
					// line 113: kwds.pop('suffix', '') # this is *always* '.py'
					πF.SetLineno(113)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ßsuffix.ToObject()
					πTemp002[1] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 114: alias = kwds.pop('alias', '') #XXX: include an alias so a name is known
					πF.SetLineno(114)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ßalias.ToObject()
					πTemp002[1] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µalias = πTemp003
					// line 115: name = str(alias) or getname(object)
					πF.SetLineno(115)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp002[0] = µalias
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πTemp005
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp002[0] = µobject
					if πE = πg.CheckLocal(πF, µgetname, "getname"); πE != nil {
						continue
					}
					if πTemp003, πE = µgetname.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πTemp003
				Label1:
					µname = πTemp001
					// line 116: name = "\n#NAME: %s\n" % name
					πF.SetLineno(116)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("\n#NAME: %s\n").ToObject(), µname); πE != nil {
						continue
					}
					µname = πTemp001
					// line 118: file = tempfile.NamedTemporaryFile(suffix='.py', **kwds)
					πF.SetLineno(118)
					πTemp006 = πg.KWArgs{
						{"suffix", πg.NewStr(".py").ToObject()},
					}
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtempfile, "tempfile"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtempfile, ßNamedTemporaryFile, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp001, nil, nil, πTemp006, µkwds); πE != nil {
						continue
					}
					µfile = πTemp003
					// line 119: file.write(b(''.join([importable(object, alias=alias),name])))
					πF.SetLineno(119)
					πTemp002 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					πTemp009 = make([]*πg.Object, 2)
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp010[0] = µobject
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"alias", µalias},
					}
					if πE = πg.CheckLocal(πF, µimportable, "importable"); πE != nil {
						continue
					}
					if πTemp001, πE = µimportable.Call(πF, πTemp010, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp009[0] = πTemp001
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp009[1] = µname
					πTemp001 = πg.NewList(πTemp009...).ToObject()
					πTemp008[0] = πTemp001
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßb); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfile, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 120: file.flush()
					πF.SetLineno(120)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfile, ßflush, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 121: return file
					πF.SetLineno(121)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πR = µfile
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdump_source.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 80: """write object source to a NamedTemporaryFile (instead of dill.dump)
			πF.SetLineno(80)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("write object source to a NamedTemporaryFile (instead of dill.dump)\nLoads with \"import\" or \"dill.temp.load_source\".  Returns the filehandle.\n\n    >>> f = lambda x: x**2\n    >>> pyfile = dill.temp.dump_source(f, alias='_f')\n    >>> _f = dill.temp.load_source(pyfile)\n    >>> _f(4)\n    16\n\n    >>> f = lambda x: x**2\n    >>> pyfile = dill.temp.dump_source(f, dir='.')\n    >>> modulename = os.path.basename(pyfile.name).split('.py')[0]\n    >>> exec('from %s import f as _f' % modulename)\n    >>> _f(4)\n    16\n\nOptional kwds:\n    If 'alias' is specified, the object will be renamed to the given string.\n\n    If 'prefix' is specified, the file name will begin with that prefix,\n    otherwise a default prefix is used.\n    \n    If 'dir' is specified, the file will be created in that directory,\n    otherwise a default directory is used.\n    \n    If 'text' is specified and true, the file is opened in text\n    mode.  Else (the default) the file is opened in binary mode.  On\n    some operating systems, this makes no difference.\n\nNOTE: Keep the return value for as long as you want your file to exist !\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßdump_source); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
				continue
			}
			// line 123: def load(file, **kwds):
			πF.SetLineno(123)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "file", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("load", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/temp.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfile *πg.Object = πArgs[0]; _ = µfile
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var µpickle *πg.Object = πg.UnboundLocal; _ = µpickle
				var µmode *πg.Object = πg.UnboundLocal; _ = µmode
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 πg.KWArgs
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 124: """load an object that was stored with dill.temp.dump
					πF.SetLineno(124)
					// line 133: import dill as pickle
					πF.SetLineno(133)
					if πTemp002, πE = πg.ImportModule(πF, "dill"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µpickle = πTemp001
					// line 134: mode = kwds.pop('mode', 'rb')
					πF.SetLineno(134)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ßmode.ToObject()
					πTemp002[1] = ßrb.ToObject()
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmode = πTemp003
					// line 135: name = getattr(file, 'name', file) # name=file.name or name=file (if str)
					πF.SetLineno(135)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp002[0] = µfile
					πTemp002[1] = ßname.ToObject()
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp002[2] = µfile
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µname = πTemp003
					// line 136: return pickle.load(open(name, mode=mode, **kwds))
					πF.SetLineno(136)
					πTemp002 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004[0] = µname
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"mode", µmode},
					}
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp001, πTemp004, nil, πTemp005, µkwds); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µpickle, "pickle"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickle, ßload, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
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
			if πE = πF.Globals().SetItem(πF, ßload.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 124: """load an object that was stored with dill.temp.dump
			πF.SetLineno(124)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("load an object that was stored with dill.temp.dump\n\n    file: filehandle\n    mode: mode to open the file, one of: {'r', 'rb'}\n\n    >>> dumpfile = dill.temp.dump([1, 2, 3, 4, 5])\n    >>> dill.temp.load(dumpfile)\n    [1, 2, 3, 4, 5]\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßload); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
			// line 138: def dump(object, **kwds):
			πF.SetLineno(138)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("dump", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/temp.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var µpickle *πg.Object = πg.UnboundLocal; _ = µpickle
				var µtempfile *πg.Object = πg.UnboundLocal; _ = µtempfile
				var µfile *πg.Object = πg.UnboundLocal; _ = µfile
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
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
					// line 139: """dill.dump of object to a NamedTemporaryFile.
					πF.SetLineno(139)
					// line 162: import dill as pickle
					πF.SetLineno(162)
					if πTemp002, πE = πg.ImportModule(πF, "dill"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µpickle = πTemp001
					// line 163: import tempfile
					πF.SetLineno(163)
					if πTemp002, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µtempfile = πTemp001
					// line 164: file = tempfile.NamedTemporaryFile(**kwds)
					πF.SetLineno(164)
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtempfile, "tempfile"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtempfile, ßNamedTemporaryFile, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp001, nil, nil, nil, µkwds); πE != nil {
						continue
					}
					µfile = πTemp003
					// line 165: pickle.dump(object, file)
					πF.SetLineno(165)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp002[0] = µobject
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp002[1] = µfile
					if πE = πg.CheckLocal(πF, µpickle, "pickle"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickle, ßdump, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 166: file.flush()
					πF.SetLineno(166)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfile, ßflush, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 167: return file
					πF.SetLineno(167)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πR = µfile
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdump.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 139: """dill.dump of object to a NamedTemporaryFile.
			πF.SetLineno(139)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("dill.dump of object to a NamedTemporaryFile.\nLoads with \"dill.temp.load\".  Returns the filehandle.\n\n    >>> dumpfile = dill.temp.dump([1, 2, 3, 4, 5])\n    >>> dill.temp.load(dumpfile)\n    [1, 2, 3, 4, 5]\n\nOptional kwds:\n    If 'suffix' is specified, the file name will end with that suffix,\n    otherwise there will be no suffix.\n    \n    If 'prefix' is specified, the file name will begin with that prefix,\n    otherwise a default prefix is used.\n    \n    If 'dir' is specified, the file will be created in that directory,\n    otherwise a default directory is used.\n    \n    If 'text' is specified and true, the file is opened in text\n    mode.  Else (the default) the file is opened in binary mode.  On\n    some operating systems, this makes no difference.\n\nNOTE: Keep the return value for as long as you want your file to exist !\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßdump); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 169: def loadIO(buffer, **kwds):
			πF.SetLineno(169)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "buffer", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("loadIO", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/temp.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µbuffer *πg.Object = πArgs[0]; _ = µbuffer
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var µpickle *πg.Object = πg.UnboundLocal; _ = µpickle
				var µStringIO *πg.Object = πg.UnboundLocal; _ = µStringIO
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 bool
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
					// line 170: """load an object that was stored with dill.temp.dumpIO
					πF.SetLineno(170)
					// line 178: import dill as pickle
					πF.SetLineno(178)
					if πTemp002, πE = πg.ImportModule(πF, "dill"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µpickle = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 179: if PY3:
					πF.SetLineno(179)
				Label1:
					// line 180: from io import BytesIO as StringIO
					πF.SetLineno(180)
					if πTemp002, πE = πg.ImportModule(πF, "io"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp004, πE = πg.GetAttrImport(πF, πTemp001, ßBytesIO); πE != nil {
						continue
					}
					µStringIO = πTemp004
					goto Label3
				Label2:
					// line 182: from StringIO import StringIO
					πF.SetLineno(182)
					if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp004, πE = πg.GetAttrImport(πF, πTemp001, ßStringIO); πE != nil {
						continue
					}
					µStringIO = πTemp004
					goto Label3
				Label3:
					// line 183: value = getattr(buffer, 'getvalue', buffer) # value or buffer.getvalue
					πF.SetLineno(183)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
						continue
					}
					πTemp002[0] = µbuffer
					πTemp002[1] = ßgetvalue.ToObject()
					if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
						continue
					}
					πTemp002[2] = µbuffer
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µvalue = πTemp004
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µvalue, µbuffer); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					goto Label5
					// line 184: if value != buffer: value = value() # buffer.getvalue()
					πF.SetLineno(184)
				Label4:
					// line 184: if value != buffer: value = value() # buffer.getvalue()
					πF.SetLineno(184)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp001, πE = µvalue.Call(πF, nil, nil); πE != nil {
						continue
					}
					µvalue = πTemp001
					goto Label5
				Label5:
					// line 185: return pickle.load(StringIO(value))
					πF.SetLineno(185)
					πTemp002 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp005[0] = µvalue
					if πE = πg.CheckLocal(πF, µStringIO, "StringIO"); πE != nil {
						continue
					}
					if πTemp001, πE = µStringIO.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µpickle, "pickle"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickle, ßload, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
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
			if πE = πF.Globals().SetItem(πF, ßloadIO.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 170: """load an object that was stored with dill.temp.dumpIO
			πF.SetLineno(170)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("load an object that was stored with dill.temp.dumpIO\n\n    buffer: buffer object\n\n    >>> dumpfile = dill.temp.dumpIO([1, 2, 3, 4, 5])\n    >>> dill.temp.loadIO(dumpfile)\n    [1, 2, 3, 4, 5]\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßloadIO); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 187: def dumpIO(object, **kwds):
			πF.SetLineno(187)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("dumpIO", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/temp.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var µpickle *πg.Object = πg.UnboundLocal; _ = µpickle
				var µStringIO *πg.Object = πg.UnboundLocal; _ = µStringIO
				var µfile *πg.Object = πg.UnboundLocal; _ = µfile
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 bool
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
					// line 188: """dill.dump of object to a buffer.
					πF.SetLineno(188)
					// line 195: import dill as pickle
					πF.SetLineno(195)
					if πTemp002, πE = πg.ImportModule(πF, "dill"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µpickle = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 196: if PY3:
					πF.SetLineno(196)
				Label1:
					// line 197: from io import BytesIO as StringIO
					πF.SetLineno(197)
					if πTemp002, πE = πg.ImportModule(πF, "io"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp004, πE = πg.GetAttrImport(πF, πTemp001, ßBytesIO); πE != nil {
						continue
					}
					µStringIO = πTemp004
					goto Label3
				Label2:
					// line 199: from StringIO import StringIO
					πF.SetLineno(199)
					if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp004, πE = πg.GetAttrImport(πF, πTemp001, ßStringIO); πE != nil {
						continue
					}
					µStringIO = πTemp004
					goto Label3
				Label3:
					// line 200: file = StringIO()
					πF.SetLineno(200)
					if πE = πg.CheckLocal(πF, µStringIO, "StringIO"); πE != nil {
						continue
					}
					if πTemp001, πE = µStringIO.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfile = πTemp001
					// line 201: pickle.dump(object, file)
					πF.SetLineno(201)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp002[0] = µobject
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp002[1] = µfile
					if πE = πg.CheckLocal(πF, µpickle, "pickle"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickle, ßdump, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 202: file.flush()
					πF.SetLineno(202)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfile, ßflush, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 203: return file
					πF.SetLineno(203)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πR = µfile
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdumpIO.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 188: """dill.dump of object to a buffer.
			πF.SetLineno(188)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("dill.dump of object to a buffer.\nLoads with \"dill.temp.loadIO\".  Returns the buffer object.\n\n    >>> dumpfile = dill.temp.dumpIO([1, 2, 3, 4, 5])\n    >>> dill.temp.loadIO(dumpfile)\n    [1, 2, 3, 4, 5]\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßdumpIO); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 205: def loadIO_source(buffer, **kwds):
			πF.SetLineno(205)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "buffer", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("loadIO_source", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/temp.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µbuffer *πg.Object = πArgs[0]; _ = µbuffer
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var µalias *πg.Object = πg.UnboundLocal; _ = µalias
				var µsource *πg.Object = πg.UnboundLocal; _ = µsource
				var µtag *πg.Object = πg.UnboundLocal; _ = µtag
				var µstub *πg.Object = πg.UnboundLocal; _ = µstub
				var µlocal *πg.Object = πg.UnboundLocal; _ = µlocal
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
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
				var πTemp007 *πg.Dict
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 206: """load an object that was stored with dill.temp.dumpIO_source
					πF.SetLineno(206)
					// line 217: alias = kwds.pop('alias', None)
					πF.SetLineno(217)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßalias.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µalias = πTemp003
					// line 218: source = getattr(buffer, 'getvalue', buffer) # source or buffer.getvalue
					πF.SetLineno(218)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
						continue
					}
					πTemp001[0] = µbuffer
					πTemp001[1] = ßgetvalue.ToObject()
					if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
						continue
					}
					πTemp001[2] = µbuffer
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsource = πTemp003
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, µsource, µbuffer); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 219: if source != buffer: source = source() # buffer.getvalue()
					πF.SetLineno(219)
				Label1:
					// line 219: if source != buffer: source = source() # buffer.getvalue()
					πF.SetLineno(219)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = µsource.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsource = πTemp002
					goto Label2
				Label2:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 220: if PY3: source = source.decode() # buffer to string
					πF.SetLineno(220)
				Label3:
					// line 220: if PY3: source = source.decode() # buffer to string
					πF.SetLineno(220)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsource, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsource = πTemp003
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µalias); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 221: if not alias:
					πF.SetLineno(221)
				Label5:
					// line 222: tag = source.strip().splitlines()[-1].split()
					πF.SetLineno(222)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µsource, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßsplitlines, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtag = πTemp003
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µtag, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, πTemp005, πg.NewStr("#NAME:").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 223: if tag[0] != '#NAME:':
					πF.SetLineno(223)
				Label7:
					// line 224: stub = source.splitlines()[0]
					πF.SetLineno(224)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µsource, ßsplitlines, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
						continue
					}
					µstub = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstub, "stub"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("unknown name for code: %s").ToObject(), µstub); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 225: raise IOError("unknown name for code: %s" % stub)
					πF.SetLineno(225)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label8
				Label8:
					// line 226: alias = tag[-1]
					πF.SetLineno(226)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µtag, πTemp002); πE != nil {
						continue
					}
					µalias = πTemp003
					goto Label6
				Label6:
					// line 227: local = {}
					πF.SetLineno(227)
					πTemp007 = πg.NewDict()
					πTemp002 = πTemp007.ToObject()
					µlocal = πTemp002
					// line 228: exec(source, local)
					πF.SetLineno(228)
					πE = πF.RaiseType(πg.NotImplementedErrorType, "exec is not available on Grumpy. Maybe never be.")
					continue
					// line 229: _ = eval("%s" % alias, local)
					πF.SetLineno(229)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s").ToObject(), µalias); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µlocal, "local"); πE != nil {
						continue
					}
					πTemp001[1] = µlocal
					if πTemp002, πE = πg.ResolveGlobal(πF, ßeval); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µ_ = πTemp003
					// line 230: return _
					πF.SetLineno(230)
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					πR = µ_
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßloadIO_source.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 206: """load an object that was stored with dill.temp.dumpIO_source
			πF.SetLineno(206)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("load an object that was stored with dill.temp.dumpIO_source\n\n    buffer: buffer object\n    alias: string name of stored object\n\n    >>> f = lambda x:x**2\n    >>> pyfile = dill.temp.dumpIO_source(f, alias='_f')\n    >>> _f = dill.temp.loadIO_source(pyfile)\n    >>> _f(4)\n    16\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßloadIO_source); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 232: def dumpIO_source(object, **kwds):
			πF.SetLineno(232)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("dumpIO_source", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/temp.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var µimportable *πg.Object = πg.UnboundLocal; _ = µimportable
				var µgetname *πg.Object = πg.UnboundLocal; _ = µgetname
				var µStringIO *πg.Object = πg.UnboundLocal; _ = µStringIO
				var µalias *πg.Object = πg.UnboundLocal; _ = µalias
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µfile *πg.Object = πg.UnboundLocal; _ = µfile
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
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
					// line 233: """write object source to a buffer (instead of dill.dump)
					πF.SetLineno(233)
					// line 245: from .source import importable, getname
					πF.SetLineno(245)
					if πTemp002, πE = πg.ImportModule(πF, ".source"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßimportable); πE != nil {
						continue
					}
					µimportable = πTemp003
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßgetname); πE != nil {
						continue
					}
					µgetname = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 246: if PY3:
					πF.SetLineno(246)
				Label1:
					// line 247: from io import BytesIO as StringIO
					πF.SetLineno(247)
					if πTemp002, πE = πg.ImportModule(πF, "io"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßBytesIO); πE != nil {
						continue
					}
					µStringIO = πTemp003
					goto Label3
				Label2:
					// line 249: from StringIO import StringIO
					πF.SetLineno(249)
					if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßStringIO); πE != nil {
						continue
					}
					µStringIO = πTemp003
					goto Label3
				Label3:
					// line 250: alias = kwds.pop('alias', '') #XXX: include an alias so a name is known
					πF.SetLineno(250)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ßalias.ToObject()
					πTemp002[1] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µalias = πTemp003
					// line 251: name = str(alias) or getname(object)
					πF.SetLineno(251)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp002[0] = µalias
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πTemp005
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp002[0] = µobject
					if πE = πg.CheckLocal(πF, µgetname, "getname"); πE != nil {
						continue
					}
					if πTemp003, πE = µgetname.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πTemp003
				Label4:
					µname = πTemp001
					// line 252: name = "\n#NAME: %s\n" % name
					πF.SetLineno(252)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("\n#NAME: %s\n").ToObject(), µname); πE != nil {
						continue
					}
					µname = πTemp001
					// line 254: file = StringIO()
					πF.SetLineno(254)
					if πE = πg.CheckLocal(πF, µStringIO, "StringIO"); πE != nil {
						continue
					}
					if πTemp001, πE = µStringIO.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfile = πTemp001
					// line 255: file.write(b(''.join([importable(object, alias=alias),name])))
					πF.SetLineno(255)
					πTemp002 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					πTemp008 = make([]*πg.Object, 2)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp009[0] = µobject
					if πE = πg.CheckLocal(πF, µalias, "alias"); πE != nil {
						continue
					}
					πTemp010 = πg.KWArgs{
						{"alias", µalias},
					}
					if πE = πg.CheckLocal(πF, µimportable, "importable"); πE != nil {
						continue
					}
					if πTemp001, πE = µimportable.Call(πF, πTemp009, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp008[0] = πTemp001
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp008[1] = µname
					πTemp001 = πg.NewList(πTemp008...).ToObject()
					πTemp007[0] = πTemp001
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp006[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßb); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfile, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 256: file.flush()
					πF.SetLineno(256)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfile, ßflush, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 257: return file
					πF.SetLineno(257)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πR = µfile
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdumpIO_source.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 233: """write object source to a buffer (instead of dill.dump)
			πF.SetLineno(233)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("write object source to a buffer (instead of dill.dump)\nLoads by with dill.temp.loadIO_source.  Returns the buffer object.\n\n    >>> f = lambda x:x**2\n    >>> pyfile = dill.temp.dumpIO_source(f, alias='_f')\n    >>> _f = dill.temp.loadIO_source(pyfile)\n    >>> _f(4)\n    16\n\nOptional kwds:\n    If 'alias' is specified, the object will be renamed to the given string.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßdumpIO_source); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
				continue
			}
			// line 260: del contextlib
			πF.SetLineno(260)
			if πE = πg.DelVar(πF, πF.Globals(), ßcontextlib); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("temp", Code)
}

