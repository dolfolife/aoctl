package aoc

import (
    "context"
    "io"
    "log"
    "net/http"
    "time"
)


func getBodyFromUrl(url string, cookie string) []byte {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

    if err != nil {
        log.Fatalf("Error making request: %s", err)
    }

    if cookie == "" {
        log.Fatal("Advent of Code requires a session and your cookie is empty")
    }

    sessionCookie := http.Cookie{Name: "session", Value: cookie}
    req.AddCookie(&sessionCookie)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatalf("Error fetching the page: %s", err)
    }

    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    
    if err != nil {
        log.Fatalf("Error reading the body of the request: %s", err)
    }

    return body
}
