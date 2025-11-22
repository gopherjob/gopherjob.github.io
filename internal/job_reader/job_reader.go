package jobReader

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/milad-rasouli/bluejob/internal/entity"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var (
	remoteKeywords = [...]string{
		"remote",
	}

	relocationKeywords = [...]string{
		"visa",
		"relocation",
		"work permit",
		"relocate",
		"sponsorship",
		"relocation",
	}
)

var (
	ErrPathNotFound = errors.New("path not found")
)

type JobReader interface {
	ReadAllFiles(path string) ([]*entity.Job, error)
}

type jobReader struct {
	logger *slog.Logger
}

func NewJobReader(logger *slog.Logger) JobReader {
	return &jobReader{logger: logger}
}

func (j *jobReader) ReadAllFiles(path string) ([]*entity.Job, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrPathNotFound, path)
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("expected directory, got file: %s", path)
	}

	var allJobs []*entity.Job

	err = filepath.Walk(path, func(p string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(p) != ".csv" {
			return nil
		}

		j.logger.Info("reading CSV file", slog.String("file", p))

		f, err := os.Open(p)
		if err != nil {
			return fmt.Errorf("failed opening file %s: %w", p, err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				j.logger.Warn("failed closing file", slog.String("file", p))
			}
		}()

		r := csv.NewReader(f)
		records, err := r.ReadAll()
		if err != nil {
			return fmt.Errorf("failed reading CSV %s: %w", p, err)
		}

		if len(records) <= 1 {
			return nil
		}

		// iterate rows, skip header
		for _, row := range records[1:] {
			if len(row) < 21 {
				continue
			}

			desc := strings.ToLower(row[19])
			tit := strings.ToLower(row[4])

			//j.logger.Debug("job info",row[0], row[2], j.containsAny(desc, remoteKeywords[:]), j.containsAny(tit, remoteKeywords[:]))
			remote := j.containsAny(desc, remoteKeywords[:]) || j.containsAny(tit, remoteKeywords[:])
			relocate := j.containsAny(desc, relocationKeywords[:]) || j.containsAny(tit, relocationKeywords[:])

			job := &entity.Job{
				ID:              row[0],
				URL:             row[2],
				Title:           row[4],
				Company:         row[5],
				Location:        row[6],
				Type:            row[8],
				Level:           row[15],
				Description:     row[19],
				CompanyIndustry: row[20],
				Remote:          remote,
				Relocation:      relocate,
			}

			allJobs = append(allJobs, job)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return j.filter(allJobs), nil
}

func (j *jobReader) filter(jobs []*entity.Job) []*entity.Job {
	filtered := make([]*entity.Job, 0, len(jobs))

	for _, job := range jobs {
		if job.Remote || job.Relocation {
			filtered = append(filtered, job)
		}
	}

	return filtered
}

func (j *jobReader) containsAny(input string, list []string) bool {
	for _, kw := range list {
		if strings.Contains(input, kw) {
			return true
		}
	}
	return false
}
