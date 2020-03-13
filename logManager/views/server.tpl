<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>编辑服务器</title>
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
            当前位置：日志管理>添加主机
        </div>
        <div class="pannel">
            <table class="common_table">
                <tr>
                    <th width="10%">id</th>
                    <th width="35%">主机名称</th>
                    <th width="35%">主机地址</th>
                    <th width="20%">管理操作</th>
                </tr>

                <tr>
                    <td>1</td>
                    <td>测试服务器</td>
                    <td>172.16.10.1</td>
                    <td><a href="javascript:;" class="edit">删除</a></td>
                </tr>

                <tr>
                    <td colspan="4">
                    <form action="/category" method="post">
                        <input type="text" class="type_txt" placeholder="主机名称" name="ServerName">
                        <input type="text" class="type_txt" placeholder="主机地址" name="ServerAddress">
                        <input type="submit" class="addtype" value="添加主机">
                    </form>
                    </td>
                </tr>
            </table>
        </div>
    </div>



</body>
</html>
