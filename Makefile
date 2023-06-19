GODBL_FOLDER = godbl/*
SEED_FOLDER = seed/*
CONFIG_FOLDER = ./config

test: $(GODBL_FOLDER)
	CONFIG_FOLDER=../$(CONFIG_FOLDER) go test ./godbl -v

testmongo: $(GODBL_FOLDER)
	CONFIG_FOLDER=../$(CONFIG_FOLDER) go test ./godbl/adapters/mongo -v

build: $(GODBL_FOLDER)
	CONFIG_FOLDER=../$(CONFIG_FOLDER) go build -a -o bin/godbl ./godbl
