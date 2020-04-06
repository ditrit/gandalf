//BUILD <br />
go build -tags libsqlite3 <br />
<br />
//CLUSTER <br />
./garcimore cluster init cluster 127.0.0.1:9000 <br />
./garcimore cluster join cluster 127.0.0.1:9001 127.0.0.1:9000 <br />
./garcimore cluster join cluster 127.0.0.1:9002 127.0.0.1:9000 <br />
<br />
//AGGREGATOR <br />
./garcimore aggregator init agg1 titi 127.0.0.1:8000 127.0.0.1:9000 <br />
./garcimore aggregator join agg1 titi 127.0.0.1:8001 127.0.0.1:9000 127.0.0.1:8000 <br />
./garcimore aggregator init agg2 titi 127.0.0.1:8100 127.0.0.1:9000 <br />
./garcimore aggregator join agg2 titi 127.0.0.1:8101 127.0.0.1:9000 127.0.0.1:8100 <br />
<br />
//CONNECTOR <br />
./garcimore connector init con1 titi 127.0.0.1:7000 127.0.0.1:7010 127.0.0.1:8000 <br />
./garcimore connector join con1 titi 127.0.0.1:7001 127.0.0.1:7011 127.0.0.1:8000 127.0.0.1:7000 <br />
./garcimore connector init con2 titi 127.0.0.1:7100 127.0.0.1:7110 127.0.0.1:8100 <br />
./garcimore connector join con2 titi 127.0.0.1:7101 127.0.0.1:7111 127.0.0.1:8100 127.0.0.1:7100 <br />
<br />
//WORKER CMD <br />
./garcimore test send cmd test test <br />
./garcimore test receive cmd test <br />
<br />
//WORKER EVT <br />
./garcimore test send evt test test test <br />
./garcimore test receive evt test test <br />