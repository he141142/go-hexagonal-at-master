package config

import (
	"fmt"
	"github.com/pingcap/log"
	"github.com/spf13/viper"
	"sync"
)

type IViper interface {
	InitEnvFiles()
	Instance() *viper.Viper
	IsSet(key string) bool
	UseEnv() IViper
	Get(key string) string
}

type _viper struct {
	IViper
	instance  *viper.Viper

}

var _Viper = &_viper{
	instance: viper.New(),
}

var once sync.Once


func (vp *_viper) Get(key string) string {

	log.Debug(vp.instance.GetString(key))
	value := ""
	value = vp.instance.GetString(key)
	//log.Info(fmt.Sprintf("key: %v value: %v", key, value))
	fmt.Println("======================")
	fmt.Println(value)
	return value
}



func (vp *_viper) Instance() *viper.Viper {
	return vp.instance
}

func (vp *_viper) IsSet(key string) bool {
	return vp.instance.IsSet(key) && vp.instance.Get(key) != ""
}

// cf : config file directory
func (vp *_viper) UseEnv() IViper {
	vp.instance.SetConfigFile("dev.env")
	//vp.instance.SetEnvPrefix("dev")
	vp.instance.AutomaticEnv()
	if err := vp.instance.ReadInConfig(); err != nil {
		panic(err.(any))
	}
	fmt.Println(vp.instance.Get("DB_HOST"))
	return vp
}

func (vp *_viper) setConfigFile(file string) {
	vp.instance.SetConfigFile(file)
}

func Init() *_viper {
	// Set the default configuration file paths. for development env
	//you can config it later
	_Viper.UseEnv()

	// Load the configuration values.
	if err := _Viper.instance.ReadInConfig(); err != nil {
		panic(err.(any))
	}

	return _Viper
}


func Viper() IViper {
	once.Do(func() {
		//default env config
		_Viper = Init()
	})
	return _Viper
}
