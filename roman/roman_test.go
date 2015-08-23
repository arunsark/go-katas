package roman

import "testing"

func TestRomanRepresentations(t *testing.T) {
    var num uint16 = 1
    testRoman(t, num, "I")
    num = 2
    testRoman(t, num, "II")
    num = 3
    testRoman(t, num, "III")
    num = 4
    testRoman(t, num, "IV")

}

func testRoman(t *testing.T, input uint16, want string) {
    got := Roman(input)

    if got != want {
        t.Errorf("Got Roman(%q) == %q, want %q", input, got, want)
    }
}
