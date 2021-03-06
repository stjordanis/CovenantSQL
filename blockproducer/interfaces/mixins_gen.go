package interfaces

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	hsp "github.com/CovenantSQL/HashStablePack/marshalhash"
)

// MarshalHash marshals for hash
func (z *TransactionTypeMixin) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82)
	o = hsp.AppendTime(o, z.Timestamp)
	if oTemp, err := z.TxType.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TransactionTypeMixin) Msgsize() (s int) {
	s = 1 + 10 + hsp.TimeSize + 7 + z.TxType.Msgsize()
	return
}
