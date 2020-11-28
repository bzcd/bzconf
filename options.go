package bzconf

// OptionsList option of ListAll
type OptionsList struct {
	Prefix string // has Prefix
	Suffix string // has suffix

	IncludeDir bool // include dir, default false
	Depth      int  // depth

	Recursion   bool // recursion, default false,
	IgnoreError bool // return if false, default true
}

// DefaultOptionsList default OptionsList
func DefaultOptionsList() *OptionsList { return &OptionsList{IgnoreError: true} }

func (opt *OptionsList) WithPrefix(prefix string) *OptionsList { opt.Prefix = prefix; return opt }

func (opt *OptionsList) WithSuffix(suffix string) *OptionsList { opt.Suffix = suffix; return opt }

func (opt *OptionsList) WithRecursion(recursion bool) *OptionsList {
	opt.Recursion = recursion
	return opt
}

func (opt *OptionsList) WithIgnoreError(ignoreerror bool) *OptionsList {
	opt.IgnoreError = ignoreerror
	return opt
}

func (opt *OptionsList) WithIncludeDir(includeDir bool) *OptionsList {
	opt.IncludeDir = includeDir
	return opt
}

func (opt *OptionsList) WithDepth(depth int) *OptionsList {
	opt.Depth = depth
	return opt
}
