package main

import "github.com/kataras/iris"

func Hello(ctx *iris.Context) {
	ctx.SetContentType("text/html")
	ctx.Write("<!doctype html>" +
		"<script type='text/javascript'>" +
		"function al(){alert('HI GOLANGERS...ARE YOU OK WITH ME?');}" +
		"</script>" +
		"<body bgcolor='black'>" +
		"<h1 style='color:white'>Great</h1><button onclick='al()'>Click</button>" +
		"</body>")
}

func Welcome(ctx *iris.Context) {
	//iris.Config.IsDevelopment = true
	ctx.SetContentType("text/html")
	ctx.MustRender("welcome.html", struct { /*  ./templates/welcome.thml  */
		Name     string
		Message  string
		NamesArr []string
	}{"Rishikesh", "Welocme Golangers", []string{"Rishikesh", "Darshan"}})
	//}{"Rishikesh", "Welocme Golangers", "<h3 style='color:green;'>Programmers</h3>"})
}

func main() {
	iris.Get("/v1/hello", Hello)
	iris.Get("/v1/welcome", Welcome)
	iris.Listen(":3000")
}
