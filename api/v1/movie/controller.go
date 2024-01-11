package movie

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mashbens/restfull2/api/common/obj"
	_response "github.com/mashbens/restfull2/api/common/response"
	"github.com/mashbens/restfull2/api/v1/movie/request"
	"github.com/mashbens/restfull2/api/v1/movie/resp"
	service "github.com/mashbens/restfull2/business/movie"
)

type MovieController struct {
	movieService service.MovieService
}

func NewMovieController(
	movieService service.MovieService,

) *MovieController {
	return &MovieController{
		movieService: movieService,
	}
}

var validate *validator.Validate

func (controller *MovieController) InsertMovie(c echo.Context) error {
	payload := new(request.MovieRequest)
	err := c.Bind(payload)

	validate = validator.New()
	err = validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
	if payload.Title == "" {
		response := _response.BuildErrorResponse("Bad Request", "title cannot be null")
		return c.JSON(http.StatusBadRequest, response)
	}
	res, err := controller.movieService.InsertMovie(request.ToService(*payload))
	if err != nil {
		fmt.Println(err)
	}
	data := resp.FromService(res)
	response := _response.BuildSuccsessResponse("Succses", data)
	return c.JSON(http.StatusOK, response)
}

func (controller *MovieController) FindAll(c echo.Context) error {
	res, _ := controller.movieService.FindAll()
	data := resp.FromeServiceSlice(res)
	response := _response.BuildSuccsessResponse("Succses", data)
	return c.JSON(http.StatusOK, response)
}

func (controller *MovieController) FindMovieById(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	res, err := controller.movieService.FindMovieById(intID)
	if err != nil {
		response := _response.BuildErrorResponse("Not Found", "Activity with ID "+id+" Not Found")
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromService(res)
	response := _response.BuildSuccsessResponse("Succses", data)
	return c.JSON(http.StatusOK, response)
}

func (controller *MovieController) UpdateMovieById(c echo.Context) error {
	payload := new(request.MovieRequest)
	err := c.Bind(payload)
	if payload.Title == "" {
		response := _response.BuildErrorResponse("Bad Request", "title cannot be null")
		return c.JSON(http.StatusBadRequest, response)
	}

	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	res, err := controller.movieService.UpdateMovieById(intID, request.ToService(*payload))
	if err != nil {
		response := _response.BuildErrorResponse("Not Found", "Activity with ID "+id+" Not Found")
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromService(res)
	response := _response.BuildSuccsessResponse("Succses", data)
	return c.JSON(http.StatusOK, response)
}

func (controller *MovieController) DeleteMovieById(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	err := controller.movieService.DeleteMovieById(intID)
	if err != nil {
		response := _response.BuildErrorResponse("Not Found", "Activity with ID "+id+" Not Found")
		return c.JSON(http.StatusBadRequest, response)
	}
	response := _response.BuildSuccsessResponse("Succses", obj.EmptyObj{})
	return c.JSON(http.StatusOK, response)
}
