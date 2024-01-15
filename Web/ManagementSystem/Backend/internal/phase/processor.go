package phase

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	m "course.project/management_system/internal/common/model"
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
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, StartCourseSelectionPhaseResponse{
			ErrorCode:    m.REQUEST_BODY_ERROR,
			ErrorMessage: m.REQUEST_BODY_ERROR_MSG,
		})
		return
	}
	if req.RoleID == "" || req.EndTime == 0 {
		c.JSON(http.StatusOK, StartCourseSelectionPhaseResponse{
			ErrorCode:    m.COMPULSORY_FIELD_MISSING,
			ErrorMessage: m.COMPULSORY_FIELD_MISSING_MSG,
		})
		return
	}
	if !role.IsAdmin(req.RoleID) {
		c.JSON(http.StatusOK, StartCourseSelectionPhaseResponse{
			ErrorCode:    m.PERMISSON_ERROR,
			ErrorMessage: m.PERMISSON_ERROR_MSG,
		})
		return
	}

	// to make sure the singleton was created
	GetCurrentPhase()
	lock.Lock()
	defer lock.Unlock()
	phaseSingleton.phaseType = COURSE_SELECTION_PHASE
	phaseSingleton.endTime = req.EndTime
	c.JSON(http.StatusOK, StartCourseSelectionPhaseResponse{
		ErrorCode:    0,
		ErrorMessage: "",
	})
}

// /api/v1/end_course_selection
