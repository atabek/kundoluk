<html>
  <head>
    <title>Go Logbook</title>
  </head>
  <body>
<form action="/sign" method="post">
      <div><input type="text" name="Kreading"></input></div>
      <div><input type="text" name="Rreading"></input></div>
      <div><input type="text" name="Creading"></input></div>
      <div><input type="text" name="Mreading"></input></div>
      <div><input type="text" name="Preading"></input></div>
      <div><input type="text" name="Technight"></input></div>
      <div><input type="text" name="Orozo"></input></div>
      <div><input type="text" name="Video"></input></div>
      <div><input type="text" name="Taspi"></input></div>
      <div><input type="text" name="Kaza"></input></div>
      <div><input type="text" name="Ezb"></input></div>
      <div><input type="text" name="Kop"></input></div>
      <div><input type="text" name="Sabak"></input></div>
      <div><input type="submit" value="Submit"></div>
</form>
<h3>Your logs:</h3>
<table>
<tr>
<th>Date</th>
<th>KK</th>
<th>RSL</th>
<th>CVN</th>
<th>Meal</th>
<th>PNT</th>
<th>THC</th>
<th>Orozo</th>
<th>Video</th>
<th>Taspi</th>
<th>Kaza</th>
<th>Ezb</th>
<th>Kop</th>
<th>Sabak</th>
</tr>

{{range .}}
<tr>
<td>{{.Date}}</td>
<td>{{.Kreading}}</td>
<td>{{.Rreading}}</td>
<td>{{.Creading}}</td>
<td>{{.Mreading}}</td>
<td>{{.Preading}}</td>
<td>{{.Technight}}</td>
<td>{{.Orozo}}</td>
<td>{{.Video}}</td>
<td>{{.Taspi}}</td>
<td>{{.Kaza}}</td>
<td>{{.Ezb}}</td>
<td>{{.Kop}}</td>
<td>{{.Sabak}}</td>
</tr>
{{end}}
</table>

</body>
</html>
