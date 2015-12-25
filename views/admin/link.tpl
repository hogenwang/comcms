<div class="row">
    <ol class="breadcrumb m-b-sm">
      <li><a href="/admin">后台首页</a></li>
      <li><a href="/admin/link">其他</a></li>
      <li class="active">友情链接管理</li>
    </ol>

	<div class="btn-group m-b-md" role="group" aria-label="工具栏">
        <a class="btn btn-primary" href="/admin/link/add" role="button"><span class="fa fa-folder-o"></span> 添加</a>
	</div>
    <div class="table-responsive">
        <table class="table table-striped">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>站点名称</th>
                  <th>排序</th>
                  <th>站点地址</th>
                  <th>站点Logo</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
              {{range $index,$v := .List}}
                <tr id="row_{{$v.Id}}">
                  <td>{{$v.Id}}</td>
                  <td>{{$v.Title}}</td>
                  <td>{{$v.Rank}}</td>
                  <td>{{$v.Url}}</td>
                  <td>{{$v.Logo}}</td>
                  <td><a class="btn btn-success btn-sm" href="/admin/link/edit/{{$v.Id}}" role="button"><span class="fa fa-pencil-square-o"></span> 修改</a> <a class="btn btn btn-danger btn-sm" href="javascript:;" onClick="doDel('/admin/link/del',{{$v.Id}})" role="button"><span class="fa fa-times"></span> 删除</a></td>
                </tr>
                {{end}}
              </tbody>
        </table>
        <nav class="">
        {{template "admin/pages.tpl" .}}
       </nav>
    </div>
</div>