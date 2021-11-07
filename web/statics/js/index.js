/*
 * @Description: 
 * @Author: gphper
 * @Date: 2021-11-01 22:33:22
 */
var formloading = false;
$('#addPreKeyForm').submit(function () {
    
    $(this).find('.help-block').remove();
    $(this).find('.has-error').removeClass('has-error');

    if (formloading) {
        return false;
    }

    var loading = layer.msg('系统处理中...', {icon: 16, shade: 0, time: 0});
    form = $(this).serializeArray();
    $.ajax({
        type: $(this).prop('method'),
        url: $(this).prop('action'),
        data: $(this).serialize(),
        dataType: 'json',
        beforeSend: function () {
            formloading = true;
        },
        success: function (data) {
            if (data.status) {
                addItem($("button[tag='"+form[1].value+"']"),form[0].value,form[2].value);
                layer.msg(data.msg, {
                    icon: 6, scrollbar: false, time: 1000, shade: [0.3, '#393D49'], end: function () {
                      $('#addPreKeyModal').modal('hide');
                    }
                });
            } else {
                layer.msg(data.msg, {icon: 5, scrollbar: false, time: 2000, shade: [0.3, '#393D49']});
            }
        },
        complete: function () {
            formloading = false;
            layer.close(loading);
        }
    });
    return false;
});


var formloading = false;
$('#addKeyForm').submit(function () {
    
    $(this).find('.help-block').remove();
    $(this).find('.has-error').removeClass('has-error');

    if (formloading) {
        return false;
    }

    var loading = layer.msg('系统处理中...', {icon: 16, shade: 0, time: 0});
    form = $(this).serializeArray();
    $.ajax({
        type: $(this).prop('method'),
        url: $(this).prop('action'),
        data: $(this).serialize(),
        dataType: 'json',
        beforeSend: function () {
            formloading = true;
        },
        success: function (data) {
            if (data.status) {
                addTopItem($("#keys"),0,form[0].value,form[1].value);
                layer.msg(data.msg, {
                    icon: 6, scrollbar: false, time: 1000, shade: [0.3, '#393D49'], end: function () {
                      $('#addKeyModal').modal('hide');
                    }
                });
            } else {
                layer.msg(data.msg, {icon: 5, scrollbar: false, time: 2000, shade: [0.3, '#393D49']});
            }
        },
        complete: function () {
            formloading = false;
            layer.close(loading);
        }
    });
    return false;
});


var treeData = Object;
var formloading = false;
$('.search_key').submit(function () {
    $(this).find('.help-block').remove();
    $(this).find('.has-error').removeClass('has-error');

    if (formloading) {
        return false;
    }

    var loading = layer.msg('系统处理中...', {icon: 16, shade: 0, time: 0});
    $.ajax({
        type: $(this).prop('method'),
        url: $(this).prop('action'),
        data: $(this).serialize(),
        dataType: 'json',
        beforeSend: function () {
            formloading = true;
        },
        success: function (data){
            if (data.status) {
              treeData = data.data;
              showTree("keys",data.data)
            } else {
                layer.msg(data.data, {icon: 5, scrollbar: false, time: 2000, shade: [0.3, '#393D49']});
            }
        },
        error: function (jqXHR, textStatus, errorThrown) {
            
        },
        complete: function () {
            formloading = false;
            layer.close(loading);
        }
    });
    return false;
});



$("#exampleFormControlSelect1").change(function(){
if($(this).val() == "zset"){
  $("#score").show();
  $("#id").hide();
  $("#field").hide();
}else if($(this).val() == "hash"){
  $("#score").hide();
  $("#id").hide();
  $("#field").show();
}else if($(this).val() == "stream"){
  $("#id").show();
  $("#field").hide();
  $("#score").hide();
}else{
  $("#score").hide();
  $("#id").hide();
  $("#field").hide();
}
})

$("#exampleFormControlSelect1Pre").change(function(){
if($(this).val() == "zset"){
  $("#scorePre").show();
  $("#idPre").hide();
  $("#fieldPre").hide();
}else if($(this).val() == "hash"){
  $("#scorePre").hide();
  $("#idPre").hide();
  $("#fieldPre").show();
}else if($(this).val() == "stream"){
  $("#idPre").show();
  $("#fieldPre").hide();
  $("#scorePre").hide();
}else{
  $("#scorePre").hide();
  $("#idPre").hide();
  $("#fieldPre").hide();
}
})



function addTopItem(that,layer,keys,type){
    arr = keys.split(":");
    var idDiv = "";
    
    for (i = 0;i < arr.length;i++){
        layer ++;
        if(i == 0){
          idDiv = arr[i];
          if($("#"+idDiv).length == 0){
            //文件夹不存在
            that.append(temp(sec2,{name:arr[i],id:idDiv,item:"",pleft:layer*20,tag:arr[i],layer:layer}));
          }
        }else if(i == arr.length - 1){
          pre = idDiv.replaceAll(":",1)
          idDiv = idDiv+":"+arr[i];
          ids = idDiv.replaceAll(":",1)
          //钥匙
          $("#"+pre).append(temp(sec1,{name:arr[i],pleft:layer*20,tag:idDiv,keytype:type}));
        }else{
          pre = idDiv
          idDiv = idDiv+":"+arr[i];
          ids = idDiv.replaceAll(":",1)
          if($("#"+ids).length == 0){
            //文件夹不存在
            $("#"+pre).append(temp(sec2,{name:arr[i],id:ids,item:"",pleft:layer*20,tag:arr[i],layer:layer}));
          }
        }
    }
    loadHandleFunc()
    event.stopPropagation();
}


function addItem(that,key,type){

    var layer = that.attr('layer')
    var id = that.attr('tag')
    var str = key;
    arr = str.split(":");
    var pre = id.replaceAll(":","1");
    for (i = 0;i < arr.length;i++){
        layer ++;
        id = id + ":" + arr[i];
        idDiv = id.replaceAll(":","1")
        if(i != arr.length-1){
            //文件夹
            if($("#"+idDiv).length == 0){
                //文件夹不存在
                $("#"+pre).append(temp(sec2,{name:arr[i],id:idDiv,item:"",pleft:layer*20,tag:id,layer:layer}));
            }
        }else{
                //钥匙
                $("#"+pre).append(temp(sec1,{name:arr[i],pleft:layer*20,tag:id,keytype:type}));
        }
            pre = idDiv
    }
    loadHandleFunc()
    event.stopPropagation();
}

function loadHandleFunc(){
  $(".father-node").click(function(){
      if($(this).children("div").children('i').attr('class') == "bi bi-folder2-open"){
          $(this).children("div").children('i').attr('class',"bi bi-folder")
      }else{
          $(this).children("div").children('i').attr('class',"bi bi-folder2-open")
      }
  })

  $(".addbtn").click(function (event){
    $("#prekey").html($(this).attr('tag')+":");
    $("input[name=prekey]").val($(this).attr('tag'));
    $("input[name=key]").val("")
    $("#addPreKeyModal").modal("show");
    event.stopPropagation();
  })

  $(".seebtn").click(function (event){
      document.getElementById("iframepage").src='/show_key/'+'?key='+$(this).attr("tag");
      event.stopPropagation();
  })

  $(".delbtn").click(function (event){
      var loading = layer.msg('系统处理中...', {icon: 16, shade: 0, time: 0});
      var jq = $(this);
      $.ajax({
          type: "post",
          url: "/del_key",
          data: {"key":$(this).attr("tag"),"type":$(this).attr("key-type")},
          dataType: 'json',
          beforeSend: function () {
              formloading = true;
          },
          success: function (data){
              if (data.status) {
                var pre = "";
                var that = jq.parent();
                var parent = jq.parent().parent();
                for(i = 0;i<20;i++){
                
                  if(parent.children("li").length == 1){
                    if(parent.attr("id") == "keys"){
                      parent.remove()
                      break;
                    }
                    pre = parent
                  }else{
                    that.remove()
                    if(that.attr('href')){
                      $(that.attr('href')).remove()
                    }
                    if(pre){
                      $("li[href='#"+pre.attr('id')+"']").remove()
                      if(pre.attr('id') != "main"){
                        pre.remove()
                      }else{
                        pre.children().first().remove()
                      }
                      
                    }
                    break;
                  }
                  parent = parent.parent()
                }

              } else {
                  layer.msg(data.data, {icon: 5, scrollbar: false, time: 2000, shade: [0.3, '#393D49']});
              }
          },
          error: function (jqXHR, textStatus, errorThrown) {
              
          },
          complete: function () {
              formloading = false;
              layer.close(loading);
          }
      });
      
      event.stopPropagation();
  })
}
