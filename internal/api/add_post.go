package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/v1tbrah/media-service/mpbapi"
)

func (a *API) AddPost(ctx context.Context, req *mpbapi.AddPostRequest) (*mpbapi.AddPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, ErrEmptyRequest.Error())
	}

	guid, err := a.storage.AddPost(ctx, req.GetData())
	if err != nil {
		log.Err(err).Msg("storage.AddPost")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &mpbapi.AddPostResponse{Guid: guid}, nil
}
