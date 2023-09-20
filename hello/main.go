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
		From("alpine@sha256:c5c5fda71656f28e49ac9c5416b3643eaa6a108a8093151d6d1afc9463be8e33").
		WithExec([]string{"apk", "add", "curl"}).
		WithExec([]string{"curl", "https://dagger.io"}).
		Stdout(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Println(output[:300])
}

