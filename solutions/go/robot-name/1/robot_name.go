package robotname

import (
    "errors"
    "math/rand/v2"
)

// Define the Robot type here.
type Robot struct {
    name string
}

var usedNames map[string]bool = make(map[string]bool)


func generateRandomName() string {
    const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits = "0123456789"
    nameBytes := make([]byte, 5)
    for i := 0; i < 2; i++ {
        nameBytes[i] = uppercase[rand.IntN(len(uppercase))]
    }
    for i := 2; i < 5; i++ {
    	nameBytes[i] = digits[rand.IntN(len(digits))]
    }
    return string(nameBytes)
}

func (r *Robot) Name() (string, error) {
    if r.name == "" {
    	if len(usedNames) == maxNames {
        	return "", errors.New("all possible names have been exhausted")
    	}
        for r.name == "" {
            name := generateRandomName()
            if _, ok := usedNames[name]; !ok {
                usedNames[name] = true
                r.name = name
            }
         }
    }
    return r.name, nil

}

func (r *Robot) Reset() {
	r.name = ""
}
