rsync -rzh --progress \
    --exclude 'postgres-data' \
    --exclude '.git' \
    /home/thanhpp/go/src/github.com/thanhpp/zola HauPCDroplet2:/root/thanhpp/