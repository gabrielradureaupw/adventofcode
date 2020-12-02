package inputs

import (
	"fmt"
	"io"
	"net/http"
)

const inputURLFormat = "https://adventofcode.com/2020/day/%d/input"
const session = "53616c7465645f5f75b77582051a3867a4aac617492d96e39cf78b7f19f6dc846df755d5472c952504abcf4ae859d161" // expires in nov 2030

// FetchInput _
func FetchInput(day int) io.Reader {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(inputURLFormat, day), nil)
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	return resp.Body
}
