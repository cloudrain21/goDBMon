package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var QueryInterval int

type configManager interface {
	readConfig() error
	showAllDBConfig()
}

type xmlConfigManager struct {
	configFile string
	dbMap      map[string]Database
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

	r := strings.NewReader(string(bytes))
	parser := xml.NewDecoder(r)

	for {
		token, err := parser.Token()
		if err != nil {
			return fmt.Errorf("Could not parse token : %s", err.Error())
		}
		switch t := token.(type) {
		case xml.StartElement:
			elmt := xml.StartElement(t)
			name := elmt.Name.Local
			if name == "QueryIntervalSec" {
				decErr := parser.DecodeElement(&QueryInterval, &elmt)
				if decErr != nil {
					panic("Could not decode element" + decErr.Error())
				}
				fmt.Println(QueryInterval)
			} else if name == "Database" {
				database := &Database{}
				decErr := parser.DecodeElement(database, &elmt)
				if decErr != nil {
					panic("Could not decode element" + decErr.Error())
				}
				c.dbMap[database.Name] = *database
			}
		case xml.EndElement:
		case xml.CharData:
		case xml.Comment:
		case xml.ProcInst:
		case xml.Directive:
			continue
		default:
			return fmt.Errorf("Unknown token")
		}
	}

	return nil
}

func (c *xmlConfigManager) showAllDBConfig() {
	fmt.Printf("QueryInterval : %d\n", QueryInterval)

	for _, db := range c.dbMap {
		fmt.Println(db)
	}
}

func (c *jsonConfigManager) readConfig() error {
	fmt.Println("ReadConfig function : json")
	return nil
}

func (c *jsonConfigManager) showAllDBConfig() {
	fmt.Println("ShowDBConfig function : json")
}
