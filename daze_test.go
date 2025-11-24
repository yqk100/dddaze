package daze

import (
	"bytes"
	"context"
	"os/exec"
	"testing"

	"github.com/libraries/daze/lib/doa"
)

const (
	DazeServerListenOn = "127.0.0.1:28080"
	CurlDest           = "https://www.zhihu.com"
	HostLookup         = "google.com"
)

func TestLocaleHTTP(t *testing.T) {
	locale := NewLocale(DazeServerListenOn, &Direct{})
	defer locale.Close()
	locale.Run()

	cmd := exec.Command("curl", "-x", "http://"+DazeServerListenOn, CurlDest)
	out := doa.Try(cmd.Output())
	if !bytes.Contains(out, []byte("zhihu")) {
		t.FailNow()
	}
}

func TestLocaleSocks4(t *testing.T) {
	locale := NewLocale(DazeServerListenOn, &Direct{})
	defer locale.Close()
	locale.Run()

	cmd := exec.Command("curl", "-x", "socks4://"+DazeServerListenOn, CurlDest)
	out := doa.Try(cmd.Output())
	if !bytes.Contains(out, []byte("zhihu")) {
		t.FailNow()
	}
}

func TestLocaleSocks4a(t *testing.T) {
	locale := NewLocale(DazeServerListenOn, &Direct{})
	defer locale.Close()
	locale.Run()

	cmd := exec.Command("curl", "-x", "socks4a://"+DazeServerListenOn, CurlDest)
	out := doa.Try(cmd.Output())
	if !bytes.Contains(out, []byte("zhihu")) {
		t.FailNow()
	}
}

func TestLocaleSocks5(t *testing.T) {
	locale := NewLocale(DazeServerListenOn, &Direct{})
	defer locale.Close()
	locale.Run()

	cmd := exec.Command("curl", "-x", "socks5://"+DazeServerListenOn, CurlDest)
	out := doa.Try(cmd.Output())
	if !bytes.Contains(out, []byte("zhihu")) {
		t.FailNow()
	}
}

func TestResolverDns(t *testing.T) {
	dns := ResolverDns(ResolverPublic.Cloudflare.Dns)
	_, err := dns.LookupHost(context.Background(), HostLookup)
	if err != nil {
		t.FailNow()
	}
}

func TestResolverDot(t *testing.T) {
	dot := ResolverDot(ResolverPublic.Cloudflare.Dot)
	_, err := dot.LookupHost(context.Background(), HostLookup)
	if err != nil {
		t.FailNow()
	}
}

func TestResolverDoh(t *testing.T) {
	doh := ResolverDoh(ResolverPublic.Cloudflare.Doh)
	_, err := doh.LookupHost(context.Background(), HostLookup)
	if err != nil {
		t.FailNow()
	}
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
		_, err := dns.LookupHost(context.Background(), HostLookup)
		if err != nil {
			t.FailNow()
		}
	}
}
