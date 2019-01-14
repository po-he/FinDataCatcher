PROJ_DIR := $(shell pwd)

BIN_DIR = ${PROJ_DIR}/bin
BUILD_DIR = ${PROJ_DIR}/build

#GO=go
#GOPATH := $(shell pwd)

APPS = tx_catcher_test tx_catcher_main
ALL: $(APPS)

$(BUILD_DIR)/tx_catcher_main:	$(wildcard apps/tx_catcher_main/*.go tx_catcher/*.go utils/*/*.go base/*/*.go)
$(BUILD_DIR)/tx_catcher_test:	$(wildcard apps/tx_catcher_test/*.go tx_catcher/*.go utils/*/*.go)


$(BUILD_DIR)/%:
	@mkdir -p $(dir $@)
	go build ${GOFLAG} -o $@ ./apps/$*

$(APPS): %: $(BUILD_DIR)/%

clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(BIN_DIR)

.PHONY: install clean ALL 
.PHONY: $(APPS)

install: $(APPS)
	install -m 755 -d ${BIN_DIR}
	for APP in $^ ; do install -m 755 ${BUILD_DIR}/$$APP ${BIN_DIR}/$$APP ; done 

all: clean ALL install



