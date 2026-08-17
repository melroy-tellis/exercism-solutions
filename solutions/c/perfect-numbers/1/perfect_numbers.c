#include "perfect_numbers.h"
#include <stdio.h>

kind classify_number(int n)
{
    if (n <= 0)
    {
        return ERROR;
    }

    if (n == 1)
    {
        return DEFICIENT_NUMBER;
    }

    int aliquot_sum = 1;
    int factor = 2;
    while (factor*factor <= n)
    {
        if (!(n%factor))
        {
            aliquot_sum += factor;
            int other_factor = n/factor;
            if (other_factor > factor)
            {
                aliquot_sum += other_factor;    
            }
        }
        ++factor;
    }
    printf("%d\n", aliquot_sum);
    if (aliquot_sum == n)
    {
        return PERFECT_NUMBER;
    }
    if (aliquot_sum > n)
    {
        return ABUNDANT_NUMBER;
    }
    return DEFICIENT_NUMBER;
    
}