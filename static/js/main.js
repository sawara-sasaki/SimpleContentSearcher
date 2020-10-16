var Search = function() {
  const data = {action:"search", parameters:[$("#search-word").val()]}
  request(data, (res)=>{
    if(res.data.length > 0) {
      showInfoMessage(res.data);
    } else {
      showDangerMessage("No Data.");
    }
  }, onerror);
}
var onError = function(e) {
  if (!!e.responseJSON) {
    console.log(e.responseJSON.message);
    showDangerMessage(e.responseJSON.message);
  } else {
    console.log(e.message);
    showDangerMessage(e.message);
  }
};
var request = function(data, callback, onerror) {
  $.ajax({
    type:          'POST',
    dataType:      'json',
    contentType:   'application/json',
    scriptCharset: 'utf-8',
    data:          JSON.stringify(data),
    url:           "./action"
  })
  .done(function(res) {
    callback(res);
  })
  .fail(function(e) {
    onerror(e);
  });
};
var showDangerMessage = function(str) {
  $("#message-danger").text(str);
  $("#message-danger-container").removeClass("hide");
}
var showInfoMessage = function(elmList) {
  elmList.forEach(elm => $("#message-info").append("<p>" + elm + "</p>"));
  $("#message-info-container").removeClass("hide");
}
var CloseMessage = function(elm) {
  $(elm).parent().addClass("hide");
}
$(function(){
  $("input"). keydown(function(e) {
    if ((e.which && e.which === 13) || (e.keyCode && e.keyCode === 13)) {
      Search();
      return false;
    } else if ((e.which && e.which === 27) || (e.keyCode && e.keyCode === 27)) {
      $('#main-form').submit();
      return false;
    } else {
      return true;
    }
  });
});
