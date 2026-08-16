package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if (n < 1) {
        return -1, fmt.Errorf("%d is not a positive integer", n)
    }
    steps := 0
    for n > 1 {
        steps++
        if n & 1 == 0 {
            n >>= 1
        } else {
            n = (n<<1) + n + 1
        }
    }
    return steps, nil
}
