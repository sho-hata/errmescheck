# errmesscheck
Go linter that checks error message format.

supported package: https://github.com/pkg/errors

## Motivation
We want to unify the format of error messages.

OK pattern🕺
```
failed to xxxxxxx.
```

NG pattern👮🏻
```
xxxxxxx failed.
```
```go

```
## Usage
```go
package a

import (
	"strconv"

	"github.com/pkg/errors"
)

func Sample() error {
	if _, err := strconv.Atoi("1"); err != nil {
		return errors.Wrap(err, "failed to strconv.Atoi")
	}

	return nil
}
```
## Analysis
```bash
$ errorsmescheck ./...
./a.go:11:10: The prefix of the error message should be 'failed to ...'
```

## Installation
```
go install github.com/sho-hata/errmescheck/cmd/errmescheck@latest
```

## Contribution
1. Fork (https://github.com/sho-hata/errmescheck/fork)
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the go `test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

## License
[MIT](https://github.com/sho-hata/errmescheck/blob/main/LICENSE)

## Author
[sho-hata](https://github.com/sho-hata)
