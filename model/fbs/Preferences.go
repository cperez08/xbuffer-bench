// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Preferences struct {
	_tab flatbuffers.Table
}

func GetRootAsPreferences(buf []byte, offset flatbuffers.UOffsetT) *Preferences {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Preferences{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Preferences) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Preferences) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Preferences) DarkMode() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Preferences) MutateDarkMode(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *Preferences) Language() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Preferences) Notification(obj *Notification) *Notification {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Notification)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func PreferencesStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func PreferencesAddDarkMode(builder *flatbuffers.Builder, darkMode bool) {
	builder.PrependBoolSlot(0, darkMode, false)
}
func PreferencesAddLanguage(builder *flatbuffers.Builder, language flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(language), 0)
}
func PreferencesAddNotification(builder *flatbuffers.Builder, notification flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(notification), 0)
}
func PreferencesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
