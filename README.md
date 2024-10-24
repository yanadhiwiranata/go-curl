go-curl
=======

[![Build Status](https://secure.travis-ci.org/andelf/go-curl.png?branch=master)](http://travis-ci.org/andelf/go-curl)

go-curl is a GoLang interface to [libcurl](https://curl.haxx.se/libcurl/),
the multiprotocol file transfer library. Similar to the HTTP
support in [net/http] (https://pkg.go.dev/net/http), go-curl can be used to
fetch objects from a Go program. While go-curl can provide simple fetches,
it also exposes most of the functionality of libcurl, including:

 * Speed - libcurl is very fast.
 * Multiple protocol (not just HTTP).
 * SSL, authentication and proxy support.
 * Support for libcurl's callbacks.

This said, libcurl API can be less easy to learn than net/http.

LICENSE
-------

go-curl is licensed under the Apache License, Version 2.0 (http://www.apache.org/licenses/LICENSE-2.0.html).

Current Development Status
--------------------------

 * currently stable
 * READ, WRITE, HEADER, PROGRESS function callback
 * a Multipart Form supports file uploading
 * Most curl_easy_setopt option
 * partly implement share & multi interface
 * new callback function prototype

Requirements
------------
 * Any version of Go
 * libcurl 7.x or higher (including development headers and static/dynamic libs)
 * Python 3 (used only by configure scripts)


How to Install
--------------

    $ go get -u github.com/yanadhiwiranata/go-curl

Current Status
--------------

 * Linux x64
   * passed go1 (ArchLinux)
 * Windows x86
   * passed go1 (win7, mingw-gcc 4.5.2, curl 7.22.0)
 * Mac OS
   * passed go1 (Mac OS X 10.7.3, curl 7.21.4)

Sample Program
--------------

```go
package main

import (
    "fmt"
    curl "github.com/yanadhiwiranata/go-curl"
)

func main() {
    easy := curl.EasyInit()
    defer easy.Cleanup()

    easy.Setopt(curl.OPT_URL, "https://www.baidu.com/")

    // OPTIONAL - make a callback function
    fooTest := func (buf []byte, userdata interface{}) bool {
        println("DEBUG: size=>", len(buf))
        println("DEBUG: content=>", string(buf))
        return true
    }

    easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

    if err := easy.Perform(); err != nil {
        fmt.Printf("ERROR: %v\n", err)
    }
}
```

See also the [examples](./examples/) directory!
