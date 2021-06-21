package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rushminator/carz/data"
	"github.com/rushminator/carz/pb"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type Cars struct {
	log *zap.Logger
}

// NewProducts creates a products handler with the given logger
func NewCars(l *zap.Logger) *Cars {
	return &Cars{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (c *Cars) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/Carz/")

	c.log.Info("Success..",
		zap.String("ID", string(id)))

	if id != "" {
		enc := r.Header.Get("accept-encoding")
		if enc == "application/protobuf" || enc == "" {
			c.log.Info("proto")
			response, err := proto.Marshal(c.findCar(id))
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			rw.Write([]byte(response))
		} else if enc == "application/json" {
			c.log.Info("json")
			response, err := json.Marshal(c.findCar(id))
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			rw.Write([]byte(response))
		} else {
			c.log.Info("none")
			rw.WriteHeader(http.StatusUnsupportedMediaType)
			rw.Write([]byte("Unsupported Media Type"))
		}
	} else {
		c.getCars(rw, r)
	}
}

func (c *Cars) getCars(rw http.ResponseWriter, r *http.Request) {
	carList := data.GetCars()
	data, err := json.Marshal(carList)
	if err != nil {
		c.log.Info("Server Error")
		http.Error(rw, "Error", http.StatusInternalServerError)

	}
	rw.Write(data)
}

func (c *Cars) findCar(id string) *pb.Car {
	carList := data.GetCars()
	for i := range carList {
		if carList[i].Id == id {
			return carList[i] // Found!
		}
	}
	return nil
}
