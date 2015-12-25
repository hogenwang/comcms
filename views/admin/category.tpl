<script src="/static/js/treegrid/jquery.treegrid.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/js/treegrid/css/jquery.treegrid.css">
<div class="row">
    <ol class="breadcrumb m-b-sm">
      <li><a href="/admin">后台首页</a></li>
      <li class="active">文章栏目管理</li>
    </ol>

	<div class="btn-group m-b-md" role="group" aria-label="工具栏">
        <a class="btn btn-primary" href="/admin/category/add" role="button"><span class="fa fa-folder-o"></span> 添加</a>
        <a class="btn btn-success" href="#" role="button"><span class="fa fa-pencil-square-o"></span> 修改</a>
        <a class="btn btn btn-danger" href="#" role="button"><span class="fa fa-times"></span> 删除</a>
	</div>
    <div class="table-responsive">
        <table class="table table-striped tree">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>栏目名称</th>
                  <th>排序</th>
                  <th>文章数量</th>
                  <th>隐藏</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
              {{range $index,$v := .List}}
                <tr class="treegrid-{{$v.Id}} {{if gt $v.Pid 0}} treegrid-parent-{{$v.Pid}} {{end}}" id="row_{{$v.Id}}">
                  <td>{{$v.Id}}</td>
                  <td>{{$v.Title}}</td>
                  <td>{{$v.Counts}}</td>
                  <td>{{$v.Rank}}</td>
                  <td>
                      {{if eq $v.IsHide 1}}
                      <span class="label label-default" title="隐藏">隐</span>
                      {{end}}
                  </td>
                  <td><a class="btn btn-success btn-sm" href="/admin/category/edit/{{$v.Id}}" role="button"><span class="fa fa-pencil-square-o"></span> 修改</a> <a class="btn btn btn-danger btn-sm"  href="javascript:;" onClick="doDel('/admin/category/del',{{$v.Id}})" role="button"><span class="fa fa-times"></span> 删除</a></td>
                </tr>
                {{end}}
              </tbody>
        </table>
   </div>
</div>
<script type="text/javascript">
  $('.tree').treegrid();
</script>