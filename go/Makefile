SUBDIR = list tree tls socket

.PHONY: all clean $(SUBDIR)

all: $(SUBDIR)

$(SUBDIR):
	$(MAKE) -C $@ $(MAKECMDGOALS)

clean: $(SUBDIR)
