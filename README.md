# WWM-BB

[![Build Status](https://travis-ci.org/WeisswurstSystems/WWM-BB.svg?branch=master)](https://travis-ci.org/WeisswurstSystems/WWM-BB)

[![Coverage Status](https://coveralls.io/repos/github/WeisswurstSystems/WWM-BB/badge.svg?branch=master)](https://coveralls.io/github/WeisswurstSystems/WWM-BB?branch=master)

# Endpunkte

```
users/                [get -> alle Nutzer lesen, post -> registrieren]
meetings/             [get -> meetings in Kurzform lesen, post -> Neues Meeting erstellen]
  |- {id}/            [get -> Einzelnes meeting mit Details lesen, patch -> Meeting bearbeiten]
      |- pay/         [post -> Bestellung als bezahlt markieren]
      |- invite/      [post mit user-id Liste -> Einladung versenden]
      |- notify/      [post mit message -> Update-Informationen an alle teilnehmenden Nutzer]
```
