package a

import (
	"strconv"

	"github.com/pkg/errors"
)

func TestSingleError_Success() error {
	if _, err := strconv.Atoi("1"); err != nil {
		return errors.Wrap(err, "failed to strconv.Atoi")
	}

	return nil
}

func TestSingleError_Failed() error {
	if _, err := strconv.Atoi("1"); err != nil {
		return errors.Wrap(err, "strconv.Atoi failed") // want "The prefix of the error message should be 'failed to ...'"
	}

	return nil
}

func TestSingleError_Not_target() error {
	if _, err := strconv.Atoi("1"); err != nil {
		return errors.New("failed to strconv.Atoi")
	}

	return nil
}

func TestMultipleError_Success() (bool, error) {
	if _, err := strconv.Atoi("1"); err != nil {
		return false, errors.Wrap(err, "failed to strconv.Atoi")
	}

	return true, nil
}

func TestMultipleError_Failed() (bool, error) {
	if _, err := strconv.Atoi("1"); err != nil {
		return false, errors.Wrap(err, "strconv.Atoi failed") // want "The prefix of the error message should be 'failed to ...'"
	}

	return true, nil
}

func TestMultipleError_Not_target() (bool, error) {
	if _, err := strconv.Atoi("1"); err != nil {
		return false, errors.New("failed to strconv.Atoi")
	}

	return true, nil
}
