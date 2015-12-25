<!DOCTYPE html>
<html>
<head>
	<title>信息提示</title>
</head>
<style type="text/css" media="screen">
body{font-family: 'Microsoft YaHei',Helvetica,Arial,sans-serif; padding:0; margin:0; border:0; color:#333;background-color: #eee;}
.container{ padding-top: 100px; width:400px; margin: 0 auto;}
p{ padding: 0; margin: 0}
.h1404{font-size: 64px; text-align: center; margin: 0px 0px 30px 0px; font-weight: normal;}
.nofound{height: 100px; background: url('/static/images/hehe.jpg') left center no-repeat; padding-left: 100px; line-height: 2em;}
a{ color: #0275d8;text-decoration:none;}
a:hover{ color: #0275d8;}
</style>
<body>
<div class="container">
	<h1 class="h1404">执行成功啦！</h1>
	<div class="nofound">
		<p>{{.Tip.Message}}</p>
        {{if .Tip.ReturnUrl eq ""}}
        <p><a href="javascript:;" onClick="window.history.go(-1)" title="返回">点击这里返回吧</a></p>
        <script>
		setTimeout(function(){
			window.history.go(-1);
		},3000);
		</script>
        {{else}}
		<p><a href="{{.Tip.ReturnUrl}}" title="返回">点击这里返回吧</a></p>
        <script>
		setTimeout(function(){
			window.location.href= "{{.Tip.ReturnUrl}}";
		},3000);
		</script>
        {{end}}
	</div>
</div>
</body>
</html>