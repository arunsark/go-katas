package string

import "testing"

func Test(t *testing.T) {

    var tests = []struct{
        s, want string
    }{
        {"Backward", "drawkcaB"},
        {"Ho ñ荷", "荷ñ oH"},
    }

    for _, c := range tests {
        got := Reverse(c.s)
        if got != c.want {
            t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
        }
    }
}
