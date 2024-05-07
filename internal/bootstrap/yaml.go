package bootstrap

type yamlData struct {
	Access  [][]string
	Develop struct {
		UserId int64 `yaml:"user_id,omitempty"`
	}
	Snowflake struct {
		Node int64 `json:"node,omitempty"`
	}
}
