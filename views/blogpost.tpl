<!DOCTYPE html>
<head>
	<meta charset="utf-8">
	<title>Kate</title>
	<link rel="stylesheet" type="text/css" href="/static/css/style.css" />
	<link rel="stylesheet" type="text/css" href="/static/css/user.css" />
	<link rel="stylesheet" type="text/css" href="/static/css/post.css" />
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
  
    <div class="user-area">
        <h1>{{.PostName}}</h1>
        <br>
        <p>{{.Content}}</p>
        <br>
        <br>
        {{if .Owner}}
            <a class="post-button" href="../delete/post/{{.Id}} ">{{.DeletePost}}</a>
            <a class="post-button" href="../edit/post/{{.Id}} ">{{.Edit}}</a>
        {{end}}
     </div>
     <div>
       <form id="posting" method="post">
		    <label for="content">{{.Comment}}:</label><p><textarea name="content" cols="50" rows="5" placeholder={{.Comment}}></textarea></p>
		    <input type="submit" value={{.Send}}>
	    </form>
     </div>
     <div>
         {{$sess := .Sess}}
         {{$delete := .Delete}}
         {{range $id, $comment := .PostComments}}
             {{$owner := $comment.Owner}}
             {{$content := $comment.Content}}
             {{$name := $comment.OwnerName}}
             <div class="user-area">
                  <a class="content-link" href="../user/{{$owner}}">{{$name}}</a>
                  <br>
                  <br>
                  <p class="postarea">{{str2html $content}}</p>
                  <br>
                  {{if eq $sess $owner}}
                      {{$d := $id | pos}}
                      <a class="post-button" href="/delete/comment/{{$d}}">{{$delete}}</a>
                  {{end}}
             </div>
          {{end}}
          <br>
          <br>
    </div>
</div>
</body>
</html> 
