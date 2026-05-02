package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/open-portfolios/review/api/review/v1"
	"github.com/open-portfolios/review/internal/biz"
	"github.com/open-portfolios/review/internal/data/model"
)

type ReviewService struct {
	pb.UnimplementedReviewServer

	uc  *biz.ReviewUsecase
	log *log.Helper
}

func NewReviewService(uc *biz.ReviewUsecase, logger log.Logger) *ReviewService {
	return &ReviewService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *ReviewService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewReply, error) {
	s.log.WithContext(ctx).Debugf("service.CreateReview %v", req)
	review, err := s.uc.CreateReview(ctx, &model.ReviewInfo{})
	if err != nil {
		return nil, err
	}
	return &pb.CreateReviewReply{ReviewID: review.ReviewID}, nil
}
func (s *ReviewService) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.UpdateReviewReply, error) {
	return &pb.UpdateReviewReply{}, nil
}
func (s *ReviewService) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewReply, error) {
	return &pb.DeleteReviewReply{}, nil
}
func (s *ReviewService) GetReview(ctx context.Context, req *pb.GetReviewRequest) (*pb.GetReviewReply, error) {
	return &pb.GetReviewReply{}, nil
}
func (s *ReviewService) ListReview(ctx context.Context, req *pb.ListReviewRequest) (*pb.ListReviewReply, error) {
	return &pb.ListReviewReply{}, nil
}
