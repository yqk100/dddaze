package daze

import (
	"bytes"
	"context"
	"os/exec"
	"testing"

	"github.com/libraries/daze/lib/doa"
)

const (
	CurlDest           = "https://www.zhihu.com"
	DazeLocaleListenOn = "127.0.0.1:28080"
	HostLookup         = "google.com"
)

func TestLocaleHttp(t *testing.T) {
	locale := NewLocale(DazeLocaleListenOn, &Direct{})
	defer locale.Close()
	locale.Run()

	cmd := exec.Command("curl", "-x", "http://"+DazeLocaleListenOn, CurlDest)
	out := doa.Try(cmd.Output())
	doa.Doa(bytes.Contains(out, []byte("zhihu")))
}

func TestLocaleSocks4(t *testing.T) {
	locale := NewLocale(DazeLocaleListenOn, &Direct{})
	defer locale.Close()
	locale.Run()

	cmd := exec.Command("curl", "-x", "socks4://"+DazeLocaleListenOn, CurlDest)
	out := doa.Try(cmd.Output())
	doa.Doa(bytes.Contains(out, []byte("zhihu")))
}

func TestLocaleSocks4a(t *testing.T) {
	locale := NewLocale(DazeLocaleListenOn, &Direct{})
	defer locale.Close()
	locale.Run()

	cmd := exec.Command("curl", "-x", "socks4a://"+DazeLocaleListenOn, CurlDest)
	out := doa.Try(cmd.Output())
	doa.Doa(bytes.Contains(out, []byte("zhihu")))
}

func TestLocaleSocks5(t *testing.T) {
	locale := NewLocale(DazeLocaleListenOn, &Direct{})
	defer locale.Close()
	locale.Run()

	cmd := exec.Command("curl", "-x", "socks5://"+DazeLocaleListenOn, CurlDest)
	out := doa.Try(cmd.Output())
	doa.Doa(bytes.Contains(out, []byte("zhihu")))
}

func TestResolverAll(t *testing.T) {
	for _, url := range []string{
		ResolverPublic.Alidns.Dns,
		ResolverPublic.Alidns.Dot,
		ResolverPublic.Alidns.Doh,
		ResolverPublic.Cloudflare.Dns,
		ResolverPublic.Cloudflare.Dot,
		ResolverPublic.Cloudflare.Doh,
		ResolverPublic.Google.Dns,
		ResolverPublic.Google.Dot,
		ResolverPublic.Google.Doh,
		ResolverPublic.Tencent.Dns,
		ResolverPublic.Tencent.Dot,
		ResolverPublic.Tencent.Doh,
	} {
		dns := ResolverAny(url)
		err := doa.Err(dns.LookupHost(context.Background(), HostLookup))
		doa.Nil(err)
	}
}
