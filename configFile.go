package main

type ConfigFile struct {
	Port            int `yaml:"port,omitempty"`
	SharedSecret    string `yaml:"shared_secret,omitempty"`
	MongoDbServer   string `yaml:"mongodb_server,omitempty"`
	MongoDbDatabase string `yaml:"mongodb_database,omitempty"`
}