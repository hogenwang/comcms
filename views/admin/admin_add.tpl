
<div class="row">
    <ol class="breadcrumb m-b-sm">
      <li><a href="/admin">后台首页</a></li>
      <li><a href="/admin/admin">用户</a></li>
      {{if eq .Action "edit" }}
      <li class="active">修改管理员</li>
      {{else}}
      <li class="active">添加管理员</li>
      {{end}}
    </ol>
    
    <form method="post" {{if eq .Action "edit" }} action="/admin/admin/edit/{{.Entity.Id}}" {{else}} action="/admin/admin/add" {{end}} name="adminForm" id="adminForm">
    	<ul class="nav nav-tabs" role="tablist">
          <li class="nav-item">
            <a class="nav-link active" href="#base" role="tab" data-toggle="tab">管理员信息</a>
            <input type="hidden" name="action" id="action" value="{{.Action}}" />
          </li>
        </ul>
        <div class="tab-content">
          <div role="tabpanel" class="tab-pane active p-t-md" id="base">
          	<div class="form-group row">
            	<label for="Title" class="col-sm-2 form-control-label text-right">用户名：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="UserName" name="UserName" placeholder="请输入登录用户名，不少于5个字符" value="{{.Entity.UserName}}" required>
                </div>
            </div>
            {{if eq .Action "edit" }}
            <div class="form-group row">
            	<label class="col-sm-2  form-control-label text-right">登录密码：</label>
                <div class="col-sm-10">
                     <input type="password" class="form-control" id="PassWord" name="PassWord" placeholder="输入则修改该管理员密码，留空则不修改"  data-toggle="tooltip" data-placement="top" >
                     <input type="hidden" name="Id" id="Id" value="{{.Entity.Id}}" />
                </div>
            </div>
            {{else}}
          	<div class="form-group row">
            	<label class="col-sm-2  form-control-label text-right">登录密码：</label>
                <div class="col-sm-10">
                     <input type="password" class="form-control" id="PassWord" name="PassWord" placeholder="请输入该管理员登录密码" required >
                </div>
            </div>
            {{end}}
            <div class="form-group row">
            	<label class="col-sm-2  form-control-label text-right">昵称：</label>
                <div class="col-sm-10">
                     <input type="text" class="form-control" id="NickName" name="NickName" placeholder="可以选填该管理员昵称" value="{{.Entity.NickName}}" >
                </div>
            </div>
            {{if eq .Action "edit" }}
            <div class="form-group row">
            	<label class="col-sm-2  form-control-label text-right">上次登录时间：</label>
                <div class="col-sm-10">
                     <input type="text" class="form-control" id="ThisLoginTime" name="ThisLoginTime" placeholder="上次登录时间" value="{{date .Entity.ThisLoginTime "Y-m-d H:i:s"}}" readonly >
                </div>
            </div>
            <div class="form-group row">
            	<label class="col-sm-2  form-control-label text-right">上次登录IP：</label>
                <div class="col-sm-10">
                     <input type="text" class="form-control" id="ThisLoginIP" name="ThisLoginIP" placeholder="上次登录IP" value="{{.Entity.ThisLoginIP}}" readonly >
                </div>
            </div>
            <div class="form-group row">
            	<label class="col-sm-2  form-control-label text-right">登录次数：</label>
                <div class="col-sm-10">
                     <input type="text" class="form-control" id="LoginCount" name="LoginCount" placeholder="登录次数" value="{{.Entity.LoginCount}}" readonly >
                </div>
            </div>
            {{end}}
          </div>
		<div>
        	<button type="submit" class="btn btn-primary"><i class="fa fa-floppy-o"></i> 提交</button>
            <a href="javascript:;" onClick="window.history.go(-1)" class="btn btn-secondary"><i class="fa fa-undo"></i> 返回</a>
        </div>
        </div>
    </form>
	<script>
	$(document).ready(function(e) {
        DoPost('adminForm');
    });
	</script>
</div>