package logging_test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"git.intra.weibo.com/adx/logging"
)

func checkLogFiles(l *logging.LogFile) (int, error) {
	// Extract the file extention
	fileExt := filepath.Ext(l.FileName)
	// If we have no file extension we append .log
	if fileExt == "" {
		fileExt = ".log"
	}
	// Remove the file extention from the filename
	fileName := strings.TrimSuffix(l.FileName, fileExt)
	// format -> /dir/logfilename-*.log
	pattern := fmt.Sprintf("%s/%s-*%s", l.LogPath, fileName, fileExt)
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		return 0, err
	}
	return len(filenames), nil
}

func TestRotateFileScanDisk(t *testing.T) {
	logFile := &logging.LogFile{
		FileName:           "adx.log",
		LogPath:            "/tmp/test-adx-logging",
		FileNameDateFormat: "20060102.150405",
		FileNameDateAlign:  true,
		RotationDuration:   logging.Duration{1 * time.Second},
		RotationCount:      2,
	}
	if err := logFile.Validate(); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		fmt.Fprintln(logFile, i)
		time.Sleep(500 * time.Millisecond) // 0.5s
	}
}
