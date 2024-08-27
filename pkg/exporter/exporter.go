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
	Registry          *prometheus.Registry
	balances          *prometheus.GaugeVec
	availableBalances *prometheus.GaugeVec
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
			[]string{"domain", "account_name", "currency"},
		),
		availableBalances: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "available_balance",
				Help:      "Available account balance",
			},
			[]string{"domain", "account_name", "currency"},
		),
	}
	exporter.Registry.MustRegister(exporter.balances)
	exporter.Registry.MustRegister(exporter.availableBalances)

	return exporter
}

func (e *Exporter) Export(accounts *simplefin.Accounts) error {
	for _, accItem := range accounts.Accounts {

		bal, err := strconv.ParseFloat(accItem.Balance, 32)
		if err != nil {
			log.Error().Err(err).Msgf("Could not parse balance from %v - %v)",
				accItem.Org.Domain, accItem.Name)

		} else {
			e.balances.WithLabelValues(accItem.Org.Domain, accItem.Name, accItem.Currency).Set(bal)
		}

		availBal, err := strconv.ParseFloat(accItem.AvailableBalance, 32)
		if err != nil {
			log.Error().Err(err).Msgf("Could not parse available balance from %v - %v)",
				accItem.Org.Domain, accItem.Name)

		} else {
			e.availableBalances.WithLabelValues(accItem.Org.Domain, accItem.Name, accItem.Currency).Set(availBal)
		}

	}
	return nil
}
