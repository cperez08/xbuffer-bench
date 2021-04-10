// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Location struct {
	_tab flatbuffers.Table
}

func GetRootAsLocation(buf []byte, offset flatbuffers.UOffsetT) *Location {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Location{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Location) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Location) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Location) Address() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Location) Latitude() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Location) MutateLatitude(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *Location) Longitude() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Location) MutateLongitude(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func (rcv *Location) IsAtive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Location) MutateIsAtive(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func LocationStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func LocationAddAddress(builder *flatbuffers.Builder, address flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(address), 0)
}
func LocationAddLatitude(builder *flatbuffers.Builder, latitude float64) {
	builder.PrependFloat64Slot(1, latitude, 0.0)
}
func LocationAddLongitude(builder *flatbuffers.Builder, longitude float64) {
	builder.PrependFloat64Slot(2, longitude, 0.0)
}
func LocationAddIsAtive(builder *flatbuffers.Builder, isAtive bool) {
	builder.PrependBoolSlot(3, isAtive, false)
}
func LocationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
