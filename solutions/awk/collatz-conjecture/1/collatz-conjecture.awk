$0 <= 0 { print "Error: Only positive numbers are allowed"; exit 1; }

{
    n = $0
    steps = 0
    while (n > 1) {
        ++steps
        if (n % 2 == 0) {
            n = n / 2
        } else {
            n = 3 * n + 1
        }
    }
    print steps
}