SRC = $(wildcard *.c)
OBJ = $(SRC:.c=.o)
EXE = $(subst .o,,$(OBJ))

SUB = lib
LIB = lib/libmisc.a

CFLAGS += -g -O2 -Wall -Werror -I$(SUB)

.PHONY: all clean $(SUB)

all: $(SUB) $(OBJ) $(EXE)

$(SUB):
	$(MAKE) -C $@ $(MAKECMDGOALS)

%.o: %.c $(SUB)
	$(CC) $(CFLAGS) -c -o $@ $<

%: %.o $(SUB)
	$(CC) -o $@ $< $(LIB) $(LDFLAGS)

clean: $(SUB)
	rm -rf $(OBJ) $(EXE)
