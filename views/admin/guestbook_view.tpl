
<div class="row">
    <ol class="breadcrumb m-b-sm">
      <li><a href="/admin">后台首页</a></li>
      <li><a href="/admin/link">其他</a></li>

      <li class="active">查看留言板详情</li>
   </ol>
    
    <form method="post" name="guestbookForm" id="guestbookForm" onSubmit="return flase;">
    	<ul class="nav nav-tabs" role="tablist">
          <li class="nav-item">
            <a class="nav-link active" href="#base" role="tab" data-toggle="tab">查看留言板详情</a>
            <input type="hidden" name="action" id="action" value="{{.Action}}" />
          </li>
        </ul>
        <div class="tab-content">
          <div role="tabpanel" class="tab-pane active p-t-md" id="base">
          	<div class="form-group row">
            	<label for="UserName" class="col-sm-2 form-control-label text-right">姓名：</label>
                <div class="col-sm-10">
                  {{.Entity.UserName}}
                </div>
            </div>
            <div class="form-group row">
            	<label for="Tel" class="col-sm-2 form-control-label text-right">电话：</label>
                <div class="col-sm-10">
                  {{.Entity.Tel}}
                </div>
            </div>
            <div class="form-group row">
            	<label for="Email" class="col-sm-2 form-control-label text-right">邮箱：</label>
                <div class="col-sm-10">
                  {{.Entity.Email}}
                </div>
            </div>
            <div class="form-group row">
            	<label for="QQ" class="col-sm-2 form-control-label text-right">QQ：</label>
                <div class="col-sm-10">
                  {{.Entity.QQ}}
                </div>
            </div>
            <div class="form-group row">
            	<label for="Title" class="col-sm-2 form-control-label text-right">留言标题：</label>
                <div class="col-sm-10">
                  {{.Entity.Title}}
                </div>
            </div>
            <div class="form-group row">
            	<label for="IP" class="col-sm-2 form-control-label text-right">IP：</label>
                <div class="col-sm-10">
                  {{.Entity.IP}}
                </div>
            </div>
            <div class="form-group row">
            	<label for="Created" class="col-sm-2 form-control-label text-right">留言时间：</label>
                <div class="col-sm-10">
                  {{date .Entity.Created "Y-m-d H:i:s"}}
                </div>
            </div>
            <div class="form-group row">
            	<label for="Content" class="col-sm-2 form-control-label text-right">留言详情：</label>
                <div class="col-sm-10">
                  {{.Entity.Content}}
                </div>
            </div>
          </div>
		<div>
           <a href="javascript:;" onClick="window.history.go(-1)" class="btn btn-secondary"><i class="fa fa-undo"></i> 返回</a>
        </div>
        </div>
    </form>
</div>