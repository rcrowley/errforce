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
	rm -f $(GOROOT)/pkg/$(GOOS)_$(GOARCH)/github.com/rcrowley/$(TARG).a
	rm -rf $(GOROOT)/src/pkg/github.com/rcrowley/$(TARG)

errno.go:
	sh errno.sh >$@

.PHONY: all uninstall
