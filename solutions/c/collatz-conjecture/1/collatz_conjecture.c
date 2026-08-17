#include "collatz_conjecture.h"

int steps(int start)
{
    if (start < 1)
    {
        return -1;
    }

    int result = 0;
    while (start != 1)
    {
        ++result;
        if (start%2)
        {
            start = 3*start + 1;
        }
        else
        {
            start = start/2;
        }
    }
    return result;
}
