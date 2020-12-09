package info
import (
	πg "github.com/pygolin/runtime"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/home/gmonroe/src/github.com/pygolin/pypi-dill/src/dill-0.3.3/dill/info.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß__doc__ := πg.InternStr("__doc__")
		ßlicense := πg.InternStr("license")
		ßreadme := πg.InternStr("readme")
		ßstable_version := πg.InternStr("stable_version")
		ßthis_version := πg.InternStr("this_version")
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 2: this_version = '0.3.3'
			πF.SetLineno(2)
			// line 2: this_version = '0.3.3'
			πF.SetLineno(2)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("0.3.3").ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßthis_version.ToObject(), πg.NewStr("0.3.3").ToObject()); πE != nil {
				continue
			}
			// line 3: stable_version = '0.3.3'
			πF.SetLineno(3)
			if πE = πF.Globals().SetItem(πF, ßstable_version.ToObject(), πg.NewStr("0.3.3").ToObject()); πE != nil {
				continue
			}
			// line 4: readme = '''-----------------------------
			πF.SetLineno(4)
			if πE = πF.Globals().SetItem(πF, ßreadme.ToObject(), πg.NewStr("-----------------------------\ndill: serialize all of python\n-----------------------------\n\nAbout Dill\n==========\n\n``dill`` extends python's ``pickle`` module for serializing and de-serializing\npython objects to the majority of the built-in python types. Serialization\nis the process of converting an object to a byte stream, and the inverse\nof which is converting a byte stream back to a python object hierarchy.\n\n``dill`` provides the user the same interface as the ``pickle`` module, and\nalso includes some additional features. In addition to pickling python\nobjects, ``dill`` provides the ability to save the state of an interpreter\nsession in a single command.  Hence, it would be feasable to save a\ninterpreter session, close the interpreter, ship the pickled file to\nanother computer, open a new interpreter, unpickle the session and\nthus continue from the 'saved' state of the original interpreter\nsession.\n\n``dill`` can be used to store python objects to a file, but the primary\nusage is to send python objects across the network as a byte stream.\n``dill`` is quite flexible, and allows arbitrary user defined classes\nand functions to be serialized.  Thus ``dill`` is not intended to be\nsecure against erroneously or maliciously constructed data. It is\nleft to the user to decide whether the data they unpickle is from\na trustworthy source.\n\n``dill`` is part of ``pathos``, a python framework for heterogeneous computing.\n``dill`` is in active development, so any user feedback, bug reports, comments,\nor suggestions are highly appreciated.  A list of issues is located at https://github.com/uqfoundation/dill/issues, with a legacy list maintained at https://uqfoundation.github.io/pathos-issues.html.\n\n\nMajor Features\n==============\n\n``dill`` can pickle the following standard types:\n\n    - none, type, bool, int, long, float, complex, str, unicode,\n    - tuple, list, dict, file, buffer, builtin,\n    - both old and new style classes,\n    - instances of old and new style classes,\n    - set, frozenset, array, functions, exceptions\n\n``dill`` can also pickle more 'exotic' standard types:\n\n    - functions with yields, nested functions, lambdas,\n    - cell, method, unboundmethod, module, code, methodwrapper,\n    - dictproxy, methoddescriptor, getsetdescriptor, memberdescriptor,\n    - wrapperdescriptor, xrange, slice,\n    - notimplemented, ellipsis, quit\n\n``dill`` cannot yet pickle these standard types:\n\n    - frame, generator, traceback\n\n``dill`` also provides the capability to:\n\n    - save and load python interpreter sessions\n    - save and extract the source code from functions and classes\n    - interactively diagnose pickling errors\n\n\nCurrent Release\n===============\n\nThis documentation is for version ``dill-0.3.3``.\n\nThe latest released version of ``dill`` is available from:\n\n    https://pypi.org/project/dill\n\n``dill`` is distributed under a 3-clause BSD license.\n\n    >>> import dill\n    >>> dill.license()\n\n\nDevelopment Version \n===================\n\nYou can get the latest development version with all the shiny new features at:\n\n    https://github.com/uqfoundation\n\nIf you have a new contribution, please submit a pull request.\n\n\nInstallation\n============\n\n``dill`` is packaged to install from source, so you must\ndownload the tarball, unzip, and run the installer::\n\n    [download]\n    $ tar -xvzf dill-0.3.3.tar.gz\n    $ cd dill-0.3.3\n    $ python setup py build\n    $ python setup py install\n\nYou will be warned of any missing dependencies and/or settings\nafter you run the \"build\" step above. \n\nAlternately, ``dill`` can be installed with ``pip`` or ``easy_install``::\n\n    $ pip install dill\n\n\nRequirements\n============\n\n``dill`` requires:\n\n    - ``python``, **version == 2.7** or **version >= 3.5**, or ``pypy``\n\nOptional requirements:\n\n    - ``setuptools``, **version >= 0.6**\n    - ``pyreadline``, **version >= 1.7.1** (on windows)\n    - ``objgraph``, **version >= 1.7.2**\n\n\nMore Information\n================\n\nProbably the best way to get started is to look at the documentation at\nhttp://dill.rtfd.io. Also see ``dill.tests`` for a set of scripts that\ndemonstrate how ``dill`` can serialize different python objects. You can\nrun the test suite with ``python -m dill.tests``. The contents of any\npickle file can be examined with ``undill``.  As ``dill`` conforms to\nthe ``pickle`` interface, the examples and documentation found at\nhttp://docs.python.org/library/pickle.html also apply to ``dill``\nif one will ``import dill as pickle``. The source code is also generally\nwell documented, so further questions may be resolved by inspecting the\ncode itself. Please feel free to submit a ticket on github, or ask a\nquestion on stackoverflow (**@Mike McKerns**).\nIf you would like to share how you use ``dill`` in your work, please send\nan email (to **mmckerns at uqfoundation dot org**).\n\n\nCitation\n========\n\nIf you use ``dill`` to do research that leads to publication, we ask that you\nacknowledge use of ``dill`` by citing the following in your publication::\n\n    M.M. McKerns, L. Strand, T. Sullivan, A. Fang, M.A.G. Aivazis,\n    \"Building a framework for predictive science\", Proceedings of\n    the 10th Python in Science Conference, 2011;\n    http://arxiv.org/pdf/1202.1056\n\n    Michael McKerns and Michael Aivazis,\n    \"pathos: a framework for heterogeneous computing\", 2010- ;\n    https://uqfoundation.github.io/pathos.html\n\nPlease see https://uqfoundation.github.io/pathos.html or\nhttp://arxiv.org/pdf/1202.1056 for further information.\n\n").ToObject()); πE != nil {
				continue
			}
			// line 164: license = '''Copyright (c) 2004-2016 California Institute of Technology.
			πF.SetLineno(164)
			if πE = πF.Globals().SetItem(πF, ßlicense.ToObject(), πg.NewStr("Copyright (c) 2004-2016 California Institute of Technology.\nCopyright (c) 2016-2020 The Uncertainty Quantification Foundation.\nAll rights reserved.\n\nThis software is available subject to the conditions and terms laid\nout below. By downloading and using this software you are agreeing\nto the following conditions.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions\nare met::\n\n    - Redistribution of source code must retain the above copyright\n      notice, this list of conditions and the following disclaimer.\n\n    - Redistribution in binary form must reproduce the above copyright\n      notice, this list of conditions and the following disclaimer in the\n      documentations and/or other materials provided with the distribution.\n\n    - Neither the names of the copyright holders nor the names of any of\n      the contributors may be used to endorse or promote products derived\n      from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS\n\"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED\nTO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR\nPURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR\nCONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,\nEXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,\nPROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;\nOR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,\nWHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR\nOTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF\nADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n\n").ToObject()); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("info", Code)
}

