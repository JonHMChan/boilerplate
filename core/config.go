package core

import (
	"log"
	"os"
	"reflect"
)

const (
	LOCAL    = "LOCAL"
	DEV      = "DEVELOPMENT"
	PROD     = "PRODUCTION"
	TIER_KEY = "TIER_KEY"
)

var Config ConfigStruct

func (c *ConfigStruct) Start() {
	Config.Env = EnvStruct{}
	Config.Env.Check()

	Config.Location = LocationStruct{
		Host: Config.Location.getHost(),
	}

	Config.Settings = SettingsStruct{}
	Config.Settings.Init()
}

func (c *ConfigStruct) Local() bool {
	return c.Env.TIER == LOCAL
}

func (c *ConfigStruct) Dev() bool {
	return c.Env.TIER == DEV
}

func (c *ConfigStruct) Prod() bool {
	return c.Env.TIER == PROD
}

type ConfigStruct struct {
	Location LocationStruct
	Env      EnvStruct
	Settings SettingsStruct
}

type LocationStruct struct {
	Host string
}

func (location *LocationStruct) getHost() string {
	host := os.Getenv("HOST")
	if host == "localhost" {
		host = host + ":" + os.Getenv("PORT")
	}
	return host
}

type EnvStruct struct {
	HOST     string
	TIER     string
	TIER_KEY string
	GIN_MODE string
	PORT     string

	CSRF_SECRET string

	DATABASE_URL string
}

func (e *EnvStruct) Check() {
	ok := true
	s := reflect.ValueOf(e).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		f.SetString(os.Getenv(typeOfT.Field(i).Name))
		value := f.Interface().(string)
		if len(value) == 0 {
			log.Printf("%s %s\n",
				typeOfT.Field(i).Name, f.Type())
			ok = false
		}
	}
	if !ok {
		log.Fatal("\033[0;33mThe application is missing the above environment variables. Check your .env or deployment environment variables.\033[0m")
	}
}
