package main

import (
	"context"
	"fmt"
	pb "github.com/adfinitas-app/prometer/models/reviews"
)

type server struct {
	pb.UnimplementedReviewServiceServer
}

var reviews = []*pb.ReviewModel{
	{
		Id:        "1",
		ProductId: "1",
		UserName:  "User 1",
		Rating:    8,
		Text:      "Amaizing",
	},
	{
		Id:        "1",
		ProductId: "2",
		UserName:  "User 24",
		Rating:    7,
		Text:      "Amaizing !",
	},
	{
		Id:        "3",
		ProductId: "1",
		UserName:  "User 3",
		Rating:    1,
		Text:      "Nul",
	},
	{
		Id:        "4",
		ProductId: "1",
		UserName:  "User 0283",
		Rating:    10,
		Text:      "Perfect",
	},
}

func (s *server) GetReview(_ context.Context, in *pb.ReviewRequest) (*pb.ReviewResponse, error) {
	if in.ReviewId == "" {
		return &pb.ReviewResponse{Review: nil}, nil
	}

	for _, r := range reviews {
		if r.Id == in.ReviewId {
			return &pb.ReviewResponse{Review: r}, nil
		}
	}
	return &pb.ReviewResponse{Review: nil}, nil
}

func (s *server) GetReviews(_ context.Context, in *pb.ReviewsRequest) (*pb.ReviewsResponse, error) {
	if len(in.ReviewIds) == 0 {
		return &pb.ReviewsResponse{Reviews: reviews}, nil
	}

	var res []*pb.ReviewModel
	for _, id := range in.ReviewIds {
		for _, r := range reviews {
			if r.Id == id {
				res = append(res, r)
			}
		}
	}
	return &pb.ReviewsResponse{Reviews: res}, nil
}

func (s *server) GetReviewsByProducts(_ context.Context, in *pb.ProductsRequest) (*pb.ReviewsResponse, error) {
	if len(in.ProductIds) == 0 {
		return nil, fmt.Errorf("ProductIds is empty")
	}

	var res []*pb.ReviewModel
	for _, id := range in.ProductIds {
		for _, r := range reviews {
			if r.ProductId == id {
				res = append(res, r)
			}
		}
	}
	return &pb.ReviewsResponse{Reviews: res}, nil
}

func (s *server) GetReviewsByProduct(_ context.Context, in *pb.ProductRequest) (*pb.ReviewsResponse, error) {
	if in.ProductId == "" {
		return nil, fmt.Errorf("ProductId is empty")
	}

	var res []*pb.ReviewModel
	for _, r := range reviews {
		if r.ProductId == in.ProductId {
			res = append(res, r)
		}
	}
	return &pb.ReviewsResponse{Reviews: res}, nil
}
