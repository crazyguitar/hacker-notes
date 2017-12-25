#include "profile.h"

#ifndef ARRAY_SIZE
#define ARRAY_SIZE(arr) sizeof(arr) / sizeof(arr[0])
#endif

static char *ftable[] = {
#define PROFILE_FUNC(func) #func,
#include "profile_func.h"
#undef PROFILE_FUNC
	NULL
};

static profile_info ptable[] = {
#define PROFILE_FUNC(func) { .count = 0, .cost_sec = 0.0 },
#include "profile_func.h"
#undef PROFILE_FUNC
	{ .count = 0, .cost_sec = 0.0 }
};

void profile_update(profile_idx idx, double cost_sec)
{
	if (idx <= FUNC_MIN || idx >= FUNC_MAX)
		abort();

	ptable[idx].count++;
	ptable[idx].cost_sec += cost_sec;
}

void profile_dump()
{
	uint32_t i = 0;

	printf("%s %-15s %-15s %-15s\n", "[idx]", "name", "count", "cost");
	for (i = 0; i < ARRAY_SIZE(ptable) - 1; i++) {
		printf("[%03d] %-15s %-15u %10.5f\n",
			i, ftable[i], ptable[i].count, ptable[i].cost_sec);
	}
}
