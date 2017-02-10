<!doctype html>
<html class="no-js" lang="en">
<head>
  {{.Head}}
</head>
<body>
  {{.Nav}}

  <br>

  {{.FlashMessages}}

  {{if .IsHome}}
    {{.Featured}}
    <div class="row column">
      <hr>
    </div>
  {{end}}

  <div class="row main-content">
    <div class="medium-8 column">
      {{.LayoutContent}}
    </div>

    <div class="medium-4 column">
      <div class="row column">
        <p class="lead">Class Details:</p>
      </div>
      <div class="row column callout">
        <h1>Side bar</h1>
      </div>
    </div>
  </div>


  {{.Footer}}
  {{.Scripts}}

</body>
</html>
