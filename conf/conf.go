//Package setting provide goil's settings
package conf

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Unknwon/goconfig"
	"github.com/howeyc/fsnotify"
	"github.com/artong0416/goweb/daemon"
)

const (
	APP_VER = "0.10"
)

var (
	AppName   string
	AppHost   string
	AppVer    string
	IsProMode bool
	TimeZone  string
)

var (
	Cfg *goconfig.ConfigFile
)

var (
	AppConfPath = "conf/app.ini"
	HookReload  []func()
	//UrlList     []string
)

// LoadConfig loads configuration file.
func LoadConfig() *goconfig.ConfigFile {
	var err error

	if fh, _ := os.OpenFile(AppConfPath, os.O_RDONLY|os.O_CREATE, 0600); fh != nil {
		fh.Close()
	}

	// Load configuration, set app version and Log level.
	Cfg, err = goconfig.LoadConfigFile(AppConfPath)

	if err != nil {
		//log.Log.Error("Fail to load configuration file: " + err.Error())
		os.Exit(2)
	}

	//Cfg.BlockMode = false

	// set time zone of wetalk system
	TimeZone = Cfg.MustValue("app", "time_zone", "UTC")
	if _, err := time.LoadLocation(TimeZone); err == nil {
		os.Setenv("TZ", TimeZone)
	} else {
		//log.Log.Error("Wrong time_zone: " + TimeZone + " " + err.Error())
		os.Exit(2)
	}


	// Trim 4th part.
	AppVer = strings.Join(strings.Split(APP_VER, ".")[:2], ".")
	AppHost = Cfg.MustValue("app", "app_host", "11333")

	IsProMode = Cfg.MustValue("app", "run_mode") == "pro"

	if daemon.IsDebug {
		IsProMode = !daemon.IsDebug
	}

	reloadConfig()
	configWatcher()

	return Cfg
}

func reloadConfig() {
	AppName = Cfg.MustValue("app", "app_name", "WeTalk Community")

	//UrlList = strings.Split(Cfg.MustValue("urllist", "urls", ""), ",")

	for _, f := range HookReload {
		f()
	}
}

var eventTime = make(map[string]int64)

func configWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic("Failed start app watcher: " + err.Error())
	}

	go func() {
		for {
			select {
			case event := <-watcher.Event:
				switch filepath.Ext(event.Name) {
				case ".ini":
					if checkEventTime(event.Name) {
						continue
					}

					//Log.Info(event.String())
					if err := Cfg.Reload(); err != nil {
						//Log.Error("Conf Reload: ", err)
					}

					reloadConfig()
					//Log.Info("Config Reloaded")
				}
			}
		}
	}()

	if err := watcher.WatchFlags("conf", fsnotify.FSN_MODIFY); err != nil {
		//Log.Error(err.Error())
	}
}

// checkEventTime returns true if FileModTime does not change.
func checkEventTime(name string) bool {
	mt := getFileModTime(name)
	if eventTime[name] == mt {
		return true
	}

	eventTime[name] = mt
	return false
}

// getFileModTime retuens unix timestamp of `os.File.ModTime` by given path.
func getFileModTime(path string) int64 {
	path = strings.Replace(path, "\\", "/", -1)
	f, err := os.Open(path)
	if err != nil {
		//Log.Error("Fail to open file[ %s ]\n", err.Error())
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		//Log.Error("Fail to get file information[ %s ]\n", err.Error())
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

/*
  #Created by Luheng on 2017/5/23.
  #Arguments:
  #Return:
  #Description:配置初始化
*/
func init() {
	daemon.Daemon()
	LoadConfig()
}
