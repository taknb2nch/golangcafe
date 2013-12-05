初めてのテンプレート。名前は {{.Name}} です。
所属しているグループは、{{len .Groups}} 件です。
{{range $index, $group := .Groups}}{{$index}} {{$group.Code}} {{$group.Name}} {{if $group.Leader}}リーダー{{end}} {{$group.Display}}
{{end}}
