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
	last_updated      *prometheus.GaugeVec
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
		last_updated: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "last_updated",
				Help:      "Last updated, in Epoch Unitx Timestamp as reported by simplefin",
			},
			[]string{"domain", "account_name"},
		),
	}
	exporter.Registry.MustRegister(exporter.balances)
	exporter.Registry.MustRegister(exporter.availableBalances)
	exporter.Registry.MustRegister(exporter.last_updated)

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

		e.last_updated.WithLabelValues(accItem.Org.Domain, accItem.Name).Set(float64(accItem.BalanceDate))

	}
	return nil
}
