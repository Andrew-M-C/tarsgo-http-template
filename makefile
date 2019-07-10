APP       := amc
TARGET    := HttpSvr
MFLAGS    :=
DFLAGS    :=
CONFIG    :=
STRIP_FLAG:= N
J2GO_FLAG:=

GOPATH ?= $(HOME)/go
GOROOT ?= /usr/local/go
libpath=${subst :, ,$(GOPATH)}
$(foreach path,$(libpath),$(eval -include $(path)/src/github.com/TarsCloud/TarsGo/tars/makefile.tars))

TARS_SRS_LONG := $(wildcard tars/*.tars) $(wildcard tars/*.jce)
TARS_SRS = $(subst tars/,, $(TARS_SRS_LONG))
PROXY_DIR = $(PWD)/vendor/proxy/

TARSBUILD:
	-@rm -rf $(PROXY_DIR) 2>/dev/null
	@mkdir -p $(PROXY_DIR); cd tars && $(TARS2GO) -outdir=$(PROXY_DIR) $(TARS_SRS)

$(warning Please ignore the warning above)
