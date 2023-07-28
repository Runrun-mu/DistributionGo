package registry
import (
	"sync"
	"net/http"
	"encoding/json"
	"log"
)

const ServerPort = ":3000"
const Serveices = "http://localhost" + ServerPort + "/services"

type registry struct {
	registrations []Registration
	mutex *sync.Mutex
}

func(r *registry) add(reg Registration){
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
}

var reg = registry{
	registrations: make([]Registration, 0),
	mutex: new(sync.Mutex),
}

type RegistryService struct {}

func(s RegistryService)ServeHTTP(w http.ResponseWriter, r *http.Request){
	log.Println("Request received")
	switch r.Method{
		case http.MethodPost:
			dec := json.NewDecoder(r.Body)
			var r Registration
			err := dec.Decode(&r)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Printf("Adding service: %v with URL: %s\n", r.ServiceName, r.ServiceURL)
			reg.add(r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
	}
}