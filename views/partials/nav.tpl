<!-- Navigation -->
<div class="title-bar" data-responsive-toggle="realEstateMenu" data-hide-for="small">
  <button class="menu-icon" type="button" data-toggle></button>
  <div class="title-bar-title">Menu</div>
</div>

<div class="top-bar" id="realEstateMenu">
  <div class="top-bar-left">
    <ul class="menu" data-responsive-menu="accordion">
      <li class="menu-text">
        <a href="/">
          <i class="fa fa-home fa-lg"></i>
          Qaamuuska Ciyaalka Xaafada!
        </a>
      </li>
      <li>
        <a href="#">
          <i class="fa fa-newspaper-o fa-lg"></i>
          Ereyo Cusub
        </a>
      </li>
      <li>
        <a href="#">
          <i class="fa fa-star-o fa-lg"></i>
          Ereyada ugu Caansan
        </a>
      </li>
      <li>
        <a href="#">
          <i class="fa fa-plus fa-lg"></i>
          Geli Erey Cusub
        </a>
      </li>
    </ul>
  </div>
  <div class="top-bar-right">
    <ul class="menu">
      <li>
        {{.IsLoggedin}}
      </li>
      {{if .IsLoggedin}}
      <li><a href="#">Your Name</a></li>

      <li>
        <a class="large button" href="/logout">
          <i class="fa fa-user-circle fa-lg"></i>
          Logout
        </a>
      </li>
      {{end}}
      {{if not .IsLoggedin}}
      <li><a href="/register">Register</a></li>

      <li>
        <a class="large button" href="/login">
          <i class="fa fa-user-circle fa-lg"></i>
          Login
        </a>
      </li>
      {{end}}

    </ul>
  </div>
</div>
<!-- /Navigation -->
