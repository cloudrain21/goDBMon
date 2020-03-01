package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type configManager interface {
	readConfig() error
	showAllDBConfig()
}

type xmlConfigManager struct {
	configFile string
	db         dbConfig
}

type jsonConfigManager struct {
	configFile string
	db         dbConfig
}

type dbConfig struct {
	dbName           string
	commonConfig     map[string]string
	connectionConfig map[string]string
	monQueryConfig   []map[string]string
}

type database_mon struct {
	common_configs
	mon_target_dbs
}

type common_configs struct {
	query_interval_sec int
}

type mon_target_dbs struct {
	targetDBs []database
}

type database struct {
	name             string
	driver_name      string
	ip               string
	port             int
	user             string
	pass             string
	conn_opt         string
	mon_query_config string
}

func (c *xmlConfigManager) readConfig() error {
	fmt.Printf("ReadConfig function : %s\n", c.configFile)
	f, err := os.Open(c.configFile)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	v := database_mon{}
	err = xml.Unmarshal([]byte(data), &v)
	if err != nil {
		return err
	}

	fmt.Println(v)

	return nil
}

func (c *xmlConfigManager) showAllDBConfig() {
	fmt.Printf("showAllDBConfig function : %s\n", c.configFile)
}

func (c *jsonConfigManager) readConfig() error {
	fmt.Println("ReadConfig function : json")
	return nil
}

func (c *jsonConfigManager) showAllDBConfig() {
	fmt.Println("ShowDBConfig function : json")
}
