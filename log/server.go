package log

import (stlog "log"
		"os"
		"net/http"
		"io/ioutil"
)

var Log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0600)
	if err != nil{
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

func Run(destination string) {
	Log = stlog.New(fileLog(destination), "Go", stlog.LstdFlags)
}

func RegisterHandlers(){
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request){//针对/log路径
		switch r.Method{
		case "POST":
			msg, err := ioutil.ReadAll(r.Body)
			if err != nil || len(msg) == 0{
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(message string){
	Log.Printf("%v\n", message)
}