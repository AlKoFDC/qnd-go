package markdown

import (
	"fmt"
	"strings"
	"net/http"
	"log"
)

func ParseToHTML(md string) string {
	const template = "<div>%s</div>"
	var html []string
	for _, line := range strings.Split(md, "\n") {
		switch {
		case len(strings.Trim(line, " \t\r")) < 1:
			// Skip empty lines.
		case len(line) > 1 && line[0:2] == "# ":
			// Heading 1
			html = append(html, "<h1>"+bold(line[2:])+"</h1>")
		case len(line) > 2 && line[0:3] == "## ":
			// Heading 2
			html = append(html, "<h2>"+bold(line[3:])+"</h2>")
		default:
			// Paragraph
			html = append(html, "<p>"+bold(line)+"</p>")
		}
	}
	return fmt.Sprintf(template, strings.Join(html, ""))
}

func bold(in string) string {
	if strings.Count(in, "*") < 2 {
		return in
	}
	in2 := strings.Replace(in, "*", "<strong>", 1)
	return bold(strings.Replace(in2, "*", "</strong>", 1))
}

func handleParse (w http.ResponseWriter, r *http.Request){
	md := r.URL.Query().Get("md")
	w.Write([]byte(ParseToHTML(md)))
}

func StartServer() {
	http.HandleFunc("/parse", handleParse)
	log.Fatal(http.ListenAndServe(":8080", nil))
}