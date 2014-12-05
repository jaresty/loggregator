package elector_test

import (
	"errors"
	"github.com/cloudfoundry/gosteno"
	"github.com/cloudfoundry/storeadapter"
	"github.com/cloudfoundry/storeadapter/fakestoreadapter"
	"syslog_drain_binder/elector"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Elector", func() {
	var fakeStore *fakestoreadapter.FakeStoreAdapter
	var logger *gosteno.Logger
	var testingSink *gosteno.TestingSink

	BeforeEach(func() {
		gosteno.EnterTestMode()
		testingSink = gosteno.GetMeTheGlobalTestSink()

		fakeStore = fakestoreadapter.New()
		logger = gosteno.NewLogger("test")
	})

	Context("at initialization", func() {
		It("connects to the store", func() {
			elector.NewElector("name", fakeStore, 1*time.Millisecond, logger)
			Expect(fakeStore.DidConnect).To(BeTrue())
		})

		Context("when store connection fails", func() {
			BeforeEach(func() {
				fakeStore.ConnectErr = errors.New("connection error")
			})

			It("logs an error", func() {
				go elector.NewElector("name", fakeStore, 100*time.Millisecond, logger)

				Eventually(testingSink.Records).Should(HaveLen(1))
				var messages []string
				for _, record := range testingSink.Records() {
					messages = append(messages, record.Message)
				}

				Expect(messages).To(ContainElement("Unable to connect to store: 'connection error'"))
			})

			It("reconnects on an interval", func() {
				go elector.NewElector("name", fakeStore, 10*time.Millisecond, logger)

				Eventually(testingSink.Records).Should(HaveLen(2))
			})
		})
	})

	Describe("RunForElection", func() {
		var candidate *elector.Elector

		BeforeEach(func() {
			candidate = elector.NewElector("name", fakeStore, time.Second, logger)
		})

		It("makes a bid to become leader", func() {
			err := candidate.RunForElection()
			Expect(err).NotTo(HaveOccurred())

			node, err := fakeStore.Get("syslog_drain_binder/leader")
			Expect(err).NotTo(HaveOccurred())
			Expect(node.Value).To(BeEquivalentTo("name"))
			Expect(node.TTL).To(Equal(uint64(1)))
		})

		It("re-attempts on an interval if key already exists", func() {
			err := fakeStore.Create(storeadapter.StoreNode{
				Key:   "syslog_drain_binder/leader",
				Value: []byte("some-other-instance"),
			})
			Expect(err).NotTo(HaveOccurred())

			go candidate.RunForElection()

			Eventually(func() int { return len(testingSink.Records()) }, 3).Should(BeNumerically(">=", 2))
			for _, record := range testingSink.Records() {
				Expect(record.Message).To(Equal("Lost election"))
			}
		})

		It("returns an error if any other error occurs while setting key", func() {
			testError := errors.New("test error")
			fakeStore.CreateErrInjector = fakestoreadapter.NewFakeStoreAdapterErrorInjector("syslog_drain_binder", testError)

			err := candidate.RunForElection()
			Expect(err).To(Equal(testError))
		})

		It("returns nil when it eventually wins", func() {
			preExistingNode := storeadapter.StoreNode{
				Key:   "syslog_drain_binder/leader",
				Value: []byte("some-other-instance"),
			}
			err := fakeStore.Create(preExistingNode)
			Expect(err).NotTo(HaveOccurred())

			errChan := make(chan error)
			go func() {
				errChan <- candidate.RunForElection()
			}()

			time.Sleep(1 * time.Second)
			fakeStore.Delete(preExistingNode.Key)
			Eventually(errChan, 2).Should(Receive(BeNil()))
		})
	})

	Describe("StayAsLeader", func() {
		var candidate *elector.Elector

		BeforeEach(func() {
			fakeStore.Create(storeadapter.StoreNode{
				Key:   "syslog_drain_binder/leader",
				Value: []byte("candidate1"),
			})

		})

		Context("when already leader", func() {
			It("maintains leadership of cluster if successful", func() {
				candidate = elector.NewElector("candidate1", fakeStore, time.Second, logger)

				err := candidate.StayAsLeader()
				Expect(err).NotTo(HaveOccurred())

				node, _ := fakeStore.Get("syslog_drain_binder/leader")
				Expect(node.Value).To(BeEquivalentTo("candidate1"))
				Expect(node.TTL).To(Equal(uint64(1)))
			})
		})

		Context("when not the cluster leader", func() {
			It("returns an error", func() {
				candidate = elector.NewElector("candidate2", fakeStore, time.Second, logger)

				err := candidate.StayAsLeader()
				Expect(err).To(HaveOccurred())
			})

			It("does not replace the existing leader", func() {
				candidate = elector.NewElector("candidate2", fakeStore, time.Second, logger)

				candidate.StayAsLeader()

				node, _ := fakeStore.Get("syslog_drain_binder/leader")
				Expect(node.Value).To(BeEquivalentTo("candidate1"))
			})
		})
	})
})