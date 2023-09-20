package main

import (
	"context"
	"fmt"
	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	output, err := client.Pipeline("test").
		Container().
		From("alpine").
		WithExec([]string{"apk", "add", "curl"}).
		WithExec([]string{"curl", "https://dagger.io"}).
		Stdout(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Println(output[:300])
}

