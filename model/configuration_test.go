package model

import "testing"

func TestFor(t *testing.T) {
	conf := Configuration{}
	value := conf.For("storage.folders")

	if value != "../resource/md/;../resource/html/" {
		t.Error("Wrong value for property 'storage.folders'.")
	}
}
