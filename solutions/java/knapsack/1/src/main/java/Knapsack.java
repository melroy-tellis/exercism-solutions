import java.util.List;

class Knapsack {

    int maximumValue(int maximumWeight, List<Item> items) {
        if (items.size() == 0 || maximumWeight == 0) {
            return 0;
        }
        var first = items.get(0);
        var remaining = items.subList(1, items.size()); 
        
        final int notTaken = maximumValue(maximumWeight, remaining);
        if (first.weight > maximumWeight) {
            return notTaken;
        }
        int taken = first.value + maximumValue(maximumWeight-first.weight, remaining);
        return Math.max(notTaken, taken);
        
    }

}