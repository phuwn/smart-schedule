for seed in *.sql; do PGPASSWORD=password psql -U phuwn -d smart_schedule -f /${seed}; done