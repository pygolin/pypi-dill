package _objects
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/pkg/cStringIO"
	// _ "github.com/pygolin/stdlib/pkg/functools"
	// _ "github.com/pygolin/stdlib/pkg/datetime"
	// _ "github.com/pygolin/stdlib/pkg/contextlib"
	// _ "github.com/pygolin/stdlib/pkg/io"
	// _ "github.com/pygolin/stdlib/pkg/operator"
	// _ "github.com/pygolin/stdlib/pkg/calendar"
	// _ "github.com/pygolin/stdlib/pkg/struct_goreservedkeyword"
	// _ "github.com/pygolin/stdlib/pkg/tempfile"
	// _ "github.com/pygolin/stdlib/pkg/pprint"
	// _ "github.com/pygolin/stdlib/pkg/tarfile"
	// _ "github.com/pygolin/stdlib/pkg/re"
	// _ "github.com/pygolin/stdlib/pkg/mutex"
	// _ "github.com/pygolin/stdlib/pkg/sqlite3"
	// _ "github.com/pygolin/stdlib/pkg/collections"
	// _ "github.com/pygolin/stdlib/pkg/csv"
	// _ "github.com/pygolin/stdlib/pkg/fractions"
	// _ "github.com/pygolin/stdlib/pkg/warnings"
	// _ "github.com/pygolin/stdlib/pkg/optparse"
	// _ "github.com/pygolin/stdlib/pkg/shelve"
	// _ "github.com/pygolin/stdlib/pkg/zipfile"
	// _ "github.com/pygolin/stdlib/pkg/xdrlib"
	// _ "github.com/pygolin/stdlib/pkg/Queue"
	// _ "github.com/pygolin/stdlib/pkg/_pyio"
	// _ "github.com/pygolin/stdlib/pkg/ctypes"
	// _ "github.com/pygolin/stdlib/pkg/codecs"
	// _ "github.com/pygolin/stdlib/pkg/sys"
	// _ "github.com/pygolin/stdlib/pkg/hashlib"
	// _ "github.com/pygolin/stdlib/pkg/anydbm"
	// _ "github.com/pygolin/stdlib/pkg/logging"
	// _ "github.com/pygolin/stdlib/pkg/socket"
	// _ "github.com/pygolin/stdlib/pkg/StringIO"
	// _ "github.com/pygolin/stdlib/pkg/decimal"
	// _ "github.com/pygolin/stdlib/pkg/weakref"
	// _ "github.com/pygolin/stdlib/pkg/threading"
	// _ "github.com/pygolin/stdlib/pkg/itertools"
	// _ "github.com/pygolin/stdlib/pkg/argparse"
	// _ "github.com/pygolin/stdlib/pkg/sets"
	// _ "github.com/pygolin/stdlib/pkg/gzip"
	// _ "github.com/pygolin/stdlib/pkg/os"
	// _ "github.com/pygolin/stdlib/pkg/hmac"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ß0x2070ef0 := πg.InternStr("0x2070ef0")
		ß0x30000f0 := πg.InternStr("0x30000f0")
		ß0x30800a1 := πg.InternStr("0x30800a1")
		ß1 := πg.InternStr("1")
		ßArgDefaultsHelpFormatterType := πg.InternStr("ArgDefaultsHelpFormatterType")
		ßArgParseFileType := πg.InternStr("ArgParseFileType")
		ßArgumentDefaultsHelpFormatter := πg.InternStr("ArgumentDefaultsHelpFormatter")
		ßArrayType := πg.InternStr("ArrayType")
		ßAttrGetterType := πg.InternStr("AttrGetterType")
		ßAttributeError := πg.InternStr("AttributeError")
		ßBZ2Compressor := πg.InternStr("BZ2Compressor")
		ßBZ2CompressorType := πg.InternStr("BZ2CompressorType")
		ßBZ2Decompressor := πg.InternStr("BZ2Decompressor")
		ßBZ2DecompressorType := πg.InternStr("BZ2DecompressorType")
		ßBZ2File := πg.InternStr("BZ2File")
		ßBZ2FileType := πg.InternStr("BZ2FileType")
		ßBigEndianStructure := πg.InternStr("BigEndianStructure")
		ßBigEndianStructureType := πg.InternStr("BigEndianStructureType")
		ßBooleanType := πg.InternStr("BooleanType")
		ßBufferType := πg.InternStr("BufferType")
		ßBufferedIOBase := πg.InternStr("BufferedIOBase")
		ßBufferedIOBaseType := πg.InternStr("BufferedIOBaseType")
		ßBufferedRandomType := πg.InternStr("BufferedRandomType")
		ßBufferedReaderType := πg.InternStr("BufferedReaderType")
		ßBufferedWriterType := πg.InternStr("BufferedWriterType")
		ßBuiltinFunctionType := πg.InternStr("BuiltinFunctionType")
		ßBuiltinMethodType := πg.InternStr("BuiltinMethodType")
		ßByteArrayType := πg.InternStr("ByteArrayType")
		ßBytesIO := πg.InternStr("BytesIO")
		ßBytesType := πg.InternStr("BytesType")
		ßCBoolType := πg.InternStr("CBoolType")
		ßCByteType := πg.InternStr("CByteType")
		ßCCharArrayType := πg.InternStr("CCharArrayType")
		ßCCharPType := πg.InternStr("CCharPType")
		ßCCharType := πg.InternStr("CCharType")
		ßCDLL := πg.InternStr("CDLL")
		ßCDLLType := πg.InternStr("CDLLType")
		ßCDoubleType := πg.InternStr("CDoubleType")
		ßCFUNCTYPE := πg.InternStr("CFUNCTYPE")
		ßCFUNCTYPEType := πg.InternStr("CFUNCTYPEType")
		ßCFloatType := πg.InternStr("CFloatType")
		ßCFunctionType := πg.InternStr("CFunctionType")
		ßCIntType := πg.InternStr("CIntType")
		ßCLibraryLoaderType := πg.InternStr("CLibraryLoaderType")
		ßCLongDoubleType := πg.InternStr("CLongDoubleType")
		ßCLongLongType := πg.InternStr("CLongLongType")
		ßCLongType := πg.InternStr("CLongType")
		ßCParamType := πg.InternStr("CParamType")
		ßCSSizeTType := πg.InternStr("CSSizeTType")
		ßCSVDictReaderType := πg.InternStr("CSVDictReaderType")
		ßCSVDictWriterType := πg.InternStr("CSVDictWriterType")
		ßCSVReaderType := πg.InternStr("CSVReaderType")
		ßCSVWriterType := πg.InternStr("CSVWriterType")
		ßCShortType := πg.InternStr("CShortType")
		ßCSizeTType := πg.InternStr("CSizeTType")
		ßCUByteType := πg.InternStr("CUByteType")
		ßCUIntType := πg.InternStr("CUIntType")
		ßCULongLongType := πg.InternStr("CULongLongType")
		ßCULongType := πg.InternStr("CULongType")
		ßCUShortType := πg.InternStr("CUShortType")
		ßCVoidPType := πg.InternStr("CVoidPType")
		ßCWCharArrayType := πg.InternStr("CWCharArrayType")
		ßCWCharPType := πg.InternStr("CWCharPType")
		ßCWCharType := πg.InternStr("CWCharType")
		ßCalendar := πg.InternStr("Calendar")
		ßCalendarType := πg.InternStr("CalendarType")
		ßCallableIteratorType := πg.InternStr("CallableIteratorType")
		ßCallableProxyType := πg.InternStr("CallableProxyType")
		ßCellType := πg.InternStr("CellType")
		ßChainType := πg.InternStr("ChainType")
		ßClassInstanceType := πg.InternStr("ClassInstanceType")
		ßClassMethodDescriptorType := πg.InternStr("ClassMethodDescriptorType")
		ßClassMethodType := πg.InternStr("ClassMethodType")
		ßClassObjectType := πg.InternStr("ClassObjectType")
		ßClassType := πg.InternStr("ClassType")
		ßClosedFileType := πg.InternStr("ClosedFileType")
		ßCmpKeyObjType := πg.InternStr("CmpKeyObjType")
		ßCmpKeyType := πg.InternStr("CmpKeyType")
		ßCodeType := πg.InternStr("CodeType")
		ßCombinationsType := πg.InternStr("CombinationsType")
		ßComplexType := πg.InternStr("ComplexType")
		ßCompressType := πg.InternStr("CompressType")
		ßConnectionType := πg.InternStr("ConnectionType")
		ßCopyrightType := πg.InternStr("CopyrightType")
		ßCountType := πg.InternStr("CountType")
		ßCounter := πg.InternStr("Counter")
		ßCounterType := πg.InternStr("CounterType")
		ßCursorType := πg.InternStr("CursorType")
		ßCycleType := πg.InternStr("CycleType")
		ßDateTimeType := πg.InternStr("DateTimeType")
		ßDbmType := πg.InternStr("DbmType")
		ßDeadCallableProxyType := πg.InternStr("DeadCallableProxyType")
		ßDeadProxyType := πg.InternStr("DeadProxyType")
		ßDeadReferenceType := πg.InternStr("DeadReferenceType")
		ßDecimal := πg.InternStr("Decimal")
		ßDecimalType := πg.InternStr("DecimalType")
		ßDefaultDictType := πg.InternStr("DefaultDictType")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßDequeType := πg.InternStr("DequeType")
		ßDialectType := πg.InternStr("DialectType")
		ßDictItemsType := πg.InternStr("DictItemsType")
		ßDictKeysType := πg.InternStr("DictKeysType")
		ßDictProxyType := πg.InternStr("DictProxyType")
		ßDictProxyType2 := πg.InternStr("DictProxyType2")
		ßDictReader := πg.InternStr("DictReader")
		ßDictType := πg.InternStr("DictType")
		ßDictValuesType := πg.InternStr("DictValuesType")
		ßDictWriter := πg.InternStr("DictWriter")
		ßDictionaryItemIteratorType := πg.InternStr("DictionaryItemIteratorType")
		ßDictionaryKeyIteratorType := πg.InternStr("DictionaryKeyIteratorType")
		ßDictionaryType := πg.InternStr("DictionaryType")
		ßDictionaryValueIteratorType := πg.InternStr("DictionaryValueIteratorType")
		ßEllipsis := πg.InternStr("Ellipsis")
		ßEllipsisType := πg.InternStr("EllipsisType")
		ßException := πg.InternStr("Exception")
		ßExceptionType := πg.InternStr("ExceptionType")
		ßExitType := πg.InternStr("ExitType")
		ßFalse := πg.InternStr("False")
		ßFieldType := πg.InternStr("FieldType")
		ßFileHandler := πg.InternStr("FileHandler")
		ßFileHandlerType := πg.InternStr("FileHandlerType")
		ßFileType := πg.InternStr("FileType")
		ßFilter := πg.InternStr("Filter")
		ßFilterType := πg.InternStr("FilterType")
		ßFloatType := πg.InternStr("FloatType")
		ßFormatter := πg.InternStr("Formatter")
		ßFormatterType := πg.InternStr("FormatterType")
		ßFraction := πg.InternStr("Fraction")
		ßFractionType := πg.InternStr("FractionType")
		ßFrameType := πg.InternStr("FrameType")
		ßFrozenSetType := πg.InternStr("FrozenSetType")
		ßFuncPtrType := πg.InternStr("FuncPtrType")
		ßFunctionType := πg.InternStr("FunctionType")
		ßGeneratorContextManager := πg.InternStr("GeneratorContextManager")
		ßGeneratorContextManagerType := πg.InternStr("GeneratorContextManagerType")
		ßGeneratorType := πg.InternStr("GeneratorType")
		ßGetSetDescriptorType := πg.InternStr("GetSetDescriptorType")
		ßGzipFile := πg.InternStr("GzipFile")
		ßGzipFileType := πg.InternStr("GzipFileType")
		ßHAS_ALL := πg.InternStr("HAS_ALL")
		ßHAS_CTYPES := πg.InternStr("HAS_CTYPES")
		ßHAS_CURSES := πg.InternStr("HAS_CURSES")
		ßHMACType := πg.InternStr("HMACType")
		ßHashType := πg.InternStr("HashType")
		ßIOBase := πg.InternStr("IOBase")
		ßIOBaseType := πg.InternStr("IOBaseType")
		ßIS_PYPY := πg.InternStr("IS_PYPY")
		ßImmutableSet := πg.InternStr("ImmutableSet")
		ßImmutableSetType := πg.InternStr("ImmutableSetType")
		ßImportError := πg.InternStr("ImportError")
		ßInputType := πg.InternStr("InputType")
		ßInstanceType := πg.InternStr("InstanceType")
		ßIntType := πg.InternStr("IntType")
		ßItemGetterType := πg.InternStr("ItemGetterType")
		ßIzipType := πg.InternStr("IzipType")
		ßLPCCharObjType := πg.InternStr("LPCCharObjType")
		ßLPCCharType := πg.InternStr("LPCCharType")
		ßLambdaType := πg.InternStr("LambdaType")
		ßListIteratorType := πg.InternStr("ListIteratorType")
		ßListType := πg.InternStr("ListType")
		ßLock := πg.InternStr("Lock")
		ßLockType := πg.InternStr("LockType")
		ßLogRecordType := πg.InternStr("LogRecordType")
		ßLoggerType := πg.InternStr("LoggerType")
		ßLoggingAdapter := πg.InternStr("LoggingAdapter")
		ßLoggingAdapterType := πg.InternStr("LoggingAdapterType")
		ßLongType := πg.InternStr("LongType")
		ßMemberDescriptorType := πg.InternStr("MemberDescriptorType")
		ßMemberDescriptorType2 := πg.InternStr("MemberDescriptorType2")
		ßMemoryHandler := πg.InternStr("MemoryHandler")
		ßMemoryHandlerType := πg.InternStr("MemoryHandlerType")
		ßMemoryType := πg.InternStr("MemoryType")
		ßMemoryType2 := πg.InternStr("MemoryType2")
		ßMethodCallerType := πg.InternStr("MethodCallerType")
		ßMethodDescriptorType := πg.InternStr("MethodDescriptorType")
		ßMethodType := πg.InternStr("MethodType")
		ßMethodWrapperType := πg.InternStr("MethodWrapperType")
		ßModuleType := πg.InternStr("ModuleType")
		ßMutexType := πg.InternStr("MutexType")
		ßNameError := πg.InternStr("NameError")
		ßNamedLoggerType := πg.InternStr("NamedLoggerType")
		ßNone := πg.InternStr("None")
		ßNoneType := πg.InternStr("NoneType")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßNotImplementedType := πg.InternStr("NotImplementedType")
		ßNullHandler := πg.InternStr("NullHandler")
		ßNullHandlerType := πg.InternStr("NullHandlerType")
		ßNullPtrType := πg.InternStr("NullPtrType")
		ßNullPyObjectType := πg.InternStr("NullPyObjectType")
		ßNumber := πg.InternStr("Number")
		ßNumberType := πg.InternStr("NumberType")
		ßNumpyArrayType := πg.InternStr("NumpyArrayType")
		ßNumpyInt32Type := πg.InternStr("NumpyInt32Type")
		ßNumpyUfuncType := πg.InternStr("NumpyUfuncType")
		ßObjectType := πg.InternStr("ObjectType")
		ßOption := πg.InternStr("Option")
		ßOptionGroup := πg.InternStr("OptionGroup")
		ßOptionGroupType := πg.InternStr("OptionGroupType")
		ßOptionParser := πg.InternStr("OptionParser")
		ßOptionParserType := πg.InternStr("OptionParserType")
		ßOptionType := πg.InternStr("OptionType")
		ßOrderedDict := πg.InternStr("OrderedDict")
		ßOrderedDictType := πg.InternStr("OrderedDictType")
		ßOutputType := πg.InternStr("OutputType")
		ßPOINTER := πg.InternStr("POINTER")
		ßPROG := πg.InternStr("PROG")
		ßPY3 := πg.InternStr("PY3")
		ßPacker := πg.InternStr("Packer")
		ßPackerType := πg.InternStr("PackerType")
		ßPartialType := πg.InternStr("PartialType")
		ßPermutationsType := πg.InternStr("PermutationsType")
		ßPrettyPrinter := πg.InternStr("PrettyPrinter")
		ßPrettyPrinterType := πg.InternStr("PrettyPrinterType")
		ßProductType := πg.InternStr("ProductType")
		ßPropertyType := πg.InternStr("PropertyType")
		ßProxyType := πg.InternStr("ProxyType")
		ßPyBufferedRandomType := πg.InternStr("PyBufferedRandomType")
		ßPyBufferedReaderType := πg.InternStr("PyBufferedReaderType")
		ßPyBufferedWriterType := πg.InternStr("PyBufferedWriterType")
		ßPyDLLType := πg.InternStr("PyDLLType")
		ßPyObjectType := πg.InternStr("PyObjectType")
		ßPyTextWrapperType := πg.InternStr("PyTextWrapperType")
		ßQueue := πg.InternStr("Queue")
		ßQueueType := πg.InternStr("QueueType")
		ßQuitterType := πg.InternStr("QuitterType")
		ßRLock := πg.InternStr("RLock")
		ßRLockType := πg.InternStr("RLockType")
		ßRawDescriptionHelpFormatter := πg.InternStr("RawDescriptionHelpFormatter")
		ßRawDescriptionHelpFormatterType := πg.InternStr("RawDescriptionHelpFormatterType")
		ßRawIOBase := πg.InternStr("RawIOBase")
		ßRawIOBaseType := πg.InternStr("RawIOBaseType")
		ßRawTextHelpFormatter := πg.InternStr("RawTextHelpFormatter")
		ßRawTextHelpFormatterType := πg.InternStr("RawTextHelpFormatterType")
		ßReferenceType := πg.InternStr("ReferenceType")
		ßRepeatType := πg.InternStr("RepeatType")
		ßRotatingFileHandler := πg.InternStr("RotatingFileHandler")
		ßRotatingFileHandlerType := πg.InternStr("RotatingFileHandlerType")
		ßSREMatchType := πg.InternStr("SREMatchType")
		ßSREPatternType := πg.InternStr("SREPatternType")
		ßSREScannerType := πg.InternStr("SREScannerType")
		ßSet := πg.InternStr("Set")
		ßSetIteratorType := πg.InternStr("SetIteratorType")
		ßSetType := πg.InternStr("SetType")
		ßSetsType := πg.InternStr("SetsType")
		ßShelf := πg.InternStr("Shelf")
		ßShelveType := πg.InternStr("ShelveType")
		ßSliceType := πg.InternStr("SliceType")
		ßSocketHandler := πg.InternStr("SocketHandler")
		ßSocketHandlerType := πg.InternStr("SocketHandlerType")
		ßSocketPairType := πg.InternStr("SocketPairType")
		ßSocketType := πg.InternStr("SocketType")
		ßStaticMethodType := πg.InternStr("StaticMethodType")
		ßStreamHandler := πg.InternStr("StreamHandler")
		ßStreamHandlerType := πg.InternStr("StreamHandlerType")
		ßStreamReader := πg.InternStr("StreamReader")
		ßStringIO := πg.InternStr("StringIO")
		ßStringType := πg.InternStr("StringType")
		ßStruct := πg.InternStr("Struct")
		ßStructType := πg.InternStr("StructType")
		ßStructure := πg.InternStr("Structure")
		ßStructureType := πg.InternStr("StructureType")
		ßSuperType := πg.InternStr("SuperType")
		ßTZInfoType := πg.InternStr("TZInfoType")
		ßTarFileType := πg.InternStr("TarFileType")
		ßTarInfo := πg.InternStr("TarInfo")
		ßTarInfoType := πg.InternStr("TarInfoType")
		ßTemporaryFile := πg.InternStr("TemporaryFile")
		ßTemporaryFileType := πg.InternStr("TemporaryFileType")
		ßTextIO := πg.InternStr("TextIO")
		ßTextIOBase := πg.InternStr("TextIOBase")
		ßTextIOBaseType := πg.InternStr("TextIOBaseType")
		ßTextWrapperType := πg.InternStr("TextWrapperType")
		ßTracebackType := πg.InternStr("TracebackType")
		ßTrue := πg.InternStr("True")
		ßTupleIteratorType := πg.InternStr("TupleIteratorType")
		ßTupleType := πg.InternStr("TupleType")
		ßTypeType := πg.InternStr("TypeType")
		ßUnboundMethodType := πg.InternStr("UnboundMethodType")
		ßUnicodeIOType := πg.InternStr("UnicodeIOType")
		ßUnicodeType := πg.InternStr("UnicodeType")
		ßWeakKeyDictionary := πg.InternStr("WeakKeyDictionary")
		ßWeakKeyDictionaryType := πg.InternStr("WeakKeyDictionaryType")
		ßWeakSet := πg.InternStr("WeakSet")
		ßWeakSetType := πg.InternStr("WeakSetType")
		ßWeakValueDictionary := πg.InternStr("WeakValueDictionary")
		ßWeakValueDictionaryType := πg.InternStr("WeakValueDictionaryType")
		ßWrapperDescriptorType := πg.InternStr("WrapperDescriptorType")
		ßWrapperDescriptorType2 := πg.InternStr("WrapperDescriptorType2")
		ßXRangeIteratorType := πg.InternStr("XRangeIteratorType")
		ßXRangeType := πg.InternStr("XRangeType")
		ßZlibCompressType := πg.InternStr("ZlibCompressType")
		ßZlibDecompressType := πg.InternStr("ZlibDecompressType")
		ß_FuncPtr := πg.InternStr("_FuncPtr")
		ß_Struct := πg.InternStr("_Struct")
		ß__IPYTHON__ := πg.InternStr("__IPYTHON__")
		ß__all__ := πg.InternStr("__all__")
		ß__call__ := πg.InternStr("__call__")
		ß__closure__ := πg.InternStr("__closure__")
		ß__cmp__ := πg.InternStr("__cmp__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__lt__ := πg.InternStr("__lt__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__prepare__ := πg.InternStr("__prepare__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__slots__ := πg.InternStr("__slots__")
		ß_bytes := πg.InternStr("_bytes")
		ß_cchar := πg.InternStr("_cchar")
		ß_cdll := πg.InternStr("_cdll")
		ß_cfunc := πg.InternStr("_cfunc")
		ß_class := πg.InternStr("_class")
		ß_class2 := πg.InternStr("_class2")
		ß_cmpkey := πg.InternStr("_cmpkey")
		ß_conn := πg.InternStr("_conn")
		ß_cstrI := πg.InternStr("_cstrI")
		ß_cstrO := πg.InternStr("_cstrO")
		ß_dict := πg.InternStr("_dict")
		ß_exception := πg.InternStr("_exception")
		ß_field := πg.InternStr("_field")
		ß_fields_ := πg.InternStr("_fields_")
		ß_fileW := πg.InternStr("_fileW")
		ß_filedescrip := πg.InternStr("_filedescrip")
		ß_function := πg.InternStr("_function")
		ß_function2 := πg.InternStr("_function2")
		ß_generator := πg.InternStr("_generator")
		ß_in := πg.InternStr("_in")
		ß_instance := πg.InternStr("_instance")
		ß_instance2 := πg.InternStr("_instance2")
		ß_int := πg.InternStr("_int")
		ß_lambda := πg.InternStr("_lambda")
		ß_list := πg.InternStr("_list")
		ß_logger := πg.InternStr("_logger")
		ß_lpchar := πg.InternStr("_lpchar")
		ß_method := πg.InternStr("_method")
		ß_methodwrap := πg.InternStr("_methodwrap")
		ß_newclass := πg.InternStr("_newclass")
		ß_newclass2 := πg.InternStr("_newclass2")
		ß_numpy_array := πg.InternStr("_numpy_array")
		ß_numpy_int32 := πg.InternStr("_numpy_int32")
		ß_numpy_ufunc := πg.InternStr("_numpy_ufunc")
		ß_oparser := πg.InternStr("_oparser")
		ß_open := πg.InternStr("_open")
		ß_pydll := πg.InternStr("_pydll")
		ß_set := πg.InternStr("_set")
		ß_sock := πg.InternStr("_sock")
		ß_socket := πg.InternStr("_socket")
		ß_srepattern := πg.InternStr("_srepattern")
		ß_str := πg.InternStr("_str")
		ß_tempfile := πg.InternStr("_tempfile")
		ß_tmpf := πg.InternStr("_tmpf")
		ß_tuple := πg.InternStr("_tuple")
		ß_xrange := πg.InternStr("_xrange")
		ßa := πg.InternStr("a")
		ßanydbm := πg.InternStr("anydbm")
		ßargparse := πg.InternStr("argparse")
		ßarray := πg.InternStr("array")
		ßattrgetter := πg.InternStr("attrgetter")
		ßbool := πg.InternStr("bool")
		ßbuffer := πg.InternStr("buffer")
		ßbyref := πg.InternStr("byref")
		ßbytearray := πg.InternStr("bytearray")
		ßbz2 := πg.InternStr("bz2")
		ßc := πg.InternStr("c")
		ßc_bool := πg.InternStr("c_bool")
		ßc_byte := πg.InternStr("c_byte")
		ßc_char := πg.InternStr("c_char")
		ßc_char_p := πg.InternStr("c_char_p")
		ßc_double := πg.InternStr("c_double")
		ßc_float := πg.InternStr("c_float")
		ßc_int := πg.InternStr("c_int")
		ßc_long := πg.InternStr("c_long")
		ßc_longdouble := πg.InternStr("c_longdouble")
		ßc_longlong := πg.InternStr("c_longlong")
		ßc_short := πg.InternStr("c_short")
		ßc_size_t := πg.InternStr("c_size_t")
		ßc_ssize_t := πg.InternStr("c_ssize_t")
		ßc_ubyte := πg.InternStr("c_ubyte")
		ßc_uint := πg.InternStr("c_uint")
		ßc_ulong := πg.InternStr("c_ulong")
		ßc_ulonglong := πg.InternStr("c_ulonglong")
		ßc_ushort := πg.InternStr("c_ushort")
		ßc_void_p := πg.InternStr("c_void_p")
		ßc_wchar := πg.InternStr("c_wchar")
		ßc_wchar_p := πg.InternStr("c_wchar_p")
		ßcalendar := πg.InternStr("calendar")
		ßcdll := πg.InternStr("cdll")
		ßchain := πg.InternStr("chain")
		ßclassmethod := πg.InternStr("classmethod")
		ßclose := πg.InternStr("close")
		ßcmp_to_key := πg.InternStr("cmp_to_key")
		ßcodecs := πg.InternStr("codecs")
		ßcollections := πg.InternStr("collections")
		ßcombinations := πg.InternStr("combinations")
		ßcompile := πg.InternStr("compile")
		ßcomplex := πg.InternStr("complex")
		ßcompress := πg.InternStr("compress")
		ßcompressobj := πg.InternStr("compressobj")
		ßconnect := πg.InternStr("connect")
		ßcontextlib := πg.InternStr("contextlib")
		ßcontextmanager := πg.InternStr("contextmanager")
		ßcopyright := πg.InternStr("copyright")
		ßcount := πg.InternStr("count")
		ßcreate_string_buffer := πg.InternStr("create_string_buffer")
		ßcreate_unicode_buffer := πg.InternStr("create_unicode_buffer")
		ßcsv := πg.InternStr("csv")
		ßctypes := πg.InternStr("ctypes")
		ßcursor := πg.InternStr("cursor")
		ßcycle := πg.InternStr("cycle")
		ßd := πg.InternStr("d")
		ßdatetime := πg.InternStr("datetime")
		ßdays := πg.InternStr("days")
		ßdbm := πg.InternStr("dbm")
		ßdecimal := πg.InternStr("decimal")
		ßdecompressobj := πg.InternStr("decompressobj")
		ßdefaultdict := πg.InternStr("defaultdict")
		ßdeque := πg.InternStr("deque")
		ßdescriptor := πg.InternStr("descriptor")
		ßdevnull := πg.InternStr("devnull")
		ßdict := πg.InternStr("dict")
		ßexc_info := πg.InternStr("exc_info")
		ßexcel := πg.InternStr("excel")
		ßexec := πg.InternStr("exec")
		ßf := πg.InternStr("f")
		ßfailures := πg.InternStr("failures")
		ßfilterwarnings := πg.InternStr("filterwarnings")
		ßfinditer := πg.InternStr("finditer")
		ßfloat := πg.InternStr("float")
		ßfoo := πg.InternStr("foo")
		ßfractions := πg.InternStr("fractions")
		ßfrozenset := πg.InternStr("frozenset")
		ßfunc_closure := πg.InternStr("func_closure")
		ßfunctools := πg.InternStr("functools")
		ßgetLogger := πg.InternStr("getLogger")
		ßget_dialect := πg.InternStr("get_dialect")
		ßgi_frame := πg.InternStr("gi_frame")
		ßgzip := πg.InternStr("gzip")
		ßhandlers := πg.InternStr("handlers")
		ßhasattr := πg.InternStr("hasattr")
		ßhashlib := πg.InternStr("hashlib")
		ßhex := πg.InternStr("hex")
		ßhexversion := πg.InternStr("hexversion")
		ßhmac := πg.InternStr("hmac")
		ßignore := πg.InternStr("ignore")
		ßint := πg.InternStr("int")
		ßint32 := πg.InternStr("int32")
		ßio := πg.InternStr("io")
		ßitemgetter := πg.InternStr("itemgetter")
		ßitems := πg.InternStr("items")
		ßiter := πg.InternStr("iter")
		ßiteritems := πg.InternStr("iteritems")
		ßiterkeys := πg.InternStr("iterkeys")
		ßitertools := πg.InternStr("itertools")
		ßitervalues := πg.InternStr("itervalues")
		ßizip := πg.InternStr("izip")
		ßkeys := πg.InternStr("keys")
		ßlatin_1_encode := πg.InternStr("latin_1_encode")
		ßlen := πg.InternStr("len")
		ßlocalhost := πg.InternStr("localhost")
		ßlogging := πg.InternStr("logging")
		ßlong := πg.InternStr("long")
		ßmakeLogRecord := πg.InternStr("makeLogRecord")
		ßmatch := πg.InternStr("match")
		ßmax := πg.InternStr("max")
		ßmd5 := πg.InternStr("md5")
		ßmemoryview := πg.InternStr("memoryview")
		ßmethodcaller := πg.InternStr("methodcaller")
		ßmkstemp := πg.InternStr("mkstemp")
		ßmro := πg.InternStr("mro")
		ßmsvcrt := πg.InternStr("msvcrt")
		ßmutex := πg.InternStr("mutex")
		ßn := πg.InternStr("n")
		ßnew := πg.InternStr("new")
		ßnext := πg.InternStr("next")
		ßnumber := πg.InternStr("number")
		ßnumbers := πg.InternStr("numbers")
		ßobject := πg.InternStr("object")
		ßodict := πg.InternStr("odict")
		ßopen := πg.InternStr("open")
		ßoperator := πg.InternStr("operator")
		ßoptparse := πg.InternStr("optparse")
		ßos := πg.InternStr("os")
		ßpartial := πg.InternStr("partial")
		ßpermutations := πg.InternStr("permutations")
		ßplatform := πg.InternStr("platform")
		ßpointer := πg.InternStr("pointer")
		ßpprint := πg.InternStr("pprint")
		ßproduct := πg.InternStr("product")
		ßproperty := πg.InternStr("property")
		ßproxy := πg.InternStr("proxy")
		ßpy_object := πg.InternStr("py_object")
		ßpythonapi := πg.InternStr("pythonapi")
		ßquit := πg.InternStr("quit")
		ßr := πg.InternStr("r")
		ßrange := πg.InternStr("range")
		ßrb := πg.InternStr("rb")
		ßre := πg.InternStr("re")
		ßreader := πg.InternStr("reader")
		ßref := πg.InternStr("ref")
		ßregistered := πg.InternStr("registered")
		ßremove := πg.InternStr("remove")
		ßrepeat := πg.InternStr("repeat")
		ßscanner := πg.InternStr("scanner")
		ßset := πg.InternStr("set")
		ßsets := πg.InternStr("sets")
		ßshelve := πg.InternStr("shelve")
		ßslice := πg.InternStr("slice")
		ßsocket := πg.InternStr("socket")
		ßsocketpair := πg.InternStr("socketpair")
		ßsqlite3 := πg.InternStr("sqlite3")
		ßstaticmethod := πg.InternStr("staticmethod")
		ßstr := πg.InternStr("str")
		ßstruct := πg.InternStr("struct")
		ßsucceeds := πg.InternStr("succeeds")
		ßsuper := πg.InternStr("super")
		ßsys := πg.InternStr("sys")
		ßtarfile := πg.InternStr("tarfile")
		ßtempfile := πg.InternStr("tempfile")
		ßthreading := πg.InternStr("threading")
		ßtimedelta := πg.InternStr("timedelta")
		ßtoday := πg.InternStr("today")
		ßtype := πg.InternStr("type")
		ßtypecode := πg.InternStr("typecode")
		ßtzinfo := πg.InternStr("tzinfo")
		ßufunc := πg.InternStr("ufunc")
		ßunicode := πg.InternStr("unicode")
		ßupdate := πg.InternStr("update")
		ßvalues := πg.InternStr("values")
		ßviewitems := πg.InternStr("viewitems")
		ßviewkeys := πg.InternStr("viewkeys")
		ßviewvalues := πg.InternStr("viewvalues")
		ßw := πg.InternStr("w")
		ßwarnings := πg.InternStr("warnings")
		ßwb := πg.InternStr("wb")
		ßweakref := πg.InternStr("weakref")
		ßwin := πg.InternStr("win")
		ßwriter := πg.InternStr("writer")
		ßx := πg.InternStr("x")
		ßxdrlib := πg.InternStr("xdrlib")
		ßxrange := πg.InternStr("xrange")
		ßzip := πg.InternStr("zip")
		ßzipfile := πg.InternStr("zipfile")
		ßzlib := πg.InternStr("zlib")
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
		var πTemp006 bool
		_ = πTemp006
		var πTemp007 *πg.BaseException
		_ = πTemp007
		var πTemp008 *πg.Traceback
		_ = πTemp008
		var πTemp009 *πg.Dict
		_ = πTemp009
		var πTemp010 []πg.Param
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 []*πg.Object
		_ = πTemp014
		var πTemp015 *πg.BaseException
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 97: goto Label97
			case 132: goto Label132
			case 123: goto Label123
			case 100: goto Label100
			case 5: goto Label5
			case 103: goto Label103
			case 40: goto Label40
			case 11: goto Label11
			case 45: goto Label45
			case 17: goto Label17
			case 20: goto Label20
			case 126: goto Label126
			case 88: goto Label88
			case 25: goto Label25
			case 91: goto Label91
			case 28: goto Label28
			case 94: goto Label94
			case 52: goto Label52
			default: panic("unexpected function state")
			}
			// line 8: """
			πF.SetLineno(8)
			// line 8: """
			πF.SetLineno(8)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nall Python Standard Library objects (currently: CH 1-15 @ 2.7)\nand some other common objects (i.e. numpy.ndarray)\n").ToObject()); πE != nil {
				continue
			}
			// line 13: __all__ = ['registered','failures','succeeds']
			πF.SetLineno(13)
			πTemp001 = make([]*πg.Object, 3)
			πTemp001[0] = ßregistered.ToObject()
			πTemp001[1] = ßfailures.ToObject()
			πTemp001[2] = ßsucceeds.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 16: import warnings; warnings.filterwarnings("ignore", category=DeprecationWarning)
			πF.SetLineno(16)
			if πTemp001, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßwarnings.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 16: import warnings; warnings.filterwarnings("ignore", category=DeprecationWarning)
			πF.SetLineno(16)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßignore.ToObject()
			if πTemp002, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
				continue
			}
			πTemp003 = πg.KWArgs{
				{"category", πTemp002},
			}
			if πTemp002, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßfilterwarnings, nil); πE != nil {
				continue
			}
			if πTemp002, πE = πTemp004.Call(πF, πTemp001, πTemp003); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			// line 17: import sys
			πF.SetLineno(17)
			if πTemp001, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 18: PY3 = (hex(sys.hexversion) >= '0x30000f0')
			πF.SetLineno(18)
			πTemp001 = πF.MakeArgs(1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßhexversion, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp004, πE = πg.ResolveGlobal(πF, ßhex); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp002, πE = πg.GE(πF, πTemp005, ß0x30000f0.ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPY3.ToObject(), πTemp002); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label1
			}
			goto Label2
			// line 19: if PY3:
			πF.SetLineno(19)
		Label1:
			// line 20: import queue as Queue
			πF.SetLineno(20)
			if πTemp001, πE = πg.ImportModule(πF, "queue"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßQueue.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 21: import dbm as anydbm
			πF.SetLineno(21)
			if πTemp001, πE = πg.ImportModule(πF, "dbm"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßanydbm.ToObject(), πTemp002); πE != nil {
				continue
			}
			goto Label3
		Label2:
			// line 23: import Queue
			πF.SetLineno(23)
			if πTemp001, πE = πg.ImportModule(πF, "Queue"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßQueue.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 24: import anydbm
			πF.SetLineno(24)
			if πTemp001, πE = πg.ImportModule(πF, "anydbm"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßanydbm.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 25: import sets # deprecated/removed
			πF.SetLineno(25)
			if πTemp001, πE = πg.ImportModule(πF, "sets"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßsets.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 26: import mutex # removed
			πF.SetLineno(26)
			if πTemp001, πE = πg.ImportModule(πF, "mutex"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßmutex.ToObject(), πTemp002); πE != nil {
				continue
			}
			goto Label3
		Label3:
			// line 27: try:
			πF.SetLineno(27)
			πF.PushCheckpoint(5)
			// line 28: from cStringIO import StringIO # has StringI and StringO types
			πF.SetLineno(28)
			if πTemp001, πE = πg.ImportModule(πF, "cStringIO"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp004, πE = πg.GetAttrImport(πF, πTemp002, ßStringIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp004); πE != nil {
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
			if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label6
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 29: except ImportError: # only has StringIO type
			πF.SetLineno(29)
		Label6:
			if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label7
			}
			goto Label8
			// line 30: if PY3:
			πF.SetLineno(30)
		Label7:
			// line 31: from io import BytesIO as StringIO
			πF.SetLineno(31)
			if πTemp001, πE = πg.ImportModule(πF, "io"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp004, πE = πg.GetAttrImport(πF, πTemp002, ßBytesIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp004); πE != nil {
				continue
			}
			goto Label9
		Label8:
			// line 33: from StringIO import StringIO
			πF.SetLineno(33)
			if πTemp001, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp004, πE = πg.GetAttrImport(πF, πTemp002, ßStringIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp004); πE != nil {
				continue
			}
			goto Label9
		Label9:
			πF.RestoreExc(nil, nil)
			goto Label4
		Label4:
			// line 34: import re
			πF.SetLineno(34)
			if πTemp001, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 35: import array
			πF.SetLineno(35)
			if πTemp001, πE = πg.ImportModule(πF, "array"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßarray.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 36: import collections
			πF.SetLineno(36)
			if πTemp001, πE = πg.ImportModule(πF, "collections"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßcollections.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 37: import codecs
			πF.SetLineno(37)
			if πTemp001, πE = πg.ImportModule(πF, "codecs"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßcodecs.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 38: import struct
			πF.SetLineno(38)
			if πTemp001, πE = πg.ImportModule(πF, "struct"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßstruct.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 39: import datetime
			πF.SetLineno(39)
			if πTemp001, πE = πg.ImportModule(πF, "datetime"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßdatetime.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 40: import calendar
			πF.SetLineno(40)
			if πTemp001, πE = πg.ImportModule(πF, "calendar"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßcalendar.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 41: import weakref
			πF.SetLineno(41)
			if πTemp001, πE = πg.ImportModule(πF, "weakref"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßweakref.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 42: import pprint
			πF.SetLineno(42)
			if πTemp001, πE = πg.ImportModule(πF, "pprint"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßpprint.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 43: import decimal
			πF.SetLineno(43)
			if πTemp001, πE = πg.ImportModule(πF, "decimal"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßdecimal.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 44: import functools
			πF.SetLineno(44)
			if πTemp001, πE = πg.ImportModule(πF, "functools"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßfunctools.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 45: import itertools
			πF.SetLineno(45)
			if πTemp001, πE = πg.ImportModule(πF, "itertools"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßitertools.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 46: import operator
			πF.SetLineno(46)
			if πTemp001, πE = πg.ImportModule(πF, "operator"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßoperator.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 47: import tempfile
			πF.SetLineno(47)
			if πTemp001, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßtempfile.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 48: import shelve
			πF.SetLineno(48)
			if πTemp001, πE = πg.ImportModule(πF, "shelve"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßshelve.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 49: import zlib
			πF.SetLineno(49)
			if πTemp001, πE = πg.ImportModule(πF, "zlib"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßzlib.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 50: import gzip
			πF.SetLineno(50)
			if πTemp001, πE = πg.ImportModule(πF, "gzip"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßgzip.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 51: import zipfile
			πF.SetLineno(51)
			if πTemp001, πE = πg.ImportModule(πF, "zipfile"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßzipfile.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 52: import tarfile
			πF.SetLineno(52)
			if πTemp001, πE = πg.ImportModule(πF, "tarfile"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßtarfile.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 53: import xdrlib
			πF.SetLineno(53)
			if πTemp001, πE = πg.ImportModule(πF, "xdrlib"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßxdrlib.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 54: import csv
			πF.SetLineno(54)
			if πTemp001, πE = πg.ImportModule(πF, "csv"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßcsv.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 55: import hashlib
			πF.SetLineno(55)
			if πTemp001, πE = πg.ImportModule(πF, "hashlib"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßhashlib.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 56: import hmac
			πF.SetLineno(56)
			if πTemp001, πE = πg.ImportModule(πF, "hmac"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßhmac.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 57: import os
			πF.SetLineno(57)
			if πTemp001, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 58: import logging
			πF.SetLineno(58)
			if πTemp001, πE = πg.ImportModule(πF, "logging"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßlogging.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 59: import optparse
			πF.SetLineno(59)
			if πTemp001, πE = πg.ImportModule(πF, "optparse"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßoptparse.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 61: import threading
			πF.SetLineno(61)
			if πTemp001, πE = πg.ImportModule(πF, "threading"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßthreading.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 62: import socket
			πF.SetLineno(62)
			if πTemp001, πE = πg.ImportModule(πF, "socket"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßsocket.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 63: import contextlib
			πF.SetLineno(63)
			if πTemp001, πE = πg.ImportModule(πF, "contextlib"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßcontextlib.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 64: try:
			πF.SetLineno(64)
			πF.PushCheckpoint(11)
			// line 65: import bz2
			πF.SetLineno(65)
			if πTemp001, πE = πg.ImportModule(πF, "bz2"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßbz2.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 66: import sqlite3
			πF.SetLineno(66)
			if πTemp001, πE = πg.ImportModule(πF, "sqlite3"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßsqlite3.ToObject(), πTemp002); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label12
			}
			goto Label13
			// line 67: if PY3: import dbm.ndbm as dbm
			πF.SetLineno(67)
		Label12:
			// line 67: if PY3: import dbm.ndbm as dbm
			πF.SetLineno(67)
			if πTemp001, πE = πg.ImportModule(πF, "dbm.ndbm"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[1]
			if πE = πF.Globals().SetItem(πF, ßdbm.ToObject(), πTemp002); πE != nil {
				continue
			}
			goto Label14
		Label13:
			// line 68: else: import dbm
			πF.SetLineno(68)
			if πTemp001, πE = πg.ImportModule(πF, "dbm"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßdbm.ToObject(), πTemp002); πE != nil {
				continue
			}
			goto Label14
		Label14:
			// line 69: HAS_ALL = True
			πF.SetLineno(69)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHAS_ALL.ToObject(), πTemp002); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label10
		Label11:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label15
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 70: except ImportError: # Ubuntu
			πF.SetLineno(70)
		Label15:
			// line 71: HAS_ALL = False
			πF.SetLineno(71)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHAS_ALL.ToObject(), πTemp002); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label10
		Label10:
			// line 72: try:
			πF.SetLineno(72)
			πF.PushCheckpoint(17)
			// line 75: HAS_CURSES = True
			πF.SetLineno(75)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHAS_CURSES.ToObject(), πTemp002); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label16
		Label17:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label18
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 76: except ImportError: # Windows
			πF.SetLineno(76)
		Label18:
			// line 77: HAS_CURSES = False
			πF.SetLineno(77)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHAS_CURSES.ToObject(), πTemp002); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label16
		Label16:
			// line 78: try:
			πF.SetLineno(78)
			πF.PushCheckpoint(20)
			// line 79: import ctypes
			πF.SetLineno(79)
			if πTemp001, πE = πg.ImportModule(πF, "ctypes"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßctypes.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 80: HAS_CTYPES = True
			πF.SetLineno(80)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHAS_CTYPES.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 82: IS_PYPY = not hasattr(ctypes, 'pythonapi')
			πF.SetLineno(82)
			πTemp001 = πF.MakeArgs(2)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			πTemp001[0] = πTemp004
			πTemp001[1] = ßpythonapi.ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			πTemp002 = πg.GetBool(!πTemp006).ToObject()
			if πE = πF.Globals().SetItem(πF, ßIS_PYPY.ToObject(), πTemp002); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label19
		Label20:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label21
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 83: except ImportError: # MacPorts
			πF.SetLineno(83)
		Label21:
			// line 84: HAS_CTYPES = False
			πF.SetLineno(84)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHAS_CTYPES.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 85: IS_PYPY = False
			πF.SetLineno(85)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIS_PYPY.ToObject(), πTemp002); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label19
		Label19:
			// line 88: class _class:
			πF.SetLineno(88)
			πTemp001 = make([]*πg.Object, 0)
			πTemp009 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_class", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 89: def _method(self):
					πF.SetLineno(89)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("_method", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 90: pass
							πF.SetLineno(90)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_method.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("_class").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_class.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 97: class _class2:
			πF.SetLineno(97)
			πTemp001 = make([]*πg.Object, 0)
			πTemp009 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_class2", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 98: def __call__(self):
					πF.SetLineno(98)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__call__", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 99: pass
							πF.SetLineno(99)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("_class2").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_class2.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 100: _instance2 = _class2()
			πF.SetLineno(100)
			if πTemp002, πE = πg.ResolveGlobal(πF, ß_class2); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_instance2.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 101: class _newclass(object):
			πF.SetLineno(101)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			πTemp009 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_newclass", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 102: def _method(self):
					πF.SetLineno(102)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("_method", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 103: pass
							πF.SetLineno(103)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_method.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("_newclass").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_newclass.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 110: class _newclass2(object):
			πF.SetLineno(110)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			πTemp009 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_newclass2", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 111: __slots__ = ['descriptor']
					πF.SetLineno(111)
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
			if πTemp004, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("_newclass2").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_newclass2.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 112: def _function(x): yield x
			πF.SetLineno(112)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "x", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("_function", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						default: panic("unexpected function state")
						}
						// line 112: def _function(x): yield x
						πF.SetLineno(112)
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πF.PushCheckpoint(1)
						return µx, nil
					Label1:
						πTemp001 = πSent
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_function.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 113: def _function2():
			πF.SetLineno(113)
			πTemp010 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("_function2", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µexc_info *πg.Object = πg.UnboundLocal; _ = µexc_info
				var µe *πg.Object = πg.UnboundLocal; _ = µe
				var µer *πg.Object = πg.UnboundLocal; _ = µer
				var µtb *πg.Object = πg.UnboundLocal; _ = µtb
				var πTemp001 *πg.BaseException
				_ = πTemp001
				var πTemp002 *πg.Traceback
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 114: try: raise
					πF.SetLineno(114)
					πF.PushCheckpoint(2)
					// line 114: try: raise
					πF.SetLineno(114)
					πE = πF.Raise(nil, nil, nil)
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp001, πTemp002 = πF.ExcInfo()
					goto Label3
					// line 115: except:
					πF.SetLineno(115)
				Label3:
					// line 116: from sys import exc_info
					πF.SetLineno(116)
					if πTemp004, πE = πg.ImportModule(πF, "sys"); πE != nil {
						continue
					}
					πTemp003 = πTemp004[0]
					if πTemp005, πE = πg.GetAttrImport(πF, πTemp003, ßexc_info); πE != nil {
						continue
					}
					µexc_info = πTemp005
					// line 117: e, er, tb = exc_info()
					πF.SetLineno(117)
					if πE = πg.CheckLocal(πF, µexc_info, "exc_info"); πE != nil {
						continue
					}
					if πTemp003, πE = µexc_info.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					µe = πTemp005
					µer = πTemp006
					µtb = πTemp007
					// line 118: return er, tb
					πF.SetLineno(118)
					if πE = πg.CheckLocal(πF, µer, "er"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µer, µtb).ToObject()
					πR = πTemp003
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
			if πE = πF.Globals().SetItem(πF, ß_function2.ToObject(), πTemp004); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßHAS_CTYPES); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label22
			}
			goto Label23
			// line 119: if HAS_CTYPES:
			πF.SetLineno(119)
		Label22:
			// line 120: class _Struct(ctypes.Structure):
			πF.SetLineno(120)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßStructure, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp013
			πTemp009 = πg.NewDict()
			if πTemp005, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp005); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Struct", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 121: pass
					πF.SetLineno(121)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp011, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp011 == nil {
				πTemp011 = πg.TypeType.ToObject()
			}
			if πTemp012, πE = πTemp011.Call(πF, []*πg.Object{πg.NewStr("_Struct").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Struct.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 122: _Struct._fields_ = [("_field", ctypes.c_int),("next", ctypes.POINTER(_Struct))]
			πF.SetLineno(122)
			πTemp001 = make([]*πg.Object, 2)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßc_int, nil); πE != nil {
				continue
			}
			πTemp005 = πg.NewTuple2(ß_field.ToObject(), πTemp012).ToObject()
			πTemp001[0] = πTemp005
			πTemp014 = πF.MakeArgs(1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ß_Struct); πE != nil {
				continue
			}
			πTemp014[0] = πTemp011
			if πTemp011, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßPOINTER, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp012.Call(πF, πTemp014, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp014)
			πTemp005 = πg.NewTuple2(ßnext.ToObject(), πTemp011).ToObject()
			πTemp001[1] = πTemp005
			πTemp005 = πg.NewList(πTemp001...).ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_Struct); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß_fields_, πTemp011); πE != nil {
				continue
			}
			goto Label23
		Label23:
			// line 123: _filedescrip, _tempfile = tempfile.mkstemp('r') # deleted in cleanup
			πF.SetLineno(123)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßr.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßmkstemp, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp011}, πg.TieTarget{Target: &πTemp012}}}, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_filedescrip.ToObject(), πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_tempfile.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 124: _tmpf = tempfile.TemporaryFile('w')
			πF.SetLineno(124)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßw.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßTemporaryFile, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ß_tmpf.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 127: try:
			πF.SetLineno(127)
			πF.PushCheckpoint(25)
			// line 128: from collections import OrderedDict as odict
			πF.SetLineno(128)
			if πTemp001, πE = πg.ImportModule(πF, "collections"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πTemp011, πE = πg.GetAttrImport(πF, πTemp005, ßOrderedDict); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßodict.ToObject(), πTemp011); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label24
		Label25:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label26
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 129: except ImportError:
			πF.SetLineno(129)
		Label26:
			// line 130: try:
			πF.SetLineno(130)
			πF.PushCheckpoint(28)
			// line 131: from ordereddict import OrderedDict as odict
			πF.SetLineno(131)
			if πTemp001, πE = πg.ImportModule(πF, "ordereddict"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πTemp011, πE = πg.GetAttrImport(πF, πTemp005, ßOrderedDict); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßodict.ToObject(), πTemp011); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label27
		Label28:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp015, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp015.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label29
			}
			πE = πF.Raise(πTemp015.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 132: except ImportError:
			πF.SetLineno(132)
		Label29:
			// line 133: odict = dict
			πF.SetLineno(133)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßodict.ToObject(), πTemp005); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label27
		Label27:
			πF.RestoreExc(nil, nil)
			goto Label24
		Label24:
			// line 135: registered = d = odict()
			πF.SetLineno(135)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßodict); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßregistered.ToObject(), πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßd.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 137: failures = x = odict()
			πF.SetLineno(137)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßodict); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfailures.ToObject(), πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßx.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 139: succeeds = a = odict()
			πF.SetLineno(139)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßodict); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsucceeds.ToObject(), πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßa.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 142: a['BooleanType'] = bool(1)
			πF.SetLineno(142)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBooleanType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 143: a['BuiltinFunctionType'] = len
			πF.SetLineno(143)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBuiltinFunctionType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 144: a['BuiltinMethodType'] = a['BuiltinFunctionType']
			πF.SetLineno(144)
			πTemp005 = ßBuiltinFunctionType.ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBuiltinMethodType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 145: a['BytesType'] = _bytes = codecs.latin_1_encode('\x00')[0] # bytes(1)
			πF.SetLineno(145)
			πTemp005 = πg.NewInt(0).ToObject()
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewStr("\x00").ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßcodecs); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßlatin_1_encode, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp013.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBytesType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_bytes.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 146: a['ClassType'] = _class
			πF.SetLineno(146)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_class); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßClassType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 147: a['ComplexType'] = complex(1)
			πF.SetLineno(147)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcomplex); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßComplexType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 148: a['DictType'] = _dict = {}
			πF.SetLineno(148)
			πTemp009 = πg.NewDict()
			πTemp005 = πTemp009.ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDictType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_dict.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 149: a['DictionaryType'] = a['DictType']
			πF.SetLineno(149)
			πTemp005 = ßDictType.ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDictionaryType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 150: a['FloatType'] = float(1)
			πF.SetLineno(150)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßFloatType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 151: a['FunctionType'] = _function
			πF.SetLineno(151)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_function); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßFunctionType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 152: a['InstanceType'] = _instance = _class()
			πF.SetLineno(152)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_class); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßInstanceType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_instance.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 153: a['IntType'] = _int = int(1)
			πF.SetLineno(153)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßIntType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_int.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 154: a['ListType'] = _list = []
			πF.SetLineno(154)
			πTemp001 = make([]*πg.Object, 0)
			πTemp005 = πg.NewList(πTemp001...).ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßListType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_list.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 155: a['NoneType'] = None
			πF.SetLineno(155)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßNoneType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 156: a['ObjectType'] = object()
			πF.SetLineno(156)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßObjectType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 157: a['StringType'] = _str = str(1)
			πF.SetLineno(157)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßStringType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_str.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 158: a['TupleType'] = _tuple = ()
			πF.SetLineno(158)
			πTemp005 = πg.NewTuple0().ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßTupleType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_tuple.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 159: a['TypeType'] = type
			πF.SetLineno(159)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßTypeType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label30
			}
			goto Label31
			// line 160: if PY3:
			πF.SetLineno(160)
		Label30:
			// line 161: a['LongType'] = _int
			πF.SetLineno(161)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_int); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßLongType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 162: a['UnicodeType'] = _str
			πF.SetLineno(162)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_str); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßUnicodeType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label32
		Label31:
			// line 164: a['LongType'] = long(1)
			πF.SetLineno(164)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßLongType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 165: a['UnicodeType'] = unicode(1)
			πF.SetLineno(165)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßUnicodeType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			goto Label32
		Label32:
			// line 167: a['CopyrightType'] = copyright
			πF.SetLineno(167)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcopyright); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCopyrightType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 169: a['ClassObjectType'] = _newclass # <type 'type'>
			πF.SetLineno(169)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_newclass); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßClassObjectType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 170: a['ClassInstanceType'] = _newclass() # <type 'class'>
			πF.SetLineno(170)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_newclass); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßClassInstanceType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 171: a['SetType'] = _set = set()
			πF.SetLineno(171)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßSetType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_set.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 172: a['FrozenSetType'] = frozenset()
			πF.SetLineno(172)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßFrozenSetType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 174: a['ExceptionType'] = _exception = _function2()[0]
			πF.SetLineno(174)
			πTemp005 = πg.NewInt(0).ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_function2); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßExceptionType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_exception.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 176: a['SREPatternType'] = _srepattern = re.compile('')
			πF.SetLineno(176)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ß.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßSREPatternType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_srepattern.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 178: a['ArrayType'] = array.array("f")
			πF.SetLineno(178)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßf.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßarray); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßarray, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßArrayType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 179: a['DequeType'] = collections.deque([0])
			πF.SetLineno(179)
			πTemp001 = πF.MakeArgs(1)
			πTemp014 = make([]*πg.Object, 1)
			πTemp014[0] = πg.NewInt(0).ToObject()
			πTemp005 = πg.NewList(πTemp014...).ToObject()
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcollections); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdeque, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDequeType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 180: a['DefaultDictType'] = collections.defaultdict(_function, _dict)
			πF.SetLineno(180)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_function); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_dict); πE != nil {
				continue
			}
			πTemp001[1] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcollections); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdefaultdict, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDefaultDictType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 181: a['TZInfoType'] = datetime.tzinfo()
			πF.SetLineno(181)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßtzinfo, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßTZInfoType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 182: a['DateTimeType'] = datetime.datetime.today()
			πF.SetLineno(182)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdatetime, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßtoday, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDateTimeType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 183: a['CalendarType'] = calendar.Calendar()
			πF.SetLineno(183)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcalendar); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßCalendar, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCalendarType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp011); πE != nil {
				continue
			}
			πTemp005 = πg.GetBool(!πTemp006).ToObject()
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label33
			}
			goto Label34
			// line 184: if not PY3:
			πF.SetLineno(184)
		Label33:
			// line 185: a['SetsType'] = sets.Set()
			πF.SetLineno(185)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsets); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßSet, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßSetsType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 186: a['ImmutableSetType'] = sets.ImmutableSet()
			πF.SetLineno(186)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsets); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßImmutableSet, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßImmutableSetType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 187: a['MutexType'] = mutex.mutex()
			πF.SetLineno(187)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßmutex); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßmutex, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßMutexType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label34
		Label34:
			// line 189: a['DecimalType'] = decimal.Decimal(1)
			πF.SetLineno(189)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdecimal); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßDecimal, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDecimalType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 190: a['CountType'] = itertools.count(0)
			πF.SetLineno(190)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(0).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcount, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCountType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 192: a['TarInfoType'] = tarfile.TarInfo()
			πF.SetLineno(192)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtarfile); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßTarInfo, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßTarInfoType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 194: a['LoggerType'] = logging.getLogger()
			πF.SetLineno(194)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßgetLogger, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßLoggerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 195: a['FormatterType'] = logging.Formatter() # pickle ok
			πF.SetLineno(195)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßFormatter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßFormatterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 196: a['FilterType'] = logging.Filter() # pickle ok
			πF.SetLineno(196)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßFilter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßFilterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 197: a['LogRecordType'] = logging.makeLogRecord(_dict) # pickle ok
			πF.SetLineno(197)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_dict); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßmakeLogRecord, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßLogRecordType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 198: a['OptionParserType'] = _oparser = optparse.OptionParser() # pickle ok
			πF.SetLineno(198)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßOptionParser, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßOptionParserType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_oparser.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 199: a['OptionGroupType'] = optparse.OptionGroup(_oparser,"foo") # pickle ok
			πF.SetLineno(199)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_oparser); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			πTemp001[1] = ßfoo.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßOptionGroup, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßOptionGroupType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 200: a['OptionType'] = optparse.Option('--foo') # pickle ok
			πF.SetLineno(200)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewStr("--foo").ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßOption, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßOptionType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßHAS_CTYPES); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label35
			}
			goto Label36
			// line 201: if HAS_CTYPES:
			πF.SetLineno(201)
		Label35:
			// line 202: a['CCharType'] = _cchar = ctypes.c_char()
			πF.SetLineno(202)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_char, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCCharType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_cchar.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 203: a['CWCharType'] = ctypes.c_wchar() # fail == 2.6
			πF.SetLineno(203)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_wchar, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCWCharType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 204: a['CByteType'] = ctypes.c_byte()
			πF.SetLineno(204)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_byte, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCByteType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 205: a['CUByteType'] = ctypes.c_ubyte()
			πF.SetLineno(205)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_ubyte, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCUByteType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 206: a['CShortType'] = ctypes.c_short()
			πF.SetLineno(206)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_short, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCShortType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 207: a['CUShortType'] = ctypes.c_ushort()
			πF.SetLineno(207)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_ushort, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCUShortType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 208: a['CIntType'] = ctypes.c_int()
			πF.SetLineno(208)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_int, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCIntType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 209: a['CUIntType'] = ctypes.c_uint()
			πF.SetLineno(209)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_uint, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCUIntType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 210: a['CLongType'] = ctypes.c_long()
			πF.SetLineno(210)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_long, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCLongType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 211: a['CULongType'] = ctypes.c_ulong()
			πF.SetLineno(211)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_ulong, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCULongType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 212: a['CLongLongType'] = ctypes.c_longlong()
			πF.SetLineno(212)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_longlong, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCLongLongType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 213: a['CULongLongType'] = ctypes.c_ulonglong()
			πF.SetLineno(213)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_ulonglong, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCULongLongType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 214: a['CFloatType'] = ctypes.c_float()
			πF.SetLineno(214)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_float, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCFloatType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 215: a['CDoubleType'] = ctypes.c_double()
			πF.SetLineno(215)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_double, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCDoubleType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 216: a['CSizeTType'] = ctypes.c_size_t()
			πF.SetLineno(216)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_size_t, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCSizeTType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 217: a['CLibraryLoaderType'] = ctypes.cdll
			πF.SetLineno(217)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcdll, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCLibraryLoaderType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 218: a['StructureType'] = _Struct
			πF.SetLineno(218)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_Struct); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßStructureType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp011); πE != nil {
				continue
			}
			πTemp005 = πg.GetBool(!πTemp006).ToObject()
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label37
			}
			goto Label38
			// line 219: if not IS_PYPY:
			πF.SetLineno(219)
		Label37:
			// line 220: a['BigEndianStructureType'] = ctypes.BigEndianStructure()
			πF.SetLineno(220)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßBigEndianStructure, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBigEndianStructureType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label38
		Label38:
			goto Label36
		Label36:
			// line 226: try: # python 2.6
			πF.SetLineno(226)
			πF.PushCheckpoint(40)
			// line 227: import fractions
			πF.SetLineno(227)
			if πTemp001, πE = πg.ImportModule(πF, "fractions"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßfractions.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 228: import number
			πF.SetLineno(228)
			if πTemp001, πE = πg.ImportModule(πF, "number"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßnumber.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 229: import io
			πF.SetLineno(229)
			if πTemp001, πE = πg.ImportModule(πF, "io"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßio.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 230: from io import StringIO as TextIO
			πF.SetLineno(230)
			if πTemp001, πE = πg.ImportModule(πF, "io"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πTemp011, πE = πg.GetAttrImport(πF, πTemp005, ßStringIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTextIO.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 232: a['ByteArrayType'] = bytearray([1])
			πF.SetLineno(232)
			πTemp001 = πF.MakeArgs(1)
			πTemp014 = make([]*πg.Object, 1)
			πTemp014[0] = πg.NewInt(1).ToObject()
			πTemp005 = πg.NewList(πTemp014...).ToObject()
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßbytearray); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßByteArrayType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 234: a['FractionType'] = fractions.Fraction()
			πF.SetLineno(234)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßfractions); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßFraction, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßFractionType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 235: a['NumberType'] = numbers.Number()
			πF.SetLineno(235)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßnumbers); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßNumber, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßNumberType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 237: a['IOBaseType'] = io.IOBase()
			πF.SetLineno(237)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßIOBase, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßIOBaseType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 238: a['RawIOBaseType'] = io.RawIOBase()
			πF.SetLineno(238)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßRawIOBase, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßRawIOBaseType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 239: a['TextIOBaseType'] = io.TextIOBase()
			πF.SetLineno(239)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßTextIOBase, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßTextIOBaseType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 240: a['BufferedIOBaseType'] = io.BufferedIOBase()
			πF.SetLineno(240)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßBufferedIOBase, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBufferedIOBaseType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 241: a['UnicodeIOType'] = TextIO() # the new StringIO
			πF.SetLineno(241)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTextIO); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßUnicodeIOType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 242: a['LoggingAdapterType'] = logging.LoggingAdapter(_logger,_dict) # pickle ok
			πF.SetLineno(242)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_logger); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_dict); πE != nil {
				continue
			}
			πTemp001[1] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßLoggingAdapter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßLoggingAdapterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßHAS_CTYPES); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label41
			}
			goto Label42
			// line 243: if HAS_CTYPES:
			πF.SetLineno(243)
		Label41:
			// line 244: a['CBoolType'] = ctypes.c_bool(1)
			πF.SetLineno(244)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_bool, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCBoolType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 245: a['CLongDoubleType'] = ctypes.c_longdouble()
			πF.SetLineno(245)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_longdouble, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCLongDoubleType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label42
		Label42:
			πF.PopCheckpoint()
			goto Label39
		Label40:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label43
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 246: except ImportError:
			πF.SetLineno(246)
		Label43:
			// line 247: pass
			πF.SetLineno(247)
			πF.RestoreExc(nil, nil)
			goto Label39
		Label39:
			// line 248: try: # python 2.7
			πF.SetLineno(248)
			πF.PushCheckpoint(45)
			// line 249: import argparse
			πF.SetLineno(249)
			if πTemp001, πE = πg.ImportModule(πF, "argparse"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßargparse.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 251: a['OrderedDictType'] = collections.OrderedDict(_dict)
			πF.SetLineno(251)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_dict); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcollections); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßOrderedDict, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßOrderedDictType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 252: a['CounterType'] = collections.Counter(_dict)
			πF.SetLineno(252)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_dict); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcollections); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßCounter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCounterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßHAS_CTYPES); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label46
			}
			goto Label47
			// line 253: if HAS_CTYPES:
			πF.SetLineno(253)
		Label46:
			// line 254: a['CSSizeTType'] = ctypes.c_ssize_t()
			πF.SetLineno(254)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_ssize_t, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCSSizeTType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label47
		Label47:
			// line 256: a['NullHandlerType'] = logging.NullHandler() # pickle ok  # new 2.7
			πF.SetLineno(256)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßNullHandler, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßNullHandlerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 257: a['ArgParseFileType'] = argparse.FileType() # pickle ok
			πF.SetLineno(257)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßargparse); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßFileType, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßArgParseFileType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label44
		Label45:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp011, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			πTemp005 = πg.NewTuple2(πTemp011, πTemp012).ToObject()
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label48
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 258: except (AttributeError, ImportError):
			πF.SetLineno(258)
		Label48:
			// line 259: pass
			πF.SetLineno(259)
			πF.RestoreExc(nil, nil)
			goto Label44
		Label44:
			// line 263: a['CodeType'] = compile('','','exec')
			πF.SetLineno(263)
			πTemp001 = πF.MakeArgs(3)
			πTemp001[0] = ß.ToObject()
			πTemp001[1] = ß.ToObject()
			πTemp001[2] = ßexec.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcompile); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCodeType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 264: a['DictProxyType'] = type.__dict__
			πF.SetLineno(264)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß__dict__, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDictProxyType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 265: a['DictProxyType2'] = _newclass.__dict__
			πF.SetLineno(265)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_newclass); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß__dict__, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDictProxyType2.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 266: a['EllipsisType'] = Ellipsis
			πF.SetLineno(266)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßEllipsis); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßEllipsisType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 267: a['ClosedFileType'] = open(os.devnull, 'wb', buffering=0).close()
			πF.SetLineno(267)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp001[1] = ßwb.ToObject()
			πTemp003 = πg.KWArgs{
				{"buffering", πg.NewInt(0).ToObject()},
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, πTemp003); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßclose, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßClosedFileType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 268: a['GetSetDescriptorType'] = array.array.typecode
			πF.SetLineno(268)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßarray); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßarray, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßtypecode, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßGetSetDescriptorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 269: a['LambdaType'] = _lambda = lambda x: lambda y: x #XXX: works when not imported!
			πF.SetLineno(269)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "x", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 269: a['LambdaType'] = _lambda = lambda x: lambda y: x #XXX: works when not imported!
					πF.SetLineno(269)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "y", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µy *πg.Object = πArgs[0]; _ = µy
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 269: a['LambdaType'] = _lambda = lambda x: lambda y: x #XXX: works when not imported!
							πF.SetLineno(269)
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
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßLambdaType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_lambda.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 270: a['MemberDescriptorType'] = _newclass2.descriptor
			πF.SetLineno(270)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_newclass2); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdescriptor, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßMemberDescriptorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp011); πE != nil {
				continue
			}
			πTemp005 = πg.GetBool(!πTemp006).ToObject()
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label49
			}
			goto Label50
			// line 271: if not IS_PYPY:
			πF.SetLineno(271)
		Label49:
			// line 272: a['MemberDescriptorType2'] = datetime.timedelta.days
			πF.SetLineno(272)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßtimedelta, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßdays, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßMemberDescriptorType2.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label50
		Label50:
			// line 273: a['MethodType'] = _method = _class()._method #XXX: works when not imported!
			πF.SetLineno(273)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_class); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ß_method, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßMethodType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_method.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 274: a['ModuleType'] = datetime
			πF.SetLineno(274)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßModuleType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 275: a['NotImplementedType'] = NotImplemented
			πF.SetLineno(275)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßNotImplementedType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 276: a['SliceType'] = slice(1)
			πF.SetLineno(276)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßSliceType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 277: a['UnboundMethodType'] = _class._method #XXX: works when not imported!
			πF.SetLineno(277)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_class); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß_method, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßUnboundMethodType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 278: a['TextWrapperType'] = open(os.devnull, 'r') # same as mode='w','w+','r+'
			πF.SetLineno(278)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp001[1] = ßr.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßTextWrapperType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 279: a['BufferedRandomType'] = open(os.devnull, 'r+b') # same as mode='w+b'
			πF.SetLineno(279)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp001[1] = πg.NewStr("r+b").ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBufferedRandomType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 280: a['BufferedReaderType'] = open(os.devnull, 'rb') # (default: buffering=-1)
			πF.SetLineno(280)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp001[1] = ßrb.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBufferedReaderType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 281: a['BufferedWriterType'] = open(os.devnull, 'wb')
			πF.SetLineno(281)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp001[1] = ßwb.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBufferedWriterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 282: try: # oddities: deprecated
			πF.SetLineno(282)
			πF.PushCheckpoint(52)
			// line 283: from _pyio import open as _open
			πF.SetLineno(283)
			if πTemp001, πE = πg.ImportModule(πF, "_pyio"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πTemp011, πE = πg.GetAttrImport(πF, πTemp005, ßopen); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_open.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 284: a['PyTextWrapperType'] = _open(os.devnull, 'r', buffering=-1)
			πF.SetLineno(284)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp001[1] = ßr.ToObject()
			if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003 = πg.KWArgs{
				{"buffering", πTemp005},
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_open); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, πTemp003); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßPyTextWrapperType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 285: a['PyBufferedRandomType'] = _open(os.devnull, 'r+b', buffering=-1)
			πF.SetLineno(285)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp001[1] = πg.NewStr("r+b").ToObject()
			if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003 = πg.KWArgs{
				{"buffering", πTemp005},
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_open); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, πTemp003); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßPyBufferedRandomType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 286: a['PyBufferedReaderType'] = _open(os.devnull, 'rb', buffering=-1)
			πF.SetLineno(286)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp001[1] = ßrb.ToObject()
			if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003 = πg.KWArgs{
				{"buffering", πTemp005},
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_open); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, πTemp003); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßPyBufferedReaderType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 287: a['PyBufferedWriterType'] = _open(os.devnull, 'wb', buffering=-1)
			πF.SetLineno(287)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp001[1] = ßwb.ToObject()
			if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003 = πg.KWArgs{
				{"buffering", πTemp005},
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_open); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, πTemp003); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßPyBufferedWriterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label51
		Label52:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label53
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 288: except ImportError:
			πF.SetLineno(288)
		Label53:
			// line 289: pass
			πF.SetLineno(289)
			πF.RestoreExc(nil, nil)
			goto Label51
		Label51:
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label54
			}
			goto Label55
			// line 291: if PY3:
			πF.SetLineno(291)
		Label54:
			// line 292: d['CellType'] = (_lambda)(0).__closure__[0]
			πF.SetLineno(292)
			πTemp005 = πg.NewInt(0).ToObject()
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(0).ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_lambda); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp012, πE = πg.GetAttr(πF, πTemp013, ß__closure__, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp013 = ßCellType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 293: a['XRangeType'] = _xrange = range(1)
			πF.SetLineno(293)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßXRangeType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_xrange.ToObject(), πTemp011); πE != nil {
				continue
			}
			goto Label56
		Label55:
			// line 295: d['CellType'] = (_lambda)(0).func_closure[0]
			πF.SetLineno(295)
			πTemp005 = πg.NewInt(0).ToObject()
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(0).ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_lambda); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp012, πE = πg.GetAttr(πF, πTemp013, ßfunc_closure, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp013 = ßCellType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 296: a['XRangeType'] = _xrange = xrange(1)
			πF.SetLineno(296)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßXRangeType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_xrange.ToObject(), πTemp011); πE != nil {
				continue
			}
			goto Label56
		Label56:
			if πTemp011, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp011); πE != nil {
				continue
			}
			πTemp005 = πg.GetBool(!πTemp006).ToObject()
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label57
			}
			goto Label58
			// line 297: if not IS_PYPY:
			πF.SetLineno(297)
		Label57:
			// line 298: d['MethodDescriptorType'] = type.__dict__['mro']
			πF.SetLineno(298)
			πTemp005 = ßmro.ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp013 = ßMethodDescriptorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 299: d['WrapperDescriptorType'] = type.__repr__
			πF.SetLineno(299)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß__repr__, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp013 = ßWrapperDescriptorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 300: a['WrapperDescriptorType2'] = type.__dict__['__module__']
			πF.SetLineno(300)
			πTemp005 = ß__module__.ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßWrapperDescriptorType2.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 301: d['ClassMethodDescriptorType'] = type.__dict__['__prepare__' if PY3 else 'mro']
			πF.SetLineno(301)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp012); πE != nil {
				continue
			}
			if !πTemp006 {
				goto Label59
			}
			πTemp011 = ß__prepare__.ToObject()
			goto Label60
		Label59:
			πTemp011 = ßmro.ToObject()
		Label60:
			πTemp005 = πTemp011
			if πTemp012, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp013 = ßClassMethodDescriptorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			goto Label58
		Label58:
			if πTemp011, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			πTemp005 = πTemp011
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label61
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			πTemp005 = πTemp011
		Label61:
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label62
			}
			goto Label63
			// line 303: if PY3 or IS_PYPY:
			πF.SetLineno(303)
		Label62:
			// line 304: _methodwrap = (1).__lt__
			πF.SetLineno(304)
			if πTemp005, πE = πg.GetAttr(πF, πg.NewInt(1).ToObject(), ß__lt__, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_methodwrap.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label64
		Label63:
			// line 306: _methodwrap = (1).__cmp__
			πF.SetLineno(306)
			if πTemp005, πE = πg.GetAttr(πF, πg.NewInt(1).ToObject(), ß__cmp__, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_methodwrap.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label64
		Label64:
			// line 307: d['MethodWrapperType'] = _methodwrap
			πF.SetLineno(307)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_methodwrap); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp013 = ßMethodWrapperType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 308: a['StaticMethodType'] = staticmethod(_method)
			πF.SetLineno(308)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_method); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßstaticmethod); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßStaticMethodType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 309: a['ClassMethodType'] = classmethod(_method)
			πF.SetLineno(309)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_method); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßclassmethod); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßClassMethodType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 310: a['PropertyType'] = property()
			πF.SetLineno(310)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßproperty); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßPropertyType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 311: d['SuperType'] = super(Exception, _exception)
			πF.SetLineno(311)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_exception); πE != nil {
				continue
			}
			πTemp001[1] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp013 = ßSuperType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label65
			}
			goto Label66
			// line 313: if PY3:
			πF.SetLineno(313)
		Label65:
			// line 314: _in = _bytes
			πF.SetLineno(314)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_bytes); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_in.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label67
		Label66:
			// line 316: _in = _str
			πF.SetLineno(316)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_str); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_in.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label67
		Label67:
			// line 317: a['InputType'] = _cstrI = StringIO(_in)
			πF.SetLineno(317)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_in); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßInputType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_cstrI.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 318: a['OutputType'] = _cstrO = StringIO()
			πF.SetLineno(318)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßOutputType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_cstrO.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 320: a['WeakKeyDictionaryType'] = weakref.WeakKeyDictionary()
			πF.SetLineno(320)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßWeakKeyDictionary, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßWeakKeyDictionaryType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 321: a['WeakValueDictionaryType'] = weakref.WeakValueDictionary()
			πF.SetLineno(321)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßWeakValueDictionary, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßWeakValueDictionaryType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 322: a['ReferenceType'] = weakref.ref(_instance)
			πF.SetLineno(322)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_instance); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßref, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßReferenceType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 323: a['DeadReferenceType'] = weakref.ref(_class())
			πF.SetLineno(323)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_class); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßref, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDeadReferenceType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 324: a['ProxyType'] = weakref.proxy(_instance)
			πF.SetLineno(324)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_instance); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßproxy, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßProxyType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 325: a['DeadProxyType'] = weakref.proxy(_class())
			πF.SetLineno(325)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_class); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßproxy, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDeadProxyType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 326: a['CallableProxyType'] = weakref.proxy(_instance2)
			πF.SetLineno(326)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_instance2); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßproxy, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCallableProxyType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 327: a['DeadCallableProxyType'] = weakref.proxy(_class2())
			πF.SetLineno(327)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_class2); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßproxy, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDeadCallableProxyType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 328: a['QueueType'] = Queue.Queue()
			πF.SetLineno(328)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßQueue); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßQueue, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßQueueType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 330: d['PartialType'] = functools.partial(int,base=2)
			πF.SetLineno(330)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			πTemp003 = πg.KWArgs{
				{"base", πg.NewInt(2).ToObject()},
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßfunctools); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßpartial, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, πTemp003); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp013 = ßPartialType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label68
			}
			goto Label69
			// line 331: if PY3:
			πF.SetLineno(331)
		Label68:
			// line 332: a['IzipType'] = zip('0','1')
			πF.SetLineno(332)
			πTemp001 = πF.MakeArgs(2)
			πTemp001[0] = ß0.ToObject()
			πTemp001[1] = ß1.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßIzipType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			goto Label70
		Label69:
			// line 334: a['IzipType'] = itertools.izip('0','1')
			πF.SetLineno(334)
			πTemp001 = πF.MakeArgs(2)
			πTemp001[0] = ß0.ToObject()
			πTemp001[1] = ß1.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßizip, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßIzipType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label70
		Label70:
			// line 335: a['ChainType'] = itertools.chain('0','1')
			πF.SetLineno(335)
			πTemp001 = πF.MakeArgs(2)
			πTemp001[0] = ß0.ToObject()
			πTemp001[1] = ß1.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßchain, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßChainType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 336: d['ItemGetterType'] = operator.itemgetter(0)
			πF.SetLineno(336)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(0).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßitemgetter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp013 = ßItemGetterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 337: d['AttrGetterType'] = operator.attrgetter('__repr__')
			πF.SetLineno(337)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ß__repr__.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßattrgetter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp013 = ßAttrGetterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label71
			}
			goto Label72
			// line 339: if PY3: _fileW = _cstrO
			πF.SetLineno(339)
		Label71:
			// line 339: if PY3: _fileW = _cstrO
			πF.SetLineno(339)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_cstrO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_fileW.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label73
		Label72:
			// line 340: else: _fileW = _tmpf
			πF.SetLineno(340)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_tmpf); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_fileW.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label73
		Label73:
			if πTemp005, πE = πg.ResolveGlobal(πF, ßHAS_ALL); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label74
			}
			goto Label75
			// line 342: if HAS_ALL:
			πF.SetLineno(342)
		Label74:
			// line 343: a['ConnectionType'] = _conn = sqlite3.connect(':memory:')
			πF.SetLineno(343)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewStr(":memory:").ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsqlite3); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßconnect, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßConnectionType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_conn.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 344: a['CursorType'] = _conn.cursor()
			πF.SetLineno(344)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_conn); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcursor, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCursorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label75
		Label75:
			// line 345: a['ShelveType'] = shelve.Shelf({})
			πF.SetLineno(345)
			πTemp001 = πF.MakeArgs(1)
			πTemp009 = πg.NewDict()
			πTemp005 = πTemp009.ToObject()
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßshelve); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßShelf, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßShelveType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßHAS_ALL); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label76
			}
			goto Label77
			// line 347: if HAS_ALL:
			πF.SetLineno(347)
		Label76:
			πTemp001 = πF.MakeArgs(1)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßhexversion, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp013
			if πTemp012, πE = πg.ResolveGlobal(πF, ßhex); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp011, πE = πg.LT(πF, πTemp013, ß0x2070ef0.ToObject()); πE != nil {
				continue
			}
			πTemp005 = πTemp011
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label78
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			πTemp005 = πTemp011
		Label78:
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label79
			}
			goto Label80
			// line 348: if (hex(sys.hexversion) < '0x2070ef0') or PY3:
			πF.SetLineno(348)
		Label79:
			// line 349: a['BZ2FileType'] = bz2.BZ2File(os.devnull) #FIXME: fail >= 3.3, 2.7.14
			πF.SetLineno(349)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßbz2); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßBZ2File, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBZ2FileType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label80
		Label80:
			// line 350: a['BZ2CompressorType'] = bz2.BZ2Compressor()
			πF.SetLineno(350)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßbz2); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßBZ2Compressor, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBZ2CompressorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 351: a['BZ2DecompressorType'] = bz2.BZ2Decompressor()
			πF.SetLineno(351)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßbz2); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßBZ2Decompressor, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßBZ2DecompressorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label77
		Label77:
			// line 355: a['TarFileType'] = tarfile.open(fileobj=_fileW,mode='w')
			πF.SetLineno(355)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_fileW); πE != nil {
				continue
			}
			πTemp003 = πg.KWArgs{
				{"fileobj", πTemp005},
				{"mode", ßw.ToObject()},
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtarfile); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßopen, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, πTemp003); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßTarFileType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 357: a['DialectType'] = csv.get_dialect('excel')
			πF.SetLineno(357)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßexcel.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcsv); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßget_dialect, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßDialectType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 358: a['PackerType'] = xdrlib.Packer()
			πF.SetLineno(358)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßxdrlib); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßPacker, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßPackerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 360: a['LockType'] = threading.Lock()
			πF.SetLineno(360)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßLock, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßLockType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 361: a['RLockType'] = threading.RLock()
			πF.SetLineno(361)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßRLock, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßRLockType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 363: a['NamedLoggerType'] = _logger = logging.getLogger(__name__) #FIXME: fail >= 3.2 and <= 2.6
			πF.SetLineno(363)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßgetLogger, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßNamedLoggerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_logger.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label81
			}
			goto Label82
			// line 366: if PY3:
			πF.SetLineno(366)
		Label81:
			// line 367: a['SocketType'] = _socket = socket.socket() #FIXME: fail >= 3.3
			πF.SetLineno(367)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsocket); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßsocket, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßSocketType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_socket.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 368: a['SocketPairType'] = socket.socketpair()[0] #FIXME: fail >= 3.3
			πF.SetLineno(368)
			πTemp005 = πg.NewInt(0).ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßsocket); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßsocketpair, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp013.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßSocketPairType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			goto Label83
		Label82:
			// line 370: a['SocketType'] = _socket = socket.socket()
			πF.SetLineno(370)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsocket); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßsocket, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßSocketType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_socket.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 371: a['SocketPairType'] = _socket._sock
			πF.SetLineno(371)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_socket); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß_sock, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßSocketPairType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			goto Label83
		Label83:
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label84
			}
			goto Label85
			// line 373: if PY3:
			πF.SetLineno(373)
		Label84:
			// line 374: a['GeneratorContextManagerType'] = contextlib.contextmanager(max)([1])
			πF.SetLineno(374)
			πTemp001 = πF.MakeArgs(1)
			πTemp014 = make([]*πg.Object, 1)
			πTemp014[0] = πg.NewInt(1).ToObject()
			πTemp005 = πg.NewList(πTemp014...).ToObject()
			πTemp001[0] = πTemp005
			πTemp014 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
				continue
			}
			πTemp014[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcontextlib); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcontextmanager, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp014, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp014)
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßGeneratorContextManagerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			goto Label86
		Label85:
			// line 376: a['GeneratorContextManagerType'] = contextlib.GeneratorContextManager(max)
			πF.SetLineno(376)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcontextlib); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßGeneratorContextManager, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßGeneratorContextManagerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label86
		Label86:
			// line 378: try: # ipython
			πF.SetLineno(378)
			πF.PushCheckpoint(88)
			// line 379: __IPYTHON__ is True # is ipython
			πF.SetLineno(379)
			if πTemp011, πE = πg.ResolveGlobal(πF, ß__IPYTHON__); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp005 = πg.GetBool(πTemp011 == πTemp012).ToObject()
			πF.PopCheckpoint()
			goto Label87
		Label88:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label89
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 380: except NameError:
			πF.SetLineno(380)
		Label89:
			// line 382: a['QuitterType'] = quit
			πF.SetLineno(382)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßquit); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßQuitterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 383: d['ExitType'] = a['QuitterType']
			πF.SetLineno(383)
			πTemp005 = ßQuitterType.ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp013 = ßExitType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label87
		Label87:
			// line 384: try: # numpy #FIXME: slow... 0.05 to 0.1 sec to import numpy
			πF.SetLineno(384)
			πF.PushCheckpoint(91)
			// line 385: from numpy import ufunc as _numpy_ufunc
			πF.SetLineno(385)
			if πTemp001, πE = πg.ImportModule(πF, "numpy"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πTemp011, πE = πg.GetAttrImport(πF, πTemp005, ßufunc); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_numpy_ufunc.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 386: from numpy import array as _numpy_array
			πF.SetLineno(386)
			if πTemp001, πE = πg.ImportModule(πF, "numpy"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πTemp011, πE = πg.GetAttrImport(πF, πTemp005, ßarray); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_numpy_array.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 387: from numpy import int32 as _numpy_int32
			πF.SetLineno(387)
			if πTemp001, πE = πg.ImportModule(πF, "numpy"); πE != nil {
				continue
			}
			πTemp005 = πTemp001[0]
			if πTemp011, πE = πg.GetAttrImport(πF, πTemp005, ßint32); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_numpy_int32.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 388: a['NumpyUfuncType'] = _numpy_ufunc
			πF.SetLineno(388)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_numpy_ufunc); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßNumpyUfuncType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 389: a['NumpyArrayType'] = _numpy_array
			πF.SetLineno(389)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_numpy_array); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßNumpyArrayType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 390: a['NumpyInt32Type'] = _numpy_int32
			πF.SetLineno(390)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_numpy_int32); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßNumpyInt32Type.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label90
		Label91:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label92
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 391: except ImportError:
			πF.SetLineno(391)
		Label92:
			// line 392: pass
			πF.SetLineno(392)
			πF.RestoreExc(nil, nil)
			goto Label90
		Label90:
			// line 393: try: # python 2.6
			πF.SetLineno(393)
			πF.PushCheckpoint(94)
			// line 395: a['ProductType'] = itertools.product('0','1')
			πF.SetLineno(395)
			πTemp001 = πF.MakeArgs(2)
			πTemp001[0] = ß0.ToObject()
			πTemp001[1] = ß1.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßproduct, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßProductType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 397: a['FileHandlerType'] = logging.FileHandler(os.devnull) #FIXME: fail >= 3.2 and <= 2.6
			πF.SetLineno(397)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßFileHandler, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßFileHandlerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 398: a['RotatingFileHandlerType'] = logging.handlers.RotatingFileHandler(os.devnull)
			πF.SetLineno(398)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßhandlers, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßRotatingFileHandler, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßRotatingFileHandlerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 399: a['SocketHandlerType'] = logging.handlers.SocketHandler('localhost',514)
			πF.SetLineno(399)
			πTemp001 = πF.MakeArgs(2)
			πTemp001[0] = ßlocalhost.ToObject()
			πTemp001[1] = πg.NewInt(514).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßhandlers, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßSocketHandler, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßSocketHandlerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 400: a['MemoryHandlerType'] = logging.handlers.MemoryHandler(1)
			πF.SetLineno(400)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßhandlers, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßMemoryHandler, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßMemoryHandlerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label93
		Label94:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label95
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 401: except AttributeError:
			πF.SetLineno(401)
		Label95:
			// line 402: pass
			πF.SetLineno(402)
			πF.RestoreExc(nil, nil)
			goto Label93
		Label93:
			// line 403: try: # python 2.7
			πF.SetLineno(403)
			πF.PushCheckpoint(97)
			// line 405: a['WeakSetType'] = weakref.WeakSet() # 2.7
			πF.SetLineno(405)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßWeakSet, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßWeakSetType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label96
		Label97:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label98
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 412: except AttributeError:
			πF.SetLineno(412)
		Label98:
			// line 413: pass
			πF.SetLineno(413)
			πF.RestoreExc(nil, nil)
			goto Label96
		Label96:
			// line 417: a['FileType'] = open(os.devnull, 'rb', buffering=0) # same 'wb','wb+','rb+'
			πF.SetLineno(417)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdevnull, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp001[1] = ßrb.ToObject()
			πTemp003 = πg.KWArgs{
				{"buffering", πg.NewInt(0).ToObject()},
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, πTemp003); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßFileType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 420: a['ListIteratorType'] = iter(_list) # empty vs non-empty FIXME: fail < 3.2
			πF.SetLineno(420)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_list); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßListIteratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 421: a['TupleIteratorType']= iter(_tuple) # empty vs non-empty FIXME: fail < 3.2
			πF.SetLineno(421)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_tuple); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßTupleIteratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 422: a['XRangeIteratorType'] = iter(_xrange) # empty vs non-empty FIXME: fail < 3.2
			πF.SetLineno(422)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_xrange); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßXRangeIteratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 424: a['PrettyPrinterType'] = pprint.PrettyPrinter() #FIXME: fail >= 3.2 and == 2.5
			πF.SetLineno(424)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßpprint); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßPrettyPrinter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßPrettyPrinterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 426: a['CycleType'] = itertools.cycle('0') #FIXME: fail < 3.2
			πF.SetLineno(426)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ß0.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcycle, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCycleType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 428: a['TemporaryFileType'] = _tmpf #FIXME: fail >= 3.2 and == 2.5
			πF.SetLineno(428)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_tmpf); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßTemporaryFileType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 430: a['GzipFileType'] = gzip.GzipFile(fileobj=_fileW) #FIXME: fail > 3.2 and <= 2.6
			πF.SetLineno(430)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_fileW); πE != nil {
				continue
			}
			πTemp003 = πg.KWArgs{
				{"fileobj", πTemp005},
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßgzip); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßGzipFile, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, πTemp003); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßGzipFileType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 432: a['StreamHandlerType'] = logging.StreamHandler() #FIXME: fail >= 3.2 and == 2.5
			πF.SetLineno(432)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßStreamHandler, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßStreamHandlerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 433: try: # python 2.6
			πF.SetLineno(433)
			πF.PushCheckpoint(100)
			// line 435: a['PermutationsType'] = itertools.permutations('0') #FIXME: fail < 3.2
			πF.SetLineno(435)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ß0.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßpermutations, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßPermutationsType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 436: a['CombinationsType'] = itertools.combinations('0',1) #FIXME: fail < 3.2
			πF.SetLineno(436)
			πTemp001 = πF.MakeArgs(2)
			πTemp001[0] = ß0.ToObject()
			πTemp001[1] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcombinations, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCombinationsType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label99
		Label100:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label101
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 437: except AttributeError:
			πF.SetLineno(437)
		Label101:
			// line 438: pass
			πF.SetLineno(438)
			πF.RestoreExc(nil, nil)
			goto Label99
		Label99:
			// line 439: try: # python 2.7
			πF.SetLineno(439)
			πF.PushCheckpoint(103)
			// line 441: a['RepeatType'] = itertools.repeat(0) #FIXME: fail < 3.2
			πF.SetLineno(441)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(0).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßrepeat, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßRepeatType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 442: a['CompressType'] = itertools.compress('0',[1]) #FIXME: fail < 3.2
			πF.SetLineno(442)
			πTemp001 = πF.MakeArgs(2)
			πTemp001[0] = ß0.ToObject()
			πTemp014 = make([]*πg.Object, 1)
			πTemp014[0] = πg.NewInt(1).ToObject()
			πTemp005 = πg.NewList(πTemp014...).ToObject()
			πTemp001[1] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcompress, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			πTemp013 = ßCompressType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label102
		Label103:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label104
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 444: except AttributeError:
			πF.SetLineno(444)
		Label104:
			// line 445: pass
			πF.SetLineno(445)
			πF.RestoreExc(nil, nil)
			goto Label102
		Label102:
			// line 449: x['GeneratorType'] = _generator = _function(1) #XXX: priority
			πF.SetLineno(449)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_function); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßGeneratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_generator.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 450: x['FrameType'] = _generator.gi_frame #XXX: inspect.currentframe()
			πF.SetLineno(450)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_generator); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßgi_frame, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßFrameType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 451: x['TracebackType'] = _function2()[1] #(see: inspect.getouterframes,getframeinfo)
			πF.SetLineno(451)
			πTemp005 = πg.NewInt(1).ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_function2); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßTracebackType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 455: x['SetIteratorType'] = iter(_set) #XXX: empty vs non-empty
			πF.SetLineno(455)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_set); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßSetIteratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label105
			}
			goto Label106
			// line 457: if PY3:
			πF.SetLineno(457)
		Label105:
			// line 458: x['DictionaryItemIteratorType'] = iter(type.__dict__.items())
			πF.SetLineno(458)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßitems, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictionaryItemIteratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 459: x['DictionaryKeyIteratorType'] = iter(type.__dict__.keys())
			πF.SetLineno(459)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßkeys, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictionaryKeyIteratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 460: x['DictionaryValueIteratorType'] = iter(type.__dict__.values())
			πF.SetLineno(460)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßvalues, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictionaryValueIteratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			goto Label107
		Label106:
			// line 462: x['DictionaryItemIteratorType'] = type.__dict__.iteritems()
			πF.SetLineno(462)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßiteritems, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictionaryItemIteratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 463: x['DictionaryKeyIteratorType'] = type.__dict__.iterkeys()
			πF.SetLineno(463)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßiterkeys, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictionaryKeyIteratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 464: x['DictionaryValueIteratorType'] = type.__dict__.itervalues()
			πF.SetLineno(464)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßitervalues, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictionaryValueIteratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			goto Label107
		Label107:
			// line 466: x['StructType'] = struct.Struct('c')
			πF.SetLineno(466)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßc.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßstruct); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßStruct, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßStructType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 467: x['CallableIteratorType'] = _srepattern.finditer('')
			πF.SetLineno(467)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ß.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_srepattern); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßfinditer, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCallableIteratorType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 468: x['SREMatchType'] = _srepattern.match('')
			πF.SetLineno(468)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ß.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_srepattern); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßmatch, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßSREMatchType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 469: x['SREScannerType'] = _srepattern.scanner('')
			πF.SetLineno(469)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ß.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_srepattern); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßscanner, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßSREScannerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 470: x['StreamReader'] = codecs.StreamReader(_cstrI) #XXX: ... and etc
			πF.SetLineno(470)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_cstrI); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcodecs); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßStreamReader, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßStreamReader.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßHAS_ALL); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label108
			}
			goto Label109
			// line 473: if HAS_ALL:
			πF.SetLineno(473)
		Label108:
			// line 474: x['DbmType'] = dbm.open(_tempfile,'n')
			πF.SetLineno(474)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_tempfile); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			πTemp001[1] = ßn.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdbm); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßopen, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDbmType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label109
		Label109:
			// line 478: x['ZlibCompressType'] = zlib.compressobj()
			πF.SetLineno(478)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßzlib); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcompressobj, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßZlibCompressType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 479: x['ZlibDecompressType'] = zlib.decompressobj()
			πF.SetLineno(479)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßzlib); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßdecompressobj, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßZlibDecompressType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 481: x['CSVReaderType'] = csv.reader(_cstrI)
			πF.SetLineno(481)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_cstrI); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcsv); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßreader, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCSVReaderType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 482: x['CSVWriterType'] = csv.writer(_cstrO)
			πF.SetLineno(482)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_cstrO); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcsv); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßwriter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCSVWriterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 483: x['CSVDictReaderType'] = csv.DictReader(_cstrI)
			πF.SetLineno(483)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_cstrI); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcsv); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßDictReader, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCSVDictReaderType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 484: x['CSVDictWriterType'] = csv.DictWriter(_cstrO,{})
			πF.SetLineno(484)
			πTemp001 = πF.MakeArgs(2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_cstrO); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			πTemp009 = πg.NewDict()
			πTemp005 = πTemp009.ToObject()
			πTemp001[1] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcsv); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßDictWriter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCSVDictWriterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 486: x['HashType'] = hashlib.md5()
			πF.SetLineno(486)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßhashlib); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßmd5, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßHashType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			πTemp001 = πF.MakeArgs(1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßhexversion, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp012
			if πTemp011, πE = πg.ResolveGlobal(πF, ßhex); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp005, πE = πg.LT(πF, πTemp012, ß0x30800a1.ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label110
			}
			goto Label111
			// line 487: if (hex(sys.hexversion) < '0x30800a1'):
			πF.SetLineno(487)
		Label110:
			// line 488: x['HMACType'] = hmac.new(_in)
			πF.SetLineno(488)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_in); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßhmac); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßnew, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßHMACType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label112
		Label111:
			// line 490: x['HMACType'] = hmac.new(_in, digestmod='md5')
			πF.SetLineno(490)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_in); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			πTemp003 = πg.KWArgs{
				{"digestmod", ßmd5.ToObject()},
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßhmac); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßnew, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, πTemp003); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßHMACType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label112
		Label112:
			if πTemp005, πE = πg.ResolveGlobal(πF, ßHAS_CURSES); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label113
			}
			goto Label114
			// line 492: if HAS_CURSES: pass
			πF.SetLineno(492)
		Label113:
			// line 492: if HAS_CURSES: pass
			πF.SetLineno(492)
			goto Label114
		Label114:
			if πTemp005, πE = πg.ResolveGlobal(πF, ßHAS_CTYPES); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label115
			}
			goto Label116
			// line 496: if HAS_CTYPES:
			πF.SetLineno(496)
		Label115:
			// line 497: x['CCharPType'] = ctypes.c_char_p()
			πF.SetLineno(497)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_char_p, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCCharPType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 498: x['CWCharPType'] = ctypes.c_wchar_p()
			πF.SetLineno(498)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_wchar_p, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCWCharPType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 499: x['CVoidPType'] = ctypes.c_void_p()
			πF.SetLineno(499)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_void_p, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCVoidPType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πTemp011, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(3).ToObject(), πg.None}, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp013, ßplatform, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetItem(πF, πTemp016, πTemp011); πE != nil {
				continue
			}
			if πTemp005, πE = πg.Eq(πF, πTemp012, ßwin.ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label117
			}
			goto Label118
			// line 500: if sys.platform[:3] == 'win':
			πF.SetLineno(500)
		Label117:
			// line 501: x['CDLLType'] = _cdll = ctypes.cdll.msvcrt
			πF.SetLineno(501)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcdll, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp011, ßmsvcrt, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCDLLType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_cdll.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label119
		Label118:
			// line 503: x['CDLLType'] = _cdll = ctypes.CDLL(None)
			πF.SetLineno(503)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßCDLL, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCDLLType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_cdll.ToObject(), πTemp005); πE != nil {
				continue
			}
			goto Label119
		Label119:
			if πTemp011, πE = πg.ResolveGlobal(πF, ßIS_PYPY); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp011); πE != nil {
				continue
			}
			πTemp005 = πg.GetBool(!πTemp006).ToObject()
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label120
			}
			goto Label121
			// line 504: if not IS_PYPY:
			πF.SetLineno(504)
		Label120:
			// line 505: x['PyDLLType'] = _pydll = ctypes.pythonapi
			πF.SetLineno(505)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßpythonapi, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßPyDLLType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_pydll.ToObject(), πTemp011); πE != nil {
				continue
			}
			goto Label121
		Label121:
			// line 506: x['FuncPtrType'] = _cdll._FuncPtr()
			πF.SetLineno(506)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_cdll); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß_FuncPtr, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßFuncPtrType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 507: x['CCharArrayType'] = ctypes.create_string_buffer(1)
			πF.SetLineno(507)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcreate_string_buffer, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCCharArrayType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 508: x['CWCharArrayType'] = ctypes.create_unicode_buffer(1)
			πF.SetLineno(508)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(1).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcreate_unicode_buffer, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCWCharArrayType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 509: x['CParamType'] = ctypes.byref(_cchar)
			πF.SetLineno(509)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_cchar); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßbyref, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCParamType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 510: x['LPCCharType'] = ctypes.pointer(_cchar)
			πF.SetLineno(510)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_cchar); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßpointer, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßLPCCharType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 511: x['LPCCharObjType'] = _lpchar = ctypes.POINTER(ctypes.c_char)
			πF.SetLineno(511)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_char, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßPOINTER, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßLPCCharObjType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_lpchar.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 512: x['NullPtrType'] = _lpchar()
			πF.SetLineno(512)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_lpchar); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßNullPtrType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 513: x['NullPyObjectType'] = ctypes.py_object()
			πF.SetLineno(513)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßpy_object, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßNullPyObjectType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 514: x['PyObjectType'] = ctypes.py_object(lambda :None)
			πF.SetLineno(514)
			πTemp001 = πF.MakeArgs(1)
			πTemp010 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("<lambda>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/_objects.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 514: x['PyObjectType'] = ctypes.py_object(lambda :None)
					πF.SetLineno(514)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßpy_object, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßPyObjectType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 515: x['FieldType'] = _field = _Struct._field
			πF.SetLineno(515)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_Struct); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ß_field, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßFieldType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_field.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 516: x['CFUNCTYPEType'] = _cfunc = ctypes.CFUNCTYPE(ctypes.c_char)
			πF.SetLineno(516)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßc_char, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßCFUNCTYPE, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCFUNCTYPEType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_cfunc.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 517: x['CFunctionType'] = _cfunc(str)
			πF.SetLineno(517)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_cfunc); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCFunctionType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			goto Label116
		Label116:
			// line 518: try: # python 2.6
			πF.SetLineno(518)
			πF.PushCheckpoint(123)
			// line 520: x['MethodCallerType'] = operator.methodcaller('mro') # 2.6
			πF.SetLineno(520)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßmro.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßmethodcaller, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßMethodCallerType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label122
		Label123:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label124
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 521: except AttributeError:
			πF.SetLineno(521)
		Label124:
			// line 522: pass
			πF.SetLineno(522)
			πF.RestoreExc(nil, nil)
			goto Label122
		Label122:
			// line 523: try: # python 2.7
			πF.SetLineno(523)
			πF.PushCheckpoint(126)
			// line 525: x['MemoryType'] = memoryview(_in) # 2.7
			πF.SetLineno(525)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_in); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßmemoryview); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßMemoryType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			// line 526: x['MemoryType2'] = memoryview(bytearray(_in)) # 2.7
			πF.SetLineno(526)
			πTemp001 = πF.MakeArgs(1)
			πTemp014 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_in); πE != nil {
				continue
			}
			πTemp014[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßbytearray); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp014, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp014)
			πTemp001[0] = πTemp011
			if πTemp005, πE = πg.ResolveGlobal(πF, ßmemoryview); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßMemoryType2.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label127
			}
			goto Label128
			// line 527: if PY3:
			πF.SetLineno(527)
		Label127:
			// line 528: x['DictItemsType'] = _dict.items() # 2.7
			πF.SetLineno(528)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_dict); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßitems, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictItemsType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 529: x['DictKeysType'] = _dict.keys() # 2.7
			πF.SetLineno(529)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_dict); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßkeys, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictKeysType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 530: x['DictValuesType'] = _dict.values() # 2.7
			πF.SetLineno(530)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_dict); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßvalues, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictValuesType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label129
		Label128:
			// line 532: x['DictItemsType'] = _dict.viewitems() # 2.7
			πF.SetLineno(532)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_dict); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßviewitems, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictItemsType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 533: x['DictKeysType'] = _dict.viewkeys() # 2.7
			πF.SetLineno(533)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_dict); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßviewkeys, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictKeysType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 534: x['DictValuesType'] = _dict.viewvalues() # 2.7
			πF.SetLineno(534)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_dict); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßviewvalues, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßDictValuesType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			goto Label129
		Label129:
			// line 536: x['RawTextHelpFormatterType'] = argparse.RawTextHelpFormatter('PROG')
			πF.SetLineno(536)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßPROG.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßargparse); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßRawTextHelpFormatter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßRawTextHelpFormatterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 537: x['RawDescriptionHelpFormatterType'] = argparse.RawDescriptionHelpFormatter('PROG')
			πF.SetLineno(537)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßPROG.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßargparse); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßRawDescriptionHelpFormatter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßRawDescriptionHelpFormatterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			// line 538: x['ArgDefaultsHelpFormatterType'] = argparse.ArgumentDefaultsHelpFormatter('PROG')
			πF.SetLineno(538)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßPROG.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßargparse); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßArgumentDefaultsHelpFormatter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßArgDefaultsHelpFormatterType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label125
		Label126:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label130
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 539: except NameError:
			πF.SetLineno(539)
		Label130:
			// line 540: pass
			πF.SetLineno(540)
			πF.RestoreExc(nil, nil)
			goto Label125
		Label125:
			// line 541: try: # python 2.7 (and not 3.1)
			πF.SetLineno(541)
			πF.PushCheckpoint(132)
			// line 542: x['CmpKeyType'] = _cmpkey = functools.cmp_to_key(_methodwrap) # 2.7, >=3.2
			πF.SetLineno(542)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_methodwrap); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßfunctools); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßcmp_to_key, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp005); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCmpKeyType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_cmpkey.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 543: x['CmpKeyObjType'] = _cmpkey('0') #2.7, >=3.2
			πF.SetLineno(543)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ß0.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_cmpkey); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßCmpKeyObjType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label131
		Label132:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label133
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 544: except AttributeError:
			πF.SetLineno(544)
		Label133:
			// line 545: pass
			πF.SetLineno(545)
			πF.RestoreExc(nil, nil)
			goto Label131
		Label131:
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPY3); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label134
			}
			goto Label135
			// line 546: if PY3: # oddities: removed, etc
			πF.SetLineno(546)
		Label134:
			// line 547: x['BufferType'] = x['MemoryType']
			πF.SetLineno(547)
			πTemp005 = ßMemoryType.ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp005); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßBufferType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			goto Label136
		Label135:
			// line 549: x['BufferType'] = buffer('')
			πF.SetLineno(549)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ß.ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßbuffer); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp011); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
				continue
			}
			πTemp013 = ßBufferType.ToObject()
			if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp005); πE != nil {
				continue
			}
			goto Label136
		Label136:
			// line 552: a.update(d) # registered also succeed
			πF.SetLineno(552)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßupdate, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp011, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(3).ToObject(), πg.None}, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp013, ßplatform, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetItem(πF, πTemp016, πTemp011); πE != nil {
				continue
			}
			if πTemp005, πE = πg.Eq(πF, πTemp012, ßwin.ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label137
			}
			goto Label138
			// line 553: if sys.platform[:3] == 'win':
			πF.SetLineno(553)
		Label137:
			// line 554: os.close(_filedescrip) # required on win32
			πF.SetLineno(554)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_filedescrip); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßclose, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			goto Label138
		Label138:
			// line 555: os.remove(_tempfile)
			πF.SetLineno(555)
			πTemp001 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_tempfile); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßremove, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
		}
		return nil, πE
	})
	πg.RegisterModule("_objects", Code)
}

