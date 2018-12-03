curl -X POST \
  http://localhost:8080/cards \
  -H 'content-type: application/json' \
  -d '{
	"name": "12 factors app",
	"description": "1. Codebase: \\n, 2. Dependencies : \\n, 3. Config: \\n, 4. Backing Service: \\n, 5. Build Release Run: \\n, 6. Processes: \\n, 7. Port Binding: \\n, 8. Concurrency: \\n, 9. Disposability: \\n, 10. Dev/Prod parity: \\n, 11. Logs: \\n, 12. Admin Processes: \\n",
	"category": "architecture",
	"owner": "19d62e1e-f651-11e8-8eb2-f2801f1b9fd1"
}'


curl -X POST \
  http://localhost:8080/cards \
  -H 'content-type: application/json' \
  -d '{
	"name": "Differences between Contiuous Intergartion, Delivery and Deployment",
	"description": "1. Continous Integration: Developers participating continuous integration merge their changes back to the main branch as often as possible. By doing so, you avoid the integration hell that usually happens when people wait for release day to merge their changes into the release branch.\\n 2. Continuous Delivery: is an extension of continuous integration to make sure that you can release new changes to your customers quicly in a ...",
	"category": "architecture",
	"owner": "19d62e1e-f651-11e8-8eb2-f2801f1b9fd1"
}'


curl -X POST \
  http://localhost:8080/cards \
  -H 'content-type: application/json' \
  -d '{
	"name": "12 factors app",
	"description": "An Heuristic Exception refers to a transaction participant decision to unilaterally take some action without the consensus of the transaction manager, usually as a result of some kind of catastrophic failure between the participant and the transation manager.",
	"category": "architecture",
	"owner": "19d62e1e-f651-11e8-8eb2-f2801f1b9fd1"
}'

curl -X POST \
  http://localhost:8080/cards \
  -H 'content-type: application/json' \
  -d '{
	"name": "Shared Nothing Architecture",
	"description": "A Shared Nothing Architecture (SN) is a distributed computing approach in which each node is independant and self-sufficient, and there is no single point of contention required across the system.",
	"category": "architecture",
	"owner": "19d62e1e-f651-11e8-8eb2-f2801f1b9fd1"
}'