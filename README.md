"# azure-cosmosdb-cassandra-go-sample" 

After setting up this repo on your box, make sure to fetch the cassandra go driver:
go get github.com/gocql/gocql
(It goes without saying, you already need go installed on your box to do this.)

NOTE: For using SSL, make sure to convert the default bc2025.crt to .pem (use openssl):
openssl x509 -inform DER -outform PEM -in bc2025.crt -out bc2025.pem