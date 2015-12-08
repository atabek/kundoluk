<html>
  <head>
    <title>Go Logbook</title>
  </head>
<!-- Latest compiled and minified CSS -->
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">

<!-- Optional theme -->
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap-theme.min.css" integrity="sha384-fLW2N01lMqjakBkx3l/M9EahuwpSfeNvV63J5ezn3uZzapT0u7EYsXMjQV+0En5r" crossorigin="anonymous">

<!-- Latest compiled and minified JavaScript -->
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
  <body>

<div class="container">

<h1>Logbook</h1>
<form action="/sign" method="post">
<div class="row">
      <div class="col-md-1">KK<input class="col-sm-12" type="number" name="Kreading"></input></div>
      <div class="col-md-1">RSL<input class="col-sm-12" type="number" name="Rreading"></input></div>
      <div class="col-md-1">CVN<input class="col-sm-12" type="number" name="Creading"></input></div>
      <div class="col-md-1">Meal<input class="col-sm-12" type="number" name="Mreading"></input></div>
      <div class="col-md-1">PNT<input class="col-sm-12" type="number" name="Preading"></input></div>
      <div class="col-md-1">THC<input class="col-sm-12" type="number" name="Technight"></input></div>
      <div class="col-md-1">Orozo<input class="col-sm-12" type="number" name="Orozo"></input></div>
      <div class="col-md-1">Video<input class="col-sm-12" type="number" name="Video"></input></div>
      <div class="col-md-1">Taspi<input class="col-sm-12" type="number" name="Taspi"></input></div>
      <div class="col-md-1">Kaza<input class="col-sm-12" type="number" name="Kaza"></input></div>
      <div class="col-md-1">Kop<input class="col-sm-12" type="number" name="Kop"></input></div>
      <div class="col-md-1">Sabak<input class="col-sm-12" type="number" name="Sabak"></input></div>
</div>
<input type="submit" class="btn btn-primary" value="Submit">
</form>
</div>
</div>

<div class="container">
<h3>Your logs:</h3>

<table class="table table-bordered">
	<thead>
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
			<th>Kop</th>
			<th>Sabak</th>
		</tr>
	</thead>

	<tbody>
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
			<td>{{.Kop}}</td>
			<td>{{.Sabak}}</td>
		</tr>
	{{end}}
	</tbody>
</table>
</div>

</body>
</html>
