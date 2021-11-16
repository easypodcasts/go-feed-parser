# Feed Parser

Simple feed parser for [Easy Podcasts](https://roig.is-a.dev/podcasts/)

## Download
Visit the [release page](/releases/latest)

## Building from Source

Using [goreleaser](https://goreleaser.com/)
```bash
goreleaser build --snapshot --rm-dist  
```

Golang way:
```bash
go build -ldflags "-s -w" -o feed-parser   
```

## Usage

```bash
Usage of ./feed-parser:
  -feed.url string
        feed url to parse
  -v    prints current version
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.