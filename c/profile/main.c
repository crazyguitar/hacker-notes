#include <stdio.h>
#include <misc.h>
#include <profile.h>

int main(int argc, char *argv[])
{
	int i = 0;
	int ret = 0;

	for (i = 0; i < 100000; i++) {
		ret = fib(i);
	}

	for (i = 1; i < 3; i++) {
		sleep_func(i);
	}

	profile_dump();

	return 0;
}
