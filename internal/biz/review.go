package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/open-portfolios/review/internal/data/model"
)

var (
	ErrAlreadyPublished = errors.New("review has beem published already")
)

type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error)
	GetReviewByOrderID(context.Context, int64) ([]*model.ReviewInfo, error)
}

type ReviewUsecase struct {
	repo  ReviewRepo
	flake SnowflakeRepo
	log   *log.Helper
}

func NewReviewUsecase(repo ReviewRepo, flake SnowflakeRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{
		repo:  repo,
		flake: flake,
		log:   log.NewHelper(logger),
	}
}

func (uc *ReviewUsecase) CreateReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("biz.CreateReview %v", review.OrderID)

	// Check existence
	exist, err := uc.repo.GetReviewByOrderID(ctx, review.OrderID)
	if err != nil {
		return nil, err
	}
	if len(exist) > 0 {
		return nil, ErrAlreadyPublished
	}

	// Generate Snowflake ID
	review.ReviewID, err = uc.flake.Generate(ctx)
	if err != nil {
		return nil, err
	}

	return uc.repo.SaveReview(ctx, review)
}
