
EXTS =  \
	go		\
	pl		\
	php		\
	py		\
	rb		\
	js

PREFIX = test_file

all:
	-for ext in $(EXTS); do (touch $(PREFIX).$$ext); done

clean:
	rm $(PREFIX).*
