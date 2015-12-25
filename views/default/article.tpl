<!DOCTYPE html>

<html>
<head>
  <title>{{.Website}}</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <script src="/static/js/jquery2.js" type="text/javascript" charset="utf-8"></script>
  <script type="text/javascript" src="/static/editor/ueditor/ueditor.config.js"></script>
  <script type="text/javascript" src="/static/editor/ueditor/ueditor.all.min.js"></script>
  <script type="text/javascript" src="/static/editor/ueditor/lang/zh-cn/zh-cn.js"></script>

</head>

<body>
  <header>
    <h1 class="logo">Welcome to COMCMS</h1>
  </header>
  <div>
    <textarea class="form-control" id="Content" name="Content" rows="3" style="width:99.9%"></textarea>
  </div>
  <script>
          $(document).ready(function () {var editor = new UE.ui.Editor({ initialFrameHeight: 200});editor.render('Content')})
        </script>
</body>
</html>
