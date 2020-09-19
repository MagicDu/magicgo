package main
import (
	"fmt"
	"database/db"
	"github.com/robfig/cron/v3"
)
func main() {

	c := cron.New(cron.WithSeconds())
	spec := "*/50 * * * * ?"
	c.AddFunc(spec, func() {
		err := db.Export()
		if err != nil{
			fmt.Println(err)
		}
	})
	c.Start()
	select {

	}

}
