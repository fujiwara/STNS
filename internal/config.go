package stns

import "github.com/BurntSushi/toml"

type Config struct {
	Users  map[string]*Attr
	Groups map[string]*Attr
}

type UserAttr struct {
	Group_Id  int
	Directory string
	Shell     string
	Password  string
	Gecos     string
	Keys      []string
}
type GroupAttr struct {
	Users []string
}
type Attr struct {
	Id   int
	Name string
	*UserAttr
	*GroupAttr
}

var AllConfig *Config

func LoadConfig(configFile string) {
	_, err := toml.DecodeFile(configFile, &AllConfig)
	if err != nil {
		panic(err)
	}
}