package configmanager

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// ApplicatonConfig is the configuration for the ApplicatonConfig
// it does not have interface specific configuration
type ApplicatonConfig struct {
	ProcessName string `json:"process_name,omitempty"`
	LogLevel    string `json:"log_severity,omitempty"`
	SqliteFile  string `json:"sqlite_file"`
}

// Validate validates the application config and fails if mandatory parameters are missing
func (cnf ApplicatonConfig) Validate() error {
	if cnf.LogLevel == "" ||
		cnf.ProcessName == "" {
		return errors.New("invalid Config: LogDirectory, LogLevel and ProcessName are mandatory")
	}
	return nil
}

// Config stores the configuration
var Config *ApplicatonConfig
var configFile *string

// LoadConfiguration loads configuration from file
// decodes the json config file into an instance of application config
// if the decoded config is valid it is set as config
func LoadConfiguration() error {
	if configFile == nil {
		return errors.New("config not initialized")
	}
	config := new(ApplicatonConfig)
	file, err := os.Open(*configFile)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return err
	}
	if err := config.Validate(); err != nil {
		return err
	}
	Config = config
	return nil
}

// GetConfig returns a copy of init-ed application config instance
// if not already initialized it is initialized with "." as config path
func GetConfig() ApplicatonConfig {
	if Config != nil {
		return *Config
	}
	err := InitConfig(".")
	if err != nil {
		panic(err)
	}
	return *Config
}
func setupReload() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for range c {
			log.Printf("Reloading configurations....\n")
			if configFile == nil {
				panic("Config file path not set!")
			}
			if err := LoadConfiguration(); err != nil {
				log.Printf("Error on reloading configurations, using old configuration : Error : %s\n", err.Error())
			}
		}
	}()
}

// InitConfig will initialize app config with config file name
func InitConfig(config string) error {
	if config == "" {
		config = "."
	}
	configFilepath := strings.TrimRight(config, "/") + "/config.json"
	configFile = &configFilepath
	setupReload()
	return LoadConfiguration()
}
