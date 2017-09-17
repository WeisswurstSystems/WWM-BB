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

# Umgebungsvariablen

| Variablenname | Verwendung | Beispiel |
| --- | --- | --- |
| `mail.port` | Port des zu verwendenden Mail-Servers | 587 |
| `mail.smtpServer` | Pfad des zu verwendenen Mail-Servers | smtp.gmail.com |
| `mail.username` | Nutzername des zu verwendenen E-Mail Kontos | fabian.wilms@gmail.com |
| `mail.password` | Passwort des zu verwendenen Mail-Kontos | **** |
||||
| `db.url` | Pfad zu Mongo-DB Server | something.mlab.com:12345/some-db |
| `db.name` | Name der zu verwendenden Datenbank auf dem Server | some-db |
| `db.username` | Nutzername des DB-Users | hans |
| `db.password` | Password des Users | **** |
