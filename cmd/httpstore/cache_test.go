package httpstore

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/xtestutil"

	"github.com/42milez/go-oidc-server/pkg/typedef"

	"github.com/42milez/go-oidc-server/cmd/iface"
	"github.com/42milez/go-oidc-server/pkg/xerr"
	"go.uber.org/mock/gomock"
)

func TestCache_Restore(t *testing.T) {
	t.Parallel()

	clock := xtestutil.FixedClocker{}

	wantSid := typedef.SessionID(484493849343885677)
	wantUserID := typedef.UserID(484493849343820141)
	wantAuthTime := clock.Now()

	cacheRWMock := NewMockCacheReadWriter(gomock.NewController(t))
	cacheRWMock.EXPECT().ReadHash(gomock.Any(), gomock.Any(), userIDFieldName).Return(wantUserID.String(), nil).AnyTimes()
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

	sess, ok := ctx.Value(SessionKey{}).(*Session)
	if !ok {
		t.Fatal(xerr.TypeAssertionFailed)
	}
	if wantSid != sess.ID {
		t.Errorf("want = %d; got = %d", wantSid, sess.ID)
	}

	gotSess, ok := ctx.Value(SessionKey{}).(*Session)
	if !ok {
		t.Fatal(xerr.TypeAssertionFailed)
	}
	if wantUserID != gotSess.UserID {
		t.Errorf("want = %d; got = %d", wantUserID, gotSess.UserID)
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

	idGenMock := iface.NewMockIDGenerator(ctrl)
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
