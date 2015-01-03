<!DOCTYPE html>
<head>
	<meta charset="utf-8">
	<title>Kate</title>
	<link rel="stylesheet" type="text/css" href="/static/css/style.css" />
   <link rel="stylesheet" type="text/css" href="/static/css/post.css" />
   <link rel="stylesheet" type="text/css" href="/static/css/user.css" />
   <link rel="shortcut icon" href="/static/img/fox.icon.png">
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
   
      <div class="user-area"><h2>{{.Edit}}</h2></div>
	   <form id="posting" method="post">
		   <label for="name"></label>
		   <br>
		   <input type="text" name="name" class="placeholder" value={{.Topic}}>
		   <br>
		   <br>
		   <label for="content">{{.BlogPost}}:</label><p><textarea name="content" cols="50" rows="15">{{.Content}}</textarea></p>
		   <input type="submit" value={{.Edit}}>
	   </form>
</div>
	
</body>
</html>
