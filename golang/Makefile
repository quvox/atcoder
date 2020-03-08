SHELL=/usr/local/bin/bash

env:
	@if [ -z "${user}" -o -z "${password}" ]; then \
		echo; \
		echo "XXXX user and password parameter mandatory!!!"; \
		echo "XXXX ex. make env env user=xxxx password=yyyy"; \
		echo; \
		exit 1; \
	fi
	@mkdir -p contests
	@echo "USER=${user}" > .USER
	@echo "PASSWORD=${password}" >> .USER
	cd ../tools && python3 -mvenv venv && . venv/bin/activate && pip install -r requirements.txt

prepare:
	@if [ -z "${contest}" ]; then \
		echo "XXXX contest parameter is mandatory!!!"; \
		echo "XXXX ex. make prepare contest=abc123"; \
		exit 1; \
	fi
	@echo "";
	@echo "++++ prepare envfile for contest [$${contest^^}] ++++"
	@echo "";
	@echo "CONTEST=$${contest^^}" > .CONTEST
	@mkdir -p contests/$${contest^^}

chk:
	@. .USER && if [ -z "$${USER}" ]; then \
		echo "XXX Please exec 'make env user=xxxx password=yyyy' first"; \
		exit 1; \
	fi
	@. .CONTEST && if [ -z "$${CONTEST}" ]; then \
		echo "XXX Please exec 'make prepare contest=yyyy' first"; \
		exit 1; \
	fi
	@. .QUESTION && if [ -z "$${QUESTION}" ]; then \
		echo "XXX Please exec 'make start q=zzzz' first"; \
		exit 1; \
	fi


get:
	@rm -rf samples
	@make chk
	@. .USER && . .CONTEST && . .QUESTION && cd ../tools/ && . venv/bin/activate && python get_samplecases.py -u $${USER} -p $${PASSWORD} -c $${CONTEST} -q $${QUESTION} -d ${PWD}/samples


start:
	@if [ -z "${q}" ]; then \
		echo "***** q is required. Ex. make done q=zzzz"; \
		exit 1; \
	fi
	@echo "Q=${q}" > .QUESTION
	@. .CONTEST && echo "QUESTION=$${CONTEST}_${q}" >> .QUESTION
	@if [ -f main.go ]; then \
		cp main.go main.tmp.go; \
	fi
	rm -f go.mod go.sum
	cp libs/template.go main.go
	go mod init main
	make get


done:
	@make chk
	. .CONTEST && . .QUESTION && mkdir -p contests/$${CONTEST}/$${QUESTION} && cp main.go contests/$${CONTEST}/$${QUESTION}/main.go
	rm -f go.mod go.sum .QUESTION

resume:
	@if [ -z "${q}" ]; then \
		echo "***** q is required. Ex. make resume q=zzzz"; \
		exit 1; \
	fi
	@echo "Q=${q}" > .QUESTION
	@. .CONTEST && echo "QUESTION=$${CONTEST}_${q}" >> .QUESTION
	@if [ -f main.go ]; then \
		cp main.go main.tmp.go; \
	fi
	. .CONTEST && . .QUESTION && cp contests/$${CONTEST}/$${QUESTION}/main.go main.go
	rm -f go.mod go.sum
	go mod init main
	make get


run:
	@go run main.go


test:
	@cd ../tools/ && . venv/bin/activate && python run_code.py -d ${PWD} -t ${case} -a ${answer}


test-contest:
	@for f in samples/in/*; do \
		FILENAME=`basename $${f}`; \
		echo "--------- TESTING: $${FILENAME} ---------"; \
		make test case=$${f} answer=samples/out/$${FILENAME}; \
	done


test-all:
	@make chk
	@. .CONTEST && . .QUESTION && if [ ! -d "../testcase/$${CONTEST}/$${Q}" ]; then \
		echo "***** No test cases for contest $${CONTEST}/${Q} are downloaded. Please download them in advance at the following URL."; \
		echo "https://www.dropbox.com/sh/nx3tnilzqz7df8a/AAAYlTq2tiEHl5hsESw6-yfLa?dl=0"; \
		exit 1; \
	fi
	@rm -f TEST_RESULT
	@. .USER && . .CONTEST && . .QUESTION && for f in ../testcase/$${CONTEST}/$${Q}/in/*; do \
		FILENAME=`basename $${f}`; \
		echo "--------- TESTING: $${FILENAME} ---------"; \
		make test case=$${f} answer=../testcase/$${CONTEST}/$${Q}/out/$${FILENAME} >> TEST_RESULT; \
	done
	@cat TEST_RESULT
	@if [ -n "`grep 'code_result: False' TEST_RESULT`" ]; then \
		echo "!!!!!! Some tests failed. Please check !!!!!!"; \
	fi
