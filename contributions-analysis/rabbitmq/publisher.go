package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wnn-dev/contributions-analysis/objects"
)

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

type GalleryCardMessage struct {
	ContributorID string    `json:"contributor_id"`
	HtmlContent   string    `json:"html_content"`
	ApprovedAt    time.Time `json:"approved_at"`
	Percentage    float64   `json:"percentage"`
}

func NewPublisher(amqpURL string) (*Publisher, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"approved_submissions_ex", // name
		"fanout",                  // type
		true,                      // durable
		false,                     // auto-deleted
		false,                     // internal
		false,                     // no-wait
		nil,                       // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &Publisher{
		conn:    conn,
		channel: ch,
	}, nil
}

func (p *Publisher) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
}

func (p *Publisher) PublishApprovedSubmission(ctx context.Context, sub *objects.HtmlCssSubmission, report *objects.HtmlCssAnalysisReport) error {
	msg := GalleryCardMessage{
		ContributorID: sub.ContributorID,
		HtmlContent:   sub.HtmlContent,
		ApprovedAt:    time.Now(),
		Percentage:    report.Percentage,
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// We use a fanout exchange so all gallery instances (if multiple) can get the message
	err = p.channel.PublishWithContext(ctx,
		"approved_submissions_ex", // exchange
		"",                        // routing key (ignored by fanout)
		false,                     // mandatory
		false,                     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	return nil
}
