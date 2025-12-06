package golang

import (
	"fmt"
	schema "opticode/compile/golang/golang"
	"sync"
)

//go:generate bash ./generate.sh

// Compiles buffer into go files.
// lut -> Look-up table for strings used in the program
// buf -> flatbuffer's buffer
func Compile(buf *[]byte) ([]*GoFile, error) {
	program := schema.GetRootAsProgram(*buf, 0)

	nodesLength := program.NodesLength()
	gen := NewGenerator(program, buf)

	const maxRoutines = 5 // maximum allowed concurrent goroutines

	sem := make(chan struct{}, maxRoutines) // semaphore channel
	var wg sync.WaitGroup

	for i := range nodesLength - 1 {
		sem <- struct{}{} // acquire a "slot" (pauses if full)
		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			defer func() { <-sem }() // release the slot when done

			var node *schema.Node
			program.Nodes(node, index)

			gen.Eval(node)
		}(i)
	}

	wg.Wait()
	fmt.Println("All tasks completed")

	return nil, nil
}
