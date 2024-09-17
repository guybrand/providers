# providers
Plan your code to scale - Provider Pattern demo code from Sep 2024 go meetup @Cyolo

This repo demonstrates a walkthrough from a simple implemntation of a code iterfacing with an SMS provider to a pluggable system that can easily scale,

- [v1](https://github.com/guybrand/providers/tree/main/v1) and [v2](https://github.com/guybrand/providers/tree/main/v2) are quoted within the presentation (pages 6 and 11-17 respectively)
- [v3](https://github.com/guybrand/providers/tree/main/v3) - implements Separation of concerns (page 20)
- [v4](https://github.com/guybrand/providers/tree/main/v4) - implements constructor dependency injection (page 21)
- [v5](https://github.com/guybrand/providers/tree/main/v5) - implements usage of plugins
- [v6](https://github.com/guybrand/providers/tree/main/v6) - implements usage of go.uber.org/fx dependency injection (page 35)


- [tax](https://github.com/guybrand/providers/tree/main/tax) - is just a mock used for pages 28-29.

## builing the plugins from v5 root folder:
```
go build -buildmode=plugin -o sms/plugin/twilio.so sms/plugin/twilio/twilio.go
go build -buildmode=plugin -o sms/plugin/alibaba.so sms/plugin/alibaba/alibaba.go

```


## presentation link:
[Plan your code to scale - Provider Pattern](https://docs.google.com/presentation/d/1mmD8sca-VeVZVezxtDHw2WtNHKVGUyMfptYparVR3Z8)
## meetup link:
[Sep 2024 go meetup @Cyolo](https://www.meetup.com/go-israel/events/302419516/?slug=go-israel&eventId=302419516)
