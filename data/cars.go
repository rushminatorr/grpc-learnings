package data

import "github.com/rushminator/carz/pb"

func GetCars() []*pb.Car {
	return cars
}

var cars = []*pb.Car{
	&pb.Car{
		Id:     "95",
		Driver: "Lightning McQueen",
		Status: pb.Status_FREE,
		Location: &pb.Location{
			Lat: 1,
			Lng: 1,
		},
	},
	&pb.Car{
		Id:     "301PCE",
		Driver: "Sally Carrera",
		Status: pb.Status_FREE,
		Location: &pb.Location{
			Lat: 1,
			Lng: 1,
		},
	},
	&pb.Car{
		Id:     "51HHMD",
		Driver: "Doc Hudson",
		Status: pb.Status_FREE,
		Location: &pb.Location{
			Lat: 2,
			Lng: 2,
		},
	},
	&pb.Car{
		Id:     "86",
		Driver: "Chick Hicks",
		Status: pb.Status_FREE,
		Location: &pb.Location{
			Lat: 2,
			Lng: 2,
		},
	},
}
