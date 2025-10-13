package thefarm
import (
    "fmt"
    "errors"
)

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
    fooder, err := fc.FodderAmount(cows)
    if err != nil {
        return 0, err
    }
    fatten, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    amountPerCow := (fooder * fatten) / float64(cows)
    return amountPerCow, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
    if cows > 0 {
        return DivideFood(fc, cows)
    }
    return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
    cowsNumber	int
    message		string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.cowsNumber, e.message)
}

func ValidateNumberOfCows(cows int) error {
    switch {
        case cows < 0:
    		return &InvalidCowsError{ cowsNumber: cows, message: "there are no negative cows"}
        case cows == 0:
    		return &InvalidCowsError{ cowsNumber: cows, message: "no cows don't need food"}
        default:
        	return nil
    }
}
