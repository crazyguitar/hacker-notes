#include <stdio.h>
#include <unistd.h>
#include "profile.h"

int fib(int n)
{
	int rc = 0;
	int a = 0, b = 1, tmp = 0;
#ifdef PROFILING
	PROFILE_START(fib);
#endif
	if (n < 1)
		goto end;

	while (n > 1) {
		tmp = b;
		b = a + b;
		a = tmp;
		n--;
	}
	rc = b;
end:
#ifdef PROFILING
	PROFILE_END(fib);
#endif
	return rc;
}

void sleep_func(int sec)
{
#ifdef PROFILING
	PROFILE_START(sleep_func);
#endif

	sleep(sec);

#ifdef PROFILING
	PROFILE_END(sleep_func);
#endif
}
