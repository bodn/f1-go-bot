# # Import env variables 
# include .env
# export $(shell sed 's/=.*//' .env) # nice? this imports all the

run:
	go run ./main.go