package setting

import "github.com/spf13/viper"

type Config struct {
	HomebrewList        InstallOption
	VscodeExtensionList []string
}

func (c *Config) Load(fileName string) (err error) {

	// set the path of the configuration file
	viper.AddConfigPath("../config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")

	// set the name of the configuration file
	viper.SetConfigName(fileName)

	// read the configuration file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// unmarshal the configuration file
	err = viper.Unmarshal(&c)
	if err != nil {
		return
	}

	return
}

type InstallOption struct {
	Kit  []string
	Cask []string
}
