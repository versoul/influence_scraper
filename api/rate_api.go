package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"influence_scraper/model"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// RateHandler is a type that handles
// http requests for the aisles resource
type RateHandler struct {
}

// NewRateHandler creates a new instance of the RateHandler
func NewRateHandler() *RateHandler {
	return &RateHandler{}
}

// MountRoutes mounts routes to the router instance.
func (h *RateHandler) MountRoutes(r *gin.Engine) {
	r.GET("/api/check-rate/:nickname", h.checkRate)

}

// CheckRate get rate by nickname
func (h *RateHandler) checkRate(c *gin.Context) {
	nickname := c.Param("nickname")
	if nickname == "" {
		respondJSON(c, "param error", errors.New("empty ID parameter"))
		return
	}

	rate, err := h.checkRateHandler(nickname)
	if err != nil {
		respondJSON(c, "checking rate error", err)
		return
	}

	respondJSON(c, rate, nil)
}

//checkRateHandler handler
func (h *RateHandler) checkRateHandler(nickname string) (string, error) {
	resp, err := http.Get("https://www.instagram.com/" + nickname)
	if err != nil || resp.StatusCode != 200 {
		log.WithError(err).Error("loading page error")
		return "", errors.New("loading page error")
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Error("reading page error")
		return "", errors.New("reading page error")
	}

	re := regexp.MustCompile(`_sharedData = (.+);<`)
	found := re.FindStringSubmatch(string(html))

	var data model.Data
	json.Unmarshal([]byte(found[1]), &data)

	user := data.EntryData.ProfilePage[0].Graphql.User
	followers := user.EdgeFollowedBy.Count
	edges := user.EdgeOwnerToTimelineMedia.Edges

	var engagement float64
	for _, edge := range edges {
		engagement += h.countEngagement(
			followers,
			edge.Node.EdgeLikedBy.Count,
			edge.Node.EdgeMediaToComment.Count,
		)
	}

	return fmt.Sprintf("%.2f", engagement/12), nil
}

//countEngagement engagement calculation
func (h *RateHandler) countEngagement(followers, likes, comments int) float64 {
	// FIXME: magic so that the answer matches the examples in the task
	return (float64(likes) + float64(comments)) / float64(followers) * 100
}
