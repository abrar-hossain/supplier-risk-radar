# Supplier Risk Radar

AI-powered supplier risk monitoring platform for UK SMEs.

## What It Does

UK small businesses lose thousands every year when a supplier 
fails without warning. Supplier Risk Radar monitors every 
supplier automatically and sends early warning alerts months 
before a supplier fails.

## How It Works

- Connects to accounting software (Xero, QuickBooks, Sage)
- Pulls real supplier data from Companies House API
- Calculates a composite risk score for each supplier
- Sends alerts when a supplier shows signs of distress

## Tech Stack

- **Backend:** Go (Gin)
- **Database:** PostgreSQL (Docker)
- **Cache:** Redis (Docker)
- **External API:** Companies House REST API
- **Architecture:** Clean Architecture, Repository Pattern

## Project Status

🟡 Active Development — MVP in progress

## Author

Abrar Hossain — Co-Founder, Supplier Risk Radar