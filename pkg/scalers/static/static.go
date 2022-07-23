package static

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *Scaler) HandleScale(writer http.ResponseWriter, request *http.Request) {
	var scale scaleResponse
	if !s.down {
		log.Println("Handling down scale")
		scale = scaleResponse{
			Name:     "foo",
			Replicas: -10000000,
			Region:   "east",
			Primary:  false,
		}
	} else {
		log.Println("Handling up scale")
		scale = scaleResponse{
			Name:     "foo",
			Replicas: 100000,
			Region:   "east",
			Primary:  true,
		}
	}
	response, err := json.Marshal(scale)
	if err != nil {
		log.Printf("err: %v", err)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	_, _ = writer.Write(response)
}

func (s *Scaler) HandleSwap(writer http.ResponseWriter, request *http.Request) {
	log.Println("Handling Swap")
	old := s.swap()
	response := fmt.Sprintf("From %t to %t\n", old, s.down)
	_, _ = writer.Write([]byte(response))
}

func (s *Scaler) swap() bool {
	s.readWriteLock.Lock()
	defer s.readWriteLock.Unlock()
	s.down = !s.down
	log.Printf("swap: %t\n", s.down)
	return !s.down
}

func NewStaticScaler() *Scaler {
	return &Scaler{
		down: false,
	}
}
