# Clash of Clan API for GO
This is a simple fetch API for Clash of Clans in Go. It uses goroutines to fetch all members for a clan concurrently. 

## Installation 

    go get github.com/amir20/coc-api-go

## Example Use

```go
package main

import (
	"os"

	"github.com/amir20/coc-api-go"
)

func main() {
    token := os.Getenv("API_TOKEN")

    clashClient := clash.NewClient(token)
    clan, err := clashClient.Clan("#UGJPVJR")
    if err != nil {
        panic(err)
    }

    // Fetches all clash member list
    clashClient.FetchAllPlayers(clan)

    println(clan.Name)
}

```