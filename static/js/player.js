$(function(){
  var vid = ""
  var qs = window.location.search;
  var qo = new Object();
  if (qSearch.startsWith('?https://www.youtube.com/watch?v=')) {
    vid = qSearch.substring(33).split('&')[0];
  } else {
    if(qs){
      qs = qs.substring(1);
      var pa = qs.split('&');
      for (var i = 0; i < pa.length; i++) {
        var el = pa[i].split('=');
        var pn = decodeURIComponent(el[0]);
        var pv = decodeURIComponent(el[1]);
        qo[pn] = pv;
      }
    }
    if (!!qo["v"] && qo["v"].length > 0) {
      vid = qo["v"];
    }
  }
  const data = {action:"search", parameters:["https://www.youtube.com/watch?v="+vid]}
  request(data, (res)=>{
    if(res.data.length > 0) {
      showInfoMessage(res.data);
    } else {
      showDangerMessage("No Data.");
    }
  }, onerror);
});
var Go = function() {
  location.href = "/player.html?" + $("#url-input").val();
}
var OpenForm = function(elm) {
  $(elm).parent().addClass("hide");
  $("#form-container").removeClass("hide");
}
