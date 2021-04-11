package a

type Foo struct {
	ID     string `json:"ID"`     // want `json\(camel\): got 'ID' want 'id'`
	UserID string `json:"UserID"` // want `json\(camel\): got 'UserID' want 'userId'`
	Name   string `json:"name"`
	Value  string `json:"value,omitempty"`
	Bar    Bar    `json:"bar"`
	Bur    `json:"bur"`
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
	Also    string `json:",omitempty"`
	ReqPerS string `avro:"req_per_s"`
}
