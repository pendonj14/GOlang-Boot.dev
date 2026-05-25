Return values may be given names, and if they are, then they are treated the same as if they were new variables defined at the top of the function.

Named return values are best thought of as a way to document the purpose of the returned values.

According to the tour of go:

A return statement without arguments returns the named return values. This is known as a "naked" return. Naked return statements should be used only in short functions. They can harm readability in longer functions.

Named return values are what enable naked returns. Use naked returns only in short functions where the purpose of the returned values is obvious.

func getCoords() (x, y int) {
// x and y are initialized with zero values

    return // automatically returns x and y

}

Is the same as:

func getCoords() (int, int) {
var x int
var y int
return x, y
}

In the first example, x and y are the return values. At the end of the function, we could simply write return to return the values of those two variables, rather than writing return x,y.

Assignment
One of our clients likes us to send text messages reminding users of life events coming up.

Fix the bug by adding named return values to the function signature – the bare return at the end is already a naked return that will return them. The variables need to be automatically initialized. Order them as they appear in the code. Do not alter the body of the function.
