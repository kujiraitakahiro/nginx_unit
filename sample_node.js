//nginx unitで使う場合、予め「npm install -g --unsafe-perm unit-http」or 「npm update -g --unsafe-perm unit-http」を実行
var http = require('http');
//nginx unitでは require('http') を require('unit-http') に変更する

var server = http.createServer(function (req, res) {
  res.writeHead(200, {"Content-Type": "text/plain"});
  res.end('node.js sample program\n');
});

server.listen(7778);
