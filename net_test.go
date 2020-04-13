package equalfold

import (
	"net/url"
	"testing"
)

func TestUrlParse(t *testing.T) {
	pageurl := "file:///www.halega.com./news?id=1234#big"
	uri, err := url.Parse(pageurl)
	if err != nil {
		t.Error(err)
	}
	if uri.Host != "www.halega.com." {
		t.Errorf("want %s, got %s", "www.halega.com.", uri.Host)
	}
	if uri.Scheme != "https://" {
		t.Errorf("want %s, got %s", "https://", uri.Scheme)
	}
}

func TestUrlParse_File(t *testing.T) {
	rawurl := "file:///home/havoron/Downloads/halega.kdbx"
	uri, err := url.Parse(rawurl)
	if err != nil {
		t.Error(err)
	}
	if uri.Scheme != "file" {
		t.Errorf("want %s, got %s", "file", uri.Scheme)
	}
	t.Log(uri)
}
