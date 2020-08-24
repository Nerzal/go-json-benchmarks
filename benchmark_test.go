package bench

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/francoispqt/gojay"
	jsoniter "github.com/json-iterator/go"
	pkgJson "github.com/pkg/json"
	"github.com/pquerna/ffjson/ffjson"
)

var canadaBytes []byte

func init() {
	canadaBytes, _ = ioutil.ReadFile("testfiles/canada.json")
}

func Benchmark_Std_Canada(b *testing.B) {
	b.ReportAllocs()

	var canada Canada

	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(canadaBytes, &canada)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_json_iter_Canada(b *testing.B) {
	b.ReportAllocs()

	var canada Canada
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	for i := 0; i < b.N; i++ {

		err := json.Unmarshal(canadaBytes, &canada)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_ffjson_Canada(b *testing.B) {
	b.ReportAllocs()

	var canada Canada

	for i := 0; i < b.N; i++ {
		err := ffjson.Unmarshal(canadaBytes, &canada)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_pkg_json_Canada(b *testing.B) {
	b.ReportAllocs()

	var canada Canada
	decoder := pkgJson.NewDecoder(bytes.NewReader(canadaBytes))

	for i := 0; i < b.N; i++ {

		err := decoder.Decode(&canada)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_go_jay_Canada(b *testing.B) {
	b.ReportAllocs()

	var canada Canada

	for i := 0; i < b.N; i++ {

		err := gojay.Unmarshal(canadaBytes, &canada)
		if err != nil {
			b.Error(err)
		}
	}
}
