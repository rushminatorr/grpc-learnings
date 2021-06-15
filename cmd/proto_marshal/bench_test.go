package main

import (
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rushminator/carz/pb"
)

var (
	car = &pb.Car{
		Id:     "95",
		Driver: "Lightning McQueen",
		Status: pb.Status_FREE,
		Location: &pb.Location{
			Lat: -22.95192,
			Lng: -43.21050,
		},
		Updated: timestamppb.Now(),
	}
)

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(car)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(car)
		if err != nil {
			b.Fatal(err)
		}
	}
}
