<!DOCTYPE html>
<head>
	<meta charset="utf-8">
	<title>Kate</title>
	<link rel="stylesheet" type="text/css" href="/static/css/style.css" />
   <link rel="stylesheet" type="text/css" href="/static/css/login.css" />
   <link rel="shortcut icon" href="/static/img/icon.png">
</head>
<body>

<!-- top bar -->
    <div class="topbar">
        <span class="right">
           <ul id="menu">
              <li><p>{{.Lang}}</p>
                  <ul>
                     <li><a href="/lang/eng">English</a></li>
                     <li><a href="/lang/ru">Русский</a></li>
                     <li><a href="/lang/chn">中文</a></li>
                  </ul>
              </li>      
           </ul>
           <p><a href="/">{{.Main}}</a> <a href={{.LogURL}}>{{.LogStr}}</a></p>
        </span>
    <div class="clr"></div>
    </div>
<!--/ top bar -->

	<form id="slick-login" method="post">
		<label for="name">{{.Login}}:</label><input type="text" name="name" class="placeholder" placeholder={{.Login}}>
		<label for="password">{{.Password}}</label><input type="password" name="password" class="placeholder" placeholder={{.Password}}>
		<input type="submit" value={{.Reg}}>
	</form>
</body>
</html>
