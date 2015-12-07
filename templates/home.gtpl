<html>
  <head>
    <title>Go Logbook</title>
  </head>
  <body>
<form action="/sign" method="post">
      <div><input type="text" name="Kreading"></input></div>
      <div><input type="text" name="Rreading"></input></div>
      <div><input type="text" name="Creading"></input></div>
      <div><input type="submit" value="Submit"></div>
</form>
<h3>Your logs:</h3>
<table>
<tr>
<th>Date</th>
<th>KK</th>
<th>RSL</th>
<th>CVN</th>
</tr>

{{range .}}
<tr>
<td>{{.Date}}</td>
<td>{{.Kreading}}</td>
<td>{{.Rreading}}</td>
<td>{{.Creading}}</td>
</tr>
{{end}}
</table>

</body>
</html>
