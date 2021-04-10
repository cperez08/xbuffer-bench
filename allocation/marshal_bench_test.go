package allocation

import "testing"

// Benchmarking L1
func BenchmarkMarshalNativeL1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetNativeLevel1()
	}
}

func BenchmarkMarshalProtoL1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetProtoLevel1()
	}
}

func BenchmarkMarshalFBSL1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFBSLevel1()
	}
}

// Benchamarking L2
func BenchmarkMarshalNativeL2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetNativeLevel2()
	}
}

func BenchmarkMarshalProtoL2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetProtoLevel2()
	}
}

func BenchmarkMarshalFBSL2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFBSLevel2()
	}
}

// Benchamarking L3
func BenchmarkMarshalNativeL3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetNativeLevel3()
	}
}

func BenchmarkMarshalProtoL3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetProtoLevel3()
	}
}

func BenchmarkAllocationFBSL3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFBSLevel3()
	}
}

//Gogo Protobuf bench

// Benchmarking L1
func BenchmarkMarshalGogoProtoL1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetGogoProtoLevel1()
	}
}

// Benchamarking L2
func BenchmarkMarshalGogoProtoL2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetGogoProtoLevel2()
	}
}

// Benchamarking L3
func BenchmarkMarshalGogoProtoL3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetGogoProtoLevel3()
	}
}
