package bzconf

import "testing"

func TestBZConf(t *testing.T) {
	s, err := ListAll("../..",
		DefaultOptionsList().WithRecursion(true).
			WithSuffix(".go").WithPrefix("b"),
	)

	t.Log(err, s)

	s, err = ListAll("../..",
		DefaultOptionsList().WithRecursion(true).
			WithSuffix(".go").WithDepth(3),
	)

	t.Log(err, s)
}
