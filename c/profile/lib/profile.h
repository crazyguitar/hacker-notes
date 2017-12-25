#ifndef PROF_H
#define PROF_H

#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <time.h>

typedef struct _profile_info {
	uint32_t count;
	double cost_sec;
} profile_info;


typedef enum {
	FUNC_MIN = -1,
#define PROFILE_FUNC(func) FUNC_##func,
#include "profile_func.h"
#undef PROFILE_FUNC
	FUNC_MAX
} profile_idx;

extern void profile_update(profile_idx idx, double cost_sec);
extern void profile_dump();

#ifndef PROFILING
#define PROFILE_START(func)
#define PROFILE_END(func)
#else
static double ts2sec(struct timespec* ts) __attribute__((unused));

static double ts2sec(struct timespec* ts)
{
	return (double)ts->tv_sec + (double)ts->tv_nsec / 1000000000.0;
}

#define PROFILE_START(func)                                              \
	struct timespec p_##func;                                        \
	do {                                                             \
		if(clock_gettime(CLOCK_MONOTONIC, &p_##func) != 0) {     \
			abort();                                         \
		}                                                        \
	} while(0)

#define PROFILE_END(func)                                                \
	do {                                                             \
		struct timespec p_##func_end;                            \
		if(clock_gettime(CLOCK_MONOTONIC, &p_##func_end) != 0) { \
			abort();                                         \
		}                                                        \
		profile_update(FUNC_##func,                              \
			ts2sec(&p_##func_end) - ts2sec(&p_##func));      \
	} while(0)
#endif /* PROFILING */

#endif /* PROF_H */
