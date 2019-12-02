package link

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type Link struct {
	Href string
	Text string
}

func processAnchor(n *html.Node) Link {
	link := Link{}
	for _, attribute := range n.Attr {
		if attribute.Key == "href" {
			link.Href = attribute.Val
		}
	}
	return link
}

func processNode(n *html.Node) []Link {
	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		link := processAnchor(n)
		links = append(links, link)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childLinks := processNode(c)
		links = append(links, childLinks...)
	}

	return links
}

func ExtractLinks(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("error while parsing html: %w", err)
	}

	links := processNode(doc)

	return links, nil
}
