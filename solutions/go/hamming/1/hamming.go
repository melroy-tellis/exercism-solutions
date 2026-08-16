package hamming

import "errors"

func Distance(a, b string) (int, error) {
    d := 0
    var err error
    ra, rb := []rune(a), []rune(b)
    if len(ra) == len(rb) {
		for i := 0; i < len(ra); i++ {
        	if ra[i] != rb[i] {
            	d++
        	}
    	}
    } else {
        err = errors.New("The strings must be of equal length.")
    }

    return d, err
}
