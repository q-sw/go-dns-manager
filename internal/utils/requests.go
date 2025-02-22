package utils

import (
    "bytes"
    "fmt"
    "io"
    "net/http"

    "github.com/spf13/viper"
)

type Header struct {
    Property string
    Value    string
}

func HttpRequest(httpMethod, url string, bodyResquest []byte, headers []Header) ([]byte, int) {
    r, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(bodyResquest))
    if err != nil {
        fmt.Errorf("impossible to create new request for URL %v with error %v", url, err)
    }

    r.Header.Add("authorization", fmt.Sprintf("Bearer %v", viper.GetString("apiToken")))
    r.Header.Add("content-type", "application/json")

    client := http.Client{}
    res, err := client.Do(r)
    if err != nil {
        fmt.Errorf("error during the request with err %v", err)
    }

    defer res.Body.Close()

    if res.StatusCode != 200 {
        fmt.Printf("request error %v \n", res.StatusCode)
        return []byte{}, res.StatusCode
    }

    b, _ := io.ReadAll(res.Body)
    return b, res.StatusCode
}
