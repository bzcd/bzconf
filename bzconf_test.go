package bzconf

import "testing"

func TestBZConf(t *testing.T) {
	s, err := ListAll("..",
		DefaultOptionsList().WithRecursion(true).
			WithSuffix(".go"),
	)

	t.Log(err, s)

	s, err = ListAll("..",
		DefaultOptionsList().WithRecursion(true).
			WithSuffix(".go").WithDepth(2),
	)

	t.Log(err, s)
}
