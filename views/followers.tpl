<!DOCTYPE html>
<head>
	<meta charset="utf-8">
	<title>Kate</title>
	<link rel="stylesheet" type="text/css" href="/static/css/style.css" />
        <link rel="stylesheet" type="text/css" href="/static/css/user.css" />
        <link rel="stylesheet" type="text/css" href="/static/css/content-links.css" />
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


<div>

<div id="left">
    <img src="/static/img/fox.png" alt="fox" class="fox" />
</div>
<div id="right">
    <img src="/static/img/fox.png" alt="fox" class="fox" />
</div>

<div class="user-area" id="link">
    <h1>{{.Followers}}</h1>
    <br>
    {{if .FolEx}}
    {{range $id, $name := .AllFollowers}}
    <a class ="content-link" href ="/user/{{$id}}">{{$name}}</a><br><br>
    {{end}}
    {{else}}
    {{range $id, $name := .AllFollowers}}
    <p>{{$name}}</p>
    {{end}}
    {{end}}
</div>

</div>
</body>
</html>
