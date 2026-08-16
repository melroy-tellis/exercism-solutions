package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
    acc := initial
    for _, i := range(s) {
        acc = fn(acc, i)
    }
    return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
    acc := initial
    for _, i := range(s.Reverse()) {
        acc = fn(i, acc)
    }
    return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	filtered := IntList{}
    for _, i := range(s) {
        if fn(i) {
            filtered = append(filtered, i)
        }
    }
    return filtered
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	mapped := IntList{}
    for _, i := range(s) {
        mapped = append(mapped, fn(i))
    }
    return mapped
}

func (s IntList) Reverse() IntList {
	reversed := IntList{}
    for i := s.Length() - 1; i >= 0; i-- {
        reversed = append(reversed, s[i])
    }
    return reversed
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	concatenated := s
    for _, list := range(lists) {
        concatenated = concatenated.Append(list)
    }
    return concatenated
}
