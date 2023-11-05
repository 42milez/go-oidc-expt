package httpstore

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/42milez/go-oidc-server/app/iface"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/typedef"
	"go.uber.org/mock/gomock"
)

func TestCache_Restore(t *testing.T) {
	t.Parallel()

	wantSid := typedef.SessionID(484493849343885677)
	wantUserId := typedef.UserID(484493849343820141)

	cacheRWMock := NewMockCacheReadWriter(gomock.NewController(t))
	cacheRWMock.EXPECT().ReadHash(gomock.Any(), gomock.Any(), gomock.Any()).Return(wantUserId.String(), nil).AnyTimes()

	cache := &Cache{
		repo: cacheRWMock,
	}
	req := httptest.NewRequest(http.MethodGet, "https://example.com", nil)

	gotReq, err := cache.Restore(req, wantSid)
	if err != nil {
		t.Fatal(err)
	}

	ctx := gotReq.Context()

	gotSid, ok := ctx.Value(typedef.SessionIdKey{}).(typedef.SessionID)
	if !ok {
		t.Fatal(xerr.TypeAssertionFailed)
	}
	if wantSid != gotSid {
		t.Errorf("want = %d; got = %d", wantSid, gotSid)
	}

	gotUserId, ok := ctx.Value(typedef.UserIdKey{}).(typedef.UserID)
	if !ok {
		t.Fatal(xerr.TypeAssertionFailed)
	}
	if wantUserId != gotUserId {
		t.Errorf("want = %d; got = %d", gotUserId, gotUserId)
	}
}

func TestCache_WriteUserInfo(t *testing.T) {
	t.Parallel()

	wantSid := typedef.SessionID(484493849344016749)
	uid := typedef.UserID(484493849344082285)

	ctrl := gomock.NewController(t)

	cacheRWMock := NewMockCacheReadWriter(ctrl)
	cacheRWMock.EXPECT().WriteHash(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()

	idGenMock := iface.NewMockIdGenerator(ctrl)
	idGenMock.EXPECT().NextID().Return(uint64(wantSid), nil).AnyTimes()

	cache := &Cache{
		repo:  cacheRWMock,
		idGen: idGenMock,
	}
	ctx := context.Background()

	gotSid, err := cache.WriteUserInfo(ctx, uid)
	if err != nil {
		t.Fatal(err)
	}
	if wantSid != gotSid {
		t.Errorf("want = %d; got =%d", wantSid, gotSid)
	}
}
