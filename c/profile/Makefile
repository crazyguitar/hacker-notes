SRC = $(wildcard *.c)
OBJ = $(SRC:.c=.o)
EXE = $(subst .o,,$(OBJ))

SUB = lib
LIB = lib/libmisc.a

CFLAGS += -g -O2 -Wall -Werror -I$(SUB)

.PHONY: all clean $(SUB)

all: $(EXE) $(OBJ) $(SUB)

%: %.o $(SUB)
	$(CC) -o $@ $< $(LIB) $(LDFLAGS) 

$.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<

$(SUB):
	$(MAKE) -C $@ $(MAKECMDGOALS)

clean: $(SUB)
	rm -rf $(OBJ) $(EXE)
