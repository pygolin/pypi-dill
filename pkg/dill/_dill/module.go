package _dill
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/cStringIO"
	// _ "github.com/pygolin/stdlib/functools"
	// _ "github.com/pygolin/stdlib/UserDict"
	// _ "github.com/pygolin/stdlib/multiprocessing/reduction"
	// _ "github.com/pygolin/stdlib/__builtin__"
	// _ "github.com/pygolin/stdlib/subprocess"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/_pyio"
	// _ "github.com/pygolin/stdlib/ctypes"
	// _ "github.com/pygolin/stdlib/io"
	// _ "github.com/pygolin/stdlib/operator"
	// _ "github.com/pygolin/stdlib/types"
	// _ "github.com/pygolin/stdlib/importlib"
	// _ "github.com/pygolin/stdlib/detect"
	// _ "github.com/pygolin/stdlib/logging"
	// _ "github.com/pygolin/stdlib/socket"
	// _ "github.com/pygolin/stdlib/thread"
	// _ "github.com/pygolin/stdlib/StringIO"
	// _ "github.com/pygolin/stdlib/tempfile"
	// _ "github.com/pygolin/stdlib/os"
	// _ "github.com/pygolin/stdlib/settings"
	// _ "github.com/pygolin/stdlib/weakref"
	// _ "github.com/pygolin/stdlib/threading"
	// _ "github.com/pygolin/stdlib/collections"
	// _ "github.com/pygolin/stdlib/pickle"
	// _ "github.com/pygolin/stdlib/multiprocessing"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAssertionError := πg.InternStr("AssertionError")
		ßAttrGetterType := πg.InternStr("AttrGetterType")
		ßAttributeError := πg.InternStr("AttributeError")
		ßB1 := πg.InternStr("B1")
		ßB3 := πg.InternStr("B3")
		ßBufferType := πg.InternStr("BufferType")
		ßBufferedRandomType := πg.InternStr("BufferedRandomType")
		ßBufferedReaderType := πg.InternStr("BufferedReaderType")
		ßBufferedWriterType := πg.InternStr("BufferedWriterType")
		ßBuiltinMethodType := πg.InternStr("BuiltinMethodType")
		ßBytesIO := πg.InternStr("BytesIO")
		ßCONTENTS_FMODE := πg.InternStr("CONTENTS_FMODE")
		ßC_EXTENSION := πg.InternStr("C_EXTENSION")
		ßCallableProxyType := πg.InternStr("CallableProxyType")
		ßCellType := πg.InternStr("CellType")
		ßClassMethodDescriptorType := πg.InternStr("ClassMethodDescriptorType")
		ßClassType := πg.InternStr("ClassType")
		ßCodeType := πg.InternStr("CodeType")
		ßDEFAULT_PROTOCOL := πg.InternStr("DEFAULT_PROTOCOL")
		ßDictProxyType := πg.InternStr("DictProxyType")
		ßEXTENSION_SUFFIXES := πg.InternStr("EXTENSION_SUFFIXES")
		ßEllipsis := πg.InternStr("Ellipsis")
		ßEllipsisType := πg.InternStr("EllipsisType")
		ßException := πg.InternStr("Exception")
		ßExitType := πg.InternStr("ExitType")
		ßFILE_FMODE := πg.InternStr("FILE_FMODE")
		ßFalse := πg.InternStr("False")
		ßFileNotFoundError := πg.InternStr("FileNotFoundError")
		ßFileType := πg.InternStr("FileType")
		ßFrameType := πg.InternStr("FrameType")
		ßFunctionType := πg.InternStr("FunctionType")
		ßGENERATOR_FAIL := πg.InternStr("GENERATOR_FAIL")
		ßGeneratorType := πg.InternStr("GeneratorType")
		ßGetSetDescriptorType := πg.InternStr("GetSetDescriptorType")
		ßHANDLE_FMODE := πg.InternStr("HANDLE_FMODE")
		ßHAS_CTYPES := πg.InternStr("HAS_CTYPES")
		ßHIGHEST_PROTOCOL := πg.InternStr("HIGHEST_PROTOCOL")
		ßINFO := πg.InternStr("INFO")
		ßIOError := πg.InternStr("IOError")
		ßIS_PYPY := πg.InternStr("IS_PYPY")
		ßImportError := πg.InternStr("ImportError")
		ßInputType := πg.InternStr("InputType")
		ßItemGetterType := πg.InternStr("ItemGetterType")
		ßKeyError := πg.InternStr("KeyError")
		ßLock := πg.InternStr("Lock")
		ßLockType := πg.InternStr("LockType")
		ßMappingProxyType := πg.InternStr("MappingProxyType")
		ßMemberDescriptorType := πg.InternStr("MemberDescriptorType")
		ßMetaCatchingDict := πg.InternStr("MetaCatchingDict")
		ßMethodDescriptorType := πg.InternStr("MethodDescriptorType")
		ßMethodType := πg.InternStr("MethodType")
		ßMethodWrapperType := πg.InternStr("MethodWrapperType")
		ßModuleType := πg.InternStr("ModuleType")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßNoneType := πg.InternStr("NoneType")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßNotImplementedType := πg.InternStr("NotImplementedType")
		ßNumpyArrayType := πg.InternStr("NumpyArrayType")
		ßNumpyUfuncType := πg.InternStr("NumpyUfuncType")
		ßOLD33 := πg.InternStr("OLD33")
		ßOLDER := πg.InternStr("OLDER")
		ßO_CREAT := πg.InternStr("O_CREAT")
		ßO_RDWR := πg.InternStr("O_RDWR")
		ßO_WRONLY := πg.InternStr("O_WRONLY")
		ßOutputType := πg.InternStr("OutputType")
		ßPOINTER := πg.InternStr("POINTER")
		ßPY3 := πg.InternStr("PY3")
		ßPY34 := πg.InternStr("PY34")
		ßPartialType := πg.InternStr("PartialType")
		ßPathFinder := πg.InternStr("PathFinder")
		ßPickler := πg.InternStr("Pickler")
		ßPicklingError := πg.InternStr("PicklingError")
		ßProxyType := πg.InternStr("ProxyType")
		ßPyBufferedRandomType := πg.InternStr("PyBufferedRandomType")
		ßPyBufferedReaderType := πg.InternStr("PyBufferedReaderType")
		ßPyBufferedWriterType := πg.InternStr("PyBufferedWriterType")
		ßPyTextWrapperType := πg.InternStr("PyTextWrapperType")
		ßR2 := πg.InternStr("R2")
		ßR3 := πg.InternStr("R3")
		ßRLock := πg.InternStr("RLock")
		ßRLockType := πg.InternStr("RLockType")
		ßReferenceError := πg.InternStr("ReferenceError")
		ßReferenceType := πg.InternStr("ReferenceType")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßSUCCESS := πg.InternStr("SUCCESS")
		ßSliceType := πg.InternStr("SliceType")
		ßSocketType := πg.InternStr("SocketType")
		ßStockPickler := πg.InternStr("StockPickler")
		ßStockUnpickler := πg.InternStr("StockUnpickler")
		ßStreamHandler := πg.InternStr("StreamHandler")
		ßStringIO := πg.InternStr("StringIO")
		ßStructure := πg.InternStr("Structure")
		ßSuperType := πg.InternStr("SuperType")
		ßT2 := πg.InternStr("T2")
		ßT3 := πg.InternStr("T3")
		ßTemporaryFile := πg.InternStr("TemporaryFile")
		ßTextWrapperType := πg.InternStr("TextWrapperType")
		ßTracebackType := πg.InternStr("TracebackType")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßTypeType := πg.InternStr("TypeType")
		ßUnpickler := πg.InternStr("Unpickler")
		ßUnpicklingError := πg.InternStr("UnpicklingError")
		ßUserDict := πg.InternStr("UserDict")
		ßValueError := πg.InternStr("ValueError")
		ßWARN := πg.InternStr("WARN")
		ßWrapperDescriptorType := πg.InternStr("WrapperDescriptorType")
		ßXRangeType := πg.InternStr("XRangeType")
		ß_Pickler := πg.InternStr("_Pickler")
		ß_RLock := πg.InternStr("_RLock")
		ß_RLock__owner := πg.InternStr("_RLock__owner")
		ß__IPYTHON__ := πg.InternStr("__IPYTHON__")
		ß__all__ := πg.InternStr("__all__")
		ß__bases__ := πg.InternStr("__bases__")
		ß__builtin__ := πg.InternStr("__builtin__")
		ß__builtins__ := πg.InternStr("__builtins__")
		ß__class__ := πg.InternStr("__class__")
		ß__closure__ := πg.InternStr("__closure__")
		ß__code__ := πg.InternStr("__code__")
		ß__defaults__ := πg.InternStr("__defaults__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__dill_imported := πg.InternStr("__dill_imported")
		ß__doc__ := πg.InternStr("__doc__")
		ß__file__ := πg.InternStr("__file__")
		ß__func__ := πg.InternStr("__func__")
		ß__get__ := πg.InternStr("__get__")
		ß__getattribute__ := πg.InternStr("__getattribute__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__globals__ := πg.InternStr("__globals__")
		ß__hook__ := πg.InternStr("__hook__")
		ß__import__ := πg.InternStr("__import__")
		ß__init__ := πg.InternStr("__init__")
		ß__kwdefaults__ := πg.InternStr("__kwdefaults__")
		ß__loader__ := πg.InternStr("__loader__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__missing__ := πg.InternStr("__missing__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__objclass__ := πg.InternStr("__objclass__")
		ß__prepare__ := πg.InternStr("__prepare__")
		ß__qualname__ := πg.InternStr("__qualname__")
		ß__reduce__ := πg.InternStr("__reduce__")
		ß__reduce_ex__ := πg.InternStr("__reduce_ex__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__self__ := πg.InternStr("__self__")
		ß__setstate__ := πg.InternStr("__setstate__")
		ß__slots__ := πg.InternStr("__slots__")
		ß__stderr__ := πg.InternStr("__stderr__")
		ß__stdin__ := πg.InternStr("__stdin__")
		ß__stdout__ := πg.InternStr("__stdout__")
		ß__thisclass__ := πg.InternStr("__thisclass__")
		ß__weakref__ := πg.InternStr("__weakref__")
		ß_acquire_restore := πg.InternStr("_acquire_restore")
		ß_asdict := πg.InternStr("_asdict")
		ß_attrgetter_helper := πg.InternStr("_attrgetter_helper")
		ß_byref := πg.InternStr("_byref")
		ß_create_array := πg.InternStr("_create_array")
		ß_create_cell := πg.InternStr("_create_cell")
		ß_create_code := πg.InternStr("_create_code")
		ß_create_filehandle := πg.InternStr("_create_filehandle")
		ß_create_ftype := πg.InternStr("_create_ftype")
		ß_create_function := πg.InternStr("_create_function")
		ß_create_lock := πg.InternStr("_create_lock")
		ß_create_namedtuple := πg.InternStr("_create_namedtuple")
		ß_create_rlock := πg.InternStr("_create_rlock")
		ß_create_stringi := πg.InternStr("_create_stringi")
		ß_create_stringo := πg.InternStr("_create_stringo")
		ß_create_type := πg.InternStr("_create_type")
		ß_create_typemap := πg.InternStr("_create_typemap")
		ß_create_weakproxy := πg.InternStr("_create_weakproxy")
		ß_create_weakref := πg.InternStr("_create_weakref")
		ß_dict_from_dictproxy := πg.InternStr("_dict_from_dictproxy")
		ß_diff_cache := πg.InternStr("_diff_cache")
		ß_eval_repr := πg.InternStr("_eval_repr")
		ß_extend := πg.InternStr("_extend")
		ß_fields := πg.InternStr("_fields")
		ß_fields_ := πg.InternStr("_fields_")
		ß_fmode := πg.InternStr("_fmode")
		ß_get_attr := πg.InternStr("_get_attr")
		ß_getattr := πg.InternStr("_getattr")
		ß_ignore := πg.InternStr("_ignore")
		ß_import_module := πg.InternStr("_import_module")
		ß_is_owned := πg.InternStr("_is_owned")
		ß_itemgetter_helper := πg.InternStr("_itemgetter_helper")
		ß_load_type := πg.InternStr("_load_type")
		ß_locate_function := πg.InternStr("_locate_function")
		ß_locate_object := πg.InternStr("_locate_object")
		ß_lookup_module := πg.InternStr("_lookup_module")
		ß_main := πg.InternStr("_main")
		ß_main_module := πg.InternStr("_main_module")
		ß_make := πg.InternStr("_make")
		ß_member := πg.InternStr("_member")
		ß_module_map := πg.InternStr("_module_map")
		ß_open := πg.InternStr("_open")
		ß_proxy_helper := πg.InternStr("_proxy_helper")
		ß_recurse := πg.InternStr("_recurse")
		ß_reduce_socket := πg.InternStr("_reduce_socket")
		ß_replace := πg.InternStr("_replace")
		ß_restore_modules := πg.InternStr("_restore_modules")
		ß_reverse_typemap := πg.InternStr("_reverse_typemap")
		ß_revert_extension := πg.InternStr("_revert_extension")
		ß_save_file := πg.InternStr("_save_file")
		ß_session := πg.InternStr("_session")
		ß_stash_modules := πg.InternStr("_stash_modules")
		ß_strictio := πg.InternStr("_strictio")
		ß_trace := πg.InternStr("_trace")
		ß_typemap := πg.InternStr("_typemap")
		ß_unmarshal := πg.InternStr("_unmarshal")
		ß_use_diff := πg.InternStr("_use_diff")
		ßacquire := πg.InternStr("acquire")
		ßaddHandler := πg.InternStr("addHandler")
		ßall := πg.InternStr("all")
		ßany := πg.InternStr("any")
		ßappend := πg.InternStr("append")
		ßargs := πg.InternStr("args")
		ßattrgetter := πg.InternStr("attrgetter")
		ßattrs := πg.InternStr("attrs")
		ßbase_exec_prefix := πg.InternStr("base_exec_prefix")
		ßbase_prefix := πg.InternStr("base_prefix")
		ßbool := πg.InternStr("bool")
		ßbuffer := πg.InternStr("buffer")
		ßbuiltins := πg.InternStr("builtins")
		ßbyref := πg.InternStr("byref")
		ßbytes := πg.InternStr("bytes")
		ßc_int := πg.InternStr("c_int")
		ßc_long := πg.InternStr("c_long")
		ßc_voidp := πg.InternStr("c_voidp")
		ßcall := πg.InternStr("call")
		ßcast := πg.InternStr("cast")
		ßcell_contents := πg.InternStr("cell_contents")
		ßcheck := πg.InternStr("check")
		ßclassmethod := πg.InternStr("classmethod")
		ßclear := πg.InternStr("clear")
		ßclear_memo := πg.InternStr("clear_memo")
		ßclose := πg.InternStr("close")
		ßclosed := πg.InternStr("closed")
		ßco_argcount := πg.InternStr("co_argcount")
		ßco_cellvars := πg.InternStr("co_cellvars")
		ßco_code := πg.InternStr("co_code")
		ßco_consts := πg.InternStr("co_consts")
		ßco_filename := πg.InternStr("co_filename")
		ßco_firstlineno := πg.InternStr("co_firstlineno")
		ßco_flags := πg.InternStr("co_flags")
		ßco_freevars := πg.InternStr("co_freevars")
		ßco_kwonlyargcount := πg.InternStr("co_kwonlyargcount")
		ßco_lnotab := πg.InternStr("co_lnotab")
		ßco_name := πg.InternStr("co_name")
		ßco_names := πg.InternStr("co_names")
		ßco_nlocals := πg.InternStr("co_nlocals")
		ßco_posonlyargcount := πg.InternStr("co_posonlyargcount")
		ßco_stacksize := πg.InternStr("co_stacksize")
		ßco_varnames := πg.InternStr("co_varnames")
		ßcontents := πg.InternStr("contents")
		ßcopy := πg.InternStr("copy")
		ßctypes := πg.InternStr("ctypes")
		ßdefaultdict := πg.InternStr("defaultdict")
		ßdescriptor := πg.InternStr("descriptor")
		ßdevnull := πg.InternStr("devnull")
		ßdict := πg.InternStr("dict")
		ßdiff := πg.InternStr("diff")
		ßdill := πg.InternStr("dill")
		ßdispatch := πg.InternStr("dispatch")
		ßdtype := πg.InternStr("dtype")
		ßdump := πg.InternStr("dump")
		ßdump_session := πg.InternStr("dump_session")
		ßdumps := πg.InternStr("dumps")
		ßencode := πg.InternStr("encode")
		ßendswith := πg.InternStr("endswith")
		ßeval := πg.InternStr("eval")
		ßexc_info := πg.InternStr("exc_info")
		ßexec_prefix := πg.InternStr("exec_prefix")
		ßexecutable := πg.InternStr("executable")
		ßexists := πg.InternStr("exists")
		ßexit := πg.InternStr("exit")
		ßfdel := πg.InternStr("fdel")
		ßfdopen := πg.InternStr("fdopen")
		ßfget := πg.InternStr("fget")
		ßfile_pointer := πg.InternStr("file_pointer")
		ßfind_class := πg.InternStr("find_class")
		ßfind_module := πg.InternStr("find_module")
		ßfind_spec := πg.InternStr("find_spec")
		ßflush := πg.InternStr("flush")
		ßfmode := πg.InternStr("fmode")
		ßfset := πg.InternStr("fset")
		ßfunc := πg.InternStr("func")
		ßfunc_closure := πg.InternStr("func_closure")
		ßfunc_code := πg.InternStr("func_code")
		ßfunc_defaults := πg.InternStr("func_defaults")
		ßfunc_globals := πg.InternStr("func_globals")
		ßfunc_name := πg.InternStr("func_name")
		ßgc := πg.InternStr("gc")
		ßget := πg.InternStr("get")
		ßgetLogger := πg.InternStr("getLogger")
		ßget_file_type := πg.InternStr("get_file_type")
		ßget_ipython := πg.InternStr("get_ipython")
		ßget_objects := πg.InternStr("get_objects")
		ßget_suffixes := πg.InternStr("get_suffixes")
		ßgetattr := πg.InternStr("getattr")
		ßgetsize := πg.InternStr("getsize")
		ßgetvalue := πg.InternStr("getvalue")
		ßglobalvars := πg.InternStr("globalvars")
		ßhasattr := πg.InternStr("hasattr")
		ßhex := πg.InternStr("hex")
		ßhexversion := πg.InternStr("hexversion")
		ßid := πg.InternStr("id")
		ßignore := πg.InternStr("ignore")
		ßim_class := πg.InternStr("im_class")
		ßim_func := πg.InternStr("im_func")
		ßim_self := πg.InternStr("im_self")
		ßimp := πg.InternStr("imp")
		ßimportlib := πg.InternStr("importlib")
		ßindex := πg.InternStr("index")
		ßinfo := πg.InternStr("info")
		ßint := πg.InternStr("int")
		ßint8 := πg.InternStr("int8")
		ßis_dill := πg.InternStr("is_dill")
		ßisinstance := πg.InternStr("isinstance")
		ßissubclass := πg.InternStr("issubclass")
		ßitemgetter := πg.InternStr("itemgetter")
		ßitems := πg.InternStr("items")
		ßiter := πg.InternStr("iter")
		ßiteritems := πg.InternStr("iteritems")
		ßitervalues := πg.InternStr("itervalues")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßkeywords := πg.InternStr("keywords")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßload := πg.InternStr("load")
		ßload_session := πg.InternStr("load_session")
		ßloads := πg.InternStr("loads")
		ßlocked := πg.InternStr("locked")
		ßlog := πg.InternStr("log")
		ßlogging := πg.InternStr("logging")
		ßmachinery := πg.InternStr("machinery")
		ßmarshal := πg.InternStr("marshal")
		ßmemoryview := πg.InternStr("memoryview")
		ßmode := πg.InternStr("mode")
		ßmodules := πg.InternStr("modules")
		ßmro := πg.InternStr("mro")
		ßname := πg.InternStr("name")
		ßnamedtuple := πg.InternStr("namedtuple")
		ßndarray := πg.InternStr("ndarray")
		ßndarraysubclassinstance := πg.InternStr("ndarraysubclassinstance")
		ßnormpath := πg.InternStr("normpath")
		ßnumpy := πg.InternStr("numpy")
		ßnumpyufunc := πg.InternStr("numpyufunc")
		ßob_refcnt := πg.InternStr("ob_refcnt")
		ßob_type := πg.InternStr("ob_type")
		ßobject := πg.InternStr("object")
		ßopen := πg.InternStr("open")
		ßos := πg.InternStr("os")
		ßpartial := πg.InternStr("partial")
		ßpath := πg.InternStr("path")
		ßpickle := πg.InternStr("pickle")
		ßpickle_dispatch_copy := πg.InternStr("pickle_dispatch_copy")
		ßpickles := πg.InternStr("pickles")
		ßpop := πg.InternStr("pop")
		ßprefix := πg.InternStr("prefix")
		ßproperty := πg.InternStr("property")
		ßprotocol := πg.InternStr("protocol")
		ßproxy := πg.InternStr("proxy")
		ßpy_object := πg.InternStr("py_object")
		ßpython := πg.InternStr("python")
		ßpythonapi := πg.InternStr("pythonapi")
		ßquit := πg.InternStr("quit")
		ßr := πg.InternStr("r")
		ßrange := πg.InternStr("range")
		ßraw := πg.InternStr("raw")
		ßrb := πg.InternStr("rb")
		ßread := πg.InternStr("read")
		ßreal_prefix := πg.InternStr("real_prefix")
		ßrecurse := πg.InternStr("recurse")
		ßreduce_socket := πg.InternStr("reduce_socket")
		ßref := πg.InternStr("ref")
		ßrefcount := πg.InternStr("refcount")
		ßregister := πg.InternStr("register")
		ßrepr := πg.InternStr("repr")
		ßrstrip := πg.InternStr("rstrip")
		ßsave_attrgetter := πg.InternStr("save_attrgetter")
		ßsave_builtin_method := πg.InternStr("save_builtin_method")
		ßsave_cell := πg.InternStr("save_cell")
		ßsave_classmethod := πg.InternStr("save_classmethod")
		ßsave_classobj := πg.InternStr("save_classobj")
		ßsave_code := πg.InternStr("save_code")
		ßsave_dict := πg.InternStr("save_dict")
		ßsave_dictproxy := πg.InternStr("save_dictproxy")
		ßsave_file := πg.InternStr("save_file")
		ßsave_function := πg.InternStr("save_function")
		ßsave_functor := πg.InternStr("save_functor")
		ßsave_global := πg.InternStr("save_global")
		ßsave_instancemethod := πg.InternStr("save_instancemethod")
		ßsave_instancemethod0 := πg.InternStr("save_instancemethod0")
		ßsave_itemgetter := πg.InternStr("save_itemgetter")
		ßsave_lock := πg.InternStr("save_lock")
		ßsave_module := πg.InternStr("save_module")
		ßsave_module_dict := πg.InternStr("save_module_dict")
		ßsave_property := πg.InternStr("save_property")
		ßsave_reduce := πg.InternStr("save_reduce")
		ßsave_rlock := πg.InternStr("save_rlock")
		ßsave_singleton := πg.InternStr("save_singleton")
		ßsave_slice := πg.InternStr("save_slice")
		ßsave_socket := πg.InternStr("save_socket")
		ßsave_stringi := πg.InternStr("save_stringi")
		ßsave_stringo := πg.InternStr("save_stringo")
		ßsave_super := πg.InternStr("save_super")
		ßsave_type := πg.InternStr("save_type")
		ßsave_weakproxy := πg.InternStr("save_weakproxy")
		ßsave_weakref := πg.InternStr("save_weakref")
		ßsave_wrapper_descriptor := πg.InternStr("save_wrapper_descriptor")
		ßseek := πg.InternStr("seek")
		ßsetLevel := πg.InternStr("setLevel")
		ßsettings := πg.InternStr("settings")
		ßshape := πg.InternStr("shape")
		ßsingletontypes := πg.InternStr("singletontypes")
		ßslice := πg.InternStr("slice")
		ßsocket := πg.InternStr("socket")
		ßsplit := πg.InternStr("split")
		ßstack := πg.InternStr("stack")
		ßstart := πg.InternStr("start")
		ßstartswith := πg.InternStr("startswith")
		ßstaticmethod := πg.InternStr("staticmethod")
		ßstep := πg.InternStr("step")
		ßstop := πg.InternStr("stop")
		ßstr := πg.InternStr("str")
		ßsuper := πg.InternStr("super")
		ßsys := πg.InternStr("sys")
		ßtell := πg.InternStr("tell")
		ßtmpfile := πg.InternStr("tmpfile")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßtype_obj := πg.InternStr("type_obj")
		ßufunc := πg.InternStr("ufunc")
		ßupdate := πg.InternStr("update")
		ßuse_diff := πg.InternStr("use_diff")
		ßvalues := πg.InternStr("values")
		ßverbose := πg.InternStr("verbose")
		ßw := πg.InternStr("w")
		ßwb := πg.InternStr("wb")
		ßwhats_changed := πg.InternStr("whats_changed")
		ßwrite := πg.InternStr("write")
		ßx := πg.InternStr("x")
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
		var πTemp007 bool
		_ = πTemp007
		var πTemp008 bool
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		var πTemp011 *πg.BaseException
		_ = πTemp011
		var πTemp012 *πg.Traceback
		_ = πTemp012
		var πTemp013 *πg.BaseException
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 []*πg.Object
		_ = πTemp017
		var πTemp018 *πg.Object
		_ = πTemp018
		var πTemp019 *πg.Object
		_ = πTemp019
		var πTemp020 *πg.Dict
		_ = πTemp020
		var πTemp021 πg.KWArgs
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
		var πTemp032 *πg.Object
		_ = πTemp032
		var πTemp033 *πg.Object
		_ = πTemp033
		var πTemp034 *πg.Object
		_ = πTemp034
		var πTemp035 *πg.Object
		_ = πTemp035
		var πTemp036 *πg.Object
		_ = πTemp036
		var πTemp037 *πg.Object
		_ = πTemp037
		var πTemp038 *πg.Object
		_ = πTemp038
		var πTemp039 *πg.Object
		_ = πTemp039
		var πTemp040 *πg.Object
		_ = πTemp040
		var πTemp041 *πg.Object
		_ = πTemp041
		var πTemp042 *πg.Object
		_ = πTemp042
		var πTemp043 *πg.Object
		_ = πTemp043
		var πTemp044 *πg.Object
		_ = πTemp044
		var πTemp045 *πg.Object
		_ = πTemp045
		var πTemp046 *πg.Object
		_ = πTemp046
		var πTemp047 *πg.Object
		_ = πTemp047
		var πTemp048 *πg.Object
		_ = πTemp048
		var πTemp049 *πg.Object
		_ = πTemp049
		var πTemp050 *πg.Object
		_ = πTemp050
		var πTemp051 *πg.Object
		_ = πTemp051
		var πTemp052 *πg.Object
		_ = πTemp052
		var πTemp053 *πg.Object
		_ = πTemp053
		var πTemp054 *πg.Object
		_ = πTemp054
		var πTemp055 *πg.Object
		_ = πTemp055
		var πTemp056 *πg.Object
		_ = πTemp056
		var πTemp057 *πg.Object
		_ = πTemp057
		var πTemp058 *πg.Object
		_ = πTemp058
		var πTemp059 *πg.Object
		_ = πTemp059
		var πTemp060 *πg.Object
		_ = πTemp060
		var πTemp061 *πg.Object
		_ = πTemp061
		var πTemp062 *πg.Object
		_ = πTemp062
		var πTemp063 *πg.Object
		_ = πTemp063
		var πTemp064 *πg.Object
		_ = πTemp064
		var πTemp065 *πg.Object
		_ = πTemp065
		var πTemp066 *πg.Object
		_ = πTemp066
		var πTemp067 *πg.Object
		_ = πTemp067
		var πTemp068 *πg.Object
		_ = πTemp068
		var πTemp069 *πg.Object
		_ = πTemp069
		var πTemp070 *πg.Object
		_ = πTemp070
		var πTemp071 *πg.Object
		_ = πTemp071
		var πTemp072 *πg.Object
		_ = πTemp072
		var πTemp073 *πg.Object
		_ = πTemp073
		var πTemp074 *πg.Object
		_ = πTemp074
		var πTemp075 *πg.Object
		_ = πTemp075
		var πTemp076 *πg.Object
		_ = πTemp076
		var πTemp077 *πg.Object
		_ = πTemp077
		var πTemp078 *πg.Object
		_ = πTemp078
		var πTemp079 *πg.Object
		_ = πTemp079
		var πTemp080 *πg.Object
		_ = πTemp080
		var πTemp081 *πg.Object
		_ = πTemp081
		var πTemp082 *πg.Object
		_ = πTemp082
		var πTemp083 *πg.Object
		_ = πTemp083
		var πTemp084 *πg.Object
		_ = πTemp084
		var πTemp085 *πg.Object
		_ = πTemp085
		var πTemp086 *πg.Object
		_ = πTemp086
		var πTemp087 *πg.Object
		_ = πTemp087
		var πTemp088 *πg.Object
		_ = πTemp088
		var πTemp089 *πg.Object
		_ = πTemp089
		var πTemp090 *πg.Object
		_ = πTemp090
		var πTemp091 *πg.Object
		_ = πTemp091
		var πTemp092 *πg.Object
		_ = πTemp092
		var πTemp093 *πg.Object
		_ = πTemp093
		var πTemp094 *πg.Object
		_ = πTemp094
		var πTemp095 *πg.Object
		_ = πTemp095
		var πTemp096 *πg.Object
		_ = πTemp096
		var πTemp097 *πg.Object
		_ = πTemp097
		var πTemp098 *πg.Object
		_ = πTemp098
		var πTemp099 *πg.Object
		_ = πTemp099
		var πTemp100 *πg.Object
		_ = πTemp100
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 67: goto Label67
			case 37: goto Label37
			case 70: goto Label70
			case 73: goto Label73
			case 14: goto Label14
			case 56: goto Label56
			case 26: goto Label26
			case 59: goto Label59
			case 29: goto Label29
			default: panic("unexpected function state")
			}
			// line 8: """
			πF.SetLineno(8)
			// line 8: """
			πF.SetLineno(8)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\ndill: a utility for serialization of python objects\n\nBased on code written by Oren Tirosh and Armin Ronacher.\nExtended to a (near) full set of the builtin types (in types module),\nand coded to the pickle interface, by <mmckerns@caltech.edu>.\nInitial port to python3 by Jonathan Dobson, continued by mmckerns.\nTest against \"all\" python types (Std. Lib. CH 1-15 @ 2.7) by mmckerns.\nTest against CH16+ Std. Lib. ... TBD.\n").ToObject()); πE != nil {
				continue
			}
			// line 18: __all__ = ['dump','dumps','load','loads','dump_session','load_session',
			πF.SetLineno(18)
			πTemp001 = make([]*πg.Object, 20)
			πTemp001[0] = ßdump.ToObject()
			πTemp001[1] = ßdumps.ToObject()
			πTemp001[2] = ßload.ToObject()
			πTemp001[3] = ßloads.ToObject()
			πTemp001[4] = ßdump_session.ToObject()
			πTemp001[5] = ßload_session.ToObject()
			πTemp001[6] = ßPickler.ToObject()
			πTemp001[7] = ßUnpickler.ToObject()
			πTemp001[8] = ßregister.ToObject()
			πTemp001[9] = ßcopy.ToObject()
			πTemp001[10] = ßpickle.ToObject()
			πTemp001[11] = ßpickles.ToObject()
			πTemp001[12] = ßcheck.ToObject()
			πTemp001[13] = ßHIGHEST_PROTOCOL.ToObject()
			πTemp001[14] = ßDEFAULT_PROTOCOL.ToObject()
			πTemp001[15] = ßPicklingError.ToObject()
			πTemp001[16] = ßUnpicklingError.ToObject()
			πTemp001[17] = ßHANDLE_FMODE.ToObject()
			πTemp001[18] = ßCONTENTS_FMODE.ToObject()
			πTemp001[19] = ßFILE_FMODE.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 23: import logging
			πF.SetLineno(23)
			if πTemp001, πE = πg.ImportModule(πF, "logging"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßlogging.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 24: log = logging.getLogger("dill")
			πF.SetLineno(24)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßdill.ToObject()
			if πTemp002, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetLogger, nil); πE != nil {
				continue
			}
			if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßlog.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 25: log.addHandler(logging.StreamHandler())
			πF.SetLineno(25)
			πTemp001 = πF.MakeArgs(1)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStreamHandler, nil); πE != nil {
				continue
			}
			if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßaddHandler, nil); πE != nil {
				continue
			}
			if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			// line 26: def _trace(boolean):
			πF.SetLineno(26)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "boolean", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("_trace", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µboolean *πg.Object = πArgs[0]; _ = µboolean
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 []*πg.Object
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
					// line 27: """print a trace through the stack when pickling; useful for debugging"""
					πF.SetLineno(27)
					if πE = πg.CheckLocal(πF, µboolean, "boolean"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µboolean); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 28: if boolean: log.setLevel(logging.INFO)
					πF.SetLineno(28)
				Label1:
					// line 28: if boolean: log.setLevel(logging.INFO)
					πF.SetLineno(28)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßINFO, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsetLevel, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label3
				Label2:
					// line 29: else: log.setLevel(logging.WARN)
					πF.SetLineno(29)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßWARN, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsetLevel, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label3
				Label3:
					// line 30: return
					πF.SetLineno(30)
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
			if πE = πF.Globals().SetItem(πF, ß_trace.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 27: """print a trace through the stack when pickling; useful for debugging"""
			πF.SetLineno(27)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("print a trace through the stack when pickling; useful for debugging").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_trace); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 32: stack = dict()  # record of 'recursion-sensitive' pickled objects
			πF.SetLineno(32)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstack.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 34: import os
			πF.SetLineno(34)
			if πTemp001, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 35: import sys
			πF.SetLineno(35)
			if πTemp001, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 36: diff = None
			πF.SetLineno(36)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdiff.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 37: _use_diff = False
			πF.SetLineno(37)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_use_diff.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 38: PY3 = (sys.hexversion >= 0x3000000)
			πF.SetLineno(38)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßhexversion, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GE(πF, πTemp006, πg.NewInt(50331648).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPY3.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 40: OLDER = (PY3 and sys.hexversion < 0x3040000) or (sys.hexversion < 0x2070ab1)
			πF.SetLineno(40)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			πTemp005 = πTemp006
			if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if !πTemp008 {
				goto Label2
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßhexversion, nil); πE != nil {
				continue
			}
			if πTemp006, πE = πg.LT(πF, πTemp010, πg.NewInt(50593792).ToObject()); πE != nil {
				continue
			}
			πTemp005 = πTemp006
		Label2:
			πTemp003 = πTemp005
			if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label1
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetAttr(πF, πTemp006, ßhexversion, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.LT(πF, πTemp009, πg.NewInt(34015921).ToObject()); πE != nil {
				continue
			}
			πTemp003 = πTemp005
		Label1:
			if πE = πF.Globals().SetItem(πF, ßOLDER.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 41: OLD33 = (sys.hexversion < 0x3030000)
			πF.SetLineno(41)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßhexversion, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.LT(πF, πTemp006, πg.NewInt(50528256).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOLD33.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 42: PY34 = (0x3040000 <= sys.hexversion < 0x3050000)
			πF.SetLineno(42)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßhexversion, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.LE(πF, πg.NewInt(50593792).ToObject(), πTemp006); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if !πTemp007 {
				goto Label3
			}
			if πTemp003, πE = πg.LT(πF, πTemp006, πg.NewInt(50659328).ToObject()); πE != nil {
				continue
			}
		Label3:
			if πE = πF.Globals().SetItem(πF, ßPY34.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label4
			}
			goto Label5
			// line 43: if PY3: #XXX: get types from .objtypes ?
			πF.SetLineno(43)
		Label4:
			// line 44: import builtins as __builtin__
			πF.SetLineno(44)
			if πTemp001, πE = πg.ImportModule(πF, "builtins"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ß__builtin__.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 45: from pickle import _Pickler as StockPickler, Unpickler as StockUnpickler
			πF.SetLineno(45)
			if πTemp001, πE = πg.ImportModule(πF, "pickle"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ß_Pickler); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStockPickler.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßUnpickler); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStockUnpickler.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 46: from _thread import LockType
			πF.SetLineno(46)
			if πTemp001, πE = πg.ImportModule(πF, "_thread"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßLockType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLockType.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßhexversion, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GE(πF, πTemp006, πg.NewInt(50462960).ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label7
			}
			goto Label8
			// line 47: if (sys.hexversion >= 0x30200f0):
			πF.SetLineno(47)
		Label7:
			// line 48: from _thread import RLock as RLockType
			πF.SetLineno(48)
			if πTemp001, πE = πg.ImportModule(πF, "_thread"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßRLock); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRLockType.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label9
		Label8:
			// line 50: from threading import _RLock as RLockType
			πF.SetLineno(50)
			if πTemp001, πE = πg.ImportModule(πF, "threading"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ß_RLock); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRLockType.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label9
		Label9:
			// line 52: from types import CodeType, FunctionType, MethodType, GeneratorType, \
			πF.SetLineno(52)
			if πTemp001, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßCodeType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCodeType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßFunctionType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFunctionType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßMethodType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMethodType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßGeneratorType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGeneratorType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßTracebackType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTracebackType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßFrameType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFrameType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßModuleType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßModuleType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßBuiltinMethodType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBuiltinMethodType.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 54: BufferType = memoryview #XXX: unregistered
			πF.SetLineno(54)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßmemoryview); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBufferType.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 55: ClassType = type # no 'old-style' classes
			πF.SetLineno(55)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßClassType.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 56: EllipsisType = type(Ellipsis)
			πF.SetLineno(56)
			πTemp001 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßEllipsis); πE != nil {
				continue
			}
			πTemp001[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßEllipsisType.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 58: NotImplementedType = type(NotImplemented)
			πF.SetLineno(58)
			πTemp001 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
				continue
			}
			πTemp001[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßNotImplementedType.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 59: SliceType = slice
			πF.SetLineno(59)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSliceType.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 60: TypeType = type # 'new-style' classes #XXX: unregistered
			πF.SetLineno(60)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTypeType.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 61: XRangeType = range
			πF.SetLineno(61)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßXRangeType.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßOLD33); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label10
			}
			goto Label11
			// line 62: if OLD33:
			πF.SetLineno(62)
		Label10:
			// line 63: DictProxyType = type(object.__dict__)
			πF.SetLineno(63)
			πTemp001 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp003, ß__dict__, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßDictProxyType.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label12
		Label11:
			// line 65: from types import MappingProxyType as DictProxyType
			πF.SetLineno(65)
			if πTemp001, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßMappingProxyType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDictProxyType.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label12
		Label12:
			goto Label6
		Label5:
			// line 67: import __builtin__
			πF.SetLineno(67)
			if πTemp001, πE = πg.ImportModule(πF, "__builtin__"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ß__builtin__.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 68: from pickle import Pickler as StockPickler, Unpickler as StockUnpickler
			πF.SetLineno(68)
			if πTemp001, πE = πg.ImportModule(πF, "pickle"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßPickler); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStockPickler.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßUnpickler); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStockUnpickler.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 69: from thread import LockType
			πF.SetLineno(69)
			if πTemp001, πE = πg.ImportModule(πF, "thread"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßLockType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLockType.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 70: from threading import _RLock as RLockType
			πF.SetLineno(70)
			if πTemp001, πE = πg.ImportModule(πF, "threading"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ß_RLock); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRLockType.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 71: from types import CodeType, FunctionType, ClassType, MethodType, \
			πF.SetLineno(71)
			if πTemp001, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßCodeType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCodeType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßFunctionType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFunctionType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßClassType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßClassType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßMethodType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMethodType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßGeneratorType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGeneratorType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßDictProxyType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDictProxyType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßXRangeType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßXRangeType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßSliceType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSliceType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßTracebackType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTracebackType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßNotImplementedType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNotImplementedType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßEllipsisType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßEllipsisType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßFrameType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFrameType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßModuleType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßModuleType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßBufferType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBufferType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßBuiltinMethodType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBuiltinMethodType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßTypeType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTypeType.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label6
		Label6:
			// line 75: from pickle import HIGHEST_PROTOCOL, PicklingError, UnpicklingError
			πF.SetLineno(75)
			if πTemp001, πE = πg.ImportModule(πF, "pickle"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßHIGHEST_PROTOCOL); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHIGHEST_PROTOCOL.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßPicklingError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPicklingError.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßUnpicklingError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnpicklingError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 76: try:
			πF.SetLineno(76)
			πF.PushCheckpoint(14)
			// line 77: from pickle import DEFAULT_PROTOCOL
			πF.SetLineno(77)
			if πTemp001, πE = πg.ImportModule(πF, "pickle"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßDEFAULT_PROTOCOL); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDEFAULT_PROTOCOL.ToObject(), πTemp005); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label13
		Label14:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp011, πTemp012 = πF.ExcInfo()
			if πTemp003, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label15
			}
			πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
			continue
			// line 78: except ImportError:
			πF.SetLineno(78)
		Label15:
			// line 79: DEFAULT_PROTOCOL = HIGHEST_PROTOCOL
			πF.SetLineno(79)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßHIGHEST_PROTOCOL); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDEFAULT_PROTOCOL.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label13
		Label13:
			// line 80: import __main__ as _main_module
			πF.SetLineno(80)
			if πTemp001, πE = πg.ImportModule(πF, "__main__"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ß_main_module.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 81: import marshal
			πF.SetLineno(81)
			if πTemp001, πE = πg.ImportModule(πF, "marshal"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßmarshal.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 82: import gc
			πF.SetLineno(82)
			if πTemp001, πE = πg.ImportModule(πF, "gc"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßgc.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 84: from weakref import ReferenceType, ProxyType, CallableProxyType
			πF.SetLineno(84)
			if πTemp001, πE = πg.ImportModule(πF, "weakref"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßReferenceType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReferenceType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßProxyType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßProxyType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßCallableProxyType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCallableProxyType.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 85: from functools import partial
			πF.SetLineno(85)
			if πTemp001, πE = πg.ImportModule(πF, "functools"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßpartial); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpartial.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 86: from operator import itemgetter, attrgetter
			πF.SetLineno(86)
			if πTemp001, πE = πg.ImportModule(πF, "operator"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßitemgetter); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßitemgetter.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßattrgetter); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßattrgetter.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßhexversion, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.LT(πF, πTemp006, πg.NewInt(50528256).ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label16
			}
			goto Label17
			// line 88: if sys.hexversion < 0x03030000:
			πF.SetLineno(88)
		Label16:
			// line 89: FileNotFoundError = IOError
			πF.SetLineno(89)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFileNotFoundError.ToObject(), πTemp003); πE != nil {
				continue
			}
			goto Label17
		Label17:
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			πTemp003 = πTemp005
			if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if !πTemp007 {
				goto Label18
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetAttr(πF, πTemp006, ßhexversion, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.LT(πF, πTemp009, πg.NewInt(50593792).ToObject()); πE != nil {
				continue
			}
			πTemp003 = πTemp005
		Label18:
			if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label19
			}
			goto Label20
			// line 90: if PY3 and sys.hexversion < 0x03040000:
			πF.SetLineno(90)
		Label19:
			// line 91: GENERATOR_FAIL = True
			πF.SetLineno(91)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGENERATOR_FAIL.ToObject(), πTemp003); πE != nil {
				continue
			}
			goto Label21
		Label20:
			// line 92: else: GENERATOR_FAIL = False
			πF.SetLineno(92)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGENERATOR_FAIL.ToObject(), πTemp003); πE != nil {
				continue
			}
			goto Label21
		Label21:
			if πTemp003, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label22
			}
			goto Label23
			// line 93: if PY3:
			πF.SetLineno(93)
		Label22:
			// line 94: import importlib.machinery
			πF.SetLineno(94)
			if πTemp001, πE = πg.ImportModule(πF, "importlib.machinery"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßimportlib.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 95: EXTENSION_SUFFIXES = tuple(importlib.machinery.EXTENSION_SUFFIXES)
			πF.SetLineno(95)
			πTemp001 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßimportlib); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßmachinery, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßEXTENSION_SUFFIXES, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßEXTENSION_SUFFIXES.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label24
		Label23:
			// line 97: import imp
			πF.SetLineno(97)
			if πTemp001, πE = πg.ImportModule(πF, "imp"); πE != nil {
				continue
			}
			πTemp003 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßimp.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 98: EXTENSION_SUFFIXES = tuple(suffix
			πF.SetLineno(98)
			πTemp001 = πF.MakeArgs(1)
			πTemp004 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsuffix *πg.Object = πg.UnboundLocal; _ = µsuffix
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µs_type *πg.Object = πg.UnboundLocal; _ = µs_type
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
						if πTemp002, πE = πg.ResolveGlobal(πF, ßimp); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_suffixes, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
								continue
							}
							µsuffix = πTemp003
							µ_ = πTemp006
							µs_type = πTemp007
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(1)            
						if πE = πg.CheckLocal(πF, µs_type, "s_type"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.ResolveGlobal(πF, ßimp); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßC_EXTENSION, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Eq(πF, µs_type, πTemp006); πE != nil {
							continue
						}
						if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp005 {
							goto Label4
						}
						goto Label5
						// line 98: EXTENSION_SUFFIXES = tuple(suffix
						πF.SetLineno(98)
					Label4:
						// line 98: EXTENSION_SUFFIXES = tuple(suffix
						πF.SetLineno(98)
						if πE = πg.CheckLocal(πF, µsuffix, "suffix"); πE != nil {
							continue
						}
						πF.PushCheckpoint(6)
						return µsuffix, nil
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
			if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßEXTENSION_SUFFIXES.ToObject(), πTemp006); πE != nil {
				continue
			}
			goto Label24
		Label24:
			// line 101: try:
			πF.SetLineno(101)
			πF.PushCheckpoint(26)
			// line 102: import ctypes
			πF.SetLineno(102)
			if πTemp001, πE = πg.ImportModule(πF, "ctypes"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßctypes.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 103: HAS_CTYPES = True
			πF.SetLineno(103)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHAS_CTYPES.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 105: IS_PYPY = not hasattr(ctypes, 'pythonapi')
			πF.SetLineno(105)
			πTemp001 = πF.MakeArgs(2)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			πTemp001[0] = πTemp006
			πTemp001[1] = ßpythonapi.ToObject()
			if πTemp006, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp007, πE = πg.IsTrue(πF, πTemp009); πE != nil {
				continue
			}
			πTemp005 = πg.GetBool(!πTemp007).ToObject()
			if πE = πF.Globals().SetItem(πF, ßIS_PYPY.ToObject(), πTemp005); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label25
		Label26:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp011, πTemp012 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label27
			}
			πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
			continue
			// line 106: except ImportError:
			πF.SetLineno(106)
		Label27:
			// line 107: HAS_CTYPES = False
			πF.SetLineno(107)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHAS_CTYPES.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 108: IS_PYPY = False
			πF.SetLineno(108)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIS_PYPY.ToObject(), πTemp005); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label25
		Label25:
			// line 109: NumpyUfuncType = None
			πF.SetLineno(109)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNumpyUfuncType.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 110: NumpyArrayType = None
			πF.SetLineno(110)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNumpyArrayType.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 111: try:
			πF.SetLineno(111)
			πF.PushCheckpoint(29)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßOLDER); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label30
			}
			goto Label31
			// line 112: if OLDER:
			πF.SetLineno(112)
		Label30:
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewStr("find_spec not found").ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			// line 113: raise AttributeError('find_spec not found')
			πF.SetLineno(113)
			πE = πF.Raise(πTemp006, nil, nil)
			continue
			goto Label31
		Label31:
			// line 114: import importlib
			πF.SetLineno(114)
			if πTemp001, πE = πg.ImportModule(πF, "importlib"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßimportlib.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßnumpy.ToObject()
			if πTemp006, πE = πg.ResolveGlobal(πF, ßimportlib); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetAttr(πF, πTemp006, ßmachinery, nil); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp009, ßPathFinder, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp006.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp009, ßfind_spec, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp007, πE = πg.IsTrue(πF, πTemp009); πE != nil {
				continue
			}
			πTemp005 = πg.GetBool(!πTemp007).ToObject()
			if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label32
			}
			goto Label33
			// line 115: if not importlib.machinery.PathFinder().find_spec('numpy'):
			πF.SetLineno(115)
		Label32:
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewStr("No module named 'numpy'").ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			// line 116: raise ImportError("No module named 'numpy'")
			πF.SetLineno(116)
			πE = πF.Raise(πTemp006, nil, nil)
			continue
			goto Label33
		Label33:
			// line 117: NumpyUfuncType = True
			πF.SetLineno(117)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNumpyUfuncType.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 118: NumpyArrayType = True
			πF.SetLineno(118)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNumpyArrayType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label28
		Label29:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp011, πTemp012 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label34
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label35
			}
			πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
			continue
			// line 119: except AttributeError:
			πF.SetLineno(119)
		Label34:
			// line 120: try:
			πF.SetLineno(120)
			πF.PushCheckpoint(37)
			// line 121: import imp
			πF.SetLineno(121)
			if πTemp001, πE = πg.ImportModule(πF, "imp"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßimp.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 122: imp.find_module('numpy')
			πF.SetLineno(122)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßnumpy.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßimp); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßfind_module, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			// line 123: NumpyUfuncType = True
			πF.SetLineno(123)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNumpyUfuncType.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 124: NumpyArrayType = True
			πF.SetLineno(124)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNumpyArrayType.ToObject(), πTemp005); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label36
		Label37:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp013, πTemp012 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label38
			}
			πE = πF.Raise(πTemp013.ToObject(), nil, πTemp012.ToObject())
			continue
			// line 125: except ImportError:
			πF.SetLineno(125)
		Label38:
			// line 126: pass
			πF.SetLineno(126)
			πF.RestoreExc(nil, nil)
			goto Label36
		Label36:
			πF.RestoreExc(nil, nil)
			goto Label28
			// line 127: except ImportError:
			πF.SetLineno(127)
		Label35:
			// line 128: pass
			πF.SetLineno(128)
			πF.RestoreExc(nil, nil)
			goto Label28
		Label28:
			// line 129: def __hook__():
			πF.SetLineno(129)
			πTemp004 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("__hook__", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 130: global NumpyArrayType, NumpyUfuncType
					πF.SetLineno(130)
					// line 131: from numpy import ufunc as NumpyUfuncType
					πF.SetLineno(131)
					if πTemp002, πE = πg.ImportModule(πF, "numpy"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßufunc); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ßNumpyUfuncType.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 132: from numpy import ndarray as NumpyArrayType
					πF.SetLineno(132)
					if πTemp002, πE = πg.ImportModule(πF, "numpy"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßndarray); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ßNumpyArrayType.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 133: return True
					πF.SetLineno(133)
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
			if πE = πF.Globals().SetItem(πF, ß__hook__.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNumpyArrayType); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label39
			}
			goto Label40
			// line 134: if NumpyArrayType: # then has numpy
			πF.SetLineno(134)
		Label39:
			// line 135: def ndarraysubclassinstance(obj):
			πF.SetLineno(135)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("ndarraysubclassinstance", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µcls *πg.Object = πg.UnboundLocal; _ = µcls
				var µNumpyInstance *πg.Object = πg.UnboundLocal; _ = µNumpyInstance
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.BaseException
				_ = πTemp009
				var πTemp010 *πg.Traceback
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßTypeType); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßClassType); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					if πTemp007, πE = πg.Contains(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label1
					}
					goto Label2
					// line 136: if type(obj) in (TypeType, ClassType):
					πF.SetLineno(136)
				Label1:
					// line 137: return False # all classes return False
					πF.SetLineno(137)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 138: try: # check if is ndarray, and elif is subclass of ndarray
					πF.SetLineno(138)
					πF.PushCheckpoint(4)
					// line 139: cls = getattr(obj, '__class__', None)
					πF.SetLineno(139)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					πTemp002[1] = ß__class__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µcls = πTemp003
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µcls == πTemp003).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeType); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µcls == πTemp003).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp008[0] = µcls
					πTemp008[1] = ßmro.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmro, nil); πE != nil {
						continue
					}
					πTemp008[2] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp007, πE = πg.Contains(πF, πTemp004, πg.NewStr("numpy.ndarray").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label7
					}
					goto Label8
					// line 140: if cls is None: return False
					πF.SetLineno(140)
				Label5:
					// line 140: if cls is None: return False
					πF.SetLineno(140)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label8
					// line 141: elif cls is TypeType: return False
					πF.SetLineno(141)
				Label6:
					// line 141: elif cls is TypeType: return False
					πF.SetLineno(141)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label8
					// line 142: elif 'numpy.ndarray' not in str(getattr(cls, 'mro', int.mro)()):
					πF.SetLineno(142)
				Label7:
					// line 143: return False
					πF.SetLineno(143)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label8
				Label8:
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßReferenceError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label9
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label10
					}
					πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
					continue
					// line 144: except ReferenceError: return False # handle 'R3' weakref in 3.x
					πF.SetLineno(144)
				Label9:
					// line 144: except ReferenceError: return False # handle 'R3' weakref in 3.x
					πF.SetLineno(144)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label3
					// line 145: except TypeError: return False
					πF.SetLineno(145)
				Label10:
					// line 145: except TypeError: return False
					πF.SetLineno(145)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					// line 147: __hook__() # import numpy (so the following works!!!)
					πF.SetLineno(147)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__hook__); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 149: NumpyInstance = NumpyArrayType((0,),'int8')
					πF.SetLineno(149)
					πTemp002 = πF.MakeArgs(2)
					πTemp001 = πg.NewTuple1(πg.NewInt(0).ToObject()).ToObject()
					πTemp002[0] = πTemp001
					πTemp002[1] = ßint8.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNumpyArrayType); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µNumpyInstance = πTemp003
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__reduce_ex__, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µNumpyInstance, "NumpyInstance"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µNumpyInstance, ß__reduce_ex__, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Eq(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label11
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__reduce__, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µNumpyInstance, "NumpyInstance"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µNumpyInstance, ß__reduce__, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Eq(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label11:
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label12
					}
					goto Label13
					// line 150: if id(obj.__reduce_ex__) == id(NumpyInstance.__reduce_ex__) and \
					πF.SetLineno(150)
				Label12:
					// line 151: id(obj.__reduce__) == id(NumpyInstance.__reduce__): return True
					πF.SetLineno(151)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label13
				Label13:
					// line 152: return False
					πF.SetLineno(152)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßndarraysubclassinstance.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 153: def numpyufunc(obj):
			πF.SetLineno(153)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("numpyufunc", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µcls *πg.Object = πg.UnboundLocal; _ = µcls
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.BaseException
				_ = πTemp009
				var πTemp010 *πg.Traceback
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßTypeType); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßClassType); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					if πTemp007, πE = πg.Contains(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label1
					}
					goto Label2
					// line 154: if type(obj) in (TypeType, ClassType):
					πF.SetLineno(154)
				Label1:
					// line 155: return False # all classes return False
					πF.SetLineno(155)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 156: try: # check if is ufunc
					πF.SetLineno(156)
					πF.PushCheckpoint(4)
					// line 157: cls = getattr(obj, '__class__', None)
					πF.SetLineno(157)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					πTemp002[1] = ß__class__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µcls = πTemp003
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µcls == πTemp003).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeType); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µcls == πTemp003).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					goto Label7
					// line 158: if cls is None: return False
					πF.SetLineno(158)
				Label5:
					// line 158: if cls is None: return False
					πF.SetLineno(158)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label7
					// line 159: elif cls is TypeType: return False
					πF.SetLineno(159)
				Label6:
					// line 159: elif cls is TypeType: return False
					πF.SetLineno(159)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label7
				Label7:
					πTemp002 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp008[0] = µcls
					πTemp008[1] = ßmro.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmro, nil); πE != nil {
						continue
					}
					πTemp008[2] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp007, πE = πg.Contains(πF, πTemp004, πg.NewStr("numpy.ufunc").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label8
					}
					goto Label9
					// line 160: if 'numpy.ufunc' not in str(getattr(cls, 'mro', int.mro)()):
					πF.SetLineno(160)
				Label8:
					// line 161: return False
					πF.SetLineno(161)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label9
				Label9:
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßReferenceError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label10
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label11
					}
					πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
					continue
					// line 162: except ReferenceError: return False # handle 'R3' weakref in 3.x
					πF.SetLineno(162)
				Label10:
					// line 162: except ReferenceError: return False # handle 'R3' weakref in 3.x
					πF.SetLineno(162)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label3
					// line 163: except TypeError: return False
					πF.SetLineno(163)
				Label11:
					// line 163: except TypeError: return False
					πF.SetLineno(163)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					// line 165: return True
					πF.SetLineno(165)
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
			if πE = πF.Globals().SetItem(πF, ßnumpyufunc.ToObject(), πTemp009); πE != nil {
				continue
			}
			goto Label41
		Label40:
			// line 167: def ndarraysubclassinstance(obj): return False
			πF.SetLineno(167)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("ndarraysubclassinstance", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 167: def ndarraysubclassinstance(obj): return False
					πF.SetLineno(167)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßndarraysubclassinstance.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 168: def numpyufunc(obj): return False
			πF.SetLineno(168)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("numpyufunc", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 168: def numpyufunc(obj): return False
					πF.SetLineno(168)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßnumpyufunc.ToObject(), πTemp014); πE != nil {
				continue
			}
			goto Label41
		Label41:
			if πTemp015, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp015); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label42
			}
			goto Label43
			// line 171: if PY3:
			πF.SetLineno(171)
		Label42:
			// line 172: CellType = type((lambda x: lambda y: x)(0).__closure__[0])
			πF.SetLineno(172)
			πTemp001 = πF.MakeArgs(1)
			πTemp015 = πg.NewInt(0).ToObject()
			πTemp017 = πF.MakeArgs(1)
			πTemp017[0] = πg.NewInt(0).ToObject()
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 172: CellType = type((lambda x: lambda y: x)(0).__closure__[0])
					πF.SetLineno(172)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "y", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µy *πg.Object = πArgs[0]; _ = µy
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 172: CellType = type((lambda x: lambda y: x)(0).__closure__[0])
							πF.SetLineno(172)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πR = µx
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
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
			if πTemp019, πE = πTemp018.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp018, πE = πg.GetAttr(πF, πTemp019, ß__closure__, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetItem(πF, πTemp018, πTemp015); πE != nil {
				continue
			}
			πTemp001[0] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßCellType.ToObject(), πTemp016); πE != nil {
				continue
			}
			goto Label44
		Label43:
			// line 174: CellType = type((lambda x: lambda y: x)(0).func_closure[0])
			πF.SetLineno(174)
			πTemp001 = πF.MakeArgs(1)
			πTemp015 = πg.NewInt(0).ToObject()
			πTemp017 = πF.MakeArgs(1)
			πTemp017[0] = πg.NewInt(0).ToObject()
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 174: CellType = type((lambda x: lambda y: x)(0).func_closure[0])
					πF.SetLineno(174)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "y", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µy *πg.Object = πArgs[0]; _ = µy
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 174: CellType = type((lambda x: lambda y: x)(0).func_closure[0])
							πF.SetLineno(174)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πR = µx
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
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
			if πTemp019, πE = πTemp018.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp018, πE = πg.GetAttr(πF, πTemp019, ßfunc_closure, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetItem(πF, πTemp018, πTemp015); πE != nil {
				continue
			}
			πTemp001[0] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßCellType.ToObject(), πTemp016); πE != nil {
				continue
			}
			goto Label44
		Label44:
			if πTemp016, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp018, πE = πg.GetAttr(πF, πTemp016, ßhexversion, nil); πE != nil {
				continue
			}
			if πTemp015, πE = πg.GE(πF, πTemp018, πg.NewInt(33882352).ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp015); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label45
			}
			goto Label46
			// line 176: if sys.hexversion >= 0x20500f0:
			πF.SetLineno(176)
		Label45:
			// line 177: from types import GetSetDescriptorType
			πF.SetLineno(177)
			if πTemp001, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp015 = πTemp001[0]
			if πTemp016, πE = πg.GetAttrImport(πF, πTemp015, ßGetSetDescriptorType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGetSetDescriptorType.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			πTemp015 = πg.GetBool(!πTemp007).ToObject()
			if πTemp007, πE = πg.IsTrue(πF, πTemp015); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label47
			}
			goto Label48
			// line 178: if not IS_PYPY:
			πF.SetLineno(178)
		Label47:
			// line 179: from types import MemberDescriptorType
			πF.SetLineno(179)
			if πTemp001, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp015 = πTemp001[0]
			if πTemp016, πE = πg.GetAttrImport(πF, πTemp015, ßMemberDescriptorType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMemberDescriptorType.ToObject(), πTemp016); πE != nil {
				continue
			}
			goto Label49
		Label48:
			// line 183: class _member(object):
			πF.SetLineno(183)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp018, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp018
			πTemp020 = πg.NewDict()
			if πTemp015, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ß__module__.ToObject(), πTemp015); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_member", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp020
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 184: __slots__ = ['descriptor']
					πF.SetLineno(184)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = ßdescriptor.ToObject()
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp002); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp016, πE = πTemp020.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp016 == nil {
				πTemp016 = πg.TypeType.ToObject()
			}
			if πTemp018, πE = πTemp016.Call(πF, []*πg.Object{πg.NewStr("_member").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp020.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_member.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 185: MemberDescriptorType = type(_member.descriptor)
			πF.SetLineno(185)
			πTemp001 = πF.MakeArgs(1)
			if πTemp015, πE = πg.ResolveGlobal(πF, ß_member); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßdescriptor, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßMemberDescriptorType.ToObject(), πTemp016); πE != nil {
				continue
			}
			goto Label49
		Label49:
			goto Label46
		Label46:
			if πTemp015, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp015); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label50
			}
			goto Label51
			// line 186: if IS_PYPY:
			πF.SetLineno(186)
		Label50:
			// line 187: WrapperDescriptorType = MethodType
			πF.SetLineno(187)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßMethodType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWrapperDescriptorType.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 188: MethodDescriptorType = FunctionType
			πF.SetLineno(188)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßFunctionType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMethodDescriptorType.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 189: ClassMethodDescriptorType = FunctionType
			πF.SetLineno(189)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßFunctionType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßClassMethodDescriptorType.ToObject(), πTemp015); πE != nil {
				continue
			}
			goto Label52
		Label51:
			// line 191: WrapperDescriptorType = type(type.__repr__)
			πF.SetLineno(191)
			πTemp001 = πF.MakeArgs(1)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp015, ß__repr__, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßWrapperDescriptorType.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 192: MethodDescriptorType = type(type.__dict__['mro'])
			πF.SetLineno(192)
			πTemp001 = πF.MakeArgs(1)
			πTemp015 = ßmro.ToObject()
			if πTemp018, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp018, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetItem(πF, πTemp019, πTemp015); πE != nil {
				continue
			}
			πTemp001[0] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßMethodDescriptorType.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 193: ClassMethodDescriptorType = type(type.__dict__['__prepare__' if PY3 else 'mro'])
			πF.SetLineno(193)
			πTemp001 = πF.MakeArgs(1)
			if πTemp018, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp018); πE != nil {
				continue
			}
			if !πTemp007 {
				goto Label53
			}
			πTemp016 = ß__prepare__.ToObject()
			goto Label54
		Label53:
			πTemp016 = ßmro.ToObject()
		Label54:
			πTemp015 = πTemp016
			if πTemp018, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp018, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetItem(πF, πTemp019, πTemp015); πE != nil {
				continue
			}
			πTemp001[0] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßClassMethodDescriptorType.ToObject(), πTemp016); πE != nil {
				continue
			}
			goto Label52
		Label52:
			// line 195: MethodWrapperType = type([].__repr__)
			πF.SetLineno(195)
			πTemp001 = πF.MakeArgs(1)
			πTemp017 = make([]*πg.Object, 0)
			πTemp015 = πg.NewList(πTemp017...).ToObject()
			if πTemp016, πE = πg.GetAttr(πF, πTemp015, ß__repr__, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßMethodWrapperType.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 196: PartialType = type(partial(int,base=2))
			πF.SetLineno(196)
			πTemp001 = πF.MakeArgs(1)
			πTemp017 = πF.MakeArgs(1)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
				continue
			}
			πTemp017[0] = πTemp015
			πTemp021 = πg.KWArgs{
				{"base", πg.NewInt(2).ToObject()},
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßpartial); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp017, πTemp021); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			πTemp001[0] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßPartialType.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 197: SuperType = type(super(Exception, TypeError()))
			πF.SetLineno(197)
			πTemp001 = πF.MakeArgs(1)
			πTemp017 = πF.MakeArgs(2)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp017[0] = πTemp015
			if πTemp015, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp017[1] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			πTemp001[0] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßSuperType.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 198: ItemGetterType = type(itemgetter(0))
			πF.SetLineno(198)
			πTemp001 = πF.MakeArgs(1)
			πTemp017 = πF.MakeArgs(1)
			πTemp017[0] = πg.NewInt(0).ToObject()
			if πTemp015, πE = πg.ResolveGlobal(πF, ßitemgetter); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			πTemp001[0] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßItemGetterType.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 199: AttrGetterType = type(attrgetter('__repr__'))
			πF.SetLineno(199)
			πTemp001 = πF.MakeArgs(1)
			πTemp017 = πF.MakeArgs(1)
			πTemp017[0] = ß__repr__.ToObject()
			if πTemp015, πE = πg.ResolveGlobal(πF, ßattrgetter); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			πTemp001[0] = πTemp016
			if πTemp015, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßAttrGetterType.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 201: def get_file_type(*args, **kwargs):
			πF.SetLineno(201)
			πTemp004 = make([]πg.Param, 0)
			πTemp015 = πg.NewFunction(πg.NewCode("get_file_type", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
				var µopen *πg.Object = πg.UnboundLocal; _ = µopen
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µt *πg.Object = πg.UnboundLocal; _ = µt
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
					// line 202: open = kwargs.pop("open", __builtin__.open)
					πF.SetLineno(202)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßopen.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß__builtin__); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßopen, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µkwargs, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µopen = πTemp003
					// line 203: f = open(os.devnull, *args, **kwargs)
					πF.SetLineno(203)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdevnull, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µopen, "open"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, µopen, πTemp001, µargs, nil, µkwargs); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µf = πTemp002
					// line 204: t = type(f)
					πF.SetLineno(204)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp001[0] = µf
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µt = πTemp003
					// line 205: f.close()
					πF.SetLineno(205)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 206: return t
					πF.SetLineno(206)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πR = µt
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßget_file_type.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 208: FileType = get_file_type('rb', buffering=0)
			πF.SetLineno(208)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßrb.ToObject()
			πTemp021 = πg.KWArgs{
				{"buffering", πg.NewInt(0).ToObject()},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßget_file_type); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp001, πTemp021); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßFileType.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 209: TextWrapperType = get_file_type('r', buffering=-1)
			πF.SetLineno(209)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßr.ToObject()
			if πTemp016, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp021 = πg.KWArgs{
				{"buffering", πTemp016},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßget_file_type); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp001, πTemp021); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßTextWrapperType.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 210: BufferedRandomType = get_file_type('r+b', buffering=-1)
			πF.SetLineno(210)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewStr("r+b").ToObject()
			if πTemp016, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp021 = πg.KWArgs{
				{"buffering", πTemp016},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßget_file_type); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp001, πTemp021); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßBufferedRandomType.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 211: BufferedReaderType = get_file_type('rb', buffering=-1)
			πF.SetLineno(211)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßrb.ToObject()
			if πTemp016, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp021 = πg.KWArgs{
				{"buffering", πTemp016},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßget_file_type); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp001, πTemp021); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßBufferedReaderType.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 212: BufferedWriterType = get_file_type('wb', buffering=-1)
			πF.SetLineno(212)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßwb.ToObject()
			if πTemp016, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp021 = πg.KWArgs{
				{"buffering", πTemp016},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßget_file_type); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp001, πTemp021); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßBufferedWriterType.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 213: try:
			πF.SetLineno(213)
			πF.PushCheckpoint(56)
			// line 214: from _pyio import open as _open
			πF.SetLineno(214)
			if πTemp001, πE = πg.ImportModule(πF, "_pyio"); πE != nil {
				continue
			}
			πTemp016 = πTemp001[0]
			if πTemp018, πE = πg.GetAttrImport(πF, πTemp016, ßopen); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_open.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 215: PyTextWrapperType = get_file_type('r', buffering=-1, open=_open)
			πF.SetLineno(215)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßr.ToObject()
			if πTemp016, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ß_open); πE != nil {
				continue
			}
			πTemp021 = πg.KWArgs{
				{"buffering", πTemp016},
				{"open", πTemp018},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßget_file_type); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp001, πTemp021); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßPyTextWrapperType.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 216: PyBufferedRandomType = get_file_type('r+b', buffering=-1, open=_open)
			πF.SetLineno(216)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewStr("r+b").ToObject()
			if πTemp016, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ß_open); πE != nil {
				continue
			}
			πTemp021 = πg.KWArgs{
				{"buffering", πTemp016},
				{"open", πTemp018},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßget_file_type); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp001, πTemp021); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßPyBufferedRandomType.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 217: PyBufferedReaderType = get_file_type('rb', buffering=-1, open=_open)
			πF.SetLineno(217)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßrb.ToObject()
			if πTemp016, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ß_open); πE != nil {
				continue
			}
			πTemp021 = πg.KWArgs{
				{"buffering", πTemp016},
				{"open", πTemp018},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßget_file_type); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp001, πTemp021); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßPyBufferedReaderType.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 218: PyBufferedWriterType = get_file_type('wb', buffering=-1, open=_open)
			πF.SetLineno(218)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßwb.ToObject()
			if πTemp016, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ß_open); πE != nil {
				continue
			}
			πTemp021 = πg.KWArgs{
				{"buffering", πTemp016},
				{"open", πTemp018},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßget_file_type); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp001, πTemp021); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßPyBufferedWriterType.ToObject(), πTemp018); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label55
		Label56:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp011, πTemp012 = πF.ExcInfo()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label57
			}
			πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
			continue
			// line 219: except ImportError:
			πF.SetLineno(219)
		Label57:
			// line 220: PyTextWrapperType = PyBufferedRandomType = PyBufferedReaderType = PyBufferedWriterType = None
			πF.SetLineno(220)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPyTextWrapperType.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPyBufferedRandomType.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPyBufferedReaderType.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPyBufferedWriterType.ToObject(), πTemp016); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label55
		Label55:
			// line 221: try:
			πF.SetLineno(221)
			πF.PushCheckpoint(59)
			// line 222: from cStringIO import StringIO, InputType, OutputType
			πF.SetLineno(222)
			if πTemp001, πE = πg.ImportModule(πF, "cStringIO"); πE != nil {
				continue
			}
			πTemp016 = πTemp001[0]
			if πTemp018, πE = πg.GetAttrImport(πF, πTemp016, ßStringIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp018); πE != nil {
				continue
			}
			πTemp016 = πTemp001[0]
			if πTemp018, πE = πg.GetAttrImport(πF, πTemp016, ßInputType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßInputType.ToObject(), πTemp018); πE != nil {
				continue
			}
			πTemp016 = πTemp001[0]
			if πTemp018, πE = πg.GetAttrImport(πF, πTemp016, ßOutputType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOutputType.ToObject(), πTemp018); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label58
		Label59:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp011, πTemp012 = πF.ExcInfo()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label60
			}
			πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
			continue
			// line 223: except ImportError:
			πF.SetLineno(223)
		Label60:
			if πTemp016, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label61
			}
			goto Label62
			// line 224: if PY3:
			πF.SetLineno(224)
		Label61:
			// line 225: from io import BytesIO as StringIO
			πF.SetLineno(225)
			if πTemp001, πE = πg.ImportModule(πF, "io"); πE != nil {
				continue
			}
			πTemp016 = πTemp001[0]
			if πTemp018, πE = πg.GetAttrImport(πF, πTemp016, ßBytesIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp018); πE != nil {
				continue
			}
			goto Label63
		Label62:
			// line 227: from StringIO import StringIO
			πF.SetLineno(227)
			if πTemp001, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
				continue
			}
			πTemp016 = πTemp001[0]
			if πTemp018, πE = πg.GetAttrImport(πF, πTemp016, ßStringIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp018); πE != nil {
				continue
			}
			goto Label63
		Label63:
			// line 228: InputType = OutputType = None
			πF.SetLineno(228)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßInputType.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOutputType.ToObject(), πTemp016); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label58
		Label58:
			if πTemp018, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp018); πE != nil {
				continue
			}
			πTemp016 = πg.GetBool(!πTemp007).ToObject()
			if πTemp007, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label64
			}
			goto Label65
			// line 229: if not IS_PYPY:
			πF.SetLineno(229)
		Label64:
			// line 230: from socket import socket as SocketType
			πF.SetLineno(230)
			if πTemp001, πE = πg.ImportModule(πF, "socket"); πE != nil {
				continue
			}
			πTemp016 = πTemp001[0]
			if πTemp018, πE = πg.GetAttrImport(πF, πTemp016, ßsocket); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSocketType.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 231: try: #FIXME: additionally calls ForkingPickler.register several times
			πF.SetLineno(231)
			πF.PushCheckpoint(67)
			// line 232: from multiprocessing.reduction import _reduce_socket as reduce_socket
			πF.SetLineno(232)
			if πTemp001, πE = πg.ImportModule(πF, "multiprocessing.reduction"); πE != nil {
				continue
			}
			πTemp016 = πTemp001[1]
			if πTemp018, πE = πg.GetAttrImport(πF, πTemp016, ß_reduce_socket); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreduce_socket.ToObject(), πTemp018); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label66
		Label67:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp011, πTemp012 = πF.ExcInfo()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label68
			}
			πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
			continue
			// line 233: except ImportError:
			πF.SetLineno(233)
		Label68:
			// line 234: from multiprocessing.reduction import reduce_socket
			πF.SetLineno(234)
			if πTemp001, πE = πg.ImportModule(πF, "multiprocessing.reduction"); πE != nil {
				continue
			}
			πTemp016 = πTemp001[1]
			if πTemp018, πE = πg.GetAttrImport(πF, πTemp016, ßreduce_socket); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreduce_socket.ToObject(), πTemp018); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label66
		Label66:
			goto Label65
		Label65:
			// line 235: try:
			πF.SetLineno(235)
			πF.PushCheckpoint(70)
			// line 236: __IPYTHON__ is True # is ipython
			πF.SetLineno(236)
			if πTemp018, πE = πg.ResolveGlobal(πF, ß__IPYTHON__); πE != nil {
				continue
			}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp016 = πg.GetBool(πTemp018 == πTemp019).ToObject()
			// line 237: ExitType = None     # IPython.core.autocall.ExitAutocall
			πF.SetLineno(237)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßExitType.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 238: singletontypes = ['exit', 'quit', 'get_ipython']
			πF.SetLineno(238)
			πTemp001 = make([]*πg.Object, 3)
			πTemp001[0] = ßexit.ToObject()
			πTemp001[1] = ßquit.ToObject()
			πTemp001[2] = ßget_ipython.ToObject()
			πTemp016 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsingletontypes.ToObject(), πTemp016); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label69
		Label70:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp011, πTemp012 = πF.ExcInfo()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label71
			}
			πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
			continue
			// line 239: except NameError:
			πF.SetLineno(239)
		Label71:
			// line 240: try: ExitType = type(exit) # apparently 'exit' can be removed
			πF.SetLineno(240)
			πF.PushCheckpoint(73)
			// line 240: try: ExitType = type(exit) # apparently 'exit' can be removed
			πF.SetLineno(240)
			πTemp001 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßexit); πE != nil {
				continue
			}
			πTemp001[0] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßExitType.ToObject(), πTemp018); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label72
		Label73:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp013, πTemp012 = πF.ExcInfo()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label74
			}
			πE = πF.Raise(πTemp013.ToObject(), nil, πTemp012.ToObject())
			continue
			// line 241: except NameError: ExitType = None
			πF.SetLineno(241)
		Label74:
			// line 241: except NameError: ExitType = None
			πF.SetLineno(241)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßExitType.ToObject(), πTemp016); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label72
		Label72:
			// line 242: singletontypes = []
			πF.SetLineno(242)
			πTemp001 = make([]*πg.Object, 0)
			πTemp016 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsingletontypes.ToObject(), πTemp016); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label69
		Label69:
			// line 247: HANDLE_FMODE = 0
			πF.SetLineno(247)
			if πE = πF.Globals().SetItem(πF, ßHANDLE_FMODE.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			// line 251: CONTENTS_FMODE = 1
			πF.SetLineno(251)
			if πE = πF.Globals().SetItem(πF, ßCONTENTS_FMODE.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			// line 253: FILE_FMODE = 2
			πF.SetLineno(253)
			if πE = πF.Globals().SetItem(πF, ßFILE_FMODE.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
				continue
			}
			// line 256: def copy(obj, *args, **kwds):
			πF.SetLineno(256)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("copy", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µargs *πg.Object = πArgs[1]; _ = µargs
				var µkwds *πg.Object = πArgs[2]; _ = µkwds
				var µignore *πg.Object = πg.UnboundLocal; _ = µignore
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
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
					// line 257: """use pickling to 'copy' an object"""
					πF.SetLineno(257)
					// line 258: ignore = kwds.pop('ignore', Unpickler.settings['ignore'])
					πF.SetLineno(258)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßignore.ToObject()
					πTemp002 = ßignore.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßUnpickler); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsettings, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
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
					µignore = πTemp003
					// line 259: return loads(dumps(obj, *args, **kwds), ignore=ignore)
					πF.SetLineno(259)
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006[0] = µobj
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdumps); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp006, µargs, nil, µkwds); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"ignore", µignore},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßloads); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcopy.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 257: """use pickling to 'copy' an object"""
			πF.SetLineno(257)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("use pickling to 'copy' an object").ToObject()); πE != nil {
				continue
			}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßcopy); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
				continue
			}
			// line 261: def dump(obj, file, protocol=None, byref=None, fmode=None, recurse=None, **kwds):#, strictio=None):
			πF.SetLineno(261)
			πTemp004 = make([]πg.Param, 6)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "file", Def: nil}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "protocol", Def: πTemp019}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "byref", Def: πTemp019}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[4] = πg.Param{Name: "fmode", Def: πTemp019}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[5] = πg.Param{Name: "recurse", Def: πTemp019}
			πTemp018 = πg.NewFunction(πg.NewCode("dump", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µfile *πg.Object = πArgs[1]; _ = µfile
				var µprotocol *πg.Object = πArgs[2]; _ = µprotocol
				var µbyref *πg.Object = πArgs[3]; _ = µbyref
				var µfmode *πg.Object = πArgs[4]; _ = µfmode
				var µrecurse *πg.Object = πArgs[5]; _ = µrecurse
				var µkwds *πg.Object = πArgs[6]; _ = µkwds
				var µsettings *πg.Object = πg.UnboundLocal; _ = µsettings
				var µ_kwds *πg.Object = πg.UnboundLocal; _ = µ_kwds
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
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 262: """pickle an object to a file"""
					πF.SetLineno(262)
					// line 263: from .settings import settings
					πF.SetLineno(263)
					if πTemp002, πE = πg.ImportModule(πF, ".settings"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßsettings); πE != nil {
						continue
					}
					µsettings = πTemp003
					// line 264: protocol = settings['protocol'] if protocol is None else int(protocol)
					πF.SetLineno(264)
					if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µprotocol == πTemp004).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label1
					}
					πTemp003 = ßprotocol.ToObject()
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µsettings, πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp004
					goto Label2
				Label1:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
						continue
					}
					πTemp002[0] = µprotocol
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πTemp004
				Label2:
					µprotocol = πTemp001
					// line 265: _kwds = kwds.copy()
					πF.SetLineno(265)
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µkwds, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µ_kwds = πTemp003
					// line 266: _kwds.update(dict(byref=byref, fmode=fmode, recurse=recurse))
					πF.SetLineno(266)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbyref, "byref"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfmode, "fmode"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"byref", µbyref},
						{"fmode", µfmode},
						{"recurse", µrecurse},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, πTemp006); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µ_kwds, "_kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µ_kwds, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 267: Pickler(file, protocol, **_kwds).dump(obj)
					πF.SetLineno(267)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp007[0] = µfile
					if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
						continue
					}
					πTemp007[1] = µprotocol
					if πE = πg.CheckLocal(πF, µ_kwds, "_kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp001, πTemp007, nil, nil, µ_kwds); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßdump, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 268: return
					πF.SetLineno(268)
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
			if πE = πF.Globals().SetItem(πF, ßdump.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 262: """pickle an object to a file"""
			πF.SetLineno(262)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πg.NewStr("pickle an object to a file").ToObject()); πE != nil {
				continue
			}
			if πTemp022, πE = πg.ResolveGlobal(πF, ßdump); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp022, ß__doc__, πTemp019); πE != nil {
				continue
			}
			// line 270: def dumps(obj, protocol=None, byref=None, fmode=None, recurse=None, **kwds):#, strictio=None):
			πF.SetLineno(270)
			πTemp004 = make([]πg.Param, 5)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp022, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "protocol", Def: πTemp022}
			if πTemp022, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "byref", Def: πTemp022}
			if πTemp022, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "fmode", Def: πTemp022}
			if πTemp022, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[4] = πg.Param{Name: "recurse", Def: πTemp022}
			πTemp019 = πg.NewFunction(πg.NewCode("dumps", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µprotocol *πg.Object = πArgs[1]; _ = µprotocol
				var µbyref *πg.Object = πArgs[2]; _ = µbyref
				var µfmode *πg.Object = πArgs[3]; _ = µfmode
				var µrecurse *πg.Object = πArgs[4]; _ = µrecurse
				var µkwds *πg.Object = πArgs[5]; _ = µkwds
				var µfile *πg.Object = πg.UnboundLocal; _ = µfile
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 271: """pickle an object to a string"""
					πF.SetLineno(271)
					// line 272: file = StringIO()
					πF.SetLineno(272)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfile = πTemp002
					// line 273: dump(obj, file, protocol, byref, fmode, recurse, **kwds)#, strictio)
					πF.SetLineno(273)
					πTemp003 = πF.MakeArgs(6)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp003[1] = µfile
					if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
						continue
					}
					πTemp003[2] = µprotocol
					if πE = πg.CheckLocal(πF, µbyref, "byref"); πE != nil {
						continue
					}
					πTemp003[3] = µbyref
					if πE = πg.CheckLocal(πF, µfmode, "fmode"); πE != nil {
						continue
					}
					πTemp003[4] = µfmode
					if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
						continue
					}
					πTemp003[5] = µrecurse
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdump); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, πTemp003, nil, nil, µkwds); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 274: return file.getvalue()
					πF.SetLineno(274)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfile, ßgetvalue, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdumps.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 271: """pickle an object to a string"""
			πF.SetLineno(271)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πg.NewStr("pickle an object to a string").ToObject()); πE != nil {
				continue
			}
			if πTemp023, πE = πg.ResolveGlobal(πF, ßdumps); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp023, ß__doc__, πTemp022); πE != nil {
				continue
			}
			// line 276: def load(file, ignore=None, **kwds):
			πF.SetLineno(276)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "file", Def: nil}
			if πTemp023, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "ignore", Def: πTemp023}
			πTemp022 = πg.NewFunction(πg.NewCode("load", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfile *πg.Object = πArgs[0]; _ = µfile
				var µignore *πg.Object = πArgs[1]; _ = µignore
				var µkwds *πg.Object = πArgs[2]; _ = µkwds
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 πg.KWArgs
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
					// line 277: """unpickle an object from a file"""
					πF.SetLineno(277)
					// line 278: return Unpickler(file, ignore=ignore, **kwds).load()
					πF.SetLineno(278)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp001[0] = µfile
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp002 = πg.KWArgs{
						{"ignore", µignore},
					}
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßUnpickler); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Invoke(πF, πTemp003, πTemp001, nil, πTemp002, µkwds); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßload, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßload.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 277: """unpickle an object from a file"""
			πF.SetLineno(277)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πg.NewStr("unpickle an object from a file").ToObject()); πE != nil {
				continue
			}
			if πTemp024, πE = πg.ResolveGlobal(πF, ßload); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp024, ß__doc__, πTemp023); πE != nil {
				continue
			}
			// line 280: def loads(str, ignore=None, **kwds):
			πF.SetLineno(280)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "str", Def: nil}
			if πTemp024, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "ignore", Def: πTemp024}
			πTemp023 = πg.NewFunction(πg.NewCode("loads", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstr *πg.Object = πArgs[0]; _ = µstr
				var µignore *πg.Object = πArgs[1]; _ = µignore
				var µkwds *πg.Object = πArgs[2]; _ = µkwds
				var µfile *πg.Object = πg.UnboundLocal; _ = µfile
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
					// line 281: """unpickle an object from a string"""
					πF.SetLineno(281)
					// line 282: file = StringIO(str)
					πF.SetLineno(282)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
						continue
					}
					πTemp001[0] = µstr
					if πTemp002, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfile = πTemp003
					// line 283: return load(file, ignore, **kwds)
					πF.SetLineno(283)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp001[0] = µfile
					if πE = πg.CheckLocal(πF, µignore, "ignore"); πE != nil {
						continue
					}
					πTemp001[1] = µignore
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßload); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, nil, nil, µkwds); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßloads.ToObject(), πTemp023); πE != nil {
				continue
			}
			// line 281: """unpickle an object from a string"""
			πF.SetLineno(281)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πg.NewStr("unpickle an object from a string").ToObject()); πE != nil {
				continue
			}
			if πTemp025, πE = πg.ResolveGlobal(πF, ßloads); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp024); πE != nil {
				continue
			}
			// line 296: def _module_map():
			πF.SetLineno(296)
			πTemp004 = make([]πg.Param, 0)
			πTemp024 = πg.NewFunction(πg.NewCode("_module_map", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdefaultdict *πg.Object = πg.UnboundLocal; _ = µdefaultdict
				var µmodmap *πg.Object = πg.UnboundLocal; _ = µmodmap
				var µitems *πg.Object = πg.UnboundLocal; _ = µitems
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var µobjname *πg.Object = πg.UnboundLocal; _ = µobjname
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
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
					case 9: goto Label9
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 297: """get map of imported modules"""
					πF.SetLineno(297)
					// line 298: from collections import defaultdict
					πF.SetLineno(298)
					if πTemp002, πE = πg.ImportModule(πF, "collections"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßdefaultdict); πE != nil {
						continue
					}
					µdefaultdict = πTemp003
					// line 299: modmap = defaultdict(list)
					πF.SetLineno(299)
					πTemp002 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µdefaultdict, "defaultdict"); πE != nil {
						continue
					}
					if πTemp001, πE = µdefaultdict.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmodmap = πTemp001
					// line 300: items = 'items' if PY3 else 'iteritems'
					πF.SetLineno(300)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					πTemp001 = ßitems.ToObject()
					goto Label2
				Label1:
					πTemp001 = ßiteritems.ToObject()
				Label2:
					µitems = πTemp001
					πTemp002 = πF.MakeArgs(2)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßmodules, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp005
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					πTemp002[1] = µitems
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp004 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label5
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µname = πTemp005
						µmodule = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(3)            
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µmodule == πTemp005).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					goto Label7
					// line 302: if module is None:
					πF.SetLineno(302)
				Label6:
					// line 303: continue
					πF.SetLineno(303)
					continue
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µmodule, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßitems, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(9)
					πTemp006 = false
				Label8:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label10
					}
					if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp009}}}, πTemp005); πE != nil {
							continue
						}
						µobjname = πTemp007
						µobj = πTemp009
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(8)            
					// line 305: modmap[objname].append((obj, name))
					πF.SetLineno(305)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(µobj, µname).ToObject()
					πTemp002[0] = πTemp005
					if πE = πg.CheckLocal(πF, µobjname, "objname"); πE != nil {
						continue
					}
					πTemp005 = µobjname
					if πE = πg.CheckLocal(πF, µmodmap, "modmap"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µmodmap, πTemp005); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßappend, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					continue
				Label9:
					if πE != nil || πR != nil {
						continue
					}
				Label10:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 306: return modmap
					πF.SetLineno(306)
					if πE = πg.CheckLocal(πF, µmodmap, "modmap"); πE != nil {
						continue
					}
					πR = µmodmap
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_module_map.ToObject(), πTemp024); πE != nil {
				continue
			}
			// line 297: """get map of imported modules"""
			πF.SetLineno(297)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πg.NewStr("get map of imported modules").ToObject()); πE != nil {
				continue
			}
			if πTemp026, πE = πg.ResolveGlobal(πF, ß_module_map); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp026, ß__doc__, πTemp025); πE != nil {
				continue
			}
			// line 308: def _lookup_module(modmap, name, obj, main_module): #FIXME: needs work
			πF.SetLineno(308)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "modmap", Def: nil}
			πTemp004[1] = πg.Param{Name: "name", Def: nil}
			πTemp004[2] = πg.Param{Name: "obj", Def: nil}
			πTemp004[3] = πg.Param{Name: "main_module", Def: nil}
			πTemp025 = πg.NewFunction(πg.NewCode("_lookup_module", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmodmap *πg.Object = πArgs[0]; _ = µmodmap
				var µname *πg.Object = πArgs[1]; _ = µname
				var µobj *πg.Object = πArgs[2]; _ = µobj
				var µmain_module *πg.Object = πArgs[3]; _ = µmain_module
				var µmodobj *πg.Object = πg.UnboundLocal; _ = µmodobj
				var µmodname *πg.Object = πg.UnboundLocal; _ = µmodname
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 309: """lookup name if module is imported"""
					πF.SetLineno(309)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002 = µname
					if πE = πg.CheckLocal(πF, µmodmap, "modmap"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µmodmap, πTemp002); πE != nil {
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
						µmodobj = πTemp003
						µmodname = πTemp006
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µmodobj, "modobj"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µmodobj == µobj).ToObject()
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µmodname, "modname"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmain_module, "main_module"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µmain_module, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, µmodname, πTemp006); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label4:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label5
					}
					goto Label6
					// line 311: if modobj is obj and modname != main_module.__name__:
					πF.SetLineno(311)
				Label5:
					// line 312: return modname
					πF.SetLineno(312)
					if πE = πg.CheckLocal(πF, µmodname, "modname"); πE != nil {
						continue
					}
					πR = µmodname
					continue
					goto Label6
				Label6:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_lookup_module.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 309: """lookup name if module is imported"""
			πF.SetLineno(309)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp026}, πg.NewStr("lookup name if module is imported").ToObject()); πE != nil {
				continue
			}
			if πTemp027, πE = πg.ResolveGlobal(πF, ß_lookup_module); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp027, ß__doc__, πTemp026); πE != nil {
				continue
			}
			// line 314: def _stash_modules(main_module):
			πF.SetLineno(314)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "main_module", Def: nil}
			πTemp026 = πg.NewFunction(πg.NewCode("_stash_modules", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmain_module *πg.Object = πArgs[0]; _ = µmain_module
				var µmodmap *πg.Object = πg.UnboundLocal; _ = µmodmap
				var µimported *πg.Object = πg.UnboundLocal; _ = µimported
				var µoriginal *πg.Object = πg.UnboundLocal; _ = µoriginal
				var µitems *πg.Object = πg.UnboundLocal; _ = µitems
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
				var µsource_module *πg.Object = πg.UnboundLocal; _ = µsource_module
				var µtypes *πg.Object = πg.UnboundLocal; _ = µtypes
				var µnewmod *πg.Object = πg.UnboundLocal; _ = µnewmod
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
				_ = πTemp004
				var πTemp005 bool
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
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 315: modmap = _module_map()
					πF.SetLineno(315)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_module_map); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µmodmap = πTemp002
					// line 316: imported = []
					πF.SetLineno(316)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µimported = πTemp001
					// line 317: original = {}
					πF.SetLineno(317)
					πTemp004 = πg.NewDict()
					πTemp001 = πTemp004.ToObject()
					µoriginal = πTemp001
					// line 318: items = 'items' if PY3 else 'iteritems'
					πF.SetLineno(318)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label1
					}
					πTemp001 = ßitems.ToObject()
					goto Label2
				Label1:
					πTemp001 = ßiteritems.ToObject()
				Label2:
					µitems = πTemp001
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmain_module, "main_module"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmain_module, ß__dict__, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					πTemp003[1] = µitems
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
							continue
						}
						µname = πTemp006
						µobj = πTemp008
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 320: source_module = _lookup_module(modmap, name, obj, main_module)
					πF.SetLineno(320)
					πTemp003 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µmodmap, "modmap"); πE != nil {
						continue
					}
					πTemp003[0] = µmodmap
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp003[1] = µname
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[2] = µobj
					if πE = πg.CheckLocal(πF, µmain_module, "main_module"); πE != nil {
						continue
					}
					πTemp003[3] = µmain_module
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_lookup_module); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µsource_module = πTemp006
					if πE = πg.CheckLocal(πF, µsource_module, "source_module"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µsource_module); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					goto Label7
					// line 321: if source_module:
					πF.SetLineno(321)
				Label6:
					// line 322: imported.append((source_module, name))
					πF.SetLineno(322)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsource_module, "source_module"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µsource_module, µname).ToObject()
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µimported, "imported"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µimported, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label8
				Label7:
					// line 324: original[name] = obj
					πF.SetLineno(324)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µobj); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoriginal, "original"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp006 = µname
					if πE = πg.SetItem(πF, µoriginal, πTemp006, πTemp002); πE != nil {
						continue
					}
					goto Label8
				Label8:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µimported, "imported"); πE != nil {
						continue
					}
					πTemp003[0] = µimported
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label9
					}
					goto Label10
					// line 325: if len(imported):
					πF.SetLineno(325)
				Label9:
					// line 326: import types
					πF.SetLineno(326)
					if πTemp003, πE = πg.ImportModule(πF, "types"); πE != nil {
						continue
					}
					πTemp001 = πTemp003[0]
					µtypes = πTemp001
					// line 327: newmod = types.ModuleType(main_module.__name__)
					πF.SetLineno(327)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmain_module, "main_module"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmain_module, ß__name__, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µtypes, "types"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtypes, ßModuleType, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µnewmod = πTemp002
					// line 328: newmod.__dict__.update(original)
					πF.SetLineno(328)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoriginal, "original"); πE != nil {
						continue
					}
					πTemp003[0] = µoriginal
					if πE = πg.CheckLocal(πF, µnewmod, "newmod"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnewmod, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 329: newmod.__dill_imported = imported
					πF.SetLineno(329)
					if πE = πg.CheckLocal(πF, µimported, "imported"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µimported); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnewmod, "newmod"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µnewmod, ß__dill_imported, πTemp001); πE != nil {
						continue
					}
					// line 330: return newmod
					πF.SetLineno(330)
					if πE = πg.CheckLocal(πF, µnewmod, "newmod"); πE != nil {
						continue
					}
					πR = µnewmod
					continue
					goto Label11
				Label10:
					// line 332: return original
					πF.SetLineno(332)
					if πE = πg.CheckLocal(πF, µoriginal, "original"); πE != nil {
						continue
					}
					πR = µoriginal
					continue
					goto Label11
				Label11:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_stash_modules.ToObject(), πTemp026); πE != nil {
				continue
			}
			// line 334: def _restore_modules(main_module):
			πF.SetLineno(334)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "main_module", Def: nil}
			πTemp027 = πg.NewFunction(πg.NewCode("_restore_modules", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmain_module *πg.Object = πArgs[0]; _ = µmain_module
				var µimports *πg.Object = πg.UnboundLocal; _ = µimports
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µmain_module, "main_module"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmain_module, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Contains(πF, πTemp002, ß__dill_imported.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 335: if '__dill_imported' not in main_module.__dict__:
					πF.SetLineno(335)
				Label1:
					// line 336: return
					πF.SetLineno(336)
					πR = πg.None
					continue
					goto Label2
				Label2:
					// line 337: imports = main_module.__dict__.pop('__dill_imported')
					πF.SetLineno(337)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = ß__dill_imported.ToObject()
					if πE = πg.CheckLocal(πF, µmain_module, "main_module"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmain_module, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µimports = πTemp001
					if πE = πg.CheckLocal(πF, µimports, "imports"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µimports); πE != nil {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp003 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label5
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
							continue
						}
						µmodule = πTemp006
						µname = πTemp007
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 339: exec("from %s import %s" % (module, name), main_module.__dict__)
					πF.SetLineno(339)
					πE = πF.RaiseType(πg.NotImplementedErrorType, "exec is not available on Grumpy. Maybe never be.")
					continue
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_restore_modules.ToObject(), πTemp027); πE != nil {
				continue
			}
			// line 342: def dump_session(filename='/tmp/session.pkl', main=None, byref=False, **kwds):
			πF.SetLineno(342)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "filename", Def: πg.NewStr("/tmp/session.pkl").ToObject()}
			if πTemp029, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "main", Def: πTemp029}
			if πTemp029, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "byref", Def: πTemp029}
			πTemp028 = πg.NewFunction(πg.NewCode("dump_session", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
				var µmain *πg.Object = πArgs[1]; _ = µmain
				var µbyref *πg.Object = πArgs[2]; _ = µbyref
				var µkwds *πg.Object = πArgs[3]; _ = µkwds
				var µsettings *πg.Object = πg.UnboundLocal; _ = µsettings
				var µprotocol *πg.Object = πg.UnboundLocal; _ = µprotocol
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µpickler *πg.Object = πg.UnboundLocal; _ = µpickler
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 343: """pickle the current state of __main__ to a file"""
					πF.SetLineno(343)
					// line 344: from .settings import settings
					πF.SetLineno(344)
					if πTemp002, πE = πg.ImportModule(πF, ".settings"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßsettings); πE != nil {
						continue
					}
					µsettings = πTemp003
					// line 345: protocol = settings['protocol']
					πF.SetLineno(345)
					πTemp001 = ßprotocol.ToObject()
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsettings, πTemp001); πE != nil {
						continue
					}
					µprotocol = πTemp003
					if πE = πg.CheckLocal(πF, µmain, "main"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µmain == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 346: if main is None: main = _main_module
					πF.SetLineno(346)
				Label1:
					// line 346: if main is None: main = _main_module
					πF.SetLineno(346)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_main_module); πE != nil {
						continue
					}
					µmain = πTemp001
					goto Label2
				Label2:
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp002[0] = µfilename
					πTemp002[1] = ßwrite.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 347: if hasattr(filename, 'write'):
					πF.SetLineno(347)
				Label3:
					// line 348: f = filename
					πF.SetLineno(348)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					µf = µfilename
					goto Label5
				Label4:
					// line 350: f = open(filename, 'wb')
					πF.SetLineno(350)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp002[0] = µfilename
					πTemp002[1] = ßwb.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µf = πTemp003
					goto Label5
				Label5:
					// line 351: try:
					πF.SetLineno(351)
					πF.PushCheckpoint(6)
					if πE = πg.CheckLocal(πF, µbyref, "byref"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µbyref); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 352: if byref:
					πF.SetLineno(352)
				Label7:
					// line 353: main = _stash_modules(main)
					πF.SetLineno(353)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmain, "main"); πE != nil {
						continue
					}
					πTemp002[0] = µmain
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_stash_modules); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmain = πTemp003
					goto Label8
				Label8:
					// line 354: pickler = Pickler(f, protocol, **kwds)
					πF.SetLineno(354)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp002[0] = µf
					if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
						continue
					}
					πTemp002[1] = µprotocol
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp001, πTemp002, nil, nil, µkwds); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µpickler = πTemp003
					// line 355: pickler._main = main     #FIXME: dill.settings are disabled
					πF.SetLineno(355)
					if πE = πg.CheckLocal(πF, µmain, "main"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µmain); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_main, πTemp001); πE != nil {
						continue
					}
					// line 356: pickler._byref = False   # disable pickling by name reference
					πF.SetLineno(356)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_byref, πTemp003); πE != nil {
						continue
					}
					// line 357: pickler._recurse = False # disable pickling recursion for globals
					πF.SetLineno(357)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_recurse, πTemp003); πE != nil {
						continue
					}
					// line 358: pickler._session = True  # is best indicator of when pickling a session
					πF.SetLineno(358)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_session, πTemp003); πE != nil {
						continue
					}
					// line 359: pickler.dump(main)
					πF.SetLineno(359)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmain, "main"); πE != nil {
						continue
					}
					πTemp002[0] = µmain
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßdump, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πF.PopCheckpoint()
					goto Label6
				Label6:
					πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µf != µfilename).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 361: if f is not filename:  # If newly opened file
					πF.SetLineno(361)
				Label9:
					// line 362: f.close()
					πF.SetLineno(362)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label10
				Label10:
					if πTemp005 != nil {
						πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
						continue
					}
					if πR != nil {
						continue
					}
					// line 363: return
					πF.SetLineno(363)
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
			if πE = πF.Globals().SetItem(πF, ßdump_session.ToObject(), πTemp028); πE != nil {
				continue
			}
			// line 343: """pickle the current state of __main__ to a file"""
			πF.SetLineno(343)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp029}, πg.NewStr("pickle the current state of __main__ to a file").ToObject()); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßdump_session); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ß__doc__, πTemp029); πE != nil {
				continue
			}
			// line 365: def load_session(filename='/tmp/session.pkl', main=None, **kwds):
			πF.SetLineno(365)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "filename", Def: πg.NewStr("/tmp/session.pkl").ToObject()}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "main", Def: πTemp030}
			πTemp029 = πg.NewFunction(πg.NewCode("load_session", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
				var µmain *πg.Object = πArgs[1]; _ = µmain
				var µkwds *πg.Object = πArgs[2]; _ = µkwds
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µunpickler *πg.Object = πg.UnboundLocal; _ = µunpickler
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 366: """update the __main__ module with the state from the session file"""
					πF.SetLineno(366)
					if πE = πg.CheckLocal(πF, µmain, "main"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µmain == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 367: if main is None: main = _main_module
					πF.SetLineno(367)
				Label1:
					// line 367: if main is None: main = _main_module
					πF.SetLineno(367)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_main_module); πE != nil {
						continue
					}
					µmain = πTemp001
					goto Label2
				Label2:
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp004[0] = µfilename
					πTemp004[1] = ßread.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 368: if hasattr(filename, 'read'):
					πF.SetLineno(368)
				Label3:
					// line 369: f = filename
					πF.SetLineno(369)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					µf = µfilename
					goto Label5
				Label4:
					// line 371: f = open(filename, 'rb')
					πF.SetLineno(371)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp004[0] = µfilename
					πTemp004[1] = ßrb.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µf = πTemp002
					goto Label5
				Label5:
					// line 372: try: #FIXME: dill.settings are disabled
					πF.SetLineno(372)
					πF.PushCheckpoint(6)
					// line 373: unpickler = Unpickler(f, **kwds)
					πF.SetLineno(373)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp004[0] = µf
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßUnpickler); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, πTemp004, nil, nil, µkwds); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µunpickler = πTemp002
					// line 374: unpickler._main = main
					πF.SetLineno(374)
					if πE = πg.CheckLocal(πF, µmain, "main"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µmain); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µunpickler, "unpickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µunpickler, ß_main, πTemp001); πE != nil {
						continue
					}
					// line 375: unpickler._session = True
					πF.SetLineno(375)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µunpickler, "unpickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µunpickler, ß_session, πTemp002); πE != nil {
						continue
					}
					// line 376: module = unpickler.load()
					πF.SetLineno(376)
					if πE = πg.CheckLocal(πF, µunpickler, "unpickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µunpickler, ßload, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µmodule = πTemp002
					// line 377: unpickler._session = False
					πF.SetLineno(377)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µunpickler, "unpickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µunpickler, ß_session, πTemp002); πE != nil {
						continue
					}
					// line 378: main.__dict__.update(module.__dict__)
					πF.SetLineno(378)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmodule, ß__dict__, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmain, "main"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmain, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 379: _restore_modules(main)
					πF.SetLineno(379)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmain, "main"); πE != nil {
						continue
					}
					πTemp004[0] = µmain
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_restore_modules); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πF.PopCheckpoint()
					goto Label6
				Label6:
					πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µf != µfilename).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label7
					}
					goto Label8
					// line 381: if f is not filename:  # If newly opened file
					πF.SetLineno(381)
				Label7:
					// line 382: f.close()
					πF.SetLineno(382)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label8
				Label8:
					if πTemp005 != nil {
						πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
						continue
					}
					if πR != nil {
						continue
					}
					// line 383: return
					πF.SetLineno(383)
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
			if πE = πF.Globals().SetItem(πF, ßload_session.ToObject(), πTemp029); πE != nil {
				continue
			}
			// line 366: """update the __main__ module with the state from the session file"""
			πF.SetLineno(366)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp030}, πg.NewStr("update the __main__ module with the state from the session file").ToObject()); πE != nil {
				continue
			}
			if πTemp031, πE = πg.ResolveGlobal(πF, ßload_session); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp031, ß__doc__, πTemp030); πE != nil {
				continue
			}
			// line 387: class MetaCatchingDict(dict):
			πF.SetLineno(387)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp032, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			πTemp001[0] = πTemp032
			πTemp020 = πg.NewDict()
			if πTemp030, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ß__module__.ToObject(), πTemp030); πE != nil {
				continue
			}
			_, πE = πg.NewCode("MetaCatchingDict", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp020
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 388: def get(self, key, default=None):
					πF.SetLineno(388)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "default", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("get", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.BaseException
						_ = πTemp003
						var πTemp004 *πg.Traceback
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 389: try:
							πF.SetLineno(389)
							πF.PushCheckpoint(2)
							// line 390: return self[key]
							πF.SetLineno(390)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp003, πTemp004 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
							continue
							// line 391: except KeyError:
							πF.SetLineno(391)
						Label3:
							// line 392: return default
							πF.SetLineno(392)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πR = µdefault
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 394: def __missing__(self, key):
					πF.SetLineno(394)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__missing__", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 395: if issubclass(key, type):
							πF.SetLineno(395)
						Label1:
							// line 396: return save_type
							πF.SetLineno(396)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsave_type); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 398: raise KeyError()
							πF.SetLineno(398)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__missing__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp031, πE = πTemp020.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp031 == nil {
				πTemp031 = πg.TypeType.ToObject()
			}
			if πTemp032, πE = πTemp031.Call(πF, []*πg.Object{πg.NewStr("MetaCatchingDict").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp020.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMetaCatchingDict.ToObject(), πTemp032); πE != nil {
				continue
			}
			// line 402: class Pickler(StockPickler):
			πF.SetLineno(402)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp032, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
				continue
			}
			πTemp001[0] = πTemp032
			πTemp020 = πg.NewDict()
			if πTemp030, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ß__module__.ToObject(), πTemp030); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Pickler", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp020
				_ = πClass
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
					// line 403: """python's Pickler extended to interpreter sessions"""
					πF.SetLineno(403)
					// line 403: """python's Pickler extended to interpreter sessions"""
					πF.SetLineno(403)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("python's Pickler extended to interpreter sessions").ToObject()); πE != nil {
						continue
					}
					// line 404: dispatch = MetaCatchingDict(StockPickler.dispatch.copy())
					πF.SetLineno(404)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßStockPickler); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdispatch, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßMetaCatchingDict); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßdispatch.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 405: _session = False
					πF.SetLineno(405)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_session.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 406: from .settings import settings
					πF.SetLineno(406)
					if πTemp001, πE = πg.ImportModule(πF, ".settings"); πE != nil {
						continue
					}
					πTemp002 = πTemp001[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßsettings); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßsettings.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 408: def __init__(self, *args, **kwds):
					πF.SetLineno(408)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkwds *πg.Object = πArgs[2]; _ = µkwds
						var µsettings *πg.Object = πg.UnboundLocal; _ = µsettings
						var µ_byref *πg.Object = πg.UnboundLocal; _ = µ_byref
						var µ_fmode *πg.Object = πg.UnboundLocal; _ = µ_fmode
						var µ_recurse *πg.Object = πg.UnboundLocal; _ = µ_recurse
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
						_ = πTemp004
						var πTemp005 *πg.Object
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
							// line 409: settings = Pickler.settings
							πF.SetLineno(409)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							µsettings = πTemp002
							// line 410: _byref = kwds.pop('byref', None)
							πF.SetLineno(410)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßbyref.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µ_byref = πTemp002
							// line 412: _fmode = kwds.pop('fmode', None)
							πF.SetLineno(412)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßfmode.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µ_fmode = πTemp002
							// line 413: _recurse = kwds.pop('recurse', None)
							πF.SetLineno(413)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßrecurse.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µ_recurse = πTemp002
							// line 414: StockPickler.__init__(self, *args, **kwds)
							πF.SetLineno(414)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp002, πTemp003, µargs, nil, µkwds); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 415: self._main = _main_module
							πF.SetLineno(415)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_main_module); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_main, πTemp002); πE != nil {
								continue
							}
							// line 416: self._diff_cache = {}
							πF.SetLineno(416)
							πTemp004 = πg.NewDict()
							πTemp001 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_diff_cache, πTemp002); πE != nil {
								continue
							}
							// line 417: self._byref = settings['byref'] if _byref is None else _byref
							πF.SetLineno(417)
							if πE = πg.CheckLocal(πF, µ_byref, "_byref"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µ_byref == πTemp005).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label1
							}
							πTemp002 = ßbyref.ToObject()
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µsettings, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp005
							goto Label2
						Label1:
							if πE = πg.CheckLocal(πF, µ_byref, "_byref"); πE != nil {
								continue
							}
							πTemp001 = µ_byref
						Label2:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_byref, πTemp002); πE != nil {
								continue
							}
							// line 418: self._strictio = False #_strictio
							πF.SetLineno(418)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_strictio, πTemp002); πE != nil {
								continue
							}
							// line 419: self._fmode = settings['fmode'] if _fmode is None else _fmode
							πF.SetLineno(419)
							if πE = πg.CheckLocal(πF, µ_fmode, "_fmode"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µ_fmode == πTemp005).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label3
							}
							πTemp002 = ßfmode.ToObject()
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µsettings, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp005
							goto Label4
						Label3:
							if πE = πg.CheckLocal(πF, µ_fmode, "_fmode"); πE != nil {
								continue
							}
							πTemp001 = µ_fmode
						Label4:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_fmode, πTemp002); πE != nil {
								continue
							}
							// line 420: self._recurse = settings['recurse'] if _recurse is None else _recurse
							πF.SetLineno(420)
							if πE = πg.CheckLocal(πF, µ_recurse, "_recurse"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µ_recurse == πTemp005).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label5
							}
							πTemp002 = ßrecurse.ToObject()
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µsettings, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp005
							goto Label6
						Label5:
							if πE = πg.CheckLocal(πF, µ_recurse, "_recurse"); πE != nil {
								continue
							}
							πTemp001 = µ_recurse
						Label6:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_recurse, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 422: def dump(self, obj): #NOTE: if settings change, need to update attributes
					πF.SetLineno(422)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "obj", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("dump", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobj *πg.Object = πArgs[1]; _ = µobj
						var µsave_numpy_ufunc *πg.Object = πg.UnboundLocal; _ = µsave_numpy_ufunc
						var µsave_numpy_array *πg.Object = πg.UnboundLocal; _ = µsave_numpy_array
						var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
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
						var πTemp006 []πg.Param
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 423: stack.clear()  # clear record of 'recursion-sensitive' pickled objects
							πF.SetLineno(423)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstack); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclear, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNumpyUfuncType); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label1
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnumpyufunc); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001 = πTemp005
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 426: if NumpyUfuncType and numpyufunc(obj):
							πF.SetLineno(426)
						Label2:
							// line 428: def save_numpy_ufunc(pickler, obj):
							πF.SetLineno(428)
							πTemp006 = make([]πg.Param, 2)
							πTemp006[0] = πg.Param{Name: "pickler", Def: nil}
							πTemp006[1] = πg.Param{Name: "obj", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("save_numpy_ufunc", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µpickler *πg.Object = πArgs[0]; _ = µpickler
								var µobj *πg.Object = πArgs[1]; _ = µobj
								var µname *πg.Object = πg.UnboundLocal; _ = µname
								var πTemp001 []*πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
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
									// line 429: log.info("Nu: %s" % obj)
									πF.SetLineno(429)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.Mod(πF, πg.NewStr("Nu: %s").ToObject(), µobj); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 430: name = getattr(obj, '__qualname__', getattr(obj, '__name__', None))
									πF.SetLineno(430)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									πTemp001[0] = µobj
									πTemp001[1] = ß__qualname__.ToObject()
									πTemp004 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									πTemp004[0] = µobj
									πTemp004[1] = ß__name__.ToObject()
									if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp004[2] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									πTemp001[2] = πTemp003
									if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									µname = πTemp003
									// line 431: StockPickler.save_global(pickler, obj, name=name)
									πF.SetLineno(431)
									πTemp001 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
										continue
									}
									πTemp001[0] = µpickler
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									πTemp001[1] = µobj
									if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
										continue
									}
									πTemp005 = πg.KWArgs{
										{"name", µname},
									}
									if πTemp002, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsave_global, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp005); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 432: log.info("# Nu")
									πF.SetLineno(432)
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr("# Nu").ToObject()
									if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 433: return
									πF.SetLineno(433)
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
							µsave_numpy_ufunc = πTemp001
							// line 427: @register(type(obj))
							πF.SetLineno(427)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsave_numpy_ufunc, "save_numpy_ufunc"); πE != nil {
								continue
							}
							πTemp004[0] = µsave_numpy_ufunc
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp008[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[0] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µsave_numpy_ufunc = πTemp002
							goto Label3
						Label3:
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNumpyArrayType); πE != nil {
								continue
							}
							πTemp002 = πTemp005
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label4
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							if πTemp005, πE = πg.ResolveGlobal(πF, ßndarraysubclassinstance); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πTemp009
						Label4:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 440: if NumpyArrayType and ndarraysubclassinstance(obj):
							πF.SetLineno(440)
						Label5:
							// line 442: def save_numpy_array(pickler, obj):
							πF.SetLineno(442)
							πTemp006 = make([]πg.Param, 2)
							πTemp006[0] = πg.Param{Name: "pickler", Def: nil}
							πTemp006[1] = πg.Param{Name: "obj", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("save_numpy_array", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µpickler *πg.Object = πArgs[0]; _ = µpickler
								var µobj *πg.Object = πArgs[1]; _ = µobj
								var µnpdict *πg.Object = πg.UnboundLocal; _ = µnpdict
								var µf *πg.Object = πg.UnboundLocal; _ = µf
								var µargs *πg.Object = πg.UnboundLocal; _ = µargs
								var µstate *πg.Object = πg.UnboundLocal; _ = µstate
								var πTemp001 []*πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 *πg.Object
								_ = πTemp005
								var πTemp006 πg.KWArgs
								_ = πTemp006
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 443: log.info("Nu: (%s, %s)" % (obj.shape,obj.dtype))
									πF.SetLineno(443)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µobj, ßshape, nil); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GetAttr(πF, µobj, ßdtype, nil); πE != nil {
										continue
									}
									πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
									if πTemp002, πE = πg.Mod(πF, πg.NewStr("Nu: (%s, %s)").ToObject(), πTemp003); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 444: npdict = getattr(obj, '__dict__', None)
									πF.SetLineno(444)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									πTemp001[0] = µobj
									πTemp001[1] = ß__dict__.ToObject()
									if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001[2] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									µnpdict = πTemp003
									// line 445: f, args, state = obj.__reduce__()
									πF.SetLineno(445)
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µobj, ß__reduce__, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
										continue
									}
									µf = πTemp002
									µargs = πTemp004
									µstate = πTemp005
									// line 446: pickler.save_reduce(_create_array, (f,args,state,npdict), obj=obj)
									πF.SetLineno(446)
									πTemp001 = πF.MakeArgs(2)
									if πTemp002, πE = πg.ResolveGlobal(πF, ß_create_array); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µnpdict, "npdict"); πE != nil {
										continue
									}
									πTemp002 = πg.NewTuple4(µf, µargs, µstate, µnpdict).ToObject()
									πTemp001[1] = πTemp002
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									πTemp006 = πg.KWArgs{
										{"obj", µobj},
									}
									if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp006); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 447: log.info("# Nu")
									πF.SetLineno(447)
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr("# Nu").ToObject()
									if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 448: return
									πF.SetLineno(448)
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
							µsave_numpy_array = πTemp002
							// line 441: @register(type(obj))
							πF.SetLineno(441)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsave_numpy_array, "save_numpy_array"); πE != nil {
								continue
							}
							πTemp004[0] = µsave_numpy_array
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp008[0] = µobj
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[0] = πTemp009
							if πTemp005, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp005, πE = πTemp009.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µsave_numpy_array = πTemp005
							goto Label6
						Label6:
							if πTemp009, πE = πg.ResolveGlobal(πF, ßGENERATOR_FAIL); πE != nil {
								continue
							}
							πTemp005 = πTemp009
							if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label7
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							if πTemp010, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp010, πE = πg.ResolveGlobal(πF, ßGeneratorType); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Eq(πF, πTemp011, πTemp010); πE != nil {
								continue
							}
							πTemp005 = πTemp009
						Label7:
							if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 450: if GENERATOR_FAIL and type(obj) == GeneratorType:
							πF.SetLineno(450)
						Label8:
							// line 451: msg = "Can't pickle %s: attribute lookup builtins.generator failed" % GeneratorType
							πF.SetLineno(451)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneratorType); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Mod(πF, πg.NewStr("Can't pickle %s: attribute lookup builtins.generator failed").ToObject(), πTemp009); πE != nil {
								continue
							}
							µmsg = πTemp005
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp004[0] = µmsg
							if πTemp005, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 452: raise PicklingError(msg)
							πF.SetLineno(452)
							πE = πF.Raise(πTemp009, nil, nil)
							continue
							goto Label10
						Label9:
							// line 454: StockPickler.dump(self, obj)
							πF.SetLineno(454)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[1] = µobj
							if πTemp005, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp005, ßdump, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp009.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label10
						Label10:
							// line 455: stack.clear()  # clear record of 'recursion-sensitive' pickled objects
							πF.SetLineno(455)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßstack); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp005, ßclear, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp009.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 456: return
							πF.SetLineno(456)
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
					if πE = πClass.SetItem(πF, ßdump.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 457: dump.__doc__ = StockPickler.dump.__doc__
					πF.SetLineno(457)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßStockPickler); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßdump, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ß__doc__, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp005); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßdump); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 458: pass
					πF.SetLineno(458)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp031, πE = πTemp020.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp031 == nil {
				πTemp031 = πg.TypeType.ToObject()
			}
			if πTemp032, πE = πTemp031.Call(πF, []*πg.Object{πg.NewStr("Pickler").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp020.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPickler.ToObject(), πTemp032); πE != nil {
				continue
			}
			// line 460: class Unpickler(StockUnpickler):
			πF.SetLineno(460)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp032, πE = πg.ResolveGlobal(πF, ßStockUnpickler); πE != nil {
				continue
			}
			πTemp001[0] = πTemp032
			πTemp020 = πg.NewDict()
			if πTemp030, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ß__module__.ToObject(), πTemp030); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Unpickler", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp020
				_ = πClass
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 461: """python's Unpickler extended to interpreter sessions and more types"""
					πF.SetLineno(461)
					// line 461: """python's Unpickler extended to interpreter sessions and more types"""
					πF.SetLineno(461)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("python's Unpickler extended to interpreter sessions and more types").ToObject()); πE != nil {
						continue
					}
					// line 462: from .settings import settings
					πF.SetLineno(462)
					if πTemp002, πE = πg.ImportModule(πF, ".settings"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßsettings); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßsettings.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 463: _session = False
					πF.SetLineno(463)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_session.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 465: def find_class(self, module, name):
					πF.SetLineno(465)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "module", Def: nil}
					πTemp004[2] = πg.Param{Name: "name", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("find_class", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmodule *πg.Object = πArgs[1]; _ = µmodule
						var µname *πg.Object = πArgs[2]; _ = µname
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µmodule, µname).ToObject()
							πTemp003 = πg.NewTuple2(ß__builtin__.ToObject(), ß__main__.ToObject()).ToObject()
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µmodule, µname).ToObject()
							πTemp003 = πg.NewTuple2(ß__builtin__.ToObject(), ßNoneType.ToObject()).ToObject()
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 466: if (module, name) == ('__builtin__', '__main__'):
							πF.SetLineno(466)
						Label1:
							// line 467: return self._main.__dict__ #XXX: above set w/save_module_dict
							πF.SetLineno(467)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_main, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__dict__, nil); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
							// line 468: elif (module, name) == ('__builtin__', 'NoneType'):
							πF.SetLineno(468)
						Label2:
							// line 469: return type(None) #XXX: special case: NoneType missing
							πF.SetLineno(469)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πR = πTemp002
							continue
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µmodule, πg.NewStr("dill.dill").ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 470: if module == 'dill.dill': module = 'dill._dill'
							πF.SetLineno(470)
						Label4:
							// line 470: if module == 'dill.dill': module = 'dill._dill'
							πF.SetLineno(470)
							µmodule = πg.NewStr("dill._dill").ToObject()
							goto Label5
						Label5:
							// line 471: return StockUnpickler.find_class(self, module, name)
							πF.SetLineno(471)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp005[1] = µmodule
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005[2] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStockUnpickler); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfind_class, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ßfind_class.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 473: def __init__(self, *args, **kwds):
					πF.SetLineno(473)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__init__", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkwds *πg.Object = πArgs[2]; _ = µkwds
						var µsettings *πg.Object = πg.UnboundLocal; _ = µsettings
						var µ_ignore *πg.Object = πg.UnboundLocal; _ = µ_ignore
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 474: settings = Pickler.settings
							πF.SetLineno(474)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							µsettings = πTemp002
							// line 475: _ignore = kwds.pop('ignore', None)
							πF.SetLineno(475)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßignore.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µ_ignore = πTemp002
							// line 476: StockUnpickler.__init__(self, *args, **kwds)
							πF.SetLineno(476)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStockUnpickler); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp002, πTemp003, µargs, nil, µkwds); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 477: self._main = _main_module
							πF.SetLineno(477)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_main_module); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_main, πTemp002); πE != nil {
								continue
							}
							// line 478: self._ignore = settings['ignore'] if _ignore is None else _ignore
							πF.SetLineno(478)
							if πE = πg.CheckLocal(πF, µ_ignore, "_ignore"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µ_ignore == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label1
							}
							πTemp002 = ßignore.ToObject()
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µsettings, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							goto Label2
						Label1:
							if πE = πg.CheckLocal(πF, µ_ignore, "_ignore"); πE != nil {
								continue
							}
							πTemp001 = µ_ignore
						Label2:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_ignore, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 480: def load(self): #NOTE: if settings change, need to update attributes
					πF.SetLineno(480)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("load", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
						var πTemp001 []*πg.Object
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
						var πTemp007 []*πg.Object
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
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							// line 481: obj = StockUnpickler.load(self)
							πF.SetLineno(481)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStockUnpickler); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßload, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µobj = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__module__, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_main_module); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = ß__name__.ToObject()
							πTemp001[2] = ß__main__.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 482: if type(obj).__module__ == getattr(_main_module, '__name__', '__main__'):
							πF.SetLineno(482)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_ignore, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 483: if not self._ignore:
							πF.SetLineno(483)
						Label3:
							// line 485: try: obj.__class__ = getattr(self._main, type(obj).__name__)
							πF.SetLineno(485)
							πF.PushCheckpoint(6)
							// line 485: try: obj.__class__ = getattr(self._main, type(obj).__name__)
							πF.SetLineno(485)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_main, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp007[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µobj, ß__class__, πTemp002); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label5
						Label6:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							if πTemp006, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 486: except (AttributeError,TypeError): pass # defined in a file
							πF.SetLineno(486)
						Label7:
							// line 486: except (AttributeError,TypeError): pass # defined in a file
							πF.SetLineno(486)
							πF.RestoreExc(nil, nil)
							goto Label5
						Label5:
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 488: return obj
							πF.SetLineno(488)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πR = µobj
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßload.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 489: load.__doc__ = StockUnpickler.load.__doc__
					πF.SetLineno(489)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßStockUnpickler); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßload, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp007, ß__doc__, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp006); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßload); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 490: pass
					πF.SetLineno(490)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp031, πE = πTemp020.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp031 == nil {
				πTemp031 = πg.TypeType.ToObject()
			}
			if πTemp032, πE = πTemp031.Call(πF, []*πg.Object{πg.NewStr("Unpickler").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp020.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnpickler.ToObject(), πTemp032); πE != nil {
				continue
			}
			// line 492: '''
			πF.SetLineno(492)
			// line 498: pickle_dispatch_copy = StockPickler.dispatch.copy()
			πF.SetLineno(498)
			if πTemp030, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
				continue
			}
			if πTemp031, πE = πg.GetAttr(πF, πTemp030, ßdispatch, nil); πE != nil {
				continue
			}
			if πTemp030, πE = πg.GetAttr(πF, πTemp031, ßcopy, nil); πE != nil {
				continue
			}
			if πTemp031, πE = πTemp030.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpickle_dispatch_copy.ToObject(), πTemp031); πE != nil {
				continue
			}
			// line 500: def pickle(t, func):
			πF.SetLineno(500)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "t", Def: nil}
			πTemp004[1] = πg.Param{Name: "func", Def: nil}
			πTemp030 = πg.NewFunction(πg.NewCode("pickle", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µt *πg.Object = πArgs[0]; _ = µt
				var µfunc *πg.Object = πArgs[1]; _ = µfunc
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
					// line 501: """expose dispatch table for user-created extensions"""
					πF.SetLineno(501)
					// line 502: Pickler.dispatch[t] = func
					πF.SetLineno(502)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfunc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdispatch, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πTemp002 = µt
					if πE = πg.SetItem(πF, πTemp003, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 503: return
					πF.SetLineno(503)
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
			if πE = πF.Globals().SetItem(πF, ßpickle.ToObject(), πTemp030); πE != nil {
				continue
			}
			// line 501: """expose dispatch table for user-created extensions"""
			πF.SetLineno(501)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp031}, πg.NewStr("expose dispatch table for user-created extensions").ToObject()); πE != nil {
				continue
			}
			if πTemp032, πE = πg.ResolveGlobal(πF, ßpickle); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp032, ß__doc__, πTemp031); πE != nil {
				continue
			}
			// line 505: def register(t):
			πF.SetLineno(505)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "t", Def: nil}
			πTemp031 = πg.NewFunction(πg.NewCode("register", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µt *πg.Object = πArgs[0]; _ = µt
				var µproxy *πg.Object = πg.UnboundLocal; _ = µproxy
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 506: """register type to Pickler's dispatch table """
					πF.SetLineno(506)
					// line 507: def proxy(func):
					πF.SetLineno(507)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "func", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("proxy", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µfunc *πg.Object = πArgs[0]; _ = µfunc
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
							// line 508: Pickler.dispatch[t] = func
							πF.SetLineno(508)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfunc); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdispatch, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp002 = µt
							if πE = πg.SetItem(πF, πTemp003, πTemp002, πTemp001); πE != nil {
								continue
							}
							// line 509: return func
							πF.SetLineno(509)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							πR = µfunc
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					µproxy = πTemp001
					// line 510: return proxy
					πF.SetLineno(510)
					if πE = πg.CheckLocal(πF, µproxy, "proxy"); πE != nil {
						continue
					}
					πR = µproxy
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßregister.ToObject(), πTemp031); πE != nil {
				continue
			}
			// line 506: """register type to Pickler's dispatch table """
			πF.SetLineno(506)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp032}, πg.NewStr("register type to Pickler's dispatch table ").ToObject()); πE != nil {
				continue
			}
			if πTemp033, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp033, ß__doc__, πTemp032); πE != nil {
				continue
			}
			// line 512: def _revert_extension():
			πF.SetLineno(512)
			πTemp004 = make([]πg.Param, 0)
			πTemp032 = πg.NewFunction(πg.NewCode("_revert_extension", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtype *πg.Object = πg.UnboundLocal; _ = µtype
				var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
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
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 513: """drop dill-registered types from pickle's dispatch table"""
					πF.SetLineno(513)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdispatch, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßitems, nil); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µtype = πTemp004
						µfunc = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfunc, ß__module__, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					goto Label5
					// line 515: if func.__module__ == __name__:
					πF.SetLineno(515)
				Label4:
					// line 516: del StockPickler.dispatch[type]
					πF.SetLineno(516)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdispatch, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtype, "type"); πE != nil {
						continue
					}
					πTemp003 = µtype
					if πE = πg.DelItem(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtype, "type"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßpickle_dispatch_copy); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, πTemp004, µtype); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					goto Label7
					// line 517: if type in pickle_dispatch_copy:
					πF.SetLineno(517)
				Label6:
					// line 518: StockPickler.dispatch[type] = pickle_dispatch_copy[type]
					πF.SetLineno(518)
					if πE = πg.CheckLocal(πF, µtype, "type"); πE != nil {
						continue
					}
					πTemp003 = µtype
					if πTemp007, πE = πg.ResolveGlobal(πF, ßpickle_dispatch_copy); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßdispatch, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtype, "type"); πE != nil {
						continue
					}
					πTemp007 = µtype
					if πE = πg.SetItem(πF, πTemp008, πTemp007, πTemp003); πE != nil {
						continue
					}
					goto Label7
				Label7:
					goto Label5
				Label5:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_revert_extension.ToObject(), πTemp032); πE != nil {
				continue
			}
			// line 513: """drop dill-registered types from pickle's dispatch table"""
			πF.SetLineno(513)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp033}, πg.NewStr("drop dill-registered types from pickle's dispatch table").ToObject()); πE != nil {
				continue
			}
			if πTemp034, πE = πg.ResolveGlobal(πF, ß_revert_extension); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp034, ß__doc__, πTemp033); πE != nil {
				continue
			}
			// line 520: def use_diff(on=True):
			πF.SetLineno(520)
			πTemp004 = make([]πg.Param, 1)
			if πTemp034, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004[0] = πg.Param{Name: "on", Def: πTemp034}
			πTemp033 = πg.NewFunction(πg.NewCode("use_diff", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µon *πg.Object = πArgs[0]; _ = µon
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 521: """
					πF.SetLineno(521)
					// line 527: global _use_diff, diff
					πF.SetLineno(527)
					// line 528: _use_diff = on
					πF.SetLineno(528)
					if πE = πg.CheckLocal(πF, µon, "on"); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ß_use_diff.ToObject(), µon); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_use_diff); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßdiff); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp004 == πTemp005).ToObject()
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 529: if _use_diff and diff is None:
					πF.SetLineno(529)
				Label2:
					// line 530: try:
					πF.SetLineno(530)
					πF.PushCheckpoint(5)
					// line 531: from . import diff as d
					πF.SetLineno(531)
					if πTemp006, πE = πg.ImportModule(πF, ""); πE != nil {
						continue
					}
					πTemp001 = πTemp006[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßdiff); πE != nil {
						continue
					}
					µd = πTemp003
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					goto Label6
					// line 532: except:
					πF.SetLineno(532)
				Label6:
					// line 533: import diff as d
					πF.SetLineno(533)
					if πTemp006, πE = πg.ImportModule(πF, "diff"); πE != nil {
						continue
					}
					πTemp001 = πTemp006[0]
					µd = πTemp001
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					// line 534: diff = d
					πF.SetLineno(534)
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ßdiff.ToObject(), µd); πE != nil {
						continue
					}
					goto Label3
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßuse_diff.ToObject(), πTemp033); πE != nil {
				continue
			}
			// line 521: """
			πF.SetLineno(521)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp034}, πg.NewStr("\n    reduces size of pickles by only including object which have changed.\n    Decreases pickle size but increases CPU time needed.\n    Also helps avoid some unpicklable objects.\n    MUST be called at start of script, otherwise changes will not be recorded.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßuse_diff); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp035, ß__doc__, πTemp034); πE != nil {
				continue
			}
			// line 536: def _create_typemap():
			πF.SetLineno(536)
			πTemp004 = make([]πg.Param, 0)
			πTemp034 = πg.NewFunction(πg.NewCode("_create_typemap", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtypes *πg.Object = πg.UnboundLocal; _ = µtypes
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var µbuiltin *πg.Object = πg.UnboundLocal; _ = µbuiltin
				var µkey *πg.Object = πg.UnboundLocal; _ = µkey
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 10: goto Label10
						case 4: goto Label4
						case 5: goto Label5
						default: panic("unexpected function state")
						}
						// line 537: import types
						πF.SetLineno(537)
						if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
							continue
						}
						πTemp001 = πTemp002[0]
						µtypes = πTemp001
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
						// line 538: if PY3:
						πF.SetLineno(538)
					Label1:
						// line 539: d = dict(list(__builtin__.__dict__.items()) + \
						πF.SetLineno(539)
						πTemp002 = πF.MakeArgs(1)
						πTemp004 = πF.MakeArgs(1)
						if πTemp005, πE = πg.ResolveGlobal(πF, ß__builtin__); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß__dict__, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßitems, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp004[0] = πTemp006
						if πTemp005, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						πTemp004 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µtypes, "types"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µtypes, ß__dict__, nil); πE != nil {
							continue
						}
						if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßitems, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp007.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp004[0] = πTemp005
						if πTemp005, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						if πTemp001, πE = πg.Add(πF, πTemp006, πTemp007); πE != nil {
							continue
						}
						πTemp002[0] = πTemp001
						if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßitems, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp001.Call(πF, nil, nil); πE != nil {
							continue
						}
						µd = πTemp005
						// line 541: builtin = 'builtins'
						πF.SetLineno(541)
						µbuiltin = ßbuiltins.ToObject()
						goto Label3
					Label2:
						// line 543: d = types.__dict__.iteritems()
						πF.SetLineno(543)
						if πE = πg.CheckLocal(πF, µtypes, "types"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.GetAttr(πF, µtypes, ß__dict__, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßiteritems, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πTemp005.Call(πF, nil, nil); πE != nil {
							continue
						}
						µd = πTemp001
						// line 544: builtin = '__builtin__'
						πF.SetLineno(544)
						µbuiltin = ß__builtin__.ToObject()
						goto Label3
					Label3:
						if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, µd); πE != nil {
							continue
						}
						πF.PushCheckpoint(5)
						πTemp003 = false
					Label4:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp003 {
							πF.PopCheckpoint()
							goto Label6
						}
						if πTemp005, πE = πg.Next(πF, πTemp001); πE != nil {
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
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp005); πE != nil {
								continue
							}
							µkey = πTemp006
							µvalue = πTemp007
						}
						if πE != nil || !πTemp008 {
							continue
						}
						πF.PushCheckpoint(4)            
						πTemp002 = πF.MakeArgs(3)
						if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
							continue
						}
						πTemp002[0] = µvalue
						πTemp002[1] = ß__module__.ToObject()
						if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp002[2] = πTemp007
						if πTemp007, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						if πE = πg.CheckLocal(πF, µbuiltin, "builtin"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.Eq(πF, πTemp009, µbuiltin); πE != nil {
							continue
						}
						πTemp005 = πTemp006
						if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
							continue
						}
						if !πTemp008 {
							goto Label7
						}
						πTemp002 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
							continue
						}
						πTemp002[0] = µvalue
						if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
							continue
						}
						πTemp006 = πg.GetBool(πTemp009 == πTemp007).ToObject()
						πTemp005 = πTemp006
					Label7:
						if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
							continue
						}
						if πTemp008 {
							goto Label8
						}
						goto Label9
						// line 546: if getattr(value, '__module__', None) == builtin \
						πF.SetLineno(546)
					Label8:
						// line 548: yield key, value
						πF.SetLineno(548)
						if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
							continue
						}
						πTemp005 = πg.NewTuple2(µkey, µvalue).ToObject()
						πF.PushCheckpoint(10)
						return πTemp005, nil
					Label10:
						πTemp006 = πSent
						goto Label9
					Label9:
						continue
					Label5:
						if πE != nil || πR != nil {
							continue
						}
					Label6:
						// line 549: return
						πF.SetLineno(549)
						πR = πg.None
						continue
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_create_typemap.ToObject(), πTemp034); πE != nil {
				continue
			}
			// line 550: _reverse_typemap = dict(_create_typemap())
			πF.SetLineno(550)
			πTemp001 = πF.MakeArgs(1)
			if πTemp035, πE = πg.ResolveGlobal(πF, ß_create_typemap); πE != nil {
				continue
			}
			if πTemp036, πE = πTemp035.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp036
			if πTemp035, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			if πTemp036, πE = πTemp035.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ß_reverse_typemap.ToObject(), πTemp036); πE != nil {
				continue
			}
			// line 551: _reverse_typemap.update({
			πF.SetLineno(551)
			πTemp001 = πF.MakeArgs(1)
			πTemp020 = πg.NewDict()
			if πTemp035, πE = πg.ResolveGlobal(πF, ßCellType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßCellType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßMethodWrapperType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßMethodWrapperType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßPartialType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßPartialType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßSuperType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßSuperType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßItemGetterType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßItemGetterType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßAttrGetterType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßAttrGetterType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßFileType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßFileType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßBufferedRandomType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßBufferedRandomType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßBufferedReaderType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßBufferedReaderType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßBufferedWriterType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßBufferedWriterType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßTextWrapperType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßTextWrapperType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßPyBufferedRandomType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßPyBufferedRandomType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßPyBufferedReaderType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßPyBufferedReaderType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßPyBufferedWriterType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßPyBufferedWriterType.ToObject(), πTemp035); πE != nil {
				continue
			}
			if πTemp035, πE = πg.ResolveGlobal(πF, ßPyTextWrapperType); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ßPyTextWrapperType.ToObject(), πTemp035); πE != nil {
				continue
			}
			πTemp035 = πTemp020.ToObject()
			πTemp001[0] = πTemp035
			if πTemp035, πE = πg.ResolveGlobal(πF, ß_reverse_typemap); πE != nil {
				continue
			}
			if πTemp036, πE = πg.GetAttr(πF, πTemp035, ßupdate, nil); πE != nil {
				continue
			}
			if πTemp035, πE = πTemp036.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp035, πE = πg.ResolveGlobal(πF, ßExitType); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp035); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label75
			}
			goto Label76
			// line 568: if ExitType:
			πF.SetLineno(568)
		Label75:
			// line 569: _reverse_typemap['ExitType'] = ExitType
			πF.SetLineno(569)
			if πTemp035, πE = πg.ResolveGlobal(πF, ßExitType); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πTemp035); πE != nil {
				continue
			}
			if πTemp037, πE = πg.ResolveGlobal(πF, ß_reverse_typemap); πE != nil {
				continue
			}
			πTemp038 = ßExitType.ToObject()
			if πE = πg.SetItem(πF, πTemp037, πTemp038, πTemp036); πE != nil {
				continue
			}
			goto Label76
		Label76:
			if πTemp035, πE = πg.ResolveGlobal(πF, ßInputType); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp035); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label77
			}
			goto Label78
			// line 570: if InputType:
			πF.SetLineno(570)
		Label77:
			// line 571: _reverse_typemap['InputType'] = InputType
			πF.SetLineno(571)
			if πTemp035, πE = πg.ResolveGlobal(πF, ßInputType); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πTemp035); πE != nil {
				continue
			}
			if πTemp037, πE = πg.ResolveGlobal(πF, ß_reverse_typemap); πE != nil {
				continue
			}
			πTemp038 = ßInputType.ToObject()
			if πE = πg.SetItem(πF, πTemp037, πTemp038, πTemp036); πE != nil {
				continue
			}
			// line 572: _reverse_typemap['OutputType'] = OutputType
			πF.SetLineno(572)
			if πTemp035, πE = πg.ResolveGlobal(πF, ßOutputType); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πTemp035); πE != nil {
				continue
			}
			if πTemp037, πE = πg.ResolveGlobal(πF, ß_reverse_typemap); πE != nil {
				continue
			}
			πTemp038 = ßOutputType.ToObject()
			if πE = πg.SetItem(πF, πTemp037, πTemp038, πTemp036); πE != nil {
				continue
			}
			goto Label78
		Label78:
			if πTemp036, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp036); πE != nil {
				continue
			}
			πTemp035 = πg.GetBool(!πTemp007).ToObject()
			if πTemp007, πE = πg.IsTrue(πF, πTemp035); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label79
			}
			goto Label80
			// line 573: if not IS_PYPY:
			πF.SetLineno(573)
		Label79:
			// line 574: _reverse_typemap['WrapperDescriptorType'] = WrapperDescriptorType
			πF.SetLineno(574)
			if πTemp035, πE = πg.ResolveGlobal(πF, ßWrapperDescriptorType); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πTemp035); πE != nil {
				continue
			}
			if πTemp037, πE = πg.ResolveGlobal(πF, ß_reverse_typemap); πE != nil {
				continue
			}
			πTemp038 = ßWrapperDescriptorType.ToObject()
			if πE = πg.SetItem(πF, πTemp037, πTemp038, πTemp036); πE != nil {
				continue
			}
			// line 575: _reverse_typemap['MethodDescriptorType'] = MethodDescriptorType
			πF.SetLineno(575)
			if πTemp035, πE = πg.ResolveGlobal(πF, ßMethodDescriptorType); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πTemp035); πE != nil {
				continue
			}
			if πTemp037, πE = πg.ResolveGlobal(πF, ß_reverse_typemap); πE != nil {
				continue
			}
			πTemp038 = ßMethodDescriptorType.ToObject()
			if πE = πg.SetItem(πF, πTemp037, πTemp038, πTemp036); πE != nil {
				continue
			}
			// line 576: _reverse_typemap['ClassMethodDescriptorType'] = ClassMethodDescriptorType
			πF.SetLineno(576)
			if πTemp035, πE = πg.ResolveGlobal(πF, ßClassMethodDescriptorType); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πTemp035); πE != nil {
				continue
			}
			if πTemp037, πE = πg.ResolveGlobal(πF, ß_reverse_typemap); πE != nil {
				continue
			}
			πTemp038 = ßClassMethodDescriptorType.ToObject()
			if πE = πg.SetItem(πF, πTemp037, πTemp038, πTemp036); πE != nil {
				continue
			}
			goto Label81
		Label80:
			// line 578: _reverse_typemap['MemberDescriptorType'] = MemberDescriptorType
			πF.SetLineno(578)
			if πTemp035, πE = πg.ResolveGlobal(πF, ßMemberDescriptorType); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πTemp035); πE != nil {
				continue
			}
			if πTemp037, πE = πg.ResolveGlobal(πF, ß_reverse_typemap); πE != nil {
				continue
			}
			πTemp038 = ßMemberDescriptorType.ToObject()
			if πE = πg.SetItem(πF, πTemp037, πTemp038, πTemp036); πE != nil {
				continue
			}
			goto Label81
		Label81:
			if πTemp035, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp035); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label82
			}
			goto Label83
			// line 579: if PY3:
			πF.SetLineno(579)
		Label82:
			// line 580: _typemap = dict((v, k) for k, v in _reverse_typemap.items())
			πF.SetLineno(580)
			πTemp001 = πF.MakeArgs(1)
			πTemp004 = make([]πg.Param, 0)
			πTemp035 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µk *πg.Object = πg.UnboundLocal; _ = µk
				var µv *πg.Object = πg.UnboundLocal; _ = µv
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
						if πTemp002, πE = πg.ResolveGlobal(πF, ß_reverse_typemap); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitems, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
							µk = πTemp003
							µv = πTemp006
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 580: _typemap = dict((v, k) for k, v in _reverse_typemap.items())
						πF.SetLineno(580)
						if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
							continue
						}
						πTemp002 = πg.NewTuple2(µv, µk).ToObject()
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
			if πTemp036, πE = πTemp035.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp036
			if πTemp036, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			if πTemp037, πE = πTemp036.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ß_typemap.ToObject(), πTemp037); πE != nil {
				continue
			}
			goto Label84
		Label83:
			// line 582: _typemap = dict((v, k) for k, v in _reverse_typemap.iteritems())
			πF.SetLineno(582)
			πTemp001 = πF.MakeArgs(1)
			πTemp004 = make([]πg.Param, 0)
			πTemp036 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µk *πg.Object = πg.UnboundLocal; _ = µk
				var µv *πg.Object = πg.UnboundLocal; _ = µv
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
						if πTemp002, πE = πg.ResolveGlobal(πF, ß_reverse_typemap); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßiteritems, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
							µk = πTemp003
							µv = πTemp006
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 582: _typemap = dict((v, k) for k, v in _reverse_typemap.iteritems())
						πF.SetLineno(582)
						if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
							continue
						}
						πTemp002 = πg.NewTuple2(µv, µk).ToObject()
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
			if πTemp037, πE = πTemp036.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp037
			if πTemp037, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			if πTemp038, πE = πTemp037.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ß_typemap.ToObject(), πTemp038); πE != nil {
				continue
			}
			goto Label84
		Label84:
			// line 584: def _unmarshal(string):
			πF.SetLineno(584)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "string", Def: nil}
			πTemp037 = πg.NewFunction(πg.NewCode("_unmarshal", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstring *πg.Object = πArgs[0]; _ = µstring
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
					// line 585: return marshal.loads(string)
					πF.SetLineno(585)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					πTemp001[0] = µstring
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmarshal); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßloads, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
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
			if πE = πF.Globals().SetItem(πF, ß_unmarshal.ToObject(), πTemp037); πE != nil {
				continue
			}
			// line 587: def _load_type(name):
			πF.SetLineno(587)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "name", Def: nil}
			πTemp038 = πg.NewFunction(πg.NewCode("_load_type", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
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
					// line 588: return _reverse_typemap[name]
					πF.SetLineno(588)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001 = µname
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_reverse_typemap); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_load_type.ToObject(), πTemp038); πE != nil {
				continue
			}
			// line 590: def _create_type(typeobj, *args):
			πF.SetLineno(590)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "typeobj", Def: nil}
			πTemp039 = πg.NewFunction(πg.NewCode("_create_type", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtypeobj *πg.Object = πArgs[0]; _ = µtypeobj
				var µargs *πg.Object = πArgs[1]; _ = µargs
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 591: return typeobj(*args)
					πF.SetLineno(591)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtypeobj, "typeobj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µtypeobj, nil, µargs, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_create_type.ToObject(), πTemp039); πE != nil {
				continue
			}
			// line 593: def _create_function(fcode, fglobals, fname=None, fdefaults=None,
			πF.SetLineno(593)
			πTemp004 = make([]πg.Param, 7)
			πTemp004[0] = πg.Param{Name: "fcode", Def: nil}
			πTemp004[1] = πg.Param{Name: "fglobals", Def: nil}
			if πTemp041, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "fname", Def: πTemp041}
			if πTemp041, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "fdefaults", Def: πTemp041}
			if πTemp041, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[4] = πg.Param{Name: "fclosure", Def: πTemp041}
			if πTemp041, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[5] = πg.Param{Name: "fdict", Def: πTemp041}
			if πTemp041, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[6] = πg.Param{Name: "fkwdefaults", Def: πTemp041}
			πTemp040 = πg.NewFunction(πg.NewCode("_create_function", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfcode *πg.Object = πArgs[0]; _ = µfcode
				var µfglobals *πg.Object = πArgs[1]; _ = µfglobals
				var µfname *πg.Object = πArgs[2]; _ = µfname
				var µfdefaults *πg.Object = πArgs[3]; _ = µfdefaults
				var µfclosure *πg.Object = πArgs[4]; _ = µfclosure
				var µfdict *πg.Object = πArgs[5]; _ = µfdict
				var µfkwdefaults *πg.Object = πArgs[6]; _ = µfkwdefaults
				var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µfdict, "fdict"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µfdict == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 597: if fdict is None: fdict = dict()
					πF.SetLineno(597)
				Label1:
					// line 597: if fdict is None: fdict = dict()
					πF.SetLineno(597)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfdict = πTemp002
					goto Label2
				Label2:
					// line 598: func = FunctionType(fcode, fglobals or dict(), fname, fdefaults, fclosure)
					πF.SetLineno(598)
					πTemp004 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µfcode, "fcode"); πE != nil {
						continue
					}
					πTemp004[0] = µfcode
					if πE = πg.CheckLocal(πF, µfglobals, "fglobals"); πE != nil {
						continue
					}
					πTemp001 = µfglobals
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp005
				Label3:
					πTemp004[1] = πTemp001
					if πE = πg.CheckLocal(πF, µfname, "fname"); πE != nil {
						continue
					}
					πTemp004[2] = µfname
					if πE = πg.CheckLocal(πF, µfdefaults, "fdefaults"); πE != nil {
						continue
					}
					πTemp004[3] = µfdefaults
					if πE = πg.CheckLocal(πF, µfclosure, "fclosure"); πE != nil {
						continue
					}
					πTemp004[4] = µfclosure
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFunctionType); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µfunc = πTemp002
					// line 599: func.__dict__.update(fdict) #XXX: better copy? option to copy?
					πF.SetLineno(599)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfdict, "fdict"); πE != nil {
						continue
					}
					πTemp004[0] = µfdict
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfunc, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.CheckLocal(πF, µfkwdefaults, "fkwdefaults"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µfkwdefaults != πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					goto Label5
					// line 600: if fkwdefaults is not None:
					πF.SetLineno(600)
				Label4:
					// line 601: func.__kwdefaults__ = fkwdefaults
					πF.SetLineno(601)
					if πE = πg.CheckLocal(πF, µfkwdefaults, "fkwdefaults"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfkwdefaults); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µfunc, ß__kwdefaults__, πTemp001); πE != nil {
						continue
					}
					goto Label5
				Label5:
					// line 602: return func
					πF.SetLineno(602)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πR = µfunc
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_create_function.ToObject(), πTemp040); πE != nil {
				continue
			}
			// line 604: def _create_code(*args):
			πF.SetLineno(604)
			πTemp004 = make([]πg.Param, 0)
			πTemp041 = πg.NewFunction(πg.NewCode("_create_code", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp004 = πF.MakeArgs(2)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µargs, πTemp003); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					πTemp004[1] = ßencode.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001 = πTemp005
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 605: if PY3 and hasattr(args[-3], 'encode'): #FIXME: from PY2 fails (optcode)
					πF.SetLineno(605)
				Label2:
					// line 606: args = list(args)
					πF.SetLineno(606)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp004[0] = µargs
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µargs = πTemp003
					// line 607: args[-3] = args[-3].encode() # co_lnotab
					πF.SetLineno(607)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßencode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πE = πg.SetItem(πF, µargs, πTemp005, πTemp001); πE != nil {
						continue
					}
					// line 608: args[-10] = args[-10].encode() # co_code
					πF.SetLineno(608)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(10).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßencode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(10).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πE = πg.SetItem(πF, µargs, πTemp005, πTemp001); πE != nil {
						continue
					}
					goto Label3
				Label3:
					πTemp004 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					πTemp004[1] = ßco_posonlyargcount.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
						goto Label4
					}
					πTemp004 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					πTemp004[1] = ßco_kwonlyargcount.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
						goto Label5
					}
					goto Label6
					// line 609: if hasattr(CodeType, 'co_posonlyargcount'):
					πF.SetLineno(609)
				Label4:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp004[0] = µargs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp004[0] = µargs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					goto Label9
					// line 610: if len(args) == 16: return CodeType(*args)
					πF.SetLineno(610)
				Label7:
					// line 610: if len(args) == 16: return CodeType(*args)
					πF.SetLineno(610)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label9
					// line 611: elif len(args) == 15: return CodeType(args[0], 0, *args[1:])
					πF.SetLineno(611)
				Label8:
					// line 611: elif len(args) == 15: return CodeType(args[0], 0, *args[1:])
					πF.SetLineno(611)
					πTemp004 = πF.MakeArgs(2)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					πTemp004[1] = πg.NewInt(0).ToObject()
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Invoke(πF, πTemp001, πTemp004, πTemp003, nil, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp005
					continue
					goto Label9
				Label9:
					// line 612: return CodeType(args[0], 0, 0, *args[1:])
					πF.SetLineno(612)
					πTemp004 = πF.MakeArgs(3)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					πTemp004[1] = πg.NewInt(0).ToObject()
					πTemp004[2] = πg.NewInt(0).ToObject()
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Invoke(πF, πTemp001, πTemp004, πTemp003, nil, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp005
					continue
					goto Label6
					// line 613: elif hasattr(CodeType, 'co_kwonlyargcount'):
					πF.SetLineno(613)
				Label5:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp004[0] = µargs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label10
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp004[0] = µargs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label11
					}
					goto Label12
					// line 614: if len(args) == 16: return CodeType(args[0], *args[2:])
					πF.SetLineno(614)
				Label10:
					// line 614: if len(args) == 16: return CodeType(args[0], *args[2:])
					πF.SetLineno(614)
					πTemp004 = πF.MakeArgs(1)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Invoke(πF, πTemp001, πTemp004, πTemp003, nil, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp005
					continue
					goto Label12
					// line 615: elif len(args) == 15: return CodeType(*args)
					πF.SetLineno(615)
				Label11:
					// line 615: elif len(args) == 15: return CodeType(*args)
					πF.SetLineno(615)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label12
				Label12:
					// line 616: return CodeType(args[0], 0, *args[1:])
					πF.SetLineno(616)
					πTemp004 = πF.MakeArgs(2)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					πTemp004[1] = πg.NewInt(0).ToObject()
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Invoke(πF, πTemp001, πTemp004, πTemp003, nil, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp005
					continue
					goto Label6
				Label6:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp004[0] = µargs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label13
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp004[0] = µargs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label14
					}
					goto Label15
					// line 617: if len(args) == 16: return CodeType(args[0], *args[3:])
					πF.SetLineno(617)
				Label13:
					// line 617: if len(args) == 16: return CodeType(args[0], *args[3:])
					πF.SetLineno(617)
					πTemp004 = πF.MakeArgs(1)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Invoke(πF, πTemp001, πTemp004, πTemp003, nil, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp005
					continue
					goto Label15
					// line 618: elif len(args) == 15: return CodeType(args[0], *args[2:])
					πF.SetLineno(618)
				Label14:
					// line 618: elif len(args) == 15: return CodeType(args[0], *args[2:])
					πF.SetLineno(618)
					πTemp004 = πF.MakeArgs(1)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Invoke(πF, πTemp001, πTemp004, πTemp003, nil, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp005
					continue
					goto Label15
				Label15:
					// line 619: return CodeType(*args)
					πF.SetLineno(619)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_create_code.ToObject(), πTemp041); πE != nil {
				continue
			}
			// line 621: def _create_ftype(ftypeobj, func, args, kwds):
			πF.SetLineno(621)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "ftypeobj", Def: nil}
			πTemp004[1] = πg.Param{Name: "func", Def: nil}
			πTemp004[2] = πg.Param{Name: "args", Def: nil}
			πTemp004[3] = πg.Param{Name: "kwds", Def: nil}
			πTemp042 = πg.NewFunction(πg.NewCode("_create_ftype", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µftypeobj *πg.Object = πArgs[0]; _ = µftypeobj
				var µfunc *πg.Object = πArgs[1]; _ = µfunc
				var µargs *πg.Object = πArgs[2]; _ = µargs
				var µkwds *πg.Object = πArgs[3]; _ = µkwds
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Dict
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
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µkwds == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 622: if kwds is None:
					πF.SetLineno(622)
				Label1:
					// line 623: kwds = {}
					πF.SetLineno(623)
					πTemp004 = πg.NewDict()
					πTemp001 = πTemp004.ToObject()
					µkwds = πTemp001
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µargs == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 624: if args is None:
					πF.SetLineno(624)
				Label3:
					// line 625: args = ()
					πF.SetLineno(625)
					πTemp001 = πg.NewTuple0().ToObject()
					µargs = πTemp001
					goto Label4
				Label4:
					// line 626: return ftypeobj(func, *args, **kwds)
					πF.SetLineno(626)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp005[0] = µfunc
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µftypeobj, "ftypeobj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µftypeobj, πTemp005, µargs, nil, µkwds); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
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
			if πE = πF.Globals().SetItem(πF, ß_create_ftype.ToObject(), πTemp042); πE != nil {
				continue
			}
			// line 628: def _create_lock(locked, *args): #XXX: ignores 'blocking'
			πF.SetLineno(628)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "locked", Def: nil}
			πTemp043 = πg.NewFunction(πg.NewCode("_create_lock", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µlocked *πg.Object = πArgs[0]; _ = µlocked
				var µargs *πg.Object = πArgs[1]; _ = µargs
				var µLock *πg.Object = πg.UnboundLocal; _ = µLock
				var µlock *πg.Object = πg.UnboundLocal; _ = µlock
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 629: from threading import Lock
					πF.SetLineno(629)
					if πTemp002, πE = πg.ImportModule(πF, "threading"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßLock); πE != nil {
						continue
					}
					µLock = πTemp003
					// line 630: lock = Lock()
					πF.SetLineno(630)
					if πE = πg.CheckLocal(πF, µLock, "Lock"); πE != nil {
						continue
					}
					if πTemp001, πE = µLock.Call(πF, nil, nil); πE != nil {
						continue
					}
					µlock = πTemp001
					if πE = πg.CheckLocal(πF, µlocked, "locked"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µlocked); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 631: if locked:
					πF.SetLineno(631)
				Label1:
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 632: if not lock.acquire(False):
					πF.SetLineno(632)
				Label3:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("Cannot acquire lock").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßUnpicklingError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 633: raise UnpicklingError("Cannot acquire lock")
					πF.SetLineno(633)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label4
				Label4:
					goto Label2
				Label2:
					// line 634: return lock
					πF.SetLineno(634)
					if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
						continue
					}
					πR = µlock
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_create_lock.ToObject(), πTemp043); πE != nil {
				continue
			}
			// line 636: def _create_rlock(count, owner, *args): #XXX: ignores 'blocking'
			πF.SetLineno(636)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "count", Def: nil}
			πTemp004[1] = πg.Param{Name: "owner", Def: nil}
			πTemp044 = πg.NewFunction(πg.NewCode("_create_rlock", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcount *πg.Object = πArgs[0]; _ = µcount
				var µowner *πg.Object = πArgs[1]; _ = µowner
				var µargs *πg.Object = πArgs[2]; _ = µargs
				var µlock *πg.Object = πg.UnboundLocal; _ = µlock
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
					// line 637: lock = RLockType()
					πF.SetLineno(637)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßRLockType); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µlock = πTemp002
					if πE = πg.CheckLocal(πF, µowner, "owner"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µowner != πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 638: if owner is not None:
					πF.SetLineno(638)
				Label1:
					// line 639: lock._acquire_restore((count, owner))
					πF.SetLineno(639)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µowner, "owner"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µcount, µowner).ToObject()
					πTemp004[0] = πTemp001
					if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µlock, ß_acquire_restore, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µowner, "owner"); πE != nil {
						continue
					}
					πTemp001 = µowner
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µlock, ß_is_owned, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp002
				Label3:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					goto Label5
					// line 640: if owner and not lock._is_owned():
					πF.SetLineno(640)
				Label4:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("Cannot acquire lock").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßUnpicklingError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 641: raise UnpicklingError("Cannot acquire lock")
					πF.SetLineno(641)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label5
				Label5:
					// line 642: return lock
					πF.SetLineno(642)
					if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
						continue
					}
					πR = µlock
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_create_rlock.ToObject(), πTemp044); πE != nil {
				continue
			}
			// line 645: def _create_filehandle(name, mode, position, closed, open, strictio, fmode, fdata): # buffering=0
			πF.SetLineno(645)
			πTemp004 = make([]πg.Param, 8)
			πTemp004[0] = πg.Param{Name: "name", Def: nil}
			πTemp004[1] = πg.Param{Name: "mode", Def: nil}
			πTemp004[2] = πg.Param{Name: "position", Def: nil}
			πTemp004[3] = πg.Param{Name: "closed", Def: nil}
			πTemp004[4] = πg.Param{Name: "open", Def: nil}
			πTemp004[5] = πg.Param{Name: "strictio", Def: nil}
			πTemp004[6] = πg.Param{Name: "fmode", Def: nil}
			πTemp004[7] = πg.Param{Name: "fdata", Def: nil}
			πTemp045 = πg.NewFunction(πg.NewCode("_create_filehandle", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µmode *πg.Object = πArgs[1]; _ = µmode
				var µposition *πg.Object = πArgs[2]; _ = µposition
				var µclosed *πg.Object = πArgs[3]; _ = µclosed
				var µopen *πg.Object = πArgs[4]; _ = µopen
				var µstrictio *πg.Object = πArgs[5]; _ = µstrictio
				var µfmode *πg.Object = πArgs[6]; _ = µfmode
				var µfdata *πg.Object = πArgs[7]; _ = µfdata
				var µnames *πg.Object = πg.UnboundLocal; _ = µnames
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µtempfile *πg.Object = πg.UnboundLocal; _ = µtempfile
				var µexists *πg.Object = πg.UnboundLocal; _ = µexists
				var µcurrent_size *πg.Object = πg.UnboundLocal; _ = µcurrent_size
				var µflags *πg.Object = πg.UnboundLocal; _ = µflags
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µFILE *πg.Object = πg.UnboundLocal; _ = µFILE
				var µPyObject *πg.Object = πg.UnboundLocal; _ = µPyObject
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.BaseException
				_ = πTemp009
				var πTemp010 *πg.Traceback
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πTemp012 []*πg.Object
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 25: goto Label25
					case 10: goto Label10
					default: panic("unexpected function state")
					}
					// line 649: names = {'<stdin>':sys.__stdin__, '<stdout>':sys.__stdout__,
					πF.SetLineno(649)
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__stdin__, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("<stdin>").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__stdout__, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("<stdout>").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__stderr__, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("<stderr>").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					µnames = πTemp002
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µnames, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp006, πE = πg.Contains(πF, πTemp005, µname); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µname, πg.NewStr("<tmpfile>").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µname, πg.NewStr("<fdopen>").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 651: if name in list(names.keys()):
					πF.SetLineno(651)
				Label1:
					// line 652: f = names[name] #XXX: safer "f=sys.stdin"
					πF.SetLineno(652)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002 = µname
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µnames, πTemp002); πE != nil {
						continue
					}
					µf = πTemp003
					goto Label5
					// line 653: elif name == '<tmpfile>':
					πF.SetLineno(653)
				Label2:
					// line 654: f = os.tmpfile()
					πF.SetLineno(654)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtmpfile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					µf = πTemp002
					goto Label5
					// line 655: elif name == '<fdopen>':
					πF.SetLineno(655)
				Label3:
					// line 656: import tempfile
					πF.SetLineno(656)
					if πTemp004, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
						continue
					}
					πTemp002 = πTemp004[0]
					µtempfile = πTemp002
					// line 657: f = tempfile.TemporaryFile(mode)
					πF.SetLineno(657)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp004[0] = µmode
					if πE = πg.CheckLocal(πF, µtempfile, "tempfile"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µtempfile, ßTemporaryFile, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µf = πTemp003
					goto Label5
				Label4:
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, µmode, ßx.ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp007).ToObject()
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label6
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßhexversion, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, πTemp008, πg.NewInt(50528256).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label6:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label7
					}
					goto Label8
					// line 660: if "x" in mode and sys.hexversion < 0x03030000:
					πF.SetLineno(660)
				Label7:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("invalid mode: '%s'").ToObject(), µmode); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 661: raise ValueError("invalid mode: '%s'" % mode)
					πF.SetLineno(661)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label8
				Label8:
					// line 662: try:
					πF.SetLineno(662)
					πF.PushCheckpoint(10)
					// line 663: exists = os.path.exists(name)
					πF.SetLineno(663)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004[0] = µname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßexists, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µexists = πTemp003
					πF.PopCheckpoint()
					goto Label9
				Label10:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					goto Label11
					// line 664: except:
					πF.SetLineno(664)
				Label11:
					// line 665: exists = False
					πF.SetLineno(665)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µexists = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label9
				Label9:
					if πE = πg.CheckLocal(πF, µexists, "exists"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µexists); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label12
					}
					goto Label13
					// line 666: if not exists:
					πF.SetLineno(666)
				Label12:
					if πE = πg.CheckLocal(πF, µstrictio, "strictio"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µstrictio); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, µmode, ßr.ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp007).ToObject()
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µfmode, "fmode"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßFILE_FMODE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, µfmode, πTemp005); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label16:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label17
					}
					goto Label18
					// line 667: if strictio:
					πF.SetLineno(667)
				Label15:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("[Errno 2] No such file or directory: '%s'").ToObject(), µname); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFileNotFoundError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 668: raise FileNotFoundError("[Errno 2] No such file or directory: '%s'" % name)
					πF.SetLineno(668)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label18
					// line 669: elif "r" in mode and fmode != FILE_FMODE:
					πF.SetLineno(669)
				Label17:
					// line 670: name = '<fdopen>' # or os.devnull?
					πF.SetLineno(670)
					µname = πg.NewStr("<fdopen>").ToObject()
					goto Label18
				Label18:
					// line 671: current_size = 0 # or maintain position?
					πF.SetLineno(671)
					µcurrent_size = πg.NewInt(0).ToObject()
					goto Label14
				Label13:
					// line 673: current_size = os.path.getsize(name)
					πF.SetLineno(673)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004[0] = µname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßgetsize, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µcurrent_size = πTemp003
					goto Label14
				Label14:
					if πE = πg.CheckLocal(πF, µposition, "position"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcurrent_size, "current_size"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µposition, µcurrent_size); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label19
					}
					goto Label20
					// line 675: if position > current_size:
					πF.SetLineno(675)
				Label19:
					if πE = πg.CheckLocal(πF, µstrictio, "strictio"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µstrictio); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label21
					}
					if πE = πg.CheckLocal(πF, µfmode, "fmode"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßCONTENTS_FMODE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µfmode, πTemp003); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label22
					}
					goto Label23
					// line 676: if strictio:
					πF.SetLineno(676)
				Label21:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("invalid buffer size").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 677: raise ValueError("invalid buffer size")
					πF.SetLineno(677)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label23
					// line 678: elif fmode == CONTENTS_FMODE:
					πF.SetLineno(678)
				Label22:
					// line 679: position = current_size
					πF.SetLineno(679)
					if πE = πg.CheckLocal(πF, µcurrent_size, "current_size"); πE != nil {
						continue
					}
					µposition = µcurrent_size
					goto Label23
				Label23:
					goto Label20
				Label20:
					// line 682: try:
					πF.SetLineno(682)
					πF.PushCheckpoint(25)
					if πE = πg.CheckLocal(πF, µfmode, "fmode"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFILE_FMODE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µfmode, πTemp003); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label26
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µname, πg.NewStr("<fdopen>").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label27
					}
					if πE = πg.CheckLocal(πF, µfmode, "fmode"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßCONTENTS_FMODE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µfmode, πTemp005); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label28
					}
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Contains(πF, µmode, ßw.ToObject()); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp011).ToObject()
					πTemp003 = πTemp005
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label29
					}
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Contains(πF, µmode, ßx.ToObject()); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp011).ToObject()
					πTemp003 = πTemp005
				Label29:
					πTemp002 = πTemp003
				Label28:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label30
					}
					goto Label31
					// line 684: if fmode == FILE_FMODE:
					πF.SetLineno(684)
				Label26:
					// line 685: f = open(name, mode if "w" in mode else "w")
					πF.SetLineno(685)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004[0] = µname
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µmode, ßw.ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label33
					}
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp002 = µmode
					goto Label34
				Label33:
					πTemp002 = ßw.ToObject()
				Label34:
					πTemp004[1] = πTemp002
					if πE = πg.CheckLocal(πF, µopen, "open"); πE != nil {
						continue
					}
					if πTemp002, πE = µopen.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µf = πTemp002
					// line 686: f.write(fdata)
					πF.SetLineno(686)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfdata, "fdata"); πE != nil {
						continue
					}
					πTemp004[0] = µfdata
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µmode, ßw.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label35
					}
					goto Label36
					// line 687: if "w" not in mode:
					πF.SetLineno(687)
				Label35:
					// line 688: f.close()
					πF.SetLineno(688)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 689: f = open(name, mode)
					πF.SetLineno(689)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004[0] = µname
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp004[1] = µmode
					if πE = πg.CheckLocal(πF, µopen, "open"); πE != nil {
						continue
					}
					if πTemp002, πE = µopen.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µf = πTemp002
					goto Label36
				Label36:
					goto Label32
					// line 690: elif name == '<fdopen>': # file did not exist
					πF.SetLineno(690)
				Label27:
					// line 691: import tempfile
					πF.SetLineno(691)
					if πTemp004, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
						continue
					}
					πTemp002 = πTemp004[0]
					µtempfile = πTemp002
					// line 692: f = tempfile.TemporaryFile(mode)
					πF.SetLineno(692)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp004[0] = µmode
					if πE = πg.CheckLocal(πF, µtempfile, "tempfile"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µtempfile, ßTemporaryFile, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µf = πTemp003
					goto Label32
					// line 693: elif fmode == CONTENTS_FMODE \
					πF.SetLineno(693)
				Label30:
					// line 696: flags = os.O_CREAT
					πF.SetLineno(696)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßO_CREAT, nil); πE != nil {
						continue
					}
					µflags = πTemp003
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µmode, πg.NewStr("+").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label37
					}
					goto Label38
					// line 697: if "+" in mode:
					πF.SetLineno(697)
				Label37:
					// line 698: flags |= os.O_RDWR
					πF.SetLineno(698)
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßO_RDWR, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IOr(πF, µflags, πTemp003); πE != nil {
						continue
					}
					µflags = πTemp002
					goto Label39
				Label38:
					// line 700: flags |= os.O_WRONLY
					πF.SetLineno(700)
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßO_WRONLY, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IOr(πF, µflags, πTemp003); πE != nil {
						continue
					}
					µflags = πTemp002
					goto Label39
				Label39:
					// line 701: f = os.fdopen(os.open(name, flags), mode)
					πF.SetLineno(701)
					πTemp004 = πF.MakeArgs(2)
					πTemp012 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp012[0] = µname
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp012[1] = µflags
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßopen, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp004[0] = πTemp002
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp004[1] = µmode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfdopen, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µf = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label40
					}
					goto Label41
					// line 703: if PY3:
					πF.SetLineno(703)
				Label40:
					// line 704: r = getattr(f, "buffer", f)
					πF.SetLineno(704)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp004[0] = µf
					πTemp004[1] = ßbuffer.ToObject()
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp004[2] = µf
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µr = πTemp003
					// line 705: r = getattr(r, "raw", r)
					πF.SetLineno(705)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[0] = µr
					πTemp004[1] = ßraw.ToObject()
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[2] = µr
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µr = πTemp003
					// line 706: r.name = name
					πF.SetLineno(706)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µname); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µr, ßname, πTemp002); πE != nil {
						continue
					}
					goto Label42
				Label41:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßHAS_CTYPES); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label43
					}
					goto Label44
					// line 708: if not HAS_CTYPES:
					πF.SetLineno(708)
				Label43:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("No module named 'ctypes'").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 709: raise ImportError("No module named 'ctypes'")
					πF.SetLineno(709)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label44
				Label44:
					// line 710: class FILE(ctypes.Structure):
					πF.SetLineno(710)
					πTemp004 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßStructure, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp008
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("FILE", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp001
						_ = πClass
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 711: _fields_ = [("refcount", ctypes.c_long),
							πF.SetLineno(711)
							πTemp001 = make([]*πg.Object, 4)
							if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßctypes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßc_long, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ßrefcount.ToObject(), πTemp004).ToObject()
							πTemp001[0] = πTemp002
							if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßctypes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpy_object, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ßtype_obj.ToObject(), πTemp004).ToObject()
							πTemp001[1] = πTemp002
							if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßctypes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßc_voidp, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ßfile_pointer.ToObject(), πTemp004).ToObject()
							πTemp001[2] = πTemp002
							if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßctypes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpy_object, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ßname.ToObject(), πTemp004).ToObject()
							πTemp001[3] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πClass.SetItem(πF, ß_fields_.ToObject(), πTemp002); πE != nil {
								continue
							}
						}
						return nil, nil
					}).Eval(πF, πF.Globals(), nil, nil)
					if πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
						continue
					}
					if πTemp003 == nil {
						πTemp003 = πg.TypeType.ToObject()
					}
					if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("FILE").ToObject(), πg.NewTuple(πTemp004...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					µFILE = πTemp005
					// line 716: class PyObject(ctypes.Structure):
					πF.SetLineno(716)
					πTemp004 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßStructure, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp008
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("PyObject", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp001
						_ = πClass
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 717: _fields_ = [
							πF.SetLineno(717)
							πTemp001 = make([]*πg.Object, 2)
							if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßctypes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßc_int, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ßob_refcnt.ToObject(), πTemp004).ToObject()
							πTemp001[0] = πTemp002
							if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßctypes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpy_object, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ßob_type.ToObject(), πTemp004).ToObject()
							πTemp001[1] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πClass.SetItem(πF, ß_fields_.ToObject(), πTemp002); πE != nil {
								continue
							}
						}
						return nil, nil
					}).Eval(πF, πF.Globals(), nil, nil)
					if πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
						continue
					}
					if πTemp003 == nil {
						πTemp003 = πg.TypeType.ToObject()
					}
					if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("PyObject").ToObject(), πg.NewTuple(πTemp004...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					µPyObject = πTemp005
					// line 723: ctypes.cast(id(f), ctypes.POINTER(FILE)).contents.name = name
					πF.SetLineno(723)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µname); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(2)
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp012[0] = µf
					if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp004[0] = πTemp005
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µFILE, "FILE"); πE != nil {
						continue
					}
					πTemp012[0] = µFILE
					if πTemp003, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßPOINTER, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp005.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp004[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßcast, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßcontents, nil); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ßname, πTemp002); πE != nil {
						continue
					}
					// line 724: ctypes.cast(id(name), ctypes.POINTER(PyObject)).contents.ob_refcnt += 1
					πF.SetLineno(724)
					πTemp004 = πF.MakeArgs(2)
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp012[0] = µname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp004[0] = πTemp003
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µPyObject, "PyObject"); πE != nil {
						continue
					}
					πTemp012[0] = µPyObject
					if πTemp002, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßPOINTER, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp004[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcast, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcontents, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßob_refcnt, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(2)
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp012[0] = µname
					if πTemp005, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp004[0] = πTemp008
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µPyObject, "PyObject"); πE != nil {
						continue
					}
					πTemp012[0] = µPyObject
					if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßPOINTER, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp008.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp004[1] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßcast, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßcontents, nil); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ßob_refcnt, πTemp003); πE != nil {
						continue
					}
					goto Label42
				Label42:
					// line 725: assert f.name == name
					πF.SetLineno(725)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßname, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp003, µname); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					goto Label32
				Label31:
					// line 727: f = open(name, mode)
					πF.SetLineno(727)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004[0] = µname
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp004[1] = µmode
					if πE = πg.CheckLocal(πF, µopen, "open"); πE != nil {
						continue
					}
					if πTemp002, πE = µopen.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µf = πTemp002
					goto Label32
				Label32:
					πF.PopCheckpoint()
					goto Label24
				Label25:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßFileNotFoundError); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					if πTemp006, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label45
					}
					πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
					continue
					// line 728: except (IOError, FileNotFoundError):
					πF.SetLineno(728)
				Label45:
					// line 729: err = sys.exc_info()[1]
					πF.SetLineno(729)
					πTemp002 = πg.NewInt(1).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp008.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					µerr = πTemp003
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					πTemp004[0] = µerr
					if πTemp002, πE = πg.ResolveGlobal(πF, ßUnpicklingError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 730: raise UnpicklingError(err)
					πF.SetLineno(730)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label24
				Label24:
					goto Label5
				Label5:
					if πE = πg.CheckLocal(πF, µclosed, "closed"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µclosed); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label46
					}
					if πE = πg.CheckLocal(πF, µposition, "position"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GE(πF, µposition, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label47
					}
					if πE = πg.CheckLocal(πF, µfmode, "fmode"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßHANDLE_FMODE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, µfmode, πTemp005); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label47:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label48
					}
					goto Label49
					// line 731: if closed:
					πF.SetLineno(731)
				Label46:
					// line 732: f.close()
					πF.SetLineno(732)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label49
					// line 733: elif position >= 0 and fmode != HANDLE_FMODE:
					πF.SetLineno(733)
				Label48:
					// line 734: f.seek(position)
					πF.SetLineno(734)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µposition, "position"); πE != nil {
						continue
					}
					πTemp004[0] = µposition
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßseek, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label49
				Label49:
					// line 735: return f
					πF.SetLineno(735)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πR = µf
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_create_filehandle.ToObject(), πTemp045); πE != nil {
				continue
			}
			// line 737: def _create_stringi(value, position, closed):
			πF.SetLineno(737)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "value", Def: nil}
			πTemp004[1] = πg.Param{Name: "position", Def: nil}
			πTemp004[2] = πg.Param{Name: "closed", Def: nil}
			πTemp046 = πg.NewFunction(πg.NewCode("_create_stringi", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µvalue *πg.Object = πArgs[0]; _ = µvalue
				var µposition *πg.Object = πArgs[1]; _ = µposition
				var µclosed *πg.Object = πArgs[2]; _ = µclosed
				var µf *πg.Object = πg.UnboundLocal; _ = µf
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
					// line 738: f = StringIO(value)
					πF.SetLineno(738)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µf = πTemp003
					if πE = πg.CheckLocal(πF, µclosed, "closed"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µclosed); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 739: if closed: f.close()
					πF.SetLineno(739)
				Label1:
					// line 739: if closed: f.close()
					πF.SetLineno(739)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label3
				Label2:
					// line 740: else: f.seek(position)
					πF.SetLineno(740)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µposition, "position"); πE != nil {
						continue
					}
					πTemp001[0] = µposition
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßseek, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label3
				Label3:
					// line 741: return f
					πF.SetLineno(741)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πR = µf
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_create_stringi.ToObject(), πTemp046); πE != nil {
				continue
			}
			// line 743: def _create_stringo(value, position, closed):
			πF.SetLineno(743)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "value", Def: nil}
			πTemp004[1] = πg.Param{Name: "position", Def: nil}
			πTemp004[2] = πg.Param{Name: "closed", Def: nil}
			πTemp047 = πg.NewFunction(πg.NewCode("_create_stringo", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µvalue *πg.Object = πArgs[0]; _ = µvalue
				var µposition *πg.Object = πArgs[1]; _ = µposition
				var µclosed *πg.Object = πArgs[2]; _ = µclosed
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 744: f = StringIO()
					πF.SetLineno(744)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µf = πTemp002
					if πE = πg.CheckLocal(πF, µclosed, "closed"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µclosed); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 745: if closed: f.close()
					πF.SetLineno(745)
				Label1:
					// line 745: if closed: f.close()
					πF.SetLineno(745)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label3
				Label2:
					// line 747: f.write(value)
					πF.SetLineno(747)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp004[0] = µvalue
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 748: f.seek(position)
					πF.SetLineno(748)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µposition, "position"); πE != nil {
						continue
					}
					πTemp004[0] = µposition
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßseek, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label3
				Label3:
					// line 749: return f
					πF.SetLineno(749)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πR = µf
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_create_stringo.ToObject(), πTemp047); πE != nil {
				continue
			}
			// line 751: class _itemgetter_helper(object):
			πF.SetLineno(751)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp050, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp050
			πTemp020 = πg.NewDict()
			if πTemp048, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ß__module__.ToObject(), πTemp048); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_itemgetter_helper", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp020
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 752: def __init__(self):
					πF.SetLineno(752)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 753: self.items = []
							πF.SetLineno(753)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßitems, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 754: def __getitem__(self, item):
					πF.SetLineno(754)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "item", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__getitem__", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µitem *πg.Object = πArgs[1]; _ = µitem
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
							// line 755: self.items.append(item)
							πF.SetLineno(755)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 756: return
							πF.SetLineno(756)
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp049, πE = πTemp020.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp049 == nil {
				πTemp049 = πg.TypeType.ToObject()
			}
			if πTemp050, πE = πTemp049.Call(πF, []*πg.Object{πg.NewStr("_itemgetter_helper").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp020.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_itemgetter_helper.ToObject(), πTemp050); πE != nil {
				continue
			}
			// line 758: class _attrgetter_helper(object):
			πF.SetLineno(758)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp050, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp050
			πTemp020 = πg.NewDict()
			if πTemp048, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp020.SetItem(πF, ß__module__.ToObject(), πTemp048); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_attrgetter_helper", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp020
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 759: def __init__(self, attrs, index=None):
					πF.SetLineno(759)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "attrs", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "index", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µattrs *πg.Object = πArgs[1]; _ = µattrs
						var µindex *πg.Object = πArgs[2]; _ = µindex
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 760: self.attrs = attrs
							πF.SetLineno(760)
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µattrs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßattrs, πTemp001); πE != nil {
								continue
							}
							// line 761: self.index = index
							πF.SetLineno(761)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µindex); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßindex, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 762: def __getattribute__(self, attr):
					πF.SetLineno(762)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "attr", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__getattribute__", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µattr *πg.Object = πArgs[1]; _ = µattr
						var µattrs *πg.Object = πg.UnboundLocal; _ = µattrs
						var µindex *πg.Object = πg.UnboundLocal; _ = µindex
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
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
							// line 763: attrs = object.__getattribute__(self, "attrs")
							πF.SetLineno(763)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ßattrs.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__getattribute__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µattrs = πTemp002
							// line 764: index = object.__getattribute__(self, "index")
							πF.SetLineno(764)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ßindex.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__getattribute__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µindex = πTemp002
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µindex == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 765: if index is None:
							πF.SetLineno(765)
						Label1:
							// line 766: index = len(attrs)
							πF.SetLineno(766)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							πTemp001[0] = µattrs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µindex = πTemp003
							// line 767: attrs.append(attr)
							πF.SetLineno(767)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp001[0] = µattr
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µattrs, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label2:
							// line 769: attrs[index] = ".".join([attrs[index], attr])
							πF.SetLineno(769)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp002 = µindex
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µattrs, πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp005[1] = µattr
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(".").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp006 = µindex
							if πE = πg.SetItem(πF, µattrs, πTemp006, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 770: return type(self)(attrs, index)
							πF.SetLineno(770)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							πTemp001[0] = µattrs
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001[1] = µindex
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__getattribute__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp049, πE = πTemp020.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp049 == nil {
				πTemp049 = πg.TypeType.ToObject()
			}
			if πTemp050, πE = πTemp049.Call(πF, []*πg.Object{πg.NewStr("_attrgetter_helper").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp020.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_attrgetter_helper.ToObject(), πTemp050); πE != nil {
				continue
			}
			if πTemp048, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp048); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label85
			}
			goto Label86
			// line 772: if PY3:
			πF.SetLineno(772)
		Label85:
			// line 773: def _create_cell(contents):
			πF.SetLineno(773)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "contents", Def: nil}
			πTemp048 = πg.NewFunction(πg.NewCode("_create_cell", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcontents *πg.Object = πArgs[0]; _ = µcontents
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
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
					// line 774: return (lambda y: contents).__closure__[0]
					πF.SetLineno(774)
					πTemp001 = πg.NewInt(0).ToObject()
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "y", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µy *πg.Object = πArgs[0]; _ = µy
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 774: return (lambda y: contents).__closure__[0]
							πF.SetLineno(774)
							if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
								continue
							}
							πR = µcontents
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ß__closure__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_create_cell.ToObject(), πTemp048); πE != nil {
				continue
			}
			goto Label87
		Label86:
			// line 776: def _create_cell(contents):
			πF.SetLineno(776)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "contents", Def: nil}
			πTemp049 = πg.NewFunction(πg.NewCode("_create_cell", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcontents *πg.Object = πArgs[0]; _ = µcontents
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
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
					// line 777: return (lambda y: contents).func_closure[0]
					πF.SetLineno(777)
					πTemp001 = πg.NewInt(0).ToObject()
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "y", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µy *πg.Object = πArgs[0]; _ = µy
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 777: return (lambda y: contents).func_closure[0]
							πF.SetLineno(777)
							if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
								continue
							}
							πR = µcontents
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßfunc_closure, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_create_cell.ToObject(), πTemp049); πE != nil {
				continue
			}
			goto Label87
		Label87:
			// line 779: def _create_weakref(obj, *args):
			πF.SetLineno(779)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp050 = πg.NewFunction(πg.NewCode("_create_weakref", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µargs *πg.Object = πArgs[1]; _ = µargs
				var µref *πg.Object = πg.UnboundLocal; _ = µref
				var µUserDict *πg.Object = πg.UnboundLocal; _ = µUserDict
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
					// line 780: from weakref import ref
					πF.SetLineno(780)
					if πTemp002, πE = πg.ImportModule(πF, "weakref"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßref); πE != nil {
						continue
					}
					µref = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µobj == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 781: if obj is None: # it's dead
					πF.SetLineno(781)
				Label1:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 782: if PY3:
					πF.SetLineno(782)
				Label3:
					// line 783: from collections import UserDict
					πF.SetLineno(783)
					if πTemp002, πE = πg.ImportModule(πF, "collections"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßUserDict); πE != nil {
						continue
					}
					µUserDict = πTemp003
					goto Label5
				Label4:
					// line 785: from UserDict import UserDict
					πF.SetLineno(785)
					if πTemp002, πE = πg.ImportModule(πF, "UserDict"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßUserDict); πE != nil {
						continue
					}
					µUserDict = πTemp003
					goto Label5
				Label5:
					// line 786: return ref(UserDict(), *args)
					πF.SetLineno(786)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µUserDict, "UserDict"); πE != nil {
						continue
					}
					if πTemp001, πE = µUserDict.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µref, πTemp002, µargs, nil, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 787: return ref(obj, *args)
					πF.SetLineno(787)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µref, πTemp002, µargs, nil, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
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
			if πE = πF.Globals().SetItem(πF, ß_create_weakref.ToObject(), πTemp050); πE != nil {
				continue
			}
			// line 789: def _create_weakproxy(obj, callable=False, *args):
			πF.SetLineno(789)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp052, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "callable", Def: πTemp052}
			πTemp051 = πg.NewFunction(πg.NewCode("_create_weakproxy", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µcallable *πg.Object = πArgs[1]; _ = µcallable
				var µargs *πg.Object = πArgs[2]; _ = µargs
				var µproxy *πg.Object = πg.UnboundLocal; _ = µproxy
				var µUserDict *πg.Object = πg.UnboundLocal; _ = µUserDict
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 790: from weakref import proxy
					πF.SetLineno(790)
					if πTemp002, πE = πg.ImportModule(πF, "weakref"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßproxy); πE != nil {
						continue
					}
					µproxy = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µobj == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 791: if obj is None: # it's dead
					πF.SetLineno(791)
				Label1:
					if πE = πg.CheckLocal(πF, µcallable, "callable"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µcallable); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 792: if callable: return proxy(lambda x:x, *args)
					πF.SetLineno(792)
				Label3:
					// line 792: if callable: return proxy(lambda x:x, *args)
					πF.SetLineno(792)
					πTemp002 = πF.MakeArgs(1)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 792: if callable: return proxy(lambda x:x, *args)
							πF.SetLineno(792)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πR = µx
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µproxy, "proxy"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µproxy, πTemp002, µargs, nil, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp001
					continue
					goto Label4
				Label4:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 793: if PY3:
					πF.SetLineno(793)
				Label5:
					// line 794: from collections import UserDict
					πF.SetLineno(794)
					if πTemp002, πE = πg.ImportModule(πF, "collections"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßUserDict); πE != nil {
						continue
					}
					µUserDict = πTemp003
					goto Label7
				Label6:
					// line 796: from UserDict import UserDict
					πF.SetLineno(796)
					if πTemp002, πE = πg.ImportModule(πF, "UserDict"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßUserDict); πE != nil {
						continue
					}
					µUserDict = πTemp003
					goto Label7
				Label7:
					// line 797: return proxy(UserDict(), *args)
					πF.SetLineno(797)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µUserDict, "UserDict"); πE != nil {
						continue
					}
					if πTemp001, πE = µUserDict.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µproxy, "proxy"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µproxy, πTemp002, µargs, nil, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 798: return proxy(obj, *args)
					πF.SetLineno(798)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µproxy, "proxy"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µproxy, πTemp002, µargs, nil, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
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
			if πE = πF.Globals().SetItem(πF, ß_create_weakproxy.ToObject(), πTemp051); πE != nil {
				continue
			}
			// line 800: def _eval_repr(repr_str):
			πF.SetLineno(800)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "repr_str", Def: nil}
			πTemp052 = πg.NewFunction(πg.NewCode("_eval_repr", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µrepr_str *πg.Object = πArgs[0]; _ = µrepr_str
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
					// line 801: return eval(repr_str)
					πF.SetLineno(801)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrepr_str, "repr_str"); πE != nil {
						continue
					}
					πTemp001[0] = µrepr_str
					if πTemp002, πE = πg.ResolveGlobal(πF, ßeval); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_eval_repr.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 803: def _create_array(f, args, state, npdict=None):
			πF.SetLineno(803)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "f", Def: nil}
			πTemp004[1] = πg.Param{Name: "args", Def: nil}
			πTemp004[2] = πg.Param{Name: "state", Def: nil}
			if πTemp054, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "npdict", Def: πTemp054}
			πTemp053 = πg.NewFunction(πg.NewCode("_create_array", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µf *πg.Object = πArgs[0]; _ = µf
				var µargs *πg.Object = πArgs[1]; _ = µargs
				var µstate *πg.Object = πArgs[2]; _ = µstate
				var µnpdict *πg.Object = πArgs[3]; _ = µnpdict
				var µarray *πg.Object = πg.UnboundLocal; _ = µarray
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
					// line 805: array = f(*args)
					πF.SetLineno(805)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µf, nil, µargs, nil, nil); πE != nil {
						continue
					}
					µarray = πTemp001
					// line 806: array.__setstate__(state)
					πF.SetLineno(806)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp002[0] = µstate
					if πE = πg.CheckLocal(πF, µarray, "array"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µarray, ß__setstate__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.CheckLocal(πF, µnpdict, "npdict"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µnpdict != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 807: if npdict is not None: # we also have saved state in __dict__
					πF.SetLineno(807)
				Label1:
					// line 808: array.__dict__.update(npdict)
					πF.SetLineno(808)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnpdict, "npdict"); πE != nil {
						continue
					}
					πTemp002[0] = µnpdict
					if πE = πg.CheckLocal(πF, µarray, "array"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µarray, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label2
				Label2:
					// line 809: return array
					πF.SetLineno(809)
					if πE = πg.CheckLocal(πF, µarray, "array"); πE != nil {
						continue
					}
					πR = µarray
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_create_array.ToObject(), πTemp053); πE != nil {
				continue
			}
			// line 811: def _create_namedtuple(name, fieldnames, modulename):
			πF.SetLineno(811)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "name", Def: nil}
			πTemp004[1] = πg.Param{Name: "fieldnames", Def: nil}
			πTemp004[2] = πg.Param{Name: "modulename", Def: nil}
			πTemp054 = πg.NewFunction(πg.NewCode("_create_namedtuple", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µfieldnames *πg.Object = πArgs[1]; _ = µfieldnames
				var µmodulename *πg.Object = πArgs[2]; _ = µmodulename
				var µclass_ *πg.Object = πg.UnboundLocal; _ = µclass_
				var µcollections *πg.Object = πg.UnboundLocal; _ = µcollections
				var µt *πg.Object = πg.UnboundLocal; _ = µt
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 812: class_ = _import_module(modulename + '.' + name, safe=True)
					πF.SetLineno(812)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodulename, "modulename"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µmodulename, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, µname); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"safe", πTemp002},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_import_module); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µclass_ = πTemp003
					if πE = πg.CheckLocal(πF, µclass_, "class_"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µclass_ != πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 813: if class_ is not None:
					πF.SetLineno(813)
				Label1:
					// line 814: return class_
					πF.SetLineno(814)
					if πE = πg.CheckLocal(πF, µclass_, "class_"); πE != nil {
						continue
					}
					πR = µclass_
					continue
					goto Label2
				Label2:
					// line 815: import collections
					πF.SetLineno(815)
					if πTemp001, πE = πg.ImportModule(πF, "collections"); πE != nil {
						continue
					}
					πTemp002 = πTemp001[0]
					µcollections = πTemp002
					// line 816: t = collections.namedtuple(name, fieldnames)
					πF.SetLineno(816)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πE = πg.CheckLocal(πF, µfieldnames, "fieldnames"); πE != nil {
						continue
					}
					πTemp001[1] = µfieldnames
					if πE = πg.CheckLocal(πF, µcollections, "collections"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcollections, ßnamedtuple, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µt = πTemp003
					// line 817: t.__module__ = modulename
					πF.SetLineno(817)
					if πE = πg.CheckLocal(πF, µmodulename, "modulename"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µmodulename); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µt, ß__module__, πTemp002); πE != nil {
						continue
					}
					// line 818: return t
					πF.SetLineno(818)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πR = µt
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_create_namedtuple.ToObject(), πTemp054); πE != nil {
				continue
			}
			// line 820: def _getattr(objclass, name, repr_str):
			πF.SetLineno(820)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "objclass", Def: nil}
			πTemp004[1] = πg.Param{Name: "name", Def: nil}
			πTemp004[2] = πg.Param{Name: "repr_str", Def: nil}
			πTemp055 = πg.NewFunction(πg.NewCode("_getattr", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobjclass *πg.Object = πArgs[0]; _ = µobjclass
				var µname *πg.Object = πArgs[1]; _ = µname
				var µrepr_str *πg.Object = πArgs[2]; _ = µrepr_str
				var µattr *πg.Object = πg.UnboundLocal; _ = µattr
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
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.BaseException
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 822: try: #XXX: works only for __builtin__ ?
					πF.SetLineno(822)
					πF.PushCheckpoint(2)
					// line 823: attr = repr_str.split("'")[3]
					πF.SetLineno(823)
					πTemp001 = πg.NewInt(3).ToObject()
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("'").ToObject()
					if πE = πg.CheckLocal(πF, µrepr_str, "repr_str"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µrepr_str, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					µattr = πTemp002
					// line 824: return eval(attr+'.__dict__["'+name+'"]')
					πF.SetLineno(824)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µattr, πg.NewStr(".__dict__[\"").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp004, µname); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr("\"]").ToObject()); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßeval); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					goto Label3
					// line 825: except:
					πF.SetLineno(825)
				Label3:
					// line 826: try:
					πF.SetLineno(826)
					πF.PushCheckpoint(5)
					// line 827: attr = objclass.__dict__
					πF.SetLineno(827)
					if πE = πg.CheckLocal(πF, µobjclass, "objclass"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobjclass, ß__dict__, nil); πE != nil {
						continue
					}
					µattr = πTemp001
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp003[0] = µattr
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßDictProxyType); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp004 == πTemp002).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label6
					}
					goto Label7
					// line 828: if type(attr) is DictProxyType:
					πF.SetLineno(828)
				Label6:
					// line 829: attr = attr[name]
					πF.SetLineno(829)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001 = µname
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µattr, πTemp001); πE != nil {
						continue
					}
					µattr = πTemp002
					goto Label8
				Label7:
					// line 831: attr = getattr(objclass,name)
					πF.SetLineno(831)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobjclass, "objclass"); πE != nil {
						continue
					}
					πTemp003[0] = µobjclass
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp003[1] = µname
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µattr = πTemp002
					goto Label8
				Label8:
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp009, πTemp007 = πF.ExcInfo()
					goto Label9
					// line 832: except:
					πF.SetLineno(832)
				Label9:
					// line 833: attr = getattr(objclass,name)
					πF.SetLineno(833)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobjclass, "objclass"); πE != nil {
						continue
					}
					πTemp003[0] = µobjclass
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp003[1] = µname
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µattr = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					// line 834: return attr
					πF.SetLineno(834)
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πR = µattr
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_getattr.ToObject(), πTemp055); πE != nil {
				continue
			}
			// line 836: def _get_attr(self, name):
			πF.SetLineno(836)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "self", Def: nil}
			πTemp004[1] = πg.Param{Name: "name", Def: nil}
			πTemp056 = πg.NewFunction(πg.NewCode("_get_attr", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µself *πg.Object = πArgs[0]; _ = µself
				var µname *πg.Object = πArgs[1]; _ = µname
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
					// line 838: return getattr(self, name, None) or getattr(__builtin__, name)
					πF.SetLineno(838)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp003[0] = µself
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp003[1] = µname
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
					πTemp003 = πF.MakeArgs(2)
					if πTemp004, πE = πg.ResolveGlobal(πF, ß__builtin__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp003[1] = µname
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
			if πE = πF.Globals().SetItem(πF, ß_get_attr.ToObject(), πTemp056); πE != nil {
				continue
			}
			// line 840: def _dict_from_dictproxy(dictproxy):
			πF.SetLineno(840)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "dictproxy", Def: nil}
			πTemp057 = πg.NewFunction(πg.NewCode("_dict_from_dictproxy", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdictproxy *πg.Object = πArgs[0]; _ = µdictproxy
				var µ_dict *πg.Object = πg.UnboundLocal; _ = µ_dict
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 841: _dict = dictproxy.copy() # convert dictproxy to dict
					πF.SetLineno(841)
					if πE = πg.CheckLocal(πF, µdictproxy, "dictproxy"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdictproxy, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µ_dict = πTemp002
					// line 842: _dict.pop('__dict__', None)
					πF.SetLineno(842)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = ß__dict__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µ_dict, "_dict"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µ_dict, ßpop, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 843: _dict.pop('__weakref__', None)
					πF.SetLineno(843)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = ß__weakref__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µ_dict, "_dict"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µ_dict, ßpop, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 844: _dict.pop('__prepare__', None)
					πF.SetLineno(844)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = ß__prepare__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µ_dict, "_dict"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µ_dict, ßpop, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 845: return _dict
					πF.SetLineno(845)
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
			if πE = πF.Globals().SetItem(πF, ß_dict_from_dictproxy.ToObject(), πTemp057); πE != nil {
				continue
			}
			// line 847: def _import_module(import_name, safe=False):
			πF.SetLineno(847)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "import_name", Def: nil}
			if πTemp059, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "safe", Def: πTemp059}
			πTemp058 = πg.NewFunction(πg.NewCode("_import_module", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µimport_name *πg.Object = πArgs[0]; _ = µimport_name
				var µsafe *πg.Object = πArgs[1]; _ = µsafe
				var µitems *πg.Object = πg.UnboundLocal; _ = µitems
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 848: try:
					πF.SetLineno(848)
					πF.PushCheckpoint(2)
					if πE = πg.CheckLocal(πF, µimport_name, "import_name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µimport_name, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 849: if '.' in import_name:
					πF.SetLineno(849)
				Label3:
					// line 850: items = import_name.split('.')
					πF.SetLineno(850)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr(".").ToObject()
					if πE = πg.CheckLocal(πF, µimport_name, "import_name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µimport_name, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µitems = πTemp004
					// line 851: module = '.'.join(items[:-1])
					πF.SetLineno(851)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µitems, πTemp001); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(".").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µmodule = πTemp004
					// line 852: obj = items[-1]
					πF.SetLineno(852)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp004
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µitems, πTemp001); πE != nil {
						continue
					}
					µobj = πTemp004
					goto Label5
				Label4:
					// line 854: return __import__(import_name)
					πF.SetLineno(854)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µimport_name, "import_name"); πE != nil {
						continue
					}
					πTemp003[0] = µimport_name
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label5
				Label5:
					// line 855: return getattr(__import__(module, None, None, [obj]), obj)
					πF.SetLineno(855)
					πTemp003 = πF.MakeArgs(2)
					πTemp005 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp005[0] = µmodule
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp005[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp005[2] = πTemp001
					πTemp006 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006[0] = µobj
					πTemp001 = πg.NewList(πTemp006...).ToObject()
					πTemp005[3] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[1] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp004, πTemp009).ToObject()
					if πTemp002, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 856: except (ImportError, AttributeError):
					πF.SetLineno(856)
				Label6:
					if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µsafe); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					goto Label8
					// line 857: if safe:
					πF.SetLineno(857)
				Label7:
					// line 858: return None
					πF.SetLineno(858)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label8
				Label8:
					// line 859: raise
					πF.SetLineno(859)
					πE = πF.Raise(nil, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_import_module.ToObject(), πTemp058); πE != nil {
				continue
			}
			// line 861: def _locate_function(obj, session=False):
			πF.SetLineno(861)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp060, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "session", Def: πTemp060}
			πTemp059 = πg.NewFunction(πg.NewCode("_locate_function", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µsession *πg.Object = πArgs[1]; _ = µsession
				var µfound *πg.Object = πg.UnboundLocal; _ = µfound
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß__module__, nil); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 2)
					πTemp003[0] = ß__main__.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πTemp005, πE = πg.Contains(πF, πTemp004, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 862: if obj.__module__ in ['__main__', None]: # and session:
					πF.SetLineno(862)
				Label1:
					// line 863: return False
					πF.SetLineno(863)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 864: found = _import_module(obj.__module__ + '.' + obj.__name__, safe=True)
					πF.SetLineno(864)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__module__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp004, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp004); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"safe", πTemp001},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_import_module); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µfound = πTemp002
					// line 865: return found is obj
					πF.SetLineno(865)
					if πE = πg.CheckLocal(πF, µfound, "found"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µfound == µobj).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_locate_function.ToObject(), πTemp059); πE != nil {
				continue
			}
			// line 879: def save_code(pickler, obj):
			πF.SetLineno(879)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp060 = πg.NewFunction(πg.NewCode("save_code", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
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
					// line 880: log.info("Co: %s" % obj)
					πF.SetLineno(880)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Co: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 881: if PY3:
					πF.SetLineno(881)
				Label1:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					πTemp001[1] = ßco_posonlyargcount.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 882: if hasattr(obj, "co_posonlyargcount"):
					πF.SetLineno(882)
				Label4:
					// line 883: args = (
					πF.SetLineno(883)
					πTemp001 = make([]*πg.Object, 16)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_argcount, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_posonlyargcount, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_kwonlyargcount, nil); πE != nil {
						continue
					}
					πTemp001[2] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_nlocals, nil); πE != nil {
						continue
					}
					πTemp001[3] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_stacksize, nil); πE != nil {
						continue
					}
					πTemp001[4] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_flags, nil); πE != nil {
						continue
					}
					πTemp001[5] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_code, nil); πE != nil {
						continue
					}
					πTemp001[6] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_consts, nil); πE != nil {
						continue
					}
					πTemp001[7] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_names, nil); πE != nil {
						continue
					}
					πTemp001[8] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_varnames, nil); πE != nil {
						continue
					}
					πTemp001[9] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_filename, nil); πE != nil {
						continue
					}
					πTemp001[10] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_name, nil); πE != nil {
						continue
					}
					πTemp001[11] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_firstlineno, nil); πE != nil {
						continue
					}
					πTemp001[12] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_lnotab, nil); πE != nil {
						continue
					}
					πTemp001[13] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_freevars, nil); πE != nil {
						continue
					}
					πTemp001[14] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_cellvars, nil); πE != nil {
						continue
					}
					πTemp001[15] = πTemp003
					πTemp002 = πg.NewTuple(πTemp001...).ToObject()
					µargs = πTemp002
					goto Label6
				Label5:
					// line 892: args = (
					πF.SetLineno(892)
					πTemp001 = make([]*πg.Object, 15)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_argcount, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_kwonlyargcount, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_nlocals, nil); πE != nil {
						continue
					}
					πTemp001[2] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_stacksize, nil); πE != nil {
						continue
					}
					πTemp001[3] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_flags, nil); πE != nil {
						continue
					}
					πTemp001[4] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_code, nil); πE != nil {
						continue
					}
					πTemp001[5] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_consts, nil); πE != nil {
						continue
					}
					πTemp001[6] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_names, nil); πE != nil {
						continue
					}
					πTemp001[7] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_varnames, nil); πE != nil {
						continue
					}
					πTemp001[8] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_filename, nil); πE != nil {
						continue
					}
					πTemp001[9] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_name, nil); πE != nil {
						continue
					}
					πTemp001[10] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_firstlineno, nil); πE != nil {
						continue
					}
					πTemp001[11] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_lnotab, nil); πE != nil {
						continue
					}
					πTemp001[12] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_freevars, nil); πE != nil {
						continue
					}
					πTemp001[13] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_cellvars, nil); πE != nil {
						continue
					}
					πTemp001[14] = πTemp003
					πTemp002 = πg.NewTuple(πTemp001...).ToObject()
					µargs = πTemp002
					goto Label6
				Label6:
					goto Label3
				Label2:
					// line 900: args = (
					πF.SetLineno(900)
					πTemp001 = make([]*πg.Object, 14)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_argcount, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_nlocals, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_stacksize, nil); πE != nil {
						continue
					}
					πTemp001[2] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_flags, nil); πE != nil {
						continue
					}
					πTemp001[3] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_code, nil); πE != nil {
						continue
					}
					πTemp001[4] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_consts, nil); πE != nil {
						continue
					}
					πTemp001[5] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_names, nil); πE != nil {
						continue
					}
					πTemp001[6] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_varnames, nil); πE != nil {
						continue
					}
					πTemp001[7] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_filename, nil); πE != nil {
						continue
					}
					πTemp001[8] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_name, nil); πE != nil {
						continue
					}
					πTemp001[9] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_firstlineno, nil); πE != nil {
						continue
					}
					πTemp001[10] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_lnotab, nil); πE != nil {
						continue
					}
					πTemp001[11] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_freevars, nil); πE != nil {
						continue
					}
					πTemp001[12] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßco_cellvars, nil); πE != nil {
						continue
					}
					πTemp001[13] = πTemp003
					πTemp002 = πg.NewTuple(πTemp001...).ToObject()
					µargs = πTemp002
					goto Label3
				Label3:
					// line 907: pickler.save_reduce(_create_code, args, obj=obj)
					πF.SetLineno(907)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_create_code); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp001[1] = µargs
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 908: log.info("# Co")
					πF.SetLineno(908)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Co").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 909: return
					πF.SetLineno(909)
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
			if πE = πF.Globals().SetItem(πF, ßsave_code.ToObject(), πTemp060); πE != nil {
				continue
			}
			// line 878: @register(CodeType)
			πF.SetLineno(878)
			πTemp001 = πF.MakeArgs(1)
			if πTemp061, πE = πg.ResolveGlobal(πF, ßsave_code); πE != nil {
				continue
			}
			πTemp001[0] = πTemp061
			πTemp017 = πF.MakeArgs(1)
			if πTemp061, πE = πg.ResolveGlobal(πF, ßCodeType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp061
			if πTemp061, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp062, πE = πTemp061.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp061, πE = πTemp062.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_code.ToObject(), πTemp061); πE != nil {
				continue
			}
			// line 912: def save_module_dict(pickler, obj):
			πF.SetLineno(912)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp061 = πg.NewFunction(πg.NewCode("save_module_dict", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 πg.KWArgs
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp003[0] = µpickler
					if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"child", πTemp004},
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßis_dill); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp006
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µpickler, ß_main, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µobj, πTemp007); πE != nil {
						continue
					}
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µpickler, ß_session, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp008).ToObject()
					πTemp001 = πTemp004
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp003[0] = µpickler
					if πTemp006, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"child", πTemp006},
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßis_dill); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp008).ToObject()
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_main_module); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µobj, πTemp007); πE != nil {
						continue
					}
					πTemp001 = πTemp004
				Label3:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µobj, ß__name__.ToObject()); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp008).ToObject()
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_main_module); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.NE(πF, µobj, πTemp007); πE != nil {
						continue
					}
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label5
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp006 = ß__name__.ToObject()
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µobj, πTemp006); πE != nil {
						continue
					}
					πTemp003[0] = πTemp007
					if πTemp006, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp007 == πTemp006).ToObject()
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(3)
					πTemp009 = πF.MakeArgs(2)
					πTemp006 = ß__name__.ToObject()
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µobj, πTemp006); πE != nil {
						continue
					}
					πTemp009[0] = πTemp007
					if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp009[1] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_import_module); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp003[0] = πTemp007
					πTemp003[1] = ß__dict__.ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp004 = πg.GetBool(µobj == πTemp007).ToObject()
					πTemp001 = πTemp004
				Label5:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					goto Label7
					// line 913: if is_dill(pickler, child=False) and obj == pickler._main.__dict__ and not pickler._session:
					πF.SetLineno(913)
				Label2:
					// line 914: log.info("D1: <dict%s" % str(obj.__repr__).split('dict')[-1]) # obj
					πF.SetLineno(914)
					πTemp003 = πF.MakeArgs(1)
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp006
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = ßdict.ToObject()
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobj, ß__repr__, nil); πE != nil {
						continue
					}
					πTemp010[0] = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp007, πE = πg.GetAttr(πF, πTemp011, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp006, πE = πg.GetItem(πF, πTemp011, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("D1: <dict%s").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 915: if PY3:
					πF.SetLineno(915)
				Label9:
					// line 916: pickler.write(bytes('c__builtin__\n__main__\n', 'UTF-8'))
					πF.SetLineno(916)
					πTemp003 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(2)
					πTemp009[0] = πg.NewStr("c__builtin__\n__main__\n").ToObject()
					πTemp009[1] = πg.NewStr("UTF-8").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbytes); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label11
				Label10:
					// line 918: pickler.write('c__builtin__\n__main__\n')
					πF.SetLineno(918)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("c__builtin__\n__main__\n").ToObject()
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label11
				Label11:
					// line 919: log.info("# D1")
					πF.SetLineno(919)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("# D1").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label8
					// line 920: elif (not is_dill(pickler, child=False)) and (obj == _main_module.__dict__):
					πF.SetLineno(920)
				Label4:
					// line 921: log.info("D3: <dict%s" % str(obj.__repr__).split('dict')[-1]) # obj
					πF.SetLineno(921)
					πTemp003 = πF.MakeArgs(1)
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp006
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = ßdict.ToObject()
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobj, ß__repr__, nil); πE != nil {
						continue
					}
					πTemp010[0] = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp007, πE = πg.GetAttr(πF, πTemp011, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp006, πE = πg.GetItem(πF, πTemp011, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("D3: <dict%s").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					goto Label13
					// line 922: if PY3:
					πF.SetLineno(922)
				Label12:
					// line 923: pickler.write(bytes('c__main__\n__dict__\n', 'UTF-8'))
					πF.SetLineno(923)
					πTemp003 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(2)
					πTemp009[0] = πg.NewStr("c__main__\n__dict__\n").ToObject()
					πTemp009[1] = πg.NewStr("UTF-8").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbytes); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label14
				Label13:
					// line 925: pickler.write('c__main__\n__dict__\n')   #XXX: works in general?
					πF.SetLineno(925)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("c__main__\n__dict__\n").ToObject()
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label14
				Label14:
					// line 926: log.info("# D3")
					πF.SetLineno(926)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("# D3").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label8
					// line 927: elif '__name__' in obj and obj != _main_module.__dict__ \
					πF.SetLineno(927)
				Label6:
					// line 930: log.info("D4: <dict%s" % str(obj.__repr__).split('dict')[-1]) # obj
					πF.SetLineno(930)
					πTemp003 = πF.MakeArgs(1)
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp006
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = ßdict.ToObject()
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobj, ß__repr__, nil); πE != nil {
						continue
					}
					πTemp010[0] = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp007, πE = πg.GetAttr(πF, πTemp011, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp006, πE = πg.GetItem(πF, πTemp011, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("D4: <dict%s").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label15
					}
					goto Label16
					// line 931: if PY3:
					πF.SetLineno(931)
				Label15:
					// line 932: pickler.write(bytes('c%s\n__dict__\n' % obj['__name__'], 'UTF-8'))
					πF.SetLineno(932)
					πTemp003 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(2)
					πTemp004 = ß__name__.ToObject()
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µobj, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("c%s\n__dict__\n").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp009[0] = πTemp001
					πTemp009[1] = πg.NewStr("UTF-8").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbytes); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label17
				Label16:
					// line 934: pickler.write('c%s\n__dict__\n' % obj['__name__'])
					πF.SetLineno(934)
					πTemp003 = πF.MakeArgs(1)
					πTemp004 = ß__name__.ToObject()
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µobj, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("c%s\n__dict__\n").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label17
				Label17:
					// line 935: log.info("# D4")
					πF.SetLineno(935)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("# D4").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label8
				Label7:
					// line 937: log.info("D2: <dict%s" % str(obj.__repr__).split('dict')[-1]) # obj
					πF.SetLineno(937)
					πTemp003 = πF.MakeArgs(1)
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp006
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = ßdict.ToObject()
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobj, ß__repr__, nil); πE != nil {
						continue
					}
					πTemp010[0] = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp007, πE = πg.GetAttr(πF, πTemp011, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp006, πE = πg.GetItem(πF, πTemp011, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("D2: <dict%s").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp003[0] = µpickler
					if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"child", πTemp004},
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßis_dill); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp006
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label18
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µpickler, ß_session, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp004
				Label18:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label19
					}
					goto Label20
					// line 938: if is_dill(pickler, child=False) and pickler._session:
					πF.SetLineno(938)
				Label19:
					// line 940: pickler._session = False
					πF.SetLineno(940)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_session, πTemp004); πE != nil {
						continue
					}
					goto Label20
				Label20:
					// line 941: StockPickler.save_dict(pickler, obj)
					πF.SetLineno(941)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp003[0] = µpickler
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[1] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßsave_dict, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 942: log.info("# D2")
					πF.SetLineno(942)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("# D2").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label8
				Label8:
					// line 943: return
					πF.SetLineno(943)
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
			if πE = πF.Globals().SetItem(πF, ßsave_module_dict.ToObject(), πTemp061); πE != nil {
				continue
			}
			// line 911: @register(dict)
			πF.SetLineno(911)
			πTemp001 = πF.MakeArgs(1)
			if πTemp062, πE = πg.ResolveGlobal(πF, ßsave_module_dict); πE != nil {
				continue
			}
			πTemp001[0] = πTemp062
			πTemp017 = πF.MakeArgs(1)
			if πTemp062, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			πTemp017[0] = πTemp062
			if πTemp062, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp063, πE = πTemp062.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp062, πE = πTemp063.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_module_dict.ToObject(), πTemp062); πE != nil {
				continue
			}
			// line 946: def save_classobj(pickler, obj): #FIXME: enable pickler._byref
			πF.SetLineno(946)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp062 = πg.NewFunction(πg.NewCode("save_classobj", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µname *πg.Object = πg.UnboundLocal; _ = µname
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß__module__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp002, ß__main__.ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 948: if obj.__module__ == '__main__': #XXX: use _main_module.__name__ everywhere?
					πF.SetLineno(948)
				Label1:
					// line 949: log.info("C1: %s" % obj)
					πF.SetLineno(949)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("C1: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 950: pickler.save_reduce(ClassType, (obj.__name__, obj.__bases__,
					πF.SetLineno(950)
					πTemp004 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßClassType); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobj, ß__bases__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µobj, ß__dict__, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(πTemp002, πTemp005, πTemp006).ToObject()
					πTemp004[1] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 953: log.info("# C1")
					πF.SetLineno(953)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("# C1").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label3
				Label2:
					// line 955: log.info("C2: %s" % obj)
					πF.SetLineno(955)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("C2: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 956: name = getattr(obj, '__qualname__', getattr(obj, '__name__', None))
					πF.SetLineno(956)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					πTemp004[1] = ß__qualname__.ToObject()
					πTemp008 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp008[0] = µobj
					πTemp008[1] = ß__name__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp008[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp004[2] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µname = πTemp002
					// line 957: StockPickler.save_global(pickler, obj, name=name)
					πF.SetLineno(957)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp004[0] = µpickler
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[1] = µobj
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"name", µname},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsave_global, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 958: log.info("# C2")
					πF.SetLineno(958)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("# C2").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label3
				Label3:
					// line 959: return
					πF.SetLineno(959)
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
			if πE = πF.Globals().SetItem(πF, ßsave_classobj.ToObject(), πTemp062); πE != nil {
				continue
			}
			// line 945: @register(ClassType)
			πF.SetLineno(945)
			πTemp001 = πF.MakeArgs(1)
			if πTemp063, πE = πg.ResolveGlobal(πF, ßsave_classobj); πE != nil {
				continue
			}
			πTemp001[0] = πTemp063
			πTemp017 = πF.MakeArgs(1)
			if πTemp063, πE = πg.ResolveGlobal(πF, ßClassType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp063
			if πTemp063, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp064, πE = πTemp063.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp063, πE = πTemp064.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_classobj.ToObject(), πTemp063); πE != nil {
				continue
			}
			// line 962: def save_lock(pickler, obj):
			πF.SetLineno(962)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp063 = πg.NewFunction(πg.NewCode("save_lock", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
					// line 963: log.info("Lo: %s" % obj)
					πF.SetLineno(963)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Lo: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 964: pickler.save_reduce(_create_lock, (obj.locked(),), obj=obj)
					πF.SetLineno(964)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_create_lock); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßlocked, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(πTemp004).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 965: log.info("# Lo")
					πF.SetLineno(965)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Lo").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 966: return
					πF.SetLineno(966)
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
			if πE = πF.Globals().SetItem(πF, ßsave_lock.ToObject(), πTemp063); πE != nil {
				continue
			}
			// line 961: @register(LockType)
			πF.SetLineno(961)
			πTemp001 = πF.MakeArgs(1)
			if πTemp064, πE = πg.ResolveGlobal(πF, ßsave_lock); πE != nil {
				continue
			}
			πTemp001[0] = πTemp064
			πTemp017 = πF.MakeArgs(1)
			if πTemp064, πE = πg.ResolveGlobal(πF, ßLockType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp064
			if πTemp064, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp065, πE = πTemp064.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp064, πE = πTemp065.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_lock.ToObject(), πTemp064); πE != nil {
				continue
			}
			// line 969: def save_rlock(pickler, obj):
			πF.SetLineno(969)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp064 = πg.NewFunction(πg.NewCode("save_rlock", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µcount *πg.Object = πg.UnboundLocal; _ = µcount
				var µowner *πg.Object = πg.UnboundLocal; _ = µowner
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 πg.KWArgs
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 970: log.info("RL: %s" % obj)
					πF.SetLineno(970)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("RL: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 971: r = obj.__repr__() # don't use _release_save as it unlocks the lock
					πF.SetLineno(971)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß__repr__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µr = πTemp003
					// line 972: count = int(r.split('count=')[1].split()[0].rstrip('>'))
					πF.SetLineno(972)
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr(">").ToObject()
					πTemp002 = πg.NewInt(0).ToObject()
					πTemp005 = πg.NewInt(1).ToObject()
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewStr("count=").ToObject()
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µr, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp006, πE = πg.GetItem(πF, πTemp009, πTemp005); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßrstrip, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcount = πTemp003
					// line 973: owner = int(r.split('owner=')[1].split()[0]) if PY3 else getattr(obj, '_RLock__owner')
					πF.SetLineno(973)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp010 {
						goto Label1
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = πg.NewInt(0).ToObject()
					πTemp006 = πg.NewInt(1).ToObject()
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("owner=").ToObject()
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µr, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp009.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp008, πE = πg.GetItem(πF, πTemp011, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp008, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp005
					goto Label2
				Label1:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					πTemp001[1] = ß_RLock__owner.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp005
				Label2:
					µowner = πTemp002
					// line 974: pickler.save_reduce(_create_rlock, (count,owner,), obj=obj)
					πF.SetLineno(974)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_create_rlock); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µowner, "owner"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µcount, µowner).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp012 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp012); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 975: log.info("# RL")
					πF.SetLineno(975)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# RL").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 976: return
					πF.SetLineno(976)
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
			if πE = πF.Globals().SetItem(πF, ßsave_rlock.ToObject(), πTemp064); πE != nil {
				continue
			}
			// line 968: @register(RLockType)
			πF.SetLineno(968)
			πTemp001 = πF.MakeArgs(1)
			if πTemp065, πE = πg.ResolveGlobal(πF, ßsave_rlock); πE != nil {
				continue
			}
			πTemp001[0] = πTemp065
			πTemp017 = πF.MakeArgs(1)
			if πTemp065, πE = πg.ResolveGlobal(πF, ßRLockType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp065
			if πTemp065, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp066, πE = πTemp065.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp065, πE = πTemp066.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_rlock.ToObject(), πTemp065); πE != nil {
				continue
			}
			if πTemp066, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp066); πE != nil {
				continue
			}
			πTemp065 = πg.GetBool(!πTemp007).ToObject()
			if πTemp007, πE = πg.IsTrue(πF, πTemp065); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label88
			}
			goto Label89
			// line 978: if not IS_PYPY:
			πF.SetLineno(978)
		Label88:
			// line 980: def save_socket(pickler, obj):
			πF.SetLineno(980)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp065 = πg.NewFunction(πg.NewCode("save_socket", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
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
					// line 981: log.info("So: %s" % obj)
					πF.SetLineno(981)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("So: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 982: pickler.save_reduce(*reduce_socket(obj))
					πF.SetLineno(982)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßreduce_socket); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Invoke(πF, πTemp002, nil, πTemp003, nil, nil); πE != nil {
						continue
					}
					// line 983: log.info("# So")
					πF.SetLineno(983)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# So").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 984: return
					πF.SetLineno(984)
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
			if πE = πF.Globals().SetItem(πF, ßsave_socket.ToObject(), πTemp065); πE != nil {
				continue
			}
			goto Label89
		Label89:
			// line 987: def save_itemgetter(pickler, obj):
			πF.SetLineno(987)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp066 = πg.NewFunction(πg.NewCode("save_itemgetter", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µhelper *πg.Object = πg.UnboundLocal; _ = µhelper
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
					// line 988: log.info("Ig: %s" % obj)
					πF.SetLineno(988)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Ig: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 989: helper = _itemgetter_helper()
					πF.SetLineno(989)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_itemgetter_helper); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µhelper = πTemp003
					// line 990: obj(helper)
					πF.SetLineno(990)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µhelper, "helper"); πE != nil {
						continue
					}
					πTemp001[0] = µhelper
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = µobj.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 991: pickler.save_reduce(type(obj), tuple(helper.items), obj=obj)
					πF.SetLineno(991)
					πTemp001 = πF.MakeArgs(2)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001[0] = πTemp003
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µhelper, "helper"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µhelper, ßitems, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 992: log.info("# Ig")
					πF.SetLineno(992)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Ig").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 993: return
					πF.SetLineno(993)
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
			if πE = πF.Globals().SetItem(πF, ßsave_itemgetter.ToObject(), πTemp066); πE != nil {
				continue
			}
			// line 986: @register(ItemGetterType)
			πF.SetLineno(986)
			πTemp001 = πF.MakeArgs(1)
			if πTemp067, πE = πg.ResolveGlobal(πF, ßsave_itemgetter); πE != nil {
				continue
			}
			πTemp001[0] = πTemp067
			πTemp017 = πF.MakeArgs(1)
			if πTemp067, πE = πg.ResolveGlobal(πF, ßItemGetterType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp067
			if πTemp067, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp068, πE = πTemp067.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp067, πE = πTemp068.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_itemgetter.ToObject(), πTemp067); πE != nil {
				continue
			}
			// line 996: def save_attrgetter(pickler, obj):
			πF.SetLineno(996)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp067 = πg.NewFunction(πg.NewCode("save_attrgetter", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µattrs *πg.Object = πg.UnboundLocal; _ = µattrs
				var µhelper *πg.Object = πg.UnboundLocal; _ = µhelper
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
					// line 997: log.info("Ag: %s" % obj)
					πF.SetLineno(997)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Ag: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 998: attrs = []
					πF.SetLineno(998)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µattrs = πTemp002
					// line 999: helper = _attrgetter_helper(attrs)
					πF.SetLineno(999)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
						continue
					}
					πTemp001[0] = µattrs
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_attrgetter_helper); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µhelper = πTemp003
					// line 1000: obj(helper)
					πF.SetLineno(1000)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µhelper, "helper"); πE != nil {
						continue
					}
					πTemp001[0] = µhelper
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = µobj.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1001: pickler.save_reduce(type(obj), tuple(attrs), obj=obj)
					πF.SetLineno(1001)
					πTemp001 = πF.MakeArgs(2)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001[0] = πTemp003
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
						continue
					}
					πTemp004[0] = µattrs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1002: log.info("# Ag")
					πF.SetLineno(1002)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Ag").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1003: return
					πF.SetLineno(1003)
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
			if πE = πF.Globals().SetItem(πF, ßsave_attrgetter.ToObject(), πTemp067); πE != nil {
				continue
			}
			// line 995: @register(AttrGetterType)
			πF.SetLineno(995)
			πTemp001 = πF.MakeArgs(1)
			if πTemp068, πE = πg.ResolveGlobal(πF, ßsave_attrgetter); πE != nil {
				continue
			}
			πTemp001[0] = πTemp068
			πTemp017 = πF.MakeArgs(1)
			if πTemp068, πE = πg.ResolveGlobal(πF, ßAttrGetterType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp068
			if πTemp068, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp069, πE = πTemp068.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp068, πE = πTemp069.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_attrgetter.ToObject(), πTemp068); πE != nil {
				continue
			}
			// line 1005: def _save_file(pickler, obj, open_):
			πF.SetLineno(1005)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp004[2] = πg.Param{Name: "open_", Def: nil}
			πTemp068 = πg.NewFunction(πg.NewCode("_save_file", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µopen_ *πg.Object = πArgs[2]; _ = µopen_
				var µposition *πg.Object = πg.UnboundLocal; _ = µposition
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µfdata *πg.Object = πg.UnboundLocal; _ = µfdata
				var µstrictio *πg.Object = πg.UnboundLocal; _ = µstrictio
				var µfmode *πg.Object = πg.UnboundLocal; _ = µfmode
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 πg.KWArgs
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
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobj, ßclosed, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 1006: if obj.closed:
					πF.SetLineno(1006)
				Label1:
					// line 1007: position = 0
					πF.SetLineno(1007)
					µposition = πg.NewInt(0).ToObject()
					goto Label3
				Label2:
					// line 1009: obj.flush()
					πF.SetLineno(1009)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobj, ßflush, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__stdout__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ß__stderr__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp004, ß__stdin__, nil); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple3(πTemp005, πTemp006, πTemp007).ToObject()
					if πTemp002, πE = πg.Contains(πF, πTemp003, µobj); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 1010: if obj in (sys.__stdout__, sys.__stderr__, sys.__stdin__):
					πF.SetLineno(1010)
				Label4:
					// line 1011: position = -1
					πF.SetLineno(1011)
					if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µposition = πTemp001
					goto Label6
				Label5:
					// line 1013: position = obj.tell()
					πF.SetLineno(1013)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobj, ßtell, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µposition = πTemp003
					goto Label6
				Label6:
					goto Label3
				Label3:
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp008[0] = µpickler
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"child", πTemp003},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßis_dill); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp008, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µpickler, ß_fmode, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßFILE_FMODE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label7:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					goto Label9
					// line 1014: if is_dill(pickler, child=True) and pickler._fmode == FILE_FMODE:
					πF.SetLineno(1014)
				Label8:
					// line 1015: f = open_(obj.name, "r")
					πF.SetLineno(1015)
					πTemp008 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobj, ßname, nil); πE != nil {
						continue
					}
					πTemp008[0] = πTemp001
					πTemp008[1] = ßr.ToObject()
					if πE = πg.CheckLocal(πF, µopen_, "open_"); πE != nil {
						continue
					}
					if πTemp001, πE = µopen_.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					µf = πTemp001
					// line 1016: fdata = f.read()
					πF.SetLineno(1016)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfdata = πTemp003
					// line 1017: f.close()
					πF.SetLineno(1017)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label10
				Label9:
					// line 1019: fdata = ""
					πF.SetLineno(1019)
					µfdata = ß.ToObject()
					goto Label10
				Label10:
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp008[0] = µpickler
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"child", πTemp001},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßis_dill); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp008, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label11
					}
					goto Label12
					// line 1020: if is_dill(pickler, child=True):
					πF.SetLineno(1020)
				Label11:
					// line 1021: strictio = pickler._strictio
					πF.SetLineno(1021)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ß_strictio, nil); πE != nil {
						continue
					}
					µstrictio = πTemp001
					// line 1022: fmode = pickler._fmode
					πF.SetLineno(1022)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ß_fmode, nil); πE != nil {
						continue
					}
					µfmode = πTemp001
					goto Label13
				Label12:
					// line 1024: strictio = False
					πF.SetLineno(1024)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µstrictio = πTemp001
					// line 1025: fmode = 0 # HANDLE_FMODE
					πF.SetLineno(1025)
					µfmode = πg.NewInt(0).ToObject()
					goto Label13
				Label13:
					// line 1026: pickler.save_reduce(_create_filehandle, (obj.name, obj.mode, position,
					πF.SetLineno(1026)
					πTemp008 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_create_filehandle); πE != nil {
						continue
					}
					πTemp008[0] = πTemp001
					πTemp010 = make([]*πg.Object, 8)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßname, nil); πE != nil {
						continue
					}
					πTemp010[0] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßmode, nil); πE != nil {
						continue
					}
					πTemp010[1] = πTemp003
					if πE = πg.CheckLocal(πF, µposition, "position"); πE != nil {
						continue
					}
					πTemp010[2] = µposition
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßclosed, nil); πE != nil {
						continue
					}
					πTemp010[3] = πTemp003
					if πE = πg.CheckLocal(πF, µopen_, "open_"); πE != nil {
						continue
					}
					πTemp010[4] = µopen_
					if πE = πg.CheckLocal(πF, µstrictio, "strictio"); πE != nil {
						continue
					}
					πTemp010[5] = µstrictio
					if πE = πg.CheckLocal(πF, µfmode, "fmode"); πE != nil {
						continue
					}
					πTemp010[6] = µfmode
					if πE = πg.CheckLocal(πF, µfdata, "fdata"); πE != nil {
						continue
					}
					πTemp010[7] = µfdata
					πTemp001 = πg.NewTuple(πTemp010...).ToObject()
					πTemp008[1] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp008, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					// line 1029: return
					πF.SetLineno(1029)
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
			if πE = πF.Globals().SetItem(πF, ß_save_file.ToObject(), πTemp068); πE != nil {
				continue
			}
			// line 1037: def save_file(pickler, obj):
			πF.SetLineno(1037)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp069 = πg.NewFunction(πg.NewCode("save_file", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µf *πg.Object = πg.UnboundLocal; _ = µf
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
					// line 1038: log.info("Fi: %s" % obj)
					πF.SetLineno(1038)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Fi: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1039: f = _save_file(pickler, obj, open)
					πF.SetLineno(1039)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp001[0] = µpickler
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[1] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_save_file); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µf = πTemp003
					// line 1040: log.info("# Fi")
					πF.SetLineno(1040)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Fi").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1041: return f
					πF.SetLineno(1041)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πR = µf
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsave_file.ToObject(), πTemp069); πE != nil {
				continue
			}
			// line 1032: @register(FileType) #XXX: in 3.x has buffer=0, needs different _create?
			πF.SetLineno(1032)
			πTemp001 = πF.MakeArgs(1)
			if πTemp070, πE = πg.ResolveGlobal(πF, ßsave_file); πE != nil {
				continue
			}
			πTemp001[0] = πTemp070
			πTemp017 = πF.MakeArgs(1)
			if πTemp070, πE = πg.ResolveGlobal(πF, ßTextWrapperType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp070
			if πTemp070, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp071, πE = πTemp070.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp070, πE = πTemp071.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_file.ToObject(), πTemp070); πE != nil {
				continue
			}
			// line 1032: @register(FileType) #XXX: in 3.x has buffer=0, needs different _create?
			πF.SetLineno(1032)
			πTemp001 = πF.MakeArgs(1)
			if πTemp070, πE = πg.ResolveGlobal(πF, ßsave_file); πE != nil {
				continue
			}
			πTemp001[0] = πTemp070
			πTemp017 = πF.MakeArgs(1)
			if πTemp070, πE = πg.ResolveGlobal(πF, ßBufferedWriterType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp070
			if πTemp070, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp071, πE = πTemp070.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp070, πE = πTemp071.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_file.ToObject(), πTemp070); πE != nil {
				continue
			}
			// line 1032: @register(FileType) #XXX: in 3.x has buffer=0, needs different _create?
			πF.SetLineno(1032)
			πTemp001 = πF.MakeArgs(1)
			if πTemp070, πE = πg.ResolveGlobal(πF, ßsave_file); πE != nil {
				continue
			}
			πTemp001[0] = πTemp070
			πTemp017 = πF.MakeArgs(1)
			if πTemp070, πE = πg.ResolveGlobal(πF, ßBufferedReaderType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp070
			if πTemp070, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp071, πE = πTemp070.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp070, πE = πTemp071.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_file.ToObject(), πTemp070); πE != nil {
				continue
			}
			// line 1032: @register(FileType) #XXX: in 3.x has buffer=0, needs different _create?
			πF.SetLineno(1032)
			πTemp001 = πF.MakeArgs(1)
			if πTemp070, πE = πg.ResolveGlobal(πF, ßsave_file); πE != nil {
				continue
			}
			πTemp001[0] = πTemp070
			πTemp017 = πF.MakeArgs(1)
			if πTemp070, πE = πg.ResolveGlobal(πF, ßBufferedRandomType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp070
			if πTemp070, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp071, πE = πTemp070.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp070, πE = πTemp071.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_file.ToObject(), πTemp070); πE != nil {
				continue
			}
			// line 1032: @register(FileType) #XXX: in 3.x has buffer=0, needs different _create?
			πF.SetLineno(1032)
			πTemp001 = πF.MakeArgs(1)
			if πTemp070, πE = πg.ResolveGlobal(πF, ßsave_file); πE != nil {
				continue
			}
			πTemp001[0] = πTemp070
			πTemp017 = πF.MakeArgs(1)
			if πTemp070, πE = πg.ResolveGlobal(πF, ßFileType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp070
			if πTemp070, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp071, πE = πTemp070.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp070, πE = πTemp071.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_file.ToObject(), πTemp070); πE != nil {
				continue
			}
			if πTemp070, πE = πg.ResolveGlobal(πF, ßPyTextWrapperType); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp070); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label90
			}
			goto Label91
			// line 1043: if PyTextWrapperType:
			πF.SetLineno(1043)
		Label90:
			// line 1048: def save_file(pickler, obj):
			πF.SetLineno(1048)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp070 = πg.NewFunction(πg.NewCode("save_file", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µf *πg.Object = πg.UnboundLocal; _ = µf
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
					// line 1049: log.info("Fi: %s" % obj)
					πF.SetLineno(1049)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Fi: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1050: f = _save_file(pickler, obj, _open)
					πF.SetLineno(1050)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp001[0] = µpickler
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[1] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_open); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_save_file); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µf = πTemp003
					// line 1051: log.info("# Fi")
					πF.SetLineno(1051)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Fi").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1052: return f
					πF.SetLineno(1052)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πR = µf
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsave_file.ToObject(), πTemp070); πE != nil {
				continue
			}
			// line 1044: @register(PyBufferedRandomType)
			πF.SetLineno(1044)
			πTemp001 = πF.MakeArgs(1)
			if πTemp071, πE = πg.ResolveGlobal(πF, ßsave_file); πE != nil {
				continue
			}
			πTemp001[0] = πTemp071
			πTemp017 = πF.MakeArgs(1)
			if πTemp071, πE = πg.ResolveGlobal(πF, ßPyTextWrapperType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp071
			if πTemp071, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp072, πE = πTemp071.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp071, πE = πTemp072.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_file.ToObject(), πTemp071); πE != nil {
				continue
			}
			// line 1044: @register(PyBufferedRandomType)
			πF.SetLineno(1044)
			πTemp001 = πF.MakeArgs(1)
			if πTemp071, πE = πg.ResolveGlobal(πF, ßsave_file); πE != nil {
				continue
			}
			πTemp001[0] = πTemp071
			πTemp017 = πF.MakeArgs(1)
			if πTemp071, πE = πg.ResolveGlobal(πF, ßPyBufferedWriterType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp071
			if πTemp071, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp072, πE = πTemp071.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp071, πE = πTemp072.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_file.ToObject(), πTemp071); πE != nil {
				continue
			}
			// line 1044: @register(PyBufferedRandomType)
			πF.SetLineno(1044)
			πTemp001 = πF.MakeArgs(1)
			if πTemp071, πE = πg.ResolveGlobal(πF, ßsave_file); πE != nil {
				continue
			}
			πTemp001[0] = πTemp071
			πTemp017 = πF.MakeArgs(1)
			if πTemp071, πE = πg.ResolveGlobal(πF, ßPyBufferedReaderType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp071
			if πTemp071, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp072, πE = πTemp071.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp071, πE = πTemp072.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_file.ToObject(), πTemp071); πE != nil {
				continue
			}
			// line 1044: @register(PyBufferedRandomType)
			πF.SetLineno(1044)
			πTemp001 = πF.MakeArgs(1)
			if πTemp071, πE = πg.ResolveGlobal(πF, ßsave_file); πE != nil {
				continue
			}
			πTemp001[0] = πTemp071
			πTemp017 = πF.MakeArgs(1)
			if πTemp071, πE = πg.ResolveGlobal(πF, ßPyBufferedRandomType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp071
			if πTemp071, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp072, πE = πTemp071.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp071, πE = πTemp072.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_file.ToObject(), πTemp071); πE != nil {
				continue
			}
			goto Label91
		Label91:
			if πTemp071, πE = πg.ResolveGlobal(πF, ßInputType); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp071); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label92
			}
			goto Label93
			// line 1058: if InputType:
			πF.SetLineno(1058)
		Label92:
			// line 1060: def save_stringi(pickler, obj):
			πF.SetLineno(1060)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp071 = πg.NewFunction(πg.NewCode("save_stringi", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
				var µposition *πg.Object = πg.UnboundLocal; _ = µposition
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
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
					// line 1061: log.info("Io: %s" % obj)
					πF.SetLineno(1061)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Io: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ßclosed, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 1062: if obj.closed:
					πF.SetLineno(1062)
				Label1:
					// line 1063: value = ''; position = 0
					πF.SetLineno(1063)
					µvalue = ß.ToObject()
					// line 1063: value = ''; position = 0
					πF.SetLineno(1063)
					µposition = πg.NewInt(0).ToObject()
					goto Label3
				Label2:
					// line 1065: value = obj.getvalue(); position = obj.tell()
					πF.SetLineno(1065)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ßgetvalue, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µvalue = πTemp003
					// line 1065: value = obj.getvalue(); position = obj.tell()
					πF.SetLineno(1065)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ßtell, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µposition = πTemp003
					goto Label3
				Label3:
					// line 1066: pickler.save_reduce(_create_stringi, (value, position, \
					πF.SetLineno(1066)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_create_stringi); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µposition, "position"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßclosed, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µvalue, µposition, πTemp003).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1068: log.info("# Io")
					πF.SetLineno(1068)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Io").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1069: return
					πF.SetLineno(1069)
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
			if πE = πF.Globals().SetItem(πF, ßsave_stringi.ToObject(), πTemp071); πE != nil {
				continue
			}
			// line 1059: @register(InputType)
			πF.SetLineno(1059)
			πTemp001 = πF.MakeArgs(1)
			if πTemp072, πE = πg.ResolveGlobal(πF, ßsave_stringi); πE != nil {
				continue
			}
			πTemp001[0] = πTemp072
			πTemp017 = πF.MakeArgs(1)
			if πTemp072, πE = πg.ResolveGlobal(πF, ßInputType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp072
			if πTemp072, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp073, πE = πTemp072.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp072, πE = πTemp073.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_stringi.ToObject(), πTemp072); πE != nil {
				continue
			}
			// line 1072: def save_stringo(pickler, obj):
			πF.SetLineno(1072)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp072 = πg.NewFunction(πg.NewCode("save_stringo", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
				var µposition *πg.Object = πg.UnboundLocal; _ = µposition
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
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
					// line 1073: log.info("Io: %s" % obj)
					πF.SetLineno(1073)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Io: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ßclosed, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 1074: if obj.closed:
					πF.SetLineno(1074)
				Label1:
					// line 1075: value = ''; position = 0
					πF.SetLineno(1075)
					µvalue = ß.ToObject()
					// line 1075: value = ''; position = 0
					πF.SetLineno(1075)
					µposition = πg.NewInt(0).ToObject()
					goto Label3
				Label2:
					// line 1077: value = obj.getvalue(); position = obj.tell()
					πF.SetLineno(1077)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ßgetvalue, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µvalue = πTemp003
					// line 1077: value = obj.getvalue(); position = obj.tell()
					πF.SetLineno(1077)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ßtell, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µposition = πTemp003
					goto Label3
				Label3:
					// line 1078: pickler.save_reduce(_create_stringo, (value, position, \
					πF.SetLineno(1078)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_create_stringo); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µposition, "position"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßclosed, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µvalue, µposition, πTemp003).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1080: log.info("# Io")
					πF.SetLineno(1080)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Io").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1081: return
					πF.SetLineno(1081)
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
			if πE = πF.Globals().SetItem(πF, ßsave_stringo.ToObject(), πTemp072); πE != nil {
				continue
			}
			// line 1071: @register(OutputType)
			πF.SetLineno(1071)
			πTemp001 = πF.MakeArgs(1)
			if πTemp073, πE = πg.ResolveGlobal(πF, ßsave_stringo); πE != nil {
				continue
			}
			πTemp001[0] = πTemp073
			πTemp017 = πF.MakeArgs(1)
			if πTemp073, πE = πg.ResolveGlobal(πF, ßOutputType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp073
			if πTemp073, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp074, πE = πTemp073.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp073, πE = πTemp074.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_stringo.ToObject(), πTemp073); πE != nil {
				continue
			}
			goto Label93
		Label93:
			// line 1084: def save_functor(pickler, obj):
			πF.SetLineno(1084)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp073 = πg.NewFunction(πg.NewCode("save_functor", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
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
					// line 1085: log.info("Fu: %s" % obj)
					πF.SetLineno(1085)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Fu: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1086: pickler.save_reduce(_create_ftype, (type(obj), obj.func, obj.args,
					πF.SetLineno(1086)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_create_ftype); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßfunc, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µobj, ßargs, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobj, ßkeywords, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple4(πTemp005, πTemp003, πTemp006, πTemp007).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp008 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp008); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1088: log.info("# Fu")
					πF.SetLineno(1088)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Fu").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1089: return
					πF.SetLineno(1089)
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
			if πE = πF.Globals().SetItem(πF, ßsave_functor.ToObject(), πTemp073); πE != nil {
				continue
			}
			// line 1083: @register(PartialType)
			πF.SetLineno(1083)
			πTemp001 = πF.MakeArgs(1)
			if πTemp074, πE = πg.ResolveGlobal(πF, ßsave_functor); πE != nil {
				continue
			}
			πTemp001[0] = πTemp074
			πTemp017 = πF.MakeArgs(1)
			if πTemp074, πE = πg.ResolveGlobal(πF, ßPartialType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp074
			if πTemp074, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp075, πE = πTemp074.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp074, πE = πTemp075.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_functor.ToObject(), πTemp074); πE != nil {
				continue
			}
			// line 1092: def save_super(pickler, obj):
			πF.SetLineno(1092)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp074 = πg.NewFunction(πg.NewCode("save_super", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
					// line 1093: log.info("Su: %s" % obj)
					πF.SetLineno(1093)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Su: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1094: pickler.save_reduce(super, (obj.__thisclass__, obj.__self__), obj=obj)
					πF.SetLineno(1094)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__thisclass__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__self__, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1095: log.info("# Su")
					πF.SetLineno(1095)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Su").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1096: return
					πF.SetLineno(1096)
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
			if πE = πF.Globals().SetItem(πF, ßsave_super.ToObject(), πTemp074); πE != nil {
				continue
			}
			// line 1091: @register(SuperType)
			πF.SetLineno(1091)
			πTemp001 = πF.MakeArgs(1)
			if πTemp075, πE = πg.ResolveGlobal(πF, ßsave_super); πE != nil {
				continue
			}
			πTemp001[0] = πTemp075
			πTemp017 = πF.MakeArgs(1)
			if πTemp075, πE = πg.ResolveGlobal(πF, ßSuperType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp075
			if πTemp075, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp076, πE = πTemp075.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp075, πE = πTemp076.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_super.ToObject(), πTemp075); πE != nil {
				continue
			}
			// line 1099: def save_builtin_method(pickler, obj):
			πF.SetLineno(1099)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp075 = πg.NewFunction(πg.NewCode("save_builtin_method", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var µ_t *πg.Object = πg.UnboundLocal; _ = µ_t
				var µ_recurse *πg.Object = πg.UnboundLocal; _ = µ_recurse
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß__self__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002 != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 1100: if obj.__self__ is not None:
					πF.SetLineno(1100)
				Label1:
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß__self__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß__builtin__); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 1101: if obj.__self__ is __builtin__:
					πF.SetLineno(1101)
				Label4:
					// line 1102: module = 'builtins' if PY3 else '__builtin__'
					πF.SetLineno(1102)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label7
					}
					πTemp001 = ßbuiltins.ToObject()
					goto Label8
				Label7:
					πTemp001 = ß__builtin__.ToObject()
				Label8:
					µmodule = πTemp001
					// line 1103: _t = "B1"
					πF.SetLineno(1103)
					µ_t = ßB1.ToObject()
					// line 1104: log.info("%s: %s" % (_t, obj))
					πF.SetLineno(1104)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_t, "_t"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µ_t, µobj).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s: %s").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label6
				Label5:
					// line 1106: module = obj.__self__
					πF.SetLineno(1106)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobj, ß__self__, nil); πE != nil {
						continue
					}
					µmodule = πTemp001
					// line 1107: _t = "B3"
					πF.SetLineno(1107)
					µ_t = ßB3.ToObject()
					// line 1108: log.info("%s: %s" % (_t, obj))
					πF.SetLineno(1108)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_t, "_t"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µ_t, µobj).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s: %s").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label6
				Label6:
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp005[0] = µpickler
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"child", πTemp001},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßis_dill); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 1109: if is_dill(pickler, child=True):
					πF.SetLineno(1109)
				Label9:
					// line 1110: _recurse = pickler._recurse
					πF.SetLineno(1110)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ß_recurse, nil); πE != nil {
						continue
					}
					µ_recurse = πTemp001
					// line 1111: pickler._recurse = False
					πF.SetLineno(1111)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_recurse, πTemp002); πE != nil {
						continue
					}
					goto Label10
				Label10:
					// line 1112: pickler.save_reduce(_get_attr, (module, obj.__name__), obj=obj)
					πF.SetLineno(1112)
					πTemp005 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_get_attr); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µmodule, πTemp002).ToObject()
					πTemp005[1] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp005[0] = µpickler
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"child", πTemp001},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßis_dill); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label11
					}
					goto Label12
					// line 1113: if is_dill(pickler, child=True):
					πF.SetLineno(1113)
				Label11:
					// line 1114: pickler._recurse = _recurse
					πF.SetLineno(1114)
					if πE = πg.CheckLocal(πF, µ_recurse, "_recurse"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µ_recurse); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_recurse, πTemp001); πE != nil {
						continue
					}
					goto Label12
				Label12:
					// line 1115: log.info("# %s" % _t)
					πF.SetLineno(1115)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_t, "_t"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("# %s").ToObject(), µ_t); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label3
				Label2:
					// line 1117: log.info("B2: %s" % obj)
					πF.SetLineno(1117)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("B2: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 1118: name = getattr(obj, '__qualname__', getattr(obj, '__name__', None))
					πF.SetLineno(1118)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005[0] = µobj
					πTemp005[1] = ß__qualname__.ToObject()
					πTemp007 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007[0] = µobj
					πTemp007[1] = ß__name__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp007[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp005[2] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µname = πTemp002
					// line 1119: StockPickler.save_global(pickler, obj, name=name)
					πF.SetLineno(1119)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp005[0] = µpickler
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005[1] = µobj
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"name", µname},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsave_global, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 1120: log.info("# B2")
					πF.SetLineno(1120)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewStr("# B2").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label3
				Label3:
					// line 1121: return
					πF.SetLineno(1121)
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
			if πE = πF.Globals().SetItem(πF, ßsave_builtin_method.ToObject(), πTemp075); πE != nil {
				continue
			}
			// line 1098: @register(BuiltinMethodType)
			πF.SetLineno(1098)
			πTemp001 = πF.MakeArgs(1)
			if πTemp076, πE = πg.ResolveGlobal(πF, ßsave_builtin_method); πE != nil {
				continue
			}
			πTemp001[0] = πTemp076
			πTemp017 = πF.MakeArgs(1)
			if πTemp076, πE = πg.ResolveGlobal(πF, ßBuiltinMethodType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp076
			if πTemp076, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp077, πE = πTemp076.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp076, πE = πTemp077.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_builtin_method.ToObject(), πTemp076); πE != nil {
				continue
			}
			// line 1124: def save_instancemethod0(pickler, obj):# example: cStringIO.StringI
			πF.SetLineno(1124)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp076 = πg.NewFunction(πg.NewCode("save_instancemethod0", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
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
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1125: log.info("Me: %s" % obj) #XXX: obj.__dict__ handled elsewhere?
					πF.SetLineno(1125)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Me: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 1126: if PY3:
					πF.SetLineno(1126)
				Label1:
					// line 1127: pickler.save_reduce(MethodType, (obj.__func__, obj.__self__), obj=obj)
					πF.SetLineno(1127)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßMethodType); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__func__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobj, ß__self__, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label3
				Label2:
					// line 1129: pickler.save_reduce(MethodType, (obj.im_func, obj.im_self,
					πF.SetLineno(1129)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßMethodType); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßim_func, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobj, ßim_self, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobj, ßim_class, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp003, πTemp005, πTemp007).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label3
				Label3:
					// line 1131: log.info("# Me")
					πF.SetLineno(1131)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Me").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1132: return
					πF.SetLineno(1132)
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
			if πE = πF.Globals().SetItem(πF, ßsave_instancemethod0.ToObject(), πTemp076); πE != nil {
				continue
			}
			// line 1123: @register(MethodType) #FIXME: fails for 'hidden' or 'name-mangled' classes
			πF.SetLineno(1123)
			πTemp001 = πF.MakeArgs(1)
			if πTemp077, πE = πg.ResolveGlobal(πF, ßsave_instancemethod0); πE != nil {
				continue
			}
			πTemp001[0] = πTemp077
			πTemp017 = πF.MakeArgs(1)
			if πTemp077, πE = πg.ResolveGlobal(πF, ßMethodType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp077
			if πTemp077, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp078, πE = πTemp077.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp077, πE = πTemp078.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_instancemethod0.ToObject(), πTemp077); πE != nil {
				continue
			}
			if πTemp078, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp079, πE = πg.GetAttr(πF, πTemp078, ßhexversion, nil); πE != nil {
				continue
			}
			if πTemp077, πE = πg.GE(πF, πTemp079, πg.NewInt(33882352).ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp077); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label94
			}
			if πTemp078, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp078); πE != nil {
				continue
			}
			πTemp077 = πg.GetBool(!πTemp007).ToObject()
			if πTemp007, πE = πg.IsTrue(πF, πTemp077); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label95
			}
			goto Label96
			// line 1134: if sys.hexversion >= 0x20500f0:
			πF.SetLineno(1134)
		Label94:
			if πTemp078, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp078); πE != nil {
				continue
			}
			πTemp077 = πg.GetBool(!πTemp007).ToObject()
			if πTemp007, πE = πg.IsTrue(πF, πTemp077); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label97
			}
			goto Label98
			// line 1135: if not IS_PYPY:
			πF.SetLineno(1135)
		Label97:
			// line 1141: def save_wrapper_descriptor(pickler, obj):
			πF.SetLineno(1141)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp077 = πg.NewFunction(πg.NewCode("save_wrapper_descriptor", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 []*πg.Object
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
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1142: log.info("Wr: %s" % obj)
					πF.SetLineno(1142)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Wr: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1143: pickler.save_reduce(_getattr, (obj.__objclass__, obj.__name__,
					πF.SetLineno(1143)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_getattr); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__objclass__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobj, ß__repr__, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp003, πTemp004, πTemp006).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1145: log.info("# Wr")
					πF.SetLineno(1145)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Wr").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1146: return
					πF.SetLineno(1146)
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
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp077); πE != nil {
				continue
			}
			// line 1136: @register(MemberDescriptorType)
			πF.SetLineno(1136)
			πTemp001 = πF.MakeArgs(1)
			if πTemp078, πE = πg.ResolveGlobal(πF, ßsave_wrapper_descriptor); πE != nil {
				continue
			}
			πTemp001[0] = πTemp078
			πTemp017 = πF.MakeArgs(1)
			if πTemp078, πE = πg.ResolveGlobal(πF, ßClassMethodDescriptorType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp078
			if πTemp078, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp079, πE = πTemp078.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp078, πE = πTemp079.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp078); πE != nil {
				continue
			}
			// line 1136: @register(MemberDescriptorType)
			πF.SetLineno(1136)
			πTemp001 = πF.MakeArgs(1)
			if πTemp078, πE = πg.ResolveGlobal(πF, ßsave_wrapper_descriptor); πE != nil {
				continue
			}
			πTemp001[0] = πTemp078
			πTemp017 = πF.MakeArgs(1)
			if πTemp078, πE = πg.ResolveGlobal(πF, ßWrapperDescriptorType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp078
			if πTemp078, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp079, πE = πTemp078.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp078, πE = πTemp079.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp078); πE != nil {
				continue
			}
			// line 1136: @register(MemberDescriptorType)
			πF.SetLineno(1136)
			πTemp001 = πF.MakeArgs(1)
			if πTemp078, πE = πg.ResolveGlobal(πF, ßsave_wrapper_descriptor); πE != nil {
				continue
			}
			πTemp001[0] = πTemp078
			πTemp017 = πF.MakeArgs(1)
			if πTemp078, πE = πg.ResolveGlobal(πF, ßMethodDescriptorType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp078
			if πTemp078, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp079, πE = πTemp078.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp078, πE = πTemp079.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp078); πE != nil {
				continue
			}
			// line 1136: @register(MemberDescriptorType)
			πF.SetLineno(1136)
			πTemp001 = πF.MakeArgs(1)
			if πTemp078, πE = πg.ResolveGlobal(πF, ßsave_wrapper_descriptor); πE != nil {
				continue
			}
			πTemp001[0] = πTemp078
			πTemp017 = πF.MakeArgs(1)
			if πTemp078, πE = πg.ResolveGlobal(πF, ßGetSetDescriptorType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp078
			if πTemp078, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp079, πE = πTemp078.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp078, πE = πTemp079.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp078); πE != nil {
				continue
			}
			// line 1136: @register(MemberDescriptorType)
			πF.SetLineno(1136)
			πTemp001 = πF.MakeArgs(1)
			if πTemp078, πE = πg.ResolveGlobal(πF, ßsave_wrapper_descriptor); πE != nil {
				continue
			}
			πTemp001[0] = πTemp078
			πTemp017 = πF.MakeArgs(1)
			if πTemp078, πE = πg.ResolveGlobal(πF, ßMemberDescriptorType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp078
			if πTemp078, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp079, πE = πTemp078.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp078, πE = πTemp079.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp078); πE != nil {
				continue
			}
			goto Label99
		Label98:
			// line 1150: def save_wrapper_descriptor(pickler, obj):
			πF.SetLineno(1150)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp078 = πg.NewFunction(πg.NewCode("save_wrapper_descriptor", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 []*πg.Object
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
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1151: log.info("Wr: %s" % obj)
					πF.SetLineno(1151)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Wr: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1152: pickler.save_reduce(_getattr, (obj.__objclass__, obj.__name__,
					πF.SetLineno(1152)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_getattr); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__objclass__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobj, ß__repr__, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp003, πTemp004, πTemp006).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1154: log.info("# Wr")
					πF.SetLineno(1154)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Wr").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1155: return
					πF.SetLineno(1155)
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
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp078); πE != nil {
				continue
			}
			// line 1148: @register(MemberDescriptorType)
			πF.SetLineno(1148)
			πTemp001 = πF.MakeArgs(1)
			if πTemp079, πE = πg.ResolveGlobal(πF, ßsave_wrapper_descriptor); πE != nil {
				continue
			}
			πTemp001[0] = πTemp079
			πTemp017 = πF.MakeArgs(1)
			if πTemp079, πE = πg.ResolveGlobal(πF, ßGetSetDescriptorType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp079
			if πTemp079, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp080, πE = πTemp079.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp079, πE = πTemp080.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp079); πE != nil {
				continue
			}
			// line 1148: @register(MemberDescriptorType)
			πF.SetLineno(1148)
			πTemp001 = πF.MakeArgs(1)
			if πTemp079, πE = πg.ResolveGlobal(πF, ßsave_wrapper_descriptor); πE != nil {
				continue
			}
			πTemp001[0] = πTemp079
			πTemp017 = πF.MakeArgs(1)
			if πTemp079, πE = πg.ResolveGlobal(πF, ßMemberDescriptorType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp079
			if πTemp079, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp080, πE = πTemp079.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp079, πE = πTemp080.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp079); πE != nil {
				continue
			}
			goto Label99
		Label99:
			// line 1158: def save_instancemethod(pickler, obj):
			πF.SetLineno(1158)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp079 = πg.NewFunction(πg.NewCode("save_instancemethod", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
					// line 1159: log.info("Mw: %s" % obj)
					πF.SetLineno(1159)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Mw: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1160: pickler.save_reduce(getattr, (obj.__self__, obj.__name__), obj=obj)
					πF.SetLineno(1160)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__self__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1161: log.info("# Mw")
					πF.SetLineno(1161)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Mw").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1162: return
					πF.SetLineno(1162)
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
			if πE = πF.Globals().SetItem(πF, ßsave_instancemethod.ToObject(), πTemp079); πE != nil {
				continue
			}
			// line 1157: @register(MethodWrapperType)
			πF.SetLineno(1157)
			πTemp001 = πF.MakeArgs(1)
			if πTemp080, πE = πg.ResolveGlobal(πF, ßsave_instancemethod); πE != nil {
				continue
			}
			πTemp001[0] = πTemp080
			πTemp017 = πF.MakeArgs(1)
			if πTemp080, πE = πg.ResolveGlobal(πF, ßMethodWrapperType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp080
			if πTemp080, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp081, πE = πTemp080.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp080, πE = πTemp081.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_instancemethod.ToObject(), πTemp080); πE != nil {
				continue
			}
			goto Label96
			// line 1164: elif not IS_PYPY:
			πF.SetLineno(1164)
		Label95:
			// line 1167: def save_wrapper_descriptor(pickler, obj):
			πF.SetLineno(1167)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp080 = πg.NewFunction(πg.NewCode("save_wrapper_descriptor", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 []*πg.Object
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
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1168: log.info("Wr: %s" % obj)
					πF.SetLineno(1168)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Wr: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1169: pickler.save_reduce(_getattr, (obj.__objclass__, obj.__name__,
					πF.SetLineno(1169)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_getattr); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__objclass__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobj, ß__repr__, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp003, πTemp004, πTemp006).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1171: log.info("# Wr")
					πF.SetLineno(1171)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Wr").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1172: return
					πF.SetLineno(1172)
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
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp080); πE != nil {
				continue
			}
			// line 1165: @register(MethodDescriptorType)
			πF.SetLineno(1165)
			πTemp001 = πF.MakeArgs(1)
			if πTemp081, πE = πg.ResolveGlobal(πF, ßsave_wrapper_descriptor); πE != nil {
				continue
			}
			πTemp001[0] = πTemp081
			πTemp017 = πF.MakeArgs(1)
			if πTemp081, πE = πg.ResolveGlobal(πF, ßWrapperDescriptorType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp081
			if πTemp081, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp082, πE = πTemp081.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp081, πE = πTemp082.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp081); πE != nil {
				continue
			}
			// line 1165: @register(MethodDescriptorType)
			πF.SetLineno(1165)
			πTemp001 = πF.MakeArgs(1)
			if πTemp081, πE = πg.ResolveGlobal(πF, ßsave_wrapper_descriptor); πE != nil {
				continue
			}
			πTemp001[0] = πTemp081
			πTemp017 = πF.MakeArgs(1)
			if πTemp081, πE = πg.ResolveGlobal(πF, ßMethodDescriptorType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp081
			if πTemp081, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp082, πE = πTemp081.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp081, πE = πTemp082.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_wrapper_descriptor.ToObject(), πTemp081); πE != nil {
				continue
			}
			goto Label96
		Label96:
			// line 1175: def save_cell(pickler, obj):
			πF.SetLineno(1175)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp081 = πg.NewFunction(πg.NewCode("save_cell", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 πg.KWArgs
				_ = πTemp004
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1176: log.info("Ce: %s" % obj)
					πF.SetLineno(1176)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Ce: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1177: f = obj.cell_contents
					πF.SetLineno(1177)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ßcell_contents, nil); πE != nil {
						continue
					}
					µf = πTemp002
					// line 1178: pickler.save_reduce(_create_cell, (f,), obj=obj)
					πF.SetLineno(1178)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_create_cell); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(µf).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1179: log.info("# Ce")
					πF.SetLineno(1179)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Ce").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1180: return
					πF.SetLineno(1180)
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
			if πE = πF.Globals().SetItem(πF, ßsave_cell.ToObject(), πTemp081); πE != nil {
				continue
			}
			// line 1174: @register(CellType)
			πF.SetLineno(1174)
			πTemp001 = πF.MakeArgs(1)
			if πTemp082, πE = πg.ResolveGlobal(πF, ßsave_cell); πE != nil {
				continue
			}
			πTemp001[0] = πTemp082
			πTemp017 = πF.MakeArgs(1)
			if πTemp082, πE = πg.ResolveGlobal(πF, ßCellType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp082
			if πTemp082, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp083, πE = πTemp082.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp082, πE = πTemp083.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_cell.ToObject(), πTemp082); πE != nil {
				continue
			}
			if πTemp083, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp083); πE != nil {
				continue
			}
			πTemp082 = πg.GetBool(!πTemp007).ToObject()
			if πTemp007, πE = πg.IsTrue(πF, πTemp082); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label100
			}
			goto Label101
			// line 1182: if not IS_PYPY:
			πF.SetLineno(1182)
		Label100:
			if πTemp083, πE = πg.ResolveGlobal(πF, ßOLD33); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp083); πE != nil {
				continue
			}
			πTemp082 = πg.GetBool(!πTemp007).ToObject()
			if πTemp007, πE = πg.IsTrue(πF, πTemp082); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label102
			}
			goto Label103
			// line 1183: if not OLD33:
			πF.SetLineno(1183)
		Label102:
			// line 1185: def save_dictproxy(pickler, obj):
			πF.SetLineno(1185)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp082 = πg.NewFunction(πg.NewCode("save_dictproxy", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
					// line 1186: log.info("Mp: %s" % obj)
					πF.SetLineno(1186)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Mp: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1187: pickler.save_reduce(DictProxyType, (obj.copy(),), obj=obj)
					πF.SetLineno(1187)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßDictProxyType); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(πTemp004).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1188: log.info("# Mp")
					πF.SetLineno(1188)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Mp").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1189: return
					πF.SetLineno(1189)
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
			if πE = πF.Globals().SetItem(πF, ßsave_dictproxy.ToObject(), πTemp082); πE != nil {
				continue
			}
			// line 1184: @register(DictProxyType)
			πF.SetLineno(1184)
			πTemp001 = πF.MakeArgs(1)
			if πTemp083, πE = πg.ResolveGlobal(πF, ßsave_dictproxy); πE != nil {
				continue
			}
			πTemp001[0] = πTemp083
			πTemp017 = πF.MakeArgs(1)
			if πTemp083, πE = πg.ResolveGlobal(πF, ßDictProxyType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp083
			if πTemp083, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp084, πE = πTemp083.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp083, πE = πTemp084.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_dictproxy.ToObject(), πTemp083); πE != nil {
				continue
			}
			goto Label104
		Label103:
			// line 1195: def save_dictproxy(pickler, obj):
			πF.SetLineno(1195)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp083 = πg.NewFunction(πg.NewCode("save_dictproxy", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µattr *πg.Object = πg.UnboundLocal; _ = µattr
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
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1196: log.info("Dp: %s" % obj)
					πF.SetLineno(1196)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Dp: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1197: attr = obj.get('__dict__')
					πF.SetLineno(1197)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ß__dict__.ToObject()
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ßget, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µattr = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp001[0] = µattr
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßGetSetDescriptorType); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µattr, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp005, ß__dict__.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µattr, ß__objclass__, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					πTemp001[1] = ß__dict__.ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[2] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp006, µobj); πE != nil {
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
					// line 1199: if type(attr) == GetSetDescriptorType and attr.__name__ == "__dict__" \
					πF.SetLineno(1199)
				Label2:
					// line 1201: pickler.save_reduce(getattr, (attr.__objclass__,"__dict__"),obj=obj)
					πF.SetLineno(1201)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µattr, ß__objclass__, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, ß__dict__.ToObject()).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1202: log.info("# Dp")
					πF.SetLineno(1202)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Dp").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1203: return
					πF.SetLineno(1203)
					πR = πg.None
					continue
					goto Label3
				Label3:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s does not reference a class __dict__").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßReferenceError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1205: raise ReferenceError("%s does not reference a class __dict__" % obj)
					πF.SetLineno(1205)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsave_dictproxy.ToObject(), πTemp083); πE != nil {
				continue
			}
			// line 1194: @register(DictProxyType)
			πF.SetLineno(1194)
			πTemp001 = πF.MakeArgs(1)
			if πTemp084, πE = πg.ResolveGlobal(πF, ßsave_dictproxy); πE != nil {
				continue
			}
			πTemp001[0] = πTemp084
			πTemp017 = πF.MakeArgs(1)
			if πTemp084, πE = πg.ResolveGlobal(πF, ßDictProxyType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp084
			if πTemp084, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp085, πE = πTemp084.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp084, πE = πTemp085.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_dictproxy.ToObject(), πTemp084); πE != nil {
				continue
			}
			goto Label104
		Label104:
			goto Label101
		Label101:
			// line 1208: def save_slice(pickler, obj):
			πF.SetLineno(1208)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp084 = πg.NewFunction(πg.NewCode("save_slice", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1209: log.info("Sl: %s" % obj)
					πF.SetLineno(1209)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Sl: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1210: pickler.save_reduce(slice, (obj.start, obj.stop, obj.step), obj=obj)
					πF.SetLineno(1210)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßstart, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ßstop, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobj, ßstep, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp003, πTemp004, πTemp005).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1211: log.info("# Sl")
					πF.SetLineno(1211)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Sl").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1212: return
					πF.SetLineno(1212)
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
			if πE = πF.Globals().SetItem(πF, ßsave_slice.ToObject(), πTemp084); πE != nil {
				continue
			}
			// line 1207: @register(SliceType)
			πF.SetLineno(1207)
			πTemp001 = πF.MakeArgs(1)
			if πTemp085, πE = πg.ResolveGlobal(πF, ßsave_slice); πE != nil {
				continue
			}
			πTemp001[0] = πTemp085
			πTemp017 = πF.MakeArgs(1)
			if πTemp085, πE = πg.ResolveGlobal(πF, ßSliceType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp085
			if πTemp085, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp086, πE = πTemp085.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp085, πE = πTemp086.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_slice.ToObject(), πTemp085); πE != nil {
				continue
			}
			// line 1217: def save_singleton(pickler, obj):
			πF.SetLineno(1217)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp085 = πg.NewFunction(πg.NewCode("save_singleton", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
					// line 1218: log.info("Si: %s" % obj)
					πF.SetLineno(1218)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Si: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1219: pickler.save_reduce(_eval_repr, (obj.__repr__(),), obj=obj)
					πF.SetLineno(1219)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_eval_repr); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__repr__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(πTemp004).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1220: log.info("# Si")
					πF.SetLineno(1220)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Si").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1221: return
					πF.SetLineno(1221)
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
			if πE = πF.Globals().SetItem(πF, ßsave_singleton.ToObject(), πTemp085); πE != nil {
				continue
			}
			// line 1214: @register(XRangeType)
			πF.SetLineno(1214)
			πTemp001 = πF.MakeArgs(1)
			if πTemp086, πE = πg.ResolveGlobal(πF, ßsave_singleton); πE != nil {
				continue
			}
			πTemp001[0] = πTemp086
			πTemp017 = πF.MakeArgs(1)
			if πTemp086, πE = πg.ResolveGlobal(πF, ßNotImplementedType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp086
			if πTemp086, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp087, πE = πTemp086.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp086, πE = πTemp087.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_singleton.ToObject(), πTemp086); πE != nil {
				continue
			}
			// line 1214: @register(XRangeType)
			πF.SetLineno(1214)
			πTemp001 = πF.MakeArgs(1)
			if πTemp086, πE = πg.ResolveGlobal(πF, ßsave_singleton); πE != nil {
				continue
			}
			πTemp001[0] = πTemp086
			πTemp017 = πF.MakeArgs(1)
			if πTemp086, πE = πg.ResolveGlobal(πF, ßEllipsisType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp086
			if πTemp086, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp087, πE = πTemp086.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp086, πE = πTemp087.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_singleton.ToObject(), πTemp086); πE != nil {
				continue
			}
			// line 1214: @register(XRangeType)
			πF.SetLineno(1214)
			πTemp001 = πF.MakeArgs(1)
			if πTemp086, πE = πg.ResolveGlobal(πF, ßsave_singleton); πE != nil {
				continue
			}
			πTemp001[0] = πTemp086
			πTemp017 = πF.MakeArgs(1)
			if πTemp086, πE = πg.ResolveGlobal(πF, ßXRangeType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp086
			if πTemp086, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp087, πE = πTemp086.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp086, πE = πTemp087.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_singleton.ToObject(), πTemp086); πE != nil {
				continue
			}
			// line 1223: def _proxy_helper(obj): # a dead proxy returns a reference to None
			πF.SetLineno(1223)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp086 = πg.NewFunction(πg.NewCode("_proxy_helper", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µ_repr *πg.Object = πg.UnboundLocal; _ = µ_repr
				var µ_str *πg.Object = πg.UnboundLocal; _ = µ_str
				var µaddress *πg.Object = πg.UnboundLocal; _ = µaddress
				var µobjects *πg.Object = πg.UnboundLocal; _ = µobjects
				var µ_obj *πg.Object = πg.UnboundLocal; _ = µ_obj
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 πg.KWArgs
				_ = πTemp011
				var πTemp012 bool
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 12: goto Label12
					case 13: goto Label13
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 1224: """get memory address of proxy's reference object"""
					πF.SetLineno(1224)
					// line 1225: _repr = repr(obj)
					πF.SetLineno(1225)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µ_repr = πTemp003
					// line 1226: try: _str = str(obj)
					πF.SetLineno(1226)
					πF.PushCheckpoint(2)
					// line 1226: try: _str = str(obj)
					πF.SetLineno(1226)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µ_str = πTemp003
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßReferenceError); πE != nil {
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
					// line 1227: except ReferenceError: # it's a dead proxy
					πF.SetLineno(1227)
				Label3:
					// line 1228: return id(None)
					πF.SetLineno(1228)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					if πE = πg.CheckLocal(πF, µ_str, "_str"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_repr, "_repr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µ_str, µ_repr); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					goto Label5
					// line 1229: if _str == _repr: return id(obj) # it's a repr
					πF.SetLineno(1229)
				Label4:
					// line 1229: if _str == _repr: return id(obj) # it's a repr
					πF.SetLineno(1229)
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
					πR = πTemp003
					continue
					goto Label5
				Label5:
					// line 1230: try: # either way, it's a proxy from here
					πF.SetLineno(1230)
					πF.PushCheckpoint(7)
					// line 1231: address = int(_str.rstrip('>').split(' at ')[-1], base=16)
					πF.SetLineno(1231)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewStr(" at ").ToObject()
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = πg.NewStr(">").ToObject()
					if πE = πg.CheckLocal(πF, µ_str, "_str"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µ_str, ßrstrip, nil); πE != nil {
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
					if πTemp003, πE = πg.GetItem(πF, πTemp010, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp011 = πg.KWArgs{
						{"base", πg.NewInt(16).ToObject()},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp011); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µaddress = πTemp003
					πF.PopCheckpoint()
					goto Label6
				Label7:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label8
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 1232: except ValueError: # special case: proxy of a 'type'
					πF.SetLineno(1232)
				Label8:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					goto Label10
					// line 1233: if not IS_PYPY:
					πF.SetLineno(1233)
				Label9:
					// line 1234: address = int(_repr.rstrip('>').split(' at ')[-1], base=16)
					πF.SetLineno(1234)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewStr(" at ").ToObject()
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = πg.NewStr(">").ToObject()
					if πE = πg.CheckLocal(πF, µ_repr, "_repr"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µ_repr, ßrstrip, nil); πE != nil {
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
					if πTemp003, πE = πg.GetItem(πF, πTemp010, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp011 = πg.KWArgs{
						{"base", πg.NewInt(16).ToObject()},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp011); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µaddress = πTemp003
					goto Label11
				Label10:
					// line 1236: objects = iter(gc.get_objects())
					πF.SetLineno(1236)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgc); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_objects, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µobjects = πTemp003
					if πE = πg.CheckLocal(πF, µobjects, "objects"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µobjects); πE != nil {
						continue
					}
					πF.PushCheckpoint(13)
					πTemp006 = false
				Label12:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label14
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µ_obj = πTemp003
					}
					if πE != nil || !πTemp012 {
						continue
					}
					πF.PushCheckpoint(12)            
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_obj, "_obj"); πE != nil {
						continue
					}
					πTemp001[0] = µ_obj
					if πTemp009, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µ_str, "_str"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp010, µ_str); πE != nil {
						continue
					}
					if πTemp012, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp012 {
						goto Label15
					}
					goto Label16
					// line 1238: if repr(_obj) == _str: return id(_obj)
					πF.SetLineno(1238)
				Label15:
					// line 1238: if repr(_obj) == _str: return id(_obj)
					πF.SetLineno(1238)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_obj, "_obj"); πE != nil {
						continue
					}
					πTemp001[0] = µ_obj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp009
					continue
					goto Label16
				Label16:
					continue
				Label13:
					if πE != nil || πR != nil {
						continue
					}
				Label14:
					// line 1240: msg = "Cannot reference object for proxy at '%s'" % id(obj)
					πF.SetLineno(1240)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Cannot reference object for proxy at '%s'").ToObject(), πTemp009); πE != nil {
						continue
					}
					µmsg = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[0] = µmsg
					if πTemp002, πE = πg.ResolveGlobal(πF, ßReferenceError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1241: raise ReferenceError(msg)
					πF.SetLineno(1241)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label11
				Label11:
					πF.RestoreExc(nil, nil)
					goto Label6
				Label6:
					// line 1242: return address
					πF.SetLineno(1242)
					if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
						continue
					}
					πR = µaddress
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_proxy_helper.ToObject(), πTemp086); πE != nil {
				continue
			}
			// line 1224: """get memory address of proxy's reference object"""
			πF.SetLineno(1224)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp087}, πg.NewStr("get memory address of proxy's reference object").ToObject()); πE != nil {
				continue
			}
			if πTemp088, πE = πg.ResolveGlobal(πF, ß_proxy_helper); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp088, ß__doc__, πTemp087); πE != nil {
				continue
			}
			// line 1244: def _locate_object(address, module=None):
			πF.SetLineno(1244)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "address", Def: nil}
			if πTemp088, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "module", Def: πTemp088}
			πTemp087 = πg.NewFunction(πg.NewCode("_locate_object", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µaddress *πg.Object = πArgs[0]; _ = µaddress
				var µmodule *πg.Object = πArgs[1]; _ = µmodule
				var µspecial *πg.Object = πg.UnboundLocal; _ = µspecial
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
				var µobjects *πg.Object = πg.UnboundLocal; _ = µobjects
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
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
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 18: goto Label18
					case 12: goto Label12
					case 13: goto Label13
					default: panic("unexpected function state")
					}
					// line 1245: """get object located at the given memory address (inverse of id(obj))"""
					πF.SetLineno(1245)
					// line 1246: special = [None, True, False] #XXX: more...?
					πF.SetLineno(1246)
					πTemp001 = make([]*πg.Object, 3)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µspecial = πTemp002
					if πE = πg.CheckLocal(πF, µspecial, "special"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µspecial); πE != nil {
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
					if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µobj = πTemp005
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp006, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.Eq(πF, µaddress, πTemp007); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 1248: if address == id(obj): return obj
					πF.SetLineno(1248)
				Label4:
					// line 1248: if address == id(obj): return obj
					πF.SetLineno(1248)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πR = µobj
					continue
					goto Label5
				Label5:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µmodule); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label6
					}
					goto Label7
					// line 1249: if module:
					πF.SetLineno(1249)
				Label6:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label9
					}
					goto Label10
					// line 1250: if PY3:
					πF.SetLineno(1250)
				Label9:
					// line 1251: objects = iter(module.__dict__.values())
					πF.SetLineno(1251)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmodule, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µobjects = πTemp005
					goto Label11
				Label10:
					// line 1253: objects = module.__dict__.itervalues()
					πF.SetLineno(1253)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmodule, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßitervalues, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					µobjects = πTemp002
					goto Label11
				Label11:
					goto Label8
				Label7:
					// line 1254: else: objects = iter(gc.get_objects())
					πF.SetLineno(1254)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgc); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßget_objects, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µobjects = πTemp005
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µobjects, "objects"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µobjects); πE != nil {
						continue
					}
					πF.PushCheckpoint(13)
					πTemp003 = false
				Label12:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label14
					}
					if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µobj = πTemp005
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(12)            
					if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp006, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.Eq(πF, µaddress, πTemp007); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label15
					}
					goto Label16
					// line 1256: if address == id(obj): return obj
					πF.SetLineno(1256)
				Label15:
					// line 1256: if address == id(obj): return obj
					πF.SetLineno(1256)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πR = µobj
					continue
					goto Label16
				Label16:
					continue
				Label13:
					if πE != nil || πR != nil {
						continue
					}
				Label14:
					// line 1258: try: address = hex(address)
					πF.SetLineno(1258)
					πF.PushCheckpoint(18)
					// line 1258: try: address = hex(address)
					πF.SetLineno(1258)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
						continue
					}
					πTemp001[0] = µaddress
					if πTemp002, πE = πg.ResolveGlobal(πF, ßhex); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µaddress = πTemp005
					πF.PopCheckpoint()
					goto Label17
				Label18:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label19
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 1259: except TypeError:
					πF.SetLineno(1259)
				Label19:
					πTemp001 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
						continue
					}
					πTemp010[0] = µaddress
					if πTemp005, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("'%s' is not a valid memory address").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1260: raise TypeError("'%s' is not a valid memory address" % str(address))
					πF.SetLineno(1260)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label17
				Label17:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Cannot reference object at '%s'").ToObject(), µaddress); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßReferenceError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1261: raise ReferenceError("Cannot reference object at '%s'" % address)
					πF.SetLineno(1261)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_locate_object.ToObject(), πTemp087); πE != nil {
				continue
			}
			// line 1245: """get object located at the given memory address (inverse of id(obj))"""
			πF.SetLineno(1245)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp088}, πg.NewStr("get object located at the given memory address (inverse of id(obj))").ToObject()); πE != nil {
				continue
			}
			if πTemp089, πE = πg.ResolveGlobal(πF, ß_locate_object); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp089, ß__doc__, πTemp088); πE != nil {
				continue
			}
			// line 1264: def save_weakref(pickler, obj):
			πF.SetLineno(1264)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp088 = πg.NewFunction(πg.NewCode("save_weakref", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µrefobj *πg.Object = πg.UnboundLocal; _ = µrefobj
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 πg.KWArgs
				_ = πTemp004
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1265: refobj = obj()
					πF.SetLineno(1265)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = µobj.Call(πF, nil, nil); πE != nil {
						continue
					}
					µrefobj = πTemp001
					// line 1266: log.info("R1: %s" % obj)
					πF.SetLineno(1266)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("R1: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 1268: pickler.save_reduce(_create_weakref, (refobj,), obj=obj)
					πF.SetLineno(1268)
					πTemp002 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_create_weakref); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µrefobj, "refobj"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple1(µrefobj).ToObject()
					πTemp002[1] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 1269: log.info("# R1")
					πF.SetLineno(1269)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("# R1").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 1270: return
					πF.SetLineno(1270)
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
			if πE = πF.Globals().SetItem(πF, ßsave_weakref.ToObject(), πTemp088); πE != nil {
				continue
			}
			// line 1263: @register(ReferenceType)
			πF.SetLineno(1263)
			πTemp001 = πF.MakeArgs(1)
			if πTemp089, πE = πg.ResolveGlobal(πF, ßsave_weakref); πE != nil {
				continue
			}
			πTemp001[0] = πTemp089
			πTemp017 = πF.MakeArgs(1)
			if πTemp089, πE = πg.ResolveGlobal(πF, ßReferenceType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp089
			if πTemp089, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp090, πE = πTemp089.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp089, πE = πTemp090.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_weakref.ToObject(), πTemp089); πE != nil {
				continue
			}
			// line 1274: def save_weakproxy(pickler, obj):
			πF.SetLineno(1274)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp089 = πg.NewFunction(πg.NewCode("save_weakproxy", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µrefobj *πg.Object = πg.UnboundLocal; _ = µrefobj
				var µ_t *πg.Object = πg.UnboundLocal; _ = µ_t
				var µcallable *πg.Object = πg.UnboundLocal; _ = µcallable
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 πg.KWArgs
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 1275: refobj = _locate_object(_proxy_helper(obj))
					πF.SetLineno(1275)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_proxy_helper); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_locate_object); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µrefobj = πTemp004
					// line 1276: try:
					πF.SetLineno(1276)
					πF.PushCheckpoint(2)
					// line 1277: _t = "R2"
					πF.SetLineno(1277)
					µ_t = ßR2.ToObject()
					// line 1278: log.info("%s: %s" % (_t, obj))
					πF.SetLineno(1278)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_t, "_t"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(µ_t, µobj).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s: %s").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßReferenceError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 1279: except ReferenceError:
					πF.SetLineno(1279)
				Label3:
					// line 1280: _t = "R3"
					πF.SetLineno(1280)
					µ_t = ßR3.ToObject()
					// line 1281: log.info("%s: %s" % (_t, sys.exc_info()[1]))
					πF.SetLineno(1281)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_t, "_t"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(1).ToObject()
					if πTemp010, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp011.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(µ_t, πTemp009).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s: %s").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßCallableProxyType); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008 == πTemp004).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label4
					}
					goto Label5
					// line 1283: if type(obj) is CallableProxyType: callable = True
					πF.SetLineno(1283)
				Label4:
					// line 1283: if type(obj) is CallableProxyType: callable = True
					πF.SetLineno(1283)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µcallable = πTemp003
					goto Label6
				Label5:
					// line 1284: else: callable = False
					πF.SetLineno(1284)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µcallable = πTemp003
					goto Label6
				Label6:
					// line 1285: pickler.save_reduce(_create_weakproxy, (refobj, callable), obj=obj)
					πF.SetLineno(1285)
					πTemp001 = πF.MakeArgs(2)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_create_weakproxy); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µrefobj, "refobj"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcallable, "callable"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µrefobj, µcallable).ToObject()
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp012 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp012); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1286: log.info("# %s" % _t)
					πF.SetLineno(1286)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_t, "_t"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("# %s").ToObject(), µ_t); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1287: return
					πF.SetLineno(1287)
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
			if πE = πF.Globals().SetItem(πF, ßsave_weakproxy.ToObject(), πTemp089); πE != nil {
				continue
			}
			// line 1272: @register(ProxyType)
			πF.SetLineno(1272)
			πTemp001 = πF.MakeArgs(1)
			if πTemp090, πE = πg.ResolveGlobal(πF, ßsave_weakproxy); πE != nil {
				continue
			}
			πTemp001[0] = πTemp090
			πTemp017 = πF.MakeArgs(1)
			if πTemp090, πE = πg.ResolveGlobal(πF, ßCallableProxyType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp090
			if πTemp090, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp091, πE = πTemp090.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp090, πE = πTemp091.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_weakproxy.ToObject(), πTemp090); πE != nil {
				continue
			}
			// line 1272: @register(ProxyType)
			πF.SetLineno(1272)
			πTemp001 = πF.MakeArgs(1)
			if πTemp090, πE = πg.ResolveGlobal(πF, ßsave_weakproxy); πE != nil {
				continue
			}
			πTemp001[0] = πTemp090
			πTemp017 = πF.MakeArgs(1)
			if πTemp090, πE = πg.ResolveGlobal(πF, ßProxyType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp090
			if πTemp090, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp091, πE = πTemp090.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp090, πE = πTemp091.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_weakproxy.ToObject(), πTemp090); πE != nil {
				continue
			}
			// line 1290: def save_module(pickler, obj):
			πF.SetLineno(1290)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp090 = πg.NewFunction(πg.NewCode("save_module", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µchanged *πg.Object = πg.UnboundLocal; _ = µchanged
				var µnames *πg.Object = πg.UnboundLocal; _ = µnames
				var µbuiltin_mod *πg.Object = πg.UnboundLocal; _ = µbuiltin_mod
				var µ_main_dict *πg.Object = πg.UnboundLocal; _ = µ_main_dict
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
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πTemp007 *πg.Object
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
				var πTemp014 bool
				_ = πTemp014
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 1291: if False: #_use_diff:
					πF.SetLineno(1291)
				Label1:
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, πTemp003, ßdill.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 1292: if obj.__name__ != "dill":
					πF.SetLineno(1292)
				Label4:
					// line 1293: try:
					πF.SetLineno(1293)
					πF.PushCheckpoint(7)
					// line 1294: changed = diff.whats_changed(obj, seen=pickler._diff_cache)[0]
					πF.SetLineno(1294)
					πTemp001 = πg.NewInt(0).ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µpickler, ß_diff_cache, nil); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"seen", πTemp005},
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßdiff); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßwhats_changed, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp007.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					µchanged = πTemp003
					πF.PopCheckpoint()
					// line 1298: log.info("M1: %s with diff" % obj)
					πF.SetLineno(1298)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("M1: %s with diff").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1299: log.info("Diff: %s", changed.keys())
					πF.SetLineno(1299)
					πTemp004 = πF.MakeArgs(2)
					πTemp004[0] = πg.NewStr("Diff: %s").ToObject()
					if πE = πg.CheckLocal(πF, µchanged, "changed"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µchanged, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004[1] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1300: pickler.save_reduce(_import_module, (obj.__name__,), obj=obj,
					πF.SetLineno(1300)
					πTemp004 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_import_module); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple1(πTemp003).ToObject()
					πTemp004[1] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µchanged, "changed"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"obj", µobj},
						{"state", µchanged},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1302: log.info("# M1")
					πF.SetLineno(1302)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("# M1").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1303: return
					πF.SetLineno(1303)
					πR = πg.None
					continue
					goto Label6
				Label7:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 1295: except RuntimeError:  # not memorised module, probably part of dill
					πF.SetLineno(1295)
				Label8:
					// line 1296: pass
					πF.SetLineno(1296)
					πF.RestoreExc(nil, nil)
					goto Label6
				Label6:
					goto Label5
				Label5:
					// line 1305: log.info("M2: %s" % obj)
					πF.SetLineno(1305)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("M2: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1306: pickler.save_reduce(_import_module, (obj.__name__,), obj=obj)
					πF.SetLineno(1306)
					πTemp004 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_import_module); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple1(πTemp003).ToObject()
					πTemp004[1] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1307: log.info("# M2")
					πF.SetLineno(1307)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("# M2").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label3
				Label2:
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					πTemp004[1] = ß__file__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
						goto Label9
					}
					goto Label10
					// line 1311: if hasattr(obj, "__file__"):
					πF.SetLineno(1311)
				Label9:
					// line 1312: names = ["base_prefix", "base_exec_prefix", "exec_prefix",
					πF.SetLineno(1312)
					πTemp004 = make([]*πg.Object, 5)
					πTemp004[0] = ßbase_prefix.ToObject()
					πTemp004[1] = ßbase_exec_prefix.ToObject()
					πTemp004[2] = ßexec_prefix.ToObject()
					πTemp004[3] = ßprefix.ToObject()
					πTemp004[4] = ßreal_prefix.ToObject()
					πTemp001 = πg.NewList(πTemp004...).ToObject()
					µnames = πTemp001
					// line 1314: builtin_mod = any(obj.__file__.startswith(os.path.normpath(getattr(sys, name)))
					πF.SetLineno(1314)
					πTemp004 = πF.MakeArgs(1)
					πTemp010 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µname *πg.Object = πg.UnboundLocal; _ = µname
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
						var πTemp007 []*πg.Object
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
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µnames); πE != nil {
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
								πTemp005 = πF.MakeArgs(2)
								if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
									continue
								}
								πTemp005[0] = πTemp004
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								πTemp005[1] = µname
								if πTemp004, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								goto Label5
								// line 1314: builtin_mod = any(obj.__file__.startswith(os.path.normpath(getattr(sys, name)))
								πF.SetLineno(1314)
							Label4:
								// line 1314: builtin_mod = any(obj.__file__.startswith(os.path.normpath(getattr(sys, name)))
								πF.SetLineno(1314)
								πTemp005 = πF.MakeArgs(1)
								πTemp007 = πF.MakeArgs(1)
								πTemp008 = πF.MakeArgs(2)
								if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
									continue
								}
								πTemp008[0] = πTemp004
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								πTemp008[1] = µname
								if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp008)
								πTemp007[0] = πTemp006
								if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßpath, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßnormpath, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp005[0] = πTemp006
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µobj, ß__file__, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßstartswith, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(6)
								return πTemp004, nil
							Label6:
								πTemp006 = πSent
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
					πTemp004[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßany); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µbuiltin_mod = πTemp005
					// line 1316: builtin_mod = (builtin_mod or obj.__file__.endswith(EXTENSION_SUFFIXES) or
					πF.SetLineno(1316)
					if πE = πg.CheckLocal(πF, µbuiltin_mod, "builtin_mod"); πE != nil {
						continue
					}
					πTemp003 = µbuiltin_mod
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					πTemp004 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßEXTENSION_SUFFIXES); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobj, ß__file__, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobj, ß__file__, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Contains(πF, πTemp007, πg.NewStr("site-packages").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp011).ToObject()
					πTemp003 = πTemp005
				Label12:
					µbuiltin_mod = πTemp003
					goto Label11
				Label10:
					// line 1319: builtin_mod = True
					πF.SetLineno(1319)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µbuiltin_mod = πTemp003
					goto Label11
				Label11:
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					πTemp013 = πg.NewTuple2(ßbuiltins.ToObject(), ßdill.ToObject()).ToObject()
					if πTemp014, πE = πg.Contains(πF, πTemp013, πTemp012); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(!πTemp014).ToObject()
					πTemp005 = πTemp007
					if πTemp011, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if !πTemp011 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µbuiltin_mod, "builtin_mod"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.IsTrue(πF, µbuiltin_mod); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(!πTemp014).ToObject()
					πTemp005 = πTemp007
				Label14:
					πTemp003 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label13
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp004[0] = µpickler
					if πTemp007, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"child", πTemp007},
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßis_dill); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp007.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp005 = πTemp012
					if πTemp011, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if !πTemp011 {
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, µpickler, ß_main, nil); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(µobj == πTemp012).ToObject()
					πTemp005 = πTemp007
				Label15:
					πTemp003 = πTemp005
				Label13:
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label16
					}
					goto Label17
					// line 1320: if obj.__name__ not in ("builtins", "dill") \
					πF.SetLineno(1320)
				Label16:
					// line 1322: log.info("M1: %s" % obj)
					πF.SetLineno(1322)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("M1: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1323: _main_dict = obj.__dict__.copy() #XXX: better no copy? option to copy?
					πF.SetLineno(1323)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					µ_main_dict = πTemp003
					// line 1324: [_main_dict.pop(item, None) for item in singletontypes
					πF.SetLineno(1324)
					πTemp010 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µitem *πg.Object = πg.UnboundLocal; _ = µitem
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
								if πTemp003, πE = πg.ResolveGlobal(πF, ßsingletontypes); πE != nil {
									continue
								}
								πTemp004 = make([]*πg.Object, 2)
								πTemp004[0] = ß__builtins__.ToObject()
								πTemp004[1] = ß__loader__.ToObject()
								πTemp005 = πg.NewList(πTemp004...).ToObject()
								if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
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
									µitem = πTemp002
								}
								if πE != nil || !πTemp007 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 1324: [_main_dict.pop(item, None) for item in singletontypes
								πF.SetLineno(1324)
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
									continue
								}
								πTemp004[0] = µitem
								if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp004[1] = πTemp002
								if πE = πg.CheckLocal(πF, µ_main_dict, "_main_dict"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µ_main_dict, ßpop, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
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
					if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
						continue
					}
					// line 1326: pickler.save_reduce(_import_module, (obj.__name__,), obj=obj,
					πF.SetLineno(1326)
					πTemp004 = πF.MakeArgs(2)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_import_module); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple1(πTemp007).ToObject()
					πTemp004[1] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_main_dict, "_main_dict"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"obj", µobj},
						{"state", µ_main_dict},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1328: log.info("# M1")
					πF.SetLineno(1328)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("# M1").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label18
				Label17:
					// line 1330: log.info("M2: %s" % obj)
					πF.SetLineno(1330)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("M2: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1331: pickler.save_reduce(_import_module, (obj.__name__,), obj=obj)
					πF.SetLineno(1331)
					πTemp004 = πF.MakeArgs(2)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_import_module); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple1(πTemp007).ToObject()
					πTemp004[1] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1332: log.info("# M2")
					πF.SetLineno(1332)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("# M2").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label18
				Label18:
					// line 1333: return
					πF.SetLineno(1333)
					πR = πg.None
					continue
					goto Label3
				Label3:
					// line 1334: return
					πF.SetLineno(1334)
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
			if πE = πF.Globals().SetItem(πF, ßsave_module.ToObject(), πTemp090); πE != nil {
				continue
			}
			// line 1289: @register(ModuleType)
			πF.SetLineno(1289)
			πTemp001 = πF.MakeArgs(1)
			if πTemp091, πE = πg.ResolveGlobal(πF, ßsave_module); πE != nil {
				continue
			}
			πTemp001[0] = πTemp091
			πTemp017 = πF.MakeArgs(1)
			if πTemp091, πE = πg.ResolveGlobal(πF, ßModuleType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp091
			if πTemp091, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp092, πE = πTemp091.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp091, πE = πTemp092.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_module.ToObject(), πTemp091); πE != nil {
				continue
			}
			// line 1337: def save_type(pickler, obj):
			πF.SetLineno(1337)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp091 = πg.NewFunction(πg.NewCode("save_type", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µ_t *πg.Object = πg.UnboundLocal; _ = µ_t
				var µ_dict *πg.Object = πg.UnboundLocal; _ = µ_dict
				var µname *πg.Object = πg.UnboundLocal; _ = µname
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
				var πTemp006 []πg.Param
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 πg.KWArgs
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 16: goto Label16
					case 15: goto Label15
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_typemap); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Contains(πF, πTemp002, µobj); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					πTemp004[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001 = πTemp005
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label2
					}
					πTemp004 = πF.MakeArgs(1)
					πTemp006 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µattr *πg.Object = πg.UnboundLocal; _ = µattr
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
								πTemp002 = πg.NewTuple4(ß_fields.ToObject(), ß_asdict.ToObject(), ß_make.ToObject(), ß_replace.ToObject()).ToObject()
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
									µattr = πTemp002
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 1343: elif issubclass(obj, tuple) and all([hasattr(obj, attr) for attr in ('_fields','_asdict','_make','_replace')]):
								πF.SetLineno(1343)
								πTemp005 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πTemp005[0] = µobj
								if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
									continue
								}
								πTemp005[1] = µattr
								if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(4)
								return πTemp006, nil
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
					if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßall); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001 = πTemp007
				Label2:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß__module__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp002, ß__main__.ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001 = πg.GetBool(µobj == πTemp007).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					goto Label6
					// line 1339: if obj in _typemap:
					πF.SetLineno(1339)
				Label1:
					// line 1340: log.info("T1: %s" % obj)
					πF.SetLineno(1340)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("T1: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1341: pickler.save_reduce(_load_type, (_typemap[obj],), obj=obj)
					πF.SetLineno(1341)
					πTemp004 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_load_type); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002 = µobj
					if πTemp008, πE = πg.ResolveGlobal(πF, ß_typemap); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple1(πTemp007).ToObject()
					πTemp004[1] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1342: log.info("# T1")
					πF.SetLineno(1342)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("# T1").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label7
					// line 1343: elif issubclass(obj, tuple) and all([hasattr(obj, attr) for attr in ('_fields','_asdict','_make','_replace')]):
					πF.SetLineno(1343)
				Label3:
					// line 1345: log.info("T6: %s" % obj)
					πF.SetLineno(1345)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("T6: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1346: pickler.save_reduce(_create_namedtuple, (getattr(obj, "__qualname__", obj.__name__), obj._fields, obj.__module__), obj=obj)
					πF.SetLineno(1346)
					πTemp004 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_create_namedtuple); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					πTemp010 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp010[0] = µobj
					πTemp010[1] = ß__qualname__.ToObject()
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					πTemp010[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß_fields, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µobj, ß__module__, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(πTemp007, πTemp002, πTemp008).ToObject()
					πTemp004[1] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1347: log.info("# T6")
					πF.SetLineno(1347)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("# T6").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1348: return
					πF.SetLineno(1348)
					πR = πg.None
					continue
					goto Label7
					// line 1349: elif obj.__module__ == '__main__':
					πF.SetLineno(1349)
				Label4:
					πTemp004 = πF.MakeArgs(2)
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp010[0] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp004[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					πTemp004[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label8
					}
					goto Label9
					// line 1350: if issubclass(type(obj), type):
					πF.SetLineno(1350)
				Label8:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp004[0] = µpickler
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"child", πTemp002},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßis_dill); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp004, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001 = πTemp007
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µpickler, ß_byref, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp011).ToObject()
					πTemp001 = πTemp002
				Label11:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label12
					}
					goto Label13
					// line 1352: if is_dill(pickler, child=True) and not pickler._byref:
					πF.SetLineno(1352)
				Label12:
					// line 1354: _t = 'T2'
					πF.SetLineno(1354)
					µ_t = ßT2.ToObject()
					// line 1355: log.info("%s: %s" % (_t, obj))
					πF.SetLineno(1355)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_t, "_t"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µ_t, µobj).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s: %s").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1356: _dict = _dict_from_dictproxy(obj.__dict__)
					πF.SetLineno(1356)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobj, ß__dict__, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_dict_from_dictproxy); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µ_dict = πTemp002
					goto Label14
				Label13:
					// line 1359: log.info("T5: %s" % obj)
					πF.SetLineno(1359)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("T5: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1360: name = getattr(obj, '__qualname__', getattr(obj, '__name__', None))
					πF.SetLineno(1360)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					πTemp004[1] = ß__qualname__.ToObject()
					πTemp010 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp010[0] = µobj
					πTemp010[1] = ß__name__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp010[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp004[2] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µname = πTemp002
					// line 1361: StockPickler.save_global(pickler, obj, name=name)
					πF.SetLineno(1361)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp004[0] = µpickler
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[1] = µobj
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"name", µname},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsave_global, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1362: log.info("# T5")
					πF.SetLineno(1362)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("# T5").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1363: return
					πF.SetLineno(1363)
					πR = πg.None
					continue
					goto Label14
				Label14:
					goto Label10
				Label9:
					// line 1365: _t = 'T3'
					πF.SetLineno(1365)
					µ_t = ßT3.ToObject()
					// line 1366: log.info("%s: %s" % (_t, obj))
					πF.SetLineno(1366)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_t, "_t"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µ_t, µobj).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s: %s").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1367: _dict = obj.__dict__
					πF.SetLineno(1367)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µobj, ß__dict__, nil); πE != nil {
						continue
					}
					µ_dict = πTemp001
					goto Label10
				Label10:
					πTemp004 = πF.MakeArgs(2)
					πTemp004[0] = ß__slots__.ToObject()
					πTemp010 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp010...).ToObject()
					πTemp004[1] = πTemp002
					if πE = πg.CheckLocal(πF, µ_dict, "_dict"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µ_dict, ßget, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Iter(πF, πTemp007); πE != nil {
						continue
					}
					πF.PushCheckpoint(16)
					πTemp003 = false
				Label15:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label17
					}
					if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
						µname = πTemp002
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(15)            
					// line 1372: del _dict[name]
					πF.SetLineno(1372)
					if πE = πg.CheckLocal(πF, µ_dict, "_dict"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002 = µname
					if πE = πg.DelItem(πF, µ_dict, πTemp002); πE != nil {
						continue
					}
					continue
				Label16:
					if πE != nil || πR != nil {
						continue
					}
				Label17:
					// line 1373: pickler.save_reduce(_create_type, (type(obj), obj.__name__,
					πF.SetLineno(1373)
					πTemp004 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_create_type); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp010[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µobj, ß__bases__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_dict, "_dict"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple4(πTemp007, πTemp002, πTemp008, µ_dict).ToObject()
					πTemp004[1] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1375: log.info("# %s" % _t)
					πF.SetLineno(1375)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_t, "_t"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("# %s").ToObject(), µ_t); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label7
					// line 1377: elif obj is type(None):
					πF.SetLineno(1377)
				Label5:
					// line 1378: log.info("T7: %s" % obj)
					πF.SetLineno(1378)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("T7: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label18
					}
					goto Label19
					// line 1379: if PY3:
					πF.SetLineno(1379)
				Label18:
					// line 1380: pickler.write(bytes('c__builtin__\nNoneType\n', 'UTF-8'))
					πF.SetLineno(1380)
					πTemp004 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(2)
					πTemp010[0] = πg.NewStr("c__builtin__\nNoneType\n").ToObject()
					πTemp010[1] = πg.NewStr("UTF-8").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbytes); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp004[0] = πTemp002
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label20
				Label19:
					// line 1382: pickler.write('c__builtin__\nNoneType\n')
					πF.SetLineno(1382)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("c__builtin__\nNoneType\n").ToObject()
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label20
				Label20:
					// line 1383: log.info("# T7")
					πF.SetLineno(1383)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("# T7").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label7
				Label6:
					// line 1385: log.info("T4: %s" % obj)
					πF.SetLineno(1385)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("T4: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1389: name = getattr(obj, '__qualname__', getattr(obj, '__name__', None))
					πF.SetLineno(1389)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					πTemp004[1] = ß__qualname__.ToObject()
					πTemp010 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp010[0] = µobj
					πTemp010[1] = ß__name__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp010[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp004[2] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µname = πTemp002
					// line 1390: StockPickler.save_global(pickler, obj, name=name)
					πF.SetLineno(1390)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp004[0] = µpickler
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[1] = µobj
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"name", µname},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsave_global, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 1391: log.info("# T4")
					πF.SetLineno(1391)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("# T4").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label7
				Label7:
					// line 1392: return
					πF.SetLineno(1392)
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
			if πE = πF.Globals().SetItem(πF, ßsave_type.ToObject(), πTemp091); πE != nil {
				continue
			}
			// line 1336: @register(TypeType)
			πF.SetLineno(1336)
			πTemp001 = πF.MakeArgs(1)
			if πTemp092, πE = πg.ResolveGlobal(πF, ßsave_type); πE != nil {
				continue
			}
			πTemp001[0] = πTemp092
			πTemp017 = πF.MakeArgs(1)
			if πTemp092, πE = πg.ResolveGlobal(πF, ßTypeType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp092
			if πTemp092, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp093, πE = πTemp092.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp092, πE = πTemp093.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_type.ToObject(), πTemp092); πE != nil {
				continue
			}
			// line 1395: def save_property(pickler, obj):
			πF.SetLineno(1395)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp092 = πg.NewFunction(πg.NewCode("save_property", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var πTemp001 []*πg.Object
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
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1396: log.info("Pr: %s" % obj)
					πF.SetLineno(1396)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Pr: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1397: pickler.save_reduce(property, (obj.fget, obj.fset, obj.fdel, obj.__doc__),
					πF.SetLineno(1397)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßproperty); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßfget, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ßfset, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µobj, ßfdel, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µobj, ß__doc__, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple4(πTemp003, πTemp004, πTemp005, πTemp006).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1399: log.info("# Pr")
					πF.SetLineno(1399)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Pr").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsave_property.ToObject(), πTemp092); πE != nil {
				continue
			}
			// line 1394: @register(property)
			πF.SetLineno(1394)
			πTemp001 = πF.MakeArgs(1)
			if πTemp093, πE = πg.ResolveGlobal(πF, ßsave_property); πE != nil {
				continue
			}
			πTemp001[0] = πTemp093
			πTemp017 = πF.MakeArgs(1)
			if πTemp093, πE = πg.ResolveGlobal(πF, ßproperty); πE != nil {
				continue
			}
			πTemp017[0] = πTemp093
			if πTemp093, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp094, πE = πTemp093.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp093, πE = πTemp094.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_property.ToObject(), πTemp093); πE != nil {
				continue
			}
			// line 1403: def save_classmethod(pickler, obj):
			πF.SetLineno(1403)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp093 = πg.NewFunction(πg.NewCode("save_classmethod", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µim_func *πg.Object = πg.UnboundLocal; _ = µim_func
				var µorig_func *πg.Object = πg.UnboundLocal; _ = µorig_func
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
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
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 1404: log.info("Cm: %s" % obj)
					πF.SetLineno(1404)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Cm: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1405: im_func = '__func__' if PY3 else 'im_func'
					πF.SetLineno(1405)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					πTemp002 = ß__func__.ToObject()
					goto Label2
				Label1:
					πTemp002 = ßim_func.ToObject()
				Label2:
					µim_func = πTemp002
					// line 1406: try:
					πF.SetLineno(1406)
					πF.PushCheckpoint(4)
					// line 1407: orig_func = getattr(obj, im_func)
					πF.SetLineno(1407)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πE = πg.CheckLocal(πF, µim_func, "im_func"); πE != nil {
						continue
					}
					πTemp001[1] = µim_func
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µorig_func = πTemp003
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 1408: except AttributeError:  # Python 2.6
					πF.SetLineno(1408)
				Label5:
					// line 1409: orig_func = obj.__get__(None, object)
					πF.SetLineno(1409)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µobj, ß__get__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µorig_func = πTemp003
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßclassmethod); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					goto Label7
					// line 1410: if isinstance(obj, classmethod):
					πF.SetLineno(1410)
				Label6:
					// line 1411: orig_func = getattr(orig_func, im_func) # Unbind
					πF.SetLineno(1411)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µorig_func, "orig_func"); πE != nil {
						continue
					}
					πTemp001[0] = µorig_func
					if πE = πg.CheckLocal(πF, µim_func, "im_func"); πE != nil {
						continue
					}
					πTemp001[1] = µim_func
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µorig_func = πTemp003
					goto Label7
				Label7:
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					// line 1412: pickler.save_reduce(type(obj), (orig_func,), obj=obj)
					πF.SetLineno(1412)
					πTemp001 = πF.MakeArgs(2)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µorig_func, "orig_func"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(µorig_func).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp008 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp008); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1413: log.info("# Cm")
					πF.SetLineno(1413)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("# Cm").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsave_classmethod.ToObject(), πTemp093); πE != nil {
				continue
			}
			// line 1401: @register(staticmethod)
			πF.SetLineno(1401)
			πTemp001 = πF.MakeArgs(1)
			if πTemp094, πE = πg.ResolveGlobal(πF, ßsave_classmethod); πE != nil {
				continue
			}
			πTemp001[0] = πTemp094
			πTemp017 = πF.MakeArgs(1)
			if πTemp094, πE = πg.ResolveGlobal(πF, ßclassmethod); πE != nil {
				continue
			}
			πTemp017[0] = πTemp094
			if πTemp094, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp095, πE = πTemp094.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp094, πE = πTemp095.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_classmethod.ToObject(), πTemp094); πE != nil {
				continue
			}
			// line 1401: @register(staticmethod)
			πF.SetLineno(1401)
			πTemp001 = πF.MakeArgs(1)
			if πTemp094, πE = πg.ResolveGlobal(πF, ßsave_classmethod); πE != nil {
				continue
			}
			πTemp001[0] = πTemp094
			πTemp017 = πF.MakeArgs(1)
			if πTemp094, πE = πg.ResolveGlobal(πF, ßstaticmethod); πE != nil {
				continue
			}
			πTemp017[0] = πTemp094
			if πTemp094, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp095, πE = πTemp094.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp094, πE = πTemp095.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_classmethod.ToObject(), πTemp094); πE != nil {
				continue
			}
			// line 1416: def save_function(pickler, obj):
			πF.SetLineno(1416)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			πTemp004[1] = πg.Param{Name: "obj", Def: nil}
			πTemp094 = πg.NewFunction(πg.NewCode("save_function", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µobj *πg.Object = πArgs[1]; _ = µobj
				var µglobalvars *πg.Object = πg.UnboundLocal; _ = µglobalvars
				var µglobs *πg.Object = πg.UnboundLocal; _ = µglobs
				var µ_byref *πg.Object = πg.UnboundLocal; _ = µ_byref
				var µ_recurse *πg.Object = πg.UnboundLocal; _ = µ_recurse
				var µ_memo *πg.Object = πg.UnboundLocal; _ = µ_memo
				var µ_super *πg.Object = πg.UnboundLocal; _ = µ_super
				var µfkwdefaults *πg.Object = πg.UnboundLocal; _ = µfkwdefaults
				var µname *πg.Object = πg.UnboundLocal; _ = µname
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
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 []*πg.Object
				_ = πTemp011
				var πTemp012 bool
				_ = πTemp012
				var πTemp013 bool
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_locate_function); πE != nil {
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
					// line 1417: if not _locate_function(obj): #, pickler._session):
					πF.SetLineno(1417)
				Label1:
					// line 1418: log.info("F1: %s" % obj)
					πF.SetLineno(1418)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("F1: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp002[0] = µpickler
					πTemp002[1] = ß_recurse.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label4
					}
					goto Label5
					// line 1419: if getattr(pickler, '_recurse', False):
					πF.SetLineno(1419)
				Label4:
					// line 1421: from .detect import globalvars
					πF.SetLineno(1421)
					if πTemp002, πE = πg.ImportModule(πF, ".detect"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßglobalvars); πE != nil {
						continue
					}
					µglobalvars = πTemp003
					// line 1422: globs = globalvars(obj, recurse=True, builtin=True)
					πF.SetLineno(1422)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"recurse", πTemp001},
						{"builtin", πTemp003},
					}
					if πE = πg.CheckLocal(πF, µglobalvars, "globalvars"); πE != nil {
						continue
					}
					if πTemp001, πE = µglobalvars.Call(πF, πTemp002, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µglobs = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstack); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label7
					}
					goto Label8
					// line 1429: if id(obj) in stack: # or obj in globs.values():
					πF.SetLineno(1429)
				Label7:
					// line 1430: globs = obj.__globals__ if PY3 else obj.func_globals
					πF.SetLineno(1430)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__globals__, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					goto Label10
				Label9:
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßfunc_globals, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label10:
					µglobs = πTemp001
					goto Label8
				Label8:
					goto Label6
				Label5:
					// line 1432: globs = obj.__globals__ if PY3 else obj.func_globals
					πF.SetLineno(1432)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__globals__, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					goto Label12
				Label11:
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßfunc_globals, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label12:
					µglobs = πTemp001
					goto Label6
				Label6:
					// line 1433: _byref = getattr(pickler, '_byref', None)
					πF.SetLineno(1433)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp002[0] = µpickler
					πTemp002[1] = ß_byref.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µ_byref = πTemp003
					// line 1434: _recurse = getattr(pickler, '_recurse', None)
					πF.SetLineno(1434)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp002[0] = µpickler
					πTemp002[1] = ß_recurse.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µ_recurse = πTemp003
					// line 1435: _memo = (id(obj) in stack) and (_recurse is not None)
					πF.SetLineno(1435)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstack); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008).ToObject()
					πTemp001 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µ_recurse, "_recurse"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µ_recurse != πTemp004).ToObject()
					πTemp001 = πTemp003
				Label13:
					µ_memo = πTemp001
					// line 1437: stack[id(obj)] = len(stack), obj
					πF.SetLineno(1437)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstack); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp004, µobj).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstack); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp009, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp007 = πTemp010
					if πE = πg.SetItem(πF, πTemp004, πTemp007, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label14
					}
					goto Label15
					// line 1438: if PY3:
					πF.SetLineno(1438)
				Label14:
					// line 1440: _super = ('super' in getattr(obj.__code__,'co_names',())) and (_byref is not None)
					πF.SetLineno(1440)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ß__code__, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					πTemp002[1] = ßco_names.ToObject()
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp002[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp008, πE = πg.Contains(πF, πTemp007, ßsuper.ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008).ToObject()
					πTemp001 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label17
					}
					if πE = πg.CheckLocal(πF, µ_byref, "_byref"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µ_byref != πTemp004).ToObject()
					πTemp001 = πTemp003
				Label17:
					µ_super = πTemp001
					if πE = πg.CheckLocal(πF, µ_super, "_super"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µ_super); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label18
					}
					goto Label19
					// line 1441: if _super: pickler._byref = True
					πF.SetLineno(1441)
				Label18:
					// line 1441: if _super: pickler._byref = True
					πF.SetLineno(1441)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_byref, πTemp003); πE != nil {
						continue
					}
					goto Label19
				Label19:
					if πE = πg.CheckLocal(πF, µ_memo, "_memo"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µ_memo); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label20
					}
					goto Label21
					// line 1442: if _memo: pickler._recurse = False
					πF.SetLineno(1442)
				Label20:
					// line 1442: if _memo: pickler._recurse = False
					πF.SetLineno(1442)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_recurse, πTemp003); πE != nil {
						continue
					}
					goto Label21
				Label21:
					// line 1443: fkwdefaults = getattr(obj, '__kwdefaults__', None)
					πF.SetLineno(1443)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					πTemp002[1] = ß__kwdefaults__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfkwdefaults = πTemp003
					// line 1444: pickler.save_reduce(_create_function, (obj.__code__,
					πF.SetLineno(1444)
					πTemp002 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_create_function); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					πTemp011 = make([]*πg.Object, 7)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__code__, nil); πE != nil {
						continue
					}
					πTemp011[0] = πTemp003
					if πE = πg.CheckLocal(πF, µglobs, "globs"); πE != nil {
						continue
					}
					πTemp011[1] = µglobs
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
						continue
					}
					πTemp011[2] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__defaults__, nil); πE != nil {
						continue
					}
					πTemp011[3] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__closure__, nil); πE != nil {
						continue
					}
					πTemp011[4] = πTemp003
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ß__dict__, nil); πE != nil {
						continue
					}
					πTemp011[5] = πTemp003
					if πE = πg.CheckLocal(πF, µfkwdefaults, "fkwdefaults"); πE != nil {
						continue
					}
					πTemp011[6] = µfkwdefaults
					πTemp001 = πg.NewTuple(πTemp011...).ToObject()
					πTemp002[1] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label16
				Label15:
					// line 1449: _super = ('super' in getattr(obj.func_code,'co_names',())) and (_byref is not None) and getattr(pickler, '_recurse', False)
					πF.SetLineno(1449)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ßfunc_code, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					πTemp002[1] = ßco_names.ToObject()
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp002[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp008, πE = πg.Contains(πF, πTemp007, ßsuper.ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008).ToObject()
					πTemp001 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label22
					}
					if πE = πg.CheckLocal(πF, µ_byref, "_byref"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µ_byref != πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label22
					}
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp002[0] = µpickler
					πTemp002[1] = ß_recurse.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πTemp004
				Label22:
					µ_super = πTemp001
					if πE = πg.CheckLocal(πF, µ_super, "_super"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µ_super); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label23
					}
					goto Label24
					// line 1450: if _super: pickler._byref = True
					πF.SetLineno(1450)
				Label23:
					// line 1450: if _super: pickler._byref = True
					πF.SetLineno(1450)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_byref, πTemp003); πE != nil {
						continue
					}
					goto Label24
				Label24:
					if πE = πg.CheckLocal(πF, µ_memo, "_memo"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µ_memo); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label25
					}
					goto Label26
					// line 1451: if _memo: pickler._recurse = False
					πF.SetLineno(1451)
				Label25:
					// line 1451: if _memo: pickler._recurse = False
					πF.SetLineno(1451)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_recurse, πTemp003); πE != nil {
						continue
					}
					goto Label26
				Label26:
					// line 1452: pickler.save_reduce(_create_function, (obj.func_code,
					πF.SetLineno(1452)
					πTemp002 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_create_function); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßfunc_code, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µglobs, "globs"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µobj, ßfunc_name, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µobj, ßfunc_defaults, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µobj, ßfunc_closure, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, µobj, ß__dict__, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple6(πTemp003, µglobs, πTemp004, πTemp007, πTemp009, πTemp010).ToObject()
					πTemp002[1] = πTemp001
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"obj", µobj},
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßsave_reduce, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label16
				Label16:
					if πE = πg.CheckLocal(πF, µ_super, "_super"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µ_super); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label27
					}
					goto Label28
					// line 1456: if _super: pickler._byref = _byref
					πF.SetLineno(1456)
				Label27:
					// line 1456: if _super: pickler._byref = _byref
					πF.SetLineno(1456)
					if πE = πg.CheckLocal(πF, µ_byref, "_byref"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µ_byref); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_byref, πTemp001); πE != nil {
						continue
					}
					goto Label28
				Label28:
					if πE = πg.CheckLocal(πF, µ_memo, "_memo"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µ_memo); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label29
					}
					goto Label30
					// line 1457: if _memo: pickler._recurse = _recurse
					πF.SetLineno(1457)
				Label29:
					// line 1457: if _memo: pickler._recurse = _recurse
					πF.SetLineno(1457)
					if πE = πg.CheckLocal(πF, µ_recurse, "_recurse"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µ_recurse); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpickler, ß_recurse, πTemp001); πE != nil {
						continue
					}
					goto Label30
				Label30:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßOLDER); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label31
					}
					if πE = πg.CheckLocal(πF, µ_byref, "_byref"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µ_byref); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp008).ToObject()
					πTemp001 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label31
					}
					if πE = πg.CheckLocal(πF, µ_super, "_super"); πE != nil {
						continue
					}
					πTemp003 = µ_super
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label32
					}
					if πE = πg.CheckLocal(πF, µ_super, "_super"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.IsTrue(πF, µ_super); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(!πTemp013).ToObject()
					πTemp004 = πTemp007
					if πTemp012, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp012 {
						goto Label33
					}
					if πE = πg.CheckLocal(πF, µ_memo, "_memo"); πE != nil {
						continue
					}
					πTemp004 = µ_memo
				Label33:
					πTemp003 = πTemp004
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label32
					}
					if πE = πg.CheckLocal(πF, µ_super, "_super"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.IsTrue(πF, µ_super); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(!πTemp013).ToObject()
					πTemp004 = πTemp007
					if πTemp012, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp012 {
						goto Label34
					}
					if πE = πg.CheckLocal(πF, µ_memo, "_memo"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.IsTrue(πF, µ_memo); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(!πTemp013).ToObject()
					πTemp004 = πTemp007
					if πTemp012, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp012 {
						goto Label34
					}
					if πE = πg.CheckLocal(πF, µ_recurse, "_recurse"); πE != nil {
						continue
					}
					πTemp004 = µ_recurse
				Label34:
					πTemp003 = πTemp004
				Label32:
					πTemp001 = πTemp003
				Label31:
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label35
					}
					goto Label36
					// line 1461: if OLDER and not _byref and (_super or (not _super and _memo) or (not _super and not _memo and _recurse)): pickler.clear_memo()
					πF.SetLineno(1461)
				Label35:
					// line 1461: if OLDER and not _byref and (_super or (not _super and _memo) or (not _super and not _memo and _recurse)): pickler.clear_memo()
					πF.SetLineno(1461)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpickler, ßclear_memo, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label36
				Label36:
					// line 1466: log.info("# F1")
					πF.SetLineno(1466)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("# F1").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label3
				Label2:
					// line 1468: log.info("F2: %s" % obj)
					πF.SetLineno(1468)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("F2: %s").ToObject(), µobj); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 1469: name = getattr(obj, '__qualname__', getattr(obj, '__name__', None))
					πF.SetLineno(1469)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					πTemp002[1] = ß__qualname__.ToObject()
					πTemp011 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp011[0] = µobj
					πTemp011[1] = ß__name__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp011[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp002[2] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µname = πTemp003
					// line 1470: StockPickler.save_global(pickler, obj, name=name)
					πF.SetLineno(1470)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					πTemp002[0] = µpickler
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[1] = µobj
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"name", µname},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßsave_global, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 1471: log.info("# F2")
					πF.SetLineno(1471)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("# F2").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label3
				Label3:
					// line 1472: return
					πF.SetLineno(1472)
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
			if πE = πF.Globals().SetItem(πF, ßsave_function.ToObject(), πTemp094); πE != nil {
				continue
			}
			// line 1415: @register(FunctionType)
			πF.SetLineno(1415)
			πTemp001 = πF.MakeArgs(1)
			if πTemp095, πE = πg.ResolveGlobal(πF, ßsave_function); πE != nil {
				continue
			}
			πTemp001[0] = πTemp095
			πTemp017 = πF.MakeArgs(1)
			if πTemp095, πE = πg.ResolveGlobal(πF, ßFunctionType); πE != nil {
				continue
			}
			πTemp017[0] = πTemp095
			if πTemp095, πE = πg.ResolveGlobal(πF, ßregister); πE != nil {
				continue
			}
			if πTemp096, πE = πTemp095.Call(πF, πTemp017, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp017)
			if πTemp095, πE = πTemp096.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßsave_function.ToObject(), πTemp095); πE != nil {
				continue
			}
			// line 1475: def pickles(obj,exact=False,safe=False,**kwds):
			πF.SetLineno(1475)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp096, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "exact", Def: πTemp096}
			if πTemp096, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "safe", Def: πTemp096}
			πTemp095 = πg.NewFunction(πg.NewCode("pickles", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µexact *πg.Object = πArgs[1]; _ = µexact
				var µsafe *πg.Object = πArgs[2]; _ = µsafe
				var µkwds *πg.Object = πArgs[3]; _ = µkwds
				var µexceptions *πg.Object = πg.UnboundLocal; _ = µexceptions
				var µpik *πg.Object = πg.UnboundLocal; _ = µpik
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var πTemp001 bool
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5: goto Label5
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 1476: """quick check if object pickles with dill"""
					πF.SetLineno(1476)
					if πE = πg.CheckLocal(πF, µsafe, "safe"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µsafe); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 1477: if safe: exceptions = (Exception,) # RuntimeError, ValueError
					πF.SetLineno(1477)
				Label1:
					// line 1477: if safe: exceptions = (Exception,) # RuntimeError, ValueError
					πF.SetLineno(1477)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(πTemp003).ToObject()
					µexceptions = πTemp002
					goto Label3
				Label2:
					// line 1479: exceptions = (TypeError, AssertionError, PicklingError, UnpicklingError)
					πF.SetLineno(1479)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßUnpicklingError); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple4(πTemp003, πTemp004, πTemp005, πTemp006).ToObject()
					µexceptions = πTemp002
					goto Label3
				Label3:
					// line 1480: try:
					πF.SetLineno(1480)
					πF.PushCheckpoint(5)
					// line 1481: pik = copy(obj, **kwds)
					πF.SetLineno(1481)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007[0] = µobj
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcopy); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp007, nil, nil, µkwds); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µpik = πTemp003
					// line 1482: try:
					πF.SetLineno(1482)
					πF.PushCheckpoint(7)
					// line 1483: result = bool(pik.all() == obj.all())
					πF.SetLineno(1483)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpik, "pik"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpik, ßall, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µobj, ßall, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					πTemp007[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µresult = πTemp003
					πF.PopCheckpoint()
					goto Label6
				Label7:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label8
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 1484: except AttributeError:
					πF.SetLineno(1484)
				Label8:
					// line 1485: result = pik == obj
					πF.SetLineno(1485)
					if πE = πg.CheckLocal(πF, µpik, "pik"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µpik, µobj); πE != nil {
						continue
					}
					µresult = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µresult); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label9
					}
					goto Label10
					// line 1486: if result: return True
					πF.SetLineno(1486)
				Label9:
					// line 1486: if result: return True
					πF.SetLineno(1486)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µexact, "exact"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µexact); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp001).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label11
					}
					goto Label12
					// line 1487: if not exact:
					πF.SetLineno(1487)
				Label11:
					// line 1488: result = type(pik) == type(obj)
					πF.SetLineno(1488)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpik, "pik"); πE != nil {
						continue
					}
					πTemp007[0] = µpik
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp007[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp002, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					µresult = πTemp002
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µresult); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label13
					}
					goto Label14
					// line 1489: if result: return result
					πF.SetLineno(1489)
				Label13:
					// line 1489: if result: return result
					πF.SetLineno(1489)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
					goto Label14
				Label14:
					// line 1491: return repr(type(pik)) == repr(type(obj)) #XXX: InstanceType?
					πF.SetLineno(1491)
					πTemp007 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpik, "pik"); πE != nil {
						continue
					}
					πTemp010[0] = µpik
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp007[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp007 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp010[0] = µobj
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp007[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp002, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label12
				Label12:
					// line 1492: return False
					πF.SetLineno(1492)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πE = πg.CheckLocal(πF, µexceptions, "exceptions"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsInstance(πF, πTemp008.ToObject(), µexceptions); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label15
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 1493: except exceptions:
					πF.SetLineno(1493)
				Label15:
					// line 1494: return False
					πF.SetLineno(1494)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpickles.ToObject(), πTemp095); πE != nil {
				continue
			}
			// line 1476: """quick check if object pickles with dill"""
			πF.SetLineno(1476)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp096}, πg.NewStr("quick check if object pickles with dill").ToObject()); πE != nil {
				continue
			}
			if πTemp097, πE = πg.ResolveGlobal(πF, ßpickles); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp097, ß__doc__, πTemp096); πE != nil {
				continue
			}
			// line 1496: def check(obj, *args, **kwds):
			πF.SetLineno(1496)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp096 = πg.NewFunction(πg.NewCode("check", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µargs *πg.Object = πArgs[1]; _ = µargs
				var µkwds *πg.Object = πArgs[2]; _ = µkwds
				var µverbose *πg.Object = πg.UnboundLocal; _ = µverbose
				var µpython *πg.Object = πg.UnboundLocal; _ = µpython
				var µsys *πg.Object = πg.UnboundLocal; _ = µsys
				var µsubprocess *πg.Object = πg.UnboundLocal; _ = µsubprocess
				var µfail *πg.Object = πg.UnboundLocal; _ = µfail
				var µ_obj *πg.Object = πg.UnboundLocal; _ = µ_obj
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					default: panic("unexpected function state")
					}
					// line 1497: """check pickling of an object across another process"""
					πF.SetLineno(1497)
					// line 1502: verbose = kwds.pop('verbose', False)
					πF.SetLineno(1502)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßverbose.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
					µverbose = πTemp003
					// line 1503: python = kwds.pop('python', None)
					πF.SetLineno(1503)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßpython.ToObject()
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
					µpython = πTemp003
					if πE = πg.CheckLocal(πF, µpython, "python"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µpython == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 1504: if python is None:
					πF.SetLineno(1504)
				Label1:
					// line 1505: import sys
					πF.SetLineno(1505)
					if πTemp001, πE = πg.ImportModule(πF, "sys"); πE != nil {
						continue
					}
					πTemp002 = πTemp001[0]
					µsys = πTemp002
					// line 1506: python = sys.executable
					πF.SetLineno(1506)
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsys, ßexecutable, nil); πE != nil {
						continue
					}
					µpython = πTemp002
					goto Label2
				Label2:
					// line 1508: isinstance(python, str)
					πF.SetLineno(1508)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpython, "python"); πE != nil {
						continue
					}
					πTemp001[0] = µpython
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1509: import subprocess
					πF.SetLineno(1509)
					if πTemp001, πE = πg.ImportModule(πF, "subprocess"); πE != nil {
						continue
					}
					πTemp002 = πTemp001[0]
					µsubprocess = πTemp002
					// line 1510: fail = True
					πF.SetLineno(1510)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µfail = πTemp002
					// line 1511: try:
					πF.SetLineno(1511)
					πF.PushCheckpoint(3)
					// line 1512: _obj = dumps(obj, *args, **kwds)
					πF.SetLineno(1512)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdumps); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, µkwds); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µ_obj = πTemp003
					// line 1513: fail = False
					πF.SetLineno(1513)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µfail = πTemp002
					πF.PopCheckpoint()
					goto Label3
				Label3:
					πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
					if πE = πg.CheckLocal(πF, µfail, "fail"); πE != nil {
						continue
					}
					πTemp002 = µfail
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µverbose, "verbose"); πE != nil {
						continue
					}
					πTemp002 = µverbose
				Label4:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 1515: if fail and verbose:
					πF.SetLineno(1515)
				Label5:
					// line 1516: print("DUMP FAILED")
					πF.SetLineno(1516)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = πg.NewStr("DUMP FAILED").ToObject()
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					goto Label6
				Label6:
					if πTemp005 != nil {
						πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
						continue
					}
					if πR != nil {
						continue
					}
					// line 1517: msg = "%s -c import dill; print(dill.loads(%s))" % (python, repr(_obj))
					πF.SetLineno(1517)
					if πE = πg.CheckLocal(πF, µpython, "python"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_obj, "_obj"); πE != nil {
						continue
					}
					πTemp001[0] = µ_obj
					if πTemp007, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp003 = πg.NewTuple2(µpython, πTemp008).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s -c import dill; print(dill.loads(%s))").ToObject(), πTemp003); πE != nil {
						continue
					}
					µmsg = πTemp002
					// line 1518: msg = "SUCCESS" if not subprocess.call(msg.split(None,2)) else "LOAD FAILED"
					πF.SetLineno(1518)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(2)
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp009[0] = πTemp007
					πTemp009[1] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µmsg, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µsubprocess, "subprocess"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µsubprocess, ßcall, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label7
					}
					πTemp002 = ßSUCCESS.ToObject()
					goto Label8
				Label7:
					πTemp002 = πg.NewStr("LOAD FAILED").ToObject()
				Label8:
					µmsg = πTemp002
					if πE = πg.CheckLocal(πF, µverbose, "verbose"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µverbose); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 1519: if verbose:
					πF.SetLineno(1519)
				Label9:
					// line 1520: print(msg)
					πF.SetLineno(1520)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[0] = µmsg
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					goto Label10
				Label10:
					// line 1521: return
					πF.SetLineno(1521)
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
			if πE = πF.Globals().SetItem(πF, ßcheck.ToObject(), πTemp096); πE != nil {
				continue
			}
			// line 1497: """check pickling of an object across another process"""
			πF.SetLineno(1497)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp097}, πg.NewStr("check pickling of an object across another process").ToObject()); πE != nil {
				continue
			}
			if πTemp098, πE = πg.ResolveGlobal(πF, ßcheck); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp098, ß__doc__, πTemp097); πE != nil {
				continue
			}
			// line 1524: def is_dill(pickler, child=None):
			πF.SetLineno(1524)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pickler", Def: nil}
			if πTemp098, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "child", Def: πTemp098}
			πTemp097 = πg.NewFunction(πg.NewCode("is_dill", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpickler *πg.Object = πArgs[0]; _ = µpickler
				var µchild *πg.Object = πArgs[1]; _ = µchild
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
					// line 1525: "check the dill-ness of your pickler"
					πF.SetLineno(1525)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µchild == πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPY34); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µpickler, ß__class__, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					πTemp005[1] = ßmro.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 1526: if (child is False) or PY34 or (not hasattr(pickler.__class__, 'mro')):
					πF.SetLineno(1526)
				Label2:
					// line 1527: return 'dill' in pickler.__module__
					πF.SetLineno(1527)
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpickler, ß__module__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp003, ßdill.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					πR = πTemp001
					continue
					goto Label3
				Label3:
					// line 1528: return Pickler in pickler.__class__.mro()
					πF.SetLineno(1528)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpickler, "pickler"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µpickler, ß__class__, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßmro, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßis_dill.ToObject(), πTemp097); πE != nil {
				continue
			}
			// line 1525: "check the dill-ness of your pickler"
			πF.SetLineno(1525)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp098}, πg.NewStr("check the dill-ness of your pickler").ToObject()); πE != nil {
				continue
			}
			if πTemp099, πE = πg.ResolveGlobal(πF, ßis_dill); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp099, ß__doc__, πTemp098); πE != nil {
				continue
			}
			// line 1530: def _extend():
			πF.SetLineno(1530)
			πTemp004 = make([]πg.Param, 0)
			πTemp098 = πg.NewFunction(πg.NewCode("_extend", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_dill.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µt *πg.Object = πg.UnboundLocal; _ = µt
				var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
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
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 1531: """extend pickle with all of dill's registered types"""
					πF.SetLineno(1531)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdispatch, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßitems, nil); πE != nil {
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
						µt = πTemp003
						µfunc = πTemp006
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 1534: try:
					πF.SetLineno(1534)
					πF.PushCheckpoint(5)
					// line 1535: StockPickler.dispatch[t] = func
					πF.SetLineno(1535)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µfunc); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßStockPickler); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßdispatch, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πTemp003 = µt
					if πE = πg.SetItem(πF, πTemp006, πTemp003, πTemp002); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					// line 1538: else: pass
					πF.SetLineno(1538)
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					goto Label6
					// line 1536: except: #TypeError, PicklingError, UnpicklingError
					πF.SetLineno(1536)
				Label6:
					// line 1537: log.info("skip: %s" % t)
					πF.SetLineno(1537)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("skip: %s").ToObject(), µt); πE != nil {
						continue
					}
					πTemp009[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlog); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 1539: return
					πF.SetLineno(1539)
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
			if πE = πF.Globals().SetItem(πF, ß_extend.ToObject(), πTemp098); πE != nil {
				continue
			}
			// line 1531: """extend pickle with all of dill's registered types"""
			πF.SetLineno(1531)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp099}, πg.NewStr("extend pickle with all of dill's registered types").ToObject()); πE != nil {
				continue
			}
			if πTemp100, πE = πg.ResolveGlobal(πF, ß_extend); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp100, ß__doc__, πTemp099); πE != nil {
				continue
			}
			// line 1541: del diff, _use_diff, use_diff
			πF.SetLineno(1541)
			if πE = πg.DelVar(πF, πF.Globals(), ßdiff); πE != nil {
				continue
			}
			if πE = πg.DelVar(πF, πF.Globals(), ß_use_diff); πE != nil {
				continue
			}
			if πE = πg.DelVar(πF, πF.Globals(), ßuse_diff); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("_dill", Code)
}

