package main

import (
	"fmt"
	"github.com/labstack/echo"

	"net/http"
	"net/smtp"
)

var (
	// htmlText = "<style>input{border-radius:5px;}</style>" +
	// 	"<body bgcolor='black'><center>" +
	// 	"<h1 style='color:lightgreen'>Hygull</h1>" +
	// 	"<form method='post' action='/v1/validation'>" +
	// 	"<input type='text' name='email' placeholder='rishikesh@gmail.com' style='height:20px'><br><br>" +
	// 	"<input type='text' name='password' placeholder='@Rishi#123$' style='height:20px'><br>" +
	// 	"<input type='submit' value='submit'>" +
	// 	"</form>" +
	// 	"</center>"

	htmlText = `<!DOCTYPE html>
<html lang='en'>
<head>
  <title>Hygull</title>
  <meta charset='utf-8'>
  <meta name='viewport' content='width=device-width, initial-scale=1'>
  <link rel='stylesheet' href='https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css'>
  <script src='https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js'></script>
  <script src='https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js'></script>
  <script src="https://apis.google.com/js/platform.js" async defer></script>
      <meta name="google-signin-scope" content="profile email">
    <meta name="427029365574-g0djrldtjgq997348odprihh3k0ujse3.apps.googleusercontent.com" content="GxcRabmezc3wSSpVZvWi2Pui.apps.googleusercontent.com">
  <style type='text/css'>
    #title{
      color:lightgreen;
    }
    body{
      background-color: white;
    }
    span#t1{
      color:green;
    }
    span#t2{
      color:orange;
    }
    #form_section{
      margin-top: 15px;
      margin-left:1px solid green;
    }
    a{
      font-weight: bold;
    }
  	body, html,div,head{
  		margin:0;
  		padding:0;
  	}
  	a{
  		text-decoration:none;

  	}
  	span{
  		font-weight:bold;
  	}
  </style>
  <script>
  function FocusOnInput(){
    document.getElementById('usr').focus();
    return true
};
  </script>
</head>
<body>

<nav class="navbar navbar-inverse">
  <div class="container-fluid">
    <div class="navbar-header">
      <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#myNavbar">
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>                        
      </button>
      <a class="navbar-brand" href="#">www.hygull.com</a>
    </div>
    <div class="collapse navbar-collapse" id="myNavbar">
      <ul class="nav navbar-nav">
        <li class="active"><a href="#">Never go back</a></li>
        <li class="dropdown">
          <a class="dropdown-toggle" data-toggle="dropdown" href="#">Imp<span class="caret"></span></a>
          <ul class="dropdown-menu">
            <li><a href="#">Quantitative aptitude</a></li>
            <li><a href="#">Logical resoning</a></li>
            <li><a href="#">English</a></li>
            <li><a href="#">Data structure</a></li>
            <li><a href="#">Progamming</a></li>
          </ul>
        </li>
        <li><a href="#">Tutorials</a></li>
        <li><a href="#">Books</a></li>
      </ul>
      <ul class="nav navbar-nav navbar-right">
        <li><a href="#"><span class="glyphicon glyphicon-user"></span> Sign Up</a></li>
        <li><a href="#usr"><span class="glyphicon glyphicon-log-in"></span>Login</a></li>
      </ul>
    </div>
  </div>
</nav>
<div class='container-fluid'>
 
 <div class='row'>
 <div class='col-sm-2 text-center'>
 	  <a href='https://slack.com/' target='_blank'><button type="button" class="btn btn-primary btn-block">Slack</button></a>
  	  <a href='https://mail.google.com/' style='text-decoration: none;' target='_blank'><button type="button" class="btn btn-default btn-block">Gmail</button></a>

 	  <a href='http://www.geeksforgeeks.org/' style='text-decoration: none;' target='_blank' ><button type="button" class="btn btn-success btn-block">Geeksforgeeks</button></a>
 	
 </div>
  <div class='col-sm-7 text-center'>
  <div class="g-signin2" data-onsuccess="onSignIn"></div>
  <h1 id='title'><span id='t1'>H</span><span id='t2'>y</span><span id='t1'>g</span><span id='t2'>u</span><span id='t1'>l</span><span id='t2'>l</span></h1>
  <center>
  <h5><span style='color:green'>Golang - </span><span>Implementation of basics</span></h5>
  <hr>
  <a href='https://gist.github.com/hygull/f8810d4f6e56bc4e2e277b33313ae9a0' target='_blank'>Operator precedence </a><br>
   <hr>
   <a href='https://gist.github.com/hygull/564a67122f359d6b8f30193e37b28c8a' target='_blank'>3 different forms of for loop</a><br>
     <hr>
    <a href='https://gist.github.com/hygull/0f7d8e3981d7690fbb4fcfd96a2df6b8' target='_blank'>Checking availability of keys in map</a><br>
      <hr>
     <a href='https://gist.github.com/hygull/7d7ea0dc03ab51fceb673d2e3bad989f' target='_blank'>strconv (Converting a string to an int)</a><br>
       <hr>
      <a href='https://gist.github.com/hygull/5470cdd04b084e78451d3360727ec4e4' target='_blank'>Pointer and memory allocation for an object  with new keyword</a><br>
        <hr>
       <a href='https://gist.github.com/hygull/53792994c9f79e09d19ca5cf2bc7b853' target='_blank'>Marshalling an array of structs into JSON</a><br>
       <!-- <hr>
        <a href='https://gist.github.com/hygull/d308d375b4d2a19bd087efc24d60b432' target='_blank'>Variable number of arguments</a><br>
          <hr>
         <a href='https://gist.github.com/hygull/d74b65b70fca4a8971014e7b44689a72' target='_blank'>3 ways of using for loop with map</a><br>
          <hr>
          <a href='https://gist.github.com/hygull/fb240bfafbd2077b702f67280e8bd1f9' target='_blank'>Printing system time and its decoration</a><br>
           <hr>
           <a href='https://gist.github.com/hygull/b2ff9914c8504882929004776fed3471' target='_blank'>Splitting a string into list of words</a><br>
  			-->
  <br>
   <button type="button" class="btn btn-success" onclick="FocusOnInput()">LOGIN for more links</button>
  <center>
  </div>
  <div class='col-sm-3' id='form_section' >
  <form class='form-horizontal' method='post' action='/v1/validate'>
    <div class='form-group'>
      <!-- <label class='control-label col-sm-2' for='email'>Email</label> -->
      <div class='col-sm-offset-2  col-sm-10'>
        <input type='text' class='form-control' id='usr' autocomplete='off' name='username' placeholder='Enter username(any)'>
      </div>
    </div>
    <div class='form-group'>
      <!-- <label class='control-label col-sm-2' for='pwd'>Password</label> -->
      <div class='col-sm-offset-2  col-sm-10'>          
        <input type='password' class='form-control' name='password' placeholder='Enter password'>
      </div>
    </div>
    <div class='form-group'>        
      <!-- <div class='col-sm-offset-2 col-sm-10'> -->
      <div class='col-sm-offset-2 col-sm-10'>
        <div class='checkbox'>
          <label><input type='checkbox'> Remember me</label>
        </div>
      </div>
    </div>
    <div class='form-group'>        
      <div class='col-sm-offset-2 col-sm-10'>
        <button type='submit' class='btn btn-success'>Submit</button>
      </div>
    </div>
  </form>
  <br>
  
  <hr>
  <br>
  <center>
    <img src='https://avatars1.githubusercontent.com/u/23253178?v=3&s=460' class="img-circle" alt="Rishikesh Agrawani" width=90px height=90px>  	
 	<br><br><a href='https://www.facebook.com/rishikesh.agrawani'><img src='https://cdn4.iconfinder.com/data/icons/social-messaging-ui-color-shapes-2-free/128/social-facebook-circle-24.png'></a>
  <a href='https://twitter.com/RishiAndCodeEx'><img src='https://cdn2.iconfinder.com/data/icons/flatte-internet-and-websites/80/13_-_Twitter_bird-32.png'></a>
  <a href='https://www.linkedin.com/in/rishikesh-agrawani-0358ba119?trk=nav_responsive_tab_profile' target='_blank'><img src='https://cdn3.iconfinder.com/data/icons/free-social-icons/67/linkedin_circle_color-32.png'></a>
  <br><br>

  <span style='color:gray'>Follow me on </span> <br> <a href='https://plus.google.com/u/0/113992208941719062433' target='_blank'><img src='https://cdn4.iconfinder.com/data/icons/iconsimple-logotypes/512/google-32.png'></a>
  
  
  </div><!--col2 end-->
  </div>
</div>

</body>
</html>
`
)

var (
	passwd = "go"
	part1  = "<body bgcolor=black style='padding-top:100px'><center>"
	part2  = "<br><a href='/v1/home' style='color:lightgreen'>Back to login</a></center></body>"
)

func main() {
	ec := echo.New()

	ec.POST("/v1/validate", func(ctx echo.Context) error {
		username := ctx.FormValue("username")
		password := ctx.FormValue("password")

		if username == "" || password == "" {
			text := "<h2 style='color:white'>Username & password both are required</h2>"
			return ctx.HTML(200, part1+text+part2)
		} else if !(len(username) > 4 && len(username) < 16) {
			fmt.Println(username)
			text := "<h2 style='color:white'>Only 5-15 characters are allowed in username</h2>"
			return ctx.HTML(200, part1+text+part2)
		} else if password != passwd {
			text := "<h2 style='color:red'>password is fixed</h2><h2 style='color:white'> request for password by </h2><h2 style='color:lightgreen'>entering your email on the box below</h2>" +
				"<form action='/v1/send/email' method='post'><br><br><input type='email' name='email' style='font-size:20px;border-radius:4px' placeholder='Enter you email'>" +
				"<br><br><input type='submit' style='font-size:15px;border-radius:4px' height=20px value='send'><br>Or<br>"
			return ctx.HTML(200, part1+text+part2)
		}
		return ctx.Redirect(302, "/v1/welcome/"+username)
	})

	var User = ""
	ec.GET("/v1/welcome/:username", func(ctx echo.Context) error {
		username := ctx.Param("username")
		User = username
		fmt.Println("Username : ", username)
		//return ctx.HTML(200, htmlText2_1+"<br>"+part1+"<center><h2 style='color:lightgreen'>Hello "+username+"</h2>"+part2+"<br>"+htmlText2_2)
		return ctx.HTML(200, getHtml(username))
	})

	ec.GET("/v1/home", func(ctx echo.Context) error {
		return ctx.HTML(http.StatusOK, htmlText)
	})

	ec.POST("/v1/send/email", func(ctx echo.Context) error {
		email := ctx.FormValue("email")
		fmt.Println("Sending an email to ", email)

		msg := "From: " + "rishikesh0014051992@gmail.com" + "\r\n" +
			"To: " + email + "\r\n" +
			"MIME-Version: 1.0" + "\r\n" +
			"Content-type: text/html" + "\r\n" +
			"Subject: Password to login into Hygull" + "\r\n\r\n" +
			"<h3 style='color:green'>" + passwd + "<h3>" + "\r\n"
		auth := smtp.PlainAuth(
			"",
			"rishikesh0014051992@gmail.com",
			"29915041",
			"smtp.gmail.com",
		)
		// Connect to the server, authenticate, set the sender and recipient,
		// and send the email all in one step.
		err := smtp.SendMail(
			"smtp.gmail.com:25",
			auth,
			"rishikesh0014051992@gmail.com",
			[]string{email},
			[]byte(msg),
		)
		if err == nil {
			// log.Fatal(err)
			return ctx.HTML(200, part1+"<h2 style='color:lightgreen'>email succesfully sent to <h2><h2 style='color:white'>"+email+"</h2>"+part2)
		}
		fmt.Println(err)
		return ctx.HTML(200, part1+"<h2 style='color:lightgreen'>Unable to send an email to "+email+"</h2>"+part2)

	})
	fmt.Println("******************** Welcome to Hygull **************************\n")
	fmt.Println("Visit   -->  127.0.0.1:8080/v1/home")
	ec.Logger.Fatal(ec.Start(":8080"))
}

func getHtml(user string) string {
	htmlText2_1 := `
<!DOCTYPE html>
<html lang='en'>
<head>
  <title>Hygull</title>
  <meta charset='utf-8'>
  <meta name='viewport' content='width=device-width, initial-scale=1'>
  <link rel='stylesheet' href='https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css'>
  <script src='https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js'></script>
  <script src='https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js'></script>
  <style type='text/css'>
    #title{
      color:lightgreen;
    }
    body{
      background-color: white;
    }
    span#t1{
      color:green;
    }
    span#t2{
      color:orange;
    }
    #form_section{
      margin-top: 15px;
      margin-left:1px solid green;
    }
    a{
      font-weight: bold;
    }
  	body, html,div,head{
  		margin:0;
  		padding:0;
  	}
  	a{
  		text-decoration:none;

  	}
  	span{
  		font-weight:bold;
  	}
  </style>
  <script>
  function FocusOnInput(){
    document.getElementById('usr').focus();
    return true
};
  </script>
</head>
<body>

<nav class="navbar navbar-inverse">
  <div class="container-fluid">
    <div class="navbar-header">
      <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#myNavbar">
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>                        
      </button>
      <a class="navbar-brand" href="#">www.hygull.com</a>
    </div>
    <div class="collapse navbar-collapse" id="myNavbar">
      <ul class="nav navbar-nav">
        <li class="active"><a href="#">Never go back</a></li>
        <li class="dropdown">
          <a class="dropdown-toggle" data-toggle="dropdown" href="#">Imp<span class="caret"></span></a>
          <ul class="dropdown-menu">
            <li><a href="#">Quantitative aptitude</a></li>
            <li><a href="#">Logical resoning</a></li>
            <li><a href="#">English</a></li>
            <li><a href="#">Data structure</a></li>
            <li><a href="#">Progamming</a></li>
          </ul>
        </li>
        <li><a href="#">Tutorials</a></li>
        <li><a href="#">Books</a></li>
      </ul>
      <ul class="nav navbar-nav navbar-right">
        <li><a href="#"><span class="glyphicon glyphicon-user"></span>` + user + `</a></li>
        <li><a href="#usr"><span class="glyphicon glyphicon-log-in"></span>Login</a></li>
      </ul>
    </div>
  </div>
</nav>
<div class='container-fluid'>
 
 <div class='row'>
 <div class='col-sm-2 text-center'>
 	  <a href='https://slack.com/' target='_blank'><button type="button" class="btn btn-primary btn-block">Slack</button></a>
  	  <a href='https://mail.google.com/' style='text-decoration: none;' target='_blank'><button type="button" class="btn btn-default btn-block">Gmail</button></a>

 	  <a href='http://www.geeksforgeeks.org/' style='text-decoration: none;' target='_blank' ><button type="button" class="btn btn-success btn-block">Geeksforgeeks</button></a>
 	
 </div>
  <div class='col-sm-7 text-center'>
  <h1 id='title'><span id='t1'>H</span><span id='t2'>y</span><span id='t1'>g</span><span id='t2'>u</span><span id='t1'>l</span><span id='t2'>l</span></h1>
<!--   <center> -->
  <h5><span style='color:green'>Golang - </span><span>Implementation of basics</span></h5>
  <hr>
  <a href='https://gist.github.com/hygull/f8810d4f6e56bc4e2e277b33313ae9a0' target='_blank'>Operator precedence </a><br>
   <hr>
   <a href='https://gist.github.com/hygull/564a67122f359d6b8f30193e37b28c8a' target='_blank'>3 different forms of for loop</a><br>
     <hr>
    <a href='https://gist.github.com/hygull/0f7d8e3981d7690fbb4fcfd96a2df6b8' target='_blank'>Checking availability of keys in map</a><br>
      <hr>
     <a href='https://gist.github.com/hygull/7d7ea0dc03ab51fceb673d2e3bad989f' target='_blank'>strconv (Converting a string to an int)</a><br>
       <hr>
      <a href='https://gist.github.com/hygull/5470cdd04b084e78451d3360727ec4e4' target='_blank'>Pointer and memory allocation for an object  with new keyword</a><br>
        <hr>
       <a href='https://gist.github.com/hygull/53792994c9f79e09d19ca5cf2bc7b853' target='_blank'>Marshalling an array of structs into JSON</a><br>
       <hr>
        <a href='https://gist.github.com/hygull/d308d375b4d2a19bd087efc24d60b432' target='_blank'>Variable number of arguments</a><br>
          <hr>
         <a href='https://gist.github.com/hygull/d74b65b70fca4a8971014e7b44689a72' target='_blank'>3 ways of using for loop with map</a><br>
          <hr>
          <a href='https://gist.github.com/hygull/fb240bfafbd2077b702f67280e8bd1f9' target='_blank'>Printing system time and its decoration</a><br>
           <hr>
           <a href='https://gist.github.com/hygull/b2ff9914c8504882929004776fed3471' target='_blank'>Splitting a string into list of words</a><br>
  			
<!--  </center> -->
  </div>
  <div class='col-sm-3' id='form_section' >
  <form class='form-horizontal' method='post' action='/v1/validate'>
    <div class='form-group'>
      <!-- <label class='control-label col-sm-2' for='email'>Email</label> -->
      <div class='col-sm-offset-2  col-sm-10'>
        <input type='text' class='form-control' id='usr' autocomplete='off' name='username' placeholder='Enter username(any)'>
      </div>
    </div>
    <div class='form-group'>
      <!-- <label class='control-label col-sm-2' for='pwd'>Password</label> -->
      <div class='col-sm-offset-2  col-sm-10'>          
        <input type='password' class='form-control' name='password' placeholder='Enter password'>
      </div>
    </div>
    <div class='form-group'>        
      <!-- <div class='col-sm-offset-2 col-sm-10'> -->
      <div class='col-sm-offset-2 col-sm-10'>
        <div class='checkbox'>
          <label><input type='checkbox'> Remember me</label>
        </div>
      </div>
    </div>
    <div class='form-group'>        
      <div class='col-sm-offset-2 col-sm-10'>
        <button type='submit' class='btn btn-success'>Submit</button>
      </div>
    </div>
  </form>
  <br>
  
  <hr>
  <br>
  <center>
    <img src='https://avatars1.githubusercontent.com/u/23253178?v=3&s=460' class="img-circle" alt="Rishikesh Agrawani" width=90px height=90px>  	
 	<br><br><a href='https://www.facebook.com/rishikesh.agrawani'><img src='https://cdn4.iconfinder.com/data/icons/social-messaging-ui-color-shapes-2-free/128/social-facebook-circle-24.png'></a>
  <a href='https://twitter.com/RishiAndCodeEx'><img src='https://cdn2.iconfinder.com/data/icons/flatte-internet-and-websites/80/13_-_Twitter_bird-32.png'></a>
  <a href='https://www.linkedin.com/in/rishikesh-agrawani-0358ba119?trk=nav_responsive_tab_profile' target='_blank'><img src='https://cdn3.iconfinder.com/data/icons/free-social-icons/67/linkedin_circle_color-32.png'></a>
  <br><br>

  <span style='color:gray'>Follow me on </span> <br> <a href='https://plus.google.com/u/0/113992208941719062433' target='_blank'><img src='https://cdn4.iconfinder.com/data/icons/iconsimple-logotypes/512/google-32.png'></a>
  
  
  </div><!--col2 end-->
  </div>
</div>

</body>
</html>
`
	return htmlText2_1
}

var htmlText2_2 = ``
