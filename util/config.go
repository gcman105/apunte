package util

import (
	"github.com/spf13/viper"
)

type Config struct {
    INpath string `unmarshal:"INpath"`
    OUTpath string `unmarshal:"OUTpath"`
    Paths struct {
        Apath string `unmarshal:"paths.Apath"`
        Bpath string `unmarshal:"paths.Bpath"`
    }
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigName("apunte")
    viper.SetConfigType("toml")
    //viper.SetConfigType("yaml")

    viper.AutomaticEnv()
    
    err = viper.ReadInConfig()
    if err != nil { // Handle errors reading the config file
        return
    }

    err = viper.Unmarshal(&config)
    return
}
