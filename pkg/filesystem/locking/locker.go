package locking

import (
	"os"

	"github.com/pkg/errors"
)

// Locker provides file locking facilities.
type Locker struct {
	// The underlying file object to be locked.
	file *os.File
}

// NewLocker attempts to create a lock with the file at the specified path,
// creating the file if necessary. The lock is returned in an unlocked state.
func NewLocker(path string, permissions os.FileMode) (*Locker, error) {
	mode := os.O_RDWR | os.O_CREATE | os.O_APPEND
	if file, err := os.OpenFile(path, mode, permissions); err != nil {
		return nil, errors.Wrap(err, "unable to open lock file")
	} else {
		return &Locker{file: file}, nil
	}
}

// Close closes the file underlying the locker. This will release any lock held
// on the file and disable future locking. On POSIX platforms, this also
// releases other locks held on the same file.
func (l *Locker) Close() error {
	return l.file.Close()
}
