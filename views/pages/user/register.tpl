<h1 class="text-center">Register</h1>


<form action="/register" method="post" id="register-form">
  {{.xsrfdata}}
  <div class="row">
    <div class="small-3 columns">
      <label for="username" class="text-right middle">Username</label>
    </div>
    <div class="small-9 columns">
      <input type="text" id="username" name="username" placeholder="Username" required>
    </div>
  </div>
  <div class="row">
    <div class="small-3 columns">
      <label for="password" class="text-right middle">Password</label>
    </div>
    <div class="small-9 columns">
      <input type="password" id="password" name="password" placeholder="Password" required>
    </div>
  </div>
  <div class="row">
    <div class="small-3 columns"></div>
    <div class="small-9 columns">
      <button class="hollow expanded button">Register</button>
      <p>
        If you're already registered <a href="/login">Login in here</a>!
      </p>
    </div>
  </div>
</form>
