package config

type CommonKey string

const (
	CommonKeyAbstractLen   CommonKey = "abstract_len"
	CommonKeyAdminPassword CommonKey = "admin_password"
	CommonKeyAdminUsername CommonKey = "admin_username"
)

type Common struct {
	IntItems map[CommonKey]int    `json:"intItems,omitempty"`
	StrItems map[CommonKey]string `json:"strItems,omitempty"`
}

func NewCommon() *Common {
	return &Common{
		IntItems: make(map[CommonKey]int),
		StrItems: make(map[CommonKey]string),
	}
}

func (c *Common) GetIntWithDefault(key CommonKey, _default int) int {
	v, ok := c.IntItems[key]
	if ok {
		return v
	}
	return _default
}

func (c *Common) GetStrWithDefault(key CommonKey, _default string) string {
	v, ok := c.StrItems[key]
	if ok {
		return v
	}
	return _default
}
