<script>
	$(function(){
		//初始化上传控件
        $(".upload-img").InitUploader({ filesize: "10240", sendurl: "/admin/webupload", swf: "/static/js/webuploader/uploader.swf", filetypes: "gif,jpg,png,bmp" });

	});
	</script>
<div class="row">
    <ol class="breadcrumb m-b-sm">
      <li><a href="/admin">后台首页</a></li>
      <li><a href="/admin/link">其他</a></li>
      {{if eq .Action "edit" }}
      <li class="active">修改友情链接详情</li>
      {{else}}
      <li class="active">添加友情链接详情</li>
      {{end}}
    </ol>
    
    <form method="post" {{if eq .Action "edit" }} action="/admin/link/edit/{{.Entity.Id}}" {{else}} action="/admin/link/add" {{end}} name="linkForm" id="linkForm">
    	<ul class="nav nav-tabs" role="tablist">
          <li class="nav-item">
            <a class="nav-link active" href="#base" role="tab" data-toggle="tab">友情链接详情</a>
            <input type="hidden" name="action" id="action" value="{{.Action}}" />
            <input type="hidden" name="Id" id="Id" value="{{.Entity.Id}}" />
          </li>
        </ul>
        <div class="tab-content">
          <div role="tabpanel" class="tab-pane active p-t-md" id="base">
          	<div class="form-group row">
            	<label for="Title" class="col-sm-2 form-control-label text-right">站点名称：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="Title" name="Title" placeholder="请输入站点名称" value="{{.Entity.Title}}" required>
                </div>
            </div>
            <div class="form-group row">
            	<label class="col-sm-2  form-control-label text-right">站点URL：</label>
                <div class="col-sm-10">
                     <input type="text" class="form-control" id="Url" name="Url" placeholder="请输入站点URL地址，如：http://www.comcms.com" value="{{.Entity.Url}}" required >
                </div>
            </div>
            <div class="form-group row">
            	<label for="Pic" class="col-sm-2 form-control-label text-right">站点Logo：</label>
                <div class="col-sm-10 input-group">
                  <input type="text" class="form-control upload-path" id="Logo" name="Logo" placeholder="请选择一张站点LOGO图片，可选" value="{{.Entity.Logo}}">
                  <span class="input-group-addon my-upload-span upload-img" id="basic-addon2"></span>
                </div>
            </div>
            <div class="form-group row">
            	<label class="col-sm-2  form-control-label text-right">排序：</label>
                <div class="col-sm-10">
                     <input type="text" class="form-control" id="Rank" name="Rank" placeholder="排序，越小排越前" value="{{.Entity.Rank}}" >
                </div>
            </div>
            <div class="form-group row">
            	<label for="Description" class="col-sm-2 form-control-label text-right">描述：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="Description" name="Description" value="{{.Entity.Description}}" placeholder="描述，可选">
                </div>
            </div>
            <div class="form-group row">
            	<label class="col-sm-2 text-right">隐藏链接：</label>
                <div class="col-sm-10">
                	<div class="checkbox">
                        <label>
                          <input type="checkbox" value="1" name="IsHide" id="IsHide" {{if eq .Entity.IsHide 1}} checked {{end}} > 隐藏 （勾选则该链接隐藏，前台不显示）
                        </label>
                    </div>
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
        DoPost('linkForm');
    });
	</script>
</div>