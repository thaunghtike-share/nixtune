package signatures

import (
	"testing"
)

func TestNginx_workerProcessesRegex(t *testing.T) {
	n := &Nginx{}
	re := n.workerProcessesRegex()

	if len(re.FindAllStringSubmatch("worker_processes 1;", -1)) != 2 {
		t.Errorf("Invalid number of worker_processes")
	}

	if len(re.FindAllStringSubmatch("worker_processes auto;", -1)) != 2 {

		t.Errorf("Invalid number of worker_processes")
	}

	if len(re.FindAllStringSubmatch("worker_processes 1", -1)) > 0 {

		t.Errorf("Invalid number of worker_processes")
	}

	if len(re.FindAllStringSubmatch("none", -1)) > 0 {

		t.Errorf("Invalid number of worker_processes")
	}
}
