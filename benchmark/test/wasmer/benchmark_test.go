package wasmer_test

import (
	"path"
	"runtime"
	"github.com/wasmerio/go-ext-wasm/wasmer"
	"testing"
)

const N = 100000

func GetBytes(wasmFile string) []byte {
	_, filename, _, _ := runtime.Caller(0)
	path := path.Join(path.Dir(filename), "wasmfile", wasmFile)
	bytes, _ := wasmer.ReadBytes(path)
	return bytes
}

func benchmarkWasmer(b *testing.B, wasmfile string, exportName string, exportValues ...interface{}) {
	// wasmBytes, _ := ioutil.ReadFile("../wasmer/wasmfile/difference.wasm")
	wasmBytes := GetBytes(wasmfile)
	for i := 0; i < b.N; i++ {
		instance, _ := wasmer.NewInstance(wasmBytes)
		result, _ := instance.Exports[exportName]
		_ = result
	}
}

func Benchmark_Wasmer_sub(b *testing.B) {
	benchmarkWasmer(b, "difference.wasm", "sub_main", N)
}

func Benchmark_Wasmer_hello(b *testing.B) {
	benchmarkWasmer(b, "string.wasm", "hello_main")
}

func Benchmark_Wasmer_math(b *testing.B) {
	benchmarkWasmer(b, "math.wasm", "eval_main", N)
}

func Benchmark_Wasmer_fib(b *testing.B) {
	benchmarkWasmer(b, "fibonacci.wasm", "fib_main", N)
}

func Benchmark_Wasmer_Cstring(b *testing.B){
	benchmarkWasmer(b, "CString.wasm", "greet")
}