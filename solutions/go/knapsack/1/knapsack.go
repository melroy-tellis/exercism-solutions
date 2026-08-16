package knapsack

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {
	if len(items) == 0 || maximumWeight == 0 {
        return 0
    }
    item, otherItems := items[0], items[1:] 

    valueWithoutItem := Knapsack(maximumWeight, otherItems)
    if maximumWeight >= item.Weight {
        valueWithItem := item.Value + Knapsack(maximumWeight-item.Weight, otherItems)
        if valueWithItem > valueWithoutItem {
            return valueWithItem
        }
    }
    return valueWithoutItem
}
