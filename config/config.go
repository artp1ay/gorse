// Copyright 2021 gorse Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/BurntSushi/toml"
)

// Config is the configuration for the engine.
type Config struct {
	Database  DatabaseConfig  `toml:"database"`
	Master    MasterConfig    `toml:"master"`
	Server    ServerConfig    `toml:"server"`
	Recommend RecommendConfig `toml:"recommend"`
}

func (config *Config) LoadDefaultIfNil() *Config {
	if config == nil {
		return &Config{
			Database:  *(*DatabaseConfig)(nil).LoadDefaultIfNil(),
			Master:    *(*MasterConfig)(nil).LoadDefaultIfNil(),
			Server:    *(*ServerConfig)(nil).LoadDefaultIfNil(),
			Recommend: *(*RecommendConfig)(nil).LoadDefaultIfNil(),
		}
	}
	return config
}

// DatabaseConfig is the configuration for the database.
type DatabaseConfig struct {
	DataStore            string   `toml:"data_store"`              // database for data store
	CacheStore           string   `toml:"cache_store"`             // database for cache store
	AutoInsertUser       bool     `toml:"auto_insert_user"`        // insert new users while inserting feedback
	AutoInsertItem       bool     `toml:"auto_insert_item"`        // insert new items while inserting feedback
	CacheSize            int      `toml:"cache_size"`              // cache size for intermediate recommendation
	PositiveFeedbackType []string `toml:"positive_feedback_types"` // positive feedback type
	PositiveFeedbackTTL  uint     `toml:"positive_feedback_ttl"`
	ItemTTL              uint     `toml:"item_ttl"`
}

// LoadDefaultIfNil loads default settings if config is nil.
func (config *DatabaseConfig) LoadDefaultIfNil() *DatabaseConfig {
	if config == nil {
		return &DatabaseConfig{
			AutoInsertUser: true,
			AutoInsertItem: true,
			CacheSize:      100,
		}
	}
	return config
}

// MasterConfig is the configuration for the master.
type MasterConfig struct {
	Port        int    `toml:"port"`
	Host        string `toml:"host"`
	HttpPort    int    `toml:"http_port"`
	HttpHost    string `toml:"http_host"`
	SearchJobs  int    `toml:"search_jobs"`
	FitJobs     int    `toml:"fit_jobs"`
	MetaTimeout int    `toml:"meta_timeout"`
}

// LoadDefaultIfNil loads default settings if config is nil.
func (config *MasterConfig) LoadDefaultIfNil() *MasterConfig {
	if config == nil {
		return &MasterConfig{
			Port:        8086,
			Host:        "127.0.0.1",
			HttpPort:    8088,
			HttpHost:    "127.0.0.1",
			SearchJobs:  1,
			FitJobs:     1,
			MetaTimeout: 60,
		}
	}
	return config
}

type RecommendConfig struct {
	PopularWindow      int `toml:"popular_window"`
	FitPeriod          int `toml:"fit_period"`
	MaxRecommendPeriod int `toml:"max_recommend_period"`
	SearchPeriod       int `toml:"search_period"`
	SearchEpoch        int `toml:"search_epoch"`
	SearchTrials       int `toml:"search_trials"`
}

// LoadDefaultIfNil loads default settings if config is nil.
func (config *RecommendConfig) LoadDefaultIfNil() *RecommendConfig {
	if config == nil {
		return &RecommendConfig{
			PopularWindow:      1,
			FitPeriod:          60,
			MaxRecommendPeriod: 1,
			SearchPeriod:       60,
			SearchEpoch:        100,
			SearchTrials:       10,
		}
	}
	return config
}

// ServerConfig is the configuration for the server.
type ServerConfig struct {
	APIKey   string `toml:"api_key"`
	DefaultN int    `toml:"default_n"`
}

// LoadDefaultIfNil loads default settings if config is nil.
func (config *ServerConfig) LoadDefaultIfNil() *ServerConfig {
	if config == nil {
		return &ServerConfig{
			APIKey:   "",
			DefaultN: 10,
		}
	}
	return config
}

// FillDefault fill default values for missing values.
func (config *Config) FillDefault(meta toml.MetaData) {
	// Default database config
	defaultDBConfig := *(*DatabaseConfig)(nil).LoadDefaultIfNil()
	if !meta.IsDefined("database", "auto_insert_user") {
		config.Database.AutoInsertUser = defaultDBConfig.AutoInsertUser
	}
	if !meta.IsDefined("database", "auto_insert_item") {
		config.Database.AutoInsertItem = defaultDBConfig.AutoInsertItem
	}
	if !meta.IsDefined("database", "cache_size") {
		config.Database.CacheSize = defaultDBConfig.CacheSize
	}
	// Default master config
	defaultMasterConfig := *(*MasterConfig)(nil).LoadDefaultIfNil()
	if !meta.IsDefined("master", "port") {
		config.Master.Port = defaultMasterConfig.Port
	}
	if !meta.IsDefined("master", "host") {
		config.Master.Host = defaultMasterConfig.Host
	}
	if !meta.IsDefined("master", "http_port") {
		config.Master.HttpPort = defaultMasterConfig.HttpPort
	}
	if !meta.IsDefined("master", "http_host") {
		config.Master.HttpHost = defaultMasterConfig.HttpHost
	}
	if !meta.IsDefined("master", "fit_jobs") {
		config.Master.FitJobs = defaultMasterConfig.FitJobs
	}
	if !meta.IsDefined("master", "search_jobs") {
		config.Master.SearchJobs = defaultMasterConfig.SearchJobs
	}
	if !meta.IsDefined("master", "meta_timeout") {
		config.Master.MetaTimeout = defaultMasterConfig.MetaTimeout
	}
	// Default server config
	defaultServerConfig := *(*ServerConfig)(nil).LoadDefaultIfNil()
	if !meta.IsDefined("server", "api_key") {
		config.Server.APIKey = defaultServerConfig.APIKey
	}
	if !meta.IsDefined("server", "default_n") {
		config.Server.DefaultN = defaultServerConfig.DefaultN
	}
	// Default recommend config
	defaultRecommendConfig := *(*RecommendConfig)(nil).LoadDefaultIfNil()
	if !meta.IsDefined("recommend", "popular_window") {
		config.Recommend.PopularWindow = defaultRecommendConfig.PopularWindow
	}
	if !meta.IsDefined("recommend", "fit_period") {
		config.Recommend.FitPeriod = defaultRecommendConfig.FitPeriod
	}
	if !meta.IsDefined("recommend", "max_recommend_period") {
		config.Recommend.MaxRecommendPeriod = defaultRecommendConfig.MaxRecommendPeriod
	}
	if !meta.IsDefined("recommend", "search_period") {
		config.Recommend.SearchPeriod = defaultRecommendConfig.SearchPeriod
	}
	if !meta.IsDefined("recommend", "search_epoch") {
		config.Recommend.SearchEpoch = defaultRecommendConfig.SearchEpoch
	}
	if !meta.IsDefined("recommend", "search_trials") {
		config.Recommend.SearchTrials = defaultRecommendConfig.SearchTrials
	}
}

// LoadConfig loads configuration from toml file.
func LoadConfig(path string) (*Config, *toml.MetaData, error) {
	var conf Config
	metaData, err := toml.DecodeFile(path, &conf)
	if err != nil {
		return nil, nil, err
	}
	conf.FillDefault(metaData)
	return &conf, &metaData, nil
}
