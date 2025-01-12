package lsproduct

import (
    "errors"
    "strconv"
    "unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
    if len(digits) < span || span <= 0 {
        return 0, errors.New("Invalid input.")
    }

    for _, r := range digits {
        if unicode.IsDigit(r) != true {
            return 0, errors.New("Invalid non-digit characters in input.")
        }
    }

    if len(digits) == span {
        return Multiply(digits), nil
    }

    var result int64
	var max int64
    
    for i := 0; i < (len(digits) - span + 1); i++ {
        sub := digits[i : i+span]
        res := Multiply(sub)

        if res >= max {
            max = res
            result = max
        }
    }

    return result, nil
}

func Multiply(digits string) int64 {
    var product int64 = 1

    for _, r := range digits {
        n, err := strconv.Atoi(string(r))
        if err != nil {
            continue
        }
        product *= int64(n)
    }

    return product
}
