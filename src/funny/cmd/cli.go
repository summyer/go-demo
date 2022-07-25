package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
	"reflect"
)

func main() {
	app := cli.NewApp()
	app.Name = "funny tool CLI"
	app.Usage = "let's you do some funny thing "
	var mHost string
	//多个参数
	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "host",          //option
			Value:       "www.baidu.com", //default arg
			Usage:       "server host",
			Destination: &mHost,
		},
		&cli.StringFlag{
			Name:        "ip",            //option
			Value:       "www.baidu.com", //default arg
			Usage:       "server host",
			Destination: &mHost,
		},
		&cli.IntFlag{
			Name:  "port", //option
			Value: 80,     //default arg
			Usage: "server port",
		},
	}
	//cmd.exe ns --host www.baidu.com
	var nsCommand cli.Command = cli.Command{
		Name:  "ns",
		Usage: "this is tool for ns",
		Flags: myFlags,
		Action: func(context *cli.Context) error {
			lHost := context.String("host")
			lPort := context.Int("port")
			fmt.Printf("current args:host[%s],port:[%d]\n", lHost, lPort)
			ns, err := net.LookupIP(lHost)
			if err != nil {
				return err
			}
			for i := 0; i < len(ns); i++ {
				fmt.Println("host:", ns[i])
			}
			return nil
		},
	}
	//var a *[]cli.Command = new([]cli.Command)
	//方法一
	app.Commands = make([]*cli.Command, 1)
	app.Commands[0] = &nsCommand
	//方法2  定义指针,指定指针类型都用*; 取变量指针都用&
	app.Commands = []*cli.Command{
		&nsCommand,
	}
	log.Println(os.Args)
	fmt.Println(reflect.TypeOf(os.Args))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
