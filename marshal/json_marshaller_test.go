package marshal_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/semihalev/incata/marshal"
	. "github.com/semihalev/incata/mocks"
)

var _ = Describe("Marshal", func() {

	It("serialize test data to json and match", func() {
		expectedString := `{"version":1,"name":"Joe","balance":12.99,"birth_date":"2015-12-13T23:59:59+02:00"}`

		location, _ := time.LoadLocation("Europe/Athens")

		testData := TestData{
			Version:   1,
			Name:      "Joe",
			Balance:   12.99,
			BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
		}

		serializedString, err := NewJSONMarshaller().Serialize(testData)

		Expect(serializedString).To(Equal(expectedString))
		Expect(err).NotTo(HaveOccurred())
	})

	It("deserialize json to test data and match", func() {
		location, _ := time.LoadLocation("Europe/Athens")

		expected := TestData{
			Version:   1,
			Name:      "Joe",
			Balance:   12.99,
			BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
		}

		actualData := `{"version":1,"name":"Joe","balance":12.99,"birth_date":"2015-12-13T23:59:59+02:00"}`
		var actual TestData

		err := NewJSONMarshaller().Deserialize(actualData, &actual)

		Expect(actual.Balance).To(Equal(expected.Balance))
		Expect(actual.BirthDate.Equal(expected.BirthDate)).To(BeTrue())
		Expect(actual.Name).To(Equal(expected.Name))
		Expect(actual.Version).To(Equal(expected.Version))
		Expect(err).NotTo(HaveOccurred())
	})

	It("deserialize fails due to invalid json", func() {

		var actual TestData
		err := NewJSONMarshaller().Deserialize(`{"version":1,"name":"Joe","balance":12.99,"birth_date":"2015-12-13T23:59:59+02:00------"}`, &actual)
		Expect(err).To(HaveOccurred())
	})

	It("deserialize wrong to the struct", func() {

		var actual TestData
		err := NewJSONMarshaller().Deserialize(123, &actual)
		Expect(err).To(HaveOccurred())
	})

	Measure("benchmarking JSON serializer", func(b Benchmarker) {

		var m = NewJSONMarshaller()
		location, _ := time.LoadLocation("Europe/Athens")

		testData := TestData{
			Version:   1,
			Name:      "Joe",
			Balance:   12.99,
			BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
		}

		runtime := b.Time("runtime", func() {

			m.Serialize(testData)
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 0.2), "Serialize shouldn't take too long.")
	}, 1000)
})
