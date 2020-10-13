$(function(){
  var qs = window.location.search;
  var qo = new Object();
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
  console.log(qo);
});
