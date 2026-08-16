# These variables are initialized on the command line (using '-v'):
# - num

BEGIN {
    result = ""
    if (num % 3 == 0) { result = "Pling" }
    if (num % 5 == 0) { result = result "Plang" }
    if (num % 7 == 0) { result = result "Plong" }
    if (length(result) == 0) { result = num }
    print result > "/dev/stderr"
}
