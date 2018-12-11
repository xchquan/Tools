package Tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pborman/uuid"
)

type TgTMInterval struct {
	TmFirst time.Time
	TmSec   time.Time
}

func InitTgTMinterval() *TgTMInterval {
	return &TgTMInterval{
		TmFirst: time.Now(),
		TmSec:   time.Now(),
	}
}

func (tgTmT *TgTMInterval) GetIntervalMis() int64 {
	tgTmT.TmSec = time.Now()

	return (tgTmT.TmSec.UnixNano() - tgTmT.TmFirst.UnixNano()) / 1000000
}

func GetTimeNow(iFmt int) string {
	tmNow := time.Now()
	var strFmt string = "2006-01-02 15:04:05"
	switch iFmt {
	case 1:
		strFmt = "20060102150405"
	case 2:
		strFmt = "2006-01-02 15:04:05"
	case 3:
		return fmt.Sprintf(GetTimeNow(2)+".%03d", tmNow.Nanosecond()/1000000)
	}
	return tmNow.Format(strFmt)
}

func GetExeDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1) + "/"
}

/// get new uuid
func GetNewUUID() string {
	return uuid.New()
}

/// get buf from special file
func GetCfgBuf(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
