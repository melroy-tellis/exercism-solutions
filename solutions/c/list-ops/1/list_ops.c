#include "list_ops.h"
#include <string.h>

list_t *new_list(size_t length, list_element_t elements[])
{
    list_t *list = malloc(sizeof(list_t) + length*sizeof(list_element_t));
    list->length = length;
    memcpy(list->elements, elements, length*sizeof(list_element_t));
    return list;
    
}

list_t *append_list(list_t *list1, list_t *list2)
{   size_t appended_length = list1->length + list2->length;
    list_t *appended = new_list(appended_length, list1->elements);
    appended->length = appended_length;
    memcpy(appended->elements+list1->length, list2->elements, list2->length*sizeof(list_element_t));
    return appended;
}

list_t *filter_list(list_t *list, bool (*filter)(list_element_t))
{
    list_t *filtered = new_list(list->length, list->elements);
    filtered->length = 0;
    for (size_t i = 0; i < list->length; ++i)
    {
        if (filter(list->elements[i]))
        {
            filtered->elements[filtered->length++] = list->elements[i];
        }
    }
    return filtered;
}

size_t length_list(list_t *list)
{
    return list->length;
}

list_t *map_list(list_t *list, list_element_t (*map)(list_element_t))
{
    list_t *mapped = new_list(list->length, list->elements);
    for (size_t i = 0; i < list->length; ++i)
    {
        mapped->elements[i] = map(list->elements[i]);
    }
    return mapped;
}

list_element_t foldl_list(list_t *list, list_element_t initial,
                          list_element_t (*foldl)(list_element_t,
                                                  list_element_t))
{
    list_element_t acc = initial;
    for (size_t i = 0; i < list->length; ++i)
    {
        acc = foldl(acc, list->elements[i]);
    }
    return acc;
}

list_element_t foldr_list(list_t *list, list_element_t initial,
                          list_element_t (*foldr)(list_element_t,
                                                  list_element_t))
{
    list_element_t acc = initial;
    for (size_t i = 0; i < list->length; ++i)
    {
        acc = foldr(list->elements[list->length-1-i], acc);
    }
    return acc;
}

list_t *reverse_list(list_t *list)
{
    list_t *reversed = new_list(list->length, list->elements);
    for (size_t i = 0; i < list->length; ++i)
    {
        reversed->elements[i] = list->elements[list->length-1-i];
    }
    return reversed;
}

void delete_list(list_t *list)
{
    free(list);
}