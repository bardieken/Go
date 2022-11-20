response=$(curl -s "https://01.kood.tech/api/graphql-engine/v1/graphql" --data '{"query":"{user(where:{login:{_eq:\"'alpbal'\"}}){id}}"}' | jq '.data.user[0].id')
echo  $response