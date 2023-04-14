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
	Name                string
	Value               string `yaml:"Value"` // want `yaml\(camel\): got 'Value' want 'value'`
	More                string `json:"-"`
	Also                string `json:",omitempty"` // want `json\(camel\): got 'Also' want 'also'`
	ReqPerS             string `avro:"req_per_s"`
	HeaderValue         string `header:"Header-Value"`
	WrongHeaderValue    string `header:"Header_Value"` // want `header\(header\): got 'Header_Value' want 'Wrong-Header-Value'`
	EnvConfigValue      string `envconfig:"ENV_CONFIG_VALUE"`
	WrongEnvConfigValue string `envconfig:"env_config_value"` // want `envconfig\(upperSnake\): got 'env_config_value' want 'WRONG_ENV_CONFIG_VALUE'`
}

type Quux struct {
	Data []byte `json:"data"`
}

// MessedUpTags struct is to validate the tool is not doing any validation about invalid tags.
// Please read the readme about this choice.
type MessedUpTags struct {
	// an invalid tag cannot be validated.
	Bad string `json:"bad`

	// a tag not supported by the rules is not validated.
	Whatever string `foo:whatever`

	// a tag supported by the rule cannot be validated because foo tag breaks the whole tags block
	Mixed string `json:"mixed" foo:mixed`
}
