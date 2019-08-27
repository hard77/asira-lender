package middlewares

import (
	"asira_lender/asira"
	"asira_lender/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type (
	AsiraKafkaHandlers struct {
		KafkaProducer     sarama.AsyncProducer
		KafkaConsumer     sarama.Consumer
		PartitionConsumer sarama.PartitionConsumer
	}
	BorrowerInfo struct {
		Info interface{} `json:"borrower_info"`
	}
)

func init() {
	topics := asira.App.Config.GetStringMap(fmt.Sprintf("%s.kafka.topics.consumes", asira.App.ENV))

	kafka := &AsiraKafkaHandlers{}
	kafka.KafkaProducer = asira.App.Kafka.Producer
	kafka.KafkaConsumer = asira.App.Kafka.Consumer

	kafka.SetPartitionConsumer(topics["new_loan"].(string))

	go func() {
		for {
			message, err := kafka.Listen()
			if err != nil {
				log.Printf("error occured when listening kafka : %v", err)
			}
			if message != nil {
				var brwr models.Borrower
				brwr, err = syncBorrowerData(message)
				if err != nil {
					log.Println(err)
				}

				err := syncLoanData(message, brwr)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()
}

func (k *AsiraKafkaHandlers) SetPartitionConsumer(topic string) (err error) {
	k.PartitionConsumer, err = k.KafkaConsumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}

	return nil
}

func (k *AsiraKafkaHandlers) Listen() ([]byte, error) {
	select {
	case err := <-k.PartitionConsumer.Errors():
		return nil, err
	case msg := <-k.PartitionConsumer.Messages():
		return msg.Value, nil
	}

	return nil, fmt.Errorf("unidentified error while listening")
}

func syncLoanData(kafkaMessage []byte, borrower models.Borrower) (err error) {
	var loan models.Loan
	err = json.Unmarshal(kafkaMessage, &loan)
	if err != nil {
		return err
	}

	loan.Bank = borrower.Bank
	loan.OwnerName = borrower.Fullname

	// loan.ID = uint64(0) // remove ID so it can create new instead of using id from borrower
	// loan.Create()
	_, err = loan.Save()
	return err
}

func syncBorrowerData(kafkaMessage []byte) (borrower models.Borrower, err error) {
	var borrowerInfo BorrowerInfo
	err = json.Unmarshal(kafkaMessage, &borrowerInfo)
	if err != nil {
		return borrower, err
	}

	marshal, err := json.Marshal(borrowerInfo.Info)
	if err != nil {
		return borrower, err
	}

	err = json.Unmarshal(marshal, &borrower)
	if err != nil {
		return borrower, err
	}

	_, err = borrower.Save()
	return borrower, err
}
