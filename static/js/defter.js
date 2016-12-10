document.addEventListener('DOMContentLoaded', function(){ 
  document.getElementById("corner").addEventListener("click",function(e) {
    window.location.href = "/";
  }, false);

  document.getElementById("file-selector").addEventListener("change",function(e) {
    alert("test");
  }, false);
}, false);
