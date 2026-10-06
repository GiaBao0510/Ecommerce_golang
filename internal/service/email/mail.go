package email

import (
	"time"

	"github.com/GiaBao0510/Ecommerce_golang/pkg/loghelper"
)

type MailConfig struct {
	ProviderConfigs map[string]any
	ProviderType    ProviderType
	MaxRetries      int
	Timeout         time.Duration
	Logger          *loghelper.ServiceLogger
}

func NewMailConfig(logger *loghelper.ServiceLogger, providerFactory ProviderFactory) (IEmailProviderService, error) {
	
}