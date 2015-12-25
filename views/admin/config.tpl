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
      <li><a href="/admin/config">系统</a></li>
      <li class="active">系统配置</li>
    </ol>
    
    <form method="post" action="/admin/config" name="configForm" id="configForm">
    	<ul class="nav nav-tabs" role="tablist">
          <li class="nav-item">
            <a class="nav-link active" href="#base" role="tab" data-toggle="tab">站点设置</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#seo" role="tab" data-toggle="tab">SEO设置</a>
         </li>
          <li class="nav-item">
            <a class="nav-link" href="#skin" role="tab" data-toggle="tab">模板设置</a>
          </li>
        </ul>
        <div class="tab-content">
          <div role="tabpanel" class="tab-pane active p-t-md" id="base">
          	<div class="form-group row">
            	<label for="Title" class="col-sm-2 form-control-label text-right">站点名称：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="SiteName" name="SiteName" placeholder="请输入站点名称" value="{{.Cfg.SiteName}}" required>
                </div>
            </div>
			<div class="form-group row">
            	<label for="Title" class="col-sm-2 form-control-label text-right">站点URL：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="SiteURL" name="SiteURL" placeholder="请输入站点URL" value="{{.Cfg.SiteURL}}" >
                </div>
            </div>
            <div class="form-group row">
            	<label for="Title" class="col-sm-2 form-control-label text-right">站点Logo：</label>
                <div class="col-sm-10 input-group">
                  <input type="text" class="form-control upload-path" id="SiteLogo" name="SiteLogo" placeholder="请输入站点Logo" value="{{.Cfg.SiteLogo}}" >
                  <span class="input-group-addon my-upload-span upload-img" id="basic-addon2"></span>
                </div>
                <script>
				var BASE_URL='/static/js/webuploader/';
				// 初始化Web Uploader
				var uploader = WebUploader.create({
				
					// 选完文件后，是否自动上传。
					auto: true,
				
					// swf文件路径
					swf: BASE_URL + 'Uploader.swf',
				
					// 文件接收服务端。
					server: '/admin/fileupload',
					fileNumLimit:1,
				
					// 选择文件的按钮。可选。
					// 内部根据当前运行是创建，可能是input元素，也可能是flash.
					pick: '#filePicker',
				
					// 只允许选择图片文件。
					accept: {
						title: 'Images',
						extensions: 'gif,jpg,jpeg,bmp,png',
						mimeTypes: 'image/*'
					}
				});
				</script>
            </div>

            <div class="form-group row">
            	<label for="Title" class="col-sm-2 form-control-label text-right">管理员Email：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="SiteEmail" name="SiteEmail" placeholder="请输入站点管理员Email" value="{{.Cfg.SiteEmail}}" >
                </div>
            </div>
            <div class="form-group row">
            	<label for="Title" class="col-sm-2 form-control-label text-right">网站备案号：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="ICP" name="ICP" placeholder="请输入网站ICP备案号" value="{{.Cfg.ICP}}" >
                </div>
            </div>
            <div class="form-group row">
            	<label for="Content" class="col-sm-2 form-control-label text-right">版权所有：</label>
                <div class="col-sm-10">
                  <textarea class="form-control" id="Copyright" name="Copyright" rows="3" style="width:99.9%">{{.Cfg.Copyright}}</textarea>
                </div>
                <script>
					$(document).ready(function () {var editor = new UE.ui.Editor({ initialFrameHeight: 200});editor.render('Copyright')})
				</script>
            </div>
            <div class="form-group row">
            	<label for="Content" class="col-sm-2 form-control-label text-right">QQ客服：</label>
                <div class="col-sm-10">
                  <textarea class="form-control" id="OnlineQQ" name="OnlineQQ" rows="3" placeholder="每一行输入一个QQ号码,QQ号码跟昵称，使用“|”分割，如：123456|张小姐，个请键入回车输入" >{{.Cfg.OnlineQQ}}</textarea>
                </div>
            </div>
          </div>
          <div role="tabpanel" class="tab-pane p-t-md" id="seo">
          	<div class="form-group row">
            	<label for="Keyword" class="col-sm-2 form-control-label text-right">首页Title：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="SiteTitle" name="SiteTitle" value="{{.Cfg.SiteTitle}}" placeholder="自定义首页标题，如果没填写，则按网站名称显示">
                </div>
            </div>
            <div class="form-group row">
            	<label for="Keyword" class="col-sm-2 form-control-label text-right">关键词：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="Keyword" name="Keyword" value="{{.Cfg.Keyword}}" placeholder="多个关键词用,(英文逗号)隔开,优化搜索引擎收录">
                </div>
            </div>
            <div class="form-group row">
            	<label for="Description" class="col-sm-2 form-control-label text-right">描述：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="Description" name="Description" value="{{.Cfg.Description}}" placeholder="优化搜索引擎收录">
                </div>
            </div>
          </div>
          <div role="tabpanel" class="tab-pane p-t-md" id="skin">
          	<div class="form-group row">
            	<label for="Keyword" class="col-sm-2 form-control-label text-right">网站风格：</label>
                <div class="col-sm-10">
                  <select class="form-control" id="Skin" name="Skin">
                  	<option value="default">default</option>
                  </select>
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
        DoPost('configForm');
    });
	</script>
</div>