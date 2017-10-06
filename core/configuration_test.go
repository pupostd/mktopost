package core

import (
	"testing"
)

func TestConfiguration(t *testing.T) {
	conf := Configuration{}
	conf.Load()

	if len(conf.Properties) != 3 {
		t.Error("Not enough properties on property file! Shoud be 3.")
	}

	var count int
	for _, v := range conf.Properties {

		if "to.parse" == v.Name {
			if "../resource/to_parse/" == v.Value {
				count++
			}
		} else if "to.publish" == v.Name {
			if "../resource/to_publish/" == v.Value {
				count++
			}
		} else if "to.publish.hidden" == v.Name {
			if "../resource/to_publish_hidden/" == v.Value {
				count++
			}
		}
	}

	if count != 3 {
		t.Error("Wrong properties defined on property file!")
	}
}
