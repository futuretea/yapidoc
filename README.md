# yapidoc
yapidoc is a doc server for [YMFE/Yapi](https://github.com/YMFE/yapi) .

## Usage
1. run a yapidoc server using docker  
```bash
docker run -itd --name yapidoc -e YAPI_URL=your_yapi_url -p 8888:8888 futuretea/yapidoc
```
2. download doc from  the yapidoc server
```javascript
http://127.0.0.1:8888/projects/your_project_token?tag=your_api_tag
```

3. stop the yapidoc server
```bash
docker stop yapidoc
docker rm yapidoc
```

## Documentation
- [FAQ](https://github.com/futuretea/yapidoc/wiki/FAQ)

## Related Projects
- [YMFE/yapi](https://github.com/YMFE/yapi)
- [futuretea/go-yapi](https://godoc.org/github.com/futuretea/go-yapi) - Go client library for YMFE/Yapi
- [futuretea/yapi](https://github.com/futuretea/yapi) - docker image for YMFE/Yapi

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).
