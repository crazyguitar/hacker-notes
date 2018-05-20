SUBDIR = aws go c

.PHONY: all clean $(SUBDIR)

all: $(SUBDIR)

$(SUBDIR):
	$(MAKE) -C $@ $(MAKECMDGOALS)

clean: $(SUBDIR)
