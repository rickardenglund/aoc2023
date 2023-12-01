package fetch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

func cacheFileName(d int) string {
	return fmt.Sprintf("inputs/day_%d.txt", d)
}

func Day(d int) string {
	if input, ok := getFromCache(d); ok {
		return input
	}

	fmt.Printf("fetching %d\n", d)
	req, err := http.NewRequestWithContext(context.Background(), "GET", fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", d), nil)
	if err != nil {
		panic(err)
	}

	cookie, err := os.ReadFile("inputs/.cookie")
	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", string(cookie))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		panic(err)
	}

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	os.Mkdir("inputs", os.ModePerm)
	if err := os.WriteFile(cacheFileName(d), bs, 0700); err != nil {
		println(err)
	}

	return string(bs)
}

func getFromCache(d int) (string, bool) {
	bs, err := os.ReadFile(cacheFileName(d))
	if err != nil {
		return "", false
	}

	return string(bs), true
}
