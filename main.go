package main

import (
	"fmt"
	"github.com/dylan96dashintha/traceable_context/context"
)

func main() {

	traceCtx := context.WithTraceId()
	fmt.Println("traceId associated with : ", traceCtx.TraceId())
}
