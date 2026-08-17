#include "reverse_string.h"
#include <string.h>
#include <stdlib.h>

char *reverse(const char *value)
{
    size_t n = strlen(value);
    char *reversed = calloc(n+1, sizeof(char));
    for (size_t i=0; i<n; ++i)
    {
        reversed[i] = value[n-1-i];        
    }
    return reversed;
}
