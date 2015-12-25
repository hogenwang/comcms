//后台所有异步提交 formId 不需要加上#
//依赖jquery jquery.form jquery.validate jquery.metadata
//验证例子：<input type="text" name="email" class="{validate:{ required:true,email:true }}" />
function DoPost(formId){
	//提交地址
	var $myform = $("#"+formId);
	var url = $myform.attr("action");
	var btn;
	if($myform.find("input[type='submit']").length >0) {
		btn = $myform.find("input[type='submit']");
	} else if ($myform.find("button[type='submit']").length >0) {
		btn = $myform.find("button[type='submit']");
	}
	//alert(queryString);
	var v = $myform.validate({
		meta:"validate",
		submitHandler: function(form) {
			//loadding效果
			var loading = layer.load(0, {
				shade: [0.2,'#000'] //0.1透明度的背景
			});
			
			btn.attr("disabled","disabled");
			//执行loadding 并且ajax提交
			var queryString = $myform.formSerialize();
			$.ajax({
				type:"POST",
				url:url,
				data:queryString,
				dataType:"JSON",
				success: function(data){
					if(data.Status == "success"){
						layer.close(loading);
						//alert(data.Message);
						//window.location.href= data.ReturnUrl;
						layer.msg(data.Message, {
							time: 1000 //2秒关闭（如果不配置，默认是3秒）
						}, function(){
							window.location.href= data.ReturnUrl;
						});
					} else {
						layer.close(loading);
						layer.alert(data.Message, {icon: 2});
						btn.removeAttr("disabled");
					}
				},
				error:function(){
					layer.close(loading);
					layer.alert('执行错误，请联系管理员！', {icon: 2});
					btn.removeAttr("disabled");
				}
			});
		}
	});
	return false;
}

//后台异步删除
function doDel(url,id){
	var loadding ="";//loadding效果
	layer.confirm('确认要删除当前记录？！', {
		icon: 3, 
		title:'系统提示',
		btn: ['确定','取消 '] //按钮
	}, function(){
		var loading = layer.load(0, {
				shade: [0.2,'#000'] //0.1透明度的背景
			});
			$.ajax({
				type:"POST",
				url:url,
				data:'id='+id,
				dataType:"JSON",
				success: function(data){
					if(data.Status == "success"){
						layer.close(loading);
						layer.msg(data.Message, {
							time: 1000 //2秒关闭（如果不配置，默认是3秒）
						}, function(){
							$("#row_"+id).fadeOut(300).remove();
							//window.location.href= data.ReturnUrl;
						});
					} else {
						layer.close(loading);
						layer.alert(data.Message, {icon: 2});

					}
				},
				error:function(){
					layer.close(loading);
					layer.alert('执行错误，请联系管理员！', {icon: 2});
				}
			});
	}, function(){
	});

}