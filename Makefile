ENCRYPTION_KEY="8a3f5c9d1b7e4f2a6c90d8b3e7f2410e"

run:
	@echo "running app"
	@ENCRYPTION_KEY=$(ENCRYPTION_KEY) go run main.go

uploadTest:
	curl -X POST http://localhost:8080/upload \
		-F 'creds={"username":"praneeth","emailid":"prax@gmai.com","password":"securePassword123"}' \
		-F 'file=@/mnt/c/Users/prap/Desktop/PersonalLearngings/FILESTORAGE/TEST-FILE-SERVER/testfile.txt'

downloadTest:
	curl -X POST "http://localhost:8080/download?filename=testfile.txt" \
		-F 'creds={"username":"praneeth","emailid":"prax@gmai.com","password":"securePassword123"}' \
		--output ../testfile.txt

push:
	@git add .
	@if [ -z "$(commit_message)" ]; then \
		echo "Error: commit_message variable is not set"; \
		exit 1; \
	else \
		git commit -m "$(commit_message)"; \
		git push origin main; \
		echo "changes pushed to github"; \
	fi
