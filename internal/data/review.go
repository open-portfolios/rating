package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/open-portfolios/review/internal/biz"
	"github.com/open-portfolios/review/internal/data/model"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewRepo) SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	if err := r.data.q.ReviewInfo.WithContext(ctx).Save(review); err != nil {
		return nil, err
	}
	return review, nil
}

func (r *reviewRepo) GetReviewByOrderID(ctx context.Context, orderID int64) ([]*model.ReviewInfo, error) {
	t := r.data.q.ReviewInfo
	return t.WithContext(ctx).Where(t.OrderID.Eq(orderID)).Find()
}
