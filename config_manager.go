package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var QueryInterval int

type configManager interface {
	readConfig() error
	showAllDBConfig()
}

type DBConfigMap map[string]string

type xmlConfigManager struct {
	configFile string
	dbMap      map[string]DBConfigMap
}

type jsonConfigManager struct {
	configFile string
	dbMap      map[string]Database
}

type DatabaseMon struct {
	CommonConfigs CommonConfigs
	MonTargetDBs  MonTargetDBs
}

type CommonConfigs struct {
	QueryIntervalSec int
}

type MonTargetDBs struct {
	Database []Database
}

type Database struct {
	Name           string
	DriverName     string
	IP             string
	Port           int
	User           string
	Pass           string
	ConnOpt        string
	MonQueryConfig string
}

func (c *xmlConfigManager) readConfig() error {
	fmt.Printf("ReadConfig function : %s\n", filepath.Base(c.configFile))
	f, err := os.Open(c.configFile)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	var DatabaseMon DatabaseMon
	err = xml.Unmarshal(bytes, &DatabaseMon)
	if err != nil {
		return err
	}

	QueryInterval = DatabaseMon.CommonConfigs.QueryIntervalSec

	for _, cfg := range DatabaseMon.MonTargetDBs.Database {
		dbName := strings.TrimSpace(cfg.Name)
		c.dbMap[dbName] = RegOneDBConfig(cfg)
	}

	fmt.Println(c.dbMap["mysql"])

	return nil
}

func RegOneDBConfig(database Database) map[string]string {
	m := make(map[string]string)

	fmt.Printf("%s\n", reflect.ValueOf(&database).Elem().Type().Field(0).Tag)

	return m
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
