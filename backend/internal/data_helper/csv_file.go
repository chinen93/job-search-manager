package datahelper

import (
	"encoding/csv"
	"errors"
	"io"
	jobmodel "job-search-manager/internal/jobs/models"
	"os"
	"path/filepath"
)

var (
	POSITIONS_FILENAME = "./data/positions.csv"
	EXPECTED_HEADERS   = []string{"Status", "Date", "Company", "Job Title", "ID", "Description"}
)

func isCVSFileValid(reader *csv.Reader) bool {

	header, err := reader.Read()

	if err != nil {
		return false
	}

	if !Equal(header, EXPECTED_HEADERS) {
		return false
	}

	return true
}

func handleCSVRows(reader *csv.Reader) ([]*jobmodel.Position, error) {

	positions := []*jobmodel.Position{}

	for {
		row, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return positions, err
		}

		position := &jobmodel.Position{}

		position.Status = SanitizeString(row[0])
		position.Date = SanitizeString(row[1])
		position.Company = SanitizeString(row[2])
		position.Title = SanitizeString(row[3])
		position.ID = SanitizeString(row[4])
		position.Description = SanitizeString(row[5])

		positions = append(positions, position)
	}
}

func parsePositions() ([]*jobmodel.Position, error) {
	f, err := os.Open(POSITIONS_FILENAME)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)

	if !isCVSFileValid(csvr) {
		return nil, errors.New("could not read CSV file")
	}

	return handleCSVRows(csvr)
}

func setPositionsFilepath(newPath string) {
	POSITIONS_FILENAME, _ = filepath.Abs(newPath)
}
