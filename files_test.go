package txttransformer

import (
	"os"
	"reflect"
	"testing"
)

func Test_FindFiles(t *testing.T) {
	type args struct {
		searchDir string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "T4",
			want: []string{"testdata/foo/bar/file.txt"},
			args: args{
				searchDir: "testdata/foo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFiles(tt.args.searchDir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_FilterPaths(t *testing.T) {
	type args struct {
		f   []string
		ext string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "T5",
			want: []string{"index.html"},
			args: args{
				ext: "html",
				f: []string{
					"index.txt",
					"index.html",
					"test.jpg",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterPaths(tt.args.f, tt.args.ext); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CreatePathFolder(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "X1", args: args{path: "testdata/foo/bar/file.txt"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreatePathFolder(tt.args.path)
			if _, err := os.Stat("testdata/foo/bar/"); os.IsNotExist(err) {
				t.Errorf("Could not find parent folder")
			}
		})
	}
}

func Test_ReadTextFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "w1", want: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?`, args: args{path: "testdata/file.txt"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := os.Stat(tt.args.path); os.IsNotExist(err) {
				t.Errorf("Could not find the file")
			}
			if got := ReadTextFile(tt.args.path); got != tt.want {
				t.Errorf("readTextFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_WriteTextFile(t *testing.T) {
	type args struct {
		path    string
		content string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "B2", args: args{
			path:    "testdata/file2.txt",
			content: "This is a file created by a test",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteTextFile(tt.args.path, tt.args.content)
			if _, err := os.Stat(tt.args.path); os.IsNotExist(err) {
				t.Errorf("Could not find the file")
			}
			if nn := ReadTextFile(tt.args.path); nn != tt.args.content {
				t.Errorf("The file is different than what I expected")
			}
		})
	}
}

func TestReadListPaths(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "y1",
			want: []string{"1.txt", "baz/2.txt", "3.txt"},
			args: args{filename: "testdata/paths.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadListPaths(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadListPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
