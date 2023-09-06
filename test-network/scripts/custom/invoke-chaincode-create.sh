peer chaincode invoke -C channel1 -n diploma-basic \
    -o localhost:7050 \
    --ordererTLSHostnameOverride orderer.example.com \
    --tls \
    --cafile $ORDERER_TLS \
    --peerAddresses localhost:7051 --tlsRootCertFiles $PEER1_TLS \
    --peerAddresses localhost:9051 --tlsRootCertFiles $PEER2_TLS \
    -c '{"Args":["CreateCertificate", "D1AE259E5D4C8FF00D641E95CA553A9140545AF2DDAB4C83C87A712BCBAE5E2DN0FKQXAxL1VTeks0RXljOWlBWDNYT05zRURnQVp5RVEzd25tQnJKOVdFbXhheG0b", "Assertion", "https://diplom.mn/must/badges/1", "BadgeClass", "Software Engineer", "Software Engineer Degree", "Has satisfactorily completed all degree requirements", "name", "Bolortoli Munkhsaikhan", "https://diplom.mn/issuer/must", "Profile", "Mongolian University of Science and Technology", "https://must.edu.mn", "contact@must.edu.mn", "diplom.mn"]}'