include $(GOROOT)/src/Make.inc

TARG=errforce
GOFILES=\
	errforce.go\
	errno.go\

include $(GOROOT)/src/Make.pkg

all: uninstall clean install
	@true

uninstall:
	rm -f $(GOROOT)/pkg/$(GOOS)_$(GOARCH)/$(TARG).a

errno.go:
	sh errno.sh >$@

.PHONY: all uninstall
