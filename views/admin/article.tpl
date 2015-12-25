<div class="row">
    <ol class="breadcrumb m-b-sm">
      <li><a href="/admin">后台首页</a></li>
      <li><a href="/admin/article">文章</a></li>
      <li class="active">文章管理</li>
    </ol>
	<div class="m-b-md">
	<div class="btn-group" role="group" aria-label="工具栏">
        <a class="btn btn-primary" href="/admin/article/add" role="button"><span class="fa fa-folder-o"></span> 添加</a>
	</div>
    	<form class="form-inline navbar-form pull-right" action="?">
			<input class="form-control" type="text" name="key" id="key" value="{{.Key}}" placeholder="文章关键字">
			<button class="btn btn-success" type="submit">搜索</button>
		</form>
    </div>
    <div class="table-responsive">
        <table class="table table-striped">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>文章标题</th>
                  <th>排序</th>
                  <th>添加时间</th>
                  <th>文章属性</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
              {{range $index,$v := .List}}
                <tr id="row_{{$v.Id}}">
                  <td>{{$v.Id}}</td>
                  <td>{{$v.Title}}</td>
                  <td>{{$v.Rank}}</td>
                  <td>{{date $v.Created "Y-m-d H:i:s"}}</td>
                  <td>
                  {{if eq $v.IsNew 1}}
                  <span class="label label-danger" title="推荐">新</span>
                  {{end}}
                  {{if eq $v.IsRecommend 1}}
                  <span class="label label-info" title="推荐">荐</span>
                  {{end}}
                  {{if eq $v.IsHide 1}}
                  <span class="label label-default" title="隐藏">隐</span>
                  {{end}}
                  </td>
                  <td><a class="btn btn-success btn-sm" href="/admin/article/edit/{{$v.Id}}" role="button"><span class="fa fa-pencil-square-o"></span> 修改</a> <a class="btn btn btn-danger btn-sm" href="javascript:;" onClick="doDel('/admin/article/del',{{$v.Id}})" role="button"><span class="fa fa-times"></span> 删除</a></td>
                </tr>
                {{end}}
              </tbody>
        </table>
        <nav class="">
        {{template "admin/pages.tpl" .}}
       </nav>
    </div>
</div>