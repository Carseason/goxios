package goxios

import (
	"bytes"
	"errors"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

type Node struct {
	node *html.Node
	err  error
}

func newNode(body []byte) *Node {
	node, err := html.Parse(bytes.NewBuffer(body))
	return &Node{
		node: node,
		err:  err,
	}
}
func (nd *Node) Node() (*html.Node, error) {
	return nd.node, nd.err
}

// css select
func (nd *Node) Query(rule string) (*html.Node, error) {
	n, err := nd.Node()
	if err != nil {
		return nil, err
	}
	sel, err := cascadia.Parse(rule)
	if err != nil {
		return nil, err
	}
	node := cascadia.Query(n, sel)
	if node == nil {
		return nil, errors.New("not found rule node")
	}
	return node, nil
}
func (nd *Node) QueryAll(rule string) ([]*html.Node, error) {
	n, err := nd.Node()
	if err != nil {
		return nil, err
	}
	sel, err := cascadia.Parse(rule)
	if err != nil {
		return nil, err
	}
	nodes := cascadia.QueryAll(n, sel)
	if nodes == nil {
		return nil, errors.New("not found rule nodes")
	}
	return nodes, nil
}
func (nd *Node) QueryRender(rule string) (string, error) {
	node, err := nd.Query(rule)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err := html.Render(&b, node); err != nil {
		return "", err
	}
	return b.String(), nil
}
func (nd *Node) QueryAttr(rule, tag string) (string, error) {
	node, err := nd.Query(rule)
	if err != nil {
		return "", err
	}
	if node != nil {
		for _, v := range node.Attr {
			if v.Key == tag {
				return v.Val, nil
			}
		}
	}
	return "", nil
}
func (nd *Node) QueryData(rule string) (string, error) {
	node, err := nd.Query(rule)
	if err != nil {
		return "", err
	}
	return node.Data, nil
}
func (nd *Node) QueryText(rule string) (string, error) {
	node, err := nd.Query(rule)
	if err != nil {
		return "", err
	}
	var result string
	if node.FirstChild != nil && len(node.FirstChild.Attr) == 0 && node.FirstChild.Data != "" {
		result = result + " " + node.FirstChild.Data
	}
	if node.LastChild != nil && len(node.LastChild.Attr) == 0 && node.LastChild.Data != "" {
		result = result + " " + node.LastChild.Data
	}
	if result != "" {
		return result, nil
	}
	var b bytes.Buffer
	if err := html.Render(&b, node); err != nil {
		return "", err
	}
	return b.String(), nil
}
func (nd *Node) QueryLastChildText(rule string) (string, error) {
	node, err := nd.Query(rule)
	if err != nil {
		return "", err
	}
	if node.LastChild != nil {
		return node.LastChild.Data, nil
	}
	return "", errors.New("not found LastChild")
}
func (nd *Node) QueryFirstChildText(rule string) (string, error) {
	node, err := nd.Query(rule)
	if err != nil {
		return "", err
	}
	if node.FirstChild != nil {
		return node.FirstChild.Data, nil
	}
	return "", errors.New("not found FirstChild")
}
