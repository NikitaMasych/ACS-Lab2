package client

import (
	"regexp"
	"server/errors"
	"sort"
	"strconv"
	"strings"
	"time"
)

var lastSortingDuration time.Duration = -1

func ProcessRequest(context string) string {
	switch {
	case context == "who":
		return whoInfo
	case context == "help":
		return helpInfo
	case context == "time":
		if lastSortingDuration == -1 {
			return noSortMeasurementInfo
		}
		return lastSortingDuration.String()
	case strings.HasPrefix(context, "sort "):
		startTime := time.Now()
		value, err := SortRequest(context)
		lastSortingDuration = time.Since(startTime)
		if err != nil {
			return err.Error()
		}
		return `"` + value + `"`
	default:
		return unknownInfo
	}
}

func SortRequest(context string) (string, error) {
	textStartIndex := strings.Index(context, "\"")
	textEndIndex := strings.LastIndex(context, "\"")
	if textStartIndex == -1 || textEndIndex == -1 {
		return "", errors.ErrNoMessage
	}
	text := context[textStartIndex+1 : textEndIndex]

	startIndex, err := fetchStartIndex(context[:textStartIndex])
	if err != nil {
		return "", err
	}
	length, err := fetchLength(context[:textStartIndex])
	if err != nil {
		return "", err
	}
	if err == nil && length == -1 {
		length = len(text) - startIndex
	}
	sortType, err := fetchSortType(context[:textStartIndex])
	if err != nil {
		return "", err
	}

	if startIndex < 0 || startIndex > len(text) {
		return "", errors.ErrInvalidStartIndex
	}
	if length < 0 || startIndex+length > len(text) {
		return "", errors.ErrInvalidLength
	}

	switch sortType {
	case descendingSortType:
		return sortDescending(text, startIndex, length), nil
	default:
		return sortAscending(text, startIndex, length), nil
	}
}

func fetchStartIndex(context string) (int, error) {
	startIndexCmd := strings.Index(context, "-s")
	if startIndexCmd == -1 {
		return 0, nil
	}

	const invalidStartIndex = -1
	decimalDigitRe := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`) // integer decimal number
	startIndex := decimalDigitRe.FindAllString(context[startIndexCmd+2:], 1)
	if len(startIndex) == 0 {
		return invalidStartIndex, errors.ErrNoStartIndex
	}
	return strconv.Atoi(startIndex[0])
}

func fetchLength(context string) (int, error) {
	const invalidLength = -1
	lengthCmd := strings.Index(context, "-l")
	if lengthCmd == -1 {
		return invalidLength, nil
	}

	decimalDigitRe := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`) // integer decimal number
	length := decimalDigitRe.FindAllString(context[lengthCmd+2:], 1)
	if len(length) == 0 {
		return invalidLength, errors.ErrInvalidLength
	}
	return strconv.Atoi(length[0])
}

const (
	ascendingSortType  = 1
	descendingSortType = -1
)

func fetchSortType(context string) (int, error) {
	sortTypeIndex := strings.Index(context, "-t")
	if sortTypeIndex == -1 {
		return ascendingSortType, nil
	}

	const invalidSortType = 0
	decimalDigitRe := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`) // integer decimal number
	sortType := decimalDigitRe.FindAllString(context[sortTypeIndex+2:], 1)
	if len(sortType) == 0 {
		return invalidSortType, errors.ErrNoSortType
	}
	switch sortType[0] {
	case "1":
		return ascendingSortType, nil
	case "-1":
		return descendingSortType, nil
	default:
		return invalidSortType, errors.ErrInvalidSortType
	}
}

func sortAscending(text string, startIndex, length int) string {
	characters := []byte(text)
	sortPart := characters[startIndex : startIndex+length]
	sort.Slice(sortPart, func(i, j int) bool {
		return sortPart[i] < sortPart[j]
	})
	return string(characters)
}

func sortDescending(text string, startIndex, length int) string {
	characters := []byte(text)
	sortPart := characters[startIndex : startIndex+length]
	sort.Slice(sortPart, func(i, j int) bool {
		return sortPart[i] > sortPart[j]
	})
	return string(characters)
}
