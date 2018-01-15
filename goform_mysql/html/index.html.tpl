<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
  <title>Goform</title>
</head>
<body>
  <div>
      <form method="POST" action="/">
          <input name="key" type="text" value="">
          <input type="submit" value="submit" />
          <input type="button" value="refresh" onClick="location.href='/'">
      </form>
      <hr>
      <h4>Post data</h4>
      {{range .}}
      <p>{{.}}</p>
      {{end}}
  </div>
</body>
</html>
