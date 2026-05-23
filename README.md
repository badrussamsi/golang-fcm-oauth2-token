# golang-fcm-oauth2-token

A lightweight Golang utility for generating a Firebase Cloud Messaging (FCM) access token for testing push notifications during backend development.

## Purpose

In many projects, backend push notification APIs are developed later in the sprint.  
This tool helps mobile developers continue testing notification flows independently.

## Prerequisites

- Go `1.20+`
- A Firebase service account JSON key file
- Firebase project with Cloud Messaging enabled

## Quick Start

1. Install dependencies:

```bash
go mod tidy
```

2. Create a `.env` file in the project root:

```env
FIREBASE_CREDENTIAL_PATH=./private_key/<your-service-account>.json
```

3. Put your Firebase service account key file in `private_key/`.

4. Run the app:

```bash
go run main.go
```

## Security Notes

- Never commit `.env` or files under `private_key/` to version control.
- Treat generated access tokens as sensitive credentials.
- Do not share access tokens in screenshots, logs, chat, or issue trackers.
