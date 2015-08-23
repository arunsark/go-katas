package fizz
import "strconv"

func FizzBuzz(n int) string {
    acc := ""
    if ( n%3 == 0) {
        acc += "fizz"
    }
    if ( n%5 == 0 ) {
        acc += "buzz"
    }
    if ( len(acc) == 0 ) {
        acc = strconv.Itoa(n)
    }
    return acc
}
