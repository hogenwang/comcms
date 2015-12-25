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
      <li class="active">修改广告详情</li>
      {{else}}
      <li class="active">添加广告详情</li>
      {{end}}
    </ol>
    
    <form method="post" {{if eq .Action "edit" }} action="/admin/ads/edit/{{.Entity.Id}}" {{else}} action="/admin/ads/add" {{end}} name="linkForm" id="linkForm">
    	<ul class="nav nav-tabs" role="tablist">
          <li class="nav-item">
            <a class="nav-link active" href="#base" role="tab" data-toggle="tab">广告详情</a>
            <input type="hidden" name="action" id="action" value="{{.Action}}" />
            <input type="hidden" name="Id" id="Id" value="{{.Entity.Id}}" />
          </li>
        </ul>
        <div class="tab-content">
          <div role="tabpanel" class="tab-pane active p-t-md" id="base">
          	<div class="form-group row">
            	<label for="Title" class="col-sm-2 form-control-label text-right">广告标题：</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="Title" name="Title" placeholder="请输入广告标题" value="{{.Entity.Title}}" required>
                </div>
            </div>
            <div class="form-group row">
            	<label for="Rank" class="col-sm-2  form-control-label text-right">排序：</label>
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
            <div class="form-group row">
            	<label for="Pid" class="col-sm-2 form-control-label text-right">广告类型：</label>
                <div class="col-sm-10">
                  <select class="form-control" id="Tid" name="Tid" onchange="ChangeType()" >
                      	<option value="0" {{if eq $.Entity.Tid 0}} selected="selected" {{end}}>HTML代码</option>
                        <option value="1" {{if eq $.Entity.Tid 1}} selected="selected" {{end}}>文字广告</option>
                        <option value="2" {{if eq $.Entity.Tid 2}} selected="selected" {{end}}>图片广告</option>
                        <option value="3" {{if eq $.Entity.Tid 3}} selected="selected" {{end}}>Flash广告</option>
                        <option value="4" {{if eq $.Entity.Tid 4}} selected="selected" {{end}}>幻灯片广告</option>
                  </select>
                </div>
            </div>
            <div id="adstype_0">
                <div class="form-group row">
                    <label for="Content" class="col-sm-2 form-control-label text-right">HTML代码：</label>
                    <div class="col-sm-10">
                      <textarea class="form-control" id="txtScript" name="txtScript" rows="3" placeholder="请输入广告HTML代码" ></textarea>
                    </div>
                </div>
            </div>
            <div id="adstype_1" class="hide">
            	<div class="form-group row">
                    <label for="txt_Txt" class="col-sm-2 form-control-label text-right">文字内容(必填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="txt_Txt" name="txt_Txt" value="" placeholder="请输入文字广告的显示内容">
                    </div>
            	</div>
                <div class="form-group row">
                    <label for="txt_Link" class="col-sm-2 form-control-label text-right">文字链接(必填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="txt_Link" name="txt_Link" value="" placeholder="请输入文字广告指向的 URL 链接地址">
                    </div>
            	</div>
                <div class="form-group row">
                    <label for="txt_Style" class="col-sm-2 form-control-label text-right">文字样式(选填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="txt_Style" name="txt_Style" value="" placeholder="请输入文字广告的内容显示字体样式，即style内容">
                    </div>
            	</div>
            </div>
            <div id="adstype_2" class="hide">
            	<div class="form-group row">
                    <label for="img_Img" class="col-sm-2 form-control-label text-right">图片地址(必填)：</label>
                    <div class="col-sm-10 input-group">
                      <input type="text" class="form-control upload-path" id="img_Img" name="img_Img" value="" placeholder="请输入图片广告的图片调用地址">
                      <span class="input-group-addon my-upload-span upload-img" id="basic-addon2"></span>
                    </div>
            	</div>
                <div class="form-group row">
                    <label for="img_Link" class="col-sm-2 form-control-label text-right">图片链接(必填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="img_Link" name="img_Link" value="" placeholder="请输入图片广告图片链接地址">
                    </div>
            	</div>
                <div class="form-group row">
                    <label for="img_Width" class="col-sm-2 form-control-label text-right">图片宽度(选填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="img_Width" name="img_Width" value="" placeholder="请输入图片广告的宽度，单位为像素">
                    </div>
            	</div>
                <div class="form-group row">
                    <label for="img_Height" class="col-sm-2 form-control-label text-right">图片高度(选填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="img_Height" name="img_Height" value="" placeholder="请输入图片广告的高度，单位为像素">
                    </div>
            	</div>
                <div class="form-group row">
                    <label for="img_Alt" class="col-sm-2 form-control-label text-right">图片替换文字(选填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="img_Alt" name="img_Alt" value="" placeholder="请输入图片广告的鼠标悬停文字信息">
                    </div>
            	</div>
            </div>
            <div id="adstype_3" class="hide">
            	<div class="form-group row">
                    <label for="flash_Swf" class="col-sm-2 form-control-label text-right">Flash 地址(必填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="flash_Swf" name="flash_Swf" value="" placeholder="请输入 Flash 广告的调用地址">
                    </div>
            	</div>
                <div class="form-group row">
                    <label for="flash_Width" class="col-sm-2 form-control-label text-right">Flash 宽度(必填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="flash_Width" name="flash_Width" value="" placeholder="请输入 Flash 广告的宽度，单位为像素">
                    </div>
            	</div>
                <div class="form-group row">
                    <label for="flash_Height" class="col-sm-2 form-control-label text-right">Flash 高度(必填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="flash_Height" name="flash_Height" value="" placeholder="请输入 Flash 广告的高度，单位为像素">
                    </div>
            	</div>
            </div>
            <div id="adstype_4" class="hide">
                <div class="form-group row">
                    <label for="slide_Width" class="col-sm-2 form-control-label text-right">幻灯片宽度(必填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="slide_Width" name="slide_Width" value="" placeholder="请输入幻灯片广告的宽度，单位为像素">
                    </div>
            	</div>
                <div class="form-group row">
                    <label for="slide_Height" class="col-sm-2 form-control-label text-right">幻灯片高度(必填)：</label>
                    <div class="col-sm-10">
                      <input type="text" class="form-control" id="slide_Height" name="slide_Height" value="" placeholder="请输入幻灯片广告的高度，单位为像素">
                    </div>
            	</div>
                <div class="form-group row">
                <label for="slide_Height" class="col-sm-2 form-control-label text-right">幻灯片详情：</label>
                    <div class="col-sm-10">
                        <table class="table table-hover">
                          <thead>
                            <tr>
                              <th style="width:30%">图片（必填）</th>
                              <th>链接地址（必填）</th>
                              <th>图片标题</th>
                              <th style="width:80px">操作</th>
                            </tr>
                          </thead>
                          <tbody id="SildeImgs">
                            <tr>
                              <td><div class="input-group"><input type="text" class="form-control upload-path" id="slide_Img_0" name="slide_Img" value="" placeholder="请输入图片地址"><span class="input-group-addon my-upload-span upload-img" id="basic-addon2"></span></div></td>
                              <td><input type="text" class="form-control" id="slide_Link_0" name="slide_Link" value="" placeholder="请输入链接地址"></td>
                              <td><input type="text" class="form-control" id="slide_Alt_0" name="slide_Alt" value="" placeholder="请输入图片标题"></td>
                              <td><button type="button" class="btn btn-sm btn-success"  onClick="addSlideRow();"><i class="fa fa-plus-square"></i> 添加</button></td>
                            </tr>
                         </tbody>
                        </table>
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
		function ChangeType() {
	        var styles, key, val;
	        val = $('#Tid').val();
	        styles = new Array('0', '1', '2', '3', '4');
	        for (key in styles) {
				if(styles[key] == val){
					$("#adstype_"+styles[key]).show();
				} else {
					$("#adstype_"+styles[key]).hide();
				}
	        }
	    }
		
		var htmlText = '<tr id="slideTR_{0}">';
	    htmlText += '<td><div class="input-group"><input type="text" class="form-control upload-path" id="slide_Img_{0}" name="slide_Img" value="" placeholder="请输入图片地址"><span class="input-group-addon my-upload-span upload-img" id="basic-addon_{0}"></span></div></td>';
		htmlText +='<td><input type="text" class="form-control" id="slide_Link_{0}" name="slide_Link" value="" placeholder="请输入链接地址"></td>';
		htmlText +='<td><input type="text" class="form-control" id="slide_Alt_{0}" name="slide_Alt" value="" placeholder="请输入图片标题"></td>';
		htmlText +='<td><button type="button" class="btn btn-sm btn-danger"><i class="fa fa-minus-square" onclick="delmySlideRow({0})"></i> 删除</button></td></tr>';
	    var items = 1;

	    function addSlideRow() {
	        var lastitem = items - 1;
	        var newHtmlTr = htmlText.replace(/\{0\}/g, items);
	        $('#SildeImgs').append(newHtmlTr);
			$("#basic-addon_"+items).InitUploader({ filesize: "10240", sendurl: "/admin/webupload", swf: "/static/js/webuploader/uploader.swf", filetypes: "gif,jpg,png,bmp" });
	        items += 1;
	    }
	    function delmySlideRow(id) {
	        $("#slideTR_" + id).remove();
	    }
	</script>
    {{if eq .Action "edit" }}
    <script>
	$(function(){
		$("#Tid").val('{{.Entity.Tid}}');
		FillAds({{.Entity.Tid}},{{.Entity.Content}});
		ChangeType();
	});
	    //JS加载广告内容
    function FillAds(tid, content) {
        var json = eval(content);
        switch (tid) {
            case 0: //代码
                $('#txtScript').val(json.Content);
                break;
            case 1: //文字类型
                $('#txt_Txt').val(json.Txt);
                $('#txt_Link').val(json.Link);
                $('#txt_Style').val(json.Style);
                break;
            case 2: //图片类
                $('#img_Img').val(json.Img);
                $('#img_Link').val(json.Link);
                if(json.width !=0){
                    $('#img_Width').val(json.Width);
                }
                if(json.height !=0){
                    $('#img_Height').val(json.Height);
                }
                $('#img_Alt').val(json.Alt);
                break;
            case 3://Flash类型
                $('#flash_Swf').val(json.Swf);
                $('#flash_Height').val(json.Height);
                $('#flash_Width').val(json.Width);
                break;
            case 4://幻灯片
                for(i=0;i<json.length;i++)
				{
					if(i ==0)
					{
						$("#slide_Width").val(json[0].Width);
						$("#slide_Height").val(json[0].Height);
					}
					else
					{
						addSlideRow();//增加一行
					}
					$("#slide_Img_"+i).val(json[i].Img);
					$("#slide_Link_"+i).val(json[i].Link);
					$("#slide_Alt_"+i).val(json[i].Alt);
				}
                break;
        }
    }
	</script>
    {{end}}
</div>