var now = Date.now();
console.log('[child] child process start!');

setInterval(function () {
  var interval = Date.now() - now;
  console.log('[child] Timer finished, time consuming: ' + interval);
}, 50);
