var Search = function() {
  console.log($("#search-word").val());
}
var showDangerMessage = function(str) {
  $("#message-danger").text(str);
  $("#message-danger-container").removeClass("hide");
}
var showInfoMessage = function(str) {
  $("#message-info").text(str);
  $("#message-info-container").removeClass("hide");
}
var CloseMessage = function(elm) {
  $(elm).parent().addClass("hide");
}
