package util

import (
	"github.com/spf13/viper"
)

type PrinterConfig struct {
    Name string `unmarshal:"name"`
}

type FiletypesConfig struct {
    FileExt string `unmarshal:"FileExt"`
}

type Config struct {
    Apath string `unmarshal:"Apath"`
    Bpath string `unmarshal:"Bpath"`
    Paths struct {
        INpath string `unmarshal:"paths.INpath"`
        OUTpath string `unmarshal:"paths.OUTpath"`
    }
    Printer PrinterConfig `mapstructure:"printer"`
    Filetypes FiletypesConfig `mapstructure:"filetypes"`
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
