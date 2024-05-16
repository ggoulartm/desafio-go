package MergeFFMPEG

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Defines the Video Structure parameters
type Video struct {
	filename string
	duration time.Duration
	scale    scale
	audio    bool
}

// Scale Parameters
type scale struct {
	width, height int
}

var (
	CIF  scale = scale{width: 352, height: 240}
	QVGA scale = scale{width: 320, height: 240}
	HD1  scale = scale{width: 480, height: 352}
	D1   scale = scale{width: 720, height: 480}
	//h960  scale = scale{width: 960, height: 480}
	FullHD scale = scale{width: 1920, height: 1080}
	//p960  scale = scale{width: 1280, height: 960}
	HD scale = scale{width: 1280, height: 720}
)

func NewVideo(filename string) *Video {
	sc, err := getVideoScale(filename)
	if err != nil {
		panic(err)
	}
	hasAudio, err := getAudioVideo(filename)
	if err != nil {
		panic(err)
	}
	duration, err := getVideoDuration(filename)
	if err != nil {
		panic(err)
	}
	return &Video{
		filename: filename,
		scale:    sc,
		audio:    hasAudio,
		duration: *duration,
	}
}

func MergeFFMPEG(videos []Video, fileName string, mode int, pad scale) (merged *Video) {
	var dur time.Duration
	var command string = " -progress " + WriteProgressPath(fileName, "merge")
	var withoutAudio bool = true
	var filtercomplex, filtercomplexStreamOutput string
	var i int
	for _, video := range videos {
		if video.filename[0] == 'c' {
			command = command + " -f lavfi"
		} else {
			hasAudio, err := getAudioVideo(video.filename)
			if err != nil {
				panic(err)
			}
			video.audio = hasAudio
		}
		command = command + " -i " + video.filename
		if video.filename[0] == 'c' && video.audio {
			filtercomplex = filtercomplex + fmt.Sprintf("[%d:v]", i) +
				fmt.Sprintf("setpts=PTS-STARTPTS,scale=%d:%d:force_original_aspect_ratio=decrease:eval=frame,pad=%d:%d:-1:-1:color=black[v%d];", pad.width, pad.height, pad.width, pad.height, i)
			filtercomplexStreamOutput = filtercomplexStreamOutput + fmt.Sprintf("[v%d][%d:a]", i, i+1)
			i = i + 2
			withoutAudio = false
		} else if video.filename[0] == 'c' && !video.audio {
			filtercomplex = filtercomplex + fmt.Sprintf("[%d:v]", i) +
				fmt.Sprintf("setpts=PTS-STARTPTS,scale=%d:%d:force_original_aspect_ratio=decrease:eval=frame,pad=%d:%d:-1:-1:color=black[v%d];", pad.width, pad.height, pad.width, pad.height, i)
			filtercomplexStreamOutput = filtercomplexStreamOutput + fmt.Sprintf("[v%d]", i)
			i++
		} else {
			filtercomplex = filtercomplex + fmt.Sprintf("[%d:v]", i) +
				fmt.Sprintf("setpts=PTS-STARTPTS,scale=%d:%d:force_original_aspect_ratio=decrease:eval=frame,pad=%d:%d:-1:-1:color=black[v%d];", pad.width, pad.height, pad.width, pad.height, i)
			filtercomplexStreamOutput = filtercomplexStreamOutput + fmt.Sprintf("[v%d][%d:a]", i, i)
			i++
			withoutAudio = false
		}

		dur += video.duration
	}
	filtercomplex = filtercomplex + filtercomplexStreamOutput
	command = command + " -filter_complex"
	var mapa string
	if !withoutAudio {
		filtercomplex = filtercomplex + "concat=n=" + fmt.Sprintf("%d", len(videos)) + ":v=1:a=1[v][a]"
		mapa = "-preset ultrafast -vsync 2 -map [v] -map [a] -c:v libx264 " + fileName + " -y"
	} else {
		filtercomplex = filtercomplex + "concat=n=" + fmt.Sprintf("%d", len(videos)) + ":v=1[v]"
		mapa = "-preset ultrafast -vsync 2 -map [v] -c:v libx264 " + fileName + " -y"
	}
	args := strings.Split(command, " ")
	maps := strings.Split(mapa, " ")
	args = append(args, filtercomplex)
	args = append(args, maps...)
	fmt.Println(args)
	procAttr := new(os.ProcAttr)

	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	process, err := os.StartProcess("/bin/ffmpeg", args, procAttr)
	if err != nil {
		panic(err)
	}
	state, err := process.Wait()
	fmt.Println("exit status code: ", state.ExitCode())
	if state.ExitCode() == 1 {
		panic(state)
	}
	if err != nil {
		panic(err)
	}
	return &Video{
		filename: fileName,
		duration: dur,
		scale: scale{
			width:  videos[0].scale.width,
			height: videos[0].scale.height,
		},
	}
}

func WriteProgressPath(filename, filter string) (progressPath string) {
	split := strings.Split(filename, "/")
	for _, I := range split[1 : len(split)-2] {
		progressPath = progressPath + "/" + I
	}
	progressPath = progressPath + "/" + filter + "ProgressInfo.txt"
	return progressPath
}

func getAudioVideo(filename string) (hasAudio bool, err error) {
	cmd := exec.Command("ffprobe",
		"-i",
		filename,
		"-show_streams",
		"-select_streams",
		"a",
		"-loglevel",
		"error")

	//running command and streamming output for scale variable
	Audio, err := cmd.Output()
	if err != nil {
		err = fmt.Errorf("getAudioVideo, error: %s", err)
		return false, err
	}
	hasAudio = len(Audio) != 0
	return hasAudio, err
}

func getVideoDuration(filename string) (*time.Duration, error) {
	cmd := exec.Command("ffprobe",
		"-i",
		filename,
		"-v",
		"error",
		"-select_streams",
		"v:0",
		"-show_entries",
		"stream=duration",
		"-of",
		"default=nw=1:nk=1")

	//running command and streamming output for duration variable
	Duration, err := cmd.Output()
	if err != nil {
		err = fmt.Errorf("getVideoDuration, error: %s", err)
		return nil, err
	}
	duration := time.Duration(Duration[0])
	return &duration, err
}

func getVideoScale(filename string) (scale, error) {
	cmd := exec.Command("ffprobe",
		"-i",
		filename,
		"-v",
		"error",
		"-select_streams",
		"v:0",
		"-show_entries",
		"stream=width,height",
		"-of",
		"default=nw=1:nk=1",
	)

	//running command and streamming output for scale variable
	scaleOut, err := cmd.Output()
	//fmt.Println(string(scaleOut))
	if err != nil {
		return scale{}, err
	}
	//extracting the data from command's output
	sc := strings.Split(string(scaleOut), "\n")
	//converting data from string to int
	width, err := strconv.Atoi(sc[0])
	if err != nil {
		return scale{}, fmt.Errorf("getVideoScale, error2: %s", err)
	}
	height, err := strconv.Atoi(sc[1])
	if err != nil {
		return scale{}, fmt.Errorf("getVideoScale, error3: %s", err)
	}

	return scale{
		width:  width,
		height: height,
	}, nil
}

// An object from nullsrc to delivery a video
// The purpose is to replace a gap on video merging or
// An entire video gap on mosaic
func NullVideoBuilder(video Video) *Video {
	fileName := "color=s=" + fmt.Sprintf("%dx%d:d=%d:r=30:c=black",
		video.scale.width,
		video.scale.height,
		video.duration)
	if video.audio {
		return nullAudioVideoBuilder(video)
	}
	return &Video{filename: fileName,
		duration: video.duration,
		scale:    video.scale,
		audio:    false,
	}
}

func nullAudioVideoBuilder(video Video) *Video { //it'll be yellow until merge with audio and black screens be implemented
	fileName := "color=s=" + fmt.Sprintf("%dx%d:d=%d:r=30:c=black -f lavfi -i anullsrc=cl=mono:d=%d",
		video.scale.width,
		video.scale.height,
		video.duration,
		video.duration)
	return &Video{filename: fileName,
		duration: video.duration,
		scale:    video.scale,
		audio:    true,
	}
}
