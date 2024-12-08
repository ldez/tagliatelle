package extended

type Foo struct {
	MyDB  string `json:"myDb"` // want `json\(goCamel\): got 'myDb' want 'myDB'`
	Dbase string `json:"dbase"`

	VirtualIPv4 string `yaml:"virtual_i_pv4"` // just to illustrate the current limitations: I think that 'virtual_ip_v4' is expected.
	VirtualIPv6 string `json:"virtualIPv6"`

	DocURLs string `json:"docUrLs"` // just to illustrate the current limitations: I think that 'docURLs' is expected.
}

type Bar struct {
	MyDB  string `json:"myDB"`
	Dbase string `json:"dbase"`
}

type FooBar struct {
	MyDB  string `sample:"myDb"` // base rule.
	Dbase string `sample:"dbase"`
}
