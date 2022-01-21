rsync -rzh --progress \
    --exclude 'postgres-data' \
    --exclude '.git' \
    --exclude 'node_modules' \
    /home/thanhpp/go/src/github.com/thanhpp/zola HauPCDroplet2:/root/thanhpp/