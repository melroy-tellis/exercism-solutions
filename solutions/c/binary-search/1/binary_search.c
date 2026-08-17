#include "binary_search.h"

const int *binary_search(int value, const int *arr, size_t length)
{
    int low = 0, high = length-1;
    while (low <= high)
    {
        int mid = low + (high-low)/2;
        if (arr[mid] < value)
            low = mid+1;
        else if (arr[mid] > value)
            high = mid-1;
        else
            return &arr[mid];
    }
    return NULL;
}


