
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
                {{ range $inx,$val := .server }}
                <tr>
                    <td>{{ $val.Id }}</td>
                    <td>{{ $val.Hostname }}</td>
                    <td>{{ $val.Address }}</td>
                    <td><a href="javascript:;" class="edit">删除</a></td>
                </tr>
                {{ end }}
                <tr>
                    <td colspan="4">
                    <form action="/admin/server/" method="post">
                        <input type="text" class="type_txt" placeholder="主机名称" name="ServerName">
                        <input type="text" class="type_txt" placeholder="主机地址" name="ServerAddress">
                        <input type="submit" class="addtype" value="添加主机">
                    </form>
                    </td>
                </tr>
            </table>
            <ul class="pagenation">
                <li><a href="/admin/server/">首页</a></li>
                <li><a href="/admin/server/?page={{ .page | prepage }}">上一页 </a> </li>
                <li> <a href="/admin/server/?page={{ nextpage .page .pageCount }}">下一页</a></li>
                <li><a href="/admin/server/?page={{.pageCount}}">末页</a></li>
                <li>共{{ .count }}条记录/共{{ .pageCount }}页/当前{{ .page }}页</li>
            </ul>
        </div>
    </div>
