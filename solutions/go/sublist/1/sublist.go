package sublist

import (
    "slices"
)

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	
    for i := 0; i <= len(l1)-len(l2); i++ {
        if slices.Equal(l1[i:i+len(l2)], l2) {
            if len(l1) == len(l2) {
                return RelationEqual
            }
            return RelationSuperlist
        }
    }
    for j := 0; j <= len(l2)-len(l1); j++ {
        if slices.Equal(l2[j:j+len(l1)], l1) {
            return RelationSublist
        }
    }
    return RelationUnequal
    
}
