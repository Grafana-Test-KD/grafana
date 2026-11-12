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
	require.NotSame(t, http.DefaultTransport, b)
}

// requireSameTransportConfig compares every exported field of two transports.
// Function fields cannot be compared with ==, so they are compared by code
// pointer. DialContext is a method value bound to a per-transport dialer, but
// every instance shares the same code pointer, so a replaced dial function is
// detected.
func requireSameTransportConfig(t *testing.T, want, got *http.Transport) {
	t.Helper()

	wv := reflect.ValueOf(want).Elem()
	gv := reflect.ValueOf(got).Elem()
	for i := 0; i < wv.NumField(); i++ {
		field := wv.Type().Field(i)
		if !field.IsExported() {
			continue
		}
		wf, gf := wv.Field(i), gv.Field(i)
		if wf.Kind() == reflect.Func {
			require.Equal(t, wf.IsNil(), gf.IsNil(), "field %s", field.Name)
			require.Equal(t, wf.Pointer(), gf.Pointer(), "field %s", field.Name)
			continue
		}
		require.Equal(t, wf.Interface(), gf.Interface(), "field %s", field.Name)
	}
}

func TestNew(t *testing.T) {
	c := New()
	require.NotNil(t, c)
	tr, ok := c.Transport.(*http.Transport)
	require.True(t, ok)

	want := NewHTTPTransport().(*http.Transport)
	require.NotSame(t, want, tr)
	require.NotNil(t, tr.DialContext)
	requireSameTransportConfig(t, want, tr)
}
