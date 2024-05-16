package main

import (
	v "MergeFilter/MergeFFMPEG"
	"fmt"
	"time"
)

func main() {
	outPut := "/home/ggmuller/Projects/output"
	teste3(outPut)
}

type TimeMesure struct {
	d          time.Duration
	method     string
	resolution string
}

func (t *TimeMesure) String() string {
	return fmt.Sprintf("Method: %s, Output Resolution: %s, Total Duration: %s", t.method, t.resolution, t.d)
}

func teste1(outPut string) {
	filename1 := "/home/ggmuller/Downloads/merge-0C54FA312DD2A020-2.mp4"
	filename2 := "/home/ggmuller/Downloads/teste1.mp4"
	v1 := v.NewVideo(filename1)
	v2 := v.NewVideo(filename2)
	Videos := []v.Video{}
	Videos = append(Videos, *v1, *v2)

	t1 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"1.mp4", 0, v.QVGA))
	D1 := time.Since(t1)
	t2 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"2.mp4", 0, v.D1))
	D2 := time.Since(t2)
	t3 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"3.mp4", 0, v.HD))
	D3 := time.Since(t3)
	t4 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"4.mp4", 0, v.FullHD))
	D4 := time.Since(t4)

	t5 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"5.mp4", 1, v.QVGA))
	D5 := time.Since(t5)
	t6 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"6.mp4", 1, v.D1))
	D6 := time.Since(t6)
	t7 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"7.mp4", 1, v.HD))
	D7 := time.Since(t7)
	t8 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"8.mp4", 1, v.FullHD))
	D8 := time.Since(t8)

	timemeasure := []TimeMesure{
		{d: D1, method: "mode 0", resolution: "QVGA"},
		{d: D2, method: "mode 0", resolution: "D1"},
		{d: D3, method: "mode 0", resolution: "HD"},
		{d: D4, method: "mode 0", resolution: "FullHD"},
		{d: D5, method: "mode 1", resolution: "QVGA"},
		{d: D6, method: "mode 1", resolution: "D1"},
		{d: D7, method: "mode 1", resolution: "HD"},
		{d: D8, method: "mode 1", resolution: "FullHD"},
	}

	for _, tm := range timemeasure {
		fmt.Println(tm.String())
	}
}

func teste2(outPut string) {
	intput := "/home/ggmuller/Projects/Projects/tests/MergeFilter/videos/video-part"
	var videos []string
	for i := 1; i < 6; i++ {
		videos = append(videos, intput+fmt.Sprintf("%d.mp4", i))
	}
	video1 := v.NewVideo(videos[0])
	video2 := v.NewVideo(videos[1])
	video3 := v.NewVideo(videos[2])
	video4 := v.NewVideo(videos[3])
	video5 := v.NewVideo(videos[4])
	Videos := []v.Video{}
	Videos = append(Videos, *video1, *video2, *video3, *video4, *video5)

	t1 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"1.mp4", 0, v.QVGA))
	D1 := time.Since(t1)
	t2 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"2.mp4", 0, v.D1))
	D2 := time.Since(t2)
	t3 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"3.mp4", 0, v.HD))
	D3 := time.Since(t3)
	t4 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"4.mp4", 0, v.FullHD))
	D4 := time.Since(t4)
	t5 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"5.mp4", 1, v.QVGA))
	D5 := time.Since(t5)
	t6 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"6.mp4", 1, v.D1))
	D6 := time.Since(t6)
	t7 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"7.mp4", 1, v.HD))
	D7 := time.Since(t7)
	t8 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"8.mp4", 1, v.FullHD))
	D8 := time.Since(t8)
	t9 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"9.mp4", -1, v.QVGA))
	D9 := time.Since(t9)

	timemeasure := []TimeMesure{
		{d: D1, method: "Old Scale Filter", resolution: "QVGA"},
		{d: D2, method: "Old Scale Filter", resolution: "D1"},
		{d: D3, method: "Old Scale Filter", resolution: "HD"},
		{d: D4, method: "Old Scale Filter", resolution: "FullHD"},
		{d: D5, method: "New Scale Filter", resolution: "QVGA"},
		{d: D6, method: "New Scale Filter", resolution: "D1"},
		{d: D7, method: "New Scale Filter", resolution: "HD"},
		{d: D8, method: "New Scale Filter", resolution: "FullHD"},
		{d: D9, method: "Without Scale Filter", resolution: "2K"},
	}

	for _, tm := range timemeasure {
		fmt.Println(tm.String())
	}
}

func teste3(outPut string) {
	intput := "/home/ggmuller/Projects/Projects/tests/MergeFilter/videos/video-part"
	var videos []string
	for i := 1; i < 6; i++ {
		videos = append(videos, intput+fmt.Sprintf("%d.mp4", i))
	}
	video1 := v.NewVideo(videos[0])
	video2 := v.NullVideoBuilder(*video1)
	video3 := v.NewVideo(videos[2])
	video4 := v.NewVideo(videos[3])
	video5 := v.NewVideo(videos[4])
	Videos := []v.Video{}
	Videos = append(Videos, *video1, *video2, *video3, *video4, *video5)
	fmt.Println(Videos)

	t1 := time.Now()
	fmt.Println(v.MergeFFMPEG(Videos, outPut+"1.mp4", 0, v.QVGA))
	D1 := time.Since(t1)

	fmt.Println("Time of execution: ", D1)
}
