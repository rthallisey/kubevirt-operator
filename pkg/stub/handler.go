package stub

import (
	"context"
	"os/exec"
	"strings"

	"github.com/rthallisey/kubevirt-operator/pkg/apis/app/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch event.Object.(type) {
	case *v1alpha1.App:
		// Create kubevirt manifest using the client
		cmd := exec.Command("kubectl", "create", "-f", "/etc/kubevirt/kubevirt.yaml")
		// Error is outputed in plain text in out
		out, _ := cmd.CombinedOutput()
		if strings.Contains(string(out), "Error from server (AlreadyExists)") {
			logrus.Debugf("Resources from kubevirt.yaml already exist")
		} else {
			logrus.Infof(string(out))
		}
	}
	return nil
}
