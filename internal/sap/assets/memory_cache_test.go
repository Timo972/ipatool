package assets

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("in-process asset cache", func() {
	It("reuses the bundle loaded earlier from the same cache directory", func() {
		directory, err := cacheDirectory()
		Expect(err).NotTo(HaveOccurred())

		loaded.mu.Lock()
		loaded.directory, loaded.bundle = directory, Bundle{AppleMediaServices: []byte("loaded")}
		loaded.mu.Unlock()
		DeferCleanup(func() {
			loaded.mu.Lock()
			loaded.directory, loaded.bundle = "", Bundle{}
			loaded.mu.Unlock()
		})

		bundle, err := Load(context.Background())

		Expect(err).NotTo(HaveOccurred())
		Expect(bundle.AppleMediaServices).To(Equal([]byte("loaded")))
	})
})
