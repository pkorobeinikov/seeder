package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/pkorobeinikov/seeder/seeder"
	"gopkg.in/yaml.v3"
)

const (
	SeederKafkaPeerEnv = "SEEDER_KAFKA_PEER"

	seederType = "kafka"
)

func Seed(ctx context.Context, cfg seeder.Config) error {

	peer, found := os.LookupEnv(SeederKafkaPeerEnv)
	if !found {
		return errors.New("connection string is not set")
	}

	b, err := os.ReadFile(cfg.File)
	if err != nil {
		return errors.Wrap(err, "read file")
	}

	p, err := newSyncProducer([]string{peer})
	if err != nil {
		return errors.Wrap(err, "new sync producer")
	}

	var payload []seed

	ext := filepath.Ext(cfg.File)
	switch ext {
	case ".json":
		fmt.Println("seeding json file:", cfg.File)

		var p []*jsonSeed
		err := json.Unmarshal(b, &p)
		if err != nil {
			return errors.Wrap(err, "unmarshal json")
		}

		for i := range p {
			payload = append(payload, p[i])
		}

	case ".yml", ".yaml":
		fmt.Println("seeding yaml file:", cfg.File)

		var p map[string]interface{}

		err := yaml.Unmarshal(b, &p)
		if err != nil {
			return errors.Wrap(err, "unmarshal yaml")
		}

		data, ok := p["data"].([]interface{})
		if !ok {
			return errors.New("bad data format, expected array")
		}

		for i := range data {
			v := data[i].(map[string]interface{})
			j := jsonSeed{
				Topic: getMapValueAsString(v, "topic"),
				Key:   getMapValueAsString(v, "key"),
				Value: getMapValueAsBytes(v, "value"),
			}
			payload = append(payload, &j)
		}

	default:
		return errors.Errorf("unsupported file type: %s", ext)
	}

	for _, v := range payload {
		_, offset, err := p.SendMessage(&sarama.ProducerMessage{
			Topic: v.GetTopic(),
			Key:   sarama.StringEncoder(v.GetKey()),
			Value: sarama.ByteEncoder(v.GetValue()),
		})
		if err != nil {
			return errors.Wrapf(err, "kafka: send message: offset=%d", offset)
		}
	}

	fmt.Println("seeded items:", len(payload))

	return nil
}

func newSyncProducer(brokerList []string) (sarama.SyncProducer, error) {

	c := sarama.NewConfig()
	c.Producer.RequiredAcks = sarama.WaitForLocal
	c.Producer.Retry.Max = 10
	c.Producer.Return.Successes = true

	p, err := sarama.NewSyncProducer(brokerList, c)
	if err != nil {
		return nil, err
	}

	return p, nil
}

type (
	seed interface {
		GetTopic() string
		GetKey() string
		GetValue() []byte
	}
	jsonSeed struct {
		Topic string          `json:"topic" yaml:"topic"`
		Key   string          `json:"key" yaml:"key"`
		Value json.RawMessage `json:"value" yaml:"value"`
	}
	msi map[string]interface{}
)

func (s *jsonSeed) GetTopic() string {
	return s.Topic
}

func (s *jsonSeed) GetValue() []byte {
	return s.Value
}

func (s *jsonSeed) GetKey() string {
	return s.Key
}

func getMapValueAsString(m msi, key string) string {
	v, ok := m[key]
	if !ok {
		return ""
	}

	s, ok := v.(string)
	if !ok {
		return ""
	}

	return s
}

// fixme: naming or move to builder
func getMapValueAsBytes(m msi, key string) []byte {
	mm, ok := m[key]
	if !ok {
		return nil
	}

	b, err := json.Marshal(mm)
	if err != nil {
		return nil
	}

	return b
}

func init() {
	seeder.DefaultRegistry().RegisterSeeder(Seed, seederType)

	seeder.DefaultRegistry().RegisterSeederHelp(
		func(w io.Writer) {
			_, _ = fmt.Fprintf(
				w,
				`Kafka seeder env variables:

- %s: peer address, example: 127.0.0.1:9092


Run example (in folder "seeder-showcase"):

$ SEEDER_KAFKA_PEER=127.0.0.1:9092 seeder -c ./401_kafka/seeder.yaml


Json seed file example:

[
    {
        "topic": "my_topic.1",    // required
        "key": "boo",             // optional
        "value": {                // required
            "id": 1,
            "name": "alice"
        }
    }
]


Yaml seed file example:

---
data:
  - topic: foo_topic.1
    value:
      id: 1
      name: "alice"
  - topic: foo_topic.2
    key: "foo"
    value:
      id: 2
      name: "bob"
...

`,
				SeederKafkaPeerEnv,
			)
		},
		seederType,
	)

}
