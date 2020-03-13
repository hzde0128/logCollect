
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

