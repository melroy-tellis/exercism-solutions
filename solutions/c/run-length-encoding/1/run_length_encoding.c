#include "run_length_encoding.h"
#include <stddef.h>
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>

char *encode(const char *text)
{
    size_t n = strlen(text);
    if (n == 0)
    {
        return calloc(1, sizeof(char));
    }
    
    char *encoding = calloc(n+1, sizeof(char));
    char previous = text[0];
    size_t run_length = 1, encoding_length = 0;
    
    for (size_t i = 1; i <= n; ++i)
    {
        if (text[i] == previous)
        {
            ++run_length;
        }
        else
        {
            if (run_length > 1)
            {
                encoding_length += sprintf(&encoding[encoding_length], "%zu", run_length);
                run_length = 1;
            }
            encoding[encoding_length++] = previous;
            previous = text[i];
        }
    }
    return encoding;
}

char *decode(const char *data)
{
    size_t n = strlen(data);
    
    size_t run_length = 0;
    size_t decoding_length = 0;
    // One pass to calculate the length of the string after decoding
    for (size_t i = 0; i < n; ++i)
    {
        if (isdigit(data[i]))
        {
            run_length = run_length*10 + (data[i]-'0');
        }
        else
        {
            if (run_length == 0)
            {
                ++run_length;
            }
            decoding_length += run_length;
            run_length = 0;
        }
    }

    char *decoding = calloc(decoding_length+1, sizeof(char));

    run_length = 0, decoding_length = 0;
    for (size_t i = 0; i < n; ++i)
    {
        if (isdigit(data[i]))
        {

            run_length = run_length*10 + (data[i]-'0');
        }
        else
        {
            if (!run_length)
            {
                ++run_length;
            }
            while (run_length)
            {
                decoding[decoding_length++] = data[i];
                run_length--;
            }
        }
    }
    return decoding;
}
