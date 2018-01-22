OS  = $(shell uname -s)

# linux only
EX  = \
arp_sniffer.go \
arp_spoofing.go \
ip_sniffer.go \
ping.go \
alg_hash_unix.go \
syscall_splice.go \
syscall_select.go

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
