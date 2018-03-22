package incata

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	. "github.com/semihalev/incata/mocks"
	. "github.com/semihalev/incata/model"
)

var _ = Describe("Retriever", func() {

	It("create a new appender without setup", func() {

		SetupRetriever(nil)
		retriever, err := NewRetriever()
		Expect(retriever).To(BeNil())
		Expect(err).To(MatchError(errors.New("retriever is not set up!")))
	})

	It("retrieve data succeeds", func() {

		var sourceID = uuid.NewV4()
		var data = make([]Event, 0)

		data = append(data, *NewEvent(uuid.NewV4(), time.Now(), GetTestData(), "TEST", 1))
		data = append(data, *NewEvent(sourceID, time.Now(), GetTestData(), "TEST", 1))
		data = append(data, *NewEvent(uuid.NewV4(), time.Now(), GetTestData(), "TEST", 1))
		data = append(data, *NewEvent(sourceID, time.Now(), GetTestData(), "TEST", 1))
		data = append(data, *NewEvent(uuid.NewV4(), time.Now(), GetTestData(), "TEST", 1))

		rd := NewMemoryReader(data)

		SetupRetriever(rd)

		r, err := NewRetriever()
		Expect(err).NotTo(HaveOccurred())

		events, err := r.Retrieve(sourceID)
		Expect(err).NotTo(HaveOccurred())

		Expect(len(events)).To(Equal(2))
	})
})
