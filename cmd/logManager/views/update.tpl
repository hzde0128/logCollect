
    <div class="main_body" id="main_body">
        <div class="breadcrub">
            当前位置：日志管理>编辑收集
        </div>
        <div class="pannel">
            <form name = logon  method="post" action="/admin/collect/{{ .collect.Id }}" enctype="multipart/form-data">
            <h3 class="review_title">编辑收集</h3>
                <div class="form_group">
                    <label>当前主机：</label>
                    <option>{{ .collect.Server.Address }}</option>
                    <input name="server" value="{{ .collect.Server.Address }}" hidden="hidden">
                </div>
                <div class="form_group">
                    <label>收集路径：</label>
                    <input type="text" class="input_txt2" name = "path" value="{{ .collect.Path }}">
                </div>
                <div class="form_group">
                    <label>Topic：</label>
                    <textarea class="input_multxt" name="topic">{{ .collect.Topic }}</textarea></textarea>
                </div>
                <div class="form_group indent_group line_top">
                    <input type="submit" value="修 改" class="confirm">
                </div>
        </form>

        </div>

    </div>
