package link

import (
	"io"
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func TestParse(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    Links
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodeToLinks(t *testing.T) {
	type args struct {
		root *html.Node
	}
	tests := []struct {
		name    string
		args    args
		want    Links
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nodeToLinks(tt.args.root)
			if (err != nil) != tt.wantErr {
				t.Errorf("nodeToLinks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodeToLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_collectATags(t *testing.T) {
	type args struct {
		root *html.Node
	}
	tests := []struct {
		name string
		args args
		want []*html.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := collectATags(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collectATags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTagText(t *testing.T) {
	type args struct {
		root *html.Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseTagText(tt.args.root); got != tt.want {
				t.Errorf("parseTagText() = %v, want %v", got, tt.want)
			}
		})
	}
}
