<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>添加收集 - 日志收集系统</title>
    <link rel="stylesheet" type="text/css" href="/static/css/reset.css">
    <link rel="stylesheet" type="text/css" href="/static/css/main.css">
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
            当前位置：日志管理>添加收集
        </div>
        <div class="pannel">
            <form method="post" action="/admin/collect/" enctype="multipart/form-data">
            <h3 class="review_title">添加收集</h3>
                <div class="form_group">
                    <label>主机选择：</label>
                    <select class="sel_opt" name="server">
                        {{ range .server }}
                            <option>{{ .Address }}</option>
                        {{ end }}
                    </select>
                </div>
            <div class="form_group">
                <label>日志路径：</label>
                <input type="text" class="input_txt2" name="filePath" >
            </div>
                <div class="form_group">
                    <label>Topic：</label>
                    <input type="text" class="input_txt2" name="topic" >
                </div>

            <div class="form_group indent_group line_top">
                <input type="submit" value="添 加" class="confirm">
            </div>
        </form>
        </div>
</div>

</body>
</html>
