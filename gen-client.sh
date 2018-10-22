#!/usr/bin/env bash

base_url=https://management.auth0.com/api/v2/api-docs

api-spec-converter ${base_url}/Client%20Grants  --from=swagger_1 --to=swagger_2 > swagger/client-grants.json
api-spec-converter ${base_url}/Clients --from=swagger_1 --to=swagger_2 > swagger/clients.json
api-spec-converter ${base_url}/Connections --from=swagger_1 --to=swagger_2 > swagger/connections.json
api-spec-converter ${base_url}/Custom%20Domains --from=swagger_1 --to=swagger_2 > swagger/custom-domains.json
api-spec-converter ${base_url}/Device%20Credentials --from=swagger_1 --to=swagger_2 > swagger/device-credentials.json
api-spec-converter ${base_url}/Grants --from=swagger_1 --to=swagger_2 > swagger/grants.json
api-spec-converter ${base_url}/Logs --from=swagger_1 --to=swagger_2 > swagger/logs.json
api-spec-converter ${base_url}/Resource%20Servers --from=swagger_1 --to=swagger_2 > swagger/resource-servers.json
api-spec-converter ${base_url}/Rules --from=swagger_1 --to=swagger_2 > swagger/rules.json
api-spec-converter ${base_url}/Rules%20Configs --from=swagger_1 --to=swagger_2 > swagger/rules-configs.json
api-spec-converter ${base_url}/User%20Blocks --from=swagger_1 --to=swagger_2 > swagger/user-blocks.json
api-spec-converter ${base_url}/Users --from=swagger_1 --to=swagger_2 > swagger/users.json
api-spec-converter ${base_url}/Users%20By%20Email --from=swagger_1 --to=swagger_2 > swagger/users-by-email.json
api-spec-converter ${base_url}/Blacklists --from=swagger_1 --to=swagger_2 > swagger/blacklists.json
api-spec-converter ${base_url}/Emails --from=swagger_1 --to=swagger_2 > swagger/emails.json
api-spec-converter ${base_url}/Email%20Templates --from=swagger_1 --to=swagger_2 > swagger/email-templates.json
api-spec-converter ${base_url}/Guardian --from=swagger_1 --to=swagger_2 > swagger/guardian.json
api-spec-converter ${base_url}/Jobs --from=swagger_1 --to=swagger_2 > swagger/jobs.json
api-spec-converter ${base_url}/Stats --from=swagger_1 --to=swagger_2 > swagger/stats.json
api-spec-converter ${base_url}/Tenants --from=swagger_1 --to=swagger_2 > swagger/tenants.json
api-spec-converter ${base_url}/Tickets --from=swagger_1 --to=swagger_2 > swagger/tickets.json

python merge-swagger.py

#swagger generate client -f swagger.json -A auth0_management --skip-validation --target .


