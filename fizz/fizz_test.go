package fizz

import "testing"
import "strconv"

func TestNoFizzBuzz(t *testing.T) {
    testFizzBuzz(t, 1, "1")
    testFizzBuzz(t, 2, "2")
    testFizzBuzz(t, 3, "fizz")
    testFizzBuzz(t, 5, "buzz")
    testFizzBuzz(t, 6, "fizz")
    testFizzBuzz(t, 7, "7")
    testFizzBuzz(t, 10, "buzz")
    testFizzBuzz(t, 15, "fizzbuzz")
}

func testFizzBuzz(t *testing.T, input int, want string) {
    got := FizzBuzz(input)

    if got != want {
        t.Errorf("Got FizzBuzz(%q) == %q, want %q", strconv.Itoa(input), got, want)
    }
}
