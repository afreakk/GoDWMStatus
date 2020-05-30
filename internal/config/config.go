package config

import (
	"encoding/json"
	"io/ioutil"
)

type Module struct {
	Cron    string `json:"cron,omitempty"`
	Sprintf string `json:"sprintf,omitempty"`

	FilePath string `json:"file_path,omitempty"`

	DateTimeFormat string `json:"datetime_format,omitempty"`

	BarWidth  int64  `json:"bar_width"`
	BarFilled string `json:"bar_filled"`
	BarEmpty  string `json:"bar_empty"`

	CommandName string   `json:"command_name,omitempty"`
	CommandArgs []string `json:"command_args,omitempty"`

	Name string `json:"name,omitempty"`
}

type Config struct {
	Modules               []Module `json:"modules"`
	MinimumRenderInterval string   `json:"minimum_render_interval"`
	WMClient              string   `json:"wmclient,omitempty"`
}

func GetConfigFromPath(configFilePath string) (Config, error) {
	var cfg Config
	cfgTxt, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return cfg, err
	}

	if err = json.Unmarshal(cfgTxt, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
