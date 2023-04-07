#!bin/bash

nohup ./bee dev > bee.log 2>&1 &  # this will directly start a bee node in dev mode(the uploaded file will be saved in dev)

sleep 5

curl http://localhost:1633 # testing whether the bee service has started

curl -s -XPOST http://localhost:1635/stamps/10000000/20  # fund this node, to make this node able to upload files

sleep 5

# it will return like this :  {"batchID":"c3f29badfb4f8decc3042d91e82ce5f698746a14395011da9a14c9ddd0112c2b","txHash":"0x0000000000000000000000000000000000000000000000000000000000000000"}


#upload a file:
# curl --data-binary @<filename e.g. bee.jpg>  -H "Swarm-Postage-Batch-Id: c3f29badfb4f8decc3042d91e82ce5f698746a14395011da9a14c9ddd0112c2b" "http://localhost:1633/bzz?name=<filename e.g. bee.jpg>"
#or upload a non-binary file:
# curl --data @<filename e.g. doc.go>  -H "Swarm-Postage-Batch-Id: a0a7776d0f959967e61c139e7d08746badcc0407fa7a7423e205388edb2783e0" "http://localhost:1633/bzz?name=<filename e.g. doc.go>"

# it will return like this {"reference":"e4f2e31d9b86c3f68c34c362c7f9ad8a0e5e18523aca56d13544fb6420c90fe8"}

#download a file:
#curl -OJ http://localhost:1633/bzz/e4f2e31d9b86c3f68c34c362c7f9ad8a0e5e18523aca56d13544fb6420c90fe8/

