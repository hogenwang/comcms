<div class="row">
    <ol class="breadcrumb m-b-sm">
      <li><a href="/admin">后台首页</a></li>
      <li><a href="/admin/link">其他</a></li>
      <li class="active">留言板管理</li>
    </ol>

	<div class="btn-group m-b-md" role="group" aria-label="工具栏">

	</div>
    <div class="table-responsive">
        <table class="table table-striped">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>姓名</th>
                  <th>电话</th>
                  <th>留言内容</th>
                  <th>留言时间</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
              {{range $index,$v := .List}}
                <tr id="row_{{$v.Id}}">
                  <td>{{$v.Id}}</td>
                  <td>{{$v.UserName}}</td>
                  <td>{{$v.Tel}}</td>
                  <td>{{$v.Title}}</td>
                  <td>{{date $v.Created "Y-m-d H:i:s"}}</td>
                  <td><a class="btn btn-success btn-sm" href="/admin/guestbook/view/{{$v.Id}}" role="button"><span class="fa fa-pencil-square-o"></span> 查看</a> <a class="btn btn btn-danger btn-sm" href="javascript:;" onClick="doDel('/admin/guestbook/del',{{$v.Id}})" role="button"><span class="fa fa-times"></span> 删除</a></td>
                </tr>
                {{end}}
              </tbody>
        </table>
        <nav class="">
        {{template "admin/pages.tpl" .}}
       </nav>
    </div>
</div>