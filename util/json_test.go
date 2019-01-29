package util

import (
	"fmt"
	//"reflect"
	"testing"
)

type T struct {
	Name  string `json:name`
	Order string `json:Order`
}

type TTestJSONMarshal struct {
	Data1 string `json:"info"`
}

func TestJSONUnmarshal(t *testing.T) {
	type args struct {
		blob []byte
		v    interface{}
	}

	var jsonBlob = []byte(`[
    {"Name": "Platypus", "Order": "Monotremata"},
    {"Name": "Quoll",    "Order": "Dasyuromorphia"}
		]`)

	var foo [2]T

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{blob: jsonBlob, v: &foo}, false},
		{"test2", args{blob: jsonBlob, v: foo}, true},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JSONUnmarshal(tt.args.blob, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("JSONUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//测试json含有多余数据 是否返回错误
func TestJSONUnmarshal2(t *testing.T) {
	type args struct {
		blob []byte
		v    interface{}
	}

	type T2 struct {
		Name string `json:"name"`
		Time string `json:"time"`
	}

	var jsonBlob = []byte(`
	{"name": "Platypus", 
	 "order": "Monotremata",
	 "time": "1234",
	 }   
		`)

	var foo T2

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{blob: jsonBlob, v: &foo}, false},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JSONUnmarshal(tt.args.blob, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("JSONUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Printf("%v", tt.args.v)
		})
	}
}

func TestJSONMarshal(t *testing.T) {
	type args struct {
		data      interface{}
		hasIndent bool
	}
	tests := []struct {
		name        string
		args        args
		wantContent []byte
		wantErr     bool
	}{
		{"case1",
			args{TTestJSONMarshal{"dGhpcyBpcyBhIGV4YW1wbGU="}, true},
			nil,
			false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContent, err := JSONMarshal(tt.args.data, tt.args.hasIndent)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONMarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(string(gotContent))
			// if !reflect.DeepEqual(gotContent, tt.wantContent) {
			// 	t.Errorf("JSONMarshal() = %v, want %v", gotContent, tt.wantContent)
			// 	fmt.Println(string(gotContent))
			// }
		})
	}
}
