package czar

import (
	"testing"

	"github.com/libraries/daze/lib/doa"
)

func TestProtocolCzarSip(t *testing.T) {
	sid := NewSip()
	for i := range Conf.StreamPool {
		doa.Doa(doa.Try(sid.Get()) == uint16(i))
	}
	doa.Doa(doa.Err(sid.Get()) != nil)
	doa.Doa(sid.Pop() == Conf.StreamPool)
	sid.Put(65)
	doa.Doa(sid.Pop() == Conf.StreamPool-1)
	sid.Put(15)
	doa.Doa(sid.Pop() == Conf.StreamPool-2)
	doa.Doa(doa.Try(sid.Get()) == 15)
	doa.Doa(sid.Pop() == Conf.StreamPool-1)
	doa.Doa(doa.Try(sid.Get()) == 65)
	doa.Doa(sid.Pop() == Conf.StreamPool)
}
