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

func TestReadSession_ReadRedirectUri(t *testing.T) {
	t.Parallel()

	wantRedirectUri := "https://example.com/cb"

	sessReaderMock := NewMockSessionReader(gomock.NewController(t))
	sessReaderMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(wantRedirectUri, nil).AnyTimes()

	rs := NewReadSession(sessReaderMock)
	clientId := "CDcp9v3Nn4i70FqWig5AuohmorD6MG"
	authCode := "EYdxIU30xstnWZKxgA54RJMz1YUR0J"

	gotRedirectUri, err := rs.ReadRedirectUri(context.Background(), clientId, authCode)
	if err != nil {
		t.Fatal(err)
	}

	if wantRedirectUri != gotRedirectUri {
		t.Errorf("want = %s; got = %s", wantRedirectUri, gotRedirectUri)
	}
}

func TestRestoreSession_Restore(t *testing.T) {
	t.Parallel()

	wantSid := typedef.SessionID(484493849343885677)
	wantUserId := typedef.UserID(484493849343820141)

	sessReaderMock := NewMockSessionReader(gomock.NewController(t))
	sessReaderMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(wantUserId.String(), nil).AnyTimes()

	rs := NewRestoreSession(sessReaderMock)
	req := httptest.NewRequest(http.MethodGet, "https://example.com", nil)

	gotReq, err := rs.Restore(req, wantSid)
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

func TestWriteSession_WriteRedirectUriAssociation(t *testing.T) {
	t.Parallel()

	clientId := "oKOhD9oEdztNOiZ1n0WmsQ2NnylaYa"
	authCode := "LkQvn1FxZdqxFl9XYMOAjUqIlXI9IC"
	uri := "https://example.com/cb"

	ctrl := gomock.NewController(t)
	sessWriterMock := NewMockSessionWriter(ctrl)
	sessWriterMock.EXPECT().Write(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
	idGenMock := NewMockIdGenerator(ctrl)
	idGenMock.EXPECT().NextID().Return(uint64(0), nil).AnyTimes()
	ctxReaderMock := iface.NewMockContextReader(ctrl)
	ctxReaderMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(typedef.SessionID(0)).AnyTimes()

	ws := NewWriteSession(sessWriterMock, ctxReaderMock, idGenMock)

	if err := ws.WriteRedirectUriAssociation(context.Background(), uri, clientId, authCode); err != nil {
		t.Error(err)
	}
}

func TestWriteSession_WriteUserId(t *testing.T) {
	t.Parallel()

	sid := typedef.SessionID(484493849344016749)
	uid := typedef.UserID(484493849344082285)

	wantSid := sid

	ctrl := gomock.NewController(t)
	sessWriterMock := NewMockSessionWriter(ctrl)
	sessWriterMock.EXPECT().Write(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
	idGenMock := NewMockIdGenerator(ctrl)
	idGenMock.EXPECT().NextID().Return(uint64(wantSid), nil).AnyTimes()
	ctxReaderMock := iface.NewMockContextReader(ctrl)
	ctxReaderMock.EXPECT().Read(gomock.Any(), gomock.Any()).Return(wantSid).AnyTimes()

	ws := NewWriteSession(sessWriterMock, ctxReaderMock, idGenMock)

	gotSid, err := ws.WriteUserId(context.Background(), uid)
	if err != nil {
		t.Fatal(err)
	}
	if wantSid != gotSid {
		t.Errorf("want = %d; got =%d", wantSid, gotSid)
	}
}
