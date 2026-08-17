#include "etl.h"
#include <string.h>
#include <stdlib.h>
#include <ctype.h>

int convert(const legacy_map *input, const size_t input_len, new_map **output)
{
    size_t output_len = 0;
    int letter_values[26] = {0};
    for (size_t i=0; i<input_len; ++i)
    {
        size_t n = strlen(input[i].keys);
        for (size_t j=0; j<n; ++j)
        {
            letter_values[input[i].keys[j]-'A'] = input[i].value;
            ++output_len;
        }
    }
    
    *output = calloc(output_len, sizeof(new_map));
    output_len = 0;
    for (size_t i=0; i<26; ++i)
    {
        if (letter_values[i] > 0)
        {
            (*output)[output_len].key = 'a' + i;
            (*output)[output_len].value = letter_values[i];
            ++output_len;
        }
    }
    return output_len;
}
