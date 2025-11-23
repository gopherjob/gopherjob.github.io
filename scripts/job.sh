#!/bin/sh

#!/usr/bin/env bash

# List of capital cities
countries=(
  "Berlin"
  "Germany"
  "Netherlands"
  "Finland"
  "Sweden"
  "Denmark"
  "Norway"
  "Estonia"
  "Ireland"
  "United Kingdom"
  "Portugal"
  "Spain"
)

# Proxy list
proxies=(
  localhost
)

# Build proxies args
proxy_args=()
for proxy in "${proxies[@]}"; do
  proxy_args+=(--proxies "$proxy")
done

# Create unified output directory
mkdir -p job_results

# Loop through countries
for country in "${countries[@]}"; do
  echo ">>> Running jobsparser for $country..."

  jobsparser --search-term '"Golang Developer" OR "Go Developer" OR "Backend Golang Developer" OR "Backend Engineer Go" OR "Software Engineer Go"' \
    --location "$country" \
    --site linkedin \
    --results-wanted 300 \
    --distance 50 \
    --job-type fulltime \
    --fetch-description \
    --hours-old 48 \
    --max-retries 5 \
    --output-dir "job_results" \
    "${proxy_args[@]}"
done