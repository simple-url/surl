# Surl JSON Metadata
Basic surl.json structure
```json
{
    "requests": [
        {
            "name": "first_request",
            "url": "https://127.0.0.1:8000/todos",
            "method": "GET"
        },
        {
            "name": "second_request",
            "url": "https://127.0.0.1:8000/todos/2",
            "method": "GET"
        },
    ]
}
```
Here are key and value for surl.json

| Key     | Value Type                                       | For                               | is required |
|---------|--------------------------------------------------|-----------------------------------|-------------|
| name    | string                                           | request indentifier               | yes         |
| url     | string                                           | request url                       | yes         |
| method  | string (GET/POST/PUT/PATCH/DELETE/OPTIONS)       | request method                    | yes         |
| body    | string                                           | request body                      | no          |
| headers | [{"key": string, "value": string}]               | set request header                | no          |
| timeout | int                                              | set timeout (default: no timeout) | no          |
| json    | any (see [json](./how-to.md#send-json-resquest)) | set json body                     | no          |
