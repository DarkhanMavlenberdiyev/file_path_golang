package main

import (
	"fmt"
	"strings"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"runtime"
)

var (
configFilePath = ""
h = '\u2500'
last = '\u2516'
mid = '\u2520'
v = '\u2503'
ress = ""
flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Destination: &configFilePath,
		},
	}
)


var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset  = ""
		Red    = ""
		Green  = ""
		Yellow = ""
		Blue   = ""
		Purple = ""
		Cyan   = ""
		Gray   = ""
		White  = ""
	}
}

func main(){
	app := cli.NewApp()
	app.Flags = flags
	app.Commands = cli.Commands{
		&cli.Command{
			Name: "printdir",
			Action: PrintDir,
		},
	}
	app.Run(os.Args)
}


func PrintDir(c *cli.Context) error {
	print(configFilePath,0)
	lis := strings.Split(ress,"\n")
	lis2 := make([][]string,0)
	for _,c := range lis {
		lis2 = append(lis2,strings.Split(c,""))
	}
	for i:=0;i<len(lis2);i++  {
		for j:=0;j<len(lis2[i]);j++{
			if i<len(lis2)-2 && (lis2[i][j] ==string(mid) || lis2[i][j]==string(v)) && lis2[i+1][j]==" " {
				lis2[i+1][j]=string(v)
			}
		}
	}
	ress = ""
	for _,li := range lis2 {
		ress+= strings.Join(li,"")+"\n"
	}
	fmt.Println(ress)
	return nil
}

func print(s string, n int) {
	dir,er := ioutil.ReadDir(s)
	if er == nil {
		for i,d := range dir {
			res := ""
			for j:=0;j<n;j++{
				res+=" "
			}
			if i==len(dir)-1 {
				res += string(last)
			}else{
				res += string(mid)
			}
			if d.IsDir() {
				res+= Blue + string(h)+d.Name() + Reset
				ress+=res+"\n"
			}else{
				res+= Yellow + string(h)+string(h)+d.Name() + Reset
				ress+=res+"\n"
			}
			print(s+"/"+d.Name(),n+1)
		}
	}
}



