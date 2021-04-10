// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Customer struct {
	_tab flatbuffers.Table
}

func GetRootAsCustomer(buf []byte, offset flatbuffers.UOffsetT) *Customer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Customer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Customer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Customer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Customer) FirstName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Customer) LastName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Customer) Age() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Customer) MutateAge(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *Customer) Balance() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Customer) MutateBalance(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

func (rcv *Customer) Debt() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Customer) MutateDebt(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

func (rcv *Customer) Preferences(obj *Preferences) *Preferences {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Preferences)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Customer) Friends(obj *Customer, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Customer) FriendsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Customer) Addresses(obj *AddressKV, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Customer) AddressesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func CustomerStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func CustomerAddFirstName(builder *flatbuffers.Builder, firstName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(firstName), 0)
}
func CustomerAddLastName(builder *flatbuffers.Builder, lastName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(lastName), 0)
}
func CustomerAddAge(builder *flatbuffers.Builder, age uint32) {
	builder.PrependUint32Slot(2, age, 0)
}
func CustomerAddBalance(builder *flatbuffers.Builder, balance float64) {
	builder.PrependFloat64Slot(3, balance, 0.0)
}
func CustomerAddDebt(builder *flatbuffers.Builder, debt float64) {
	builder.PrependFloat64Slot(4, debt, 0.0)
}
func CustomerAddPreferences(builder *flatbuffers.Builder, preferences flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(preferences), 0)
}
func CustomerAddFriends(builder *flatbuffers.Builder, friends flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(friends), 0)
}
func CustomerStartFriendsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CustomerAddAddresses(builder *flatbuffers.Builder, addresses flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(addresses), 0)
}
func CustomerStartAddressesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CustomerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
