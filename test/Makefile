
EXTS =  \
	elm		\
	go		\
	js		\
	php		\
	pl		\
	py		\
	rb

PREFIX = test_file

all:
	-for ext in $(EXTS); do (touch $(PREFIX).$$ext); done

clean:
	rm $(PREFIX).*
