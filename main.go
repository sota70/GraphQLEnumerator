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
	v     *bool
	c     *bool
	u     *string
	query *string
}

func handle(args commandArgs) {
	switch {
	case *args.e:
		fmt.Println(enumeration.Enumerate(*args.u, *args.c))
	case *args.q:
		fmt.Println(query.Query(*args.u, *args.query, *args.c))
	case *args.v:
		fmt.Printf("GraphQL Visualizer: https://graphql-kit.com/graphql-voyager\n1: Go to the website\n2: Click CHANGE SCHEMA\n3: Go to INTROSPECTION\n4: Copy the introspection query provided in this section\n5: Run the query and retrieve the result\n6: Paste the result in INTROSPECTION section and click DISPLAY\n")
	default:
		fmt.Println("Usage: ./graphqlenumerator [-e | -q] [-u url] [-query query]")
	}
}

func main() {
	var (
		e     *bool   = flag.Bool("e", false, "Enumeration Mode")
		q     *bool   = flag.Bool("q", false, "Query Mode")
		v     *bool   = flag.Bool("v", false, "Shows graphql visualizer")
		c     *bool   = flag.Bool("c", false, "Copy the result to clipboard")
		query *string = flag.String("query", "{}", "GraphQL Query")
		u     *string = flag.String("u", "", "GraphQL endpoint URL")
	)
	flag.Parse()
	handle(commandArgs{e, q, v, c, u, query})
}
