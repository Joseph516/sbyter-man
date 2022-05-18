package upload

import (
	"os/exec"
)

func ExactCoverFromVideo(pathVideo, pathImg string) error {
	command := "ffmpeg"
	frameExtractionTime := "0:00:00.001"
	vframes := "1"
	qv := "2"

	// create the command
	cmd := exec.Command(command,
		"-ss", frameExtractionTime,
		"-i", pathVideo,
		"-vframes", vframes,
		"-q:v", qv,
		pathImg)

	// run the command and don't wait for it to finish. waiting exec is run
	// fmt.Println(cmd.String())
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}
