package db

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)



func backupDb(host,port,user,password,databaseName,tableName,sqlPath string,ch chan <- string)  {
	var cmd *exec.Cmd
	if tableName=="" {
		cmd=exec.Command("mysqldump","--opt","-h"+host,"-P"+port,"-u"+user,"-p"+password,databaseName)
	}else {
		cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName, tableName)
	}
	stdout, err := cmd.StdoutPipe()
	if err!=nil {
		ch <- fmt.Sprintln("Error: ", databaseName, "\t export views throw, \t", err)
		return

	}
	if err:=cmd.Start();err!=nil{
		ch <- fmt.Sprintln("Error: ", databaseName, "\t export views throw, \t", err)
		return
	}
	bytes,err:=ioutil.ReadAll(stdout)
	if err!=nil{
		ch <- fmt.Sprintln("Error: ", databaseName, "\t export views throw, \t", err)
		return
	}
	now:=time.Now().Format("20060102150405")
	var backuppath string
	if tableName==""{
		backuppath=sqlPath+databaseName+"_"+now+".sql"
	}else{
		backuppath=sqlPath+databaseName+"_"+tableName+"_"+now+".sql"
	}
	err=ioutil.WriteFile(backuppath,bytes,0644)
	if err!=nil {
		ch <- fmt.Sprintln("Error: ", databaseName, "\t export views throw, \t", err)
		return
	}
	ch <- fmt.Sprintln("Export ", databaseName, "\t success at \t", time.Now().Format("2006-01-02 15:04:05"))
}