package api

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/v1tbrah/media-service/internal/api/mocks"
	"github.com/v1tbrah/media-service/mpbapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAPI_AddPost(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name            string
		mockStorage     func(t *testing.T) *mocks.Storage
		req             *mpbapi.AddPostRequest
		expectedGuid    string
		wantErr         bool
		expectedErr     error
		expectedErrCode codes.Code
	}{
		{
			name: "nil req",
			mockStorage: func(t *testing.T) *mocks.Storage {
				return mocks.NewStorage(t)
			},
			req:             nil,
			wantErr:         true,
			expectedErr:     ErrEmptyRequest,
			expectedErrCode: codes.InvalidArgument,
		},
		{
			name: "error on AddPost",
			req:  &mpbapi.AddPostRequest{Data: []byte("test_data")},
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				testStorage.On("AddPost",
					mock.MatchedBy(func(ctx context.Context) bool { return true }), []byte("test_data")).
					Return("", errors.New("unexpected error")).
					Once()
				return testStorage
			},
			wantErr:         true,
			expectedErr:     errors.New("unexpected error"),
			expectedErrCode: codes.Internal,
		},
		{
			name: "OK",
			req:  &mpbapi.AddPostRequest{Data: []byte("test_data")},
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				testStorage.On("AddPost",
					mock.MatchedBy(func(ctx context.Context) bool { return true }), []byte("test_data")).
					Return("12345678", nil).
					Once()
				return testStorage
			},
			expectedGuid: "12345678",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{storage: tt.mockStorage(t)}
			resp, err := a.AddPost(ctx, tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErrCode, status.Code(err))
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			}

			if !tt.wantErr {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedGuid, resp.GetGuid())
			}
		})
	}
}
