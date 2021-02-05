package log

import (
	"github.com/pkg/errors"
	"log"
)

// OriginalError return original error
func OriginalError() error {
	return errors.New("error occurred")
}

// PassThroughError invoke and encapsulate OriginalError
func PassThroughError() error {
	err := OriginalError()
	// no need err check, because it work right even though nil there
	return errors.Wrap(err, "in passthrougherror")
}

func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// record any err occur.
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}
