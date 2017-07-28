ROOT = main.go parse_message.go
CONSTANTS = $(wildcard constants/*)
MODELS = $(wildcard models/*)
MODULES = $(wildcard modules/*)
UTIL = $(wildcard util/*)

chocobot: $(ROOT) $(CONSTANTS) $(MODELS) $(MODULES) $(UTIL)
	go build

clean:
	rm chocobot