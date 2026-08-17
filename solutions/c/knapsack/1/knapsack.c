#include "knapsack.h"

unsigned int maximum_value(unsigned int maximum_weight, item_t *items, size_t item_count)
{
    if (maximum_weight == 0 || item_count == 0)
    {
        return 0;
    }
    unsigned int not_taken = maximum_value(maximum_weight, &items[1], item_count-1);
    if (items[0].weight <= maximum_weight)
    {
        unsigned int taken = items[0].value + maximum_value(maximum_weight-items[0].weight, 
                                                    &items[1], 
                                                    item_count-1);
        if (taken > not_taken)
            return taken;
    }
    return not_taken;
}
