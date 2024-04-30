package setting

const (
	InstallFile = "Install.yml"
)

var (
	config *Config
)

// Init initializes the configuration
func Init() (err error) {

	// new configuration
	cfg := &Config{}

	// load Luckybo configuration
	err = cfg.Load(InstallFile)
	if err != nil {
		return
	}

	// set the configuration
	Set(cfg)
	return
}

// Set sets the configuration
func Set(cfg *Config) {
	config = cfg
}

// Get gets the configuration
func Get() *Config {
	return config
}

// GetBrewKit gets the Homebrew kit
func GetBrewKit() []string {
	return config.HomebrewList.Kit
}

// GetBrewCask gets the Homebrew cask
func GetBrewCask() []string {
	return config.HomebrewList.Cask
}

func GetVSCodeExtensions() []string {
	return config.VscodeExtensionList
}
