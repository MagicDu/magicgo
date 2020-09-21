package db

import (
	"database/common"
	"fmt"
	"strings"
)

func Export(confs map[string]interface{} ) (error) {
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
