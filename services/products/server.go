package main

import (
	"context"
	pb "github.com/adfinitas-app/prometer/models/products"
)

type server struct {
	pb.UnimplementedProductServiceServer
}

var products = []*pb.ProductModel{
	{
		Id:          "1",
		Title:       "First Product",
		Description: "Awesome product",
	},
	{
		Id:          "2",
		Title:       "Second Product",
		Description: "The next generation product",
	},
}

func (s *server) GetProducts(_ context.Context, in *pb.ProductsRequest) (*pb.ProductsResponse, error) {
	if len(in.ProductIds) == 0 {
		return &pb.ProductsResponse{Products: products}, nil
	}

	var res []*pb.ProductModel
	for _, id := range in.ProductIds {
		for _, p := range products {
			if p.Id == id {
				res = append(res, p)
			}
		}
	}
	return &pb.ProductsResponse{Products: res}, nil
}
