package main

import (
	"context"
	"os"
	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	_, err = client.Pipeline("test").
		Container().
		From("alpine@sha256:c5c5fda71656f28e49ac9c5416b3643eaa6a108a8093151d6d1afc9463be8e33").
		WithExec([]string{"apk", "add", "curl"}).
		WithExec([]string{"curl", "https://dagger.io"}).
		WithExec([]string{"sleep", "10"}).
		Sync(ctx)

	if err != nil {
		panic(err)
	}
//	fmt.Println(output)
}

