package exporter

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/eduser25/simplefin-bridge-exporter/pkg/logger"
	"github.com/eduser25/simplefin-bridge-exporter/pkg/simplefin"
)

const (
	namespace = "simplefin"
)

var (
	log = logger.NewZerologLogger()
)

type Exporter struct {
	Registry *prometheus.Registry
	balances *prometheus.GaugeVec
}

func NewExporter() *Exporter {
	exporter := &Exporter{
		Registry: prometheus.NewRegistry(),
		balances: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "balance",
				Help:      "Account balance",
			},
			[]string{"domain", "account_name"},
		),
	}
	exporter.Registry.MustRegister(exporter.balances)

	return exporter
}

func (e *Exporter) Export(accounts *simplefin.Accounts) error {
	for _, accItem := range accounts.Accounts {
		bal, err := strconv.ParseFloat(accItem.Balance, 32)
		if err != nil {
			log.Error().Err(err).Msgf("Could not parse flot from %v - %v)",
				accItem.Org.Domain, accItem.Name)
		}

		e.balances.WithLabelValues(accItem.Org.Domain, accItem.Name).Set(bal)
	}
	return nil
}
