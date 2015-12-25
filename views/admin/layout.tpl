<!DOCTYPE html>

<html>
<head>
  <title>{{.Title}}</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta http-equiv="x-ua-compatible" content="ie=edge">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="/static/js/bootstrap/css/bootstrap.css">
  <link rel="stylesheet" href="/static/admin/style.css">
  <script src="/static/js/jquery2.js" type="text/javascript" charset="utf-8"></script>
  <script src="/static/js/bootstrap/js/bootstrap.min.js" type="text/javascript" charset="utf-8" async defer></script>
  <script src="/static/js/jquery.form.min.js" async defer></script>
  <script src="/static/js/validate/jquery.validate.min.js" async defer></script>
  <script src="/static/js/validate/messages_zh.min.js" defer></script>
  <script src="/static/js/jquery.metadata.js" async defer></script>
  <script src="/static/js/admin.js?v=2015-10-11" async defer></script>
  <script src="/static/js/my97/WdatePicker.js"></script>
  <script src="/static/js/layer/layer.js"></script>
  <link rel="stylesheet" href="/static/awesome/css/font-awesome.min.css">
  <script src="/static/js/webuploader/webuploader.min.js"></script>
  <script src="/static/js/uploader.js"></script>
</head>

<body>
    <div>
		<nav class="navbar navbar-dark bg-primary navbar-inverse navbar-fixed-top">
			  <a class="navbar-brand" href="#">COMCMS</a>
			  <ul class="nav navbar-nav" id="topnavbar">
			    <li class="nav-item active" id="home">
			      <a class="nav-link" href="javascript:;">首页</a>
			    </li>
			    <li class="nav-item" id="system">
			      <a class="nav-link" href="javascript:;">系统</a>
			    </li>
			    <li class="nav-item" id="article">
			      <a class="nav-link" href="javascript:;">文章</a>
			    </li>
                <li class="nav-item" id="user">
			      <a class="nav-link" href="javascript:;">用户</a>
			    </li>
			    <li class="nav-item" id="other">
			      <a class="nav-link" href="javascript:;">其他</a>
			    </li>
			  </ul>
              <ul class="nav navbar-nav pull-right">
              	<li class="nav-item"><a class="nav-link btn btn-primary" href="#" title="修改密码">欢迎:{{.AdminName}}</a></li>
                <li class="nav-item"><a class="nav-link" href="/" target="_blank" title="网站首页">网站首页</a></li>
                <li class="nav-item"><a class="nav-link" href="/admin/logout" title="退出登录">退出登录</a></li>
              </ul>
			  <form class="form-inline navbar-form pull-right">
			    <input class="form-control" type="text" placeholder="文章关键字">
			    <button class="btn btn-success-outline" type="submit">搜索</button>
			  </form>
              
		</nav>
		<div class="container-fluid">
			<div class="row">
		        <div class="col-sm-3 col-md-2 sidebar" id="leftnav">
                  <ul class="nav nav-pills nav-stacked nav-sidebar" id="home_subul">
                      <li class="nav-item">
                        <a href="#" class="nav-link active">后台首页</a>
                      </li>
                	</ul>
                    <ul class="nav nav-pills nav-stacked nav-sidebar hide" id="system_subul">
                    	<li class="nav-item">
                        <a href="/admin/config" class="nav-link">系统设置</a>
                      </li>
                    </ul>
                    <ul class="nav nav-pills nav-stacked nav-sidebar hide" id="article_subul">
                    	<li class="nav-item">
                        <a href="/admin/category" class="nav-link">文章栏目管理</a>
                      </li>
                      <li class="nav-item">
                        <a href="/admin/article" class="nav-link">文章管理</a>
                      </li>
                    </ul>
                    <ul class="nav nav-pills nav-stacked nav-sidebar hide" id="user_subul">
                    	<li class="nav-item">
                        <a href="/admin/admin/modpwd" class="nav-link">修改密码</a>
                      </li>
                      <li class="nav-item">
                        <a href="/admin/admin/list" class="nav-link">管理员管理</a>
                      </li>
                    </ul>
                    <ul class="nav nav-pills nav-stacked nav-sidebar hide" id="other_subul">
                      <li class="nav-item">
                        <a href="/admin/link" class="nav-link">友情链接管理</a>
                      </li>
                      <li class="nav-item">
                        <a href="/admin/guestbook" class="nav-link">留言板管理</a>
                      </li>
                     <li class="nav-item">
                        <a href="/admin/ads" class="nav-link">广告管理</a>
                      </li>
                    </ul>
		        </div>
                <div class="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main">
                    {{.LayoutContent}}
                </div>
		    </div>
            <footer class="footer"><p>Copyright &copy; COMCMS 2015</p></footer>
		</div>
	</div>
    <script>
	$(function(){
		$("#topnavbar li").on("click",function(){
			var id =$(this).attr("id");
			$("#topnavbar li.active").removeClass('active');
			$(this).addClass('active');
			$("#leftnav ul").hide();
			$("#"+id+"_subul").show();
		});
	});
	</script>
</body>
</html>
