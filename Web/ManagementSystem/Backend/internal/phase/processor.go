package phase

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"course.project/management_system/internal/role"
)

const (
	NORMAL_PHASE           = 0
	COURSE_SELECTION_PHASE = 1
)

type CurrentPhase struct {
	phaseType int
	endTime   int64
}

var phaseSingleton *CurrentPhase
var lock = &sync.Mutex{}

func GetCurrentPhase() *CurrentPhase {
	lock.Lock()
	defer lock.Unlock()

	if phaseSingleton == nil {
		phaseSingleton = &CurrentPhase{
			phaseType: NORMAL_PHASE,
		}
	} else if phaseSingleton.phaseType == COURSE_SELECTION_PHASE &&
		phaseSingleton.endTime <= time.Now().Unix() {
		phaseSingleton.phaseType = NORMAL_PHASE
		phaseSingleton.endTime = 0
	}
	return phaseSingleton
}

// /api/v1/start_course_selection
func StartCourseSelectionPhase(c *gin.Context) {
	var req StartCourseSelectionPhaseRequest
	if c.ShouldBind(&req) != nil {
		c.JSON(http.StatusOK, StartCourseSelectionPhaseResponse{
			ErrorCode:    REQUEST_BODY_ERROR,
			ErrorMessage: REQUEST_BODY_ERROR_MSG,
		})
	}
	if !role.IsAdmin(req.RoleID) {
		c.JSON(http.StatusOK, StartCourseSelectionPhaseResponse{
			ErrorCode:    PERMISSON_ERROR,
			ErrorMessage: PERMISSON_ERROR_MSG,
		})
	}

	// to make sure the singleton was created
	GetCurrentPhase()
	lock.Lock()
	defer lock.Unlock()
	phaseSingleton.phaseType = COURSE_SELECTION_PHASE
	phaseSingleton.endTime = req.EndTime
}

// /api/v1/end_course_selection
