header:
	cat doc/head.md > README.md
	godoc2md github.com/jspc/routes >> README.md
