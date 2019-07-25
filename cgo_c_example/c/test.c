#include <stdio.h>
#include "../go/test.h"


int main() {
    printf("hello");

    GoString p0 = {"hello", 5};
    struct SayWhat_return result = SayWhat(p0);

    printf("\nresult %s, %d\n", result.r0, result.r1);
}