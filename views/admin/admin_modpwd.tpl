
<div class="row">
    <ol class="breadcrumb m-b-sm">
      <li><a href="/admin">后台首页</a></li>
      <li><a href="/admin/admin">用户 </a></li>
      <li class="active">修改密码</li>
    </ol>
    
    <form method="post" action="/admin/admin/modpwd" name="modpwdForm" id="modpwdForm">
    	<ul class="nav nav-tabs" role="tablist">
          <li class="nav-item">
            <a class="nav-link active" href="#base" role="tab" data-toggle="tab">修改密码</a>
            <input type="hidden" name="action" id="action" value="{{.Action}}" />
          </li>
        </ul>
        <div class="tab-content">
          <div role="tabpanel" class="tab-pane active p-t-md" id="base">
          	<div class="form-group row">
            	<label for="Title" class="col-sm-2 form-control-label text-right">用户名：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="UserName" name="UserName" placeholder="请输入登录用户名，不少于5个字符" value="{{.UserName}}" required>
                </div>
            </div>
          	<div class="form-group row">
            	<label class="col-sm-2 text-right">旧密码：</label>
                <div class="col-sm-10">
                     <input type="password" class="form-control {minlength:5}" id="oldPwd" name="oldPwd" placeholder="请输入当前登录密码" required >
                </div>
            </div>
            <div class="form-group row">
            	<label class="col-sm-2 text-right">新密码：</label>
                <div class="col-sm-10">
                     <input type="password" class="form-control" id="newPwd" name="newPwd" placeholder="请输入新的登录密码，不少于5个字符" required >
                </div>
            </div>
            <div class="form-group row">
            	<label class="col-sm-2 text-right">确认密码：</label>
                <div class="col-sm-10">
                   <input type="password" class="form-control" id="newPwd2" name="newPwd2" placeholder="请输再一次入新的登录密码"  required >
                </div>
            </div>
          </div>
		<div>
        	<button type="submit" class="btn btn-primary"><i class="fa fa-floppy-o"></i> 提交</button>
            <a href="javascript:;" onClick="window.history.go(-1)" class="btn btn-secondary"><i class="fa fa-undo"></i> 返回</a>
        </div>
        </div>
    </form>
	<script>
	$(document).ready(function(e) {
        DoPost('modpwdForm');
    });
	</script>
</div>