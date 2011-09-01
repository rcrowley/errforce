include $(GOROOT)/src/Make.inc

TARG=errforce
GOFILES=errforce.go

include $(GOROOT)/src/Make.pkg

uninstall:
	rm -f $(GOROOT)/pkg/$(GOOS)_$(GOARCH)/$(TARG).a

.PHONY: uninstall
