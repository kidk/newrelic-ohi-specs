# New Relic Infrastructure Integration for specs

Reports status and metrics for Operating System and bare metal specs.

## Requirements

* [New Relic infrastructure installed](https://docs.newrelic.com/docs/infrastructure/install-configure-manage-infrastructure)

## Installation

1. Download the latest version from [here](https://github.com/kidk/newrelic-ohi-specs/releases)
2. Unzip in a temporary location
3. Copy `newrelic-specs-definition.yml` and the entire `bin/` folder to `/var/db/newrelic-infra/custom-integrations`
4. Copy `newrelic-specs-config.yml` to `/etc/newrelic-infra/integrations.d`
5. [Restart the New Relic agent](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/configuration/start-stop-restart-check-infrastructure-agent-status)

## Compatibility

* Supported (and tested) Operating Systems: Ubuntu, Debian

