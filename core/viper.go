package core

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/kinggq/core/internal"
	"github.com/kinggq/global"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" {
		config = internal.ConfigDefaultFile
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
