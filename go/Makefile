SUBDIR = list tree socket

.PHONY: all clean $(SUBDIR)

all: $(SUBDIR)

$(SUBDIR):
	$(MAKE) -C $@ $(MAKECMDGOALS)

clean: $(SUBDIR)
