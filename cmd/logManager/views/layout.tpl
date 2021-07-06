<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>日志收集系统</title>
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
        <a href="/logout" class="logout fr">退 出</a>
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

{{ .LayoutContent}}

</body>
</html>
