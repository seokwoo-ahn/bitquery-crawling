package main

import (
	graphql "blockdata-crawling/src/graphQL"
	_ "blockdata-crawling/src/spread_sheet"
)

func main() {
	graphql.Api()
}
