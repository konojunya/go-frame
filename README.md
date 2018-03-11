# go-frame
A Go library that displays a flat structure beautifully

![](https://raw.githubusercontent.com/konojunya/go-frame/master/screenshot/screen.png)

## Description

A Go library made to display a flat structure beautifully on the console.

## Installation

```
$ go get -d github.com/konojunya/go-frame
```

## Usage

```
package main

import (
	frame "github.com/konojunya/go-frame"
)

// User User model
type User struct {
	Name string `frame:"Customer Name"`
	Age  int    `frame:"Age"`
	URL  string `frame:"HomePage URL"`
}

func main() {
	user := User{
		Name: "konojunya",
		Age:  21,
		URL:  "https://konojunya.com",
	}

	frame.Print(user)
}
```

## Development

```
$ dep ensure
```

## Contribution

Please check the [issue](https://github.com/konojunya/go-frame/issues).

1. Fork it https://github.com/konojunya/go-frame
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create new Pull Request! :)

## Licence

MIT

## Auther

- Twitter [@konojunya](https://twitter.com/konojunya)
