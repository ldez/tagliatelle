package a

type Foo struct {
	ID     string `json:"ID"`     // want `json\(camel\): got 'ID' want 'id'`
	UserID string `json:"UserID"` // want `json\(camel\): got 'UserID' want 'userId'`
	Name   string `json:"name"`
	Value  string `json:"value,omitempty"`
	Bar    Bar    `json:"bar"`
	Bur    `json:"bur"`

	Qiix Quux `json:",inline"`
	Quux `json:",inline"`
}

type Bar struct {
	Name                 string `json:"-"`
	Value                string `json:"value"`
	CommonServiceFooItem *Bir   `json:"CommonServiceItem,omitempty"` // want `json\(camel\): got 'CommonServiceItem' want 'commonServiceFooItem'`
}

type Bir struct {
	Name             string   `json:"-"`
	Value            string   `json:"value"`
	ReplaceAllowList []string `mapstructure:"replace-allow-list"`
}

type Bur struct {
	Name    string
	Value   string `yaml:"Value"` // want `yaml\(camel\): got 'Value' want 'value'`
	More    string `json:"-"`
	Also    string `json:",omitempty"` // want `json\(camel\): got 'Also' want 'also'`
	ReqPerS string `avro:"req_per_s"`
}

type Quux struct {
	Data []byte `json:"data"`
}

// MessedUpTags is to validate structtag validation is done
// without it, the tool could let think everything is ok, while it's not.
// We could validate more usecases, but we don't as everything is already validated
// by analysis/passes/structtag unit tests.
type MessedUpTags struct {
	// an invalid tag listed in the rule is supported.
	Bad string `json:"bad` // want "struct field tag `json:\"bad` not compatible with reflect.StructTag.Get: bad syntax for struct tag value"

	// an invalid tag not in the rule is supported.
	Whatever string `foo:whatever` // want "struct field tag `foo:whatever` not compatible with reflect.StructTag.Get: bad syntax for struct tag value"

	// a invalid tag supported by the rule, is not hidden by another broken tag
	Mixed string `json:"mixed" foo:mixed` // want "struct field tag `json:\"mixed\" foo:mixed` not compatible with reflect.StructTag.Get: bad syntax for struct tag value"
}
