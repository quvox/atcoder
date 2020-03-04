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

get:
	rm -rf samples
	. .USER && . .CONTEST && cd ../tools/ && . venv/bin/activate && python get_samplecases.py -u $${USER} -p $${PASSWORD} -c $${CONTEST} -q $${q} -d ${PWD}/samples


start:
	@if [ -z "${q}" ]; then \
		echo "***** q is required. Ex. make done q=a"; \
		exit 1; \
	fi
	echo ":::::"
	@if [ -f main.go ]; then \
		cp main.go main.tmp.go; \
	fi
	echo "222"
	rm -f go.mod go.sum
	cp libs/template.go main.go
	go mod init main
	make get q=${q}


done:
	@if [ -z "${q}" ]; then \
		echo "***** q is required. Ex. make done q=a"; \
		exit 1; \
	fi
	. .CONTEST && mkdir -p contests/$${CONTEST}/$${q^^} && cp main.go contests/$${CONTEST}/$${q^^}/main.go
	rm -f go.mod go.sum

resume:
	@if [ -z "${q}" ]; then \
		echo "***** q is required. Ex. make resume q=a"; \
		exit 1; \
	fi
	@if [ -f main.go ]; then \
		cp main.go main.tmp.go; \
	fi
	. .CONTEST && cp main.go contests/$${CONTEST}/$${q^^}/main.go main.go
	rm -f go.mod go.sum


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
	@if [ -z "${q}" ]; then \
		echo "***** q is required. Ex. make resume q=a"; \
		exit 1; \
	fi
	@. .CONTEST && if [ ! -d "../testcase/$${CONTEST}/${q}" ]; then \
		echo "***** No test cases for contest $${CONTEST}/${q} are downloaded. Please download them in advance at the following URL."; \
		echo "https://www.dropbox.com/sh/nx3tnilzqz7df8a/AAAYlTq2tiEHl5hsESw6-yfLa?dl=0"; \
		exit 1; \
	fi
	@rm -f TEST_RESULT
	@. .USER && . .CONTEST && for f in ../testcase/$${CONTEST}/${q}/in/*; do \
		FILENAME=`basename $${f}`; \
		echo "--------- TESTING: $${FILENAME} ---------"; \
		make test case=$${f} answer=../testcase/$${CONTEST}/${q}/out/$${FILENAME} >> TEST_RESULT; \
	done
	@cat TEST_RESULT
	@if [ -n "`grep 'code_result: False' TEST_RESULT`" ]; then \
		echo "!!!!!! Some tests failed. Please check !!!!!!"; \
	fi
