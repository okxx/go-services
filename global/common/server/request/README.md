### REQUEST

#### request package provides calls to external services 

```
调用方式

    sapi := "/user/common/getInfo"
    params := map[string]interface{}{
        "userId" : 1
    }
    request.sync(sapi string,params map[string]interface{})
    
返回值：
    json
        {
            "state" : 200,
            "msg"   : "success"
            "data"  : {
                "userId" : 1,
                "nickname" : "sapi"
            }   
        }
    
```