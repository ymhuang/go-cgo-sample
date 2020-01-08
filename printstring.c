#include "printstring.h"
#include <stdio.h>

int printstring(const char * s, char * out) 
{
    int n;

    n = sprintf(out, "%s\n", s);
    printf("in c: %s\n", s);

    return n;
}
