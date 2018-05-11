package mapTree

type MT struct {
	Root *Node
}

type Node struct {
	Childs  []*Node
	Parent  *Node
	Content interface{}
}

func NewMT(i interface{}) *MT {
	return &MT{
		Root: &Node{
			Content: i,
		},
	}
}
func (mt *MT) AddNode(i interface{}) {
	mt.Root.AddNode(i)
}
func (n *Node) AddNode(i interface{}) {
	n.Childs = append(n.Childs, &Node{
		Content: i,
		Parent:  n,
	})
}

func (mt *MT) DeleteNode(i interface{}) {
	mt.Root.DeleteNode(i)
}
func (n *Node) DeleteNode(i interface{}) {
	for ii, d := range n.Childs {
		if d.Content == i {
			n.Childs = append(n.Childs[:ii], n.Childs[ii+1:]...)
		}
	}
}

func (mt *MT) GetNode() []*Node {
	return mt.Root.Childs
}

func (n *Node) GetNode() []*Node {
	return n.Childs
}
