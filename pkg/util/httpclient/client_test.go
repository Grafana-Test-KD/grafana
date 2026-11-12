package httpclient

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewHTTPTransport(t *testing.T) {
	rt := NewHTTPTransport()
	tr, ok := rt.(*http.Transport)
	require.True(t, ok)

	require.Equal(t, reflect.ValueOf(http.ProxyFromEnvironment).Pointer(), reflect.ValueOf(tr.Proxy).Pointer())
	require.NotNil(t, tr.DialContext)
	require.True(t, tr.ForceAttemptHTTP2)
	require.Equal(t, 100, tr.MaxIdleConns)
	require.Equal(t, 90*time.Second, tr.IdleConnTimeout)
	require.Equal(t, 10*time.Second, tr.TLSHandshakeTimeout)
	require.Equal(t, time.Second, tr.ExpectContinueTimeout)
}

func TestNewHTTPTransportReturnsDistinctInstances(t *testing.T) {
	a := NewHTTPTransport()
	b := NewHTTPTransport()
	require.NotSame(t, a, b)
	require.NotSame(t, http.DefaultTransport, a)
}

func TestNew(t *testing.T) {
	c := New()
	require.NotNil(t, c)
	tr, ok := c.Transport.(*http.Transport)
	require.True(t, ok)

	want := NewHTTPTransport().(*http.Transport)
	require.NotSame(t, want, tr)
	require.Equal(t, reflect.ValueOf(want.Proxy).Pointer(), reflect.ValueOf(tr.Proxy).Pointer())
	require.NotNil(t, tr.DialContext)
	require.Equal(t, want.ForceAttemptHTTP2, tr.ForceAttemptHTTP2)
	require.Equal(t, want.MaxIdleConns, tr.MaxIdleConns)
	require.Equal(t, want.IdleConnTimeout, tr.IdleConnTimeout)
	require.Equal(t, want.TLSHandshakeTimeout, tr.TLSHandshakeTimeout)
	require.Equal(t, want.ExpectContinueTimeout, tr.ExpectContinueTimeout)
}
