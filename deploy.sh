#!/bin/bash
gcloud functions deploy Hello \
  --project random-stuff-292920 \
  --entry-point Hello \
  --runtime go113\
  --trigger-http
