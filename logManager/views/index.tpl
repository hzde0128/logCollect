
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
                        <td><a href="/admin/collect/{{ .Id }}">编辑</a></td>
                        <td>{{ .Topic }}</td>
                        <td>{{ .Server.Address }}</td>
                    </tr>
                    {{ end }}
            </table>

            <ul class="pagenation">
                <li><a href="/admin/">首页</a></li>
                <li><a href="/admin/?page={{ .page | prepage }}">上一页 </a> </li>
                <li> <a href="/admin/?page={{ nextpage .page .pageCount }}">下一页</a></li>
                <li><a href="/admin/?page={{.pageCount}}">末页</a></li>
                <li>共{{ .count }}条记录/共{{ .pageCount }}页/当前{{ .page }}页</li>
            </ul>
        </div>
    </div>