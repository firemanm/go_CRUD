{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": {
          "type": "grafana",
          "uid": "-- Grafana --"
        },
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "fiscalYearStartMonth": 0,
  "graphTooltip": 0,
  "id": 5,
  "links": [],
  "panels": [
    {
      "datasource": {
        "type": "prometheus",
        "uid": "ceznufir6eqkge"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": 0
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 7,
        "w": 6,
        "x": 0,
        "y": 0
      },
      "id": 1,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "hideZeros": false,
          "mode": "single",
          "sort": "none"
        }
      },
      "pluginVersion": "12.1.1",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "ceznufir6eqkge"
          },
          "disableTextWrap": false,
          "editorMode": "builder",
          "exemplar": false,
          "expr": "histogram_quantile(0.5, sum by(le, method) (rate(gocrudapp_http_request_duration_seconds_bucket{method=\"$api_method\"}[5m])))",
          "fullMetaSearch": false,
          "includeNullMetadata": true,
          "legendFormat": "{{method}} - q0.5",
          "range": true,
          "refId": "A",
          "useBackend": false
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "ceznufir6eqkge"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.95, sum by(le, method) (rate(gocrudapp_http_request_duration_seconds_bucket{method=\"$api_method\"}[5m])))",
          "hide": false,
          "instant": false,
          "legendFormat": "{{method}} - q0.95",
          "range": true,
          "refId": "B"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "ceznufir6eqkge"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.99, sum by(le, method) (rate(gocrudapp_http_request_duration_seconds_bucket{method=\"$api_method\"}[5m])))",
          "hide": false,
          "instant": false,
          "legendFormat": "{{method}} - q0.99",
          "range": true,
          "refId": "C"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "ceznufir6eqkge"
          },
          "disableTextWrap": false,
          "editorMode": "builder",
          "expr": "max by(method) (rate(gocrudapp_http_request_duration_seconds_sum{method=\"$api_method\"}[5m])) / (max by(method) (rate(gocrudapp_http_request_duration_seconds_count[5m])))",
          "fullMetaSearch": false,
          "hide": false,
          "includeNullMetadata": true,
          "instant": false,
          "legendFormat": "{{method}} - qMAX",
          "range": true,
          "refId": "D",
          "useBackend": false
        }
      ],
      "title": "Latency",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "ceznufir6eqkge"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": 0
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 7,
        "w": 6,
        "x": 6,
        "y": 0
      },
      "id": 2,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "hideZeros": false,
          "mode": "single",
          "sort": "none"
        }
      },
      "pluginVersion": "12.1.1",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "ceznufir6eqkge"
          },
          "disableTextWrap": false,
          "editorMode": "builder",
          "exemplar": false,
          "expr": "sum by(le, method) (rate(gocrudapp_http_requests_total{method=\"$api_method\"}[5m]))",
          "fullMetaSearch": false,
          "includeNullMetadata": true,
          "legendFormat": "{{method}} - q0.5",
          "range": true,
          "refId": "A",
          "useBackend": false
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "ceznufir6eqkge"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.95, sum by(le, method) (rate(gocrudapp_http_request_duration_seconds_bucket{method=\"$api_method\"}[5m])))",
          "hide": false,
          "instant": false,
          "legendFormat": "{{method}} - q0.95",
          "range": true,
          "refId": "B"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "ceznufir6eqkge"
          },
          "editorMode": "code",
          "expr": "histogram_quantile(0.99, sum by(le, method) (rate(gocrudapp_http_request_duration_seconds_bucket{method=\"$api_method\"}[5m])))",
          "hide": false,
          "instant": false,
          "legendFormat": "{{method}} - q0.99",
          "range": true,
          "refId": "C"
        },
        {
          "datasource": {
            "type": "prometheus",
            "uid": "ceznufir6eqkge"
          },
          "disableTextWrap": false,
          "editorMode": "builder",
          "expr": "max by(method) (rate(gocrudapp_http_request_duration_seconds_sum{method=\"$api_method\"}[5m])) / (max by(method) (rate(gocrudapp_http_request_duration_seconds_count[5m])))",
          "fullMetaSearch": false,
          "hide": false,
          "includeNullMetadata": true,
          "instant": false,
          "legendFormat": "{{method}} - qMAX",
          "range": true,
          "refId": "D",
          "useBackend": false
        }
      ],
      "title": "RPS",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "ceznufir6eqkge"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisBorderShow": false,
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "barWidthFactor": 0.6,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "insertNulls": false,
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": 0
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 7,
        "w": 6,
        "x": 12,
        "y": 0
      },
      "id": 3,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "hideZeros": false,
          "mode": "single",
          "sort": "none"
        }
      },
      "pluginVersion": "12.1.1",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "ceznufir6eqkge"
          },
          "disableTextWrap": false,
          "editorMode": "code",
          "exemplar": false,
          "expr": "rate(gocrudapp_http_requests_total{statuscode=\"500\"}[5m])",
          "format": "table",
          "fullMetaSearch": false,
          "includeNullMetadata": false,
          "instant": false,
          "legendFormat": "__auto",
          "range": true,
          "refId": "A",
          "useBackend": false
        }
      ],
      "title": "ERROR Rate",
      "type": "timeseries"
    }
  ],
  "preload": false,
  "refresh": "",
  "schemaVersion": 41,
  "tags": [],
  "templating": {
    "list": [
      {
        "current": {
          "text": "POST",
          "value": "POST"
        },
        "definition": "label_values(gocrudapp_http_request_duration_seconds_bucket,method)",
        "description": "label_values(http_request_duration_seconds_bucket, api_method)",
        "name": "api_method",
        "options": [],
        "query": {
          "qryType": 1,
          "query": "label_values(gocrudapp_http_request_duration_seconds_bucket,method)",
          "refId": "PrometheusVariableQueryEditor-VariableQuery"
        },
        "refresh": 1,
        "regex": "",
        "type": "query"
      }
    ]
  },
  "time": {
    "from": "now-30m",
    "to": "now"
  },
  "timepicker": {},
  "timezone": "browser",
  "title": "Filtering by method",
  "uid": "1640c075-3e25-49d6-9d8f-8bffe5695735",
  "version": 22
}