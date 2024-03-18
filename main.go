package main

import (
	"flag"
	"fmt"
	"os"

	"bngo/inits"
	"bngo/utils"

	"github.com/open-binance/logger"
)

var (
	appName string
)

func init() {
	utils.ServerTime()
	fmt.Println("开始程序.....")

}

func main() {
	//启动时传入参数，如果不传用默认值
	//go run main.go -c config/cfg.local.yaml -v
	cfgFile := flag.String("c", "config/cfg.local.yaml", "配置文件")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	// show version
	if *version {
		fmt.Printf("app name:\t%s\n", inits.AppName)
		fmt.Printf("version:\t%s\n", inits.Version)
		os.Exit(0)
	}

	//load config
	if err := inits.LoadConfig(*cfgFile); err != nil {
		fmt.Printf("msg=failed to load config||err=%s\n", err.Error())
		os.Exit(1)
	}

	// init logger
	logCfg := logger.NewDefaultCfg()
	if err := logCfg.SetLevel(inits.Config.Log.Level); err != nil {
		fmt.Printf("msg=failed to set log level||err=%s\n", err.Error())
		os.Exit(1)
	}

	// logCfg.Level = "debug"
	if err := logger.Init(logCfg); err != nil {
		fmt.Printf("msg=failed to init logger||err=%s\n", err.Error())
		os.Exit(1)
	}
	defer logger.Sync() // flushes buffer, called in main function is reasonable
	logger.Infof("msg=succeed to init logger||max_size=%d||max_age=%d||max_backups=%d||level=%s||path=%s||encoding=%s",
		logCfg.MaxSize, logCfg.MaxAge, logCfg.MaxBackups, logCfg.Level, logCfg.Path, logCfg.Encoding)
	if err := inits.InitFileLogger(inits.Config.File); err != nil {
		logger.Errorf("msg=failed to init file logger||err=%s", err.Error())
		fmt.Printf("msg=failed to init file logger||err=%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(logCfg.MaxSize)

}
