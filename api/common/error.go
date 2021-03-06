package common

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nccgroup/tracy/api/store"
	"github.com/nccgroup/tracy/api/types"
	"github.com/nccgroup/tracy/log"
)

// ServerError is the common function for logging an internal server error and
// serving back something generic.
func ServerError(err error) []byte {
	ref := rand.Intn(100000)
	errs := types.Error{
		ErrorID:  uint(ref),
		ErrorMsg: err.Error(),
	}

	if errd := store.DB.Create(&errs).Error; errd != nil {
		log.Error.Print(errd)
	}

	return []byte(fmt.Sprintf(`{"Message":"Internal Server Error", "Reference":"%d"}`, ref))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
