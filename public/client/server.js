const path = require("path");
const express = require("express");
const app = express();
app.use(express.static(__dirname + '/dist'));
app.get('/*', function(req,res){
  res.sendFile(path.join(__dirname, '/dist','index.html'))
});
app.listen(process.env.PORT || 5100)
// function requireHTTPS(req, res, next) {
//   // The 'x-forwarded-proto' check is for Heroku
//   if (!req.secure && req.get('x-forwarded-proto') !== 'https') {
//       return res.redirect('https://' + req.get('host') + req.url);
//   }
//   next();
// }
// const express = require('express');
// const app = express();
// app.use(requireHTTPS);
// // app.use('../public', serveStatic(path.join(__dirname, '/client')))
// // app.use(express.static(path.join(__dirname,'public')))

// app.use(express.static(__dirname + "/public/client/dist"));

// app.get('/*', function(req, res) {
//   res.sendFile(path.join(__dirname + '/dist/client/index.html'
//   ));});

//   app.listen(process.env.PORT || 4200)
