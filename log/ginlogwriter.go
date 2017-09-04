//gin的日志参数

package log

import (
	"errors"
	"os"
	"sync"
	"time"
)

type GinLoggerWriter struct {
	file     *os.File
	lastDate *time.Time
	mu       *sync.RWMutex
}

const (
	DATEFORMATE = "2006-01-02"
)

var ginlogdir string

/*
  #Created by Luheng on 2017/5/9.
  #Arguments:
  #Return:
  #Description: 初始化
*/
func (l *GinLoggerWriter) InitGinLog() error {
	if err := GinCreateDir("logs/gin"); err != nil {
		Log.Error("Gin日志初始化 创建文件假错误", err)
		return err
	}
	now := time.Now().Format(DATEFORMATE)
	fileName := ginlogdir + "/" + "gin" + "." + now
	f, erropen := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if erropen != nil {
		Log.Error("Gin日志初始化 打开文件出错误", erropen)
		return erropen
	}
	l.file = f
	t, errt := time.Parse(DATEFORMATE, now)
	if errt != nil {
		Log.Error("Gin日志初始化 日期格式化错误", errt)
		return errt
	}
	l.lastDate = &t
	return nil
}

/*
  #Created by Luheng on 2017/5/9.
  #Arguments:
  #Return:
  #Description: 按日期切割
*/
func isNewDay(l *GinLoggerWriter) bool {
	now := time.Now().Format(DATEFORMATE)
	t, err := time.Parse(DATEFORMATE, now)
	if err != nil {
		return true
	}
	if l.lastDate != nil {
		return t.After(*l.lastDate)
	}
	return true
}

/*
  #Created by Luheng on 2017/5/9.
  #Arguments:
  #Return:
  #Description: 创建日志目录
*/
func GinCreateDir(dir string) (e error) {
	ginlogdir = dir
	_, er := os.Stat(dir)
	b := er == nil || os.IsExist(er)
	if !b {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			if os.IsPermission(err) {
				e = err
			}
		}
	}
	return e
}

/*
  #Created by Luheng on 2017/5/9.
  #Arguments:
  #Return:
  #Description: 写文件
*/
func (l *GinLoggerWriter) Write(p []byte) (int, error) {
	if isNewDay(l) {
		now := time.Now().Format(DATEFORMATE)
		l.mu = new(sync.RWMutex)
		l.mu.Lock()
		defer l.mu.Unlock()
		if l.file != nil {
			l.file.Close()
		}
		t, _ := time.Parse(DATEFORMATE, now)
		l.lastDate = &t
		fileName := ginlogdir + "/" + "gin" + "." + now
		f, _err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if _err == nil {
			l.file = f
		} else {
			Log.Error("Gin日志 打开文件错误", _err.Error())
		}
	}
	if l.file != nil {
		cnt, err := l.file.Write(p)
		if err != nil {
			Log.Error("Gin日志 写日志错误", err.Error())
		}
		return cnt, err
	}
	Log.Error("Gin日志 文件描述符为空")
	return 0, errors.New("文件描述符为空")
}
