BEGIN {
    sp = 0
    opening["]"] = "["
    opening[")"] = "("
    opening["}"] = "{"
}
function push(elem) {
    stack[sp++] = elem
}
function pop() {
    top = stack[--sp]
    del stack[sp]
    return top
}

function size() {
    return sp
}

{
    for (i = 1; i <= length; ++i) {
        char = substr($0, i, 1)
        if (char ~ /[\[({]/) {
            
            push(char)
        } else if (char in opening) {
            if (size() == 0 || pop() != opening[char]) {
                print "false"
                exit
            }
            
        }
    }
    print size() == 0 ? "true" : "false"
}
