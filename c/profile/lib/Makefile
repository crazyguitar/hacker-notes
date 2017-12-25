SRC = $(wildcard *.c)
OBJ = $(SRC:.c=.o)
LIB = libmisc.a

.PHONY: all clean

CFLAGS += -g -O2 -Wall -Werror -Wextra -DPROFILING

all: $(LIB) $(OBJ)

$(LIB): $(OBJ)
	ar rcs $@ $^

%.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $^

clean:
	rm -rf $(LIB) $(OBJ)
