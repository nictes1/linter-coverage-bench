package typeNumber

type Test struct {
	in  int
	out string
}

var tests = []Test{{-1, "negative"}, {5, "positive"}}

func TestTypeNumber(t *testing.T) {
	for i, test := range tests {
		size := TypeNumber(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}
