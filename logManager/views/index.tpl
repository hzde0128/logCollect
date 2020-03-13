<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>日志列表 - 日志收集系统</title>
    <link rel="stylesheet" type="text/css" href="/static/css/reset.css">
    <link rel="stylesheet" type="text/css" href="/static/css/main.css">
    <script type="text/javascript" src="/static/js/jquery-1.12.4.min.js"></script>
    <script type="text/javascript">
        window.onload = function () {
            $(".dels").click(function () {
                if(!confirm("是否删除?")){
                    return false
                }
            })
            $("#select").change(function () {
                $("#form").submit()
            })
        }
    </script>
</head>
<body>

    <div class="header">
        <a href="#" class="logo fl"><img src="/static/img/logo.png" alt="logo"></a>
        <a href="#" class="logout fr">退 出</a>
    </div>

    <div class="side_bar">
        <div class="user_info">
            <img src="/static/img/person.png" alt="">
            <p>欢迎你 <em>admin</em></p>
        </div>

        <div class="menu_con">
            <div class="first_menu active"><a href="javascript:;" class="icon02">日志管理</a></div>
            <ul class="sub_menu show">
                <li><a href="/admin/" class="icon031">日志列表</a></li>
                <li><a href="/admin/collect/" class="icon032">添加收集</a></li>
                <li><a href="/admin/server/" class="icon034">添加主机</a></li>
            </ul>
        </div>
    </div>

    <div class="main_body" id="main_body">
        <div class="breadcrub">
            当前位置：日志管理>日志列表
        </div>
        <div class="pannel">
            <span class="sel_label">请选择主机：</span>
            <form id="form" method="post" action="/admin/">
                <select name="select" id="select" class="sel_opt">
                    {{ range .server }}
                        <option selected="true">{{ .Address }}</option>
                    {{ end }}
                </select>
                <input type="submit" hidden="hidden">
            </form>

            <table class="common_table">
                <tr>
                    <th width="45%">日志路径</th>
                    <th width="15%">添加时间</th>
                    <th width="7%">删除</th>
                    <th width="7%">编辑</th>
                    <th width="10%">Topic</th>
                    <th width="16%">主机地址</th>
                </tr>
                    {{ range .collect }}
                    <tr>
                        <td>{{ .Path }}</td>
                        <td>{{ .CreateTime.Format "2006-01-02 13:04:05" }}</td>
                        <td><a href="#" class="dels">删除</a></td>
                        <td><a href="/admin/colletc/{{ .Id }}">编辑</a></td>
                        <td>{{ .Topic }}</td>
                        <td>{{ .Server.Address }}</td>
                    </tr>
                    {{ end }}
            </table>

            <ul class="pagenation">
                <li><a href="/admin/">首页</a></li>
{{/*                <li><a href="/article?pageIndex={{.pageIndex | ShowPrePage }}">上一页 </a> </li>*/}}
{{/*                <li> <a href="/article?pageIndex={{ ShowNextPage .pageIndex .pageCount }}">下一页</a></li>*/}}
{{/*                <li><a href="/article?pageIndex={{.pageCount}}">末页</a></li>*/}}
{{/*                <li>共{{.count}}条记录/共{{.pageCount}}页/当前{{.pageIndex}}页</li>*/}}
            </ul>
        </div>
    </div>
</body>
</html>
