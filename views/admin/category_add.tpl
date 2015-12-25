    <script type="text/javascript" src="/static/editor/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" src="/static/editor/ueditor/ueditor.all.min.js"></script>
    <script type="text/javascript" src="/static/editor/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script>
	$(function(){
		//初始化上传控件
        $(".upload-img").InitUploader({ filesize: "10240", sendurl: "/admin/webupload", swf: "/static/js/webuploader/uploader.swf", filetypes: "gif,jpg,png,bmp" });

	});
	</script>
<div class="row">
    <ol class="breadcrumb m-b-sm">
      <li><a href="/admin">后台首页</a></li>
      <li><a href="/admin/category">文章栏目管理</a></li>
      {{if eq .Action "edit" }}
      <li class="active">编辑栏目</li>
      {{else}}
      <li class="active">添加栏目</li>
      {{end}}
    </ol>
    
    <form method="post" {{if eq .Action "edit" }} action="/admin/category/edit/{{.Entity.Id}}" {{else}} action="/admin/category/add" {{end}} name="categoryForm" id="categoryForm">
    	<ul class="nav nav-tabs" role="tablist">
          <li class="nav-item">
            <a class="nav-link active" href="#base" role="tab" data-toggle="tab">栏目基本设置</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#detail" role="tab" data-toggle="tab">栏目详细设置</a>
            <input type="hidden" name="action" id="action" value="{{.Action}}" />
            <input type="hidden" name="Id" id="Id" value="{{.Entity.Id}}" />
          </li>
        </ul>
        <div class="tab-content">
          <div role="tabpanel" class="tab-pane active p-t-md" id="base">
          	<div class="form-group row">
            	<label for="Title" class="col-sm-2 form-control-label text-right">栏目名称：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="Title" name="Title" placeholder="请输入栏目名称" value="{{.Entity.Title}}" required>
                </div>
            </div>
            <div class="form-group row">
            	<label for="Pid" class="col-sm-2 form-control-label text-right">上级栏目：</label>
                <div class="col-sm-10">
                  <select class="form-control" id="Pid" name="Pid">
                      <option value="0">一级栏目</option>
                      {{range .PCategories}}
                      	<option value="{{.Id}}" {{if eq $.Entity.Pid .Id}} selected="selected" {{end}}>{{.Title}}</option>
                      {{end}}
                  </select>
                </div>
            </div>
          	<div class="form-group row">
            	<label class="col-sm-2 text-right">文章列表页：</label>
                <div class="col-sm-10">
                	<div class="checkbox">
                        <label>
                          <input type="checkbox" value="1" name="IsList" id="IsList" {{if eq .Entity.IsList 1}} checked {{end}} > 文章列表页 （勾选择该栏目为文章列表）
                        </label>
                    </div>
                </div>
            </div>
            <div class="form-group row">
            	<label class="col-sm-2 text-right">隐藏栏目：</label>
                <div class="col-sm-10">
                	<div class="checkbox">
                        <label>
                          <input type="checkbox" value="1" name="IsHide" id="IsHide" {{if eq .Entity.IsHide 1}} checked {{end}} > 隐藏 （勾选择该栏目为隐藏，前台无法访问）
                        </label>
                    </div>
                </div>
            </div>
            <div class="form-group row">
            	<label for="PageSize" class="col-sm-2 form-control-label text-right">每页文章数：</label>
                <div class="col-sm-10">
                  <input type="number" class="form-control" id="PageSize" name="PageSize" placeholder=""  value="{{.Entity.PageSize}}"  required>
                </div>
            </div>
            <div class="form-group row">
            	<label for="Rank" class="col-sm-2 form-control-label text-right">排序：</label>
                <div class="col-sm-10">
                  <input type="number" class="form-control" id="Rank" name="Rank" placeholder="越小则排得越前"  value="{{.Entity.Rank}}" required>
                </div>
            </div>
            <div class="form-group row">
            	<label for="DetailTemplateFile" class="col-sm-2 form-control-label text-right">栏目模板：</label>
                <div class="col-sm-10">
                  <select class="form-control" id="TemplateFile" name="TemplateFile">
                  	{{range .TemplatesList}}
                      <option value="{{.}}" {{if eq . $.Entity.TemplateFile}} selected="selected" {{end}}>{{.}}</option>
                    {{end}}
                  </select>
                </div>
            </div>
            <div class="form-group row">
            	<label for="TemplateFile" class="col-sm-2 form-control-label text-right">文章详情模板：</label>

                <div class="col-sm-10">
                  <select class="form-control" id="DetailTemplateFile" name="DetailTemplateFile">
                    {{range .TemplatesList}}
                      <option value="{{.}}" value="{{.}}" {{if eq . $.Entity.DetailTemplateFile}} selected="selected" {{end}}>{{.}}</option>
                    {{end}}
                  </select>
                </div>
            </div>
            <div class="form-group row">
            	<label for="LinkUrl" class="col-sm-2 form-control-label text-right">跳转链接：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="LinkUrl" name="LinkUrl"  value="{{.Entity.LinkUrl}}" placeholder="填写则点击该栏目自动跳转到该地址">
                </div>
            </div>

          </div>
          <div role="tabpanel" class="tab-pane p-t-md" id="detail">
            <div class="form-group row">
            	<label for="Keyword" class="col-sm-2 form-control-label text-right">关键词：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="Keyword" name="Keyword" value="{{.Entity.Keyword}}" placeholder="多个关键词用,(英文逗号)隔开,优化搜索引擎收录">
                </div>
            </div>
            <div class="form-group row">
            	<label for="Description" class="col-sm-2 form-control-label text-right">描述：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="Description" name="Description" value="{{.Entity.Description}}" placeholder="优化搜索引擎收录">
                </div>
            </div>
            <div class="form-group row">
            	<label for="Pic" class="col-sm-2 form-control-label text-right">Banner图片：</label>
                <div class="col-sm-10 input-group">
                  <input type="text" class="form-control upload-path" id="Pic" name="Pic" placeholder="请选择一张栏目banner图片" value="{{.Entity.Pic}}">
                  <span class="input-group-addon my-upload-span upload-img" id="basic-addon2"></span>
                </div>
            </div>
            <div class="form-group row">
            	<label for="Content" class="col-sm-2 form-control-label text-right">栏目介绍：</label>
                <div class="col-sm-10">
                  <textarea class="form-control" id="Content" name="Content" rows="3" style="width:99.9%">{{.Entity.Content}}</textarea>
                </div>
                <script>
					$(document).ready(function () {var editor = new UE.ui.Editor({ initialFrameHeight: 200});editor.render('Content')})
				</script>
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
        DoPost('categoryForm');
    });
	</script>
</div>