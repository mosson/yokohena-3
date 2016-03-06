package route

import "testing"

var tests = map[string]string{
	"b":               "AB",
	"l":               "AD",
	"r":               "AC",
	"bbb":             "ABAB",
	"rrr":             "ACBA",
	"blll":            "ABCAB",
	"llll":            "ADEBA",
	"rbrl":            "ACADE",
	"brrrr":           "ABEDAB",
	"llrrr":           "ADEFDE",
	"lrlll":           "ADFEDF",
	"lrrrr":           "ADFCAD",
	"rllll":           "ACFDAC",
	"blrrrr":          "ABCFEBC",
	"brllll":          "ABEFCBE",
	"bbbrrlrl":        "ABABEDFCB",
	"rbllrrrr":        "ACABCFEBC",
	"lbrlrrblr":       "ADABCFEFCA",
	"rlbrrrrbl":       "ACFCADFCFD",
	"bllrlrbrrb":      "ABCADEFEBCB",
	"rllrllllbb":      "ACFDEBADEDE",
	"blblrlrrlbr":     "ABCBEDFCABAC",
	"lrlrrrrrbrb":     "ADFEBCFEBEDE",
	"rblllrlrrlrr":    "ACABCADEFDABE",
	"rbrrlrblrllb":    "ACADFEBEFDACA",
	"lrrrlrllrrllr":   "ADFCABEFCADEBC",
	"rrlblllrrlrrb":   "ACBEBADEFDABEB",
	"brbllrrbbrlrll":  "ABEBADFCFCABEFC",
	"rrrbbrlbrlblrb":  "ACBABACFCABADFD",
	"lllllllllllblrr": "ADEBADEBADEBEFDE",
	"llllllrllrlbrrr": "ADEBADEFCBADABED",
}

func TestWalk(t *testing.T) {
	s := NewStranger()

	r := s.Walk("b")

	if r != "B" {
		t.Errorf("expected A, actual %v", r)
	}

	s = NewStranger()

	r = s.Walk("l")
}

func TestSolve(t *testing.T) {
	for k, v := range tests {
		s := NewStranger()
		r := s.solve(k)
		if r != v {
			t.Errorf("%v: expected %v, actual %v", k, v, r)
		}
	}

}
