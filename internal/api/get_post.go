package api

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/v1tbrah/media-service/internal/storage"
	"github.com/v1tbrah/media-service/mpbapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *API) GetPost(ctx context.Context, req *mpbapi.GetPostRequest) (*mpbapi.GetPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, ErrEmptyRequest.Error())
	}

	data, err := a.storage.GetPost(ctx, req.GetGuid())
	if err != nil {
		if errors.Is(err, storage.ErrNotFoundByGUID) {
			return nil, status.Error(codes.NotFound, storage.ErrNotFoundByGUID.Error())
		}

		log.Err(err).Str("guid", req.GetGuid()).Msg("storage.GetPost")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &mpbapi.GetPostResponse{Data: data}, nil
}
