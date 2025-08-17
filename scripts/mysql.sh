set -a
source "$(dirname "$0")/../.env.local"
set +a

MYSQL_PWD="${DB_PASSWORD}" mysql \
  --protocol=TCP \
  -h "${DB_HOST}" \
  -P "${DB_PORT}" \
  -u "${DB_USERNAME}" \
  "${DB_NAME}"