package url

import (
	"testing"
)

func TestURLEnsureValidNilInvalid(t *testing.T) {
	var invalid *URL
	if invalid.EnsureValid() == nil {
		t.Error("nil URL marked as valid")
	}
}

func TestURLEnsureValidLocalUsernameInvalid(t *testing.T) {
	invalid := &URL{
		User: "george",
		Path: "some/path",
	}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidLocalHostnameInvalid(t *testing.T) {
	invalid := &URL{
		Host: "somehost",
		Path: "some/path",
	}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidLocalPortInvalid(t *testing.T) {
	invalid := &URL{
		Port: 22,
		Path: "some/path",
	}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidLocalEmptyPathInvalid(t *testing.T) {
	invalid := &URL{}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidLocalEnvironmentVariablesInvalid(t *testing.T) {
	invalid := &URL{
		Path: "some/path",
		Environment: map[string]string{
			"key": "value",
		},
	}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidLocal(t *testing.T) {
	valid := &URL{Path: "some/path"}
	if err := valid.EnsureValid(); err != nil {
		t.Error("valid URL classified as invalid")
	}
}

func TestURLEnsureValidSSHEmptyHostnameInvalid(t *testing.T) {
	invalid := &URL{
		Protocol: Protocol_SSH,
		Path:     "some/path",
	}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidSSHLargePortInvalid(t *testing.T) {
	invalid := &URL{
		Protocol: Protocol_SSH,
		Host:     "washington",
		Port:     65536,
		Path:     "some/path",
	}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidSSHEmptyPathInvalid(t *testing.T) {
	invalid := &URL{
		Protocol: Protocol_SSH,
	}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidSSHEnvironmentVariablesInvalid(t *testing.T) {
	invalid := &URL{
		Protocol: Protocol_SSH,
		User:     "george",
		Host:     "washington",
		Port:     22,
		Path:     "~/path",
		Environment: map[string]string{
			"key": "value",
		},
	}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidSSH(t *testing.T) {
	valid := &URL{
		Protocol: Protocol_SSH,
		User:     "george",
		Host:     "washington",
		Port:     22,
		Path:     "~/path",
	}
	if err := valid.EnsureValid(); err != nil {
		t.Error("valid URL classified as invalid")
	}
}

func TestURLEnsureValidDockerPortInvalid(t *testing.T) {
	invalid := &URL{
		Protocol: Protocol_Docker,
		User:     "george",
		Host:     "washington",
		Port:     50,
		Path:     "~/path",
	}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidDockerEmptyPathInvalid(t *testing.T) {
	invalid := &URL{
		Protocol: Protocol_Docker,
		User:     "george",
		Host:     "washington",
		Path:     "",
	}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidDockerBadPathInvalid(t *testing.T) {
	invalid := &URL{
		Protocol: Protocol_Docker,
		User:     "george",
		Host:     "washington",
		Path:     "$path",
	}
	if invalid.EnsureValid() == nil {
		t.Error("invalid URL classified as valid")
	}
}

func TestURLEnsureValidDockerHomeRelativePath(t *testing.T) {
	valid := &URL{
		Protocol: Protocol_Docker,
		User:     "george",
		Host:     "washington",
		Path:     "~/path",
	}
	if err := valid.EnsureValid(); err != nil {
		t.Error("valid URL classified as invalid")
	}
}

func TestURLEnsureValidDockerUserRelativePath(t *testing.T) {
	valid := &URL{
		Protocol: Protocol_Docker,
		User:     "george",
		Host:     "washington",
		Path:     "~otheruser/path",
	}
	if err := valid.EnsureValid(); err != nil {
		t.Error("valid URL classified as invalid")
	}
}

func TestURLEnsureValidDockerWindowsPath(t *testing.T) {
	valid := &URL{
		Protocol: Protocol_Docker,
		User:     "george",
		Host:     "washington",
		Path:     `C:\path`,
	}
	if err := valid.EnsureValid(); err != nil {
		t.Error("valid URL classified as invalid")
	}
}

func TestURLEnsureValidDocker(t *testing.T) {
	valid := &URL{
		Protocol: Protocol_Docker,
		User:     "george",
		Host:     "washington",
		Path:     "/path",
	}
	if err := valid.EnsureValid(); err != nil {
		t.Error("valid URL classified as invalid")
	}
}
