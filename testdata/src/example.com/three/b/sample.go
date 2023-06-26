package b

type Foo struct {
	ID     string `json:"ID"`
	UserID string `json:"UserID"`
	Name   string `json:"NAME"`
	Value  string `json:"VALUE,omitempty"`
	Bar    Bar    `json:"BAR"`
	Bur    `json:"BUR"`

	Qiix Quux `json:",inline"`
	Quux `json:",inline"`
}

type Bar struct {
	Name                 string `json:"-"`
	Value                string `json:"VALUE"`
	CommonServiceFooItem *Bir   `json:"COMMON_SERVICE_ITEM,omitempty"`
}

type Bir struct {
	Name             string   `json:"-"`
	Value            string   `json:"VALUE"`
	ReplaceAllowList []string `mapstructure:"replace-allow-list"`
}

type Bur struct {
	Name                string
	Value               string `yaml:"Value"`
	More                string `json:"-"`
	Also                string `json:",omitempty"`
	ReqPerS             string `avro:"req_per_s"`
	HeaderValue         string `header:"Header-Value"`
	WrongHeaderValue    string `header:"Header_Value"`
	EnvConfigValue      string `envconfig:"ENV_CONFIG_VALUE"`
	WrongEnvConfigValue string `envconfig:"env_config_value"`
	EnvValue            string `env:"ENV_VALUE"`
	WrongEnvValue       string `env:"env_value"`
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
