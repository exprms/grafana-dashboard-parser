# Grafana Dashboard Parser

## Ziel des Projekts

Der **Grafana Dashboard Parser** analysiert Grafana-Dashboard-Dateien und extrahiert die verwendeten Prometheus-Metriken aus den enthaltenen PromQL-Abfragen.

Das Projekt richtet sich insbesondere an Administratoren und Entwickler, die einen Überblick darüber erhalten möchten,

* welche Metriken in Grafana verwendet werden,
* in welchen Dashboards und Panels diese vorkommen,
* welche Auswirkungen Änderungen an Metriken oder Recording Rules haben können.

Die Auswertung erfolgt ausschließlich anhand der Dashboard-JSON-Dateien. Eine Verbindung zu Grafana oder Prometheus ist nicht erforderlich.

---

## Funktionen

Der aktuelle Entwicklungsstand unterstützt:

* rekursives Einlesen eines Dashboard-Verzeichnisses
* Verarbeitung beliebig vieler Dashboard-Dateien
* Auswertung verschachtelter Panels
* Analyse von PromQL-Ausdrücken mit dem offiziellen Prometheus PromQL Parser
* Extraktion der tatsächlich verwendeten Metriken (keine Funktionen oder Schlüsselwörter)

### Verfügbare Reports

#### Dashboard-Übersicht

Zeigt für jedes Dashboard

* alle Panels
* die verwendeten PromQL-Abfragen
* die darin enthaltenen Metriken

Beispiel:

```text
Dashboard: Kubernetes

Panel: CPU

Query:
sum(rate(node_cpu_seconds_total{mode!="idle"}[5m]))

Metrics:
- node_cpu_seconds_total
```

---

#### Metrik-Übersicht

Listet jede verwendete Metrik genau einmal auf und zeigt anschließend

* das Dashboard
* das Panel
* die vollständige PromQL-Abfrage

in der die Metrik verwendet wird.

Dadurch lässt sich schnell nachvollziehen, welche Dashboards von Änderungen an einer Metrik betroffen sind.

---

## Projektstruktur

```text
grafana-dashboard-parser/
│
├── cmd/
│   └── analyzer/
│
├── internal/
│   ├── analyzer/
│   ├── dashboard/
│   ├── model/
│   └── report/
│
└── dashboards/
```

---

## Verwendung

### Dashboards einlesen

Standardmäßig werden lediglich die gefundenen Dashboards ausgegeben.

```bash
go run ./cmd/analyzer
```

Alternativ kann ein anderes Verzeichnis angegeben werden:

```bash
go run ./cmd/analyzer -dir /pfad/zu/meinen/dashboards
```

---

### Dashboard-Report

```bash
go run ./cmd/analyzer -report dashboards
```

---

### Metrik-Report

```bash
go run ./cmd/analyzer -report metrics
```

---

## Voraussetzungen

* Go 1.24 oder neuer
* Grafana-Dashboards im JSON-Format

Die PromQL-Auswertung erfolgt mit dem offiziellen Prometheus-Paket.

---

## Roadmap

Geplante Erweiterungen:

* CSV-Export
* Markdown-Export
* HTML-Report
* Statistik über die Verwendung einzelner Metriken
* Suche nach Dashboards anhand einer Metrik
* Suche nach ungenutzten bzw. selten verwendeten Metriken
* Erkennung identischer oder ähnlicher PromQL-Abfragen
* Analyse der verwendeten Labels
* Unterstützung weiterer Datasources

---

## Projektstatus

Das Projekt befindet sich derzeit in aktiver Entwicklung. Die Grundarchitektur für das Einlesen, Analysieren und Auswerten von Grafana-Dashboards ist implementiert und wird schrittweise um weitere Analyse- und Reporting-Funktionen erweitert.

