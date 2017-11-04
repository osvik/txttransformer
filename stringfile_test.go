package txttransformer

import (
	"testing"
)

func Test_ReplaceAllText(t *testing.T) {
	type args struct {
		fileText   string
		searchText string
		newText    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "T1", want: "go my go", args: args{
			fileText:   "fi my fi",
			searchText: "fi",
			newText:    "go"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceAllText(tt.args.fileText, tt.args.searchText, tt.args.newText); got != tt.want {
				t.Errorf("replaceAllText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ReplaceMatchedText(t *testing.T) {
	type args struct {
		fileText   string
		regularExp string
		newText    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "T2", want: "fi my fi go", args: args{
			fileText:   "fi my fi fi",
			regularExp: "fi$",
			newText:    "go"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceMatchedText(tt.args.fileText, tt.args.regularExp, tt.args.newText); got != tt.want {
				t.Errorf("replaceMatchedText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ReplaceMatchedTextFunction(t *testing.T) {
	type args struct {
		fileText   string
		regularExp string
		f          func(a string) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "R1", want: "Do is the great thing we want", args: args{
			regularExp: `do\!$`,
			fileText:   "Do is the great thing we do!",
			f: func(a string) string {
				return "want"
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceMatchedTextFunction(tt.args.fileText, tt.args.regularExp, tt.args.f); got != tt.want {
				t.Errorf("replaceMatchedTextFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
