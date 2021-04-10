package allocation

import (
	"testing"
)

func TestGetObjectsLevel1(t *testing.T) {
	native := GetNativeLevel1()
	proto := GetProtoLevel1()
	fbs := GetFBSLevel1()
	gogo := GetGogoProtoLevel1()

	t.Log("Native L1 Len: ", len(native))
	t.Log("Proto L1 Len: ", len(proto))
	t.Log("Fbs L1 Len: ", len(fbs))
	t.Log("Gogo L1 Len: ", len(gogo))
}

func TestGetObjectsLevel2(t *testing.T) {
	native := GetNativeLevel2()
	proto := GetProtoLevel2()
	fbs := GetFBSLevel2()
	gogo := GetGogoProtoLevel2()

	t.Log("Native L2 Len: ", len(native))
	t.Log("Proto L2 Len: ", len(proto))
	t.Log("Fbs L2 Len: ", len(fbs))
	t.Log("Gogo L2 Len: ", len(gogo))
}

func TestGetObjectsLevel3(t *testing.T) {
	native := GetNativeLevel3()
	proto := GetProtoLevel3()
	fbs := GetFBSLevel3()
	gogo := GetGogoProtoLevel3()

	t.Log("Native L3 Len: ", len(native))
	t.Log("Proto L3 Len: ", len(proto))
	t.Log("Fbs L3 Len: ", len(fbs))
	t.Log("Gogo L3 Len: ", len(gogo))
}
