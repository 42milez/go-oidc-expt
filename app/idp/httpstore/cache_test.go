package httpstore

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"go.uber.org/mock/gomock"
)

func TestCache_Restore(t *testing.T) {
	t.Parallel()

	clock := xtestutil.FixedClocker{}

	wantSid := typedef.SessionID(484493849343885677)
	wantUserId := typedef.UserID(484493849343820141)
	wantAuthTime := clock.Now()

	cacheRWMock := NewMockCacheReadWriter(gomock.NewController(t))
	cacheRWMock.EXPECT().ReadHash(gomock.Any(), gomock.Any(), userIdFieldName).Return(wantUserId.String(), nil).AnyTimes()
	cacheRWMock.EXPECT().ReadHash(gomock.Any(), gomock.Any(), authTimeFieldName).Return(strconv.FormatInt(wantAuthTime.Unix(), 10), nil).AnyTimes()

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

	gotSess, ok := ctx.Value(SessionKey{}).(*Session)
	if !ok {
		t.Fatal(xerr.TypeAssertionFailed)
	}
	if wantUserId != gotSess.UserID {
		t.Errorf("want = %d; got = %d", wantUserId, gotSess.UserID)
	}
	if !wantAuthTime.Equal(gotSess.AuthTime) {
		t.Errorf("want = %v; got = %v", wantAuthTime, gotSess.AuthTime)
	}
}

func TestCache_CreateSession(t *testing.T) {
	t.Parallel()

	wantSid := typedef.SessionID(484493849344016749)
	uid := typedef.UserID(484493849344082285)

	ctrl := gomock.NewController(t)

	cacheRWMock := NewMockCacheReadWriter(ctrl)
	cacheRWMock.EXPECT().WriteHash(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	idGenMock := iface.NewMockIdGenerator(ctrl)
	idGenMock.EXPECT().NextID().Return(uint64(wantSid), nil).AnyTimes()

	cache := &Cache{
		repo:  cacheRWMock,
		idGen: idGenMock,
	}
	ctx := context.Background()

	gotSid, err := cache.CreateSession(ctx, uid)
	if err != nil {
		t.Fatal(err)
	}
	if wantSid != gotSid {
		t.Errorf("want = %d; got =%d", wantSid, gotSid)
	}
}
