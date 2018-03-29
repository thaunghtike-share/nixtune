all: dev

release:
	cd mental_models && serverless deploy --stage prod

dev:
	cd mental_models && serverless deploy

deps:
	npm install -g serverless@beta
	pip install astroid isort pylint

lint:
	find . -name "*py" | grep -v "psycopg2" | xargs pylint
