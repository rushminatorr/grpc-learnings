package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rushminator/carz/pb"
)

func main() {
	//var loc pb.Location
	loc := pb.Location{
		Lat: -22.95192,
		Lng: -43.21050,
	}

	data, err := proto.Marshal(&loc)
	if err != nil {
		log.Fatal(err)
	}
	// os.Stdout.Write(data)
	fmt.Println("protobuf size:", len(data))

	jdata, err := json.Marshal(&loc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON size:", len(jdata))

	var loc2 pb.Location
	if err := proto.Unmarshal(data, &loc2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", &loc2)

	s := pb.Status_FREE
	fmt.Printf("free: %s (%d)\n", s, s)

	car := &pb.Car{
		Id:       "95",
		Driver:   "Lightening McQueen",
		Status:   pb.Status_FREE,
		Location: &loc,
		Updated:  timestamppb.Now(),
	}

	fmt.Println(car)
	fmt.Printf("%#v\n", car) // not readable

	t := car.Updated.AsTime()
	fmt.Println(t)

	jdata, err = protojson.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jdata))

	t := car.Updated.AsTime()
	fmt.Println(t)
	tz, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	us := t.In(tz)
	fmt.Println(us)

	data, err = proto.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("proto size: %d\n", len(data))

	jdata, err = json.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("JSON size: %d\n", len(jdata))
	fmt.Println("encoding/json", string(jdata))
	jdata, err = protojson.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("protojson size: %d\n", len(jdata))
	fmt.Println("protojson", string(jdata))

	json.NewEncoder(os.Stdout).Encode(time.Now())

}
