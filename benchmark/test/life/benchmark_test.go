package wasmer_test

import (
	life "github.com/perlin-network/life/exec"
	"io/ioutil"
	"testing"
	"path"
	"runtime"
)

const N = 100000

func GetBytes(wasmFile string) []byte {
	_, filename, _, _ := runtime.Caller(0)
	path := path.Join(path.Dir(filename), "wasmfile", wasmFile)
	bytes, _ := ioutil.ReadFile(path)
	return bytes
}

func benchmarkLife(b *testing.B, wasmfile string, exportName string, exportValues ...int64) {
	// wasmBytes, _ := ioutil.ReadFile("../wasmer/wasmfile/benchmark.wasm")
	wasmBytes := GetBytes(wasmfile)
	for i := 0; i < b.N; i++ {
		vm, _ := life.NewVirtualMachine(wasmBytes, life.VMConfig{}, &life.NopResolver{}, nil)
		entryID, _ := vm.GetFunctionExport(exportName)
		result, _ := vm.Run(entryID, exportValues...)
		_ = result
	}
}

func Benchmark_Wasmer_sub(b *testing.B) {
	benchmarkLife(b, "difference.wasm", "sub_main")
}

func Benchmark_Wasmer_hello(b *testing.B) {
	benchmarkLife(b, "string.wasm", "hello_main", 0)
}

func Benchmark_Wasmer_math(b *testing.B) {
	benchmarkLife(b, "math.wasm", "eval_main")
}

func Benchmark_Wasmer_fib(b *testing.B) {
	benchmarkLife(b, "fibonacci.wasm", "fib_main")
}