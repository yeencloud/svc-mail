package metrics

type MailSentMetrics struct {
	Address string `metric:"address"`
	Subject string `metric:"subject"`

	Status string `metric:"status"`
}
