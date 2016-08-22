package main

import "github.com/kataras/iris"

type VideoAPI struct {
	*iris.Context
}

//Render Developer portal
//contains samples of video format
func developerPortal(ctx *iris.Context) {
	ctx.Render("developerportal.html", struct{ Name string }{Name: "iris"})
}

/**************/
//Video API
/*************/

// GET /videos
func (u VideoAPI) Get() {
	u.Write("/assets/videosdatabase.json")
	// u.JSON(iris.StatusOK,myDb.AllVideos())

}

// GET /videos/:param1 which its value passed to the id argument
func (u VideoAPI) GetBy(id string) { // id equals to u.Param("param1")
	u.Write("Get from /videos/%s", id)
	// u.JSON(iris.StatusOK, myDb.GetUserById(id))

}

//Get videos and return in JSON format
func getVideos(ctx *iris.Context) {
	ctx.Render("developerportal.html", struct{ Name string }{Name: "iris"})
}

/**************/
//Video API
/*************/

func main() {

	// Get theme
	iris.StaticServe("./static", "/assets") //Serve static assets

	iris.Get("/", developerPortal)

	iris.API("/videos", VideoAPI{}) //handles /videos endpoint
	iris.Listen(":8080")
}
