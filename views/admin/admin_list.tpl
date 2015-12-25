<div class="row">
    <ol class="breadcrumb m-b-sm">
      <li><a href="/admin">后台首页</a></li>
      <li><a href="/admin/admin/list">用户</a></li>
      <li class="active">管理员列表</li>
    </ol>

	<div class="btn-group m-b-md" role="group" aria-label="工具栏">
        <a class="btn btn-primary" href="/admin/admin/add" role="button"><span class="fa fa-folder-o"></span> 添加</a>
	</div>
    <div class="table-responsive">
        <table class="table table-striped">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>用户名</th>
                  <th>最后登录时间</th>
                  <th>最后登录IP</th>
                  <th>登录次数</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
              {{range $index,$v := .List}}
                <tr id="row_{{$v.Id}}">
                  <td>{{$v.Id}}</td>
                  <td>{{$v.UserName}}</td>
                  <td>{{date $v.ThisLoginTime "Y-m-d H:i:s"}}</td>
                  <td>{{$v.ThisLoginIP}}</td>
                  <td>{{$v.LoginCount}}</td>
                  <td><a class="btn btn-success btn-sm" href="/admin/admin/edit/{{$v.Id}}" role="button"><span class="fa fa-pencil-square-o"></span> 修改</a> <a class="btn btn btn-danger btn-sm" href="javascript:;" onClick="doDel('/admin/admin/del',{{$v.Id}})" role="button"><span class="fa fa-times"></span> 删除</a></td>
                </tr>
                {{end}}
              </tbody>
        </table>
        <nav class="">
        {{template "admin/pages.tpl" .}}
       </nav>
    </div>
</div>