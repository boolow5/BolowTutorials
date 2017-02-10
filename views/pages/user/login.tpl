<h1 class="text-center">Login</h1>

<form action="/login" method="post" class="login-form">
  {{.xsrfdata}}
  <div class="row">
    <div class="small-3 columns">
      <label for="username" class="text-right middle">Username</label>
    </div>
    <div class="small-9 columns">
      <input type="text" id="username" name="username" placeholder="Username">
    </div>
  </div>
  <div class="row">
    <div class="small-3 columns">
      <label for="password" class="text-right middle">Password</label>
    </div>
    <div class="small-9 columns">
      <input type="password" id="password" name="password" placeholder="Password">
    </div>
  </div>
  <div class="row">
    <div class="small-3 columns"></div>
    <div class="small-9 columns">
      <button class="hollow expanded button">Login</button>
      <p>
        If you're not registered <a href="/register">Register here</a>!
      </p>
    </div>
  </div>
</form>
