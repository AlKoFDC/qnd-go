package markdown

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestShouldParseMarkdown(t *testing.T) {
	for _, testIO := range []struct {
		in     string
		expect string
	}{
		{in: "hello, world", expect: `<div><p>hello, world</p></div>`},
		{
			in:     "line1\nline2",
			expect: "<div><p>line1</p><p>line2</p></div>",
		},
		{
			in:     "paragraph1\n\nparagraph2",
			expect: "<div><p>paragraph1</p><p>paragraph2</p></div>",
		},
		{
			in:     "x",
			expect: "<div><p>x</p></div>",
		},
		{
			in:     "# headline",
			expect: "<div><h1>headline</h1></div>",
		},
		{
			in:     "# headline\nline1\nline2",
			expect: "<div><h1>headline</h1><p>line1</p><p>line2</p></div>",
		},
		{
			in:     "## headline\nline1\nline2",
			expect: "<div><h2>headline</h2><p>line1</p><p>line2</p></div>",
		},
		{
			in:     "# headline1\nline\n## headline2",
			expect: "<div><h1>headline1</h1><p>line</p><h2>headline2</h2></div>",
		},
		{
			in:     "*bold*",
			expect: "<div><p><strong>bold</strong></p></div>",
		},
		{
			in:     "*",
			expect: "<div><p>*</p></div>",
		},
		{
			in:     "bold*",
			expect: "<div><p>bold*</p></div>",
		},
		{
			in:     "**bold*",
			expect: "<div><p><strong></strong>bold*</p></div>",
		},
		{
			in:     "*bold**",
			expect: "<div><p><strong>bold</strong>*</p></div>",
		},
		{
			in:     "*bold*not bold*strong*",
			expect: "<div><p><strong>bold</strong>not bold<strong>strong</strong></p></div>",
		},
	} {
		if got := ParseToHTML(testIO.in); got != testIO.expect {
			t.Errorf("For '%s' expected '%s', but got '%s'.", testIO.in, testIO.expect, got)
		}
	}
}

func TestShouldServeParser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handleParse))
	defer ts.Close()

	for _, testIO := range []struct {
		query  string
		expect string
	}{
		{query: "# hello, world", expect: "<div><h1>hello, world</h1></div>"},
		{query: "# hello, *world*", expect: "<div><h1>hello, <strong>world</strong></h1></div>"},
		{query: "## hello, *world*2", expect: "<div><h2>hello, <strong>world</strong>2</h2></div>"},
	} {
		res, err := http.Get(ts.URL + "?md=" + url.QueryEscape(testIO.query))
		if err != nil {
			t.Fatal(err)
		}
		result, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
		if resString := string(result); resString != testIO.expect {
			t.Errorf("For %s expect %s but got %s.", testIO.query, testIO.expect, resString)
		}
	}
}
