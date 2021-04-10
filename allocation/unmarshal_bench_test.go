package allocation

import "testing"

// Benchmarking L1
func BenchmarkUnmarshalnNativeL1(b *testing.B) {
	bt := GetNativeLevel1()
	for i := 0; i < b.N; i++ {
		UnmarshalNative(bt)
	}
}

func BenchmarkUnmarshalProtoL1(b *testing.B) {
	bt := GetProtoLevel1()
	for i := 0; i < b.N; i++ {
		UnmarshalProto(bt)
	}
}

func BenchmarkUnmarshalFBSL1(b *testing.B) {
	bt := GetFBSLevel1()
	for i := 0; i < b.N; i++ {
		UnmarshalFBS(bt)
	}
}

// Benchmarking L2
func BenchmarkUnmarshalnNativeL2(b *testing.B) {
	bt := GetNativeLevel2()
	for i := 0; i < b.N; i++ {
		UnmarshalNative(bt)
	}
}

func BenchmarkUnmarshalProtoL2(b *testing.B) {
	bt := GetProtoLevel2()
	for i := 0; i < b.N; i++ {
		UnmarshalProto(bt)
	}
}

func BenchmarkUnmarshalFBSL2(b *testing.B) {
	bt := GetFBSLevel2()
	for i := 0; i < b.N; i++ {
		UnmarshalFBS(bt)
	}
}

// Benchmarking L3
func BenchmarkUnmarshalnNativeL3(b *testing.B) {
	bt := GetNativeLevel3()
	for i := 0; i < b.N; i++ {
		UnmarshalNative(bt)
	}
}

func BenchmarkUnmarshalProtoL3(b *testing.B) {
	bt := GetProtoLevel3()
	for i := 0; i < b.N; i++ {
		UnmarshalProto(bt)
	}
}

func BenchmarkUnmarshalFBSL3(b *testing.B) {
	bt := GetFBSLevel3()
	for i := 0; i < b.N; i++ {
		UnmarshalFBS(bt)
	}
}

// Gogo protobuf bench

// Benchmarking L1

func BenchmarkUnmarshalGogoProtoL1(b *testing.B) {
	bt := GetProtoLevel1()
	for i := 0; i < b.N; i++ {
		UnmarshalGogoProto(bt)
	}
}

// Benchmarking L2
func BenchmarkUnmarshalGogoProtoL2(b *testing.B) {
	bt := GetProtoLevel2()
	for i := 0; i < b.N; i++ {
		UnmarshalGogoProto(bt)
	}
}

// Benchmarking L3
func BenchmarkUnmarshalGogoProtoL3(b *testing.B) {
	bt := GetProtoLevel3()
	for i := 0; i < b.N; i++ {
		UnmarshalGogoProto(bt)
	}
}
