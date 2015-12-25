<div class="row">
    <ol class="breadcrumb m-b-sm">
      <li><a href="/admin">后台首页</a></li>
      <li><a href="/admin/link">其他</a></li>
      <li class="active">广告管理</li>
    </ol>

	<div class="btn-group m-b-md" role="group" aria-label="工具栏">
        <a class="btn btn-primary" href="/admin/ads/add" role="button"><span class="fa fa-folder-o"></span> 添加</a>
	</div>
    <div class="table-responsive">
        <table class="table table-striped">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>广告标题</th>
                  <th>排序</th>
                  <th>广告类型</th>
                  <th>是否隐藏</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
              {{range $index,$v := .List}}
                <tr id="row_{{$v.Id}}">
                  <td>{{$v.Id}}</td>
                  <td>{{$v.Title}}</td>
                  <td>{{$v.Rank}}</td>
                  <td>
                  {{if eq $v.Tid 0}}
                  <span class="label label-info" title="HTML代码">HTML代码</span>
                  {{end}}
                  {{if eq $v.Tid 1}}
                  <span class="label label-info" title="文字广告">文字广告</span>
                  {{end}}
                  {{if eq $v.Tid 2}}
                  <span class="label label-info" title="图片广告">图片广告</span>
                  {{end}}
                  {{if eq $v.Tid 3}}
                  <span class="label label-info" title="Flash广告">Flash广告</span>
                  {{end}}
                  {{if eq $v.Tid 4}}
                  <span class="label label-info" title="幻灯片广告">幻灯片广告</span>
                  {{end}}
                  </td>
                  <td>
                  {{if eq $v.IsHide 1}}
                  <span class="label label-default" title="隐藏">隐</span>
                  {{end}}
                  </td>
                  <td><a class="btn btn-success btn-sm" href="/admin/ads/edit/{{$v.Id}}" role="button"><span class="fa fa-pencil-square-o"></span> 修改</a> <a class="btn btn btn-danger btn-sm" href="javascript:;" onClick="doDel('/admin/ads/del',{{$v.Id}})" role="button"><span class="fa fa-times"></span> 删除</a></td>
                </tr>
                {{end}}
              </tbody>
        </table>
        <nav class="">
        {{template "admin/pages.tpl" .}}
       </nav>
    </div>
</div>