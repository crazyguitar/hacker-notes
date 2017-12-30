OS  = $(shell uname -s)

EX  = arp_sniffer.go
ifeq ($(OS), Darwin)
SRC = $(filter-out $(EX), $(wildcard *.go))
else
SRC = $(wildcard *.go)
endif

EXE = $(subst .go,, $(SRC))

.PHONY: all clean

all: $(EXE)

%: %.go
	golint $^
	go build -o $@ $^

clean:
	rm -rf $(EXE)
