peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID channel1 --name diploma-basic --version 1.0 --sequence 1 --tls --cafile $ORDERER_TLS --peerAddresses localhost:7051 --tlsRootCertFiles $PEER1_TLS --peerAddresses localhost:9051 --tlsRootCertFiles $PEER2_TLS