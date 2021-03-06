package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Recommendation struct {
	ID          int       `json:"ID" bson:"ID"`
	Usr         *User     `json:"User" bson:"User"`
	Stck        *Stock    `json:"Stock" bson:"Stock"`
	Mt          *Meet     `json:"Meet" bson:"Meet"`
	Created     time.Time `json:"Created" bson:"Created"`
	LastUpdated time.Time `json:"LastUpdated" bson:"LastUpdated"`
	URL         string    `json:"URL" bson:"URL"`
}

type User struct {
	ID          int       `json:"ID" bson:"ID"`
	Name        string    `json:"Name" bson:"Name"`
	Phone       string    `json:"Phone" bson:"Phone"`
	Mail        string    `json:"Mail" bson:"Mail"`
	Score       int       `json:"Score" bson:"Score"`
	Created     time.Time `json:"Created" bson:"Created"`
	LastUpdated time.Time `json:"LastUpdated" bson:"LastUpdated"`
	URL         string    `json:"URL" bson:"URL"`
}

type Stock struct {
	ID                 int       `json:"ID" bson:"ID"`
	Name               string    `json:"Name" bson:"Name"`
	Symbol             string    `json:"Symbol" bson:"Symbol"`
	LastTradePriceOnly float64   `json:"LastTradePriceOnly" bson:"LastTradePriceOnly"`
	Change             string    `json:"Change" bson:"Change"`
	BuyPrice           float64   `json:"BuyPrice" bson:"BuyPrice"`
	NumberOfShares     float64   `json:"NumberOfShares" bson:"NumberOfShares"`
	Color              string    `json:"Color" bson:"Color"`
	Created            time.Time `json:"created" bson:"Created"`
	SalesPrice         float64   `json:"SalesPrice" bson:"SalesPrice"`
	LastUpdated        time.Time `json:"LastUpdated" bson:"LastUpdated"`
	URL                string    `json:"URL" bson:"URL"`
}

type Meet struct {
	ID          int       `json:"ID" bson:"ID"`
	Location    string    `json:"Location" bson:"Location"`
	Date        time.Time `json:"Date" bson:"Date"`
	Text        string    `json:"Text" bson:"Text"`
	Created     time.Time `json:"Created" bson:"Created"`
	LastUpdated time.Time `json:"LastUpdated" bson:"LastUpdated"`
	URL         string    `json:"URL" bson:"URL"`
	IDUser      int       `json:"IDuser" bson:"IDuser"`
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081" //GG
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "static")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.GET("/start", allRecs)
	router.GET("/users/single/:id", userRecs)
	router.GET("/users/leaderboard", GetLeaderboardEndpoint)
	router.GET("/meet/:id", meetRecs)
	router.GET("/meets", GetMeetsEndpoint)

	router.Run(":" + port)

}

//ginFunc returns a gin context
func allRecs(c *gin.Context) {

	request := "https://rocky-wildwood-11035.herokuapp.com/recommendations"
	resp, err := http.Get(request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	defer resp.Body.Close()
	var recs []*Recommendation

	if err := json.NewDecoder(resp.Body).Decode(&recs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	layoutData := struct {
		ThreadID int
		Posts    []*Recommendation
	}{
		ThreadID: 1,
		Posts:    recs,
	}

	c.HTML(http.StatusOK, "start.html", layoutData)

}

func userRecs(c *gin.Context) {

	symbol := c.Param("id")
	request := "https://rocky-wildwood-11035.herokuapp.com/recommendations/user/" + symbol
	resp, err := http.Get(request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	defer resp.Body.Close()
	var recs []*Recommendation

	if err := json.NewDecoder(resp.Body).Decode(&recs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	layoutData := struct {
		ThreadID int
		Posts    []*Recommendation
	}{
		ThreadID: 1,
		Posts:    recs,
	}

	c.HTML(http.StatusOK, "userrec.html", layoutData)

}

func meetRecs(c *gin.Context) {

	symbol := c.Param("id")
	request := "https://rocky-wildwood-11035.herokuapp.com/recommendations/meet/" + symbol
	resp, err := http.Get(request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	defer resp.Body.Close()
	var recs []*Recommendation

	if err := json.NewDecoder(resp.Body).Decode(&recs); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	layoutData := struct {
		ThreadID int
		Posts    []*Recommendation
	}{
		ThreadID: 1,
		Posts:    recs,
	}

	c.HTML(http.StatusOK, "meetrec.html", layoutData)

}

func GetMeetsEndpoint(c *gin.Context) {

	request := "https://rocky-wildwood-11035.herokuapp.com/meet"
	resp, err := http.Get(request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	defer resp.Body.Close()
	var meets []*Meet

	if err := json.NewDecoder(resp.Body).Decode(&meets); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	layoutData := struct {
		ThreadID int
		Posts    []*Meet
	}{
		ThreadID: 1,
		Posts:    meets,
	}

	c.HTML(http.StatusOK, "meets.html", layoutData)

}

func GetLeaderboardEndpoint(c *gin.Context) {

	request := "https://rocky-wildwood-11035.herokuapp.com/users/leaderboard"
	resp, err := http.Get(request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	defer resp.Body.Close()
	var users []*User

	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	layoutData := struct {
		ThreadID int
		Posts    []*User
	}{
		ThreadID: 1,
		Posts:    users,
	}

	c.HTML(http.StatusOK, "leaderboard.html", layoutData)

}
