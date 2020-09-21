package main
import (
	"encoding/json"
	"fmt"
	"database/db"
	"github.com/robfig/cron/v3"
	"os"
)
func main() {
	fmt.Println("start database backup ")
	var configs interface{}
	fr, err := os.Open("./config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	decoder := json.NewDecoder(fr)
	err = decoder.Decode(&configs)
	if err != nil {
		fmt.Println(err)
		return
	}
	confs := configs.(map[string]interface{})
	spec := confs["spec"].(string)
	c := cron.New(cron.WithSeconds())
	_, _ = c.AddFunc(spec, func() {
		err := db.Export(confs)
		if err != nil {
			fmt.Println(err)
		}
	})
	c.Start()
	select {

	}

}
