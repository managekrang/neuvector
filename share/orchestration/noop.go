package orchestration

import (
	"net"
	"strings"

	"github.com/codeskyblue/go-sh"

	"github.com/neuvector/neuvector/share"
	"github.com/neuvector/neuvector/share/container"
	sk "github.com/neuvector/neuvector/share/system/sidekick"
)

func shell(cmd string) ([]byte, error) {
	// log.Printf("shell: %v\n", cmd)
	c := strings.Split(cmd, " ")
	return sh.Command(c[0], c[1:]).Output()
}

func shellCombined(cmd string) ([]byte, error) {
	// log.Printf("shell: %v\n", cmd)
	c := strings.Split(cmd, " ")
	return sh.Command(c[0], c[1:]).CombinedOutput()
}

// --

type noop struct {
	platform, flavor, network string
}

func (d *noop) GetVersion() (string, string) {
	return "", ""
}

func (d *noop) GetServiceSubnet(envs []string) *net.IPNet {
	return nil
}

func (d *noop) GetService(meta *container.ContainerMeta) *Service {
	return nil
}

func (d *noop) GetPlatformRole(m *container.ContainerMeta) (string, bool) {
	return "", true
}

func (d *noop) GetDomain(labels map[string]string) string {
	return ""
}

func (d *noop) SetIPAddrScope(ports map[string][]share.CLUSIPAddr,
	meta *container.ContainerMeta, nets map[string]*container.Network,
) {
	return
}

func (d *noop) GetHostTunnelIP(links map[string]sk.NetIface) []net.IPNet {
	return nil
}

func (d *noop) IgnoreConnectFromManagedHost() bool {
	return true
}

func (d *noop) ConsiderHostsAsInternal() bool {
	return true
}

func (d *noop) ApplyPolicyAtIngress() bool {
	return false
}

func (d *noop) SupportKubeCISBench() bool {
	return false
}

func (d *noop) CleanupHostPorts(hostPorts map[string][]share.CLUSIPAddr) error {
	return ErrMethodNotSupported
}
