package db
import (
	"strings"
	"encoding/json"
	"os"
	"database/common"
	"fmt"
)

func Export() (error) {
	var configs interface{}
	fr, err := os.Open("./config.json")
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(fr)
	err = decoder.Decode(&configs)
	if err != nil {
		return err
	}
	confs := configs.(map[string]interface{})
	sqlPath := confs["sqlPath"].(string)
	ch := make(chan string)
	for key, value := range confs {
		if strings.HasPrefix(key, "db_") {
			dbConf := value.(map[string]interface{})
			dbConn := common.DbConnFields{
				DbHost:    dbConf["db_host"].(string),
				DbPort:    dbConf["db_port"].(string),
				DbUser:    dbConf["db_user"].(string),
				DbPass:    dbConf["db_pass"].(string),
				DbName:    dbConf["db_name"].(string),
			}
			go backupDb(dbConn.DbHost, dbConn.DbPort, dbConn.DbUser, dbConn.DbPass,dbConn.DbName,"",sqlPath,ch)
		}
	}
	for key := range confs {
		if strings.HasPrefix(key, "db_") {
			fmt.Print( <- ch )
		}
	}
	return nil
}
