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
	s.log.WithContext(ctx).Debugf("service.CreateReview %v", req.OrderID)

	var anonymous int32
	if req.Anonymous {
		anonymous = 1
	}
	review, err := s.uc.CreateReview(ctx, &model.ReviewInfo{
		UserID:       req.UserID,
		OrderID:      req.OrderID,
		Score:        req.Score,
		ServiceScore: req.ServiceScore,
		ExpressScore: req.ExpressScore,
		Content:      req.Content,
		PicInfo:      req.PicInfo,
		VideoInfo:    req.VideoInfo,
		Anonymous:    anonymous,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateReviewReply{ReviewID: review.ReviewID}, nil
}
