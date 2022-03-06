package kafka

import (
	"context"
	"os"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/pkorobeinikov/seeder/seeder"
)

const (
	extJson = "json"
	extYaml = "yaml"
)

func Test_Seed_SeedFileParse(t *testing.T) {
	type fields struct {
		sp sarama.SyncProducer
	}
	type args struct {
		ctx context.Context
		cfg seeder.Config
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		{
			"valid json file with topic and value",
			fields{sp: new(MockSyncProducer)},
			args{
				ctx: context.Background(),
				cfg: seeder.Config{
					File: tempFile(
						extJson,
						`[
							{
								"topic": "my_topic.1",
								"value": {
									"id": 1,
									"name": "alice"
								}
							}
						]`),
				},
			},
			1,
			false,
		},
		{
			"valid json file with topic and unset value",
			fields{sp: new(MockSyncProducer)},
			args{
				ctx: context.Background(),
				cfg: seeder.Config{
					File: tempFile(
						extJson,
						`[
							{
								"topic": "my_topic.1"
							}
						]`),
				},
			},
			1,
			false,
		},
		{
			"valid json file with topic, key and value",
			fields{sp: new(MockSyncProducer)},
			args{
				ctx: context.Background(),
				cfg: seeder.Config{
					File: tempFile(
						extJson,
						`[
							{
								"topic": "my_topic.1",
								"key": "boo",
								"value": {
									"id": 1,
									"name": "alice"
								}
							}
						]`),
				},
			},
			1,
			false,
		},
		{
			"invalid json file",
			fields{sp: new(MockSyncProducer)},
			args{
				ctx: context.Background(),
				cfg: seeder.Config{
					File: tempFile(
						extJson,
						`[}`),
				},
			},
			-1,
			true,
		},
		{
			"valid json file without topic",
			fields{sp: new(MockSyncProducer)},
			args{
				ctx: context.Background(),
				cfg: seeder.Config{
					File: tempFile(
						extJson,
						`[
							{
								"no-topic": "my_topic.1",
								"value": {
									"id": 1,
									"name": "alice"
								}
							}
						]`),
				},
			},
			-1,
			true,
		},
		{
			"valid yaml file with topic and value",
			fields{sp: new(MockSyncProducer)},
			args{
				ctx: context.Background(),
				cfg: seeder.Config{
					File: tempFile(
						extYaml,
						`---
data:
  - topic: foo_topic.1
    value:
      id: 1
      name: "alice"
...`),
				},
			},
			1,
			false,
		},
		{
			"valid yaml file without topic",
			fields{sp: new(MockSyncProducer)},
			args{
				ctx: context.Background(),
				cfg: seeder.Config{
					File: tempFile(
						extYaml,
						`---
data:
  - no-topic: foo_topic.1
    value:
      id: 1
      name: "alice"
...`),
				},
			},
			-1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &kafkaSeeder{
				sp: tt.fields.sp,
			}
			defer os.Remove(tt.args.cfg.File)

			n, err := s.Seed(tt.args.ctx, tt.args.cfg)
			if err != nil {
				t.Log("err:", err)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Seed() error = %v, wantErr %v", err, tt.wantErr)
			}

			if n != tt.wantN {
				t.Errorf("Seed() n = %d, wantN %d", n, tt.wantN)
			}
		})
	}
}

func tempFile(ext, content string) (filename string) {
	f, err := os.CreateTemp("", "*-test."+strings.TrimPrefix(ext, "."))
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}

	return f.Name()
}

type MockSyncProducer struct {
	offset int64
}

func (m *MockSyncProducer) SendMessage(_ *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return 0, atomic.AddInt64(&m.offset, 1), nil
}

func (m *MockSyncProducer) SendMessages(_ []*sarama.ProducerMessage) error {
	return nil
}

func (m *MockSyncProducer) Close() error {
	return nil
}
