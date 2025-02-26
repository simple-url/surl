# SURL

Simple curl (SURL). Simple http client that aim to be between curl and desktop http client (postman, insomnia).

## Why using surl
When we want to make a http request, usually we gonna use an application such as postman or insomnia.
Think about it, for a simple http request we have to open a full fledge app, That many feature in that app
that we don't use (in simple term bloated). Another alternative are using cli app like curl or wget. The downside
of curl and wget is we have to type everything in cli. Make it prone to typo and make it hard for sending request
with complicated body (json, xml, file, etc).

## So how surl solve this problem??
Surl is a cli tool like curl or wget. Instead of type all command on cli, surl read from json file who has http request spec.
it make it easy to reproduce http request. Here an example of surl.json file.
```json
{
    "requests": [
        {
            "name": "first",
            "url": "https://jsonplaceholder.typicode.com/todos/2",
            "method": "GET"
        },
        {
            "name": "second",
            "url": "https://jsonplaceholder.typicode.com/posts",
            "method": "POST",
            "headers": [
                {
                    "key": "content-type",
                    "value": "application/json"
                }
            ],
            "body": "{\"hello\":\"world\"}",
            "timeout": 1
        }
    ]
}
```
After that run it using `surl run <name>`
```sh
surl run first
```
```sh
{
  "userId": 1,
  "id": 2,
  "title": "quis ut nam facilis et officia qui",
  "completed": false
}
```
## Installation
### Using Go Install
make sure yo have minimum go version 1.23 
(for mac and linux make sure you have setup your Go Path see [setup go path](https://stackoverflow.com/questions/36083542/error-command-not-found-after-installing-go-eval))
```sh
go install github.com/simple-url/surl
```
then you are good to go
```sh
surl help
```
### Build From Source
make sure yo have minimum go version 1.23 
- clone this repository `git clone https://github.com/simple-url/surl.git`
- build project using go `go build`, it will produce `surl` executable
- run executable `surl help`
- now is up to you where and how you want to use the executable

## Read More
- [How To](./docs/how-to.md)
- [surl cli command](./docs/cli_command.md)
- [surl.json metadata](./docs/surl_json.md)
