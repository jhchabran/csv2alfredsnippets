# csv2alfredsnippets

A small script that converts snippets formatted in CSV to Alfred 3 snippets format. Mostly useful to programmatically
generate snippets.

## Getting started

### Prerequisites

Go >= 1.11 is installed on your system.

### Installing

```shell
GO111MODULE=off go get -u github.com/jhchabran/csv2alfredsnippets
```

### Running

```shell
csv2alfredsnippets snippets.csv snippets.alfredsnippets
```

## CSV format

- `name`
  - what Alfred will display to describe the snippet
- `keyword`
  - what you need to type to trigger the snippet
- `snippet`
  - what it will expand to

| name | keyword | snippet |
| ---- | ------- | ------- |
| My address | myaddr | 123 somewhere, 4567 Internet |
| Bob weird github | gh:bob | @that_weird_github_handle |

:arrow_down: Raw csv

```csv
My address,myaddr,123 somewhere 4567 Internet
Bob weird github,gh:bob,@that_weird_github_handle
```

:bulb: _Having the name show somehow the keyword is helpful to learn it overtime_.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details
