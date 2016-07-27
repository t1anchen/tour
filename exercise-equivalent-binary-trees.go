package main

import (
	"golang.org/x/tour/tree"
	"log"
	"os"
	"text/template"
)

// Inspired from solutions on
// https://github.com/golang/tour/blob/master/solutions/binarytrees_quit.go

func WalkImpl(t *tree.Tree, ch, q chan int) {
	if t == nil {
		return
	}

	WalkImpl(t.Left, ch, q)

	select {
	case ch <- t.Value:
		// Value sent
	case <-q:
		// No extra elements
		return
	}

	WalkImpl(t.Right, ch, q)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch, q chan int) {
	WalkImpl(t, ch, q)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	q := make(chan int)
	defer close(q)

	go Walk(t1, ch1, q)
	go Walk(t2, ch2, q)

	for {
		node1, ok1 := <-ch1
		node2, ok2 := <-ch2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		return node1 == node2
	}
}

func main() {
	const t1 = "tree.New(1) {{if .Same}}equals{{- else}}does not equal{{- end}} tree.New(1)\n"
	const t2 = "tree.New(1) {{if .Same}}equals{{- else}}does not equal{{- end}} tree.New(2)\n"
	var same bool

	t_engine := template.New("tree")
	t1_tmpl := template.Must(t_engine.Parse(t1))
	t2_tmpl := template.Must(t_engine.Parse(t2))

	same = Same(tree.New(1), tree.New(1))
	if err := t1_tmpl.Execute(os.Stdout, map[string]bool{"Same": same}); err != nil {
		log.Fatal("t1: ", err)
	}

	same = Same(tree.New(1), tree.New(2))
	if err := t2_tmpl.Execute(os.Stdout, map[string]bool{"Same": same}); err != nil {
		log.Fatal("t2: ", err)
	}
}
