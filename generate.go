package kurirmoo

//go:generate swagger generate server --exclude-main -A kurirmoo-server -t gen -f ./api/swagger.yml  --principal models.Principal
//go:generate swagger -q generate client -A kurirmoo -f api/swagger.yml -c pkg/client -m gen/models --principal models.Principal
//go:generate swagger -q generate client -A otp -f api/otp.yml -c pkg/client/otp_service_client -m gen/models --principal models.Principal
