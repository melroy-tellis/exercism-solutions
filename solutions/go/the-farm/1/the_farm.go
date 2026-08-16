package thefarm

import (
    "errors"
    "fmt"
)

type InvalidCowsError struct {
    nCows int;
    message string;
}

func (err InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", err.nCows, err.message)
    
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, nCows int) (float64, error) {
    var fodderAmount, fatteningFactor, foodPerCow float64
    var err error
    if fodderAmount, err = fc.FodderAmount(nCows); err == nil {
        if fatteningFactor, err = fc.FatteningFactor(); err == nil {
            foodPerCow = (fodderAmount * fatteningFactor)/float64(nCows)
        }
    }
    return foodPerCow, err
    
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, nCows int) (float64, error) {
    if nCows <= 0 {
        return 0.0, errors.New("invalid number of cows")
    }
    return DivideFood(fc, nCows)
    
}


// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(nCows int) error {
    if (nCows < 0) {
        return &InvalidCowsError{nCows, "there are no negative cows"}
    } else if (nCows == 0) {
        return &InvalidCowsError{nCows, "no cows don't need food"}
    }
    return nil
    
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
