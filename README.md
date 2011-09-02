Errforce
========

Utilities for working effectively with Go errors and UNIX errnos
----------------------------------------------------------------

Go's flat namespace and relatively primitive type assertions make it tedious to work with Go's standard `os.Error` and classic UNIX errnos.  This library bridges the gap.

Installation
------------

	goinstall github.com/rcrowley/errforce

Usage
-----

	import "github.com/rcrowley/errforce"

	func doIt() bool {
		err := doSomethingThatReturnsAnError()
		if errforce.IsENOENT(err) { return false }
		return true
	}
