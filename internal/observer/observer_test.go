package observer_test

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/nokamoto/shogithink/internal/observer"
)

func getFreePort() int {
	// Listen on a random port and close immediately to get a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}

var _ = Describe("Observer", func() {
	It("should serve logs over HTTP", func(ctx SpecContext) {
		port := getFreePort()
		obs, err := observer.New(port)
		Expect(err).NotTo(HaveOccurred())
		DeferCleanup(func() { obs.Stop() })

		obs.Log("hello %s", "world")

		Eventually(func() string {
			resp, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
			if err != nil {
				return ""
			}
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			return string(body)
		}).WithContext(ctx).WithTimeout(10 * time.Second).Should(ContainSubstring("hello world"))
	})
})

func TestObserver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Observer Suite")
}
