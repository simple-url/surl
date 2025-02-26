# How To

## Send Raw Json Resquest
```json
{
    "requests": [
        {
            "name": "request_name",
            "url": "http://127.0.0.1:8000/posts",
            "method": "POST",
            "headers": [
                {
                    "key": "content-type",
                    "value": "application/json"
                }
            ],
            "body": "{\"hello\":\"world\"}",
        }
    ]
}
```
