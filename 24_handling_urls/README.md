# Handling URLs In Go


URLs provide a uniform way to locate resources. Here’s how to parse URLs in Go.

## Parsing The URL And Getting The Components

* It basically means extracting the information from thr URL.
* It is breaking down the URL into its individual components, such as the scheme, host, port, path, query, and fragment.
    ```go
    const myUrl = "postgres://user:pass@host.com:5432/path?k=v#f"

    /* parsing */
    result, _ := url.Parse(myUrl) // Here, we have ignored the error. It's better to handle the same.

    fmt.Println("Scheme:", result.Scheme)
    fmt.Println("Host:", result.Host)
    fmt.Println("Port:", result.Port())
    fmt.Println("Path:", result.Path)

    /* extract the path and the fragment after the #. */
    fmt.Println("Fragment:", result.Fragment)
    fmt.Println("Raw Query:", result.RawQuery)
    ```


## Parse Query Params Into A Map

* The parsed query param **maps** are from strings to **slices of strings**, so index into `[0]`, if you only want the first value.

    ```go
    queryParams := result.Query()

    fmt.Printf("Type of queryParams: %T\n", queryParams)
    fmt.Println(queryParams["k"])
    fmt.Println(queryParams["k"][0])
    ```


## Making URL From Components

* Just like URL can be parsed into components, the vice-versa is also possible.
    ```go
    partsOfUrl := &url.URL{
        Scheme:   "https",
        Host:     "abcd.dev",
        Path:     "/path",
        RawQuery: "user=goku",
    }

    fullUrl := partsOfUrl.String()
    fmt.Println("full url: ", fullUrl)
    ```