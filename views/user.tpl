<!DOCTYPE html>
<head>
	<meta charset="utf-8">
	<title>Kate</title>
	<link rel="stylesheet" type="text/css" href="/static/css/style.css" />
   <link rel="stylesheet" type="text/css" href="/static/css/user.css" />
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
<div class="user-area">
	
    <h1>{{.Name}}</h1>
    <br>
    {{if .Exist}} 
       <a class="user-button" href = "/followers/{{.Id}}">{{.Followers}}</a>
       <br>
       <a class="user-button" href = "/following/{{.Id}}">{{.Following}}</a>
       <br>
       {{if .Owner}}
          <a class="user-button" href="../add"> {{.Posting}}</a>
       {{else}}    
          <a class="user-button" href="../follow/{{.Id}}">{{.Follow}}</a>
       {{end}}
    {{end}}

</div>

<div>
{{range $id, $post := .UserPosts}}
  {{range $name, $content := $post}}
    {{$uid := $id | pos}}
    <div class="user-area" id="link" class="post-area">
        <a class="post-header"; href="../posts/{{$uid}}">{{$name}}</a>
        <p>{{$content}}</p>
    </div>
  {{end}}
{{end}}
</div>
<br>
<br>
</div>
</body>
</html>
