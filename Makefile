permissions:
	find ./scripts -type f -exec chmod +x {} \;
	bash ./permissions.sh

start : 
	bash ./scripts/start.sh