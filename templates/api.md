# {{ .Name }} API文档
{{- range $_, $service := .Services}}
## {{ $service.Name }}
{{- range $_, $api := $service.APIS}}
### {{ $api.Title }}

> 基本信息

**Path：** {{ $api.Path }}
**Method：** {{ $api.Method }}
**status：** {{ $api.Status }}

> 请求参数

**Headers**
| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
{{- range $_, $item := $api.ReqHeaders}}
| {{$item.Name}}  | {{$item.Value}}  | {{if eq $item.Required "1"}}必须{{else}}非必须{{end}} | {{$item.Example}}  |  {{$item.Desc}} |
{{- end}}

**路径参数**
| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
{{- range $_, $item := $api.ReqParams}}
| {{$item.Name}} | {{$item.Example}}  |  {{$item.Desc}} |
{{- end}}

**Query**
| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
{{- range $_, $item := $api.ReqQuery}}
| {{$item.Name}} | {{if eq $item.Required "1"}}必须{{else}}非必须{{end}} | {{$item.Example}}  |  {{$item.Desc}} |
{{- end}}

**Body**
{{- if eq $api.ReqBodyType "form"}}
| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
{{- range $_, $item := $api.ReqBodyForm}}
| {{$item.Name}}  | {{$item.Value}}  | {{$item.Required}}  | {{$item.Example}}  |  {{$item.Desc}} |
{{- end}}
{{- else}}
```{{if eq $api.ReqBodyType "json"}}json5{{else}}text{{end}}
{{ $api.ReqBodyOther }}
```
{{- end}}

> 返回数据

```{{if eq $api.ResBodyType "json"}}json5{{else}}text{{end}}
{{ $api.ResBody }}
```
{{- end}}
{{- end}}