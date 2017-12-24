SRC = $(wildcard *.go)
EXE = $(subst .go,, $(SRC))

.PHONY: all clean

all: $(EXE)

%: %.go
	golint $^
	go build -o $@ $^

clean:
	rm -rf $(EXE)
