package main

import (
	"flag"
	"fmt"
	"graphqlenumerator/enumeration"
	"graphqlenumerator/query"
)

type commandArgs struct {
	e     *bool
	q     *bool
	u     *string
	query *string
}

func handle(args commandArgs) {
	switch {
	case *args.e:
		fmt.Println(enumeration.Enumerate(*args.u))
	case *args.q:
		fmt.Println(query.Query(*args.u, *args.query))
	default:
		fmt.Println("Usage: ./graphqlenumerator [-e | -q] [-u url] [-query query]")
	}
}

func main() {
	var (
		e     *bool   = flag.Bool("e", false, "Enumeration Mode")
		q     *bool   = flag.Bool("q", false, "Query Mode")
		query *string = flag.String("query", "{}", "GraphQL Query")
		u     *string = flag.String("u", "", "GraphQL endpoint URL")
	)
	flag.Parse()
	handle(commandArgs{e, q, u, query})
}
