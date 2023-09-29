package config

import (
	"fmt"
	"os"
	"time"

	"jtso/logger"

	"github.com/spf13/viper"
)

type PortalConfig struct {
	Port int
}

type NetconfConfig struct {
	Port       int
	RpcTimeout int
}

type GnmiConfig struct {
	Port       int
	SkipVerify bool
}

type EnricherConfig struct {
	Folder   string
	Interval time.Duration
	Workers  int
	Port     int
}
type ConfigContainer struct {
	//Instances []*InstanceConfig
	Enricher *EnricherConfig
	Portal   *PortalConfig
	Netconf  *NetconfConfig
	Gnmi     *GnmiConfig
}

func NewConfigContainer(f string) *ConfigContainer {
	viper.SetConfigFile(f)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Log.Errorf("Fatal error config file %v", err)
		fmt.Println("Fatal error config file: default \n", err)
		os.Exit(1)
	}

	logger.Log.Info("Read configuration file")

	// Ser default value for portal
	viper.SetDefault("modules.portal.port", 8080)

	// Ser default value for enricher
	viper.SetDefault("modules.enricher.folder", "./")
	viper.SetDefault("modules.enricher.interval", 720*time.Minute)
	viper.SetDefault("modules.enricher.workers", 4)
	viper.SetDefault("modules.enricher.port", 10000)

	// Set default value for Netconf
	viper.SetDefault("protocols.netconf.port", 830)
	viper.SetDefault("protocols.netconf.rpc_timeout", 60)

	// Set default value for gnmi
	viper.SetDefault("protocols.gnmi.port", 9339)
	viper.SetDefault("protocols.gnmi.skip_verify", true)

	return &ConfigContainer{
		//Instances: inst,
		Portal: &PortalConfig{
			Port: viper.GetInt("modules.portal.port"),
		},
		Enricher: &EnricherConfig{
			Folder:   viper.GetString("modules.enricher.folder"),
			Interval: viper.GetDuration("modules.enricher.interval") * time.Minute,
			Workers:  viper.GetInt("modules.enricher.workers"),
			Port:     viper.GetInt("modules.enricher.port"),
		},
		Netconf: &NetconfConfig{
			Port:       viper.GetInt("protocols.netconf.port"),
			RpcTimeout: viper.GetInt("protocols.netconf.rpc_timeout"),
		},
		Gnmi: &GnmiConfig{
			Port:       viper.GetInt("protocols.gnmi.port"),
			SkipVerify: viper.GetBool("protocols.gnmi.skip_verify"),
		},
	}
}
