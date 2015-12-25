<!DOCTYPE html>

<html>
<head>
  <title>登录系统管理</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta http-equiv="x-ua-compatible" content="ie=edge">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="/static/js/bootstrap/css/bootstrap.css">
  <link rel="stylesheet" href="/static/admin/style.css">
  <script src="/static/js/jquery2.js" type="text/javascript" charset="utf-8"></script>
  <script src="/static/js/bootstrap/js/bootstrap.min.js" type="text/javascript" charset="utf-8" async defer></script>
  <link rel="stylesheet" href="/static/awesome/css/font-awesome.min.css">
  <script src="/static/js/jquery.form.min.js" ></script>
  <script src="/static/js/validate/jquery.validate.min.js" ></script>
  <script src="/static/js/validate/messages_zh.min.js" ></script>
  <script src="/static/js/layer/layer.js"></script>
  <script src="/static/js/admin.js?v=2015-10-11"></script>
</head>

<body class="loginbody">
    <div class="container">

      <form class="form-signin" name="loginForm" id="loginForm" action="/admin/login" method="post">
        <h2 class="form-signin-heading">管理员登录</h2>
        <label for="username" class="sr-only">管理账号</label>
        <input type="text" id="username" name="username" class="form-control" placeholder="请输入管理账号" required autofocus>
        <label for="password" class="sr-only">管理密码</label>
        <input type="password" id="password" name="password" class="form-control" placeholder="请输入登录密码" required>
        <div class="tip">
          <span class="label label-pill label-info">!</span>
          <label>
            {{.Tip}}
          </label>
        </div>
        <button class="btn btn-lg btn-primary btn-block" type="submit">立刻登录</button>
      </form>
    </div> <!-- /container -->
    <script>
	$(document).ready(function(e) {
        DoPost('loginForm');
    });
	</script>
</body>
</html>
