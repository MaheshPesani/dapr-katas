# Echo Service

```
@startuml
python_sender -> node_receiver: Hello, World
node_receiver --> python_sender: Received

python_sender -> node_receiver: Request
python_sender <-- node_receiver: Response
@enduml
```
