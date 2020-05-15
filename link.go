package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Link is representation of link and text structure
type Link struct {
	Href string `json:"href,omitempty"`
	Text string `json:"text,omitempty"`
}

// Links is a collection of Link structs
type Links []Link

// Parse creates a collection of "Link" (type "Links")
// base on the contents of HTML file
func Parse(r io.Reader) (Links, error) {

	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	return nodeToLinks(root)
}

func nodeToLinks(root *html.Node) (Links, error) {

	aTags := collectATags(root)

	links := make(Links, len(aTags))
	for i, tag := range aTags {
		link := &links[i]

		// has HREF
		for _, atr := range tag.Attr {
			if atr.Key == "href" {
				link.Href = atr.Val
			}
		}

		// has text to it
		link.Text = parseTagText(tag)
	}

	return links, nil
}

func collectATags(root *html.Node) []*html.Node {
	// base case
	if root == nil {
		return nil
	}

	// root node is an <a></a> node
	var aTagNodes []*html.Node
	if root.DataAtom == atom.A {
		aTagNodes = append(aTagNodes, root)
		return aTagNodes
	}

	// root node not an <a></a> but might have a child that is
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		childLinks := collectATags(c)
		if childLinks != nil {
			aTagNodes = append(aTagNodes, childLinks...)
		}
	}

	return aTagNodes
}

func parseTagText(root *html.Node) string {
	if root == nil {
		return ""
	}

	var text string
	if root.Type == html.TextNode {
		text += strings.TrimSpace(root.Data)
	}

	for c := root.FirstChild; c != nil; c = c.NextSibling {
		text += " " + parseTagText(c)
	}

	return text
}
