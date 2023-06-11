function requireHTTPS(req, res, next) {
  // The 'x-forwarded-proto' check is for Heroku
  if (!req.secure && req.get('x-forwarded-proto') !== 'https') {
      return res.redirect('https://' + req.get('host') + req.url);
  }
  next();
}
const express = require('express');
const app = express();
app.use(requireHTTPS);
app.use('/lafam', serveStatic(path.join(__dirname, '/public')))
app.use(express.static(path.join(__dirname,'public')))

app.use(express.static(__dirname + "/public/client/dist"));

app.get('/*', function(req, res) {
  res.sendFile(path.join(__dirname + '/dist/client/index.html'
  ));});

  app.listen(process.env.PORT || 4200)
