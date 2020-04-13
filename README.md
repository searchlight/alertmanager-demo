# Alertmanager-Demo

### Step 1: Get your slack webhook URL

- (https://api.slack.com/incoming-webhooks)

### Step 2: Download prometheus alert manager and move the binary to project directory
  
- (https://prometheus.io/download/#alertmanager)
    
### Step 3: Configure alert manager for sending alert to specified slack URL

alertmanager.yaml
```yaml
global:
  slack_api_url: 'https://hooks.slack.com/services/../../...'

route:
  receiver: 'slack-notifications'
  group_by: [alertname, datacenter, app]

receivers:
  - name: 'slack-notifications'
    slack_configs:
      - channel: '#alerts'
        text: 'https://mywebsite.com/alerts/{{ .GroupLabels.app }}/{{ .GroupLabels.alertname }}'
```
### Step 4: Run alert manager 

```bash
./alertManager/alertmanager --config.file=alertManager/alertmanager.yaml

```
### Step 5: Run main.go to send alert to alert manager by POST request 

```bash
go run alertSender/main.go

```
### Step 6: Run alertLister to list all alerts

``` bash
go run alertLister/listAlert.go
```

## Architecture

https://github.com/prometheus/alertmanager#architecture

## API Reference
- https://github.com/prometheus/alertmanager#api
- https://github.com/prometheus/alertmanager/tree/master/api/v2
- http://petstore.swagger.io/?url=https://raw.githubusercontent.com/prometheus/alertmanager/master/api/v2/openapi.yaml
