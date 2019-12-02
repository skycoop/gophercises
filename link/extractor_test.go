package link

import (
	"os"
	"reflect"
	"testing"
)

type testCase struct {
	name     string
	htmlFile string
	want     []Link
	wantErr  bool
}

var tests = []testCase{
	{
		name:     "ex1",
		htmlFile: "link/test_data/ex1.html",
		want: []Link{
			{
				Href: "/other-page",
				Text: "A link to another page",
			},
		},
		wantErr: false,
	},
	{
		name:     "ex2",
		htmlFile: "link/test_data/ex2.html",
		want: []Link{
			{
				Href: "https://www.twitter.com/joncalhoun",
				Text: "Check me out on twitter",
			},
			{
				Href: "https://github.com/gophercises",
				Text: "Gophercises is on Github!",
			},
		},
		wantErr: false,
	},
	{
		name:     "ex3",
		htmlFile: "link/test_data/ex3.html",
		want: []Link{
			{
				Href: "#",
				Text: "Login",
			},
			{
				Href: "/lost",
				Text: "Lost? Need help?",
			},
			{
				Href: "https://twitter.com/marcusolsson",
				Text: "@marcusolsson",
			},
		},
		wantErr: false,
	},
	{
		name:     "ex4",
		htmlFile: "link/test_data/ex4.html",
		want: []Link{
			{
				Href: "/dog-cat",
				Text: "dog cat",
			},
		},
		wantErr: false,
	},
}

func TestExtractLinks(t *testing.T) {

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r, err := os.Open(test.htmlFile)
			if err != nil {
				t.Fatalf("Failed to load data file %s", test.htmlFile)
			}

			got, err := ExtractLinks(r)
			if (err != nil) != test.wantErr {
				t.Errorf("ExtractLinks() error = %+v, wantErr %+v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("ExtractLinks() got = %+v, want %+v", got, test.want)
			}
		})
	}
}
